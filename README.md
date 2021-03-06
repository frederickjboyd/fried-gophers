# fried-gophers

_A command line tool for deep-frying images, written in Go._

## Installation

Download and install `fried-gophers` using the built-in command for Go:

```
go get https://github.com/frederickjboyd/fried-gophers
```

## Example

Using default values:

| Before                                                         | After                                                         |
| -------------------------------------------------------------- | ------------------------------------------------------------- |
| ![Imgur](https://i.imgur.com/scU2Dhc.jpg "Before deep frying") | ![Imgur](https://i.imgur.com/nWq2n4H.jpg "After deep frying") |

## Usage

### Flags

To view all flags and their descriptions, run:

`fried-gophers --help` or `fried-gophers -h`

### Compatibility

`fried-gophers` should work on Windows, OS X, and Linux. Preliminary tests have
been done on Windows and Linux, but more rigorous testing will be done in the
near future.

#### Supported Image Formats:

- JPEG

## Todo

- Support deep frying multiple images in the same directory simultaneously
- Support for additional image formats (e.g. PNG)
- Verify Go 11 support
