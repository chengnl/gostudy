package crawl

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

type result struct {
	Html    string
	More    bool
	Offset  int
	Enabled bool
}

type gallery struct {
	Gallery_title string
	PicInfo       []*duoimage
	HiidoId       []string
}

type duoimage struct {
	Title string
	//Pic_id      string
	Ding string
	Cai  string
	//Old         string
	Cover_url string
	File_url  string
	//File_width  string
	//File_height string
	Url    string
	Source string
	//Sort        string
	Cmt_md5     string
	Cmt_url     string
	Add_intro   string
	Comment_url string
}

func Crawl(baseUrl string, imagechan chan<- string) *result {
	resp, err := http.Get(baseUrl)
	defer resp.Body.Close()
	if err != nil {
		log.Print(err)
		return nil
	}
	if resp.StatusCode != http.StatusOK {
		return nil
	}

	var jsonObj result
	errr := json.NewDecoder(resp.Body).Decode(&jsonObj)
	if errr != nil {
		log.Fatal(errr)
		return nil
	}
	doc, err := html.Parse(strings.NewReader(jsonObj.Html))
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, attr := range n.Attr {
				if attr.Key == "href" {
					url := attr.Val
					if strings.HasPrefix(url, "http://tu.duowan.com/gallery/") && n.Parent.Data == "em" {
						url = strings.Replace(url, ".com", ".cn", 1)
						fmt.Println(url)
						bigImage(url, imagechan)
					}
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
	return &jsonObj
}

func bigImage(url string, imagechan chan<- string) {
	resp, err := http.Get(url)
	if err != nil {
		log.Print(err)
		return
	}
	if resp.StatusCode != http.StatusOK {
		return
	}

	doc, err := html.Parse(resp.Body)
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "script" {
			if child := n.FirstChild; child != nil {
				data := child.Data
				if strings.EqualFold(data[5:12], "imgJson") {
					jsonStr := data[15 : len(data)-2]
					//fmt.Println(jsonStr)
					var imagesObj gallery
					err := json.Unmarshal([]byte(jsonStr), &imagesObj)
					if err != nil {
						log.Fatal(err)
						return
					}
					for _, imageinfo := range imagesObj.PicInfo {
						imagechan <- imageinfo.Source
					}
				}
			}

		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
}
