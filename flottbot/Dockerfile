FROM golang:1.13-alpine

RUN apk add libc6-compat
RUN wget https://github.com/target/flottbot/releases/download/0.3.1/flottbot-linux-amd64.tgz
RUN tar -zxvf flottbot-linux-amd64.tgz

COPY / .

CMD ["./flottbot"]
