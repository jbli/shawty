package data

import "time"

type Shawty struct {
	ID        uint64
	Rand      string
	Hits      uint64
	Url       string
	CreatorIP string
	CreatedOn time.Time
}
