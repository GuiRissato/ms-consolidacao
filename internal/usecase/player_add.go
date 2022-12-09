package usecase
// caso de uso, pegar o player e executar
import (
	"context"

	"github.com/GuiRissato/ms-consolidacao/internal/domain/entity"
	"github.com/GuiRissato/ms-consolidacao/internal/domain/repository"
	"github.com/GuiRissato/ms-consolidacao/pkg/uow"
)

type AddPlayerInput struct{
	ID string
	Name string
	InitialPrice float64
}

type AddPlayerUseCase struct{
	Uow uow.UowInterface
}

func (a *AddPlayerUseCase) Execute(ctx context.Context, input AddPlayerInput) error{
	playerRepository := a.getPlayerRepository(ctx)
	player := entity.NewPlayer(input.ID,input.Name,input.InitialPrice)
	err := playerRepository.Create(ctx,player)
	if err != nil{
		return err
	}
	// isso é o que o unit of work, caso de uso
	return a.Uow.CommitOrRollback()
	
}

func (a *AddPlayerUseCase) getPlayerRepository(ctx context.Context) repository.PlayerRepositoryInterface{
	playerRepository, err := a.Uow.GetRepository(ctx, "PlayerRepository")
	if err != nil{
		panic(err)
	}
	return playerRepository.(repository.PlayerRepositoryInterface)
}