# Usage

import "github.com/rpagliuca/gojicorsmiddleware"

mux := goji.NewMux()

mux.Use(gojicorsmiddleware.CorsMiddlewareHandler)
