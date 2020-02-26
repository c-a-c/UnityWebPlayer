MAIN = server.go

default:
	go run ${MAIN} -p 8000

rd:
	docker-compose build
	docker-compose up -d

sd:
	docker-compose down

test: default


