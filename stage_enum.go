package main

type fash_stage_enum struct{}

func (fsg fash_stage_enum) murder() string {
	return "murder"
}
func (fsg fash_stage_enum) spyguy() string {
	return "spyguy"

}
func (fsg fash_stage_enum) spydeck() string {
	return "spydeck"

}
func (fsg fash_stage_enum) special_election() string {
	return "special_election"

}
func (fsg fash_stage_enum) veto() string {
	return "veto"
}

func (fsg fash_stage_enum) none() string {
	return "none"
}

type game_stage_enum struct {
}

func (gse game_stage_enum) election() string {
	return "election"
}

//	func (gse game_stage_enum) specialelection() string {
//		return "specialelection"
//	}
func (gse game_stage_enum) policy() string {
	return "policy"
}
func (gse game_stage_enum) fascistpower() string {
	return "fascist power"
}
