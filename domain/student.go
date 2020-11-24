package domain

type Student struct {
	Name string `json:"name"`
	Age int `json:"age"`
}

type Registry interface {
	NameValidate(string)bool
	AgeValidate(int)bool
	SaveStudent(Student)bool
}

type Fetch interface {
	GetAllStudent()[]Student
}
