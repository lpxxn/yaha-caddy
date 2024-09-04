package module

import (
	"fmt"
	"net/http"

	"github.com/caddyserver/caddy/v2"
	"github.com/caddyserver/caddy/v2/caddyconfig/caddyfile"
	"github.com/caddyserver/caddy/v2/caddyconfig/httpcaddyfile"
	"github.com/caddyserver/caddy/v2/modules/caddyhttp"
)

// MyHandler 是一个自定义的 HTTP handler 模块
type MyHandler struct {
	Message string `json:"message"`
}

// CaddyModule 返回 Caddy 模块的信息
func (MyHandler) CaddyModule() caddy.ModuleInfo {
	return caddy.ModuleInfo{
		ID:  "http.handlers.myhandler",
		New: func() caddy.Module { return new(MyHandler) },
	}
}

// ServeHTTP 实现了 caddyhttp.MiddlewareHandler 接口
func (h MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request, next caddyhttp.Handler) error {
	fmt.Fprintf(w, "Hello, %s!", h.Message)
	return nil
}

func parseMyHandlerCaddyfile(h httpcaddyfile.Helper) (caddyhttp.MiddlewareHandler, error) {
	handler := new(MyHandler)
	handler.UnmarshalCaddyfile(h.Dispenser)
	return handler, nil
}

// UnmarshalCaddyfile 实现了解析 Caddyfile 配置的方法
func (h *MyHandler) UnmarshalCaddyfile(d *caddyfile.Dispenser) error {
	for d.Next() {
		for d.NextBlock(0) {
			switch d.Val() {
			case "message":
				if !d.Args(&h.Message) {
					return d.ArgErr()
				}
			default:
				return d.Errf("unexpected token: %s", d.Val())
			}
		}
	}
	return nil
}

func init() {
	// 注册自定义模块
	caddy.RegisterModule(MyHandler{})
	httpcaddyfile.RegisterHandlerDirective("myhandler", parseMyHandlerCaddyfile)
}
