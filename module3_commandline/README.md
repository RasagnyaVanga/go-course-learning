1. Accepts a comma-separated string of numbers using the -numbers flag.
2. Converts the string into an array of strings using strings.Split.
3. Converts the array of strings into an array of integers using strconv.Atoi.
4. Skips invalid inputs and prints a warning message.
5. Panics if no valid numbers are entered, caught with a defer + recover block.
6. Displays the sum of all valid numbers.