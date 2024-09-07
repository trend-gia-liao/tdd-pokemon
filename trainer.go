package tdd_pokemon

import "fmt"

type Trainer struct {
	storage Storage
}

type Storage struct {
	pokemon []Pokemon
	sizeMax int
}

func (s *Storage) add(pokemon Pokemon) {
	s.pokemon = append(s.pokemon, pokemon)
}

func (s *Storage) getLeftSpace() int {
	return s.sizeMax - len(s.pokemon)
}

func (s *Storage) CountPVP(no int) int {
	count := 0
	for _, pokemon := range s.pokemon {
		if pokemon.no == no && pokemon.iv.isPVP() {
			count++
		}
	}
	return count
}

func (s *Storage) CountBetter(pokemon Pokemon) int {
	count := 0
	for _, p := range s.pokemon {
		if p.no == pokemon.no && p.iv.Sum() > pokemon.iv.Sum() {
			count++
		}
	}
	return count
}

func (s *Storage) CountShinyBetter(pokemon Pokemon) int {
	count := 0
	for _, p := range s.pokemon {
		if p.no == pokemon.no && p.isShiny() && p.iv.Sum() > pokemon.iv.Sum() {
			count++
		}
	}
	return count
}

type Pokemon struct {
	no    int
	iv    IV
	shiny bool
	event bool
}

var legendaryNo = []int{144, 145, 146, 151}

func (p *Pokemon) isLegendary() bool {
	for _, no := range legendaryNo {
		if p.no == no {
			return true
		}
	}
	return false
}

func (p *Pokemon) isShiny() bool {
	return p.shiny
}

func (p *Pokemon) isEvent() bool {
	return p.event
}

type IV struct {
	Attack  int
	Defense int
	HP      int
}

func (iv *IV) Sum() int {
	return iv.Attack + iv.Defense + iv.HP
}

func (iv *IV) isThreeStar() bool {
	return iv.Sum() >= 37
}

func (iv *IV) isExcellent() bool {
	return iv.Sum() >= 42
}

func (iv *IV) isPVP() bool {
	return iv.Attack < iv.Sum()/3
}

func NewTrainer() *Trainer {
	return &Trainer{storage: Storage{sizeMax: 1000}}
}

func (t *Trainer) Play() error {
	if t.storage.getLeftSpace() <= 0 {
		managed := t.Manage()
		if managed == false {
			return fmt.Errorf("your storage is full. Upgrade storage")
		}
	}

	pokemon, index := t.Gotcha()
	if t.QuickFilter(pokemon) == false {
		_ = t.Transfer(index)
	}

	return nil
}

func (t *Trainer) Manage() bool {
	managed := false
	for i := len(t.storage.pokemon) - 1; i >= 0; i-- {
		pokemon := t.storage.pokemon[i]
		if t.PreciseFilter(pokemon) == false {
			_ = t.Transfer(i)
			managed = true
		}
	}
	return managed
}

func (t *Trainer) QuickFilter(pokemon Pokemon) bool {
	if pokemon.iv.isThreeStar() {
		return true
	}
	if pokemon.iv.isPVP() {
		return true
	}
	if pokemon.isLegendary() {
		return true
	}
	if pokemon.isShiny() {
		return true
	}
	if pokemon.isEvent() {
		return true
	}
	return false
}

func (t *Trainer) PreciseFilter(pokemon Pokemon) bool {
	if pokemon.iv.isExcellent() {
		return true
	}
	if pokemon.iv.isPVP() && t.storage.CountPVP(pokemon.no) <= 1 {
		return true
	}
	if pokemon.isLegendary() && t.storage.CountBetter(pokemon) <= 2 {
		return true
	}
	if pokemon.isShiny() && t.storage.CountShinyBetter(pokemon) <= 0 {
		return true
	}
	if pokemon.isShiny() && pokemon.isLegendary() && t.storage.CountShinyBetter(pokemon) <= 2 {
		return true
	}
	return false
}

func (t *Trainer) Gotcha() (Pokemon, int) {
	pokemon := Pokemon{}
	t.storage.add(pokemon)
	return pokemon, len(t.storage.pokemon) - 1
}

func (t *Trainer) Transfer(index int) error {
	if index < 0 || index >= len(t.storage.pokemon) {
		return fmt.Errorf("index out of range")
	}
	t.storage.pokemon = append(t.storage.pokemon[:index], t.storage.pokemon[index+1:]...)
	return nil
}
