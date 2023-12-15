package pkglogger

import "testing"

func Test(t *testing.T) {
	NewPkgLogger("pkglogger", "Test").Identifier("id").Detail("One", "Two").Detail("Three", "Four").Info().Msg("My message")
}
