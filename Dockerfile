# Usa una imagen base de Ubuntu
FROM ubuntu:latest

# Instalar Go, nmap y dependencias necesarias
RUN apt-get update && apt-get install -y \
    golang-go \
    nmap \
    && rm -rf /var/lib/apt/lists/*

# Crear y establecer el directorio de trabajo
WORKDIR /app

# Copiar el código fuente de tu aplicación Go al contenedor
COPY . .

# Exponer el puerto en el que tu aplicación Go escuchará (si es necesario)
EXPOSE 8080

# Compilar la aplicación Go
RUN go build -o main .

# Ejecutar la aplicación Go
CMD ["./main"]
