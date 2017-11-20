package palindrome

func FindLongestQuadratic(s string) string {
    if len(s) == 0 {
        return ""
    }

    s2 := addPaddings(s)

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

        // For each pointer, find the longest palindrome pivoted around the current pointer
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
    if len(s) == 0 {
        return ""
    }

    s2 := addPaddings(s)
    p := make([]int, len(s2))

    // Center & Right of the rightermost palindrome so far
    c,rb := 0, 0
    l,r := 0, 0
    for i := 1; i < len(s2); i++ {
        // Set the walking indices l & r centered around i
        if i > rb {
            // If the current pointer extends beyond the right boundary of the current rightermost palindrome,
            // naively expand around the current pointer i
            p[i] = 0;
            l = i - 1;
            r = i + 1;
        } else {
            // If the current pointer lies within the right boundary of the current rightermost palindrome,
            // we can examine the already calculated values to avoid re-doing the same calculations

            // j is the mirrored position of i around c
            // (e.g. when c = 5, (i,j) = (6,4), (7,3), (8,2)...)
            j := c * 2 - i

            if p[j] < (rb-i-1) {
                // If the longest palindrome around j is within the bounds of the rightermost palindrome,
                // it has the same length as its mirror so we can simply copy the value
                p[i] = p[j]
                l = -1 // Skip the expand loop below
            } else {
                // If the mirror palindrome extends past the right boundary of the rightermost palindrome,
                // we can only assume up to the boundary, and then actually caculate the rest
                // But we stil save steps by staring past the right boundary
                p[i] = rb-i
                r = rb+1
                l = i*2-r
            }
        }

        // Count the length of the longest palindrome pivoted around the current pointer i
        for l >=0 && r < len(s2) && s2[l] == s2[r] {
            p[i] += 1
            l -= 1
            r += 1
        }

        // If the current palindrome extends beyond the right boundary, update the rightermost palindrome
        if (i + p[i]) > rb {
            c = i
            rb = i + p[i]
        }

    }

    // Find the pivot with the longest palindorme & its length
    longest, c := 0, 0
    for i := 1; i < len(s2); i++ {
        if  longest < p[i] {
            longest = p[i]
            c = i
        }
    }

    return removePaddings(s2[c-longest:c+longest+1])
}

/**
 * Explode s to insert padding between each char
 * e.g. "aba" -> ".a.b.a.", "anna" -> ".a.n.n.a."
 */
func addPaddings(s string) []rune {
    s2 := make([]rune, len(s) * 2 + 1)
    i := 1
    for _, c := range s {
        s2[i] = c
        i += 2
    }
    return s2
}

func removePaddings(s []rune) string {
    s2 := make([]rune, len(s) / 2)
    for i,j := 1, 0; i < len(s); i,j = i+2, j+1 {
        s2[j] = s[i]
    }
    return string(s2)
}
