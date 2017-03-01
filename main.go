package main

import (
  "fmt"
  "io/ioutil"
  "strings"
  "strconv"
  "os"
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
  length, _ := strconv.Atoi(str[1])
  //fmt.Printf("%d\n", length)
  tabFinal := finalTab(length)
  var i = 0
  var j = 0
  var l = 0
  var k = 2
  tabInit := [][]int{}
  for i < length {
    //fmt.Print(str[k])
    strFirstSplit := strings.Split(str[k], "#")
    strSplit := strings.Split(strFirstSplit[0], " ")
    var tabLigne = make([]int, length)
    taille := len(strSplit)
    for j < taille {
      if strSplit[j] != "" {
        nb,err := strconv.Atoi(strSplit[j])
        if err != nil {
          fmt.Print("File is not well formated3\n")
          os.Exit(3)
        }
        tabLigne[l] = nb
        l++
      }
        j++
    }
    if l != length {
      fmt.Print("File is not well formated2\n")
      os.Exit(3)
    }
    l = 0
    j = 0
    tabInit = append(tabInit, tabLigne)
    k++
    i++
  }
  return tabInit, tabFinal, length
}

func main () {
    dat, err := ioutil.ReadFile("/tmp/npuzzle-4-u-test.txt")
    if err != nil {
      fmt.Print("No such file or directory")
      os.Exit(3)
    }
    check(err)

    datString := string(dat)
    datSplitted := strings.Split(datString, "\n")
    _, erreur := strconv.Atoi(datSplitted[1]);
    if datSplitted != nil && datSplitted[0][0] != '#' || erreur != nil {
        fmt.Printf("File is not well formated\n")
        os.Exit(3)
    }

    tabInit, tabFinal, len := parsing(datSplitted)
    fmt.Printf("\n%d\n", len)
    fmt.Printf("\nFirst:\n%d\n%d\n%d\n", tabInit[0][0], tabInit[0][1], tabInit[0][2])
    fmt.Printf("\nSecond:\n%d\n%d\n%d\n", tabInit[1][0], tabInit[1][1], tabInit[1][2])
    fmt.Printf("\nThird:\n%d\n%d\n%d\n", tabInit[2][0], tabInit[2][1], tabInit[2][2])

    fmt.Printf("\nFirst:\n%d\n%d\n%d\n", tabFinal[0][0], tabFinal[0][1], tabFinal[0][2])
    fmt.Printf("\nSecond:\n%d\n%d\n%d\n", tabFinal[1][0], tabFinal[1][1], tabFinal[1][2])
    fmt.Printf("\nThird:\n%d\n%d\n%d\n", tabFinal[2][0], tabFinal[2][1], tabFinal[2][2])
    os.Exit(3)
//  astar(tab, len)
  return
}
