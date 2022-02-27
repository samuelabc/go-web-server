package helloModel

type HelloRequest struct {
	Name string `json:"name" validate:"required"`
	Age  int    `json:"age"`
}

type HelloResponse struct {
	Name        string `json:"name" validate:"required"`
	Age         int    `json:"age"`
	Description string `json:"description"`
	Timestamp   string `json:"timestamp"`
}
