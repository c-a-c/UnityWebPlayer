# Base of Docker images
FROM golang:latest

# Work directory
WORKDIR /go/src/

# Copy host directories to Docker
COPY . .

# Gopath
# ENV GOPATH $GOPATH:$HOME/go

# Download packages
# RUN go get github.com/jinzhu/gorm
# RUN go get github.com/gin-gonic/gin

# Execute server.go ( main )
# CMD [ "cd app/server" ]

RUN go build -o main server.go
CMD [ "./main" ]

