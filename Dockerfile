FROM golang AS builder

WORKDIR /go/src

COPY Makefile .

RUN make depend

COPY . .

RUN make && mv .build/autoscale /go/bin

FROM golang

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /go/bin/autoscale /go/bin/autoscale

ENTRYPOINT ["/go/bin/autoscale"]
