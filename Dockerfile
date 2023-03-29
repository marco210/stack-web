FROM golang:1.20-alpine

WORKDIR /app

RUN export GO111MODULE=on
RUN go get github.com/marco210/stack-web

RUN cd /app  && git clone https://github.com/marco210/stack-web.git

RUN cd /app/stack-web/public && go build

EXPOSE 9000

ENTRYPOINT [ "/app/stack-web/public/main" ]