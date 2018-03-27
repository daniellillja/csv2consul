package main

import (
	"testing"
)

func Test_EnvironmentKvpPut(t *testing.T) {
	kvpEnv := EnvironmentKvp{
		Environment: "env",
		Key:         "key",
		Value:       "value",
	}

	sut := NewConsulClient()
	err := sut.EnvironmentKvpPut(kvpEnv)
	if err != nil {
		t.Error(err)
	}
}
