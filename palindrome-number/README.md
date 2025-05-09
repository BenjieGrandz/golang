Given an integer x, return true if x is a palindrome, and false otherwise.

 

Example 1:

Input: x = 121
Output: true
Explanation: 121 reads as 121 from left to right and from right to left.
Example 2:

Input: x = -121
Output: false
Explanation: From left to right, it reads -121. From right to left, it becomes 121-. Therefore it is not a palindrome.
Example 3:

Input: x = 10
Output: false
Explanation: Reads 01 from right to left. Therefore it is not a palindrome.


## Solution

- Start off by handling errors (negative/tens and hundreds)
- Then declare the rev in to hold the reversed number
- use a for loop so as to cut the number by half
- inside do a reversal using (rev := (rev*10) + x%d) (n /= 10 )
- return x == rev for even numbers and x == rev/10 for odd numbers