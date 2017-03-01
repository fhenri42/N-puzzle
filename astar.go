package main

import (
  "fmt"
  "math"
)

type Heuristique struct {
  x int
  y int
  count int
  heuristioque int

}

type Noeu struct {
  x int
  y int
}

type Heu struct {
  H int
  G int
}

type state struct {
  grid [][]int
  h int
  g int
  index int
  len int
  look bool
  //  parent *state
  pos Noeu
}

func findBest(sta []state) state  {

  b := sta[0]

  for _, s := range sta {
    if s.h + s.g < b.h + b.g { b = s }
  }
  return b
}

func deplacementUp(tab [][]int, cur Noeu )  [][]int{
  tab[cur.x][cur.y] = tab[cur.x - 1][cur.y]
  tab[cur.x - 1][cur.y] = 0
  return tab
}

func deplacementRight(tab [][]int, cur Noeu )  [][]int{
  tab[cur.x][cur.y] = tab[cur.x][cur.y + 1]
  tab[cur.x][cur.y + 1] = 0
  return tab
}

func deplacementLeft(tab [][]int, cur Noeu )  [][]int{
  tab[cur.x][cur.y] = tab[cur.x][cur.y - 1]
  tab[cur.x][cur.y - 1] = 0
  return tab
}

func deplacementDown(tab [][]int, cur Noeu )  [][]int{
  tab[cur.x][cur.y] = tab[cur.x + 1][cur.y]
  tab[cur.x + 1][cur.y] = 0
  return tab
}
// state->parent && state->parent->pos->i == state->pos->i - 1

func MoveGrid(best state) ([][][]int, int) {
  maxTab := make([][][]int, 5)
  var t = 0
  var count = 0
  tmp := best.pos
  //fmt.Printf("parent = %d\n", best.parent)
  for t < 5 {

    if t == 0 && tmp.x - 1 > 0 {
      a := deplacementUp(best.grid, tmp)
      maxTab[count] = a
      count++
    }
    if  t == 1  && tmp.y + 1 < best.len {
      b := deplacementRight(best.grid, tmp)
      maxTab[count] = b
      count++
    }
    if t == 2  && tmp.x + 1 < best.len{
      c := deplacementDown(best.grid, tmp)
      maxTab[count] = c
      count++
    }
    if t == 3  && tmp.y - 1 > 0{
      d := deplacementLeft(best.grid, tmp)
      maxTab[count] = d
      count++
    }
    t++

  }
  return maxTab, count
}


func inList(aList []state, bList [][]int)  state {

  for _,list := range aList {
    var x = 0
    for x < len(list.grid) {
      var y = 0
      for y < len(list.grid[x]) {
        if list.grid[x][y] != bList[x][y] {
          list.look = true
          fmt.Printf("in the true")
          return list
        }
        y++
      }
      x++
    }
  }
  aList[0].look = false
  return aList[0]

}
func checKGood( tab [][]int, goodTab[][]int, len int) bool  {
  var x = 0
  for x < len {
    var y = 0
    for y < len {
      if tab[x][y] != goodTab[x][y] { return false }
      y++
    }
    x++
  }
  return true
}
func shouldBe(tab [][]int, len int) Noeu  {
  var x = 0
  var y = 0
  var tmp Noeu
  for x < len {
    for y < len {
      if tab[x][y] == 0 {
        tmp.x = x
        tmp.y = y
        return tmp
      }
      y++
    }
    x++
  }
  return tmp
}

func getPos(t int, tab[][]int, len int) Noeu  {
     var x = 0
     var tmp Noeu
     for  x < len {
       var y = 0
       for y < len {
          if tab[x][y] == t {
            tmp.x = x
            tmp.y = y
            return tmp
          }
          y++
       }
       x++
     }
     return tmp
}

func getDist(x int, open state, goodTab[][]int)  int{
tmp1 := getPos(x, open.grid, open.len)
tmp2 := getPos(x, goodTab, open.len)
H := math.Abs(float64(tmp2.x) - float64(tmp1.x)) + math.Abs(float64(tmp2.y) - float64(tmp1.y))
return int(H)
}

func calculeManhattan(open state, goodTab[][]int) int {

  var x = 0
  var res = 0
  for x < len(goodTab[x]) {
    res += getDist(x, open, goodTab)
    x++
  }


  return res
}

func astar(tab [][]int, goodTab [][]int, len int) int  {

  var cal = 0
  openList := make([]state, 1)
  openList[0].grid  = tab
  openList[0].h = 0
  openList[0].g = 0
  openList[0].len = len
  openList[0].index = 0
  openList[0].pos.x = 0
  openList[0].pos.y = 0
  openList[0].look = false
  var succes = false

  closeList := make([]state, 1)

  //  aff(tab, len)
  //for succes != true {
  for _, open := range openList {
    fmt.Printf("%d\n",open)
    best := findBest(openList)
    fmt.Printf("best = %d\n",best.index)
    if checKGood(best.grid, goodTab, len) == true {
      succes = true
      } else {
        //  openList = append(openList[:0], openList[1:])
        closeList = append(closeList, best)
        move, cr := MoveGrid(best)
        var l = 0
        for l < cr {

          fmt.Printf("move = %d\n",move[l])
          find := inList(openList, move[l])

          if find.look == false  {
            var newList state
            tmp := shouldBe(best.grid, best.len)
            fmt.Printf("tmp = %d\n",tmp)
            newList.grid = open.grid
            //TODO change h + g
            newList.h = calculeManhattan(open, goodTab)
            newList.g = open.g + 1
            newList.len = open.len
            newList.look = false
            newList.index = open.index + 1
            newList.pos.x = tmp.x
            newList.pos.y = tmp.y
            openList = append(openList, newList)

            } else {
              fmt.Printf("maison \n")
              if find.g + find.h > find.h + best.g + 1 {
                find.g = find.h + best.g + 1
                fmt.Printf("maison1 \n")
                find2 :=  inList(closeList, find.grid)
                if find2.look == false {
                  fmt.Printf("maison2 \n")
                  closeList = append(closeList[:0], closeList[find.index])
                  openList = append(openList, find)
                }
              }
            }
            l++
          }
        }
        fmt.Printf("h = %d\n", open.h)
      }
      //}
      fmt.Printf("nb of calcule = %d, %d\n", cal, succes)
      return 0
    }



    func  aff(tab [][]int, len int)  {
      var x = 0
      fmt.Printf("====================\n")
      for x < len {
        fmt.Printf("%d\n",tab[x])
        x++
      }
      fmt.Printf("====================\n")

    }



    // func astarApllication(tab [][]int, at Noeu, togo Noeu) ([]Noeu, int)  {
    //   var tmp Noeu
    //   openList := []Noeu{}
    //   var deplacement = 1
    //   tmp.x = at.x
    //   tmp.y = at.y
    //   for (tmp.x  != togo.x || tmp.y != togo.y)  {
    //     a := calculeManhattan(deplacement, tmp.x + 1, tmp.y, tmp)
    //     b := calculeManhattan(deplacement, tmp.x - 1, tmp.y, tmp)
    //     c := calculeManhattan(deplacement, tmp.x, tmp.y + 1, tmp)
    //     d := calculeManhattan(deplacement, tmp.x, tmp.y - 1, tmp)
    //     fmt.Printf("a == %d\nb == %d\nc == %d\nd == %d\n=== \n", a, b, c, d)
    //     if (a <= b  && a <= c && a <= d && a >= 0 && tmp.x != togo.x) {
    //       //    fmt.Printf("A")
    //       tmp.x += 1
    //       openList = append(openList, tmp)
    //       } else if (b <= a  && b <= c && b <= d && b >= 0 && tmp.x != togo.x) {
    //         //    fmt.Printf("B")
    //         tmp.x -= 1
    //         openList = append(openList, tmp)
    //         } else if (c <= b  && c <= a && c <= d && c >= 0 && tmp.y != togo.y) {
    //           //    fmt.Printf("C")
    //
    //           openList = append(openList, tmp)
    //           tmp.y += 1
    //           }  else if (d <= b  && d <= c && d <= a && d >= 0 && tmp.y != togo.y) {
    //             //    fmt.Printf("D")
    //             tmp.y -= 1
    //             openList = append(openList, tmp)
    //           }
    //
    //           deplacement++
    //         }
    //         var l = 0
    //         for l < len(openList)  {
    //           fmt.Printf("[%d %d]\n",openList[l].x, openList[l].y)
    //           l++
    //         }
    //         return openList, len(openList)
    //       }

    // func findToGo(tab [][]int,goodTab[][]int, len int ) []Noeu, int  {
    //
    //   n :=[]Noeu{}
    //   var x = 0
    //   var l = 0
    //
    //   for x < len {
    //     var y = 0
    //     for y < len {
    //       if tab[x][y] != goodTab[x][y]{
    //         var tmp Noeu
    //         tmp.x = x
    //         tmp.y = y
    //         tmp.value = tab[x][y]
    //         n = append(n, tmp)
    //         l++
    //       }
    //       y++
    //     }
    //     x++
    //   }
    //   return n, l
    // }

    //TODO Dude, suckin' at something is the first step to being sorta good at something



    // for !checKGood(tab,goodTab,len) && cal < 500  {
    //   var x = 0
    //   for x < len  {
    //     var y = 0
    //     for y < len {
    //       if tab[x][y] == 0 {
    //         var shouldBe  = goodTab[x][y]
    //         //TODO if shouldBe == 0 bug
    //         //TODO revoir lalgo
    //         if shouldBe == 0 {
    //           tab[x][y] = tab[x + 1][y]
    //           tab[x + 1][y] = 0
    //           } else {
    //              var at Noeu
    //              at.x = x
    //              at.y = y
    //              togo, cr := findToGo(tab, goodTab, len)
    //           //   fmt.Printf("\n at.x = %d, at.y = %d togo.x = %d, togo.y = %d  \n",at.x,at.y,togo.x,togo.y)
    //              openList, l:= astarApllication(tab, at, togo)
    //             // var k = 0
    //             // var xT = x
    //             // var yT = y
    //             // for k < l {
    //             //   tmp := tab[openList[k].x][openList[k].y]
    //             //   tab[xT][yT] = tmp
    //             //   tab[openList[k].x][openList[k].y] = 0
    //             //   xT = openList[k].x
    //             //   yT = openList[k].y
    //             //   k++
    //             //   aff(tab, len)
    //             // }
    //           }
    //         }
    //         y++
    //         aff(tab, len)
    //         cal++
    //       }
    //       x++
    //     }
    //   }
