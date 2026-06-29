# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [1.0.0] - 2026-06-28
### Added
- Python `ApiKeyGen` class with raw key and SHA-256 hash generation.
- Node.js `ApiKeyGen` class utilizing native `crypto` module.
- Go `ApiKeyGen` struct utilizing `crypto/rand` and `crypto/sha256`.
- Comprehensive test suites for all three languages.
- CLI argument support for customizable key prefixes.
- 
