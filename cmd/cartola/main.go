package main

import (
	"context"
	"database/sql"

	"github.com/GuiRissato/ms-consolidacao/internal/infra/db"
	"github.com/GuiRissato/ms-consolidacao/internal/infra/repository"
	"github.com/GuiRissato/ms-consolidacao/pkg/uow"
)

func main() {
	ctx := context.Background()
	// abre o db
	dtb, err := sql.Open("mysql","root:root@tcp(localhost:3306)/cartola?parseTime=true")
	if err != nil{
		panic(err)
	}
	// fecha o db
	defer dtb.Close()

	uow, err := uow.NewUow(ctx,dtb)
	if err != nil{
		panic(err)
	}
	registerRepositories(uow)
}

func registerRepositories(uow *uow.Uow){
	
	uow.Register("PlayersRepository", func(tx *sql.Tx) interface{}{
		repo := repository.NewPlayerRepository(uow.Db)
		repo.Queries = db.New(tx)
		return repo
	})

	uow.Register("MatchRepository", func(tx *sql.Tx) interface{}{
		repo := repository.NewMatchRepository(uow.Db)
		repo.Queries = db.New(tx)
		return repo
	})

	uow.Register("TeamRepository", func(tx *sql.Tx) interface{}{
		repo := repository.NewTeamRepository(uow.Db)
		repo.Queries = db.New(tx)
		return repo
	})

	uow.Register("MyTeamRepository", func(tx *sql.Tx) interface{}{
		repo := repository.NewMyTeamRepository(uow.Db)
		repo.Queries = db.New(tx)
		return repo
	})
}