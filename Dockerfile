FROM golang:1.12 as builder
RUN mkdir -p $GOPATH/src/github.com/sandeeplamb/pod-labeller
ADD . $GOPATH/src/github.com/sandeeplamb/pod-labeller/
WORKDIR $GOPATH/src/github.com/sandeeplamb/pod-labeller
RUN go get -d -v ./... 
# RUN go install -v ./...
# RUN sed -i "s/klog.V(9)/klog.V(9).Enabled()/g" /go/src/k8s.io/client-go/transport/round_trippers.go \
#  && sed -i "s/klog.V(8)/klog.V(9).Enabled()/g" /go/src/k8s.io/client-go/transport/round_trippers.go \
#  && sed -i "s/klog.V(7)/klog.V(9).Enabled()/g" /go/src/k8s.io/client-go/transport/round_trippers.go \
#  && sed -i "s/klog.V(6)/klog.V(9).Enabled()/g" /go/src/k8s.io/client-go/transport/round_trippers.go \
#  && sed -i "s/klog.V(10)/klog.V(10).Enabled()/g" /go/src/k8s.io/client-go/rest/request.go \
#  && sed -i "s/klog.V(9)/klog.V(9).Enabled()/g" /go/src/k8s.io/client-go/rest/request.go \
#  && sed -i "s/klog.V(8)/klog.V(8).Enabled()/g" /go/src/k8s.io/client-go/rest/request.go 

RUN go get k8s.io/klog && cd $GOPATH/src/k8s.io/klog && git checkout v0.4.0
RUN go install -v ./...
RUN go build main.go && ls -l && echo $GOPATH

FROM golang:1.12
RUN mkdir /app
COPY --from=builder /go/src/github.com/sandeeplamb/pod-labeller/main /app/
WORKDIR /app
CMD ["./main"]