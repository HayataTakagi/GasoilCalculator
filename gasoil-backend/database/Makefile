export GO111MODULE := off

include ../.env

DATA_SOURCE:=root:$(DB_ROOT_PASS)@tcp(127.0.0.1:$(DB_PORT))/$(DB_NAME)

install:
	which goose || GO111MODULE=off go get -u github.com/pressly/goose/cmd/goose

docker-compose/up:
	docker-compose up -d

docker-compose/down:
	docker-compose down

.PHONY: mysql
mysql:
	mysql -u root -h localhost --protocol tcp -p $(DB_NAME) -P $(DB_PORT)

migrate/init:
	mysql -u root -h localhost --protocol tcp -e "create database \`$(DB_NAME)\`" -p  -P $(DB_PORT)

migrate/status:
	@goose -dir migrations mysql "$(DATA_SOURCE)" status

migrate/up:
	@goose -dir migrations mysql "$(DATA_SOURCE)" up

__migrate/down:
	@goose -dir migrations mysql "$(DATA_SOURCE)" down

init: install docker-compose/up migrate/init migrate/up
