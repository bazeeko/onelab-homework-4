package main

import (
	"reflect"
	"testing"
)

func TestRemoveCyrillic(t *testing.T) {
	c := "HELLO ЩЩЩЩЩССФЫЩ ahhahahha"
	someStruct := Alo{
		A: "HELLO ЩЩЩЩЩССФЫЩ ahhahahha",
		B: "HELLO ЩЩЩЩЩССФЫЩ ahhahahha",
		C: &c,
		D: "HELLO ЩЩЩЩЩССФЫЩ ahhahahha",
	}

	nestedStruct := Ola{
		E: someStruct,
		F: "ddo туттутрара лфыл askl",
	}

	expectedC := "HELLO  ahhahahha"
	expectedStruct := Alo{
		A: "HELLO  ahhahahha",
		B: "HELLO  ahhahahha",
		C: &expectedC,
		D: "HELLO  ahhahahha",
	}

	expectedNestedStruct := Ola{
		E: expectedStruct,
		F: "ddo   askl",
	}

	if err := RemoveCyrillic(&nestedStruct); err != nil {
		t.Errorf("TestRemoveCyrillic: %s", err)
	}

	if !reflect.DeepEqual(expectedNestedStruct, nestedStruct) {
		t.Errorf("TestRemoveCyrillic: expected(%v), got(%v)", expectedNestedStruct, nestedStruct)
	}
}
