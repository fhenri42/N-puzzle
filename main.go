package main

import (
  "fmt"
  "io/ioutil"
  "strings"
  "strconv"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func finalTab(len int) ([][]int) {
  var i = 0
  var j = 0
  tabFinal := [][]int{}

  for i < len {
    var tabSemiFinal = make([]int, len)
    for j < len {
      tabSemiFinal[j] = 0
      j++
    }
    tabFinal = append(tabFinal, tabSemiFinal)
    i++
  }
  i = 0
  j = 0
  var k = 0
  var count = 1
  tabFinal[i][j] = 1
  for k < len * len - 2 {
    count += 1
    if j < len - 1 && tabFinal[i][j + 1] == 0 && tabFinal[i + 1][j + 1] == 0 {
      j++
      tabFinal[i][j] += count
    } else if i < len - 1 && tabFinal[i + 1][j] == 0 {
      i++
      tabFinal[i][j] += count
    } else if j > 0 && tabFinal[i][j - 1] == 0 {
      j--
      tabFinal[i][j] += count
    } else if i > 0 && tabFinal[i - 1][j] == 0 {
      i--
      tabFinal[i][j] += count
    }
    //fmt.Printf("\nTOUR\n%d\n%d", i, j)
    k++
  }
  return tabFinal
}

func parsing(str []string) ([][]int, [][]int, int) {
  len, _ := strconv.Atoi(str[1])
  //fmt.Printf("%d\n", len)
  tabFinal := finalTab(len)
  var i = 0
  var j = 0
  var l = 0
  var k = 2
  tabInit := [][]int{}
  for i < len {
    //fmt.Print(str[k])
    strSplit := strings.Split(str[k], " ")
    var tabLigne = make([]int, len)
    for l < len {
      if strSplit[j] != "" {
        nb,_ := strconv.Atoi(strSplit[j])
        tabLigne[l] = nb
        l++
      }
        j++
    }
    l = 0
    j = 0
    tabInit = append(tabInit, tabLigne)
    k++
    i++
  }


  return tabInit, tabFinal, len
}

func main () {
    dat, err := ioutil.ReadFile("/tmp/npuzzle-12.txt")
    check(err)
    datString := string(dat)
    //fmt.Print(datString)
    datSplitted := strings.Split(datString, "\n")
    tabInit, tabFinal, len := parsing(datSplitted)
    fmt.Printf("\n%d\n", len)
    fmt.Printf("\nFirst:\n%d\n%d\n%d\n", tabInit[0][0], tabInit[0][1], tabInit[0][2])
    fmt.Printf("\nSecond:\n%d\n%d\n%d\n%d\n", tabInit[1][0], tabInit[1][1], tabInit[1][2], tabInit[1][3])
    fmt.Printf("\nThird:\n%d\n%d\n%d\n", tabInit[2][0], tabInit[2][1], tabInit[2][2])

    fmt.Printf("\nFirst:\n%d\n%d\n%d\n", tabFinal[0][0], tabFinal[0][1], tabFinal[0][2])
    fmt.Printf("\nSecond:\n%d\n%d\n%d\n", tabFinal[10][0], tabFinal[9][0], tabFinal[8][0])
    fmt.Printf("\nThird:\n%d\n%d\n%d\n", tabFinal[11][11], tabFinal[11][10], tabFinal[11][0])
//  astar(tab, len)
  return
}
