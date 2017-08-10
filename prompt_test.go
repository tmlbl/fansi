package fansi

import "testing"
import "strings"

func askYesNoTest(t *testing.T, input string, expect bool) {
	p := NewTestPrompt(input)
	ans := p.AskYesNo("")
	if ans != expect {
		t.Errorf("%s should have produced %v, but didn't",
			strings.TrimSpace(input), expect)
	}
}

func TestAskYesNo(t *testing.T) {
	// Responds correctly to yes
	askYesNoTest(t, "Y\n", true)
	askYesNoTest(t, "y\n", true)
	askYesNoTest(t, "yYy\n", true)
	askYesNoTest(t, "yes\n", true)
	askYesNoTest(t, "YES\n", true)
	// Responds correctly to no
	askYesNoTest(t, "N\n", false)
	askYesNoTest(t, "n\n", false)
	askYesNoTest(t, "nnn\n", false)
	askYesNoTest(t, "no\n", false)
	askYesNoTest(t, "NO\n", false)
	askYesNoTest(t, "NONONONONONO\n", false)
}

func TestGetInput(t *testing.T) {
	p := NewTestPrompt("hello\n")
	str, err := p.GetInput("")
	if err != nil {
		t.Error(err)
	}
	if str != "hello" {
		t.Errorf("Expected string hello but got %s", str)
	}
}
