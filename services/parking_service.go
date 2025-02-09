package services

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"query-availability-microservice/config"

	"github.com/joho/godotenv"
)

// Load environment variables
func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("⚠️ Warning: No .env file found, using system environment variables")
	}
}

// checkParkingLotExists verifies if the parking lot exists by calling the external API
func checkParkingLotExists(id int) (bool, error) {
	loadEnv() // Ensure environment variables are loaded

	parkingServiceBaseURL := os.Getenv("PARKINGLOT_SERVICE_URL")
	if parkingServiceBaseURL == "" {
		return false, errors.New("missing PARKINGLOT_SERVICE_URL in environment variables")
	}

	parkingServiceURL := fmt.Sprintf("%s/%d", parkingServiceBaseURL, id)

	resp, err := http.Get(parkingServiceURL)
	if err != nil {
		return false, fmt.Errorf("error contacting parking lot service: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return false, nil // The parking lot does not exist
	}

	if resp.StatusCode != http.StatusOK {
		return false, fmt.Errorf("unexpected response from parking lot service: %d", resp.StatusCode)
	}

	// Read and parse response to verify it's a valid JSON
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false, fmt.Errorf("error reading response from parking lot service: %v", err)
	}

	var response map[string]interface{}
	if err := json.Unmarshal(body, &response); err != nil {
		return false, fmt.Errorf("invalid response format from parking lot service")
	}

	return true, nil
}

// FetchParkingCapacity retrieves both capacity and total_space for a parking lot
func FetchParkingCapacity(id int) (int, int, error) {
	// Step 1: Verify if the parking lot exists
	exists, err := checkParkingLotExists(id)
	if err != nil {
		return 0, 0, err
	}
	if !exists {
		return 0, 0, errors.New("parking lot not found")
	}

	// Step 2: Fetch data from database
	var capacity, totalSpace int
	err = config.DB.QueryRow(
		"SELECT capacity, total_space FROM ParkingLot WHERE id = ?", id,
	).Scan(&capacity, &totalSpace)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, 0, errors.New("parking lot data not found in database")
		}
		return 0, 0, fmt.Errorf("database error: %v", err)
	}

	return capacity, totalSpace, nil
}
