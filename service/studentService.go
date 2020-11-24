package service

import (
	"golang-clean-architecture/domain"
	_interface "golang-clean-architecture/interface"
	"golang-clean-architecture/usecase"
)

func RegisterStudent(student domain.Student)bool{
	registryUsecase:=usecase.RegisterStudent{
		Repo: _interface.MySqlRegistryRepo,
	}
	if registryUsecase.RegisterStudent(student){
		return _interface.KafkaProducer.Produce(student)
	}
	return false
}

func GetAllStudent()[]domain.Student{
	fetchUseCase:=usecase.GetStudent{
		Repo:_interface.MySqlFetchRepo,
	}
	return fetchUseCase.GetAllStudent()
}
