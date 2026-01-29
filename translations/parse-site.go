//go:build parsesite

package main

import (
	"encoding/csv"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const ARROW = "→"

var (
	fontSizeRe = regexp.MustCompile(`font-size:[^;" ]*;?`)
	aOpenRe    = regexp.MustCompile(`<a[^>]*>`)
	aCloseRe   = regexp.MustCompile(`</a[^>]*>`)
	imgRe      = regexp.MustCompile(`<img[^>]*>`)
)

func cleanHTML(html string) string {
	html = fontSizeRe.ReplaceAllString(html, "")
	html = imgRe.ReplaceAllString(html, "")
	html = aOpenRe.ReplaceAllString(html, "<i>")
	html = aCloseRe.ReplaceAllString(html, "</i>")
	return html
}

func main() {
	doc, err := goquery.NewDocumentFromReader(os.Stdin)
	if err != nil {
		log.Fatalln(err)
	}
	name_ru := doc.Find("h3#sites-page-title-header span#sites-page-title").Text()
	src := doc.Find("#sites-canvas-main-content > table > tbody > tr > td > div > div:last-child > ul > li > font").Text()
	name_en := strings.TrimSpace(src[strings.Index(src, ARROW)+len(ARROW):])
	descr := ""
	doc.Find("#sites-canvas-main-content > table > tbody > tr > td > div > div > div > div").Each(func(i int, s *goquery.Selection) {
		if s.Text() == "ОПИСАНИЕ" {
			s.NextAll().Each(func(i int, c *goquery.Selection) {
				html, err := goquery.OuterHtml(c)
				if err != nil {
					log.Println(err)
				}
				descr += cleanHTML(html)
			})
		}
	})

	w := csv.NewWriter(os.Stdout)
	w.Write([]string{name_en, name_ru, descr})
	w.Flush()
}
