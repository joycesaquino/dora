FROM golang:1.17-alpine3.14 as builder

ADD . /app

WORKDIR /app

RUN CGO_ENABLED=0 GO111MODULE=on GOOS=linux GOARCH=amd64 go build -o dora-api ./cmd/dora-api

FROM alpine:3.14

COPY --from=builder /app/dora-api /bin/dora-api

RUN chmod +x /bin/dora-api

CMD [ "start" ]

ENTRYPOINT [ "dora-api" ]