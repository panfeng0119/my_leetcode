![242. Valid Anagram](https://leetcode.com/problems/valid-anagram/)

Given two strings s and t , write a function to determine if t is an anagram of s.

```text
Example 1:

Input: s = "anagram", t = "nagaram"
Output: true
Example 2:

Input: s = "rat", t = "car"
Output: false
```

思考

You may assume the string contains only lowercase alphabets.

What if the inputs contain unicode characters? How would you adapt your solution to such case?

参考大神的代码，开始思路是对字符串排序，而大神直接生成长度为26的map

每次映射k，成功则v++

第二个字符串，映射成功则v--
