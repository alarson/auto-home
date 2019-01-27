FROM golang:alpine as builder
RUN mkdir /build 
COPY services/foo/app/ /build
COPY src/ /go/src/
#ADD . /build/
WORKDIR /build 
RUN go build -o main .
FROM alpine
RUN adduser -S -D -H -h /app appuser
USER appuser
COPY --from=builder /build/main /app/
COPY --from=builder /build/run.sh /app/
USER root
RUN apk add --no-cache bash
WORKDIR /app
CMD ["./main"]

