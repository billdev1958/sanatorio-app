FROM golang:1.24.0 AS builder

# Establecer el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copiar los archivos go.mod y go.sum primero para aprovechar la caché de Docker
COPY go.mod go.sum ./

# Descargar las dependencias
RUN go mod download

# Copiar el código fuente al contenedor
COPY . .

# Compilar el binario
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o sanatorio-back ./cmd/main.go

# Etapa 2: Creación de la imagen ligera utilizando Distroless
FROM gcr.io/distroless/base-debian12

# Establecer el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copiar el binario desde la etapa de construcción
COPY --from=builder /app/sanatorio-back /app/sanatorio-back

COPY internal/domain/email/plantillaConfirmacion.html email/
COPY internal/domain/email/logo.png email/
COPY internal/domain/email/logo_cms.png email/

# Exponer el puerto en el que corre tu aplicación (ajusta según tu configuración)
EXPOSE 8080

# Comando para ejecutar el binario
CMD ["/app/sanatorio-back"]
