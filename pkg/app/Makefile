MAIN = server.go

default:
	go run ${MAIN} -p 8000

run:
	docker-compose build
	docker-compose up -d

stop:
	docker-compose down

test: default


