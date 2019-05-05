package object

type Person struct {
	Age int
}

func (p *Person) SetAge(age int) {
	//(*p).Age = age
	p.Age = age
}


