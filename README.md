# logWriter

专门为日志场景实现的一个io.writer，新建一个writer之后可以设置日志保存相关参数。

## 用法

NewLogger默认配置

```go
logger := NewLogger("./test.log")
log.SetOutput(logger)
logger.SetMaxSize(1)
logger.SetMaxBackup(3)
logger.SetMaxAge(3)
logger.DisCompress()
logger.DisLocalTime()
logger.SetFormatTime("2006-01-02-15-04-05")

logger.Write([]byte("hello worldhello worldhello worldhello world\n"))
```

### 参数介绍：

```
MaxSize: 日志文件的最大保存大小，超过大小会进行日志保存。默认100M
MaxBackups：MaxBackups是保留旧日志文件的最大数量。默认的全部保存
MaxAge：最大保存天数。 默认不进行删除。
BackupTimeFormat:备份的时间戳格式。默认 2006-01-02_15:04:05
LocalTime:格式化时间戳格式.备份文件是计算机的本地时间。默认是使用UTC.
Compress:是否压缩，压缩格式是.gz 默认进行压缩
```

## Tip:

* BackupTimeFormat参数必须和time.Time{}.format中的格式一致
* 保存格式为：**文件名-创建时间.后缀**
