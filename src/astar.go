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

func astar(tab [][]int, goodTab [][]int, len int, choice int) int  {

  openList := make([]state, 1)
  openList[0].grid  = tab
  openList[0].g = 0
  openList[0].len = len;
  openList[0].parent.x = -2
  openList[0].parent.y = -2
  openList[0].index = 0
  openList[0].pos = shouldBe(tab, len)
  if choice == 1 { openList[0].h = calculeManhattan(openList[0], goodTab)  * 10
  } else if choice == 2 { openList[0].h = misplaced(openList[0], goodTab)  * 10 } else { openList[0].h = h_row_column(openList[0], goodTab)  * 10 }
  var nbrState = 1
  var nbrMove = 0
  var indexState = 1
  var succes = false
  closeList := make([]state, 0)

  for succes != true {

    for index := range openList {

      best := findBest(openList)

      if (index > indexState ) { indexState = index}
      if checKGood(best.grid, goodTab, len) == true {
        succes = true
        var nbrMoveRec = 1
        for _, close := range closeList {
          aff(close.grid, len, nbrMoveRec)
          nbrMoveRec++
        }
        aff(best.grid, len, nbrMoveRec)
        fmt.Printf("\033[1;32mSUCCESS !\033[m\n\n\033[32;35mWe did \033[m\033[32;25m%d\033[m \033[32;35mdifferente state (comlexity is size).\nNumber of move to find the solution: \033[m \033[32;25m%d\033[m \033[32;35m\nNumber of move from intiale state to the solution: \033[m \033[32;25m%d\033[m \033[32;35m\nThe maximum state in the memory (complexity in time) at the same time was: \033[m\033[32;25m%d\033[m\n\n",nbrState, nbrMove,nbrMoveRec,indexState)
        return 1
        } else {

          openList = removeList(openList, best.index)
          move, cr := MoveGrid(best)
          closeList = append(closeList, best)

          var l = 0

          for l < cr {
            nbrMove ++
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
              if choice == 1 { newList.h = calculeManhattan(newList, goodTab) * 10
              } else if choice == 2 { newList.h = misplaced(newList, goodTab) * 10
              } else { newList.h = h_row_column(newList, goodTab) * 10 }
              openList = append(openList, newList)
              nbrState++
              } else {

                found := inListToFind(closeList, move[l])
                if found.len == 0 {found = inListToFind(openList, move[l]) }
                if found.g  > best.g + 1 {
                  found.g = best.g + 1
                  found.parent.x = best.pos.x
                  found.parent.y = best.pos.y
                  if err1 == 0 {
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

    func  aff(tab [][]int, len int, state int)  {
      var x = 0
      fmt.Printf("reconstruct path: \033[32;25m%d\033[m\n\n", state)
      for x < len {
        fmt.Printf("\033[1;33m%d\033[m\n",tab[x])
        x++
      }
      fmt.Printf("\n")
    }
