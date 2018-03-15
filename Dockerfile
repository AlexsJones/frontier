FROM golang:1.8

WORKDIR /go/src/github.com/AlexsJones/frontier
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["frontier"]
