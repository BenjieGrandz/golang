## Solution

- make a map to store the words using wordCOuntMap := make(map[string]int)
- create a words variable to store result of strconv.Fields(s)
- loop through the words increasing count using wordCountMap[word]++
- loop through the wordCountMap and print the key(word) and value(count)
- return the map now



## Encounter

- encountered strconv.Fields() splits  by whitespaces and tabs