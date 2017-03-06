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
  parent Noeu
  pos Noeu
}

func findBest(sta []state) state  {

  best := sta[0]
  for _, s := range sta {
      if s.h + s.g < best.h + best.g { best = s }
  }
  return best
}

      func inList(aList []state, bList [][]int) int {

        if len(aList) == 0  {
          return  0
        }
        for _,list := range aList {
          var x = 0
          for x < len(list.grid) {
            var y = 0
            for y < len(list.grid[x]) {
              if list.grid[x][y] != bList[x][y] {
                return 0
              }
              y++
            }
            x++
          }
        }
        return 1
      }

      func inListToFind(aList []state, bList [][]int) state {

        for _,list := range aList {
          var x = 0
          for x < len(list.grid) {
            var y = 0
            for y < len(list.grid[x]) {
              if list.grid[x][y] != bList[x][y] {
                return list
              }
              y++
            }
            x++
          }
        }
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
        fmt.Printf("checKGood")
        return true
      }
      func shouldBe(tab [][]int, len int) Noeu  {
        var x = 0
        var tmp Noeu
        for x < len {
          var y = 0
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
        fmt.Printf("baddddddddd")
        tmp.y = 0
        tmp.x = 0
        return tmp
      }

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

      func removeList(openList []state, index int)[]state  {
        var x = 0
        newList := make([]state, 0)
        for x < len(openList) {
          if x != index {
            newList = append(newList, openList[x])
          }
          x++
        }
        x = 0
        for x < len(newList) {
    //       fmt.Printf("list = %d\n",list.index)
    //       fmt.Printf("index = %d\n", index)
           newList[x].index = x
           x++
        }
        return newList
      }

      func  indexShouldBe(openList[]state) int {
        var k = 0
        for _,_ = range openList {
        k++
        }
      fmt.Printf("indexShouldBe = %d \n", k)
        return k
      }
      func astar(tab [][]int, goodTab [][]int, len int) int  {

        var cal = 0
        openList := make([]state, 1)
        openList[0].grid  = tab
        openList[0].g = 0
        openList[0].len = len
        openList[0].parent.x = -2
        openList[0].parent.y = -2
        openList[0].index = 0
        openList[0].pos = shouldBe(tab, len)
        openList[0].h = calculeManhattan(openList[0], goodTab)
        var p = 0
        var succes = false

        closeList := make([]state, 0)

        //  for succes != true {
        for p < 100 {
          fmt.Printf("p== %d\n",p);
          p++;
        //  fmt.Printf("openList == %d\n",openList);
          for _, open := range openList {

            best := findBest(openList)
  //          fmt.Printf("best ==%d\n", best)

            aff(best.grid, len)

            if checKGood(best.grid, goodTab, len) == true {
              succes = true
              fmt.Printf("SUCCESSE")
              return 1
              } else {
            //    fmt.Printf("best == %d\n", best)
            //    fmt.Printf("openList == %d\n", openList)
                openList = removeList(openList, best.index)
            //    fmt.Printf("openList == %d\n, best.index == %d\n", openList,best.index)
                move, cr := MoveGrid(best)
                closeList = append(closeList, best)
              //  fmt.Printf("closeList == %d\n", closeList)
                var l = 0

                for l < cr {

                  fmt.Printf("move ==%d\n",move[l])
                  err := inList(openList, move[l])
                  err1 := inList(closeList, move[l])

                  if err == 0 && err1 == 0  {

                    var newList state
                    newList.grid = move[l]
                    newList.parent.x = best.pos.x
                    newList.parent.y = best.pos.y
                    newList.g = best.g + 1
                    newList.len = best.len
                    newList.index = indexShouldBe(openList)
                    newList.pos = shouldBe(move[l],best.len)
                    newList.h = calculeManhattan(newList, goodTab)
                    openList = append(openList, newList)

                    } else {
                      find := inListToFind(closeList, move[l])
                      fmt.Printf("maison=  find.g = %d find.h = %d && find.h = %d ,best.g = %d + 1:",find.g, find.h, find.h, best.g )

                      if find.g + find.h > find.h + best.g + 1 {
                        find.g = find.h + best.g + 1
                        fmt.Printf("maison1 \n")
                        err3 :=  inList(closeList, find.grid)
                        if err3 != 0 {
                          fmt.Printf("maison2 \n")
                          fmt.Printf("maison2 \n")
                          fmt.Printf("maison2 \n")
                          fmt.Printf("maison2 \n")
                          fmt.Printf("maison2 \n")
                          fmt.Printf("maison2 \n")
                          fmt.Printf("maison2 \n")
                          fmt.Printf("maison2 \n")
                          fmt.Printf("maison2 \n")
                          fmt.Printf("maison2 \n")
                          fmt.Printf("maison2 \n")
                          fmt.Printf("maison2 \n")
                          fmt.Printf("maison2 \n")
                          fmt.Printf("maison2 \n")
                          closeList = removeList(closeList, find.index)
                          openList = append(openList, find)
                        }
                      }
                    }
                    l++
                  }
                }
                fmt.Printf("h = %d\n", open.h)
                //      }
              }
            }
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
