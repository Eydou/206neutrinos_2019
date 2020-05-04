//
// EPITECH PROJECT, 2020
// 202unsold_2019
// File description:
// main
//

package main

import (
	"fmt"
	"os"
	"strconv"

	functions "./functions"
)

func help() {
	fmt.Printf("USAGE\n   ./206neutrinos n a h sd\n")
	fmt.Printf("\nDESCRIPTION\n")
	fmt.Printf("   n\t    number of values\n")
	fmt.Printf("   a\t    arithmetic mean\n")
	fmt.Printf("   h\t    harmonic mean\n")
	fmt.Printf("   sd\t    standard deviation\n")
	os.Exit(0)
}

func main() {
	args := os.Args

	if len(args) == 2 {
		if args[1] == "-h" || args[1] == "--help" {
			help()
		}
	}
	if _, err := functions.ErrorArgs(args); err != nil {
		fmt.Fprintf(os.Stderr, "\033[31mX\033[0m Error: %s\n", err)
		os.Exit(84)
	}
	numbers := make([]float64, 0)
	for i := 1; i != len(args); i++ {
		nb, err := strconv.Atoi(args[i])
		if err != nil {
			fmt.Println(err, nb)
			os.Exit(84)
		}
		if nb < 0 {
			os.Exit(84)
		}
		numbers = append(numbers, float64(nb))
	}
	os.Exit(functions.LoopNeutrinos(numbers))
}
