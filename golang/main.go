package main

import (
	"flag"
	"os"

	"github.com/VyacheslavIsWorkingNow/tfl/lab2/internal/benchmark"
	"github.com/VyacheslavIsWorkingNow/tfl/lab2/internal/reggen"
)

const (
	countWords  = 100
	maxDumpSize = 200
)

// TODO: можно сделать красивый логер для всего этого
// и это было бы приятным дополнением
// TODO: для увеличения скорости построения слов перейти на strings.Builder

func main() {
	var simplifier_path string
	var countRegex, alphabetSize, starHeight, letterCount int

	flag.StringVar(&simplifier_path, "binpath", "", "regex simplifier binary path")
	flag.IntVar(&countRegex, "regex-n", 1, "number of regexes")
	flag.IntVar(&alphabetSize, "alphabet-size", 1, "alphabet size")
	flag.IntVar(&starHeight, "max-height", 1, "max star height")
	flag.IntVar(&letterCount, "max-len", 1, "max regex length")
	flag.Parse()

	if len(os.Args) == 1 {
		flag.Usage()
		return
	}
	if letterCount <= 0 {
		panic("max regex length cannot be <= 0")
	}
	if alphabetSize <= 0 {
		panic("alphabet size cannot be <= 0")
	}

	regGenerator, _ := reggen.New(countRegex, alphabetSize, starHeight, letterCount)

	err := benchmark.Start(regGenerator, simplifier_path, countWords, maxDumpSize)
	if err != nil {
		panic("error: " + err.Error())
	}
}
