FROM alpine:20200626
# RUN apk update
ADD plugin.run /bin/
ENTRYPOINT ["/bin/plugin.run"]
