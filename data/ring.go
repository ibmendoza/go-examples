//https://graysonkoonce.com/solving-ring-shaped-problems-with-golangs-container-ring/

package main

import (  
    "container/ring"
    "fmt"
)

func main() {  
    games := []int{100, 5, 6, 7, 10, 500}

    for _, playerAmount := range games {
        survivor := findSurvivor(playerAmount)
        fmt.Println("Given", playerAmount, "players, the surviving player was", survivor)
    }
}

func findSurvivor(playerAmount int) int {  
    players := ring.New(playerAmount)
    for i := 1; i <= players.Len(); i++ {
        players = players.Next()
        players.Value = i
    }

    skip := 1
    for players.Len() > 1 {
        players.Unlink(1)
        players = players.Move(skip)
        skip++
    }
    return players.Value.(int)
}
