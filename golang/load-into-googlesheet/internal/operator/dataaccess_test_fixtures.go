package operator

type MockConnection struct{}

func (c *MockConnection) Execute(q Query, result interface{}) {}

type MockQuery struct{}

func (q *MockQuery) Get(query interface{}) error {
	return nil
}
