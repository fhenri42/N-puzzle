package main

import (
  "fmt"
)

func main () {
  fmt.Printf("Welecome to Npuzzl\n")
//  tab, len := parcing()
 tabGood :=[][]int{}
 tabx0 := []int{1, 2, 3}
 tabx1 := []int{8, 0, 4}
 tabx2 := []int{7, 6, 5}
 tabGood = append(tabGood, tabx0)
 tabGood = append(tabGood, tabx1)
 tabGood = append(tabGood, tabx2)

// 4 2 5
// 1 3 7
 //6 8 0


//5 3 2
//6 0 8
//4 1 7
 tab :=[][]int{}
 tab0 := []int{5, 3, 2}
 tab1 := []int{6, 0, 8}
 tab2 := []int{4, 1, 7}
 tab = append(tab, tab0)
 tab = append(tab, tab1)
 tab = append(tab, tab2)
 astar(tab, tabGood,3)
  return
}
