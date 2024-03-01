# FROM golang:latest

# WORKDIR /app

# RUN mkdir "/build"

# COPY . .

# RUN go mod download

# RUN go build -o /build/app

# EXPOSE 8080

# CMD ["go","run","main.go"]




FROM golang:latest

ENV PROJECT_DIR=/app \
    GO111MODULE=on \
    CGO_ENABLED=0

WORKDIR /app

RUN mkdir "/build"

COPY . .

RUN go get github.com/githubnemo/CompileDaemon

RUN go install github.com/githubnemo/CompileDaemon

EXPOSE 8080

ENTRYPOINT CompileDaemon -build="go build -buildvcs=false -o /build/app" -command="/build/app"