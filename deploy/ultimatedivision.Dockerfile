# BUILDER Image. Used to download all dependenices, etc
FROM golang:1.17.4-alpine3.15 as ultimatedivision_builder

WORKDIR /app
# Copy all files to root directory
COPY . .
# Collect and download dependances
RUN go mod vendor
# Building application
RUN CGO_ENABLED=0 go build -mod vendor -o main ./cmd/ultimatedivision/main.go

# Result image
FROM alpine:3.15.4

# Volume directorys
ARG APP_DATA_DIR=/data
ARG APP_CONFIGS=/config

RUN mkdir -p ${APP_CONFIGS}
RUN mkdir -p ${APP_DATA_DIR}
# Criating volumes
VOLUME ["${APP_CONFIGS}", "${APP_DATA_DIR}"]

# Copy executable file (builded application) from builder to root directory
COPY --from=ultimatedivision_builder /app/main .
COPY --from=ultimatedivision_builder /app/web/admin ./web/admin
COPY --from=ultimatedivision_builder /app/web/console/dist ./web/console/dist
# Builded application running with config directory as argument
ENTRYPOINT ["/main", "run", "--config='${APP_CONFIGS}'"]
# ports open
EXPOSE 8080 8081