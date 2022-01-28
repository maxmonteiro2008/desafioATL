# syntax=docker/dockerfile:1

# Alpine is chosen for its small footprint
# compared to Ubuntu
FROM golang:1.17.6-bullseye

WORKDIR /app

# Download necessary Go modules


COPY ./src/desafio/ *.go ./ 
RUN true

RUN true
COPY  ./src/desafio/api ./
RUN true
COPY ./src/desafio/go ./
RUN true
#COPY ./src/go.mod ./
#RUN true
#COPY ./src/go.sum ./

RUN go mod init
RUN go mod download


RUN go build -o /docker-desafio-api

EXPOSE 8080/tcp

CMD [ "/docker-desafio-api" ]



