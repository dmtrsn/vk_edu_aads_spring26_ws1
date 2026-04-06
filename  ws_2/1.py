def longest_substr(s):
    chars = set()
    l = 0
    max_length = 0

    for r in range(len(s)):
        while s[r] in chars:
            chars.remove(s[l])
            l += 1
        
        chars.add(s[r])
        max_length = max(max_length, r - l + 1)

    return max_length


s = "ababcdefabb"
print(longest_substr(s))

s = "abc"
print(longest_substr(s))

s = "aa"
print(longest_substr(s))

s = "a"
print(longest_substr(s))