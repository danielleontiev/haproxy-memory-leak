FROM golang:1.17.2

COPY . /app
WORKDIR /app

RUN go build -o app main.go

CMD ./app
