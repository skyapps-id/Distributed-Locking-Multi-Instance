# Build
FROM golang:alpine AS build

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o binary main.go 


# Final
FROM alpine:latest

WORKDIR /app

COPY --from=build /app/binary /app/binary

ENTRYPOINT ["/app/binary"]