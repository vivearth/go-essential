package main

func findMax(input []int) int {
	var biggestSofar = input[0]
	for i := 1; i < len(input); i++ {
		if input[i] > biggestSofar {
			biggestSofar = input[i]
		}
	}
	return biggestSofar
}
