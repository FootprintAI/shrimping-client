# Changelog

## [1.1.5] - 2021-12-28
### Added
- Added client reconnect mechanism when the server is unavailable.
- Added empty parameter check, and return `http:400` when it is empty parameters.
- Added swagger resources for its corresponding grpc stubs.

## [1.1.0] - 2021-12-14
### Changed
- Use shortcode in `Post` Call where previously we allows it to be username.

### Improved
- Add async behavior to deliver results from callback streaming.

## [1.0.3] - 2021-11-26
### Added
- Added Topsearch with keywords or hashtags.
- Added crawled timestamp on each response object.
- Added proto definitions to `/proto` folder.

### Improved
- Enforce server side compression to reduce bandwidth overhead.
