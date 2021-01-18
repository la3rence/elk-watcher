# ELK watcher (GO 实现)

对接 Elastic Search 的 REST API，发送 HTTP 请求对应的索引数据，定时执行以实现周期监控.

```shell
go build
chmod +x ./elk-watcher
./elk-watcher
```

## ENV export

在执行二进制文件之前，需要设置环境变量.

```shell
export ELASTIC_USERNAME="username"
export ELASTIC_PASSWORD="password"
export ELASTIC_URL="http://elastic_hostname/index_name/_search"
export DINGTALK_TOKEN="access_token"
```


