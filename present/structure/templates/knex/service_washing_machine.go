package knex

type ServiceWashingMachine struct {
	Info        WashingMachineInfo   `json:"info"`
	HistoryInfo []WashingMachineInfo `json:"history_info"`
	Ops         WashingMachineOps    `json:"ops"`
}
