/*package main

import "fmt"

func main() {
	//m := [5]int{1, 2, 3, 4, 5}
	var e [5][5]int
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			e[i][j] = i + j
		}
	}
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			fmt.Printf("%v\t", e[i][j])
		}
		fmt.Println()
	}
	fmt.Println(len(e))
}*/

//m := [5]int{1, 2, 3, 4, 5}
//this is an array bcoz size is mentioned explicitly
//m := []int{1, 2, 3, 4, 5}
//this is a slice as the size isnt mentioned and can chnage overtime

/*package main

import "fmt"

func main() {
	// s := []string{"k", "jj", "nue"}
	// fmt.Println(len(s))
	var s []string
	s = make([]string, 3)
	fmt.Println(s, len(s), s == nil)
	s[0] = "hiii"
	s[1] = "wassup"
	s[2] = "ntng"
	fmt.Println(s)
	//s[3] = "dumdmdmdmddmddm" wronng way of extending the slice
	// append is the right way to do it make sure to hold the value of the new slice
	s = append(s, "ddimdiddm")
	fmt.Println(s)
}*/

//maps

/*package main

import (
	"fmt"
)

func main() {
	a := make(map[string]int)
	a["first"] = 1
	a["second"] = 2
	fmt.Println(a)
	fmt.Println(a["okok"])
	b := map[string]string{"apple": "fruit", "cucumber": "veggie"}
	fmt.Println(b)
	_, prs := b["mango"]
	fmt.Println(prs)
	delete(b, "apple")
	fmt.Println(b)
	clear(a)
	fmt.Println(b)
}*/

/*package main

import "fmt"

func main() {
	a := make(map[string]string)
	a["apple"] = "magi"
	a["magi"] = "bagu"
	a["bagu"] = "dumb"
	for i, j := range a {
		fmt.Println(i, j)
	}
}*/

/*package main
//this comes under the lazy evaluation please do consider uploading it on git it will be really helpful in tracking the errors and the people who have contributed to the code thanks a lot for asking i will get back to you
//this for sure will turn out tp be a productive session i will gaurantee you of that :wink
import "fmt"

func add(a, b, c int) {
	//return a + b - c
	fmt.Println("umm to lazy to calclularte")
}
func main() {
	var a, b, c int = 1, 2, 3
	//d := add(a, b, c)
	//fmt.Println(d)
	add(a, b, c)
}*/

// multiple return values
/*package main

import "fmt"

func twonums(a, b int) (int, int) {
	return a, b
}
func main() {
	var a, b int = 10, 20
	m, n := twonums(a, b)
	fmt.Println(m, n)
	_, c := twonums(m, n)
	// blank identifier It allows you to ignore values returned from a function or assigned to variables when you don't intend to use them.
	fmt.Println(c)

}*/

//variadic functions

/*package main

import "fmt"

func sums(nums ...int) int {

	res := 0
	for i := range nums {
		res = res + nums[i]
	}

	return res
}
func main() {
	fmt.Println(sums(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
	e := []int{1, 2, 3, 4, 5}
	fmt.Println(sums(e...))
}*/

/*package main

import "fmt"

func intseq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
//val maintains the internal state and therefore the value is updated every time
func main() {
	val := intseq()
	fmt.Println(val())
	fmt.Println(val())
	fmt.Println(val())

}*/

// recusrsion
/*package main

import "fmt"

func fact(n int) int {
	if n == 0 {
		return 1
	} else {
		return n * fact(n-1)
	}
}
func fib(n int) int {
	if n < 2 {
		return n
	} else {
		return fib(n-1) + fib(n-2)
	}
}

func main() {
	fmt.Println(fact(5))
	fmt.Println(fib(7))
	var fib1 func(n int) int
	fib1 = func(n int) int {
		if n < 2 {
			return n
		} else {
			return fib(n-1) + fib(n-2)
		}
	}
	fmt.Println(fib1(7))

}*/

/*package main

import "fmt"

func zeroval(i int) {
	i = 420
}
func zeroptr(i *int) {
	*i = 420
}
func main() {
	i := 69
	fmt.Println(i)
	zeroval(i)
	fmt.Println(i)
	zeroptr(&i)
	fmt.Println(i)
	//zeroval(i)

}*/

// rune in go lang represents a unicode code point used to represent a single unicode charecter
// unicode aims to give every character in the Unicode standard a unique integer code
// rune in lamen terms it represents the current character

/*package main

import "fmt"

func main() {
	str := "hi how are you ನಮಸ್ಕಾರ"
	for _,i := range str {
		fmt.Printf("%c\t", str[i])
	}
}*/

// simple structure
/*package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	person := Person{Name: "khushi", Age: 20}
	fmt.Println(person.Name, person.Age)
}*/

/*package main

import "fmt"

type Person struct {
	name string
	age  int
}

func newpern(name string, age int) *Person {
	p := Person{name: name, age: age}

	return &p
}

func main() {
	fmt.Println(Person{"khuhsiu", 20})
	fmt.Println(Person{name: "khshiuii0", age: 50})
	fmt.Println(newpern("wowow", 20))
	//structure for a single entity
	dog := struct {
		name string
		age  int
	}{
		"timmy",
		50,
	}
	fmt.Println(dog)
}*/

/*package main

import "fmt"

type own struct {
	Name   string
	Age    int
	Hobby  string
	Gender string
}

func wow(name string, age int) *own {
	p := own{Name: name}
	p.Age = age
	p.Hobby = "aimless"
	p.Gender = "neutral"
	return &p

}
func main() {
	p := own{"khushi", 19, "cooking", "F"}
	q := own{Gender: "Fak", Name: "midu", Age: 12, Hobby: "art and craft"}
	fmt.Println(p.Name)
	fmt.Println(q.Gender)
	a := wow("bruhhhh", 55)
	fmt.Println(a)

}*/

/*package main

import "fmt"

type idk struct {
	width   int
	breadth int
}

func (r *idk) area() int {
	return r.width * r.breadth
}
func main() {
	p := idk{10, 20}
	q := idk{breadth: 20, width: 60}
	fmt.Println(p.area())
	fmt.Println(q.area())
	m := &q
	fmt.Println()
	fmt.Println(m.area())

}*/

/*package main

import "fmt"

type Shape interface {
	Area() float64
	Perimeter() float64
}
type Circle struct {
	radius float64
}

func (c *Circle) Area() float64 {
	return 3.14 * c.radius * c.radius
}

type Rect struct {
	width, breadth float64
}

func (r *Rect) Area() float64 {
	return r.width * r.breadth
}
func main() {
	cc := Circle{5.0}
	fmt.Println(cc.Area())
}*/

//struct embedding

