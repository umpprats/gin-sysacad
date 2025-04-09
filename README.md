# SYSACAD en Gin Framework
https://gin-gonic.com/es/docs/quickstart/

## Structura del Proyecto

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
1. Tener instalado Golang version 1.24

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

3. **Ejecutar aplicacion:**
   ```
   go run cmd/main.go
   ```
