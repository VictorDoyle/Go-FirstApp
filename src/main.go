package main

import (
	"fmt"
	"math"
)

/* Must intl "nil" with an explicit type */
func main() {
	/* var declares multiple + the := syntax declaeres and initializes variables */
	var a = "initial variable"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	/* base value for int is 0 */
	var d int
	fmt.Println(d)

	const n = 500000000

	const f = 3e20 / n
	/* expect a float64 -- calc returns -0.28470407323754404 */
	fmt.Println(math.Sin(n))

	var x interface{} = nil

	/* for loops */
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	for {
		fmt.Println("loop")
		/* stop loop */
		break
	}

	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			/* continue cmd for loop */
			continue
		}
		fmt.Println(n)
	}
	_ = x
}
