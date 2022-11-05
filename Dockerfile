FROM golang:1.19.3 AS builder
COPY . /workspace
WORKDIR /workspace
ENV GO111MODULE=on
RUN CGO_ENABLED=0 GOOS=linux go build -o app

FROM scratch
COPY --from=builder /workspace/certs/ca.crt /etc/ssl/certs/
COPY --from=builder /workspace/app /
CMD ["/app"]