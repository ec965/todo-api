package models

import (
	"testing"
)

func TestTableInsertStr(t *testing.T) {
	tt := table{"test", []string{"t1", "t2", "t3"}}
	s := tt.insertStr()
	testStr := "INSERT INTO test ( t1, t2, t3 ) VALUES ( $1, $2, $3 ) RETURNING id"
	if s != testStr {
		t.Errorf("Got '%v', wanted '%v'", s, testStr)
	}
}

func TestTableUpdateStr(t *testing.T) {
	tt := table{"test", []string{"t1", "t2", "t3"}}
	s := tt.updateStr()
	testStr := "UPDATE test SET t1 = $1, t2 = $2, t3 = $3 WHERE id = $4"
	if s != testStr {
		t.Errorf("Got '%v', wanted '%v'", s, testStr)
	}
}

func TestTableSelectByIdStr(t *testing.T){
	tt := table{"test", []string{"t1", "t2", "t3"}}
	s := tt.selectByStr("id")
	testStr := "SELECT t1, t2, t3 FROM test WHERE id = $1"
	if s != testStr {
		t.Errorf("Got '%v', wanted '%v'", s, testStr)
	}
}