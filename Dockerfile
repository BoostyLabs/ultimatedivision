# BUILDER Image. Used to download all dependenices, etc
FROM golang:1.14.15-alpine3.13 as builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/currencysigner/main.go

# Result image
FROM alpine:3.13

ARG APP_DATA_DIR=/app/data

RUN mkdir -p ${APP_DATA_DIR}

COPY --from=builder /app/config_currency_signer.json .
COPY --from=builder /app/main .

# NO NEEDED PORTS
#EXPOSE 8187 8188

VOLUME ["${APP_DATA_DIR}"]

ENTRYPOINT ["/main", "run"]

