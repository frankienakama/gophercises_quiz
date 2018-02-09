package main

import "flag" // specify options for command-line programs

func main() {
	csvFilename := flag.String("csv", "problems.csv", "a csv file in the format of a 'question,answer'")
	flag.Parse()
	_ = csvFilename
}
