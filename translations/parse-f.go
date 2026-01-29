//go:build parsef

package main

import (
	"bufio"
	"encoding/csv"
	"os"
	"strings"
)

var headers = []string{
	"Дальность:",
	"Компоненты:",
	"Наведение Заклинания:",
	"Продолжительность:",
	"Скорость сотворения:",
	"Сопротивляемость Заклинаниям:",
	"Спасбросок:",
	"Уровень:",
	"Школа:",
	"Школа(Подшкола):",
}

func isHeader(str string) bool {
	for _, h := range headers {
		if strings.HasPrefix(str, h) {
			return true
		}
	}
	return false
}

func main() {
	writer := csv.NewWriter(os.Stdout)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	for {
		nameEn := strings.Trim(scanner.Text(), "() ")
		scanner.Scan()
		nameRu := strings.TrimSpace(scanner.Text())
		scanner.Scan()
		for isHeader(scanner.Text()) {
			scanner.Scan()
		}

		var descrRu strings.Builder
		var moreData bool
		for {
			if scanner.Text() != "" {
				descrRu.WriteString("<p>")
				descrRu.WriteString(scanner.Text())
				descrRu.WriteString("</p>")
			}
			moreData = scanner.Scan()
			if !moreData || strings.HasPrefix(scanner.Text(), "(") {
				break
			}
		}

		writer.Write([]string{nameEn, nameRu, descrRu.String()})
		if !moreData {
			break
		}
	}
	writer.Flush()
}
