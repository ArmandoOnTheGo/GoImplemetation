package morestrings //Package Name

func ReverseRunes (s string) string {
	r := []rune(s)
	//starts by instanstianting i to 0 and j to len(r) 
	//then the condition to stay in the for loop "i < len(r)"
	//then i and j are equal to i + 1 and j - 1 each iteration
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
