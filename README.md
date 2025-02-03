### Ejecución product-api-meli


    
1. ejecutar el comando para instalar dependencias
    
    ```bash
    go mod tidy 
    ```
    
2. se debe crear un archivo .env en la raiz del proyecto con las siguientes variables 
    
    ```bash
    DB_HOST=localhost
    DB_USER=postgresql
    DB_PASS=root
    DB_NAME=products_db
    SSL_MODE=disable
    
    DB_PORT=5432
    KAFKA_BROKERS=localhost:9092
    KAFKA_RETRY=5
    
    REDIS_HOST=localhost
    REDIS_PORT=6379
    KAFKA_TOPIC=price
    ```
    
3. iniciar docker, puede ser docker desktop o simplemente estar ejecutando docker en tu maquina
4. ejecutar el siguiente comando para levantar los servicios de docker necesarios
    
    ```bash
    make requirements-up
    ```
    
5. luego ejecutar las migraciones de base de datos con el siguiente comando 
    
    ```bash
    make migrateup
    ```
    
6. y por ultimo, ejecutar el comando que correra la api de go
    
    ```bash
    go run cmd/main.go
    ```
