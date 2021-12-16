package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	var days int
	switch os.Args[1] {
	case "1":
		days = 80
	case "2":
		days = 256
	default:
		log.Fatal(errors.New(fmt.Sprintf("unknown part number %v", os.Args[1])))
	}

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var line string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line = strings.TrimSpace(scanner.Text())
		if len(line) > 0 {
			break
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	parts := strings.Split(line, ",")
	fish := make(map[int]int)
	for _, p := range parts {
		num, err := strconv.Atoi(strings.TrimSpace(p))
		if err != nil {
			log.Fatal(err)
		}
		fish[num] += 1
	}

	for day := 0; day < days; day++ {
		moms := fish[0]
		for i := 0; i < 8; i++ {
			fish[i] = fish[i+1]
		}
		fish[6] += moms
		fish[8] = moms
	}

	count := 0
	for i := 0; i <= 8; i++ {
		count += fish[i]
	}
	fmt.Println(count)
}