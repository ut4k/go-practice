package person

type Person struct {
	firstName string
	lastName  string
}

func (p *Person) FirstName() string { return p.firstName }

func (p *Person) LastName() string { return p.lastName }
