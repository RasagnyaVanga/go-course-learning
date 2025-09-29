# Word Count

## Introduction
Imagine you are teaching English as a foreign language to high school students.  
To make lessons engaging, you’ve decided to use TV shows as the basis of your curriculum.  

To choose the simplest shows first and gradually increase difficulty, you need to analyze subtitles:  
- Which words are used?  
- How often are they repeated?  

This project provides a tool to count word frequencies in subtitles.


## Instructions
Your task is to count how many times each word occurs in a subtitle of a drama.  

### Rules
- Subtitles contain only **ASCII characters**.  
- **Contractions** (e.g. *they're*, *it's*) are treated as single words.  
- Words are separated by **punctuation** (like `:`, `!`, `?`) or **whitespace** (like tabs `\t`, newlines `\n`, or spaces `" "`).  
- The only punctuation that does **not** separate words is the **apostrophe** (`'`) in contractions.  
- **Numbers** are considered words.  
  - Example: `"It costs 100 dollars."` → `100` is its own word.  
- **Case insensitive**: `"You"`, `"you"`, and `"YOU"` are all the same word.  
- The **order** of results does not matter.  

## Example
Sentence: "You come back, you hear me? DO YOU HEAR ME?"
## output 
- you: 3
- come: 1
- back: 1
- hear: 2
- me: 2
- do: 1

