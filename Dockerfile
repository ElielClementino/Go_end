FROM golang:1.18

WORKDIR App

COPY go.mod go.sum ./

RUN go mod download

COPY . ./

RUN GOOS=linux go build -o Task

EXPOSE 8080

CMD ["./Task"]