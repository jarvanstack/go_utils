# 高性能 logger 使用方法

## 1.配置是否(默认打开)打印 debug 日志(debug 是最低级的日志)

文件地址: config/config.yaml

```yaml
# 是否输出 Debug() 级别的日志
isDebug: true
```

## 2. 使用

```go
func main() {
	logger.Info("info")
	logger.Debug("debug")
	logger.Error("error")
}
```