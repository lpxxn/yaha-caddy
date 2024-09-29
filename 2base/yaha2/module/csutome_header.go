package module

import (
	"net/http"

	"github.com/caddyserver/caddy/v2"
	"github.com/caddyserver/caddy/v2/caddyconfig/caddyfile"
	"github.com/caddyserver/caddy/v2/modules/caddyhttp"
)

func init() {
	caddy.RegisterModule(CustomHeader{})
}

// CustomHeader is an example handler that adds a custom header to the response
type CustomHeader struct {
	HeaderName  string `json:"header_name,omitempty"`
	HeaderValue string `json:"header_value,omitempty"`
}

// CaddyModule returns the Caddy module information.
func (CustomHeader) CaddyModule() caddy.ModuleInfo {
	return caddy.ModuleInfo{
		ID:  "http.handlers.custom_header",
		New: func() caddy.Module { return new(CustomHeader) },
	}
}

// ServeHTTP implements caddyhttp.MiddlewareHandler.
func (ch CustomHeader) ServeHTTP(w http.ResponseWriter, r *http.Request, next caddyhttp.Handler) error {
	w.Header().Set(ch.HeaderName, ch.HeaderValue)
	return next.ServeHTTP(w, r)
}

// UnmarshalCaddyfile implements caddyfile.Unmarshaler.
func (ch *CustomHeader) UnmarshalCaddyfile(d *caddyfile.Dispenser) error {
	for d.Next() {
		if !d.NextArg() {
			return d.ArgErr()
		}
		ch.HeaderName = d.Val()

		if !d.NextArg() {
			return d.ArgErr()
		}
		ch.HeaderValue = d.Val()
	}
	return nil
}

// Interface guards
var (
	_ caddy.Module                = (*CustomHeader)(nil)
	_ caddyhttp.MiddlewareHandler = (*CustomHeader)(nil)
	_ caddyfile.Unmarshaler       = (*CustomHeader)(nil)
)
