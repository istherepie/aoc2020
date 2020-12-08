package solution

import (
	"bufio"
	"fmt"
	"io"
	"strings"
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

func (p *Passport) isValid() bool {
	return p.byr != "" && p.iyr != "" && p.eyr != "" && p.hgt != "" && p.hcl != "" && p.ecl != "" && p.pid != ""
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

		if passport.isValid() {
			d.Passports[passport] = Empty{}
		}
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
	fmt.Printf("=> Answer: %d passports are valid\n", len(d.Passports))
}
