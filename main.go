package main

import "github.com/moskyb/sap-solver/solver"

func main() {
	bat := &solver.Battle{
		LeftTeam: &solver.Team{
			Side: solver.Left,
			Friends: []*solver.Pet{
				solver.NewBeaver(),
				solver.NewPig(),
				solver.NewOtter(),
			},
		},
		RightTeam: &solver.Team{
			Side: solver.Left,
			Friends: []*solver.Pet{
				solver.NewBeaver(),
				solver.NewOtter(),
			},
		},
	}

	bat.Run()
}
