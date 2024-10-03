package utils

type ErrorRes struct {
	IsError    	bool        `json:"isError"`
	StatusCode 	int         `json:"statusCode"`
	Message		string		`json:"message"`
	Error    	interface{} `json:"error"`
}
