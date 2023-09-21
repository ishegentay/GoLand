package Engineer

type Engineer struct {
	position string
	salary   float64
	address  string
}

func NewEngineer(position string, salary float64, address string) *Engineer {
	return &Engineer{
		position: position,
		salary:   salary,
		address:  address,
	}
}

func (d *Engineer) GetPosition() string {
	return d.position
}

func (d *Engineer) SetPosition(position string) {
	d.position = position
}

func (d *Engineer) GetSalary() float64 {
	return d.salary
}

func (d *Engineer) SetSalary(salary float64) {
	d.salary = salary
}

func (d *Engineer) GetAddress() string {
	return d.address
}

func (d *Engineer) SetAddress(address string) {
	d.address = address
}
