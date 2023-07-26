package cpf

import (
	"backend/internal/pkg/shared"
	"fmt"
	"strconv"
	"strings"
)

type Cpf struct {
	Number string
}

func NewCpf(number string) Cpf {
	return Cpf{Number: number}
}

func (cpf Cpf) IsValid() bool {
	return validate(cpf.Number)
}

func validate(cpf string) bool {
	cpf = shared.ClearString(cpf, ".", "-", " ")
	fmt.Println("cpf: ", cpf)
	if shared.IsEmpty(cpf) {
		return false
	}

	if !isValidLength(cpf) {
		return false
	}

	if allDigitsAreEqual(cpf) {
		return false
	}

	dg1 := calculateDigit(cpf, 10)
	dg2 := calculateDigit(cpf, 11)
	checkDigits := extractCheckDigits(cpf)
	nDigResult := fmt.Sprintf("%d%d", dg1, dg2)

	return checkDigits == nDigResult
}

func isValidLength(cpf string) bool {
	return len(cpf) == 11
}

func allDigitsAreEqual(cpf string) bool {
	return strings.HasPrefix(strings.Repeat(string(cpf[0]), len(cpf)), cpf)
}

func extractCheckDigits(cpf string) string {
	return cpf[len(cpf)-2:]
}

func calculateDigit(cpf string, factor int) int {
	total := getTotal(cpf, factor)
	var rest = total % 11
	if rest < 2 {
		return 0
	} else {
		return 11 - rest
	}
}

func getTotal(cpf string, factor int) int {
	var total = 0
	for _, cpfByte := range cpf {
		if factor > 1 {
			digit, _ := strconv.Atoi(string(cpfByte))
			total += digit * factor
		}
		factor = factor - 1
	}
	return total
}
