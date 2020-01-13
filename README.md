# Usage

## 1) Import package

`import "github.com/rpagliuca/gojicorsmiddleware"`

## 2) Enable middleware on Goji mux

`// mux := goji.NewMux()`

`mux.Use(gojicorsmiddleware.CorsMiddlewareHandler)`

# Example

Check `main_test.go` for full example.
