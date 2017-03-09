package main

import (
  "fmt"
  "strings"
  "strconv"
  "os"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func checkTabInit(tabInit [][]int, len int) {
  var i = 0
  var j = 0
  var k = 0
  for i < len {
    for j < len {
      if tabInit[i][j] >= len * len {
        fmt.Print("\n\033[1;31mFile is not well formated\033[m\n")
        os.Exit(3)
      }
      tmp := tabInit[i][j]
      l := i
      m := j
      for l < len {
          for m < len {
            if tabInit[l][m] == tmp {
              k += 1
            }
            if k > 1 {
              fmt.Print("\n\033[1;31mFile is not well formated\033[m\n")
              os.Exit(3)
            }
            m++
          }
          m = 0
        l++
      }
      k = 0
      j++
    }
    j = 0
    i++
  }
}


func parsing(str []string) ([][]int, int) {
  length, _ := strconv.Atoi(str[1])
  var i = 0
  var j = 0
  var l = 0
  var k = 2
  tabInit := [][]int{}
  
  for i < length {
    strFirstSplit := strings.Split(str[k], "#")
    strSplit := strings.Split(strFirstSplit[0], " ")
    var tabLigne = make([]int, length)
    taille := len(strSplit)
    for j < taille {
      if strSplit[j] != "" {
        nb,err := strconv.Atoi(strSplit[j])
        if err != nil {
          fmt.Print("\n\033[1;31mFile is not well formated\033[m\n")
          os.Exit(3)
        }
        tabLigne[l] = nb
        l++
      }
        j++
    }
    if l != length {
      fmt.Print("\n\033[1;31mFile is not well formated\033[m\n")
      os.Exit(3)
    }
    l = 0
    j = 0
    tabInit = append(tabInit, tabLigne)
    k++
    i++
  }
  checkTabInit(tabInit, length)
  return tabInit, length
}
