package antifraurd

import (
	"solution/internal/repository"
	def "solution/internal/service"
)

var _ def.AntifraudService = (*service)(nil)

type service struct {
	antifraudRepository repository.AntifraudRepository
}

func NewService(antifraudRepository repository.AntifraudRepository) *service {
	return &service{
		antifraudRepository: antifraudRepository,
	}
}
