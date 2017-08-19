FROM golang:1.8-alpine as builder
WORKDIR /go/src/github.com/Hendra-Huang/httpsocket-vs-unixsocket/unixsocket
COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM nginx:1.13-alpine
COPY --from=builder /go/src/github.com/Hendra-Huang/httpsocket-vs-unixsocket/unixsocket/app /bin/app
COPY --from=builder /go/src/github.com/Hendra-Huang/httpsocket-vs-unixsocket/unixsocket/dockerbuild.sh ./dockerbuild.sh
COPY --from=builder /go/src/github.com/Hendra-Huang/httpsocket-vs-unixsocket/unixsocket/nginx.conf /etc/nginx/nginx.conf
COPY --from=builder /go/src/github.com/Hendra-Huang/httpsocket-vs-unixsocket/unixsocket/default.conf /etc/nginx/conf.d/default.conf
CMD sh ./dockerbuild.sh
