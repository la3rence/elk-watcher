![GitHub release (latest by date)](https://img.shields.io/github/v/release/Lonor/elk-watcher?color=%23FFFFFF)
![Build docker image](https://github.com/Lonor/elk-watcher/workflows/Build%20docker%20image/badge.svg)

# ELK watcher by Go

对接 Elastic Search 的 REST API，发送 HTTP 请求对应的索引数据，定时执行以实现周期监控. 效果类似于 ELK 付费白金版的 watcher 功能.

```shell
go build
chmod +x ./elk-watcher
```

## Export environment variable

在执行二进制文件之前，需要设置环境变量.

```shell
export ELASTIC_USERNAME="username"
export ELASTIC_PASSWORD="password"
# URL for ElasticSearch REST API
export ELASTIC_URL="http://elastic_hostname/index_name/_search"
export DINGTALK_TOKEN="access_token"
# 执行周期 (写法: 10s, 1m, 3h)
export DURATION="60s"
```

在上述环境下直接运行:

```shell
./elk-watcher
```

## Run as container

Docker 镜像，可以直接传递一样的环境变量作为容器运行（镜像体积几乎和编译后产生的二进制文件一样大）：

```shell
docker run --name=watcher -itd -e ELASTIC_USERNAME="username" \
  ELASTIC_PASSWORD="password" \
  ELASTIC_URL="http://elastic_hostname/index_name/_search" \
  DINGTALK_TOKEN="access_token" \
  DURATION="60s" \
  lawrence2018/watcher:latest
```

有镜像后，当然可以通过 Kubernetes 来编排。

```shell
kubectl apply -f deployment.yml
```
