package main

import (
	"fmt"
	"strconv"
)

func getBonus(theRange int, maxRatingAdd int, ratingOpponent int, ownRating int, result string) int {
	var bonus int
	perRating := (2 * theRange) / maxRatingAdd
	low := ownRating - theRange
	if ratingOpponent <= ownRating - theRange {
		bonus = 0
	} else if ratingOpponent >= ownRating + theRange {
		bonus = maxRatingAdd
	} else {
		bonus = (ratingOpponent - low) / perRating
	}
	if result == "w" {
		return bonus
	} else if result == "d" {
		if ratingOpponent >= ownRating {
			return bonus / 2
		} else {
			return 0 - ((((maxRatingAdd - bonus) / 2) * 75) / 100)
		}
	} else {
		return 0 - (((maxRatingAdd - bonus) * 75) / 100)
	}
}

func main() {
	var theRange, maxRatingAdd, ownRating, ratingOpponent, ratingAdd int
	var result string
	var input string = "0"
	for {
		fmt.Print("\n\nOwn Rating: ")
		fmt.Scan(&ownRating)
		//Hoeveel rating boven eigen rating krijg je max punten?
		theRange = 675
		//Hoeveel punten krijg je max erbij?
		maxRatingAdd = 40
		for input != "." {
			fmt.Print("Rating tegenstander of . als je klaar bent: ")
			fmt.Scan(&input)
			if input == "." {
				break
			} else {
				ratingOpponent,_ = strconv.Atoi(input)
				result = "."
				for result != "w" && result != "d" && result != "l" {
					fmt.Print("Resultaat (w (winst), d (gelijkspel), l (verlies)): ")
					fmt.Scan(&result)
				}
				add := getBonus(theRange, maxRatingAdd, ratingOpponent, ownRating, result)
				fmt.Println("Rating erbij:", add)
				ratingAdd += add
			}
		}
		fmt.Println("Totaal Rating erbij:", ratingAdd)
		newRating := ownRating + ratingAdd
		fmt.Println("Nieuwe Rating:", newRating)
	}
}
