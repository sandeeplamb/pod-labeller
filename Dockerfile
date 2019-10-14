FROM golang:1.12 as builder
RUN mkdir -p $GOPATH/src/github.com/sandeeplamb/pod-labeller
ADD . $GOPATH/src/github.com/sandeeplamb/pod-labeller
WORKDIR $GOPATH/src/github.com/sandeeplamb/pod-labeller
RUN go get -d -v ./... && go install -v ./...
RUN go build main.go && ls -l && echo $GOPATH

FROM golang:1.12
RUN mkdir /app
COPY --from=builder /go/src/github.com/sandeeplamb/pod-labeller/main /app/
WORKDIR /app
CMD ["./main"]