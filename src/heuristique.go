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
  var misplaced = 1

  var tmp1 = getPos(nb, grid1, size)
  var tmp2 = getPos(nb, grid2, size)
  if (tmp1.x == tmp2.x && tmp1.y == tmp2.y) {
    misplaced = 0
  }
  return misplaced
}

func misplaced(open state, grid2[][]int) int {
  var res = 0
  var i = 0

  for i < open.len * open.len {
    res += isMisplaced(i, open.grid, grid2, open.len)
    i++
  }
  return res
}

/*
** row_column Heuristic
*/

func get_out_row_column(i int, grid1[][]int, grid2[][]int, size int) int {
  var p1 = getPos(i, grid1, size)
  var p2 = getPos(i, grid2, size)
  var out_row = 0
  var out_column = 0

  if (p1.x != p2.x) {
    out_row = 1
  }
  if (p1.y != p2.y) {
    out_column = 1
  }
  return out_column + out_row
}

func h_row_column(open state, grid2[][]int) int {
  var res = 0
  var i = 0
  for i < open.len * open.len {
    res += get_out_row_column(i, open.grid, grid2, open.len)
    i++
  }
  return res
}
