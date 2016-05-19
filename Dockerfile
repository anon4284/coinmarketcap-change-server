FROM busybox
COPY ca-certificates.crt /etc/ssl/certs/
COPY app /
EXPOSE 5000
CMD ["/app"]
