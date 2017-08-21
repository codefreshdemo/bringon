FROM golang:1.7.3 AS builder
WORKDIR /go/src/github.com/antweiss/bringon
ADD . .
#RUN go get -d -v github.com/gorilla/mux && go get gopkg.in/mgo.v2
RUN go get -u github.com/golang/dep/cmd/dep && dep ensure -v
RUN CGO_ENABLED=0 GOOS=linux go build -o bringon

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
EXPOSE 8091
COPY --from=builder /go/src/github.com/antweiss/bringon/bringon .
CMD ["./bringon"]
