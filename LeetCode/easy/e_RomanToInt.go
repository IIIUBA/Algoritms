package main

func romanToInt(s string) int {
	sum := 0
	m := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	for i := 0; i < len(s); i++ {
		if i+1 < len(s) && m[s[i]] < m[s[i+1]] {
			sum -= m[s[i]]
		} else {
			sum += m[s[i]]
		}
	}

	return sum
}

// func main() {
// 	s := "MCMXCIV"
// 	fmt.Println(romanToInt(s))
// }

// создадим мапу
// нужно как то объяснить коду что нужно вычитать числа которые меньше следующего числа и продолжать вычитать если не увеличилось
