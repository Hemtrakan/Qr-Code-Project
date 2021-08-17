package computer

type Computer struct {
	Info        Info   `json:"info"`
	HistoryInfo []Info `json:"history_info"`
	Ops         Ops    `json:"ops"`
}
