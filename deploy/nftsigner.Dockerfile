# BUILDER Image. Used to download all dependenices, etc
FROM golang:1.17.4-alpine3.15 as nftsigner_builder

WORKDIR /app
# Copy all files to root derictory
COPY . .
# Collect and download dependances
RUN go mod vendor
# Building application
RUN CGO_ENABLED=0 go build -o main ./cmd/currencysigner/main.go

# Result image
FROM alpine:3.15.4

WORKDIR /app
# Volume directory
ARG APP_DATA_DIR=/data

RUN mkdir -p ${APP_DATA_DIR}
# Copy executable file (builded application) from builder to root directory
COPY --from=cnftsigner_builder /app/main .
# Criating volume
VOLUME ["${APP_DATA_DIR}"]
# Builded application running with config file as argument
ENTRYPOINT ["/main", "run", "--config=./configs"]