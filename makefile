DB_USER := $(shell grep DB_USER .env | cut -d '=' -f 2)
DB_PASSWORD := $(shell grep DB_PASSWORD .env | cut -d '=' -f 2)
DB_HOST := $(shell grep DB_HOST .env | cut -d '=' -f 2)
DB_PORT := $(shell grep DB_PORT .env | cut -d '=' -f 2)
DB_NAME := $(shell grep DB_NAME .env | cut -d '=' -f 2)


requirements-up:
	docker-compose -f docker-compose.yml up -d


migration-files:
	@read -p "Enter the title: " title; \
	migrate create -ext sql -dir ./db/migrations -seq $$title

migrateup:
	migrate -path ./db/migrations -database "postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable" -verbose up

migratedown:
	migrate -path ./db/migrations -database "postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable" -verbose down
