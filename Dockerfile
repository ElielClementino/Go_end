FROM golang:1.18

WORKDIR App

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . ./


RUN go build -o Task

EXPOSE 8080

CMD ["./Task"]