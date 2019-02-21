FROM golang:latest

RUN mkdir -p /go/src/barnett/api

WORKDIR /go/src/barnett/api

COPY . /go/src/barnett/api

RUN go install

CMD /go/bin/api

EXPOSE 8080