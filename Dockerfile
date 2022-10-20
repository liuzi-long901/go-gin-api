FROM golang:1.19-alpine as builder

COPY . .
RUN go mod download && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /main main.go

FROM scratch
WORKDIR /app
#COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder main /bin/main
COPY ./init/web-docker.yaml /init/web.yaml
ENTRYPOINT ["/bin/main"]



