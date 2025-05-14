# SOLUTION





## Discoveries

- os.Args creates a slice of string
- tasks := append(tasks, Task{Description: desc})
- The ... (ellipsis) unpacks the slice so it appends each element individually.

🧩 tasks[:index-1]
Gives you all elements before the one to delete.

Since index = 3, index-1 = 2, so:

tasks[:2] // gets tasks[0] and tasks[1]
