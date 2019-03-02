package 第0阶段

import "fmt"

type Person struct {
	name string
	sex  byte
	int
}

func main01() {
	array := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s1 := array[:6:8]
	fmt.Println(len(s1))
	fmt.Println(cap(s1))
	s := []int{1, 2, 3, 4, 5}
	fmt.Println(s)
}
func sortScores(scores *[10]int) {
	for i := 0; i < len(scores); i++ {
		for j := i; j < len(scores); j++ {
			if scores[i] < scores[j] {
				scores[i], scores[j] = scores[j], scores[i]
			}
		}

	}
}
func main() {
	var scores [10]int = [10]int{80, 99, 39, 23, 91, 23, 56, 32, 56, 67}
	sortScores(&scores)
	fmt.Println(scores[:3])
}
