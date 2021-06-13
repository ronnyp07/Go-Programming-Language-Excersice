package main

import (
	"Go-Programming-Language-Excersice/ch4/github"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	now := time.Now()
	beforeMonth := now.AddDate(0, -1, 0)

	fmt.Printf("%d Issues: \n", result.TotalCount)
	for _, item := range result.Items {
		fmt.Printf("#%-5d %9.9s %.55s  created_at: %s \n", item.Number, item.User.Login, item.Title, item.CreatedAt)
	}

	fmt.Println("created less than a month")
	for _, item := range result.Items {
		if item.CreatedAt.After(beforeMonth) {
			fmt.Printf("#%-5d %9.9s %.55s  created_at: %s \n", item.Number, item.User.Login, item.Title, item.CreatedAt)
		}
	}

	fmt.Printf("created less than a year \n")
	beforeYear := now.AddDate(-1, 1, 0)
	for _, item := range result.Items {
		if item.CreatedAt.After(beforeYear) {
			fmt.Printf("#%-5d %9.9s %.55s  created_at: %s \n", item.Number, item.User.Login, item.Title, item.CreatedAt)
		}
	}

	fmt.Printf("created more than a year \n")
	for _, item := range result.Items {
		if item.CreatedAt.Before(beforeYear) {
			fmt.Printf("#%-5d %9.9s %.55s  created_at: %s \n", item.Number, item.User.Login, item.Title, item.CreatedAt)
		}
	}

}
