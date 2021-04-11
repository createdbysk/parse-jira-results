package operator

type Connection interface {
	Get(impl interface{})
}

type Output interface {
	Write(c Connection, data interface{}) error
}
