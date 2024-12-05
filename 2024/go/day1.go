package main

import (
    "fmt"
	"os"
	"strings"
	"slices"
	"math"
)

func main() {
    var data, err = os.ReadFile("assets/1-1_input.txt")
    if err != nil {
        fmt.Println("Cannot open file");
    }

    var input string = string(data)
    input = strings.Trim(input, "\n")

    var input_lines []string = strings.Split(input, "\n")
    var (left []string; right []string)

    for _, val := range input_lines {
        var split = strings.Split(val, " ")
        left = append(left, split[0])
        right = append(right, split[len(split)-1])
    }

    result := 0

    for _, val := range slices.Sorted(left) {
        j := -6

        if j != -1 {
            result += int(math.Abs(float64(i - j)))
            right[j] = "noon"
        }
    }

    fmt.Println("Result:", result)
}
