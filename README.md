![GitHub release (latest by date)](https://img.shields.io/github/v/release/Lonor/elk-watcher?color=%23FFFFFF)
![Build docker image](https://github.com/Lonor/elk-watcher/workflows/Build%20docker%20image/badge.svg)

# ELK watcher by Go

This project calls the REST API of Elastic Search, especially the index data, at regular intervals. It may look like the
paid features in the ELK stack -- watcher.

```shell
go build
chmod +x ./elk-watcher
```

## Export environment variable

Before starting this app, the following env variables need to be set.

```shell
export ELASTIC_USERNAME="username"
export ELASTIC_PASSWORD="password"
# URL for ElasticSearch REST API
export ELASTIC_URL="http://elastic_hostname/index_name/_search"
export DINGTALK_TOKEN="access_token"
export DINGTALK_SECRET="dingtalk_secret"
# interval (pattern with `10s`, `1m`, `3h`)
export DURATION="60s"
```

Run it by this command:

```shell
./elk-watcher
```

## Run as container

```shell
docker run --name=watcher -itd -e ELASTIC_USERNAME="username" \
  ELASTIC_PASSWORD="password" \
  ELASTIC_URL="http://elastic_hostname/index_name/_search" \
  DINGTALK_TOKEN="access_token" \
  DINGTALK_SECRET="dingtalk_secret"
  DURATION="60s" \
  lawrence2018/watcher:latest
```

Of course, we can deploy this with Kubernetes.

```shell
kubectl apply -f deployment.yml
```
