package solution

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type Empty struct{}

type Password struct {
	Min   int
	Max   int
	Char  string
	Value string
}

func (t *Password) GetMinMax(expr string) error {
	values := strings.Split(expr, "-")

	min, minErr := strconv.ParseInt(values[0], 10, 32)

	max, maxErr := strconv.ParseInt(values[1], 10, 32)

	if minErr != nil || maxErr != nil {
		return errors.New("Conversion error")
	}

	t.Min = int(min)
	t.Max = int(max)

	return nil
}

func (t *Password) GetChar(char string) error {
	if char == "" || len(char) > 2 {
		return errors.New("INVALID FORMAT")
	}

	sanitized := strings.Replace(char, ":", "", -1)

	if len(sanitized) > 1 {
		return errors.New("INVALID FORMAT")
	}

	t.Char = sanitized

	return nil
}

type Day2 struct {
	Passwords []Password
}

func (d *Day2) Input(data io.Reader) error {
	scanner := bufio.NewScanner(data)

	for scanner.Scan() {
		line := scanner.Text()

		tokens := strings.Split(line, " ")

		token := Password{Value: tokens[2]}

		minMaxErr := token.GetMinMax(tokens[0])
		charErr := token.GetChar(tokens[1])

		if minMaxErr != nil || charErr != nil {
			continue
		}

		d.Passwords = append(d.Passwords, token)
	}

	return nil
}

func (d *Day2) ValidatePasswords() int {

	// Poor mans SET
	validatedPasswords := make(map[string]struct{})

	for _, passwd := range d.Passwords {
		count := strings.Count(passwd.Value, passwd.Char)

		if count < passwd.Min || count > passwd.Max {
			continue
		}

		validatedPasswords[passwd.Value] = Empty{}
	}

	return len(validatedPasswords)
}

func (d *Day2) Output() {
	passwordsValidated := d.ValidatePasswords()

	fmt.Printf("ANSWER: %d passwords are valid\n", passwordsValidated)
}
