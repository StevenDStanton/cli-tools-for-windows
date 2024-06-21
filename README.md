# CLI Tools for Windows

I named this repo Linux Tools for Windows because initially I planned to just reverse engineer various linux commands missing on Windows as I learned Go. However, I have found a few other projects and tools I wanted to add and have renamed it CLI Tools for Windows.

Bringing beloved Linux utilities to the Windows environment, rewritten from scratch, along with some unique applications of my own, such as the Text to Speech program.

## Planned Tools

- [x] wc
- [x] tts
- [x] crypto
- [ ] compress
- [ ] cut
- [ ] sort
- [ ] grep
- [ ] uniq
- [ ] diff
- [ ] cat
- [ ] tail
- [ ] traceroute
- [ ] jq

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

## Testing

[![Coverage Status](https://coveralls.io/repos/github/StevenDStanton/cli-tools-for-windows/badge.svg?branch=master)](https://coveralls.io/github/StevenDStanton/cli-tools-for-windows?branch=master)

I have written all tests to use Fuzz. However, this is not set up in the pipeline due to how expensive those tests are to run.

[Fuzz Testing](https://go.dev/doc/security/fuzz/)

### Normal Test Running

```bash
go test -v
```

### Fuzz Testing

```bash
go test --fuzz=Fuzz -fuzztime=1m
```

## Tools Included

- wc: Supports counting bytes, characters, words, lines, and the maximum line length.
- tts: Converts Markdown files to speech using the OpenAI API.

## Installation

[Download the latest cli-tools-for-windows-windows-amd64.zip](https://github.com/StevenDStanton/cli-tools-for-windows/releases)

Unzip and add the executables to your Windows `PATH`.

## Software

You can use --help with any of the below packages

### wc

```bash
wc [OPTION]... [FILE]...
wc [OPTION]... --files0-from=F
```

### crypto

Provide Crypto/Fiat currency abbreviations.

```bash
crypto BTC/USD
crypto BTC/USD DOGE/USD SHIB/USD
```

### Compress

Note: This is not a traditional Linux tool; it's a command line Huffman Encoder/Decoder I developed for learning purposes.

### tts

A simple CLI tool for converting text files to speech using the OpenAI API. The tool reads a text file (including Markdown), sends its content to the OpenAI API for text-to-speech conversion, and saves the generated audio file.

```bash
tts -f filename.md -o filename.mp3
```

## License

This project is released under the MIT License. For more details, see the [LICENSE](LICENSE) file.

## Resources

[POSIX.1-2017](https://pubs.opengroup.org/onlinepubs/9699919799.2018edition/)
