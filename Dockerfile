FROM golang:alpine3.20 AS build 
RUN apk update && apk add --no-cache git

WORKDIR /app
COPY . .

RUN go build -v -o fc-go-challenge .

FROM scratch
COPY --from=build /app/fc-go-challenge ./fc-go-challenge

EXPOSE 8080
ENTRYPOINT ["/app/fc-go-challenge"]
