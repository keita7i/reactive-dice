FROM golang:1.14 AS builder

WORKDIR /build

COPY . /build

RUN go build -o reactive-dice --ldflags "-linkmode 'external' -extldflags '-static'" .

FROM scratch

WORKDIR /bin

COPY --from=builder /etc/ssl/certs /etc/ssl/certs
COPY --from=builder /build/reactive-dice /bin/reactive-dice
COPY ./assets /usr/share/reactive-dice/assets

EXPOSE 80

ENTRYPOINT [ "/bin/reactive-dice" ]
