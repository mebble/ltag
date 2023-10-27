# ltag

[![CI Tests](https://github.com/mebble/ltag/actions/workflows/test.yml/badge.svg)](https://github.com/mebble/ltag/actions/workflows/test.yml)
[![CI Release](https://github.com/mebble/ltag/actions/workflows/release.yml/badge.svg)](https://github.com/mebble/ltag/actions/workflows/release.yml)
[![Downloads on latest release](https://img.shields.io/github/downloads-pre/mebble/ltag/latest/total)](https://github.com/mebble/ltag/releases/latest)

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

For more examples, check out the test data at `./test/testdata/`. For more usage information, run `ltag --help`.

## Features

- Tag the lines to the values of headings, sub-headings and inline tags
- Tags are slugified
- Customise the string pattern used to identify and to format the tags
- Trim off the tags from lines that have already been ltagged

## Development

_Scaffolded with `go mod init github.com/mebble/ltag`_

### Testing

```sh
make test
```

### Benchmarking

```sh
make bench
```
