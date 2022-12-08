package service

import (
	"errors"

	"github.com/GuiRissato/Imersao10-consolidacao/internal/domain/entity"
)

func ChoosePlayers(myTeam entity.MyTeam, players []entity.Player) error{
	// venda e compra de jogadores
	totalCost :=0.0
	totalEarned  := 0.0

	for _, player := range players{
		// venda
		if playerInMyTeam(player, myTeam) && !playerInPlayerList(player, &players){
			totalEarned += player.Price
		}
		//compra
		if !playerInMyTeam(player, myTeam) && playerInPlayerList(player, &players){
			totalCost += player.Price
		}
	}

	if totalCost > myTeam.Score + totalEarned{
		return errors.New("not enough money")
	}

	myTeam.Score += totalEarned - totalCost
	myTeam.Players = []string{}
	// atualiza a lista de jogadores no time
	for _, player := range players{
		myTeam.Players = append(myTeam.Players, player.ID)
	}

	return nil
}

func playerInMyTeam(player entity.Player, myTeam entity.MyTeam) bool{
	for _, playerID := range myTeam.Players {
		if player.ID == playerID{
			return true
		}
	}
	return false
}

func playerInPlayerList(player entity.Player, players *[]entity.Player) bool{
	for _,p := range *players{
		if player.ID	== p.ID{
			return true
		}
	}
	return false
}