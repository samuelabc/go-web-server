package helloRoute

type PostHelloSchema struct {
	Name string `json:"name" validate:"required,min=3,max=10"`
	Age  int    `json:"age"`
}
