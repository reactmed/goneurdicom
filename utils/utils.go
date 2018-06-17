package utils

import (
	"testing"
)

func AssertEqual(actual interface{}, expected interface{}, t *testing.T, ) {
	if actual != expected {
		t.Errorf("%v expected, but actual is %v", expected, actual)
	}
}

func AssertEqualf(a interface{}, b interface{}, t *testing.T, format string, args ...interface{}) {
	if a != b {
		t.Errorf(format, args)
	}
}

func AssertNil(a interface{}, t *testing.T) {
	if a != nil {
		t.Error("Should be nil")
	}
}

func AssertNilf(a interface{}, t *testing.T, format string, args ...interface{}) {
	if a != nil {
		t.Error(format, args)
	}
}

func AssertNotNil(a interface{}, t *testing.T) {
	if a == nil {
		t.Error("Should not be nil")
	}
}

func AssertNotNilf(a interface{}, t *testing.T, format string, args ...interface{}) {
	if a == nil {
		t.Error(format, args)
	}
}

func AssertTrue(actual bool, t *testing.T, ) {
	if !actual {
		t.Errorf("True is expected")
	}
}

func AssertTruef(a bool, t *testing.T, format string, args ...interface{}) {
	if !a {
		t.Errorf(format, args)
	}
}

func AssertFalse(actual bool, t *testing.T, ) {
	if actual {
		t.Errorf("False is expected")
	}
}

func AssertFalsef(a bool, t *testing.T, format string, args ...interface{}) {
	if a {
		t.Errorf(format, args)
	}
}

func AssertHasLen(a []interface{}, length int, t *testing.T) {
	if len(a) != length {
		t.Errorf("Should has length %d", length)
	}
}

func AssertHasLenf(a []interface{}, length int, t *testing.T, format string, args ...interface{}) {
	if len(a) != length {
		t.Errorf(format, args)
	}
}

func AssertNotErr(err error, t *testing.T){
	if err != nil{
		t.Error(err)
	}
}