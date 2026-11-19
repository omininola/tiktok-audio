package types

type Monitor struct {
	ID int `json:"id"`
	X  int `json:"x"`
}

type Workspace struct {
	ID              int    `json:"id"`
	MonitorID       int    `json:"monitorID"`
	LastWindowTitle string `json:"lastwindowtitle"`
}

type Audio struct {
	ID         int             `json:"index"`
	Properties AudioProperties `json:"properties"`
}

type AudioProperties struct {
	Media string `json:"media.name"`
}
