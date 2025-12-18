package middlewaresUseCases

import "github.com/Nattakornn/cache/modules/middlewares/middlewaresRepositories"

type IMiddlewaresUseCase interface {
}

type middlewaresUseCase struct {
	middlewaresRepository middlewaresRepositories.IMiddlewaresRepository
}

func MiddlewaresUseCase(middlewaresRepository middlewaresRepositories.IMiddlewaresRepository) IMiddlewaresUseCase {
	return &middlewaresUseCase{
		middlewaresRepository: middlewaresRepository,
	}
}
