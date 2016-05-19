FROM busybox
COPY app /
EXPOSE 5000
CMD ["./app"]
