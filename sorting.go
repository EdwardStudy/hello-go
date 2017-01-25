package main

import "fmt"
import "sort"

func main(){
  strs := []string{"c", "b", "a"}
  sort.Strings(strs)
  fmt.Println("Strings:", strs)

  ints := []int{3, 4, 1}
  sort.Ints(ints)
  fmt.Println("Integers: ", ints)

  isSorted := sort.IntsAreSorted(ints)
  fmt.Println("Sorted: ", isSorted)
}