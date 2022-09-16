FROM golang:1.19-alpine as builder

WORKDIR /build
COPY go.mod .
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /main main.go


FROM scratch
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder main /bin/main
COPY ./init/web.yaml /build/init/web.yaml
ENTRYPOINT ["/bin/main"]



