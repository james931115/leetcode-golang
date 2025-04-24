package main

import "fmt"

// TwoSum function
func twoSum(nums []int, target int) []int {
    m := make(map[int]int)
    for i, num := range nums {
        if index, found := m[target-num]; found {
            return []int{index, i}
        }
        m[num] = i
    }
    return nil
}

// main function to run the program
func main() {
    fmt.Println(twoSum([]int{2, 7, 11, 15}, 9)) // Output: [0 1]
}
