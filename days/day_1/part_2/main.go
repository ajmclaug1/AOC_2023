package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	number_map := map[string]string{"one": "1", "two": "2", "three": "3", "four": "4", "five": "5",
		"six": "6", "seven": "7", "eight": "8", "nine": "9"}

	readFile, err := os.Open("data.txt")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	running_total := 0

	for fileScanner.Scan() {
		string_to_inspect := string(fileScanner.Text())
		string_to_inspect_list := strings.Split(string_to_inspect, "")
		var found = false
		var first_num = ""
		i := 0
		for !found {
			for k, v := range number_map {
				if strings.Contains(strings.Join(string_to_inspect_list[0:i], ""), k) || strings.Contains(strings.Join(string_to_inspect_list[0:i], ""), v) {
					first_num = v
					found = true
				}
			}
			i++
		}
		found = false
		var last_num = ""
		i = len(string_to_inspect)
		for !found {
			for k, v := range number_map {
				if strings.Contains(strings.Join(string_to_inspect_list[i:len(string_to_inspect)], ""), k) || strings.Contains(strings.Join(string_to_inspect_list[i:len(string_to_inspect)], ""), v) {
					last_num = v
					found = true
				}
			}
			i--
		}

		first_and_last := first_num + last_num
		number_to_add, _ := strconv.Atoi(first_and_last)
		running_total += number_to_add

		fmt.Println(number_to_add, running_total)

	}

}
