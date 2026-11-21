package types

type Monitor struct {
	ID int `json:"id"`
	X  int `json:"x"`
}

type Client struct {
	Title   string `json:"title"`
	Monitor int    `json:"monitor"`
	PID     int    `json:"pid"`
}

type Audio struct {
	ID         int             `json:"index"`
	Properties AudioProperties `json:"properties"`
}

type AudioProperties struct {
	Media string `json:"media.name"`
	PID   string `json:"application.process.id"`
}
