FROM golang:alpine as build-step
RUN apk add --update --no-cache ca-certificates git
LABEL maintainer="guzhongren@live.cn"
RUN mkdir /cms
WORKDIR /cms
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 go build -o /go/bin/cms

FROM scratch
COPY  --from=build-step /go/bin/cms /go/bin/cms
COPY ./conf.yaml /go/bin/conf.yaml
EXPOSE 1234
CMD ["/go/bin/cms"]