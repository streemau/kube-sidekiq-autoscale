FROM golang

WORKDIR /go/src

COPY Makefile .

RUN make depend

COPY . .

RUN make && mv .build/autoscale /go/bin

ENTRYPOINT ["/go/bin/autoscale"]
