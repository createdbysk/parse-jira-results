package operator

type Connection interface {
	Execute() Iterator
}

type Query interface{}

type Iterator interface {
	Next(v interface{}) bool
}
