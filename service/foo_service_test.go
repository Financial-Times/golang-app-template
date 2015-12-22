package actions

import "testing"

func TestDidItDoTheThing(t *testing.T) {
	fooService := NewFooService("field1", "field2")
	if len(fooService.DoTheThing("param")) <= 0 {
		t.Error("The thing was not done!")
	}
}
