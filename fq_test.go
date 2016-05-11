package main

import (
	"bufio"
	"os"
	"strings"
	"strconv"
)

func readLine(path string) string {
	inFile, _ := os.Open(path)
	defer inFile.Close()
	scanner := bufio.NewScanner(inFile)
	scanner.Split(bufio.ScanLines)

	line := 0
	pos := 1

	lastGood := 0

	for scanner.Scan() {

		line++

		text := scanner.Text()

		switch pos {
		case 1:
			if (!strings.HasPrefix(text, "@")) {
				return "scan failed at line " + strconv.Itoa(line) + " should be @ but got " + string(text[0]) + ", last good read ended at " + strconv.Itoa(lastGood * 4)
			}

		case 3:

			if (!strings.HasPrefix(text, "+")) {
				return "scan failed at line " + strconv.Itoa(line) + " should be + but got " + string(text[0]) + ", last good read ended at " + strconv.Itoa(lastGood * 4)
			}
		}

		if (pos == 4) {
			lastGood++
			pos = 1
		} else {
			pos++
		}
	}
	return "unknown out"
}

func main() {
	argsWithoutProg := os.Args[1]
	println(readLine(argsWithoutProg));
}
