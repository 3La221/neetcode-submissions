


func isAnagram(s string, t string) bool {
if len(s) != len(t){
	return false
}

check := make(map[rune]int)

for _,v := range s {
	check[v]++
}

for _,v := range t {
	if check[v] == 0 {
		return false
	}

	check[v]--
}

return true




}
