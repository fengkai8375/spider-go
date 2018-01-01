package Spider

import (
	"strings"
	"time"
	"fmt"
	"net/http"
	"github.com/PuerkitoBio/goquery"
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
)

type Config struct {
	MaxDepth int
	MaxConnections int
	StartUrl string
	FetchOutsideLinks bool
	StoreTable string
	DbConfig string
	UserAgent string
}

type HtmlPage struct{
	url string
	title string
	content string
	statusCode int
	refer string
}

type HtmlUrl struct{
	url string
	depth int
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func Fetch(htmlUrl HtmlUrl ,queueUrls chan HtmlUrl,fetchResults chan HtmlPage, config Config){
	if htmlUrl.depth <= config.MaxDepth {
		client := &http.Client{}
		req,err := http.NewRequest("GET", htmlUrl.url, nil)
		if config.UserAgent != ""{
			req.Header.Add("User-Agent", config.UserAgent)
		}else{
			req.Header.Add("User-Agent", "SpiderGo V1.0")
		}
		//resp, err := http.Get(htmlUrl.url)
		resp, err := client.Do(req)
		if err != nil {
			panic(err)
		}

		document,err := goquery.NewDocumentFromResponse(resp) //该方法调用完成后会执行 resp.Body.Close()

		statusCode := resp.StatusCode

		//抓取结果
		page := HtmlPage{}
		page.url = htmlUrl.url
		documentTitle := document.Find("title").Eq(0).Text()
		page.title = documentTitle
		page.statusCode = statusCode
		page.content = document.Text()
		fetchResults <- page

		db, err := sql.Open("mysql", config.DbConfig)
		checkErr(err)

		stmt, err := db.Prepare("INSERT " + config.StoreTable +" SET url=?,title=?,content=?,statusCode=?,depth=?")
		checkErr(err)

		stmt.Exec(page.url, page.title, page.content, page.statusCode, htmlUrl.depth)

		db.Close()

		//页面链接处理
		if htmlUrl.depth <= config.MaxDepth - 1 {
			document.Find("a").Each(func(i int, selection *goquery.Selection) {
				href,err := selection.Attr("href")
				if err == false{
					//fmt.Println("link invalid")
				}
				if !strings.HasPrefix(href,"#") && !strings.HasPrefix(href,"mailto") && !strings.EqualFold(href,"/") {
					if strings.HasPrefix(href, "/"){
						href = config.StartUrl + href
					}
					//是否抓取外链
					if config.FetchOutsideLinks || strings.Contains(href ,config.StartUrl){
						newHtmlUrl := HtmlUrl{}
						newHtmlUrl.url = href
						newHtmlUrl.depth = htmlUrl.depth + 1
						queueUrls <- newHtmlUrl
					}
				}
			})
		}
	}
}

func SpiderGo(config Config){

	firstHtml := HtmlUrl{}
	firstHtml.url = config.StartUrl
	firstHtml.depth = 1

	queueUrls := make(chan  HtmlUrl, config.MaxConnections)
	fetchResults := make(chan HtmlPage, 10000)

	queueUrls <- firstHtml

	for true{

		select {
		case <- time.After(time.Second * 1) :
			fmt.Println("no url in queue")
		case htmlUrl := <- queueUrls :
			go Fetch(htmlUrl, queueUrls, fetchResults, config)
		}

		select {
		case <- time.After(time.Second * 5) :
			fmt.Println("no fetch result yet")
		case result := <- fetchResults :
			fmt.Println("Fetched:",result.title)
		}
	}
}