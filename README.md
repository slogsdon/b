# b [![Build Status](https://travis-ci.org/slogsdon/b.svg)](https://travis-ci.org/slogsdon/b) [![GoDoc](https://godoc.org/github.com/slogsdon/b?status.png)](http://godoc.org/github.com/slogsdon/b) [![Coverage Status](https://coveralls.io/repos/slogsdon/b/badge.png?branch=master)](https://coveralls.io/r/slogsdon/b?branch=master)

A static-ish blog application. Can be run as a standalone application/server or be used to locally manage and deploy posts to a remote server.

## Building

If you wish to build b you'll need Go version 1.2+ installed.

Please check your installation with:

```
go version
```

## Testing

If you wish to run the tests, either run them using `go test ./...` in the project directory, or using the included `Guardfile`, run [`guard`](https://github.com/guard/guard) in the project directory to have the tests run automatically when `*.go` files are modified.

## TODO

- Create new posts
- Edit posts
- Admin interface
- Generation of static files
- Automated revisioning with git
- Some form of authentication (?)
- Deployment
    + S3
    + SSH
- Manage other content
    + Images
    + Themes
- Modules/Plugins

## License

Licensed under the MIT License.

See the [LICENSE](https://github.com/slogsdon/b/blob/master/LICENSE) file for details.