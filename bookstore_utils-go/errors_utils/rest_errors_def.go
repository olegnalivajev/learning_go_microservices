package errors_utils

type RestErr struct {
	Message string `json:"message"`
	Status 	int	   `json:"status"`
	Error	string `json:"error"`
}