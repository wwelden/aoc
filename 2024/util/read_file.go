package util

import "os"

type Solution interface {
	Solve(string)
}

func ReadFile(path string) string {
	inputFile, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return string(inputFile)
}
