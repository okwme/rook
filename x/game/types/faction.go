package types

func NewFaction(player string, resources *ResourceSet, startingPosition *Position) *Faction {
	return &Faction{
		Player:    player,
		Resources: resources,
		Population: []*Populace{
			{
				Amount:   resources.Population,
				Position: startingPosition,
			},
		},
	}

}

func (f *Faction) Reap() {

}
