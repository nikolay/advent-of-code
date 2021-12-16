package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	last := -1
	count := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		if last != -1 && num > last {
			count++
		}
		last = num
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(count)
}