package business_ru_api_integration_go

type DataEmployees = struct {
	Status    string     `json:"status"`
	Employees []Employee `json:"result"`
}

type Employee = struct {
	Id       string `json:"id"`
	LastName string `json:"last_name"`
}
