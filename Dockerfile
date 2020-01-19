FROM golang:1.12.5 as builder

ENV GO111MODULE=on

COPY time_exporter.go util.go go.mod /build/
WORKDIR /build/
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o time_exporter .

FROM alpine:3.9.2

RUN apk add --no-cache tzdata
COPY --from=builder /build/ .

ENTRYPOINT [ "./time_exporter" ]

EXPOSE 9001
