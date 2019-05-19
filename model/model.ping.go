package model

type PingData struct {
	PingData string `json:"ping_data"`
}

func (p *PingData) GetPing() *PingData {
	return &PingData{PingData: "Pong"}
}
