# ltag

> A CLI tool to tag your lines based on their headings

## Basic Usage

Imagine you have a file `info.txt`:

```
# animals
elephant
# felines
cats
tigers
# canines
dogs
wolves
```

Pipe its contents to ltag:

```sh
cat info.txt | ltag
```

The output:

```
elephant #animals
cats #animals #felines
tigers #animals #felines
dogs #animals #canines
wolves #animals #canines
```

For more usage information, run `ltag --help`.

## Features

- Tag the lines to the values of headings, sub-headings and inline tags
- Tags are slugified
- Customise the string pattern used to identify and to format the tags
- Trim off the tags from lines that have already been ltagged

## Development

_Scaffolded with `go mod init github.com/mebble/ltag`_

### Testing

```sh
go test ./src/...
```

### Benchmarking

```sh
source ./benchmarks/run.sh
```
