FROM golang:bookworm AS build

WORKDIR /build

COPY . .

RUN go build 

FROM alpine:3.21 AS final

COPY --from=BUILD /build/mconfig /usr/local/bin

ENTRYPOINT [ "mconfig" ]


