# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/), and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [0.4.0] - 2021-09-01

### Added

- Each `VarintCodec` now supports reading from ByteReaders and writing to ByteWriters

### Changed

- The regular byte slice methods have been renamed to `EncodeToBytes` and `DecodeBytes`, to match `encoding/hex` and similar packages

## [0.3.0] - 2021-08-30

### Changed

- the module has been restructured - instead of submodules /signed and /unsigned, there is now just one top-level module `varint` that exports two variables `Signed` and `Unsigned` that both implement the `VarintCodec` interface.

## [0.2.0] - 2021-08-30

### Added

- This changlog!
- An index package `bigvarint` at `github.com/joeltg/go-big-varint`

[unreleased]: https://github.com/joeltg/go-big-varint/compare/v0.4.0...HEAD
[0.4.0]: https://github.com/joeltg/go-big-varint/compare/v0.4.0
[0.3.0]: https://github.com/joeltg/go-big-varint/releases/tag/v0.3.0
[0.2.0]: https://github.com/joeltg/go-big-varint/releases/tag/v0.2.0
