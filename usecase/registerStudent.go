package usecase

import "golang-clean-architecture/domain"

type RegisterStudent struct {
	Repo domain.Registry
}

func (r RegisterStudent)RegisterStudent(student domain.Student)bool{
	if r.Repo.AgeValidate(student.Age){
		if r.Repo.NameValidate(student.Name){
			 return r.Repo.SaveStudent(student)
		}
	}
	return false
}
