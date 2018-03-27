package main

import (
	"strings"
	"testing"
)

const testCsv = `key,env1,env2
key,val1,val2`

func Test_LoadCsv(t *testing.T) {
	strReader := strings.NewReader(testCsv)
	sut := NewKvpEnvironmentReader(strReader)
	models, err := sut.KvpEnvironmentRead()
	if err != nil {
		t.Error(err)
	}

	if len(models) != 2 {
		t.Error("should return 2 records")
	}

	firstRecord := models[0]
	expectedKey := "key"
	actualKey := firstRecord.Key
	if expectedKey != actualKey {
		t.Error("key incorrect")
	}
	expectedEnv := "env1"
	actualEnv := firstRecord.Environment
	if expectedEnv != actualEnv {
		t.Error("environment incorrect")
	}
	expectedVal := "val1"
	actualVal := firstRecord.Value
	if expectedVal != actualVal {
		t.Error("value incorrect")
	}
}
