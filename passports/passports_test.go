package passports

import (
	"strings"
	"testing"
)

func testPassports() []Passport {
	r := strings.NewReader(`ecl:gry pid:860033327 eyr:2020 hcl:#fffffd
byr:1937 iyr:2017 cid:147 hgt:183cm

iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884
hcl:#cfa07d byr:1929

hcl:#ae17e1 iyr:2013
eyr:2024
ecl:brn pid:760753108 byr:1931
hgt:179cm

hcl:#cfa07d eyr:2025 pid:166559648
iyr:2011 ecl:brn hgt:59in
`)
	passports, err := ParseBatch(r)
	if err != nil {
		panic(err)
	}
	return passports
}

func TestParsing(t *testing.T) {
	passports := testPassports()
	if got := len(passports); got != 4 {
		t.Fatalf("Expected 4 passports, got %d: %v", got, passports)
	}

	want := Passport{
		fields: map[string]string{
			"ecl": "gry", "pid": "860033327", "eyr": "2020", "hcl": "#fffffd",
			"byr": "1937", "iyr": "2017", "cid": "147", "hgt": "183cm",
		}}

	for k, v := range want.fields {
		got, ok := passports[0].fields[k]
		if !ok {
			t.Errorf("Passport %v is missing expected field %q", passports[0], k)
			continue
		}
		if got != v {
			t.Errorf("Passport %v has unexpected value %q (instead of %q) for field %q", passports[0], got, v, k)
		}
	}
}

func TestRequiredFieldsPresent(t *testing.T) {
	passports := testPassports()
	tests := []bool{true, false, true, false}
	for idx, want := range tests {
		if got := passports[idx].RequiredFieldsPresent(); got != want {
			t.Errorf("Passport %d: got %v, want %v", idx, got, want)
		}
	}
}
