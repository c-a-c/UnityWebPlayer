# Base of Docker images
FROM golang:latest

# Work directory
WORKDIR /go/src/

# Change Current Directory to Source
# RUN cd ../pkg

# Copy host directories to Docker
COPY . .

# Exec DB Migration
# WORKDIR /go/src/app/db
# RUN ls
# RUN bash migrate.sh

# Execute server.go ( main )
RUN cd app && \
    go build -o ../main server.go
CMD [ "./main" ]

