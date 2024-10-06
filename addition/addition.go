// Take an array of integers and return the sum of all even numbers
package main

import (
    "fmt"
)

func main() {
    sum := 0
    array_to_sort := []int{1,2,3,4,5}
    for i := 0; i < len(array_to_sort); i++ {
        if array_to_sort[i]%2 == 0 {
            sum += array_to_sort[i]
        }
    }
    fmt.Println("The sum of even elements in the array is", sum)
}