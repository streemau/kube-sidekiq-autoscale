FROM golang:1.13 AS builder

WORKDIR /go/src/app

ENV GO111MODULE=on

COPY Makefile .
COPY go.mod .
COPY go.sum .

RUN make depend

COPY . .

RUN make && mv .build/autoscale /go/bin

FROM golang:1.13

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /go/bin/autoscale /go/bin/autoscale

ENTRYPOINT ["/go/bin/autoscale"]
