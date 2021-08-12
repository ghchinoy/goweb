# Go Webserver

Simple webserver demonstrating Go 1.16's embed feature - the webserver embeds all assets into a single binary.

Add your HTML etc into the `public` directory and compile, then serve via Cloud Run, etc.


As cool as [Justine Tunney](https://justine.lol/)'s [redbean](https://redbean.dev/)? Probably not.

## Deploy

Deploy to Google [Cloud Run with the single command](https://cloud.google.com/blog/products/serverless/build-and-deploy-an-app-to-cloud-run-with-a-single-command):

```
gcloud beta run deploy goweb --source . --allow-unauthenticated --region us-central1
```

## Notes

A brief note on how this works. The [embed](https://pkg.go.dev/embed) package provides static files embedded into the go binary at compile time. This fantastic feature was introduced in [Go 1.16](https://blog.golang.org/go1.16).

The folder `public` is designated as a static, embeded filesystem by what appears to be a comment [`//go:embed public/*`](https://github.com/ghchinoy/goweb/blob/486018f07221e1ef47984f5261301029beab82be/main.go#L10) but is actually a special directive (i.e. a comment that starts with `go:`) and then later [the `public` path prefix is removed](https://github.com/ghchinoy/goweb/blob/486018f07221e1ef47984f5261301029beab82be/main.go#L20) when the HTTP server provides the filesystem.