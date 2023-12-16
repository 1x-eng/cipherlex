# CipherLex

## Overview
CipherLex is a Go application that identifies occurrences of dictionary words within input text strings. It supports both original and scrambled word forms.

## System Requirements
Go (1.x or higher)

## How does it work?

```mermaid
graph TD
    A[Start] --> B[Load Dictionary]
    B --> C[Process Dictionary Words]
    C --> D[Load Input File]
    D --> E[Split Input into Chunks]
    E --> |For each chunk| F[Process Chunks in Parallel]
    subgraph Parallel Processing
        F --> G[Chunk 1]
        F --> H[Chunk 2]
        F --> I[Chunk n]
    end
    G --> J[Merge Results]
    H --> J
    I --> J
    J --> K[Count Unique Matches]
    K --> L[Output Results]
    L --> M[End]
``````

- **Start**: The beginning of the program.
- **Load Dictionary**: Reads the dictionary file.
- **Process Dictionary Words**: Applies constraints and processes dictionary words.
- **Load Input File**: Reads the input file (line by line, this is serial atm, we could leverage concurrency here as well. Its my todo.)
- **Split Input into Chunks**: Divides the input text into chunks for parallel processing.
- **Process Chunks in Parallel**: Concurrently processes each chunk to find matches.
- **Merge Results**: Combines results from all chunks (more akin of 'reduce' step of mapR)
- **Count Unique Matches**: Counts the unique dictionary words found.
- **Output Results**: Formats and outputs the results per line.
- **End**: The end of the program.

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
