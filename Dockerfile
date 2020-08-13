FROM golang:1.14.6-alpine3.12 as builder

COPY ./src/greeting/main.go .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o app

FROM scratch

COPY --from=builder /go/app .

EXPOSE 8000

CMD ["./app"]
