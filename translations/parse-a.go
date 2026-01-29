//go:build parsea

package main

import (
	"bufio"
	"encoding/csv"
	"os"
	"regexp"
	"strings"
)

var (
	noteRegexp     = regexp.MustCompile(`{{\?{1,2}\|[\pL\pN ]*}}`)
	untranslatedRe = regexp.MustCompile(`{{\?{3}\|([\pL\pN ]*)}}`)
	italicRegexp   = regexp.MustCompile(`''([\pL\pN\ ]*)''`)
)

func norm(str string) string {
	str = noteRegexp.ReplaceAllString(str, "")
	str = untranslatedRe.ReplaceAllString(str, "<i>$1</i>")
	str = italicRegexp.ReplaceAllString(str, "<i>$1</i>")
	return str
}

func main() {
	writer := csv.NewWriter(os.Stdout)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	for {
		nameRu := norm(strings.Trim(scanner.Text(), "= "))
		scanner.Scan()
		nameEn := norm(strings.Trim(scanner.Text(), "{?|}"))
		scanner.Scan()
		for strings.HasPrefix(scanner.Text(), ": '''") {
			scanner.Scan()
		}

		descrRu := "<p>"
		scanner.Scan()
		if scanner.Text() != "" {
			descrRu += norm(scanner.Text())
		}
		inParagraph := false
		var moreData bool
		for {
			moreData = scanner.Scan()
			if !moreData || strings.HasPrefix(scanner.Text(), "===") {
				break
			}
			if scanner.Text() == "" {
				descrRu += "</p>"
				inParagraph = true
			} else {
				if inParagraph {
					descrRu += "<p>"
					inParagraph = false
				}
				descrRu += norm(scanner.Text())
			}
		}

		writer.Write([]string{nameEn, nameRu, descrRu})
		if !moreData {
			break
		}
	}
	writer.Flush()
}
