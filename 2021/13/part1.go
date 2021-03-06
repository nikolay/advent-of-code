/*
 *   Copyright (c) 2021
 *   All rights reserved.
 */
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func expandMatrix(matrix [][]bool, width, height int) [][]bool {
	for len(matrix) < height {
		matrix = append(matrix, make([]bool, width))
	}
	for y := 0; y < height; y++ {
		for len(matrix[y]) < width {
			matrix[y] = append(matrix[y], false)
		}
	}
	return matrix
}

func printMatrix(matrix [][]bool) {
	for y := 0; y < len(matrix); y++ {
		for x := 0; x < len(matrix[y]); x++ {
			if matrix[y][x] {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	matrix := make([][]bool, 0)
	width, height := 0, 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) == 0 {
			break
		}
		coords := strings.Split(line, ",")
		x, _ := strconv.Atoi(coords[0])
		y, _ := strconv.Atoi(coords[1])
		if height < y+1 {
			height = y + 1
			matrix = expandMatrix(matrix, width, height)
		}
		if width < x+1 {
			width = x + 1
			matrix = expandMatrix(matrix, width, height)
		}
		matrix[y][x] = true
	}

	r := regexp.MustCompile(`^fold along (x|y)=(\d+)$`)

	folds := 0
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) == 0 {
			break
		}
		matches := r.FindStringSubmatch(line)
		switch matches[1] {
		case "y":
			foldY, _ := strconv.Atoi(matches[2])
			newHeight := max(foldY, height-1-foldY)
			newMatrix := expandMatrix([][]bool{}, width, newHeight)
			for y := 0; y < newHeight; y++ {
				topY, bottomY := foldY-(y+1), foldY+(y+1)
				for x := 0; x < width; x++ {
					newMatrix[newHeight-(y+1)][x] = topY >= 0 && matrix[topY][x] || bottomY < height && matrix[bottomY][x]
				}
			}
			height = newHeight
			matrix = newMatrix
		case "x":
			foldX, _ := strconv.Atoi(matches[2])
			newWidth := max(foldX, width-1-foldX)
			newMatrix := expandMatrix([][]bool{}, newWidth, height)
			for x := 0; x < newWidth; x++ {
				leftX, rightX := foldX-(x+1), foldX+(x+1)
				for y := 0; y < height; y++ {
					newMatrix[y][newWidth-(x+1)] = leftX >= 0 && matrix[y][leftX] || rightX < width && matrix[y][rightX]
				}
			}
			width = newWidth
			matrix = newMatrix
		}
		folds++
		if folds == 1 {
			count := 0
			for y := 0; y < height; y++ {
				for x := 0; x < width; x++ {
					if matrix[y][x] {
						count++
					}
				}
			}
			fmt.Println("Part 1", "=", count)
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Part 2", "=")
	printMatrix(matrix)
}
