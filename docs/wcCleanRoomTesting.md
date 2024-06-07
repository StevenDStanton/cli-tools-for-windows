# WC Observations

All testing was done on Windows using WSL and Ubuntu

## help

`wc --help`

This provided the following scenarios to test

```bash
wc [OPTION] [FILE]
wc [OPTION] --files0-from=F
```

With the following notes:

- Print newline, word, and byte counts for each FILE, and a total line if more than one FILE is specified.
- A word is a non-zero-length sequence of characters delimited by white space.
- With no FILE, or when FILE is -, read standard input.
- Output order: newline, word, character, byte, maximum line length.

### Flags to test

- [x] --bytes
- [x] --chars
- [x] --lines
- [x] --words
- [x] --max-line-length
- [x] --version
- [x] --help
- [ ] --files0-from=F
- [x] -c
- [x] -m
- [x] -l
- [x] -w
- [x] -L
- [x] No arguments
- [x] Single File
- [x] Multiple Files

### Flag Combos to Test

- [x] --bytes --words
- [x] -cm
- [x] --nonExistentFlag
- [x] --bytes --nonExistentFlag
- [x] -cwx

### Edge behavior to test

- [x] --version with other flags
- [x] --help after other flags
- [x] --option flags after positional input
- [x] - in the middle of two files
- [x] nonExistent file with two good files
- [x] GoodFile - Good File
- [x] short flags, file, long flags mixed

## black box testing

- Test:

  - Command: `wc README.md`
  - Output: `50  415 2899 README.md`
  - Notes: Need to make sure we have left padding on all outputs except file name.

- Test:

  - Command: `wc -c READEME.md`
  - Output: `2899 README.md`
  - Notes: As Expected

- Test:

  - Command: `wc -c -m README.md`
  - Output: `2899 2899 README.md`
  - Notes: As expected as all characters in the README.md are 1 byte

- Test:
  - Command: `wc -cm README.md`
  - Output: `2899 2899 README.md`
  - Notes: As expected same results as `-c -m`
- Test:
  - Command: `wc -c -m README.md --version --help`
  - Output: Only prints version statement
  - Notes: Version overrides all other commands
- Test:
  - Command: `wc`
  - Output: Awaited Standard Input and entered. `Hello World` Output of `0       2      12`
  - Notes: Output is on the same line as standard input which I view as a bug. Will fix in my implementation
- Test:
  - Command:`wc -c -m README.md --words`
  - Output `415 2899 2899 README.md`
  - Notes: Outputs characters, bytes, and words as expected
- Test:
  - Command `wc README.md --bytes --chars --lines --words --max-line-length`
  - Output:`50  415 2899 2899  466 README.md`
  - Notes: Expected output per documentation
- Test:
  - Command: `wc README.md -cmlwL`
  - Output:`50  415 2899 2899  466 README.md`
  - Notes: Expected output per documentation
- Test:
  - Command: `wc README.md`
  - Output: `50  415 2899 README.md`
  - Notes: As Expected
- Test:
  - Command: `wc LICENSE README.md`
  - Output:

```bash
  21  169 1092 LICENSE
  50  415 2899 README.md
  71  584 3991 total

```

- Notes: Expected Per Documentation

- Test:
  - Command: `wc LICENSE README.md nonExistentFilePAth`
  - Output

```bash
  21  169 1092 LICENSE
  50  415 2899 README.md
wc: nonExistentFilePAth: No such file or directory
  71  584 3991 total
```

- Notes: Tested multiple variations and will display that error anywhere in the list it can't open the file

- Test:
  - Command: `wc LICENSE README.md -wcx`
  - Output:

````bash
wc: invalid option -- 'x'
Try 'wc --help' for more information.
```

- Test:
  - Command: `wc LICENSE README.md --bytes --unknownTest`
  - Output:

```bash
wc: unrecognized option '--unknownTest'
Try 'wc --help' for more information.
````

- Test:
  - Command: `wc LICENSE - README.md`
  - Output:

```bash
     21     169    1092 LICENSE
Hello World      0       2      11 -
     50     415    2899 README.md
     71     586    4002 total
```

- Notes: I am not a fan of how it includes the text in the lines as I feel it makes it more cluttered and I will fix in my version

- Test:
  - Command: `wc LICENSE - README.md -c --help`
  - Output: Displays Help Files
  - Notes: As expected

All other tests ran were variations of these tests.
