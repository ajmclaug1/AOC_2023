package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	colours_dict := map[string]int{"red": 0, "blue": 0, "green": 0}
	//user_input := os.Args
	var input = colours_dict
	input["red"] = 3
	input["green"] = 3
	input["blue"] = 3

	readFile, err := os.Open("data_example.txt")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)
	//listofgames := []string{}
	for fileScanner.Scan() {
		var new_dict = colours_dict
		game := strings.Split(fileScanner.Text(), ":")
		game[1] = strings.ReplaceAll(game[1], ";", ",")
		game_list := strings.Split(strings.TrimSpace(game[1]), ",") 
		for _, v := range game_list{
			fmt.Print(v)
			dict_list := strings.Split(v, " ")
			new_dict[dict_list[0]] += 3

		}
		//listofgames = append(listofgames, game[1])
		fmt.Println(new_dict)

	}
}
