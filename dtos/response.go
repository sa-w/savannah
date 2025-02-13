package dtos

//Generic type for all responses
type ResponseDTO[T any] struct {
	Status  bool
	Code    int
	Message string
	Data    T
}
