package pkg

type ErrorReport struct {
	Msg  string `json:"msg"`
	Err  string `json:"err"`
	Code int    `json:"code,omitempty"`
}

type TrafficReport struct {
	Msg string `json:"msg"`
}

type LogEntry struct {
	Timestamp string
	Entry     interface{}
}
