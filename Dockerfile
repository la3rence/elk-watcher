FROM golang:1.15-alpine as builder
WORKDIR /usr/src/app
RUN apk add --no-cache upx
COPY ./go.mod ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -ldflags "-s -w" -o watcher &&\
  upx --best watcher -o _upx_watcher && \
  mv -f _upx_watcher watcher

FROM scratch as runner
COPY --from=builder /usr/src/app/watcher /opt/app/
CMD ["/opt/app/wathcer"]
