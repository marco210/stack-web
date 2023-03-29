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

RUN go build -o main ./public/main.go

WORKDIR /dist

RUN cp /app/main .

EXPOSE 9000
CMD [ "/dist/main" ]