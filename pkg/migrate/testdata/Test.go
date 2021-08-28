package testdata

type Test struct {
	ID int `json:"id"`
}

func (Test) TableName() string {
	return "test"
}
