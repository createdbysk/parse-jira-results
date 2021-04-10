package operator

type Query interface {
	Get(result interface{}) error
}

type Connection interface {
	Execute(q Query, data interface{})
}
