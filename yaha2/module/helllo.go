package module

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/caddyserver/caddy/v2"
	"github.com/caddyserver/caddy/v2/caddyconfig/caddyfile"
	"github.com/caddyserver/caddy/v2/caddyconfig/httpcaddyfile"
	"github.com/caddyserver/caddy/v2/modules/caddyhttp"
)

func init() {
	caddy.RegisterModule(helloWorld{})
	httpcaddyfile.RegisterHandlerDirective("helloworld", parseCaddyfile)
}

var (
	_ caddy.Module                = (*helloWorld)(nil)
	_ caddy.Provisioner           = (*helloWorld)(nil)
	_ caddyhttp.MiddlewareHandler = (*helloWorld)(nil)
	_ caddyfile.Unmarshaler       = (*helloWorld)(nil)
)

type helloWorld struct {
	Name string `json:"name,omitempty"`
}

func (h helloWorld) CaddyModule() caddy.ModuleInfo {
	return caddy.ModuleInfo{
		ID: "http.handlers.helloworld",
		New: func() caddy.Module {
			return new(helloWorld)
		},
	}
}

func (h *helloWorld) Provision(context caddy.Context) error {
	if h.Name == "" {
		h.Name = "default name"
	}
	return nil
}

// Validate validates that the module has a usable config.
func (h helloWorld) Validate() error {
	if h.Name == "" {
		return errors.New("the name is must!!!")
	}
	return nil
}
func (h helloWorld) ServeHTTP(writer http.ResponseWriter, request *http.Request, handler caddyhttp.Handler) error {
	fmt.Fprintf(writer, "Hello, %s!", h.Name)
	return handler.ServeHTTP(writer, request)
}

func parseCaddyfile(h httpcaddyfile.Helper) (caddyhttp.MiddlewareHandler, error) {
	handler := new(helloWorld)
	handler.UnmarshalCaddyfile(h.Dispenser)
	return handler, nil
}

func (h *helloWorld) UnmarshalCaddyfile(d *caddyfile.Dispenser) error {
	for d.Next() {
		if !d.NextArg() {
			return d.ArgErr()
		}
		h.Name = d.Val()
		//if !d.AllArgs(&h.Name) {
		//	return d.ArgErr()
		//}
	}
	return nil
}
