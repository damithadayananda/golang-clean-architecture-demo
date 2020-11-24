package testMysql

import "golang-clean-architecture/domain"

var StudentMap map[string]domain.Student

type TestRegistryRepo struct {

}

func (t TestRegistryRepo)NameValidate(string)bool{
	return true
}
func (t TestRegistryRepo)AgeValidate(int)bool  {
	return true
}
func (t TestRegistryRepo)SaveStudent(s domain.Student)bool {
	StudentMap[s.Name]=s
	return true
}





func init()  {
	StudentMap= make(map[string]domain.Student)
}