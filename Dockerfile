FROM alpine:3.16

RUN apk --no-cache upgrade && apk --no-cache add ca-certificates
WORKDIR /efucloud

COPY ./http-demo /usr/local/bin/http-demo
EXPOSE 9000
ENTRYPOINT ["http-demo"]

