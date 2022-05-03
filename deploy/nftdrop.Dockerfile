# BUILDER Image. Used to download all dependenices, etc
FROM golang:1.17.4-alpine3.15 as nftdrop_builder

WORKDIR /app
# Copy all files to root derictory
COPY . .
# Download dependances
RUN go mod download
# Building application
RUN GOOS=linux go build -o main ./cmd/nftdrop/main.go

# Result image
FROM alpine:3.15.4

WORKDIR /app
# Volume directory
ARG APP_DATA_DIR=/data

RUN mkdir -p ${APP_DATA_DIR}
# Copy executable file (builded application) from builder to root directory
COPY --from=nftdrop_builder /app/main .
# Criating volume
VOLUME ["${APP_DATA_DIR}"]
# Builded application running with config file as argument
ENTRYPOINT ["/main", "run", "--config=./configs"]
