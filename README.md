# Coolsnippets

I built this project to solve my problems using multiple apps(sublime text, neovim, ...) to manipulate data with common tasks.

CoolSnippets is a command-line tool written in Go that provides various text processing utilities, such as removing HTML tags, unescaping HTML entities, extracting URLs, and more. 

It is designed to be fast and efficient.

Neovim plugin will be released soon.

## Installation

### 1. Download Pre-built Binaries

You can install the `coolsnippets` tool by downloading the pre-built binaries from the releases.

### 2. Install From Source

Go must be installed on your system. Then, you can clone the repository and build the project:

```bash
git clone github.com/matusstafura/coolsnippets
cd coolsnippets
go build -o coolsnippets
```

or via makefile:

```bash
git clone github.com/matusstafura/coolsnippets
cd coolsnippets
make build
```

or via go install:

```bash
git clone github.com/matusstafura/coolsnippets
cd coolsnippets
go install .
```

## Usage
 
 ```bash
coolsnippets -u <utility> -s <sourcetext> <rest of parameters>
 ```

```bash
# Basic usage example
coolsnippets -u remove-html -s "hello<br>world"

You can also pipe input

# from standard input
echo "hello<br/>world" | ./coolsnippets -u remove-html

# or multiple utilities in a pipeline
cat file.txt | ./coolsnippets -u strip-tags-newline | sort | uniq | ./coolsnippets -u unescape-html

# or directly from a file
./coolsnippets -u strip-tags-newline < file.txt
```

## Current Features

### Strip Tags

Strip HTML tags from the input string.

```bash
coolsnippets -u strip-tags -s '<p>Hello<br>World!</p>'
# Output: HelloWorld!
```

### Strip Tags Newline

Strip HTML tags and replace tags with newlines.

```bash
coolsnippets -u strip-tags-newline -s '<p>Hello<br>World!</p>'
# Output: Hello
#         World!
```

### Unescape HTML entities

Unescape HTML entities in the input string.

```bash
coolsnippets -u unescape-html -s 'Hello &amp; World &lt;3'
# Output: Hello & World <3
```

### Extract URLs

Extract URLs from the input string.

```bash
coolsnippets -u extract-urls -s 'Visit https://example.com and http://test.com'
# Output:
# https://example.com
# http://test.com
```

### Backlink generation

Generate backlinks from a list of URLs.

```bash
coolsnippets -u backlink -s <sourcetext> <keyword> <n> <url>
# n is the occurrence number of the keyword to be replaced with backlink
```

```bash
coolsnippets -u backlink -s "this is a product." "product" "1" "http://example.com"
# Output:
# this is a <a href="http://example.com">product</a>.
```

### Strip Attributes

Strip style and script tags from HTML content.

```bash
coolsnippets -u strip-attributes -s '<div style="color:red;">Hello <span style="font-weight:bold;">World!</span></div>'
# Output: <div>Hello <span>World!</span></div>
```

## Development

To set up a development environment, clone the repository and build the project:

```bash
git clone github.com/matusstafura/coolsnippets
cd coolsnippets
go build -o coolsnippets
```

### Benchmarks and Tests

You can run benchmarks and tests using the provided [Makefile](Makefile):

```bash
# run tests
make test

# run benchmarks
make bench

# run tests with coverage
make cover

# run bats test
make test-bats
```

## Contributing & Support

Contributions are welcome! Please feel free to submit a Pull Request or open an issue for any bugs or feature requests.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
