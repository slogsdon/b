# b
[![GoDoc](https://godoc.org/github.com/slogsdon/b?status.png)](http://godoc.org/github.com/slogsdon/b)
[![Build Status](https://travis-ci.org/slogsdon/b.svg)](https://travis-ci.org/slogsdon/b)
[![Coverage Status](https://coveralls.io/repos/slogsdon/b/badge.png?branch=master)](https://coveralls.io/r/slogsdon/b?branch=master)

A static-ish blog application. Can be run as a standalone application/server or be used to locally manage and deploy posts to a remote server.

## Building

```
$ cd $GOPATH/src/github.com/slogsdon/b
$ go build -o b-serve cmd/b-serve/main.go
```

> If you wish to build b you'll need Go version 1.2+ installed.
>
> Please check your installation with:
>
> ```
> $ go version
> ```

## Testing

```
$ go test ./...
```

## TODO

- Create new posts - 50% (api endpoint complete)
- Edit posts - 75% (api endpoints [show, update] and admin endpoint [edit] complete)
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
