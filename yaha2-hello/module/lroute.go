package module

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/caddyserver/caddy/v2"
	"github.com/caddyserver/caddy/v2/caddyconfig/caddyfile"
	"github.com/caddyserver/caddy/v2/modules/caddyhttp"
)

func init() {
	caddy.RegisterModule(LRoute{})
}

var (
	_ caddy.Module                = (*LRoute)(nil)
	_ caddy.Provisioner           = (*LRoute)(nil)
	_ caddyhttp.MiddlewareHandler = (*LRoute)(nil)
	_ caddyfile.Unmarshaler       = (*LRoute)(nil)
)

type LRoute struct {
	Protocol  string          `json:"protocol,omitempty"`
	UriRegexp []string        `json:"uri_regexp,omitempty"`
	Raw       json.RawMessage `json:"raw,omitempty" caddy:"namespace=http.handlers inline_key=module"`

	moduleInfo caddy.Module
}

func (h LRoute) CaddyModule() caddy.ModuleInfo {
	return caddy.ModuleInfo{
		ID: "http.handlers.lroute",
		New: func() caddy.Module {
			return new(LRoute)
		},
	}
}

func (h *LRoute) Provision(ctx caddy.Context) error {
	return nil
}

// Validate validates that the module has a usable config.
func (h LRoute) Validate() error {
	if string(h.Raw) == "" {
		return errors.New("the name is must!!!")
	}
	return nil
}
func (h LRoute) ServeHTTP(writer http.ResponseWriter, request *http.Request, handler caddyhttp.Handler) error {
	return handler.ServeHTTP(writer, request)
}
