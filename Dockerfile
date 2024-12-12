# Usa una imagen base de Go
FROM golang:1.20-alpine

# Establecer el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copiar los archivos del proyecto al contenedor
COPY . .

# Instalar dependencias de Go (como "gorilla/websocket")
RUN go mod tidy

# Compilar la aplicación Go
RUN go build -o websocket-server .

# Exponer el puerto en el que el servidor WebSocket estará escuchando
EXPOSE 8080

# Comando para ejecutar el servidor cuando el contenedor se inicie
CMD ["./websocket-server"]
