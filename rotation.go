
package main

import ()

func makeASpare(tab[][]int) [][]int {

  tmp := make([][]int,len(tab))

  var x = 0
  for x < len(tab) {
    var y = 0
    tmp[x] = make([]int,len(tab))
    for y < len(tab) {

      tmp[x][y] = tab[x][y]
      y++
    }
    x++
  }
  return tmp
}

func deplacementUp(tab [][]int, cur Noeu, parent Noeu)  [][]int {

  tmp := makeASpare(tab)
  if (cur.x - 1 < 0 || parent.x  == cur.x - 1)  {
    tmp[0][0] = -4
    return  tmp
  }
  tmp[cur.x][cur.y] = tmp[cur.x - 1][cur.y]
  tmp[cur.x - 1][cur.y] = 0
  return tmp
}

func deplacementRight(tab [][]int, cur Noeu, parent Noeu,len int)  [][]int{

  tmp := makeASpare(tab)
  if (cur.y + 1 >= len || parent.y  == cur.y + 1)  {
    tmp[0][0] = -4
    return  tmp
  }
  tmp[cur.x][cur.y] = tmp[cur.x][cur.y + 1]
  tmp[cur.x][cur.y + 1] = 0
  return tmp
}

func deplacementLeft(tab [][]int, cur Noeu, parent Noeu)  [][]int{

  tmp := makeASpare(tab)
  if (cur.y - 1 < 0 || parent.y == cur.y - 1)  {
    tmp[0][0] = -4
    return tmp
  }
  tmp[cur.x][cur.y] = tmp[cur.x][cur.y - 1]
  tmp[cur.x][cur.y - 1] = 0

  return tmp
}

func deplacementDown(tab [][]int, cur Noeu, parent Noeu,len int)  [][]int{

  tmp := makeASpare(tab)
  if (cur.x + 1 >= len || parent.x  == cur.x + 1)  {
    tmp[0][0] = -4
    return  tmp
  }
  tmp[cur.x][cur.y] = tmp[cur.x + 1][cur.y]
  tmp[cur.x + 1][cur.y] = 0
  return tmp
}

func MoveGrid(best state) ([][][]int, int) {
  maxTab := make([][][]int, 5)
  var t = 0
  var count = 0
  tmp := best.pos

  for t < 4 {

    if t == 0 {
      a := deplacementUp(best.grid, tmp, best.parent)
      if a[0][0] != -4 {
        maxTab[count] = a
        count++
      }
      } else if  t == 1 {
        b := deplacementRight(best.grid, tmp, best.parent, best.len)
        if b[0][0] != -4 {
          maxTab[count] = b
          count++
        }
        } else  if t == 2  {
          c := deplacementDown(best.grid, tmp, best.parent, best.len)
          if c[0][0] != -4 {
            maxTab[count] = c
            count++
          }
          } else if t == 3 {
            d := deplacementLeft(best.grid, tmp, best.parent)
            if d[0][0] != -4 {
              maxTab[count] = d
              count++
            }
          }
          t++
        }
        return maxTab, count
      }
