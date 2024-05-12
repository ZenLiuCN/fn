package info

import (
	"bytes"
	"fmt"
	"github.com/ZenLiuCN/fn"
	"go/ast"
	"go/types"
	"golang.org/x/tools/go/packages"
	"os"
	"testing"
)

var (
	typing = map[Id]Type{"net/http.Client": Type{
		Name: "Client",
		Path: "net/http",
		Fields: []Field{
			{
				Name:     "Transport",
				Tag:      "",
				Id:       "net/http.RoundTripper",
				Embedded: false,
				Offset:   0,
				Size:     16,
			},
			{
				Name:     "CheckRedirect",
				Tag:      "",
				Id:       "func(req *net/http.Request, via []*net/http.Request) error",
				Embedded: false,
				Offset:   16,
				Size:     8,
			},
			{
				Name:     "Jar",
				Tag:      "",
				Id:       "net/http.CookieJar",
				Embedded: false,
				Offset:   24,
				Size:     16,
			},
			{
				Name:     "Timeout",
				Tag:      "",
				Id:       "time.Duration",
				Embedded: false,
				Offset:   40,
				Size:     8,
			},
		},
	},
		"net/http.Cookie": Type{
			Name: "Cookie",
			Path: "net/http",
			Fields: []Field{
				{
					Name:     "Name",
					Tag:      "",
					Id:       "string",
					Embedded: false,
					Offset:   0,
					Size:     16,
				},
				{
					Name:     "Value",
					Tag:      "",
					Id:       "string",
					Embedded: false,
					Offset:   16,
					Size:     16,
				},
				{
					Name:     "Path",
					Tag:      "",
					Id:       "string",
					Embedded: false,
					Offset:   32,
					Size:     16,
				},
				{
					Name:     "Domain",
					Tag:      "",
					Id:       "string",
					Embedded: false,
					Offset:   48,
					Size:     16,
				},
				{
					Name:     "Expires",
					Tag:      "",
					Id:       "time.Time",
					Embedded: false,
					Offset:   64,
					Size:     24,
				},
				{
					Name:     "RawExpires",
					Tag:      "",
					Id:       "string",
					Embedded: false,
					Offset:   88,
					Size:     16,
				},
				{
					Name:     "MaxAge",
					Tag:      "",
					Id:       "int",
					Embedded: false,
					Offset:   104,
					Size:     8,
				},
				{
					Name:     "Secure",
					Tag:      "",
					Id:       "bool",
					Embedded: false,
					Offset:   112,
					Size:     1,
				},
				{
					Name:     "HttpOnly",
					Tag:      "",
					Id:       "bool",
					Embedded: false,
					Offset:   113,
					Size:     1,
				},
				{
					Name:     "SameSite",
					Tag:      "",
					Id:       "net/http.SameSite",
					Embedded: false,
					Offset:   120,
					Size:     8,
				},
				{
					Name:     "Raw",
					Tag:      "",
					Id:       "string",
					Embedded: false,
					Offset:   128,
					Size:     16,
				},
				{
					Name:     "Unparsed",
					Tag:      "",
					Id:       "[]string",
					Embedded: false,
					Offset:   144,
					Size:     24,
				},
			},
		},
		"net/http.MaxBytesError": Type{
			Name: "MaxBytesError",
			Path: "net/http",
			Fields: []Field{
				{
					Name:     "Limit",
					Tag:      "",
					Id:       "int64",
					Embedded: false,
					Offset:   0,
					Size:     8,
				},
			},
		},
		"net/http.ProtocolError": Type{
			Name: "ProtocolError",
			Path: "net/http",
			Fields: []Field{
				{
					Name:     "ErrorString",
					Tag:      "",
					Id:       "string",
					Embedded: false,
					Offset:   0,
					Size:     16,
				},
			},
		},
		"net/http.PushOptions": Type{
			Name: "PushOptions",
			Path: "net/http",
			Fields: []Field{
				{
					Name:     "Method",
					Tag:      "",
					Id:       "string",
					Embedded: false,
					Offset:   0,
					Size:     16,
				},
				{
					Name:     "Header",
					Tag:      "",
					Id:       "net/http.Header",
					Embedded: false,
					Offset:   16,
					Size:     8,
				},
			},
		},
		"net/http.Request": Type{
			Name: "Request",
			Path: "net/http",
			Fields: []Field{
				{
					Name:     "Method",
					Tag:      "",
					Id:       "string",
					Embedded: false,
					Offset:   0,
					Size:     16,
				},
				{
					Name:     "URL",
					Tag:      "",
					Id:       "*net/url.URL",
					Embedded: false,
					Offset:   16,
					Size:     8,
				},
				{
					Name:     "Proto",
					Tag:      "",
					Id:       "string",
					Embedded: false,
					Offset:   24,
					Size:     16,
				},
				{
					Name:     "ProtoMajor",
					Tag:      "",
					Id:       "int",
					Embedded: false,
					Offset:   40,
					Size:     8,
				},
				{
					Name:     "ProtoMinor",
					Tag:      "",
					Id:       "int",
					Embedded: false,
					Offset:   48,
					Size:     8,
				},
				{
					Name:     "Header",
					Tag:      "",
					Id:       "net/http.Header",
					Embedded: false,
					Offset:   56,
					Size:     8,
				},
				{
					Name:     "Body",
					Tag:      "",
					Id:       "io.ReadCloser",
					Embedded: false,
					Offset:   64,
					Size:     16,
				},
				{
					Name:     "GetBody",
					Tag:      "",
					Id:       "func() (io.ReadCloser, error)",
					Embedded: false,
					Offset:   80,
					Size:     8,
				},
				{
					Name:     "ContentLength",
					Tag:      "",
					Id:       "int64",
					Embedded: false,
					Offset:   88,
					Size:     8,
				},
				{
					Name:     "TransferEncoding",
					Tag:      "",
					Id:       "[]string",
					Embedded: false,
					Offset:   96,
					Size:     24,
				},
				{
					Name:     "Close",
					Tag:      "",
					Id:       "bool",
					Embedded: false,
					Offset:   120,
					Size:     1,
				},
				{
					Name:     "Host",
					Tag:      "",
					Id:       "string",
					Embedded: false,
					Offset:   128,
					Size:     16,
				},
				{
					Name:     "Form",
					Tag:      "",
					Id:       "net/url.Values",
					Embedded: false,
					Offset:   144,
					Size:     8,
				},
				{
					Name:     "PostForm",
					Tag:      "",
					Id:       "net/url.Values",
					Embedded: false,
					Offset:   152,
					Size:     8,
				},
				{
					Name:     "MultipartForm",
					Tag:      "",
					Id:       "*mime/multipart.Form",
					Embedded: false,
					Offset:   160,
					Size:     8,
				},
				{
					Name:     "Trailer",
					Tag:      "",
					Id:       "net/http.Header",
					Embedded: false,
					Offset:   168,
					Size:     8,
				},
				{
					Name:     "RemoteAddr",
					Tag:      "",
					Id:       "string",
					Embedded: false,
					Offset:   176,
					Size:     16,
				},
				{
					Name:     "RequestURI",
					Tag:      "",
					Id:       "string",
					Embedded: false,
					Offset:   192,
					Size:     16,
				},
				{
					Name:     "TLS",
					Tag:      "",
					Id:       "*crypto/tls.ConnectionState",
					Embedded: false,
					Offset:   208,
					Size:     8,
				},
				{
					Name:     "Cancel",
					Tag:      "",
					Id:       "<-chan struct{}",
					Embedded: false,
					Offset:   216,
					Size:     8,
				},
				{
					Name:     "Response",
					Tag:      "",
					Id:       "*net/http.Response",
					Embedded: false,
					Offset:   224,
					Size:     8,
				},
			},
		},
		"net/http.Response": Type{
			Name: "Response",
			Path: "net/http",
			Fields: []Field{
				{
					Name:     "Status",
					Tag:      "",
					Id:       "string",
					Embedded: false,
					Offset:   0,
					Size:     16,
				},
				{
					Name:     "StatusCode",
					Tag:      "",
					Id:       "int",
					Embedded: false,
					Offset:   16,
					Size:     8,
				},
				{
					Name:     "Proto",
					Tag:      "",
					Id:       "string",
					Embedded: false,
					Offset:   24,
					Size:     16,
				},
				{
					Name:     "ProtoMajor",
					Tag:      "",
					Id:       "int",
					Embedded: false,
					Offset:   40,
					Size:     8,
				},
				{
					Name:     "ProtoMinor",
					Tag:      "",
					Id:       "int",
					Embedded: false,
					Offset:   48,
					Size:     8,
				},
				{
					Name:     "Header",
					Tag:      "",
					Id:       "net/http.Header",
					Embedded: false,
					Offset:   56,
					Size:     8,
				},
				{
					Name:     "Body",
					Tag:      "",
					Id:       "io.ReadCloser",
					Embedded: false,
					Offset:   64,
					Size:     16,
				},
				{
					Name:     "ContentLength",
					Tag:      "",
					Id:       "int64",
					Embedded: false,
					Offset:   80,
					Size:     8,
				},
				{
					Name:     "TransferEncoding",
					Tag:      "",
					Id:       "[]string",
					Embedded: false,
					Offset:   88,
					Size:     24,
				},
				{
					Name:     "Close",
					Tag:      "",
					Id:       "bool",
					Embedded: false,
					Offset:   112,
					Size:     1,
				},
				{
					Name:     "Uncompressed",
					Tag:      "",
					Id:       "bool",
					Embedded: false,
					Offset:   113,
					Size:     1,
				},
				{
					Name:     "Trailer",
					Tag:      "",
					Id:       "net/http.Header",
					Embedded: false,
					Offset:   120,
					Size:     8,
				},
				{
					Name:     "Request",
					Tag:      "",
					Id:       "*net/http.Request",
					Embedded: false,
					Offset:   128,
					Size:     8,
				},
				{
					Name:     "TLS",
					Tag:      "",
					Id:       "*crypto/tls.ConnectionState",
					Embedded: false,
					Offset:   136,
					Size:     8,
				},
			},
		},
		"net/http.Server": Type{
			Name: "Server",
			Path: "net/http",
			Fields: []Field{
				{
					Name:     "Addr",
					Tag:      "",
					Id:       "string",
					Embedded: false,
					Offset:   0,
					Size:     16,
				},
				{
					Name:     "Handler",
					Tag:      "",
					Id:       "net/http.Handler",
					Embedded: false,
					Offset:   16,
					Size:     16,
				},
				{
					Name:     "DisableGeneralOptionsHandler",
					Tag:      "",
					Id:       "bool",
					Embedded: false,
					Offset:   32,
					Size:     1,
				},
				{
					Name:     "TLSConfig",
					Tag:      "",
					Id:       "*crypto/tls.Config",
					Embedded: false,
					Offset:   40,
					Size:     8,
				},
				{
					Name:     "ReadTimeout",
					Tag:      "",
					Id:       "time.Duration",
					Embedded: false,
					Offset:   48,
					Size:     8,
				},
				{
					Name:     "ReadHeaderTimeout",
					Tag:      "",
					Id:       "time.Duration",
					Embedded: false,
					Offset:   56,
					Size:     8,
				},
				{
					Name:     "WriteTimeout",
					Tag:      "",
					Id:       "time.Duration",
					Embedded: false,
					Offset:   64,
					Size:     8,
				},
				{
					Name:     "IdleTimeout",
					Tag:      "",
					Id:       "time.Duration",
					Embedded: false,
					Offset:   72,
					Size:     8,
				},
				{
					Name:     "MaxHeaderBytes",
					Tag:      "",
					Id:       "int",
					Embedded: false,
					Offset:   80,
					Size:     8,
				},
				{
					Name:     "TLSNextProto",
					Tag:      "",
					Id:       "map[string]func(*net/http.Server, *crypto/tls.Conn, net/http.Handler)",
					Embedded: false,
					Offset:   88,
					Size:     8,
				},
				{
					Name:     "ConnState",
					Tag:      "",
					Id:       "func(net.Conn, net/http.ConnState)",
					Embedded: false,
					Offset:   96,
					Size:     8,
				},
				{
					Name:     "ErrorLog",
					Tag:      "",
					Id:       "*log.Logger",
					Embedded: false,
					Offset:   104,
					Size:     8,
				},
				{
					Name:     "BaseContext",
					Tag:      "",
					Id:       "func(net.Listener) context.Context",
					Embedded: false,
					Offset:   112,
					Size:     8,
				},
				{
					Name: "ConnContext",
					Tag:  "",
					Id:   "func(ctx context.Context, c net.Conn) context.Context",

					Embedded: false,
					Offset:   120,
					Size:     8,
				},
			},
		},
		"net/http.Transport": Type{
			Name: "Transport",
			Path: "net/http",
			Fields: []Field{
				{
					Name:     "Proxy",
					Tag:      "",
					Id:       "func(*net/http.Request) (*net/url.URL, error)",
					Embedded: false,
					Offset:   112,
					Size:     8,
				},
				{
					Name:     "OnProxyConnectResponse",
					Tag:      "",
					Id:       "func(ctx context.Context, proxyURL *net/url.URL, connectReq *net/http.Request, connectRes *net/http.Response) error",
					Embedded: false,
					Offset:   120,
					Size:     8,
				},
				{
					Name:     "DialContext",
					Tag:      "",
					Id:       "func(ctx context.Context, network string, addr string) (net.Conn, error)",
					Embedded: false,
					Offset:   128,
					Size:     8,
				},
				{
					Name:     "Dial",
					Tag:      "",
					Id:       "func(network string, addr string) (net.Conn, error)",
					Embedded: false,
					Offset:   136,
					Size:     8,
				},
				{
					Name:     "DialTLSContext",
					Tag:      "",
					Id:       "func(ctx context.Context, network string, addr string) (net.Conn, error)",
					Embedded: false,
					Offset:   144,
					Size:     8,
				},
				{
					Name:     "DialTLS",
					Tag:      "",
					Id:       "func(network string, addr string) (net.Conn, error)",
					Embedded: false,
					Offset:   152,
					Size:     8,
				},
				{
					Name:     "TLSClientConfig",
					Tag:      "",
					Id:       "*crypto/tls.Config",
					Embedded: false,
					Offset:   160,
					Size:     8,
				},
				{
					Name:     "TLSHandshakeTimeout",
					Tag:      "",
					Id:       "time.Duration",
					Embedded: false,
					Offset:   168,
					Size:     8,
				},
				{
					Name:     "DisableKeepAlives",
					Tag:      "",
					Id:       "bool",
					Embedded: false,
					Offset:   176,
					Size:     1,
				},
				{
					Name:     "DisableCompression",
					Tag:      "",
					Id:       "bool",
					Embedded: false,
					Offset:   177,
					Size:     1,
				},
				{
					Name:     "MaxIdleConns",
					Tag:      "",
					Id:       "int",
					Embedded: false,
					Offset:   184,
					Size:     8,
				},
				{
					Name:     "MaxIdleConnsPerHost",
					Tag:      "",
					Id:       "int",
					Embedded: false,
					Offset:   192,
					Size:     8,
				},
				{
					Name:     "MaxConnsPerHost",
					Tag:      "",
					Id:       "int",
					Embedded: false,
					Offset:   200,
					Size:     8,
				},
				{
					Name:     "IdleConnTimeout",
					Tag:      "",
					Id:       "time.Duration",
					Embedded: false,
					Offset:   208,
					Size:     8,
				},
				{
					Name:     "ResponseHeaderTimeout",
					Tag:      "",
					Id:       "time.Duration",
					Embedded: false,
					Offset:   216,
					Size:     8,
				},
				{
					Name:     "ExpectContinueTimeout",
					Tag:      "",
					Id:       "time.Duration",
					Embedded: false,
					Offset:   224,
					Size:     8,
				},
				{
					Name:     "TLSNextProto",
					Tag:      "",
					Id:       "map[string]func(authority string, c *crypto/tls.Conn) net/http.RoundTripper",
					Embedded: false,
					Offset:   232,
					Size:     8,
				},
				{
					Name:     "ProxyConnectHeader",
					Tag:      "",
					Id:       "net/http.Header",
					Embedded: false,
					Offset:   240,
					Size:     8,
				},
				{
					Name:     "GetProxyConnectHeader",
					Tag:      "",
					Id:       "func(ctx context.Context, proxyURL *net/url.URL, target string) (net/http.Header, error)",
					Embedded: false,
					Offset:   248,
					Size:     8,
				},
				{
					Name:     "MaxResponseHeaderBytes",
					Tag:      "",
					Id:       "int64",
					Embedded: false,
					Offset:   256,
					Size:     8,
				},
				{
					Name:     "WriteBufferSize",
					Tag:      "",
					Id:       "int",
					Embedded: false,
					Offset:   264,
					Size:     8,
				},
				{
					Name:     "ReadBufferSize",
					Tag:      "",
					Id:       "int",
					Embedded: false,
					Offset:   272,
					Size:     8,
				},
				{
					Name:     "ForceAttemptHTTP2",
					Tag:      "",
					Id:       "bool",
					Embedded: false,
					Offset:   313,
					Size:     1,
				},
			},
		},
	}
)

func TestBuild(t *testing.T) {
	p := fn.Panic1(packages.Load(&packages.Config{
		Mode: packages.NeedTypes | packages.NeedTypesInfo,
	}, os.ExpandEnv("${GOROOT}/src/net/http")))
	pkg := p[0]
	if pkg.TypesSizes != nil {
		Sizes = pkg.TypesSizes
	} else {
		Sizes = types.SizesFor("gc", "amd64")
	}
	scope := pkg.Types.Scope()
	b := new(bytes.Buffer)
	x := new(bytes.Buffer)
	for _, s := range scope.Names() {
		if ast.IsExported(s) {
			o := scope.Lookup(s)
			if t, ok := o.(*types.TypeName); ok {
				if st, ok := t.Type().(*types.Named); ok {
					if v, ok := st.Underlying().(*types.Struct); ok {
						if BuildType(x, t.Pkg().Path(), t.Name(), v, true) > 0 {
							x.WriteByte(',')
							_, _ = fmt.Fprintf(b, "\"%s.%s\":", t.Pkg().Path(), t.Name())
							_, _ = x.WriteTo(b)
							b.WriteByte('\n')
						}
						x.Reset()
					}
				}
			}
		}
	}

	println(b.String())
}
func TestParseId(t *testing.T) {
	//for id := range typing {
	//	fmt.Printf("%s=>\t%#+v\n", id, ParseId(id))
	//}
	for id, t := range typing {
		for _, f := range t.Fields {
			fmt.Printf("%s#%s=>\t%s\n", id, f.Name, ParseId(f.Id))
		}
	}
}
