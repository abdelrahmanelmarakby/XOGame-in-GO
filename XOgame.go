package main

import (
	"fmt"
)

func main() {

	for {

		//creating the board in which you will play
		XOboard := [3][3]string{}
		Player := "X"
		fmt.Println("You are playing with :", Player)

		//creating Rows to be used for mapping where to play
		var row int
		fmt.Println("Enter row Number 0,1 or 2 :")
		fmt.Scanln(&row)
		if row > 2 || row < 0 {
			fmt.Println("are you playing in another board??")
			continue
		}
		//creating Columns to be used for mapping where to play
		var column int
		fmt.Println("Enter Column Number :")
		fmt.Scanln(&column)
		if column > 2 || column < 0 {
			fmt.Println("are you playing in another board??")
			continue
		}

		fmt.Println("Column Number is:", column, "   ", "Row Number is:", row)

		//mapping where to play and make sure that you don't override the othe Player's value
		if XOboard[row][column] == "" {
			XOboard[row][column] = Player
		} else {
			fmt.Println("not an empty space")
		}

		for i := 0; i < 3; i++ {
			// check rows || columns
			if (XOboard[i][0] == XOboard[i][1] && XOboard[i][1] == XOboard[i][2] && XOboard[i][2] == Player) ||
				(XOboard[0][i] == XOboard[1][i] && XOboard[1][i] == XOboard[2][i] && XOboard[2][i] == Player) {
				fmt.Println("Game ended, winner is Player:", Player)
				// we use return to exit both for loops and the app
				return

			}
		}
		if (XOboard[0][0] == XOboard[1][1] && XOboard[1][1] == XOboard[2][2] && XOboard[2][2] == Player) || (XOboard[0][2] == XOboard[1][1] && XOboard[1][1] == XOboard[2][0] && XOboard[2][0] == Player) {
			fmt.Println("Game ended, winner is Player:", Player)
			// end the game and exit the app
			// we can use "break" or "return"
			return
		}
		fmt.Println(XOboard)
		//Switching Players
		if Player == "X" {
			Player = "O"
		} else {
			Player = "X"
		}
	}
}
