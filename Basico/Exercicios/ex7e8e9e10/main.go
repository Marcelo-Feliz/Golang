package main

import "fmt"

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}
type sedan struct {
	vehicle
	luxury bool
}

func (t truck) transportationDevice() string {
	return fmt.Sprint("I am a Truck with ", t.doors, " doors and i am ", t.color)
}

func (s sedan) transportationDevice() string {
	return fmt.Sprint("I am a Sedan with ", s.doors, " doors and i am ", s.color)
}

type transportation interface {
	transportationDevice() string
}

func report(t transportation) {
	fmt.Println(t.transportationDevice())
}

func main() {

	t1 := truck{
		vehicle{
			5,
			"Black",
		},
		true,
	}
	s1 := sedan{
		vehicle{
			7,
			"Yellow",
		},
		true,
	}

	fmt.Println(t1)
	fmt.Println(s1)

	fmt.Println(t1.color)
	fmt.Println(s1.color)

	s := t1.transportationDevice()
	fmt.Println(s)
	s = s1.transportationDevice()
	fmt.Println(s)

	report(t1)
	report(s1)

}
