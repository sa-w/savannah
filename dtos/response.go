package dtos

type ResponseDTO[T any] struct {
	Status  bool
	Code    int
	Message string
	Data    T
}
