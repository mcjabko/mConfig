FROM golang:bullseye as BUILD

COPY src /src

WORKDIR /src

RUN go build 

FROM alpine:3.18 AS FINAL

COPY --from=BUILD /src/mconfig /usr/local/bin

ENTRYPOINT [ "mconfig" ]


