package main

import (
//  "fmt"
)

func findBest(sta []state) state  {

  best := sta[0]
  for _, s := range sta {
      if s.h + s.g < best.h + best.g { best = s }
  }
  return best
}

func inList(aList []state, bList [][]int) int {

  if len(aList) == 0  { return  0 }
  var toNo = 0
  for _, list := range aList {
    toNo = 0
    var x = 0
    for x < list.len {
      var y = 0
      for y < list.len {
        if list.grid[x][y] == bList[x][y] {
          toNo++
        }
        y++
      }
      x++
    }
    if toNo == list.len * list.len { return 1 }
  }
  return 0
}

func inListToFind(aList []state, bList [][]int) state {

  var toNo = 0
  var tmp state
  for _,list := range aList {
    var x = 0
    for x < len(list.grid) {
      var y = 0
      for y < len(list.grid[x]) {
        if list.grid[x][y] == bList[x][y] {
          toNo++
        }
        y++
      }
      x++
    }
    if toNo == list.len * list.len { return list }
  }
  return tmp
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

func removeList(openList []state, index int) []state  {

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
  return k
}
