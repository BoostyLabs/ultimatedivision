# BUILDER Image. Used to download all dependenices, etc
FROM golang:1.14.15-alpine3.13 as builder

WORKDIR /app

COPY . .

RUN go mod download


RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/ultimatedivision/main.go

# Result image
FROM alpine:3.13

ARG APP_DATA_DIR=/app/data

RUN mkdir -p ${APP_DATA_DIR}

COPY --from=builder /app/main .
COPY --from=builder /app/web/admin ./web/admin
COPY --from=builder /app/web/console/dist ./web/console/dist

EXPOSE 8087 8088

VOLUME ["${APP_DATA_DIR}"]

ENTRYPOINT ["/main", "run", "--config=./config"]