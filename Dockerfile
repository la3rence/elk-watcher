FROM golang:1.15-alpine as builder
WORKDIR /usr/src/app
RUN apk add --no-cache upx
RUN apk add --update tzdata
COPY ./go.mod ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -ldflags "-s -w" -o watcher &&\
  upx --best watcher -o _upx_watcher && \
  mv -f _upx_watcher watcher

FROM scratch as runner
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /usr/src/app/watcher /opt/app/
CMD ["/opt/app/watcher"]
