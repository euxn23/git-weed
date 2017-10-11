package main

import (
	"github.com/urfave/cli"
)

var (
	commitMessage string
	amend         bool

	strict bool
	random bool

	year   int
	month  int
	day    int
	hour   int
	minute int
	second int

	yearsAgo   int
	monthsAgo  int
	daysAgo    int
	hoursAgo   int
	minutesAgo int
	secondsAgo int
)

var Flags = []cli.Flag{
	cli.StringFlag{
		Name:        "message, m",
		Usage:       "commit message",
		Destination: &commitMessage,
	},
	//cli.BoolFlag{
	//	Name: "amend",
	//	Usage: "amend previous commit",
	//	Destination: &amend,
	//},

	cli.BoolFlag{
		Name:        "strict, s",
		Usage:       "Not allow future commit.",
		Destination: &strict,
	},
	cli.BoolFlag{
		Name:        "random, r",
		Usage:       "Generate random time after HEAD, before than now.",
		Destination: &random,
	},

	cli.IntFlag{
		Name:        "year",
		Usage:       "The year.",
		Destination: &year,
	},
	cli.IntFlag{
		Name:        "month",
		Usage:       "The month.",
		Destination: &month,
	},
	cli.IntFlag{
		Name:        "day",
		Usage:       "The day.",
		Destination: &day,
	},
	cli.IntFlag{
		Name:        "hour",
		Usage:       "The hour.",
		Destination: &hour,
	},
	cli.IntFlag{
		Name:        "minute",
		Usage:       "The minute.",
		Destination: &minute,
	},
	cli.IntFlag{
		Name:        "second",
		Usage:       "The second.",
		Destination: &second,
	},

	cli.IntFlag{
		Name:        "years-ago",
		Usage:       "How many years ago.",
		Destination: &yearsAgo,
	},
	cli.IntFlag{
		Name:        "months-ago",
		Usage:       "How many months ago.",
		Destination: &monthsAgo,
	},
	cli.IntFlag{
		Name:        "days-ago",
		Usage:       "How many days ago.",
		Destination: &daysAgo,
	},
	cli.IntFlag{
		Name:        "hours-ago",
		Usage:       "How many hours ago.",
		Destination: &hoursAgo,
	},
	cli.IntFlag{
		Name:        "minutes-ago",
		Usage:       "How many minutes ago.",
		Destination: &minutesAgo,
	},
	cli.IntFlag{
		Name:        "second-ago",
		Usage:       "How many seconds ago.",
		Destination: &secondsAgo,
	},
}
