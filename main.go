package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

const TIME_FORMAT = "2006-01-02T15:04:05Z"

func print_usage() {
	fmt.Println("Usage: tso <tso|time>")
	fmt.Println("Example:")
	fmt.Println("  tso 412345678901234567")
	fmt.Println("  tso 2023-11-10T23:08:03Z")
}

func get_stdin() (string, bool) {
	fi, err := os.Stdin.Stat()
	if err != nil {
		panic(err)
	}
	if (fi.Mode() & os.ModeCharDevice) == 0 {
		buf := make([]byte, 1024)
		n, err := os.Stdin.Read(buf)
		if err != nil {
			panic(err)
		}
		return string(buf[:n]), true
	}
	return "", false
}

func main() {
	input := strings.Join(os.Args[1:], " ")

	if input == "" {
		// check if stdin is piped
		if stdin, ok := get_stdin(); ok {
			input = stdin
		} else {
			print_usage()
			os.Exit(1)
		}
	}

	// trim any extra whitespace
	input = strings.Trim(input, " \n\r\t\"'")

	if v, err := strconv.Atoi(input); err == nil {
		fmt.Fprintf(os.Stderr, "Converting input '%s' to RFC3339...\n", input)
		formatted_time := tsoToTime(uint64(v)).Format(TIME_FORMAT)
		fmt.Println(formatted_time)
	} else {
		fmt.Fprintf(os.Stderr, "Converting input '%s' to TSO...\n", input)
		t, err := time.Parse(TIME_FORMAT, input)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Could not parse time: '%s'\n", input)
			fmt.Println(err)
			os.Exit(1)
		}
		tso_uint := timeToTSO(t)
		fmt.Println(tso_uint)
	}
}

func timeToTSO(t time.Time) uint64 {
	ts := t.UnixNano()
	val := uint64((ts / int64(time.Millisecond)) << 18)
	return val
}

func tsoToTime(tso uint64) time.Time {
	ts := int64(tso>>18) * int64(time.Millisecond)
	return time.Unix(0, ts).UTC()
}
