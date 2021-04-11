package operator

type Query interface {
	Get(query interface{}) error
}

type Connection interface {
	Execute(q Query, result interface{})
}
