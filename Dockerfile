FROM alpine:latest
ADD app /go/bin/app
CMD /go/bin/app
