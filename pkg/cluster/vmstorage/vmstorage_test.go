package vmstorage

import (
	"net/url"
	"testing"
)

func Test_storage_Health(t *testing.T) {
	u, _ := url.Parse("http://localhost" + ":8482")
	u = u.JoinPath("/metrics")
	t.Log(u.String())
}
