FROM alpine

ADD deligo /go/bin/
ADD header.html /go/bin/
ADD dist /go/bin/dist
ADD delicious.json /go/bin/
ADD ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

WORKDIR /go/bin/

ENTRYPOINT /go/bin/deligo
