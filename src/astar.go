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

func  isPossible(close state ,son state) bool  {
  zero := shouldBe(son.grid, son.len)
  if zero.x + 1 < close.len && close.grid[zero.x + 1][zero.y] == 0 { return true }
  if zero.x - 1 >= 0 && close.grid[zero.x - 1][zero.y] == 0 { return true }
  if zero.y + 1 < close.len && close.grid[zero.x][zero.y + 1] == 0 { return true }
  if zero.y - 1 >= 0 && close.grid[zero.x][zero.y - 1] == 0 { return true }
  return false
}
func reconstructGrid(closeList []state, len int, best state) int {

  var nbrMoveRec = 1
  var son state
  for index, close := range closeList {
    if index == 0 {
      son = close
      aff(close.grid, len, nbrMoveRec)
      nbrMoveRec++
    }
    if isPossible(close, son) == true {
      aff(close.grid, len, nbrMoveRec)
      son = close
      nbrMoveRec++
    }
  }
  aff(best.grid, len, nbrMoveRec)
  return nbrMoveRec
}
func  endfunc(closeList []state, best state, len int, nbrState int, nbrMove int, indexState int)  {
  nbrMoveRec := reconstructGrid(closeList, len, best)
  fmt.Printf("\033[1;32mSUCCESS !\033[m\n\n\033[32;35mWe did \033[m\033[32;25m%d\033[m \033[32;35mdifferente state (comlexity is size).\nNumber of move to find the solution: \033[m \033[32;25m%d\033[m \033[32;35m\nNumber of move from intiale state to the solution: \033[m \033[32;25m%d\033[m \033[32;35m\nThe maximum state in the memory (complexity in time) at the same time was: \033[m\033[32;25m%d\033[m\n\n",nbrState, nbrMove,nbrMoveRec,indexState)
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
  var index = 1
  closeList := make([]state, 0)

  for 1 != 2 {

    for i := range openList {

      best := findBest(openList)
  //aff(best.grid, len, 0)
      if (i > indexState ) { indexState = i }
      if checKGood(best.grid, goodTab, len) == true {
fmt.Printf("oeuoeu")
        endfunc(closeList, best, len, nbrState, nbrMove, indexState);   return 1
      } else {

          openList = removeList(openList, best.index)
          closeList = append(closeList, best)
          move, cr := MoveGrid(best)
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
              newList.pos = shouldBe(move[l], best.len)
              if choice == 1 { newList.h = calculeManhattan(newList, goodTab) * 10
              } else if choice == 2 { newList.h = misplaced(newList, goodTab) * 10
              } else { newList.h = h_row_column(newList, goodTab) * 10 }
              openList = append(openList, newList)
              nbrState++
              index++
              } else {
            found := inListToFind(closeList, move[l])
              if found.len == 0 {
                found = inListToFind(openList, move[l])
            }
  //           //fmt.Printf("move = %d\n ", move[l])
  //           //fmt.Printf("found = %d\n ", found)
  //           //fmt.Printf("NEXT \n")
  //           //return 1
  //
              if found.g  > best.g + 1 {
                found.g = best.g + 1
                found.parent.x = best.pos.x
                found.parent.y = best.pos.y
                if err1 == 0 {

                  openList = append(openList, found)
                  closeList = removeList(closeList, found.index)
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
