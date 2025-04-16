FROM alpine:3.21.3
RUN apk add --no-cache go git
WORKDIR /app
COPY . .
RUN go build -o main main.go
EXPOSE 8001
CMD ["./main"]