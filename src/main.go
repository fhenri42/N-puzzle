package main

import (
  "fmt"
  "strings"
  "strconv"
  "io/ioutil"
  "os"
)
func  createGood( len int)[][]int  {
  doubleLen := len * len
  var cur = 1
  var x = 0
  var ix = 1
  var y = 0
  var iy = 0
  simpleTab := make([]int, doubleLen)
  for x < doubleLen {
    simpleTab[x] = -1
    x++
  }

  x = 0
  var continu = true

  for continu == true {
    simpleTab[x + y * len] = cur
    if cur == 0 {
      continu = false
    }
    cur += 1
    if x + ix == len || x + ix < 0 || (ix != 0 && simpleTab[x + ix + y * len] != -1) {
      iy = ix
      ix = 0
      } else if y + iy == len || y + iy < 0 || (iy != 0 && simpleTab[x + (y + iy) * len] != -1) {
        ix = -iy
        iy = 0
      }
      x += ix
      y += iy
      if cur == doubleLen { cur = 0 }
    }

    goodTab := make([][]int, len)
    x = 0
    var  t = 0
    for x < len {
      goodTab[x] = make([]int, len)
      y = 0
      for y < len {
        goodTab[x][y] = simpleTab[t]
        t++
        y++
      }
      x++
    }
    return goodTab
  }

  func main () {

    if 2 == len(os.Args) {
      dat, err := ioutil.ReadFile(os.Args[1])
      if err != nil {
        fmt.Print("\n\033[1;31mNo such file or directory\033[m\n")
        os.Exit(3)
      }
      check(err)

      var i int
      fmt.Print("\n\n\033[1;32mWelcome to Npuzzle !\033[m\n\n\033[1;36mChoose the heuristic that you wish to use.\033[m\n\n\033[1;33mPress 1 to use Manhattan\033[m\n\033[1;33mPress 2 to use Misplaced\033[m\n\033[1;33mPress 3 to use Get column\033[m\n\n\033[1;32mYour choice:(1, 2, 3) => \033[m")
    _, err2 := fmt.Scanf("%d", &i)

    if err2 != nil {
      fmt.Print("\n\033[1;31mSorry, this choice is not possible.\033[m\n")
      os.Exit(3)
    }
    if i < 1 || i > 3 {
      fmt.Print("\n\n\033[1;36mSorry, You can only pick those 3 choices:\033[m\n\n\033[1;33mPress 1 to use Manhattan\033[m\n\033[1;33mPress 2 to use Misplaced\033[m\n\033[1;33mPress 3 to use Get column\033[m\n\n\033[1;32mYour choice:(1, 2, 3) => \033[m")
      _, err3 := fmt.Scanf("%d", &i)
      if err3 != nil || i < 1 || i > 3 {
        fmt.Print("\n\033[1;31mSorry, this choice is not possible.\033[m\n")
        os.Exit(3)
      }
    }

      datString := string(dat)
      datSplitted := strings.Split(datString, "\n")
      check, erreur := strconv.Atoi(datSplitted[1]);
      if datSplitted != nil && datSplitted[0][0] != '#' || erreur != nil || len(datSplitted) != check + 3 {
        fmt.Printf("\n\033[1;31mFile is not well formated\033[m\n")
        os.Exit(3)
      }
      tabInit, len := parsing(datSplitted)
      fmt.Printf("\nlen of: %d\n", len)
      var x = 0
      for x < len {
        var y = 0
        fmt.Printf("[")
        for y < len {
          fmt.Printf(" %d ",tabInit[x][y])
          y++
        }
        fmt.Printf("]\n")
        x++
      }

      tabGood := createGood(len)
      if solvabiliter(tabInit, tabGood, len) == false{
        fmt.Printf("\n\033[1;31mSorry this is not a solvable puzzle\033[m\n\n")
        os.Exit(3)
      }
      astar(tabInit, tabGood,  len, i)
      os.Exit(3)
      } else {
        fmt.Print("\n\033[1;31mYou need only one args.\033[m\n")
      }
      return
    }
