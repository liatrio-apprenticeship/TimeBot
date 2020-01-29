FROM golang:1.13-alpine

RUN apk add libc6-compat
RUN wget https://github.com/target/flottbot/releases/download/0.3.1/flottbot-linux-amd64.tgz
RUN tar -zxvf flottbot-linux-amd64.tgz
RUN apk update && apk add bash
RUN apk --update add git less openssh && \
    rm -rf /var/lib/apt/lists/* && \
    rm /var/cache/apk/*
RUN go get go.mongodb.org/mongo-driver; exit 0

COPY / .

CMD ["./flottbot"]
