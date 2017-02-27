package main

import (
  "fmt"
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

func  aff(tab [][]int, len int)  {
  var x = 0
  fmt.Printf("====================\n")
  for x < len {
    fmt.Printf("%d\n",tab[x])
    x++
  }
  fmt.Printf("====================\n")

}
// func noeuComparaison(n1 Heuristique, n2 Heuristique) int  {
//   if n1.heuristique < n2.heuristique  {
//     return 1
//   } else if n1.heuristique  == n2.heuristique {
//     return 0
//   } else {
//     return -1
//   }
//
// }
//
func calculeHeu(deplacement int, nextPos Noeu, togo Noeu) Heu  {
  var heu Heu
  somX =  nextPos.x - togo.x
  somY = nextPos.y - togo.y
  heu.H = somX - somY
  heu.G = deplacement
}

func astarApllication(tab [][]int, at Noeu, togo Noeu) Noeu  {
  var tmp Noeu
  var deplacement  = 1
  tmp.x = at.x
  tmp.y = at.y
  for tmp.x  != togo.x && tmp.y != togo.y {
     a := calculeHeu(deplacement,  tmp, togo)
     b:= calculeHeu(deplacement,  tmp, togo)
     c := calculeHeu(deplacement,  tmp, togo)
     d := calculeHeu(deplacement,  tmp, togo)

     deplacement++
  }
  return at
}

func findToGo(tab [][]int,nbr, len int ) Noeu  {
  var n Noeu
  n.x = 0
  n.y = 0
  var x = 0
  for x < len {
    var y = 0
    for y < len {
      if tab[x][y] == nbr {
        n.x = x
        n.y = y
      }
      y++
    }
    x++
  }
  return n
}

func astar(tab [][]int, goodTab [][]int, len int) int  {

  var cal = 0
  aff(tab, len)
  for !checKGood(tab,goodTab,len) && cal < 5  {
    var x = 0
    for x < len  {
      var y = 0
      for y < len {
        if tab[x][y] == 0 {
          var shouldBe  = goodTab[x][y]
          if y + 1 < len && tab[x][y + 1] == shouldBe {
            tab[x][y] = shouldBe
            tab[x][y + 1] = 0

            } else if y - 1 >=  0 && tab[x][y - 1] == shouldBe {

              tab[x][y] = shouldBe
              tab[x][y - 1] = 0

              } else if x + 1 < len && tab[x + 1][y] == shouldBe {

                tab[x][y] = shouldBe
                tab[x + 1][y] = 0

                } else if  x - 1 >= 0 && tab[x - 1][y] == shouldBe {

                  tab[x][y] = shouldBe
                  tab[x - 1][y] = 0
                } else {
                  var at Noeu
                  at.x = x
                  at.y = y
                  togo := findToGo(tab, shouldBe, len)
                  fmt.Printf("\n at.x = %d, at.y = %d togo.x = %d, togo.y = %d  \n",at.x,at.y,togo.x,togo.y)
          //        astarApllication(tab, at, togo)
                }
              }
              y++
              aff(tab, len)
              cal++
            }
            x++
          }
        }
        fmt.Printf("nb of calcule = %d\n", cal)
        return 0
      }
