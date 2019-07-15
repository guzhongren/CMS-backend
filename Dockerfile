FROM golang:alpine as builder
RUN apk add --update --no-cache ca-certificates git
LABEL maintainer="guzhongren@live.cn"
RUN mkdir /cms
WORKDIR /cms
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 go build -o /go/bin/cms

FROM alpine:latest
RUN apk --no-cache add ca-certificates
ENV TZ "Asia/Shanghai"
WORKDIR /root/
COPY ./conf.yaml .
COPY  --from=builder /go/bin/cms .
VOLUME [ "/root/assets/" ]
EXPOSE 1234
CMD ["./cms"]