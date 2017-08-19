FROM golang:1.8-alpine as builder
WORKDIR /go/src/github.com/Hendra-Huang/httpsocket-vs-unixsocket/fasthttp
COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM nginx:1.13-alpine
COPY --from=builder /go/src/github.com/Hendra-Huang/httpsocket-vs-unixsocket/fasthttp/app /bin/app
COPY --from=builder /go/src/github.com/Hendra-Huang/httpsocket-vs-unixsocket/fasthttp/dockerbuild.sh ./dockerbuild.sh
COPY --from=builder /go/src/github.com/Hendra-Huang/httpsocket-vs-unixsocket/fasthttp/nginx.conf /etc/nginx/nginx.conf
COPY --from=builder /go/src/github.com/Hendra-Huang/httpsocket-vs-unixsocket/fasthttp/default.conf /etc/nginx/conf.d/default.conf
CMD sh ./dockerbuild.sh
