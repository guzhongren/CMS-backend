FROM golang:1.12.5-alpine3.9 as build-step
RUN apk add --update --no-cache ca-certificates git
MAINTAINER guzhongren "guzhongren@live.cn"
RUN mkdir /cms
WORKDIR /cms
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o /go/bin/cms
FROM scratch
COPY  --from=build-step /go/bin/cms /go/bin/cms
EXPOSE 1234
ENTRYPOINT ["/go/bin/cms"]