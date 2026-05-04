package model

type ResponseApi[T any] struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    T      `json:"data"`
}

type ResponseApiWithoutData struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

type ResponseApiError struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Error   string `json:"error"`
}

type WithPagination[T any] struct {
	List T `json:"list"`
	Pagination Pagination `json:"pagination"`
}

type Pagination struct {
	Page int `json:"page"`
	PerPage int `json:"perPage"`
	Total int `json:"total"`
	LastPage int `json:"lastPage"`
}