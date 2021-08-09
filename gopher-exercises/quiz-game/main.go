package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
)
var Lines struct {
	Question string
	Answer string
}
func main() {
	filePath := flag.String("file", "./problems.csv", "input file path")
	duration := flag.Int64("duration", 30, "time limit for the quizz")
	flag.Parse()
	fmt.Println("this is the duration",*duration)

	lines := ReadCsvFile(*filePath)
	fmt.Println("Let's Start The Quizz")

	StartGame(lines)
	

}
func StartGame(lines [][]string) {
	var right int64
	var wrong int64
	for _, value := range lines {
		if value[0] != "question" && value[1] != "answer" {
			fmt.Println("What is the answer to ", value[0], "?")
			scanner := bufio.NewScanner(os.Stdin)
			
			scanner.Scan()
			answer := scanner.Text()
			
			if answer == value[1] {
				right++
			} else {
				wrong++
			}
		}
	}
	fmt.Println("correct answers: ", right)
	fmt.Println("incorrect answers: ", wrong)

}
func ReadCsvFile(path string) [][]string {
	file, err := os.Open(path)
	
	if err != nil {
		log.Fatal("Error occured while opening file:  ", err)
		os.Exit(1)
	}
	
	defer file.Close()
	csvReader := csv.NewReader(file)
	
	records, err := csvReader.ReadAll()
	
	if err != nil {
		log.Fatal("Unable to parsea CSV for: ", path, err)
	}

	return records



}