//go:build printbook

package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

const (
	colName = iota
	colSchool
	colLevel
	colCastTime
	colDistance
	colTarget
	colDuration
	colSave
	colResist
	colComponent
	colArea
	colEffect
	colDescription
)

var printBookParams = []struct {
	label  string
	col    int
	suffix string
}{
	{"Школа", colSchool, ";"}, {"Круг", colLevel, ""}, {"Время сотворения", colCastTime, ""},
	{"Компоненты", colComponent, ""}, {"Дистанция", colDistance, ""}, {"Цель", colTarget, ""},
	{"Область", colArea, ""}, {"Эффект", colEffect, ""}, {"Длит.", colDuration, ""},
	{"Испытание", colSave, ""}, {"Устойчивость к магии", colResist, ""},
}

func main() {
	reader := csv.NewReader(os.Stdin)
	reader.Read()
	for {
		r, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		id := strings.ReplaceAll(strings.ToLower(r[colName]), " ", "-")
		fmt.Printf("<h3 id=\"%s\">%s</h3>\n", id, r[colName])
		for _, p := range printBookParams {
			if r[p.col] != "" {
				fmt.Printf("<p><strong>%s: </strong>%s%s</p>\n", p.label, r[p.col], p.suffix)
			}
		}
		fmt.Println(r[colDescription])
		fmt.Println()
	}
}
