# SYSACAD en Gin Framework
https://gin-gonic.com/es/docs/quickstart/

## Estructura del Proyecto

```
gin-sysacad
├── cmd
│   └── main.go          
├── internal
│   ├── handlers
│   │   └── handlers.go  
│   ├── middleware
│   │   └── middleware.go 
│   └── models
│       └── models.go    
├── config
│   └── config.go        
├── go.mod                
├── go.sum                
└── README.md             
```
## Requisitos previos
1. Tener instalado Golang version 1.24 https://go.dev/doc/install

## Crear Proyecto Nuevo
```
go mod init https://github.com/usuario/proyecto
```

## Instrucciones para el Proyecto

1. **Clone the repository:**
   ```
   git clone https://github.com/umpprats/gin-sysacad
   cd gin-sysacad
   ```

2. **Instalar dependencias:**
   ```
   go mod tidy
   ```
3. **Ejecutar Test**
   ```
   # Ejecuta todos los test en un directorio/carpeta
   go test
   # Ejecuta un test específico
   go test -run UniversidadTest
   ```
4. **Ejecutar aplicación:**
   ```
   go run main.go
   ```
