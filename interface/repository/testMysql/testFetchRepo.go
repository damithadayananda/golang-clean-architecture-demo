package testMysql

import "golang-clean-architecture/domain"

type FetchRepo struct {

}

func (f FetchRepo)GetAllStudent()[]domain.Student  {
	var result []domain.Student
	for _,v:=range StudentMap{
		result=append(result,v)
	}
	return result
}
