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
    //   fmt.Printf("blank = %d\n", blank)

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

  //fmt.Printf("%d\n", tmp)
  return nbr, blank

}

func solvabiliter(tab [][]int, goodTab [][]int, len int) bool  {

  tabInversion, blankPosition := inversions(tab, len)
  goodTabInversieon, blankPosition1 := inversions(goodTab,len)

  if len % 2 == 0 {

    tabInversion =  tabInversion + (blankPosition / len)
    goodTabInversieon = goodTabInversieon + (blankPosition1 / len)

  //  fmt.Printf("%d\n", tabInversion)
  //  fmt.Printf("%d\n", goodTabInversieon)
  }

  if tabInversion % 2 == goodTabInversieon % 2 {
//    fmt.Printf("solulebe")
  return true
}
//fmt.Printf("not solulebe")
  return false
}

//If the grid width is odd, then the number of inversions in a solvable situation is even.
//If the grid width is even, and the blank is on an even row counting from the bottom (second-last, fourth-last etc), then the number of inversions in a solvable situation is odd.
//If the grid width is even, and the blank is on an odd row counting from the bottom (last, third-last, fifth-last etc) then the number of inversions in a solvable situation is even.
