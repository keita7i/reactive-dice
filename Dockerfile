FROM node:12 as node-builder

WORKDIR /build

COPY ./react /build

RUN npm install . && npm run build

FROM golang:1.14 AS golang-builder

WORKDIR /build

COPY . /build

RUN go build -o reactive-dice --ldflags "-linkmode 'external' -extldflags '-static'" .

FROM scratch

WORKDIR /bin

COPY --from=node-builder /build/dist /usr/share/reactive-dice/assets
COPY --from=golang-builder /etc/ssl/certs /etc/ssl/certs
COPY --from=golang-builder /build/reactive-dice /bin/reactive-dice

EXPOSE 80

ENTRYPOINT [ "/bin/reactive-dice" ]
