FROM golang:bookworm as BUILD

WORKDIR /build

COPY . .

RUN go build 

FROM alpine:3.21 AS FINAL

COPY --from=BUILD /src/mconfig /usr/local/bin

ENTRYPOINT [ "mconfig" ]


