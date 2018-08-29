package goyz

import (
	"testing"
)

func TestQueryBuild(t *testing.T) {
	client := NewClient("testToken")
	queryString := client.buildQuery(map[string]interface{}{"key1": "value1", "key2": 1});
	if queryString != "key1=value1&key2=1" {
		t.Errorf("expected < key1=value1&key2=1 >, got < " + queryString + " >")
	}
}