package palindrome

import "testing"

func TestFindLongestQuadratic(t *testing.T) {
    test(FindLongestQuadratic, t)
}

func TestFindLongestLinear(t *testing.T) {
    test(FindLongestLinear, t)
}

func test(f func(string) string, t *testing.T) {
    cases := []struct {
        in, expected string
    }{
        {"", ""},
        {"a", "a"},
        {"ab", "a"},
        {"abab", "aba"},
        {"abba", "abba"},
        {"ababcbab", "babcbab"},
        {"abaaaaaa", "aaaaaa"},
        {"abcdcbacac", "abcdcba"},
    }
    for _, c := range cases {
        actual := f(c.in)
        if actual != c.expected {
            t.Errorf("FindLongest(%q) == %q, expected %q", c.in, actual, c.expected)
        }
    }
}
