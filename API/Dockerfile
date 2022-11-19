FROM golang:1.19-alpine as builder
WORKDIR /build/
COPY . .
RUN apk add --no-cache make gcc musl-dev git
RUN go build server.go

FROM alpine:latest
LABEL author=Nathan13888
WORKDIR /app
RUN apk add --no-cache tzdata
RUN cp /usr/share/zoneinfo/America/Toronto /etc/localtime

ENV USER=escapar
ENV UID=12345
ENV GID=23456
RUN addgroup --gid $GID $USER
RUN adduser \
    --disabled-password \
    --gecos "" \
    --home "$(pwd)" \
    --ingroup "$USER" \
    --no-create-home \
    --uid "$UID" \
    "$USER"

RUN chown -R escapar:escapar .
RUN chmod -R 755 .
USER escapar

COPY --from=builder --chown=escapar /build/bin/escapar /app/escapar

VOLUME /app/.escapar
EXPOSE 3000

ENTRYPOINT [ "/app/escapar" ]

