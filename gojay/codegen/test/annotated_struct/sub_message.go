package annotated_struct

import "time"

type SubMessage struct {
	Id          int        `json:"id"`
	Description string     `json:"description"`
	StartTime   time.Time  `json:"startDate" timeFormat:"yyyy-MM-dd HH:mm:ss"`
	EndTime     *time.Time `json:"endDate" timeLayout:"2006-01-02 15:04:05"`
}
