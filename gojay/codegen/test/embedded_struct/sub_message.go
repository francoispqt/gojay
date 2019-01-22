package embedded_struct

import "time"

type SubMessage struct {
	Description string
	StartTime   time.Time
	EndTime     *time.Time
}
