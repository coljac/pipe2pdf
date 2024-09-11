# pipe2pdf

A command-line tool for putting text in PDF form (for printing, etc). Pass it one or more text files, or pipe something from stdin, and it will produce a text-only PDF file.

Handy for quickly printing plain text content in a clean and predictable format. Specify font, font size, orientation, and paper size from the command line.

## Installation

You can:

- `go install github.com/coljac/pipe2pdf@latest`
- Download a binary
- Execute `curl ... | bash` on Linux

## Command-line usage

```
pipe2pdf --help
```

```
pipe2pdf -t "My Title" -o "output.pdf" "file1.txt" "file2.txt"
```


Or pipe something from stdin:

```
echo "Hello, world!" | pipe2pdf -o "output.pdf"
```

## Fonts

pipe2pdf comes with a few fonts built-in, but you can also specify a custom font.

```
pipe2pdf -f "customfont.ttf" -o "output.pdf" "file1.txt" "file2.txt"
```

pipe2pdf comes with a few fonts built-in, but you can also specify a custom font.

```
pipe2pdf -f "customfont.ttf" -o "output.pdf" "file1.txt" "file2.txt"
```
