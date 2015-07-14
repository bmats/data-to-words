# data-to-words

Humans are not very skilled at remembering or communicating 1's and 0's, but we are pretty skilled at communicating using words. This project is an attempt to convert data into a format that is easier to communicate without using a computer.

`data-to-words` is a command-line utility which converts an input stream into words. It was mainly an exercise to get to know Go better.

## Usage

```
$ data-to-words --help
data-to-words 0.1.0
Usage: data-to-words [options] [data string]

Pipe some bytes to me or pass bytes to me as an argument.

Options:
  -delim=" ": Word delimiter
  -seed=0: Seed for choosing words
  -size=0: Dictionary size. 0 to select the maximum
```

```bash
$ cat /dev/random | data-to-words
$ data-to-words "bnlhbiBueWFuIG55YW4gbnlhbiBueWFu"
$ data-to-words -seed=2485 "bnlhbiBueWFuIG55YW4gbnlhbiBueWFu"
$ cat nyan.gif | data-to-words
$ echo "the\nquick\nbrown\nfox\njumps\nover\nthe\nlazy\ndog" > dictionary.txt && data-to-words "hello"
```

## To do

- Implement decoder
- Prevent duplicate terms when sorted by alpha by checking the previous term
- Form actual sentences?
