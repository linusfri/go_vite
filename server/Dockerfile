FROM golang:alpine

WORKDIR /app

ENV WATCHEXEC_VERSION 1.14.0
RUN apk update && apk add curl && curl -L https://github.com/watchexec/watchexec/releases/download/${WATCHEXEC_VERSION}/watchexec-${WATCHEXEC_VERSION}-x86_64-unknown-linux-musl.tar.xz | tar xvJ && \
    mv watchexec-${WATCHEXEC_VERSION}-x86_64-unknown-linux-musl/watchexec /bin/watchexec

COPY go.mod .
RUN go mod download

COPY . .

EXPOSE 8080