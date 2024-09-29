package module

import (
	"log"

	"github.com/caddyserver/caddy/v2/caddyconfig/caddyfile"
	"github.com/caddyserver/caddy/v2/caddyconfig/httpcaddyfile"
	"github.com/caddyserver/caddy/v2/modules/caddyhttp"
)

func init() {
	httpcaddyfile.RegisterHandlerDirective("helloworld", parseCaddyfile)
}

// -- 解析 caddy file
func parseCaddyfile(h httpcaddyfile.Helper) (caddyhttp.MiddlewareHandler, error) {
	handler := new(LRoute)
	handler.UnmarshalCaddyfile(h.Dispenser)
	return handler, nil
}

func (h *LRoute) UnmarshalCaddyfile(d *caddyfile.Dispenser) error {
	for d.Next() {
		if !d.NextArg() {
			return d.ArgErr()
		}
		log.Printf("val: %s", d.Val())
	}
	return nil
}
