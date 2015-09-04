FROM busybox
ADD ./authit_linux-amd64 /app
CMD ["/app"]
