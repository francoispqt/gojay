package basic_struct

import "time"

type SubMessage struct {
	Id          int
	Description string
	StartTime   time.Time
	EndTime     *time.Time
}
