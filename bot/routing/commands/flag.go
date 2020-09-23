package commands

import (
	flag "github.com/spf13/pflag"
)

type BoolFlag struct {
	LongFlag  string
	ShortFlag string
	FlagUsage string
	Value     bool
}

func (b BoolFlag) Long() string {
	return b.LongFlag
}

func (b BoolFlag) Short() string {
	return b.ShortFlag
}

func (b BoolFlag) Usage() string {
	return b.FlagUsage
}

func (b BoolFlag) Register(f *flag.FlagSet) {
	f.BoolP(b.LongFlag, b.ShortFlag, b.Value, b.FlagUsage)
}

type StringFlag struct {
	LongFlag  string
	ShortFlag string
	FlagUsage string
	Value     string
}

func (s StringFlag) Long() string {
	return s.LongFlag
}

func (s StringFlag) Short() string {
	return s.ShortFlag
}

func (s StringFlag) Usage() string {
	return s.FlagUsage
}

func (s StringFlag) Register(f *flag.FlagSet) {
	f.StringP(s.LongFlag, s.ShortFlag, s.Value, s.FlagUsage)
}

type UInt64Flag struct {
	LongFlag  string
	ShortFlag string
	FlagUsage string
	Value     uint64
}

func (s UInt64Flag) Long() string {
	return s.LongFlag
}

func (s UInt64Flag) Short() string {
	return s.ShortFlag
}

func (s UInt64Flag) Usage() string {
	return s.FlagUsage
}

func (s UInt64Flag) Register(f *flag.FlagSet) {
	f.Uint64P(s.LongFlag, s.ShortFlag, s.Value, s.FlagUsage)
}
