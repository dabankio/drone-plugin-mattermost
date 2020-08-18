FROM alpine:20200626
ADD plugin.run /bin/
ENTRYPOINT ["/bin/plugin.run"]
