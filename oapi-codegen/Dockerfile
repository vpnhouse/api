FROM amd64/golang:1.19.5-alpine3.17

RUN go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@v1.9.1
RUN ls /go/bin

WORKDIR /app

COPY ./update.sh /app/update.sh

RUN chmod +x update.sh
