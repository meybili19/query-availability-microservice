# Usa la imagen oficial de Go para compilar la aplicación
FROM golang:1.21 AS builder

# Establece el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copia los archivos del proyecto al contenedor
COPY go.mod go.sum ./
RUN go mod tidy

COPY . .

# Compila la aplicación en un ejecutable binario
RUN go build -o main .

# Crea la imagen final mínima para ejecutar el servicio
FROM debian:bullseye-slim

# Establece el directorio de trabajo en la imagen final
WORKDIR /app

# Copia el binario compilado desde la etapa anterior
COPY --from=builder /app/main .

# Copia el archivo .env
COPY .env .env

# Expone el puerto en el que corre el servicio
EXPOSE 6005

# Comando para ejecutar la aplicación
CMD ["./main"]
