# CipherLex

## Overview
CipherLex is a Go application that identifies occurrences of dictionary words within input text strings. It supports both original and scrambled word forms.

## System Requirements
Go (1.x or higher)

## Getting Started

- Download & build
```bash
git clone https://github.com/1x-eng/cipherlex
cd cipherlex
go build -o cipherlex ./cmd
```

- Run
```bash
./cipherlex path/to/dictionary.txt path/to/input.txt
```


### Dictionary File Format
- One word per line.
- No duplicates.
- Words must be 2 to 20 characters long.
- Maximum of 100 words.


### Input File Format
- One line of text per line.
- Maximum of 100 lines.
- Each line must be 2 to 500 characters long.


## Output
Outputs the number of unique dictionary words (in original or scrambled form) found in each line of the input file, formatted as:

```
Case #1: [count]
Case #2: [count]
```

### Configuration
Configurable parameters (via environment variables):

- MIN_WORD_LENGTH: Minimum length of dictionary words.
- MAX_WORD_LENGTH: Maximum length of dictionary words.
- MAX_DICTIONARY_SIZE: Maximum number of words in the dictionary.
- MIN_LINE_LENGTH: Minimum length of input text lines.
- MAX_LINE_LENGTH: Maximum length of input text lines.
- MAX_LINE_COUNT: Maximum number of lines in the input file.
- CHUNK_SIZE: Size of chunks for processing input text.
