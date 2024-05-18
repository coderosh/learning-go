package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/go-resty/resty/v2"
	strip "github.com/grokify/html-strip-tags-go" // => strip
	"golang.org/x/net/html"
)

var client = resty.New()

type ManganatoResponse struct {
	LinkSearchByKw string  `json:"link_search_by_kw"`
	SearchList     []Manga `json:"searchlist"`
}

type Manga struct {
	Author       string `json:"author"`
	Id           string `json:"id"`
	Image        string `json:"image"`
	LastChapter  string `json:"lastchapter"`
	Name         string `json:"name"`
	NameUnsigned string `json:"nameunsigned"`
	UrlStory     string `json:"url_story"`
}

func (m Manga) Title() string {
	return m.Name
}

func (m Manga) Description() string {
	return m.Author
}

func (m Manga) FilterValue() string {
	return m.Name
}

type Chapter struct {
	Href string
	Name string
}

func (m Chapter) Title() string {
	return m.Name
}

func (m Chapter) Description() string {
	return m.Href
}

func (m Chapter) FilterValue() string {
	return m.Name
}

func SearchManga(search string) []Manga {
	req := client.R().
		SetFormData(map[string]string{
			"searchword": search,
		}).
		SetHeaders(map[string]string{
			"Content-Type": "application/x-www-form-urlencoded",
			"Referer":      "https://manganato.com/",
			"Origin":       "https://manganato.com",
		})

	resp, err := req.Post("https://manganato.com/getstorysearchjson")
	if err != nil {
		log.Fatal(err)
	}

	mresponse := ManganatoResponse{}
	err = json.Unmarshal([]byte(resp.String()), &mresponse)
	if err != nil {
		log.Fatal(err)
	}

	for i := range mresponse.SearchList {
		mresponse.SearchList[i].Name = strip.StripTags(mresponse.SearchList[i].Name)
		mresponse.SearchList[i].Author = strip.StripTags(mresponse.SearchList[i].Author)
	}

	return mresponse.SearchList
}

func GetChapters(manga *Manga) []Chapter {
	resp, err := client.R().Get(manga.UrlStory)
	if err != nil {
		log.Fatal(err)
	}

	node, err := html.Parse(strings.NewReader(resp.String()))
	if err != nil {
		log.Fatal(err)
	}

	doc := goquery.NewDocumentFromNode(node)

	chapters := []Chapter{}

	doc.Find(".row-content-chapter li a:first-child").Each(func(_ int, g *goquery.Selection) {
		href, _ := g.Attr("href")
		name := g.Text()

		chapter := Chapter{Name: name, Href: href}
		chapters = append(chapters, chapter)
	})

	return chapters
}

func GetChapterImagesUrl(chapter *Chapter) []string {
	resp, err := client.R().Get(chapter.Href)
	if err != nil {
		log.Fatal(err)
	}

	node, err := html.Parse(strings.NewReader(resp.String()))
	if err != nil {
		log.Fatal(err)
	}

	doc := goquery.NewDocumentFromNode(node)

	images := []string{}
	doc.Find(".container-chapter-reader img").Each(func(_ int, s *goquery.Selection) {
		src, _ := s.Attr("src")
		images = append(images, src)
	})

	return images
}

func CreatePdfFromImages(images []string, fileName string) {
	err := os.RemoveAll("/tmp/manganato/")
	if err != nil {
		log.Fatal(err)
	}

	err = os.MkdirAll("/tmp/manganato", 0777)
	if err != nil {
		log.Fatal(err)
	}

	files := []string{}

	for i, image := range images {
		resp, err := client.R().SetHeader("Referer", "https://manganato.com/").Get(image)
		if err != nil {
			log.Fatal(err)
		}

		fileName := fmt.Sprintf("/tmp/manganato/%d.png", i)

		files = append(files, fileName)

		os.WriteFile(fileName, resp.Body(), 0644)
	}

	script := fmt.Sprintf("img2pdf %v -o %v.pdf", strings.Join(files, " "), fileName)
	cmd := exec.Command("sh", "-c", script)

	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}
