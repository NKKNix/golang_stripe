FROM golang:1.23-alpine3.20 AS builder

RUN apk update && apk add --no-cache ca-certificates

WORKDIR /app
COPY . .

COPY ./src/docs/swagger.yaml /app/src/docs/swagger.yaml

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /main .

FROM scratch
WORKDIR /
COPY --from=builder /main /main
COPY --from=builder /app/src/docs /src/docs
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
ENTRYPOINT ["/main"]

