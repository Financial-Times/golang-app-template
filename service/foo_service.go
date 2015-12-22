package actions

// FooService does the thing
type FooService struct {
	myField1 string
	myField2 string
}

// NewFooService creates a new FooService
func NewFooService(field1 string, field2 string) *FooService {
	return &FooService{field1, field2}
}

func (se *FooService) DoTheThing(parameter string) string {
	return "I did the thing!"
}
