package main

import "fmt"

func main() {
	// F STRINGS
	player_name := "Kobe Bryant"
	player_nba_years := 20
	player_pts_avg := 25.0
	fmt.Printf("\nMy name is %v and I played %v years in the NBA.", player_name, player_nba_years)
	fmt.Printf("\nName is of type %T and years in the NBA is of type %T", player_name, player_nba_years)
	fmt.Printf("\n%v averaged %0.1f points per game for his career.", player_name, player_pts_avg)

	// SPRINTF
	var saved_string string = fmt.Sprintf("\n%v averaged %0.1f points per game for his career.", player_name, player_pts_avg)
	fmt.Print("\nThe saved string is:", saved_string)

}
