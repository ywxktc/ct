package main

import (
	"flag"
	"fmt"
	"time"
)

//var timestamp int64

type TimeFlag struct {
	Value     int64
	LongName  string
	ShortName string
	Usage     string
}

func NewTimeFlag() *TimeFlag {
	t := &TimeFlag{
		Value:     time.Now().Unix(),
		LongName:  "timestamp",
		ShortName: "t",
		Usage:     "timestamp",
	}

	flag.Int64Var(&t.Value, t.LongName, t.Value, t.Usage)
	flag.Int64Var(&t.Value, t.ShortName, t.Value, fmt.Sprintf("see -%s", t.LongName))
	flag.Parse()

	return t
}

func (t *TimeFlag) String() string {
	const format = "2006-01-02 03:04:05"
	return time.Unix(t.Value, 0).Format(format)
}

func main() {
	t := NewTimeFlag()
	fmt.Printf("time:      %s\ntimestamp: %d\n", t, t.Value)
}
