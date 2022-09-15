package ports

import "sodouku/internal/core/domain"

type GameService interface {
	Create(name string, size uint, dificulty string) (domain.Game, error)
}
