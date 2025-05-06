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

## Conexión a través de un ORM PostgreSQL
   ORM simplifica el acceso a los datos al permitir a los desarrolladores interactuar con bases de datos mediante objetos y métodos de su lenguaje de programación preferido, en lugar de escribir consultas SQL.
   - GORM (Golang): [https://gorm.io/docs/](https://gorm.io/docs/)
   - Hibernate (Java): [hibernate.org/orm/documentation/](hibernate.org/orm/documentation/)
   - Entity Framework (Microsoft .NET): [learn.microsoft.com/en-us/ef/](https://learn.microsoft.com/es-mx/ef/)
   - SQLAlchemy (Python): [www.sqlalchemy.org/library.html#tutorials](https://www.sqlalchemy.org/library.html#tutorials)
   - Django ORM (Python/Django): [docs.djangoproject.com/en/stable/topics/db/](https://docs.djangoproject.com/en/5.2/topics/db/)
   - Sequelize (Node.js): [sequelize.org/docs/v6/](https://sequelize.org/docs/v6/)
   - TypeORM (TypeScript/JavaScript): [typeorm.io/](https://typeorm.io/)
   - Doctrine (PHP): [www.doctrine-project.org/projects/orm.html](https://www.doctrine-project.org/projects/orm.html)  
   ```
   go get gorm.io/gorm
   go get gorm.io/driver/postgres
   ```
   Para obtener los valores de variables de entorno para el proyecto:
   ```
   go get github.com/spf13/viper
   ```
    