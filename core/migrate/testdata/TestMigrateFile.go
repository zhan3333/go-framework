package testdata

import (
	gdb2 "go-framework/core/gdb"
)

type TestFile struct {
}

func (TestFile) Key() string {
	return "TestFile"
}
func (TestFile) Up(db *gdb2.Entry) error {
	db.Exec("create table test (id int)")
	return nil
}
func (TestFile) Down(db *gdb2.Entry) error {
	db.Exec("drop table test")
	return nil
}
