package operator

type Connection interface{}
type Query interface{}

type Iterator interface {
	Next(v interface{}) bool
}
