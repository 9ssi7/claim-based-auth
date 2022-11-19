package result

type Result struct {
	Success bool   `json:"success"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type DataResult struct {
	Success bool   `json:"success"`
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func (r *Result) Error() string {
	return r.Message
}

func (r *DataResult) Error() string {
	return r.Message
}

func Success(m string, c int) *Result {
	return &Result{
		Success: true,
		Message: m,
		Code:    c,
	}
}

func Error(m string, c int) *Result {
	return &Result{
		Success: false,
		Message: m,
		Code:    c,
	}
}

func SuccessData(m string, d any, c int) *DataResult {
	return &DataResult{
		Success: true,
		Message: m,
		Data:    d,
		Code:    c,
	}
}

func ErrorData(m string, d any, c int) *DataResult {
	return &DataResult{
		Success: false,
		Message: m,
		Data:    d,
		Code:    c,
	}
}