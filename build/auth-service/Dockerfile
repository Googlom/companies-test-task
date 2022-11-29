FROM golang:1.19-alpine3.16 as BUILD-ENV

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o main ./cmd/auth-service


FROM alpine:3.16
WORKDIR /app
COPY --from=BUILD-ENV /app/main /app/main

EXPOSE 8090
CMD ["./main"]
