## Build Stage ##
# First pull Golang image
FROM golang:1.23-alpine as build-env

ENV APP_NAME go-user-service
ENV CMD_PATH src/cmd/main.go

COPY . $GOPATH/src/$APP_NAME
WORKDIR $GOPATH/src/$APP_NAME

# Build application
RUN CGO_ENABLED=0 go build -v -o /$APP_NAME $GOPATH/src/$APP_NAME/$CMD_PATH

## Run Stage ##
FROM alpine:3.14

ENV APP_NAME go-user-service
COPY --from=build-env /$APP_NAME .

EXPOSE 8083

CMD ./$APP_NAME