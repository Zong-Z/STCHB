FROM golang:1.16-alpine

ADD . /go/src/app
WORKDIR /go/src/app

RUN go get
RUN go install

RUN go build -o main .

CMD /go/src/app/main
