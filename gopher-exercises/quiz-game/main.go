package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)
type Counter struct {
	R int
	W int
}
func main() {
	// TestTimer()

	
	
	
	filePath := flag.String("file", "./problems.csv", "input file path")
	duration := flag.Int("duration", 30, "time limit for the quizz")
	
	flag.Parse()
	fmt.Println("this is the filePath",*filePath)
	fmt.Println("this is the duration",*duration)
	// timer := time.NewTimer((time.Duration(*duration) * time.Second))
	// <- timer.C
	
	lines := ReadCsvFile(filePath)
	fmt.Println("Let's Start The Quizz")
    c := Counter{R: 0, W: 0}
	// StartGame(lines, &c, duration)


	fmt.Println(duration)
	fmt.Println(*duration)
	timer := time.NewTimer(time.Duration(*duration) * time.Second)
	for index, value := range lines {
		fmt.Printf("Problem #%d: %s = \n", index+1, value[0])
		answerCh := make(chan string)
		go func() {
			scanner := bufio.NewScanner(os.Stdin)
			scanner.Scan()
			answer := scanner.Text()
			answerCh <- answer
		}()
		select {
		case <-timer.C:
			fmt.Printf("correct answers: %d\n", c.R)
			fmt.Printf("incorrect answers: %d\n", c.W)
			return
		case answer := <- answerCh:

				if strings.TrimSpace(answer) == strings.TrimSpace(value[1]) {
					c.R++
					fmt.Println("Correct!")
				} else {
					c.W++
					fmt.Println("Incorrect!")
				}
		}
	}
	

}
// func StartGame(lines [][]string, c *Counter, d *int) {
	// fmt.Println("ddd", *d)
	// timer := time.NewTimer((time.Duration(*d) * time.Second))
	// // <-timer.C
	// for index, value := range lines {
	// 	select {
	// 	case <-timer.C:
	// 		fmt.Printf("correct answers: %d\n", c.R)
	// 		fmt.Printf("incorrect answers: %d\n", c.W)
	// 		return
	// 	default:
	// 		if value[0] != "question" && value[1] != "answer" {
	// 			fmt.Printf("Problem #%d: %s = \n", index+1, value[0])
	// 			scanner := bufio.NewScanner(os.Stdin)
				
	// 			scanner.Scan()
	// 			answer := scanner.Text()
				
	// 			if strings.TrimSpace(answer) == strings.TrimSpace(value[1]) {
	// 				c.R++
	// 				fmt.Println("Correct!")
	// 			} else {
	// 				c.W++
	// 				fmt.Println("Incorrect!")
	// 			}
	// 		}
	// 	}
	// }
// }
func ReadCsvFile(path *string) [][]string {
	file, err := os.Open(*path)
	
	if err != nil {
		log.Fatal("Error occured while opening file:  ", err)
		
	}
	
	defer file.Close()
	csvReader := csv.NewReader(file)
	
	records, err := csvReader.ReadAll()
	
	if err != nil {
		log.Fatal("Unable to parsea CSV for: ", path, err)
	}

	return records



}


func TestTimer(){

	timer := time.NewTimer((5 * time.Second))

	fmt.Println(<- timer.C)


}
