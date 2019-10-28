package combination

import "fmt"

var (
	k      = 4
	values = []string{"x", "o", "-"}
)

func printAllKLength(set []string, k int) {
	n := len(set)
	printAllKLengthRec(set, "", n, k)
}

func printAllKLengthRec(set []string, prefix string, n int, k int) {
	if k == 0 {
		fmt.Println(prefix)
		return
	}
	for i := 0; i < n; i++ {
		newPrefix := prefix + set[i]
		printAllKLengthRec(set, newPrefix, n, k-1)
	}
}

func main() {
	//So much slower than python... Benchmark ?
	printAllKLength(values, k)
}
