FROM golang:alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o /docker-go-test cmd/main.go

EXPOSE 8080

CMD [ "/docker-go-test" ]