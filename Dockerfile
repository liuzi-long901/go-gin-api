FROM golang:1.19-alpine as builder

WORKDIR /app
COPY go.mod .
RUN go mod tidy
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /main main.go

FROM scratch
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder main /app/main
ENTRYPOINT ["/main"]


