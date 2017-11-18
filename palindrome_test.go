package palindrome

import "testing"

func TestFind(t *testing.T) {
    cases := []struct {
        in, expected string
    }{
        {"", ""},
        {"a", "a"},
        {"ab", "a"},
        {"abab", "aba"},
        {"ababcbab", "babcbab"},
        {"abaaaaaa", "aaaaaa"},
    }
    for _, c := range cases {
        actual := FindLongest(c.in)
        if actual != c.expected {
            t.Errorf("FindLongest(%q) == %q, expected %q", c.in, actual, c.expected)
        }
    }
}
