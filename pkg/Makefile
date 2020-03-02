MAIN = server.go

default:
	cd app && \
	go build -o ../main ${MAIN}
	./main -p 8000

run:
	docker-compose build
	docker-compose up -d

build:run

stop:
	docker-compose down

test: default


