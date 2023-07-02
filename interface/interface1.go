package interfaces

type Dog struct {
	Id   string
	Name string
	Age  string
	Sex  string
}

type DogDataBase map[string]Dog

func (d DogDataBase) Get(id string) string {
	return d[id].Name
}

func (d DogDataBase) Post(id string, name string) bool {
	d[id] = Dog{Name: name}
	return true
}

type Cat struct {
	Id   string
	Name string
	Age  string
	Sex  string
}

type CatDataBase map[string]Cat

func (c CatDataBase) Get(id string) string {
	return c[id].Name
}

func (c CatDataBase) Post(id string, name string) bool {
	c[id] = Cat{Name: name}
	return true
}

type ChooseAnimal interface {
	Get(id string) string
	Post(id string, name string) bool
}

type MyService struct {
	a ChooseAnimal
}

func NewAnimalService(a ChooseAnimal) MyService {
	return MyService{
		a: a,
	}
}

func RunMain() {
	// ddb := DogDataBase{}
	cdb := CatDataBase{}

	as := NewAnimalService(cdb)

	as.a.Get("1")
}
