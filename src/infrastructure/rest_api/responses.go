package rest

func NullableJsonResponse[T any](data T) NullableResponse[T] {
	response := NullableResponse[T]{}
	response.Body.Data = &data

	return response
}

func NonNullableJsonResponse[T any](data T) NonNullableResponse[T] {
	response := NonNullableResponse[T]{}
	response.Body.Data = &data

	return response
}

func NullJsonResponse() NullResponse {
	response := NullResponse{}
	response.Body.Data = nil

	return response
}
