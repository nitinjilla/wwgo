package main

import (
	"fmt"
	"maps"
	"slices"
)

var a int = 5

func main() {

	// hello world
	fmt.Print("Hello World!\n")

	// variable
	var aa int = 0
	ab := 10
	fmt.Println(aa, ab)

	// constants
	const d = 5
	//d = d + 10	  // error: cannot assign to d (neither addressable nor a map index expression)
	fmt.Println(d + 10) // 15

	// for loop
	for i := range 5 {
		fmt.Println(i)
	}

	/*/ infinite loop
	for {
		fmt.Println("I promise to never use infinite loop again")
	}*/

	// if - else statements
	if ia := 1; ia > 1 {
		fmt.Println(ia) // we do not see the undefined error in if-else statements
		fmt.Println("ia is greater than 1")
	} else {
		fmt.Println(ia)
		fmt.Println("ia is less than or equal to 1")
	}

	//fmt.Println(i) 		// undefined: i (scope ends in the if-else statements)

	// switch case
	switch a {
	case 1:
		fmt.Println("Not the global variable!")
	case 2:
		fmt.Println("This isn't a global variable, too!")
	default:
		// if none of the cases would fit, and default statement will be printed.
		fmt.Println("It's 5.")
	}

	// arrays
	var a [4]int
	fmt.Println(a) // [0 0 0 0]
	a[1] = 10
	fmt.Println(a)              // [0 10 0 0]
	fmt.Println(len(a), cap(a)) // 4
	b := [4]int{1, 2, 3, 4}     // [1 2 3 4]
	fmt.Println(b)

	c := make([]int, 4)
	c[3] = 5
	fmt.Printf("%v, %T\n", c, c) // [0 0 0 5], []int

	// slices
	co := make([]string, 4)
	co[0], co[1], co[2], co[3] = "a", "b", "c", "d"
	cp := make([]string, 4)
	cq := []string{"a", "b"}
	//cq[2] = "c"                                       // panic: runtime error: index out of range [2] with length 2
	fmt.Println(co, len(co), cap(co), co[:2])           // [a b c d] 4 4 [a b]
	fmt.Println(cp, len(cp), cap(cp), copy(cp, co), cp) // [a b c d] 4 4 4 [a b c d]
	fmt.Println(cp, cq, len(cq), cap(cq))               // [a b  ] [a b] 2 2
	slices.Reverse(co)
	fmt.Println(co) // [d c b a]
	cp[0], cp[1] = "p", "q"
	fmt.Println(cap(cp), slices.Clip(cp), cap(cp)) // 4 [p q c d] 4

	//maps
	m := make(map[int]string)

	m[1], m[2], m[3], m[0] = "first", "second", "delete", "void"
	fmt.Println(m, len(m)) // map[0:void 1:first 2:second 3:delete] 4

	_, n := m[2]
	fmt.Println(n) // true

	mc := maps.Clone(m)
	fmt.Println("clone:", mc) // clone: map[0:void 1:first 2:second 3:delete]

	delete(m, 3)
	fmt.Println(m) // map[0:void 1:first 2:second]

	clear(m)
	fmt.Println(m) // map[]

	// range

	for _, v := range m { // for k, v := range m - will print both keys and values
		fmt.Println(v)
	}

	for _, uc := range "message" {
		fmt.Printf("%v:%v ", string(uc), uc) // m:109 e:101 s:115 s:115 a:97 g:103 e:101
	}

}
