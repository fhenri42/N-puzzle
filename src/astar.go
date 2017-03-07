package main

import (
  "fmt"
)

type Noeu struct {
  x int
  y int
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

func astar(tab [][]int, goodTab [][]int, len int) int  {

  openList := make([]state, 1)
  openList[0].grid  = tab
  openList[0].g = 0
  openList[0].len = len
  openList[0].parent.x = -2
  openList[0].parent.y = -2
  openList[0].index = 0
  openList[0].pos = shouldBe(tab, len)
  openList[0].h = calculeManhattan(openList[0], goodTab)
  var nbrState = 1
  var nbrMove = 0
  var indexState = 1
  var succes = false
  closeList := make([]state, 0)

  for succes != true {


    //  fmt.Printf("openList == %d\n",openList);
    for index, _ := range openList {

      best := findBest(openList)
      //          fmt.Printf("best ==%d\n", best)
      if (index > indexState ) { indexState = index}
      aff(best.grid, len)
      if checKGood(best.grid, goodTab, len) == true {
        succes = true
        fmt.Printf("SUCCESSE\n\nWe did %d differente state, and %d move, to find the solution.\nThe maximum state in the memory at the same time was: %d\n\n",nbrState, nbrMove, indexState)
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
            nbrMove ++
            //          fmt.Printf("move ==%d\n",move[l])
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
              nbrState++

              } else {
                find := inListToFind(closeList, move[l])
                fmt.Printf("maison=  find.g = %d find.h = %d && find.h = %d ,best.g = %d + 1:",find.g, find.h, find.h, best.g )

                if find.g + find.h > find.h + best.g + 1 {
                  find.g = find.h + best.g + 1
                  fmt.Printf("maison1 \n")
                  err3 :=  inList(closeList, find.grid)
                  if err3 != 0 {
                    fmt.Printf("maison2 \n")
                    closeList = removeList(closeList, find.index)
                    openList = append(openList, find)
                  }
                }
              }
              l++
            }
          }
        }
      }
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
