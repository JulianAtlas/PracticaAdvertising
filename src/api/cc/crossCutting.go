package cc

type ProductDto struct {
	Id   int
	Name string
}

type ApiErr struct {
	ErrorMessage string `json:"message"`
	ErrorStatus  int    `json:"status"`
}

type MyError struct {
	Error  error
	Status int
}
