package commands

import "testing"

func TestShellSplit(t *testing.T) {
	stringArrayEquals := func(a []string, b []string) bool {
		if len(a) != len(b) {
			return false
		}
		for i, v := range a {
			if v != b[i] {
				return false
			}
		}
		return true
	}

	parseAndCompare := func(in string, out []string) {
		if !stringArrayEquals(ShellSplit(in), out) {
			t.Fatalf("expected %+v, got %+v for '%s'", out, ShellSplit(in), in)
		}
	}

	parseAndCompare("echo 'hello world'", []string{"echo", "hello world"})
	parseAndCompare(`echo "hello world"`, []string{"echo", "hello world"})
	parseAndCompare(`"heeeeeeeeeee"     "hello world"`, []string{"heeeeeeeeeee", "hello world"})
}
