import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

//function är en struct som representerar funktionen samt dess derivata och primitiva funktion som en string, exempel: function.fun = e^(2x)*x+x^3
type function struct {
	fun        string
	primitive  string
	derivative string
}

//hjälpfunktion som ser om ett fel har uppkommit och printar ut felet om det existerar
func checkError(e error) {
	if e != nil {
		fmt.Println(e)
	}
}

//hjälpfunktion som tar in en sträng och returnerar en int, om strängen är tom returneras 1
func parseIntHelp(s string) int64 {
	if s == "" {
		temp := int64(1)
		return temp
	} else {
		temp, err := strconv.ParseInt(s, 10, 64)
		checkError(err)
		return temp
	}

}

//tar in en sträng av exponentialfunktioner och returnerar en lista av function-structar som motsvarar dessa
func parse(input string) []function {
	if input == "" {
		panic("Fel: Strängen kan inte vara tom!")
	}
	re1, e1 := regexp.Compile("^\\d*e\\^\\d*x([\\+\\-\\/\\*]\\d*e(\\^\\d*x))*$")
	functionlist := []function{}
	checkError(e1)
	if re1.Match([]byte(input)) {
		re2, e2 := regexp.Compile("[\\+\\-\\/\\*]")
		partition := re2.Split(input, -1)
		checkError(e2)
		for i, v := range partition {
			for _, w := range partition[i] {
				if string(w) == "e" {
					functionlist = append(functionlist, function{v, "", ""})
				}
			}
		}
	} else {
		panic("Fel: otillåtna tecken i funktionen (endast e, x, ^+-*/, 0-9 är tillåtna)!")
	}
	return functionlist
}

//Deriverar en funktion, tar in en funktion som en sträng och returnerar dess derivata som en sträng
func Diff(input string) string {
	if input == "" {
		panic("Fel: Strängen kan inte vara tom!")
	}
	output := ""
	signslice := regexp.MustCompile("(\\d*e\\^\\d*x)*").Split(input, -1)
	partition := parse(input)
	for i, v := range partition {
		partition[i] = expDiff(v)
	}
	for i, v := range partition {
		signslice[i] = v.derivative + signslice[i+1]
	}
	output = strings.Join(signslice, "")
	return output
}

//Integrerar en funktion, tar in en funktion som en sträng och returnerar dess primitiva funktion som en sträng
func Prim(input string) string {
	if input == "" {
		panic("Fel: Strängen kan inte vara tom!")
	}
	output := ""
	signslice := regexp.MustCompile("(\\d*e\\^\\d*x)*").Split(input, -1)
	partition := parse(input)
	for i, v := range partition {
		partition[i] = expPrim(v)
	}
	for i, v := range partition {
		signslice[i] = v.primitive + signslice[i+1]
	}
	output = strings.Join(signslice, "")
	return output
}

//Hjälpfunktion som deriverar exponentialfunktioner, tar in en delfunktion som en function och returnerar dess derivata som en function
func expDiff(f function) function {
	coeff1 := regexp.MustCompile("e").Split(f.fun, 2)
	expcoeff := regexp.MustCompile("\\^").Split(string(coeff1[1]), 2)
	coeff2 := regexp.MustCompile("x").Split(string(expcoeff[1]), 2)
	coeffint1 := parseIntHelp(coeff1[0])
	coeffint2 := parseIntHelp(coeff2[0])
	coeffint1 *= coeffint2
	if coeffint2 != 1 {
		f.derivative = strings.Join([]string{strconv.Itoa(int(coeffint1)), "e^", strconv.Itoa(int(coeffint2)), "x"}, "")
	} else {
		f.derivative = strings.Join([]string{strconv.Itoa(int(coeffint1)), "e^x"}, "")
	}
	return f
}

//Hjälpfunktion som integrerar exponentialfunktioner, tar in en delfunktion som en function och returnerar dess primitiva funktion som en function
func expPrim(f function) function {
	coeff1 := regexp.MustCompile("e").Split(f.fun, 2)
	expcoeff := regexp.MustCompile("\\^").Split(string(coeff1[1]), 2)
	coeff2 := regexp.MustCompile("x").Split(string(expcoeff[1]), 2)
	coeffint1 := parseIntHelp(coeff1[0])
	coeffint2 := parseIntHelp(coeff2[0])
	if coeffint1 == coeffint2 {
		f.primitive = strings.Join([]string{"e^", strconv.Itoa(int(coeffint2)), "x"}, "")
	} else {
		if coeffint2 != 1 {
			f.primitive = strings.Join([]string{strconv.Itoa(int(coeffint1)), "/", strconv.Itoa(int(coeffint2)), "e^", strconv.Itoa(int(coeffint2)), "x"}, "")
		} else {
			f.primitive = strings.Join([]string{strconv.Itoa(int(coeffint1)), "e^x"}, "")
		}
	}
	return f
}

func main() {
	funct := "2e^2x+8e^x-3e^4x/32e^x"
	fmt.Println(Diff(funct))
	fmt.Println(Prim(funct))
}