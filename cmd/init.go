package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"stocks/client"
	"stocks/io"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Returns data for a stock",
	Long: `This command will io you for a stock ticker or company name, and return a list of matching stock.
After you select a stock, it will show relevant information about the company.`,
	Run: func(cmd *cobra.Command, args []string) {
		ticker, _ := cmd.Flags().GetString("ticker")
		runCommand(ticker)
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
	initCmd.Flags().StringP("ticker", "t", "", "Ticker symbol for the stock")
}

func runCommand(input string) {
	if input == "" {
		input, _ = io.Prompt("Search for a stock or company name")
	}

	//makes a call to the SEC API and returns a list of stocks
	stocks, err := client.FetchCompanies(input)

	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		return
	}

	if len(stocks.Hits.Hits) == 0 {
		fmt.Println("No results found")
		return
	}

	index, _ := io.Select(stocks, err)

	selected := stocks.Hits.Hits[index]

	fmt.Println("Selected:" + selected.Source.Ticker)

	info, err := client.FetchCompanyInfo(selected.Id)
	if err != nil {
		return
	}

	fmt.Println(info.Ticker)
	fmt.Println(info.Name)
	fmt.Println(info.Description)
	fmt.Println(info.Addresses.Business.City)

	fmt.Println(info.Filings.Files)

	fillings := info.Filings.Recent

	io.Print(info.Cik, fillings)

}
