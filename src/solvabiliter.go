package main

import (
//  "fmt"
)

func inversions(tab [][]int, len int) (int, int) {
  var doubleLen = len * len - 1
  tmp := make([]int, doubleLen)
  var t = 0
  var x = 0
  var y = 0
  var nbr = 0
  var blank = 0
  for x < len {
    y = 0
    for y < len {
      if tab[x][y] != 0 {
      tmp[t] = tab[x][y]
      t++
    } else {
       blank = t
     }
      y++
    }
    x++
  }
  t = 0
  x = 0
  for t < doubleLen {
    y = 0
    x = t
    for x < doubleLen  {
      if tmp[t] > tmp[x]{ y++ }
      x++
    }
    nbr = nbr + y
    t++
  }
  return nbr, blank

}

func solvabiliter(tab [][]int, goodTab [][]int, len int) bool  {

  tabInversion, blankPosition := inversions(tab, len)
  goodTabInversieon, blankPosition1 := inversions(goodTab,len)

  if len % 2 == 0 {

    tabInversion =  tabInversion + (blankPosition / len)
    goodTabInversieon = goodTabInversieon + (blankPosition1 / len)
  }

  if tabInversion % 2 == goodTabInversieon % 2 {
  return true
}
  return false
}
