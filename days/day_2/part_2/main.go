package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	total := 0

	readFile, err := os.Open("data.txt")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		game := strings.Split(fileScanner.Text(), ":")
		edit_draws := strings.ReplaceAll(game[1], ";", ",")
		total += find_min_cubes(edit_draws)

	}
	fmt.Println(total)

}

func find_min_cubes(draws string) int {
	input := map[string]int{"red": 0, "blue": 0, "green": 0}
	draws_list := strings.Split(draws, ",")
	for _, v := range draws_list {
		draw_value := strings.Trim(v, " ")
		draw_value_list := strings.Split(draw_value, " ")
		no_of_dice, _ := strconv.Atoi(draw_value_list[0])
		if input[draw_value_list[1]] < no_of_dice {
			input[draw_value_list[1]] = no_of_dice
		}
	}

	return input["red"] * input["blue"] * input["green"]

}
