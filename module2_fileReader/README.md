## Overview
- Takes filepath from user.
- Reads the entire file into memory using `os.ReadFile()`.
- If the file does not exist or the path is incorrect, it prints an appropriate error message.
- If reading the file succeeds, the file content is printed as a string.
- Handles "file not found" errors using `errors.Is(err, os.ErrNotExist)`.

## Alternative Approaches
Instead of `os.ReadFile`, we could use:
- `os.Open()` to open the file.
- `file.Read()` to read bytes, or `bufio.NewScanner()` to read line by line.
- `io.ReadAll()` to read the whole file.
- Manually close the file using `file.Close()`.
