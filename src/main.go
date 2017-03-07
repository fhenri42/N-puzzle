package main

import (
  "fmt"
  "strings"
  "strconv"
  "io/ioutil"
  "os"
)
//TODO a comprendre
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
        fmt.Printf("\nSoory this is not a solavalble puzzle\n")
        os.Exit(3)
      }
      astar(tabInit, tabGood,  3)
      os.Exit(3)
      } else {
        fmt.Print("You need only one args.\n")
      }
      //  tab, len := parcing()

      return
    }
