FROM golang:1.11.0-alpine3.8 as builder

WORKDIR /go-grace/

COPY . ./

RUN go build -o app

FROM alpine:3.8

WORKDIR /root/

# If the volume is not specified, mounting it in docker-compose will not
# automatically link the directory together.
VOLUME /root/data

COPY --from=builder /go-grace/app .

CMD ["./app"]
