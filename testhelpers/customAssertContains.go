// CustomAssertContains applies to tests so it only outputs critical information when a test fails.
package testhelpers

import (
	"strings"
	"testing"
)

func CustomAssertContains(t *testing.T, s, contains, msg string, args ...interface{}) {
	if !strings.Contains(s, contains) {
		t.Logf(msg, args...)
		t.Fail()
	}
}
