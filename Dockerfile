#Stage 1: Build the Golang application
FROM golang:1.20 AS builder
WORKDIR /app
COPY .  . 
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main . 

#Stage2: Create a minimial runtime 
FROM alpine
WORKDIR /root/
COPY --from=builder /app/main .
CMD ["./main"]
