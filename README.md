# data-to-words

Humans are not very skilled at remembering or communicating 1's and 0's, but we are pretty skilled at communicating using words. This project is an attempt to convert data into a format that is easier to communicate without using a computer.

`data-to-words` is a command-line utility which converts an input stream into words. It was mainly a exercise to get to know Go better.

## Usage

```bash
$ data-to-words --help
$ cat /dev/random | data-to-words
$ data-to-words "bnlhbiBueWFuIG55YW4gbnlhbiBueWFu"
$ data-to-words -seed=2485 "bnlhbiBueWFuIG55YW4gbnlhbiBueWFu"
$ echo "the\nquick\nbrown\nfox\njumps\nover\nthe\nlazy\ndog" > dictionary.txt && data-to-words "hello"
```

## To do

- [ ] impl decoder
- [ ] prevent duplicate terms when sorted by alpha by checking the previous term
- [ ] form actual sentences?
