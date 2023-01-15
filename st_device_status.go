package smartthings

import (
	"time"
)

type StDeviceStatus struct {
	Switch struct {
		Timestamp time.Time `json:"timestamp"`
		Value     string    `json:"value"`
	} `json:"switch"`
}
