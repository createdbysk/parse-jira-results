package operator

type Query interface {
	Get(result interface{}) error
}

type Iterator interface {
	Next(v interface{}) bool
}
type Connection interface {
	Execute(q Query) Iterator
}
