package io

import (
	"os"
	"stocks/types"
	"strings"

	"github.com/jedib0t/go-pretty/v6/table"
)

func Print(ticker string, recentData types.Filling) {
	var rows []table.Row

	for index, accessionNumber := range recentData.AccessionNumber {
		accessionNumber = strings.Replace(accessionNumber, "-", "", -1)
		//println(accessionNumber)

		url := "https://www.sec.gov/Archives/edgar/data/000" + ticker + "/" + accessionNumber + "/" + recentData.PrimaryDocument[index]

		formName := ""
		switch recentData.Form[index] {
		case "4":
			formName = "\033[32m" + "Insider trading" + "\033[0m"
		default:
			formName = recentData.Form[index]

		}

		row := table.Row{
			recentData.FilingDate[index],
			formName,
			url,
		}

		rows = append(rows, row)
	}

	PrintTable(rows)
}
func PrintTable(rows []table.Row) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	row := table.Row{"Filling Date", "Form Type", "URL"}
	t.AppendHeader(row)
	t.AppendRows(rows)
	//t.AppendSeparator()
	t.Render()
}
