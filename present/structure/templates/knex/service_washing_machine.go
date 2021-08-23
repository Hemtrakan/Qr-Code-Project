package knex

type ServiceWashingMachine struct {
	Info        WashingMachineInfo   `json:"info"`
	Ops         WashingMachineOps    `json:"ops"`
}
