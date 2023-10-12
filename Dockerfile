FROM golang:1.21-bookworm AS base

WORKDIR /app

# Install dependencies
RUN apt-get update && apt-get install -y curl 
COPY . .

RUN go build -o main cmd/SeeCV/main.go
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.16.2/migrate.linux-amd64.tar.gz | tar xvz


FROM debian:bookworm-slim
WORKDIR /app
RUN apt-get update && apt-get install -y poppler-utils tidy wv unrtf netcat-openbsd  ca-certificates && apt-get clean && rm -rf /var/lib/apt/lists/*

COPY --from=base /app/main .
COPY --from=base /app/migrate .
COPY config.toml .
COPY wait-for.sh .
COPY start.sh .
COPY internal/db/migration ./migration


EXPOSE 8080

CMD ["/app/main"]
ENTRYPOINT [ "/app/start.sh" ]