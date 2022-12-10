From golang:1.18.9-buster as builder

WORKDIR /go/src/pizza/src
COPY . /go/src/pizza

RUN go get -d -v
RUN go build -o /go/bin/pizza

FROM gcr.io/distroless/base-debian10
COPY --from=builder /go/bin/pizza /
CMD ["/pizza"]
