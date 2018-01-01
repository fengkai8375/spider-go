# spider-go
Golang Edition Spider

### 功能需求

- [x] 页面抓取：内容、状态码
- [x] 页面内容解析：DOM、正则
- [x] 抓取深度控制
- [x] 抓取内容存储：文件、数据库
- [x] 并发处理控制
- [ ] User-Agent
- [ ] 代理：设置、频繁更换
- [ ] POST数据
- [ ] cookie处理：接收、发送


### 性能需求

1. 内存占用低 ==整体资源消耗==
2. 阻塞轻 ==中间件==
3. 网络负载不高  ==读取/生产组件==
4. 数据库负载低  ==写入/消费组件==

### 依赖库

1. github.com/PuerkitoBio/goquery
2. github.com/go-sql-driver/mysql
