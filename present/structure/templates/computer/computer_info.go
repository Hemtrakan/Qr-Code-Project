package computer

type Info struct {
	Case        string `json:"case"`
	PowerSupply string `json:"power_supply"`
	MainBoar    string `json:"main_boar"`
	CPU         string `json:"cpu"`
	RAM         string `json:"ram"`
	GraphicCard string `json:"graphic_card"`
	HardDisk    string `json:"hard_disk"`
}