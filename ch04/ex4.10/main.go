package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"bootcamp/ch04/github"
)

func queryAndReportIssues(args []string) {
	result, err := github.SearchIssues(args)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, item := range result.Items {
		fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
	}
}

func main() {
	args := os.Args[1:]

	now := time.Now()
	oneMonthBefore := now.Add(-time.Duration(1) * time.Hour * 24 * 30)
	args = append(args, "created:>"+oneMonthBefore.Format("2006-01-02"))
	fmt.Printf("Issues created within a month:\n")
	queryAndReportIssues(args)

	oneYearBefore := now.Add(-time.Duration(1) * time.Hour * 24 * 365)
	args = append(os.Args[1:], "created:>"+oneYearBefore.Format("2006-01-02"))
	fmt.Printf("Issues created within a year:\n")
	queryAndReportIssues(args)

	args = append(os.Args[1:], "created:<="+oneYearBefore.Format("2006-01-02"))
	fmt.Printf("Issues created over a year ago:\n")
	queryAndReportIssues(args)
}
