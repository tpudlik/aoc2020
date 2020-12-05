package passports

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type Passport struct {
	fields map[string]string
}

func (p Passport) RequiredFieldsPresent() bool {
	// All fields are required, except for cid.
	required := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
	for _, field := range required {
		if _, ok := p.fields[field]; !ok {
			return false
		}
	}
	return true
}

func (p Passport) Valid() bool {
	return (p.RequiredFieldsPresent() &&
		p.validYearField("byr", 1920, 2002) &&
		p.validYearField("iyr", 2010, 2020) &&
		p.validYearField("eyr", 2020, 2030) &&
		p.validHeight() &&
		p.validHairColor() &&
		p.validEyeColor() &&
		p.validPassportId())
}

func (p Passport) validYearField(field string, at_least, at_most int) bool {
	if len(p.fields[field]) != 4 {
		return false
	}
	value, err := strconv.Atoi(p.fields[field])
	if err != nil {
		return false
	}
	if value < at_least || value > at_most {
		return false
	}
	return true
}

func (p Passport) validHeight() bool {
	height := p.fields["hgt"]
	if len(height) < 4 {
		return false
	}
	value, err := strconv.Atoi(height[0 : len(height)-2])
	if err != nil {
		return false
	}
	switch height[len(height)-2 : len(height)] {
	case "cm":
		return value >= 150 && value <= 193
	case "in":
		return value >= 59 && value <= 76
	default:
		return false
	}
}

func (p Passport) validHairColor() bool {
	if len(p.fields["hcl"]) != 7 {
		return false
	}
	if p.fields["hcl"][0] != '#' {
		return false
	}
	for _, c := range p.fields["hcl"][1:] {
		if !strings.Contains("0123456789abcdef", string(c)) {
			return false
		}
	}
	return true
}

func (p Passport) validEyeColor() bool {
	validColors := []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}
	for _, c := range validColors {
		if p.fields["ecl"] == c {
			return true
		}
	}
	return false
}

func (p Passport) validPassportId() bool {
	if len(p.fields["pid"]) != 9 {
		return false
	}
	for _, c := range p.fields["pid"] {
		if !strings.Contains("0123456789", string(c)) {
			return false
		}
	}
	return true
}

func ParseBatch(r io.Reader) ([]Passport, error) {
	out := []Passport{}
	scanner := bufio.NewScanner(r)
	var p Passport
	in_passport := false
	for scanner.Scan() {
		txt := scanner.Text()
		if txt == "" {
			// Either blank line between two passports, or blank lines at the
			// beginning or end of file.
			if in_passport {
				// End of current passport.
				out = append(out, p)
				in_passport = false
			}
			continue
		}
		if !in_passport {
			// Start of a new passport
			p = Passport{fields: map[string]string{}}
			in_passport = true
		}

		tokens := strings.Split(txt, " ")
		for _, token := range tokens {
			components := strings.Split(token, ":")
			if len(components) != 2 {
				return nil, fmt.Errorf("Cannot parse %v as passport token", token)
			}
			p.fields[components[0]] = components[1]
		}
	}

	// If the input doesn't end in a blank line, we may have one last passport
	// to export.
	if in_passport {
		out = append(out, p)
	}
	return out, nil
}
