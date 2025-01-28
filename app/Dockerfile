FROM golang:1.9.6 AS builder 

WORKDIR /app
COPY main.go .
COPY go.mod .

RUN CGO_ENABLED=0 GOOS=linux go build -o go-challenge .

FROM scratch
COPY --from=builder /app/go-challenge ./go-challenge

ENTRYPOINT [ "./go-challenge" ]
