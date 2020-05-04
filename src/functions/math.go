//
// EPITECH PROJECT, 2020
// 201yams_2019
// File description:
// error
//

package functions

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

//Calcul calcul
func Calcul(value int, numbers []float64, index int) {
	Amean := (float64(numbers[1]) * (float64(numbers[0]) - 1)) + float64(value)
	Hmean := float64(numbers[0]) / ((1 / float64(value)) + ((float64(numbers[0]) - 1) / float64(numbers[2])))
	Rmean := math.Sqrt(((math.Pow(numbers[3], 2)+math.Pow(numbers[1], 2))*(numbers[0]-1) + math.Pow(float64(value), 2)) / numbers[0])
	newStdDev := math.Sqrt((((math.Pow(float64(numbers[3]), 2)+math.Pow(float64(numbers[1]), 2))*(float64(numbers[0])-1) + math.Pow(float64(value), 2)) / (float64(numbers[0]))) - math.Pow(((float64(numbers[1])*(float64(numbers[0])-1))+(float64(value)))/(float64(numbers[0])), 2))
	fmt.Printf("\tNumber of values:\t%.0f\n", numbers[0])
	fmt.Printf("\tStandard deviation:\t%.2f\n", float64(newStdDev))
	fmt.Printf("\tArithmetic mean:\t%.2f\n", Amean/float64(numbers[0]))
	fmt.Printf("\tRoot mean square:\t%.2f\n", Rmean)
	fmt.Printf("\tHarmonic mean:\t%.2f\n\n", Hmean)
	numbers[3] = float64(newStdDev)
	numbers[1] = Amean / float64(numbers[0])
	numbers[2] = Hmean
}

//LoopNeutrinos loop
func LoopNeutrinos(numbers []float64) int {
	reader := bufio.NewReader(os.Stdin)
	for i := 0; i != -1; i++ {
		fmt.Printf("Enter next value: ")
		str, err := reader.ReadString('\n')
		if err != nil {
			os.Exit(84)
		}
		str = strings.Replace(str, "\n", "", -1)
		if strings.Compare("END", str) == 0 {
			break
		}
		value, err := strconv.Atoi(str)
		if err != nil {
			os.Exit(84)
		}
		numbers[0]++
		Calcul(value, numbers, i)
	}
	return 0
}
