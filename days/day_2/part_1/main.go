package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	colours_dict := map[string]string{"red": "0", "blue": "0", "green": "0"}
	user_input := os.Args
	input := colours_dict
	input["red"] = user_input[1]
	input["green"] = user_input[2]
	input["blue"] = user_input[3]

	readFile, err := os.Open("data_example.txt")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		fmt.Print(fileScanner.Text())
	}
}
