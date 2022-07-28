package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

// vscode에서 go.mode file not found ... 오류가 나면
// 아래 터미널 창에서 go env -w GO111MODULE=auto 실행 후 해보시길.

var baseURL = "https://kr.indeed.com/jobs?q=python&start="

type joblist struct {
	jobtitle string
	location string
	link     string
}

func main() {

	for i := 0; i < 10; i++ {
		getPages(i)
	}

}

func getPages(page int) {
	var jobs []joblist
	c := make(chan joblist)
	pageURL := baseURL + strconv.Itoa(page*10)
	res, err := http.Get(pageURL)
	checkErr(err)
	checkCode(res)

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	searchData := doc.Find("div.job_seen_beacon")
	searchData.Each(func(i int, s *goquery.Selection) {
		go getJobs(s, c)
		job := <-c
		jobs = append(jobs, job)
	})

	fmt.Println(jobs)
}

func getJobs(card *goquery.Selection, c chan<- joblist) {
	title := card.Find("h2.jobTitle a>span").Text()
	location := card.Find("div.companyLocation").Text()
	datajk, _ := card.Find("a").Attr("data-jk")
	c <- joblist{
		jobtitle: title,
		location: location,
		link:     datajk,
	}
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func checkCode(res *http.Response) {
	if res.StatusCode != 200 {
		log.Fatalln("error status", res.StatusCode)
	}
}
