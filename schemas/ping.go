package schemas

import "encoding/json"

type Ping struct {
	Success bool `json:"success"`
}

func Create() ([]byte, error) {
	return json.Marshal(&Ping{
		Success: true,
	})
}
