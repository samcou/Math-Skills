package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

// ReadNums reads a file and returns a slice of numbers
func ReadNums(filename string) (nums []float64, err error) {
	// Open the file based off argument
	x, err := os.ReadFile(os.Args[1])
	// If there is an error, return it
	if err != nil {
		return nil, err
	}
	// Split the file into a slice of strings
	lines := strings.Split(string(x), "\n")
	// Make a cap to avoid resizing array on evert append
	nums = make([]float64, 0, len(lines))
	for _, l := range lines {
		// empty line appears at the end of the file when we use Split
		if len(l) == 0 {
			continue
		}
		// conv strings to float64
		n, err := strconv.ParseFloat(l, 64)
		if err != nil {
			return nil, err
		}
		nums = append(nums, n)
	}
	return nums, nil
}

func Average(arr []float64) float64 {
	var result float64
	result = 0
	// var LenArr = float64(len(arr))
	for _, i := range arr {
		result += i
	}
	return result / float64(len(arr))
}

// Median returns the median of a slice of numbers
func Median(arr []float64) float64 {
	// make a copy of the array as to not mess up the initial data
	arrCopy := make([]float64, len(arr))
	copy(arrCopy, arr)
	sort.Float64s(arrCopy)
	var median float64
	len := len(arrCopy)
	if len == 0 {
		return 0
	} else if len%2 == 0 {
		median = (arrCopy[len/2-1] + arrCopy[len/2]) / 2
	} else {
		median = arrCopy[len/2]
	}
	return median
}

func Variance(nums []float64) float64 {
	var sum float64
	var mean float64
	var variance float64
	for _, x := range nums {
		sum += x
	}
	mean = sum / float64(len(nums))
	for _, x := range nums {
		variance += (x - mean) * (x - mean)
	}
	return variance / float64(len(nums))
}

func main() {
	nums, err := ReadNums(os.Args[1])
	if err != nil {
		panic(err)
	}
	fmt.Println("Average:", math.Round(Average(nums)))
	fmt.Println("Median:", math.Round(Median(nums)))
	fmt.Println("Variance:", int(math.Round(Variance(nums))))
	fmt.Println("Standard Deviation:", math.Round(math.Sqrt(Variance(nums))))
}
