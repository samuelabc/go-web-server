package healthRoute

type PostHealthSchema struct {
	Name string `json:"name" validate:"required,min=3,max=10"`
	Age  int    `json:"age"`
}

func GetPostHealthSchemaObject(x interface{}) {
	x = PostHealthSchema{}
}
