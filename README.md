# b [![wercker status](https://app.wercker.com/status/eaa45a2252df7c1535fddd9cced59e91/s/ "wercker status")](https://app.wercker.com/project/bykey/eaa45a2252df7c1535fddd9cced59e91)

A static-ish blog application. Can be run as a standalone application/server or be used to manage and deploy posts to a remote server.

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
- Some form of authentication (?)
- Generation of static files
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