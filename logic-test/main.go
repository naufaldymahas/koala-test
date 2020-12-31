package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	fmt.Println(fibonacci0To1000())
	fmt.Println(primeNumber0To1000())
	fmt.Println(isPalindrome("Dennis, Nell, Edna, Leon, Nedra, Anita, Rolf, Nora, Alice, Carol, Leo, Jane, Reed, Dena, Dale, Basil, Rae, Penny, Lana, Dave, Denny, Lena, Ida, Bernadette, Ben, Ray, Lila, Nina, Jo, Ira, Mara, Sara, Mario, Jan, Ina, Lily, Arne, Bette, Dan, Reba, Diane, Lynn, Ed, Eva, Dana, Lynne, Pearl, Isabel, Ada, Ned, Dee, Rena, Joel, Lora, Cecil, Aaron, Flora, Tina, Arden, Noel, and Ellen sinned"))
}

func fibonacciSequence(n int32) int32 {
	if n <= 1 {
		return n
	}
	return fibonacciSequence(n-1) + fibonacciSequence(n-2)
}

func fibonacci0To1000() (result []int32) {
	for i := int32(0); true; i++ {
		if fibonacciSequence(i) > 1000 {
			break
		}
		result = append(result, fibonacciSequence(i))
	}
	return
}

func isPrimeNumber(n int32) bool {
	if n <= 1 {
		return false
	}

	if n == 2 || n == 3 {
		return true
	}

	for i := int32(2); i < n; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}

func primeNumber0To1000() (result []int32) {
	for i := int32(0); i <= 1000; i++ {
		if isPrimeNumber(i) {
			result = append(result, i)
		}
	}
	return
}

func isPalindrome(str string) bool {
	regex := regexp.MustCompile(`[^a-zA-Z]`)
	normalizeString := strings.ToLower(regex.ReplaceAllString(str, ""))
	builder := strings.Builder{}
	for i := len(normalizeString) - 1; i >= 0; i-- {
		builder.WriteByte(normalizeString[i])
	}

	if normalizeString == builder.String() {
		return true
	}
	return false
}
