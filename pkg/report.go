package pkg

type ErrorReport struct {
	Msg  string `json:"msg"`
	Err  error  `json:"err"`
	Code int    `json:"code,omitempty"`
}

type TrafficReport struct{}
