// primeira linha sempre o nome do pacote, no caso a pasta em que esta
package entity

type Player struct{
	ID string
	Name string
	Price float64
}

func NewPlayer(id,name string, price float64) *Player { //ponteiro
	return &Player{
		ID: id,
		Name: name,
		Price: price,
	}
}