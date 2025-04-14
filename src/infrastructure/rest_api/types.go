package rest

type NullableResponse[T any] struct {
	Body struct {
		Data *T `json:"data" required:"false"`
	}
}

type NonNullableResponse[T any] struct {
	Body struct {
		Data *T `json:"data" required:"true" nullable:"false"`
	}
}

type NullResponse struct {
	Body struct {
		Data interface{} `json:"data" required:"true" nullable:"false"`
	}
}

type RequestBody[T any] struct {
	Body T
}

type ReferenceByPathIdRequest struct {
	Id string `path:"id" doc:"The id that uniquely identifies the record."`
}
