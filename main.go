package main

import (
  "fmt"
  "io/ioutil"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

// func parsing(str []string) ([][]int, [][]int, len int) {
// len = str[1]
// fmt.Printf("%d\n", len)
// }

func main () {
  fmt.Printf("Welecome to Npuzzl")
    dat, err := ioutil.ReadFile("/tmp/npuzzle-3.txt")
    check(err)
    fmt.Print("%s",dat)
//  tabInit, tabFinal, len := parsing()
//  astar(tab, len)
  return
}
