package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	user_input := os.Args
	input := map[string]int{"red": 0, "blue": 0, "green": 0}
	red, _ := strconv.Atoi(user_input[1])
	green, _ := strconv.Atoi(user_input[2])
	blue, _ := strconv.Atoi(user_input[3])
	input["red"] = red
	input["green"] = green
	input["blue"] = blue
	running_total := 0

	readFile, err := os.Open("data.txt")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		game := strings.Split(fileScanner.Text(), ":")
		edit_draws := strings.ReplaceAll(game[1], ";", ",")

		possible := find_impossible(input, edit_draws)
		if possible {
			number_to_add, _ := strconv.Atoi(strings.ReplaceAll(game[0], "Game ", ""))
			running_total += number_to_add
		}

	}
	fmt.Println(running_total)

}

func find_impossible(input map[string]int, draws string) bool {
	draws_list := strings.Split(draws, ",")
	for _, v := range draws_list {
		draw_value := strings.Trim(v, " ")
		draw_value_list := strings.Split(draw_value, " ")
		no_of_dice, _ := strconv.Atoi(draw_value_list[0])
		if input[draw_value_list[1]] < no_of_dice {
			return false
		}
	}

	return true

}
