# spider-go
Golang Edition Spider

### 功能需求

- [x] 页面抓取：内容、状态码
- [x] 页面内容解析：DOM、正则
- [x] 抓取深度控制
- [x] 抓取内容存储：文件、数据库
- [x] 并发处理控制
- [x] User-Agent
- [ ] 代理：设置、频繁更换
- [ ] 表单提交
- [ ] cookie处理：接收、发送


### 性能需求

1. 内存占用低 ==整体资源消耗==
2. 阻塞轻 ==中间件==
3. 网络负载不高  ==读取/生产组件==
4. 数据库负载低  ==写入/消费组件==

### 依赖库

1. github.com/PuerkitoBio/goquery
2. github.com/go-sql-driver/mysql


### How to Use

```
func main(){
    config := Spider.Config{}
	config.MaxDepth = 2
	config.MaxConnections = 5
	config.StartUrl = "http://www.fengkai.info"
	config.StoreTable = "pages5"
	config.FetchOutsideLinks = false
	config.DbConfig = "user:password@tcp(127.0.0.1:3306)/test?charset=utf8"
	config.UserAgent = "your-ua"
	Spider.SpiderGo(config)
}
```
