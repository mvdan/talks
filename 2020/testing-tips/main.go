package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/chromedp/chromedp"
)

func MeetupDate(ctx context.Context, url string) (string, error) {
	ctx, cancel := chromedp.NewContext(ctx)
	defer cancel()

	var date string
	if err := chromedp.Run(ctx,
		chromedp.Navigate(url),
		chromedp.Text(".eventTimeDisplay", &date, chromedp.ByQuery),
	); err != nil {
		return "", err
	}
	return date, nil
}

func main() {
	flag.Parse()

	actx := context.Background()
	ctx, cancel := chromedp.NewContext(actx)
	defer cancel()

	for _, arg := range flag.Args() {
		fmt.Println(arg)
		date, err := MeetupDate(ctx, arg)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(date)
		fmt.Println()
	}
}
