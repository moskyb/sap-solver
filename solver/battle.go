package solver

import "fmt"

type Battle struct {
	LeftTeam  *Team
	RightTeam *Team
}

func (b *Battle) Run() {
	for {
		if b.hasEnded() {
			break
		}

		fmt.Printf("%v vs %v\n", b.LeftTeam, b.RightTeam)

		firstLeft, firstRight := b.LeftTeam.Friends[0], b.RightTeam.Friends[0]

		fmt.Printf("%v deals %d damage to %v\n", firstLeft, firstLeft.Power, firstRight)
		fmt.Printf("%v deals %d damage to %v\n", firstRight, firstRight.Power, firstLeft)
		firstLeft.Clash(firstRight)

		b.LeftTeam.CollectFainted()
		b.RightTeam.CollectFainted()
	}

	fmt.Printf("Final state: ")
}

func (b *Battle) hasEnded() bool {
	if len(b.LeftTeam.Friends) == 0 && len(b.RightTeam.Friends) == 0 {
		fmt.Println("It's a draw!")
		return true
	}

	if len(b.LeftTeam.Friends) == 0 {
		fmt.Println("Right wins!")
		return true
	}

	if len(b.RightTeam.Friends) == 0 {
		fmt.Println("Left wins!")
		return true
	}

	return false
}
