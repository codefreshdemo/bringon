FROM golang:1.7.3 AS builder
WORKDIR /go/src/github.com/antweiss/gobot
ADD . .
RUN go get -d -v github.com/gorilla/mux
RUN CGO_ENABLED=0 GOOS=linux go build -o gobot

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /go/src/github.com/antweiss/gobot/gobot .
CMD ["./gobot"]
