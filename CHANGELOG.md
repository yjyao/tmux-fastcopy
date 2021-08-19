# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## Unreleased
### Removed
- Remove `-log` flag in favor of controlling the log file with an environment
  variable.

### Fixed
- Fix crashes of the wrapped binary not getting logged.

## 0.2.0 - 2021-08-15
### Changed
- Download pre-built binary from GitHub if Go isn't available.

## 0.1.0 - 2021-08-14

- Initial release.