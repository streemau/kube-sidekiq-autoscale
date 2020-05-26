FROM golang:1.13 AS builder

WORKDIR /go/src/app

ENV GO111MODULE=on

COPY Makefile .
COPY go.mod .
COPY go.sum .

RUN go get -u k8s.io/client-go@v0.17.2 github.com/googleapis/gnostic@v0.3.1 ./...

RUN make depend

COPY . .

RUN make && mv .build/autoscale /go/bin

FROM golang:1.13

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /go/bin/autoscale /go/bin/autoscale

ENTRYPOINT ["/go/bin/autoscale"]
