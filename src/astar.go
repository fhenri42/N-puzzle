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
func myCopy(best state) state  {
  var tmp  state
  tmp.grid = best.grid
  tmp.g = best.g
  tmp.len = best.len
  tmp.parent.x = best.parent.x
  tmp.parent.y = best.parent.y
  tmp.index = best.index
  tmp.pos = best.pos
  tmp.h = best.h
  return tmp
}
func addToStart(closeList []state, best state) []state  {
  var y = 1
  var t = 0
  newClose := make([]state, len(closeList) + 1)
  newClose[0] = myCopy(best)
  for y < len(closeList) {
    newClose[y] = myCopy(closeList[t])
    y++
    t++
  }
 return newClose
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
  openList[0].h = calculeManhattan(openList[0], goodTab)  * 10
  var nbrState = 1
  var nbrMove = 0
  var indexState = 1
  var succes = false
  closeList := make([]state, 0)

  for succes != true {

    for index, _ := range openList {

      best := findBest(openList)
      if (index > indexState ) { indexState = index}
        aff(best.grid, len)
      if checKGood(best.grid, goodTab, len) == true {
        succes = true
        fmt.Printf("SUCCESSE\n\nWe did %d differente state, and %d move, to find the solution.\nThe maximum state in the memory at the same time was: %d\n\n",nbrState, nbrMove, indexState)
        return 1
        } else {

          openList = removeList(openList, best.index)
          move, cr := MoveGrid(best)
          closeList = append(closeList, best)

          var l = 0

          for l < cr {
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
              newList.h = calculeManhattan(newList, goodTab) * 10
              openList = append(openList, newList)
              nbrState++

              } else {
                nbrMove ++
                found := inListToFind(closeList, move[l])
                if found.len == 0 { found = inListToFind(openList, move[l]) }
                if found.g  > best.g + 1 {
                  found.g = best.g + 1
                  found.parent.x = best.pos.x
                  found.parent.y = best.pos.y
                  fmt.Printf("found.index == %d\n",found.index)
                  if err1 == 0 {
                     fmt.Printf("in the if after the elseaoeuaoeu\n")
                    closeList = removeList(closeList, found.index)
                    openList = append(openList, found)
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
