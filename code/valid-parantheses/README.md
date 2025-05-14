## Solution

- create stack to push and pop the brackets
- create a bracketMap to map closing brackets to opening brackets
- loop through the string
- use switch case to switch through the opening brackets and closing brackets.
- pop opening brackets into the stack, push out the brackets 
- return only if the stack in empty := retrun len(stack) == 0 // true