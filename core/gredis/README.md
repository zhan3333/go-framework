# gredis

go-redis 连接管理

## 使用

```bash
 go get github.com/zhan3333/gredis@master 
```

使用前需要初始化配置:

```go
package main
import "github.com/zhan3333/gredis"

func main() {
    gredis.Configs = map[string]gredis.Conf{
        gredis.DefaultConn: {
            Host:     "127.0.0.1",
            Password: "",
            Port:     6379,
            Database: 0,
        },
        "customize": {
            Host:     "127.0.0.1",
            Password: "",
            Port:     6379,
            Database: 1,
        },
    }
    gredis.DefaultConn = "default"
    // 若需要重置配置, 需要调用 gredis.Reset()
    gredis.Def().Set("test", "test")
    gredis.Def().Get("test")
    gredis.Conn("customize").Set("foo", "bar")
}
```