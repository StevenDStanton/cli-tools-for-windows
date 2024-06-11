# Linux Tools for Windows (LTfW)

Bringing beloved Linux utilities to the Windows environment, rewritten from scratch.

## Introduction

This project aims to provide Windows users with native versions of popular Linux command-line tools, starting with wc (word count). These tools have been re-engineered based only on their behavior and POSIX standards, without using any original source code.

## Reverse Engineering Technique

Why reverse engineer open source code? The goals are actually three fold. First, I wanted some of my favorite Linux tools available on Windows. Second, I wanted to do some high level system work for learning. Third, I am leaning Go and these seemed like a perfect way to master the language.

### Methodology

Step 1: RTFM
Step 2: Design tests
Step 3: Implement tests
Step 4: Record tests
Step 5: Write code!

I employed the black-box reverse engineering and clean room design approach in this project. This methodology is crucial in ensuring that the new software is developed independently of the original GPL-licensed code, avoiding any legal implications related to copyright infringement.

- Black-Box Reverse Engineering: In this phase, I interacted with the software purely from an external perspective. I analyzed how the software behaves in response to various inputs, what outputs it produces, and how it handles different operating conditions. This analysis was conducted without any access to the actual source code, relying solely on observable behaviors.

- Clean Room Design: To ensure compliance with copyright laws and to maintain the integrity of the reverse engineering process, I adopted a clean room strategy. This involved documenting all observations and behaviors noted during the black-box analysis phase. Afterward, using these documents as a blueprint, I developed the new software from scratch. This step was critical to ensure that no code or internal architecture from the original software was replicated.

### Implementation

The new implementation is crafted to mimic the functionality observed during the analysis phase, yet it is fundamentally built from unique code written by myself. This approach not only adheres to legal standards but also pushes the boundaries of my coding skills, allowing me to explore innovative solutions and optimizations not present in the original software.

## Features

- wc: Supports counting bytes, characters, words, lines, and the maximum line length.
- Cross-Platform: Designed to work seamlessly on all Windows versions.
- POSIX Compliance: Adheres to the POSIX.1-2017 standard for maximum compatibility and reliability.

## Installation

Coming Soon

## Software

### wc

```bash
wc [OPTION]... [FILE]...
wc [OPTION]... --files0-from=F
```

### Compress

Ok so this is not a linux tool, I just wanted to take my hand at building a command line Huffman Encoder/Decoder

## License

This project is released under the MIT License. For more details, see the [LICENSE](LICENSE) file.

## Resources

[POSIX.1-2017](https://pubs.opengroup.org/onlinepubs/9699919799.2018edition/)
