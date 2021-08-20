package main

import "fmt"

func main() {
	var music_score, math_score, sport_score int = 90, 80, 87
	var sum = music_score + math_score + sport_score
	fmt.Printf("sum = %d, avg = %f\n", sum, float64(sum)/3)
}
