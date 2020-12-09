package solution

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
	"unicode"
)

type Passport struct {
	byr string // (Birth Year)
	iyr string // (Issue Year)
	eyr string // (Expiration Year)
	hgt string // (Height)
	hcl string // (Hair Color)
	ecl string // (Eye Color)
	pid string // (Passport ID)
	cid string // (Country ID)
}

func (p *Passport) HasRequiredFields() bool {
	return p.byr != "" && p.iyr != "" && p.eyr != "" && p.hgt != "" && p.hcl != "" && p.ecl != "" && p.pid != ""
}

func (p *Passport) ValidateMinMax(value string, min int, max int) bool {

	converted, err := strconv.ParseInt(value, 10, 64)

	if err != nil {
		return false
	}

	return int(converted) >= min && int(converted) <= max
}

func (p *Passport) ValidateHeight(height string) bool {
	cm := strings.Contains(height, "cm")
	in := strings.Contains(height, "in")

	if cm {
		value := strings.Replace(height, "cm", "", 1)
		return p.ValidateMinMax(value, 150, 193)
	}

	if in {
		value := strings.Replace(height, "in", "", 1)
		return p.ValidateMinMax(value, 59, 76)
	}

	return false
}

func (p *Passport) ValidateHairColor(color string) bool {
	if !strings.HasPrefix(color, "#") {
		return false
	}

	value := strings.Replace(color, "#", "", 1)

	for _, char := range value {
		if unicode.IsSymbol(char) {
			return false
		}
	}

	return len(value) == 6
}

func (p *Passport) ValidateEyeColor(eyeColor string) bool {
	required := []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}

	for _, color := range required {
		if eyeColor == color {
			return true
		}
	}
	return false
}

func (p *Passport) IsValidPart1() bool {
	return p.byr != "" && p.iyr != "" && p.eyr != "" && p.hgt != "" && p.hcl != "" && p.ecl != "" && p.pid != ""
}

func (p *Passport) IsValidPart2() bool {
	requirements := p.HasRequiredFields()

	checkLength := len(p.byr) == 4 && len(p.iyr) == 4 && len(p.eyr) == 4 && len(p.pid) == 9

	byr := p.ValidateMinMax(p.byr, 1920, 2002)
	iyr := p.ValidateMinMax(p.iyr, 2010, 2020)
	eyr := p.ValidateMinMax(p.eyr, 2020, 2030)
	hgt := p.ValidateHeight(p.hgt)
	hcl := p.ValidateHairColor(p.hcl)
	ecl := p.ValidateEyeColor(p.ecl)

	return requirements && checkLength && byr && iyr && eyr && hgt && hcl && ecl
}

type Day4 struct {
	// We use a map to ensure uniqueness
	Passports map[Passport]struct{}
}

func (d *Day4) Input(data io.Reader) error {
	d.Passports = make(map[Passport]struct{})

	scanner := bufio.NewScanner(data)

	var fields []string
	for scanner.Scan() {
		line := scanner.Text()

		if len(line) == 0 {
			fields = nil
			continue
		}

		temp := strings.Split(line, " ")
		fields = append(fields, temp...)

		passport := d.CreatePassport(fields)
		d.Passports[passport] = Empty{}

	}

	return nil
}

func (d *Day4) CreatePassport(fields []string) Passport {

	passport := Passport{}

	for _, field := range fields {
		p := strings.Split(field, ":")

		switch p[0] {
		case "byr":
			passport.byr = p[1]
		case "iyr":
			passport.iyr = p[1]
		case "eyr":
			passport.eyr = p[1]
		case "hgt":
			passport.hgt = p[1]
		case "hcl":
			passport.hcl = p[1]
		case "ecl":
			passport.ecl = p[1]
		case "pid":
			passport.pid = p[1]
		default:
			continue
		}
	}

	return passport
}

func (d *Day4) Output() {

	var validatedPasswordsPart1 int
	var validatedPasswordsPart2 int

	for passport := range d.Passports {
		if passport.IsValidPart1() {
			validatedPasswordsPart1++
		}
	}

	for passport := range d.Passports {
		if passport.IsValidPart2() {
			validatedPasswordsPart2++
		}
	}

	fmt.Printf("=> [PART1] Answer: %d passports are valid\n", validatedPasswordsPart1)
	fmt.Printf("=> [PART2] Answer: %d passports are valid\n", validatedPasswordsPart2)
}
