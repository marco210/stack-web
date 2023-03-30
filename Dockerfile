# FROM golang:1.20-alpine
# WORKDIR /app
# COPY ./ ./
# RUN go mod download
# RUN go build -o stack-web ./public/main.go
# EXPOSE 9000
# ENTRYPOINT [ "/app/stack-web" ]

FROM golang:1.20-alpine

WORKDIR /app

COPY ./go.mod .
COPY ./go.sum .
COPY . . 

ENV DB_DRIVER mysql
ENV DB_USER root
ENV DB_PASS password
ENV DB_NAME book_example
ENV DB_HOST 0.0.0.0
ENV DB_PORT 3306

RUN go build -o main ./main.go

WORKDIR /dist

RUN cp /app/main .

EXPOSE 9000
CMD [ "/dist/main" ]