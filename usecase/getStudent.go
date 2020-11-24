package usecase

import "golang-clean-architecture/domain"

type GetStudent struct {
	Repo domain.Fetch
}

func (s GetStudent)GetAllStudent()[]domain.Student{
	return s.Repo.GetAllStudent()
}
