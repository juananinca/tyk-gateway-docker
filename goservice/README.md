# Go Service

Servicio HTTP simple en Go con dos endpoints.

## Endpoints

- `GET /hello` - Retorna un mensaje de bienvenida con timestamp
- `GET /status` - Retorna el estado del servicio, versión y uptime

## Ejecutar

### Localmente
```bash
go run main.go
```

### Con Docker
```bash
# Construir la imagen
docker build -t goservice:latest .

# Ejecutar el contenedor
docker run -p 8080:8080 goservice:latest
```

El servicio estará disponible en `http://localhost:8080`

## Ejemplos

```bash
# Endpoint Hello
curl http://localhost:8080/hello

# Endpoint Status
curl http://localhost:8080/status
```
