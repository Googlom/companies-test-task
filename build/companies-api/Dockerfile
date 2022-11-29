FROM golang:1.19-alpine3.16 as BUILD-ENV

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o main ./cmd/companies-api
RUN go build -o migrate ./cmd/migrate


FROM alpine:3.16
WORKDIR /app
COPY --from=BUILD-ENV /app/main /app/main
COPY --from=BUILD-ENV /app/migrate /app/migrate
COPY db_migration/ ./db_migration/

ENV GIN_MODE=release
CMD ["./main"]
