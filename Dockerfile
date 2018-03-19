FROM golang:1.8

WORKDIR /go/src/github.com/AlexsJones/frontier
COPY . .

# Kafka component
RUN git clone https://github.com/edenhill/librdkafka.git
RUN cd librdkafka && ./configure --prefix /usr && make \
&& make install

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["frontier"]
