FROM golang:1.19-alpine
WORKDIR /app
COPY . .
RUN go mod download && go mod tidy
RUN go build -o main ./src/main.go
CMD ["./main"]