package palindrome

import "fmt"

func FindLongest(s string) string {
    return FindLongestQuadratic(s)
}

func FindLongestQuadratic(s string) string {
    if len(s) == 0 {
        return ""
    }

    // Explode s to insert padding between each char
    // e.g. "aba" -> ".a.b.a.", "anna" -> ".a.n.n.a."
    s2 := make([]rune, len(s) * 2 + 1)
    i := 1
    for _, c := range s {
        s2[i] = c
        i = i + 2
    }

    // Find the longest palindrome around each pivot
    const EMPTY = 0
    longestPalindrome := ""
    for i := 1; i < len(s2)-1; i++ {
        // Initialize the pivot
        var palindrome string
        if s2[i] != EMPTY {
            palindrome = string(s2[i])
        } else {
            palindrome = ""
        }

        for l, r := i-1, i+1; l >= 0 && r <= len(s2)-1; l,r = l-1, r+1 {
          if s2[l] != s2[r] {
            break
          } else if s2[r] != EMPTY {
            palindrome = string(s2[l]) + palindrome + string(s2[r])
          }
        }

        if len(palindrome) > len(longestPalindrome) {
            longestPalindrome = palindrome
        }
    }

    return longestPalindrome
}

func FindLongestLinear(s string) string {
    s2 := make([]int, len(s) * 2 + 1)
    i := 1
    for range s {
        s2[i] = 1
        i += 2
    }

    fmt.Println(s2)
    return s // FIXME
}
