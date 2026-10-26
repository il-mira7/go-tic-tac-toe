# сборка: модуль лежит в src/
FROM golang:1.25-alpine AS build
WORKDIR /build
COPY src/go.mod src/go.sum ./
RUN go mod download
COPY src/ .
RUN CGO_ENABLED=0 go build -o /bin/tic-tac-toe ./cmd/app

# запуск
FROM alpine:3.20
RUN adduser -D -u 10001 app
COPY --from=build /bin/tic-tac-toe /usr/local/bin/tic-tac-toe
USER app
EXPOSE 8080
ENTRYPOINT ["tic-tac-toe"]
