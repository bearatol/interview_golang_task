package mapping

import "time"

const (
	TimeoutConnect = time.Minute * 5
	MaxMsgSize     = 1024 * 1024 * 10 // max message size 10 MiB
)
