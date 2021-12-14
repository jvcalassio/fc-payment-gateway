FROM golang:1.17

WORKDIR /go/src

RUN curl https://raw.githubusercontent.com/eficode/wait-for/v2.1.3/wait-for --output /usr/bin/wait-for && \
    chmod +x /usr/bin/wait-for

RUN apt-get update && apt-get install build-essential librdkafka-dev netcat -y

CMD ["tail", "-f", "/dev/null"]