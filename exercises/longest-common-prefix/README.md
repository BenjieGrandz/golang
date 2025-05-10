## Longest Common Prefix

Write a function to find the longest common prefix string amongst an array of strings.

If there is no common prefix, return an empty string "".

 

Example 1:

Input: strs = ["flower","flow","flight"]
Output: "fl"
Example 2:

Input: strs = ["dog","racecar","car"]
Output: ""
Explanation: There is no common prefix among the input strings.


### Solution

- sort errors first 
- find the shortest prefix
- loop through each str in the strs array while comparing the values at that point. provided the char is not the shortest
- return strs[0][:shortest]