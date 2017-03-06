package main

import (
  "fmt"
  "strings"
  "strconv"
  "io/ioutil"
  "os"
)

func  createGood(len int) [][]int {
  var x = 0
  var y = 0
  var t = 1
  goodTab := make([][]int, len)
  simpleTab := make([][]int, len)

//TODO reduire le tableau a chaque tour

  for x < len {
    y = 0
    simpleTab[x] = make([]int, len)
    for y < len {
      simpleTab[x][y] = t
      y++
      t++
    }
    x++
  }
  simpleTab[len - 1][len - 1] = 0
  x = 0
  fmt.Printf("simpleTab = %d",simpleTab)

  for x < len {
    var y = 0
    goodTab[x] := make([]int, len)
    if x == ) {
      goodTab = append(goodTab, simpleTab[0])
    }
    if y == len - 1 {
      goodTab[x][len - 1] = simpleTab[2][x]
    }
    if x == len - 1 {
      goodTab[len][y] = simpleTab[]
    }
    x++
  }
  return goodTab
}

func main () {

  if 2 == len(os.Args) {
    dat, err := ioutil.ReadFile(os.Args[1])
    if err != nil {
      fmt.Print("No such file or directory\n")
      os.Exit(3)
    }
    check(err)

    datString := string(dat)
    datSplitted := strings.Split(datString, "\n")
    check, erreur := strconv.Atoi(datSplitted[1]);
    if datSplitted != nil && datSplitted[0][0] != '#' || erreur != nil || len(datSplitted) != check + 3 {
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

    tabGood := createGood(len)
    return
    // tabx0 := []int{1, 2, 3, 4}
    // tabx1 := []int{11, 13, 14, 5}
    // tabx2 := []int{11, 0, 15, 6}
    // tabx3 := []int{10, 9, 8, 7}
    // tabGood = append(tabGood, tabx0)
    // tabGood = append(tabGood, tabx1)
    // tabGood = append(tabGood, tabx2)
    // tabGood = append(tabGood, tabx3)

    if solvabiliter(tabInit, tabGood, 4) == false{
      fmt.Printf("Soory this is not a solavalble puzzle\n")
      os.Exit(3)
    }
    //    astar(tabInit, tabGood,  3)
    os.Exit(3)
    } else {
      fmt.Print("You need only one args.\n")
    }
    //  tab, len := parcing()

    return
  }
