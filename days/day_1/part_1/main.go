package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {

	readFile, err := os.Open("data.txt")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	running_total := 0
	for fileScanner.Scan() {
		re := regexp.MustCompile("[0-9]+")
		numbers := re.FindAllString(string(fileScanner.Text()), -1)
		numberstr := strings.Join(numbers, "")
		numberlist := strings.Split(numberstr, "")
		numberstring := numberlist[0] + numberlist[len(numberlist)-1]
		number, _ := strconv.Atoi(numberstring)
		running_total += number
		fmt.Println(running_total)
	}

	readFile.Close()
}
