package main

import (
  "math"
)

func getPos(nb int, tab[][]int, len int) Noeu  {
  var x = 0
  var tmp Noeu
  for  x < len {
    var y = 0
    for y < len {
      if tab[x][y] == nb {
        tmp.x = x
        tmp.y = y
        return tmp
      }
      y++
    }
    x++
  }
  tmp.y = 0
  tmp.x = 0
  return tmp
}

/*
** Manhattan Heuristic
*/

func getDist(nb int, open state, goodTab[][]int)  int{
  tmp1 := getPos(nb, open.grid, open.len)
  tmp2 := getPos(nb, goodTab, open.len)
  H := math.Abs(float64(tmp2.x) - float64(tmp1.x)) + math.Abs(float64(tmp2.y) - float64(tmp1.y))
  return int(H)
}

func calculeManhattan(open state, goodTab[][]int) int {

  var x = 0
  var res = 0
  for x < len(goodTab) * len(goodTab) {
    res += getDist(x, open, goodTab)
    x++
  }
  return res
}


/*
** Misplaced Heuristic
*/

func isMisplaced(nb int, grid1[][]int, grid2[][]int, size int) int {
  var tmp Noeu
  var tmp2 Noeu
  var misplaced = 1

  tmp1 = getPos(nb, grid1, size)
  tmp2 = getPos(nb, grid2, size)
  if (tmp1.x == tmp2.x && tmp1.y == tmp2.y) {
    misplaced = 0
  }
  return misplaced
}

func misplaced(grid1[][]int, grid2[][]int, size int) int {
  var res = 0
  var i = 0

  while (i < size * size)
  {
    res += isMisplaced(i, grid1, grid2, size)
    i++
  }
  return res
}
