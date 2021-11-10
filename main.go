package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	format := "s"
	if len(os.Args) > 1 {
		format = os.Args[1]
	}

	ts := time.Now()
	switch format {
	case "-h", "--help", "help":
		fmt.Fprintf(os.Stderr, `Usage: %s [format]

Format can be one of:

  s:      Unix timestamp with second precision [default]
  ms:     Unix timestamp with millisecond precision
  us:     Unix timestamp with microsecond precision
  string: Go time formatting string
`, os.Args[0])
		os.Exit(2)
	case "s", "sec", "seconds", "unix":
		fmt.Println(ts.Unix())
	case "ms", "milli", "milliseconds":
		fmt.Println(ts.UnixMilli())
	case "us", "micro", "microseconds":
		fmt.Println(ts.UnixMicro())
	default:
		fmt.Println(ts.Format(format))
	}
}
