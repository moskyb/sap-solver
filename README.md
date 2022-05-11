# SAP Solver

The goal of this project is to create a system that can accurately predict the outcome of a match of the game [Super Auto Pets](https://teamwoodgames.com/).

## Currently Working

At current, this project can accurately simulate a game of Super Auto Pets where there are no abilities. That's basically it.

## Things I'd like to do

Simulate:
- Abilities (things that get triggered during combat)
- Held foods (Melon/Garlic/Coconut armor, meat bone, etc)
- Decision tree descent - ie, some abilities are non-deterministic, and mean that given two teams, the outcome of a match is determined by RNG. Ideally, we'd like to figure out, given two teams, the winrate of one team over another.

## Running the solver

This solver is implemented using Go 1.18. Install Go 1.18 and later, then run

```
go run main.go
```

to simulate a hardcoded battle: [🦦 🐷 🦫] vs [🦫 🦦]
