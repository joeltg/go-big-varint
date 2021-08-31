# go-big-varint

[![standard-readme compliant](https://img.shields.io/badge/readme%20style-standard-brightgreen.svg)](https://github.com/RichardLitt/standard-readme) [![license](https://img.shields.io/github/license/joeltg/go-big-varint)](https://opensource.org/licenses/MIT) ![latest tag](https://img.shields.io/github/v/tag/joeltg/go-big-varint)

Encode and decode arbitrarily large signed and unsigned `*big.Int` values.

## Table of Contents

- [Install](#install)
- [Usage](#usage)
- [Testing](#testing)
- [Credits](#credits)
- [Contributing](#contributing)
- [License](#license)

## Install

```
go get github.com/joeltg/go-big-varint
```

## Usage

https://pkg.go.dev/github.com/joeltg/go-big-varint

The module exports two variables `Signed` and `Unsigned` that both implement the `VarintCodec` interface:

```go
type VarintCodec interface {
	EncodingLength(i *big.Int) int
	Encode(data []byte, i *big.Int) int
	Decode(data []byte) (*big.Int, int)
}

var Signed VarintCodec
var Unsigned VarintCodec
```

## Testing

```
go test
```

## Credits

The encoding/decoding functions were translated from [chrisdickinson/varint](https://github.com/chrisdickinson/varint).

## Contributing

PRs welcome!

## License

MIT © 2021 Joel Gustafson
