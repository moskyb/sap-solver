package solver

import "fmt"

type Pet struct {
	Power  int
	Health int
	Name   string
	Emoji  string

	hasFainted bool
}

func NewPig() *Pet {
	return &Pet{
		Power:  3,
		Health: 1,
		Name:   "Pig",
		Emoji:  "🐷",
	}
}

func NewOtter() *Pet {
	return &Pet{
		Power:  1,
		Health: 2,
		Name:   "Otter",
		Emoji:  "🦦",
	}
}

func NewBeaver() *Pet {
	return &Pet{
		Power:  2,
		Health: 2,
		Name:   "Beaver",
		Emoji:  "🦫",
	}
}

func (p *Pet) Hurt(amount int) {
	p.Health -= amount
	if p.Health <= 0 {
		p.Faint()
	}
}

func (p *Pet) Faint() {
	p.hasFainted = true
}

func (p *Pet) Clash(other *Pet) {
	p.Hurt(other.Power)
	other.Hurt(p.Power)
}

func (p *Pet) String() string {
	return fmt.Sprintf("%s {%d⚔️,  %d❤️ }", p.Emoji, p.Power, p.Health)
}
