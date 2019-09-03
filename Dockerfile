FROM golang AS builder

WORKDIR /go/src

ENV GO111MODULE=on

COPY Makefile .
COPY go.mod .
COPY go.sum .

RUN make depend

COPY . .

RUN make && mv .build/autoscale /go/bin

FROM golang

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /go/bin/autoscale /go/bin/autoscale

ENTRYPOINT ["/go/bin/autoscale"]
