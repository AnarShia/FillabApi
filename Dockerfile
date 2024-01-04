FROM golang:1.21.5

WORKDIR /usr/src/app

RUN go install github.com/cosmtrek/air@latest

COPY . .
RUN go mod tidy
RUN go build -o main .

EXPOSE 8080

CMD ["./main"]


