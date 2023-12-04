package day1

import (
	"bufio"
	"fmt"
	"log"
	"strconv"
)

type Digit struct {
	digit byte
	index int
}

func CalculateCaliberationPartTwo(scanner *bufio.Scanner) int {
	ans := 0

	for scanner.Scan() {
		line := scanner.Text()
		forward := false
		backward := false
		digit := make([]Digit, 2)

		for i, j := 0, len(line)-1; i < len(line); i, j = i+1, j-1 {
			if !forward && isNumeric(line[i]) {
				digit[0].digit = line[i]
				digit[0].index = i
				forward = true
			}
			if !backward && isNumeric(line[j]) {
				digit[1].digit = line[j]
				digit[1].index = j
				backward = true
			}

			if forward && backward {
				break
			}
		}
		if !forward && !backward {
			digit[0].index = -1
			digit[1].index = -1
		}

		// digits
		// fmt.Println(digit)

		ans += findNumericChar(line, digit)
	}

	return ans
}

func findNumericChar(line string, digit []Digit) int {

	numericChar := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}

	var num1, num2 string = "-1", "-1"
	// forward pass
	if digit[0].index == 0 {
		// no need to process
		num1 = string(digit[0].digit)
	}

	limit := digit[0].index

	if digit[0].index < 0 {
		limit = len(line) - 1
	}

  outerloop:
	for i := 0; i < limit; i++ {
		for j := limit; j > i; j-- {
			substr := line[i:j]
			if val, exist := numericChar[substr]; exist {
        num1 = strconv.Itoa(val)
        break outerloop
			}
		}
	}
  
  // backward pass
  if digit[1].index == len(line) - 1 {
    // no need to process
    num2 = string(digit[1].digit)
  }
  
  limit = digit[1].index + 1
  if (digit[1].index < 0) {
    limit = 0;
  }
  outerloop1:
  for i := len(line); i > limit; i-- {
    for j := 3; j <= i - limit; j++ {
      substr := line[i  - j: i]
      if val, exist := numericChar[substr]; exist {
        num2 = strconv.Itoa(val)
        break outerloop1
      }
    }
  }

  if num1 == "-1" {
	num1 = string(digit[0].digit)
  }
  if num2 == "-1" {
	num2 = string(digit[1].digit)
  }

  val, err := strconv.Atoi(fmt.Sprintf("%s%s", num1, num2))
  if err != nil {
    log.Fatal("Unable to convert string to int")
  }

  return val
}
