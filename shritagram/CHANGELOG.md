# Changelog

## [1.3.1] - 2022-02-21
### Added
- Added webhook endpoint and apikey binding. When registered a webhook, a new apikey is returned for futher usage.
- Added priority on each request, so client can manually specify which priority should be used.

## [1.2.0] - 2022-01-21
### Added
- Added topsearch implemenation on server side.
- Added lazy initialization module.

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
