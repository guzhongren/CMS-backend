FROM alpine:latest
RUN apk --no-cache add ca-certificates
ENV TZ "Asia/Shanghai"
WORKDIR /root/
COPY ./conf.yaml .
COPY ./cms .
VOLUME [ "/root/assets/" ]
EXPOSE 1234
CMD ["./cms"]