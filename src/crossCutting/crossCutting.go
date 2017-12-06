package crossCutting

type ProductDto struct {
	Id     int
	Nombre string
}

type ApiErr struct {
	ErrorMessage string `json:"message"`
	ErrorStatus  int    `json:"status"`
}

type MyError struct {
	Error  error
	Status int
}
