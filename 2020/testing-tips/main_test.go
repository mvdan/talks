package main

import (
	"context"
	"flag"
	"testing"

	"github.com/chromedp/chromedp"
)

var headful = flag.Bool("headful", false, "")

func TestMeetupDate(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		url  string
		want string
	}{
		{
			"April2020",
			"https://www.meetup.com/GoSheffield/events/269190444/",
			"Thursday, April 2, 2020\n6:30 PM to 7:15 PM GMT+1",
		},
		{
			"May2020",
			"https://www.meetup.com/GoSheffield/events/270080480/",
			"Thursday, May 7, 2020\n6:30 PM to 7:30 PM GMT+1",
		},
		{
			"June2020",
			"https://www.meetup.com/GoSheffield/events/270842085/",
			"Thursday, June 4, 2020\n6:30 PM to 7:30 PM GMT+1",
		},
		{
			"July2020",
			"https://www.meetup.com/GoSheffield/events/271453771/",
			"Thursday, July 2, 2020\n6:30 PM to 7:30 PM GMT+1",
		},
	}

	actx := context.Background()
	if *headful {
		var cancel func()
		actx, cancel = chromedp.NewExecAllocator(actx)
		t.Cleanup(cancel)
	}

	ctx, cancel := chromedp.NewContext(actx)
	t.Cleanup(cancel)
	if err := chromedp.Run(ctx); err != nil {
		t.Fatal(err)
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test := test
			t.Parallel()

			got, err := MeetupDate(ctx, test.url)
			if err != nil {
				t.Fatal(err)
			}
			if got != test.want {
				t.Errorf("want %q, got %q", test.want, got)
			}
		})
	}
}
