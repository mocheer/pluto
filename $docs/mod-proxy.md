
# go 包管理器代理
```ps
$env:GO111MODULE="on"
$env:GOPROXY="https://goproxy.io"
go env -w GOPROXY=https://goproxy.io,direct
go env -w GOPRIVATE=*.corp.example.com
```
