package providers

import (
	"encoding/json"
	"reflect"
	"testing"
)

type injectJSONSettings struct {
	Host string
}

func (s *injectJSONSettings) UnmarshalJSON(b []byte) error {
	type alias injectJSONSettings
	return json.Unmarshal(b, (*alias)(s))
}

func TestConfigProvider_writeInjectDest_pointerFieldFromStarInt(t *testing.T) {
	c := &ConfigProvider{}
	var port *int
	v := 8080
	src := &v
	c.writeInjectDest("app.port", &port, src)
	if port == nil || *port != 8080 {
		t.Fatalf("got %v", port)
	}
	if port == src {
		t.Fatal("field must hold copied value, not GetBean pointer address")
	}
}

func TestConfigProvider_writeInjectDest_pointerFieldFromInt64Star(t *testing.T) {
	c := &ConfigProvider{}
	var port *int
	var v int64 = 9090
	c.writeInjectDest("app.port", &port, &v)
	if port == nil || *port != 9090 {
		t.Fatalf("got %v", port)
	}
}

func TestConfigProvider_writeInjectDest_pointerFieldNilSrc(t *testing.T) {
	c := &ConfigProvider{}
	v := 1
	port := &v
	c.writeInjectDest("app.port", &port, (*int)(nil))
	if port != nil {
		t.Fatal("expected nil *int")
	}
}

func TestConfigProvider_writeInjectDest_structPointerField(t *testing.T) {
	c := &ConfigProvider{}
	type holder struct {
		P *int
	}
	h := holder{}
	v := 3000
	c.writeInjectDest("app.port", &h.P, &v)
	if h.P == nil || *h.P != 3000 {
		t.Fatalf("got %v", h.P)
	}
}

func TestConfigProvider_writeInjectDest_stringFieldFromInt(t *testing.T) {
	c := &ConfigProvider{}
	var port string
	v := 8080
	c.writeInjectDest("app.port", &port, &v)
	if port != "8080" {
		t.Fatalf("got %q", port)
	}
}

func TestConfigProvider_writeInjectDest_starStringFromInt(t *testing.T) {
	c := &ConfigProvider{}
	var port *string
	v := 8080
	c.writeInjectDest("app.port", &port, &v)
	if port == nil || *port != "8080" {
		t.Fatalf("got %v", port)
	}
}

func TestConfigProvider_writeInjectDest_jsonUnmarshaler(t *testing.T) {
	c := &ConfigProvider{}
	var cfg injectJSONSettings
	raw := `{"Host":"127.0.0.1"}`
	c.writeInjectDest("app.settings", &cfg, &raw)
	if cfg.Host != "127.0.0.1" {
		t.Fatalf("got %+v", cfg)
	}
}

func TestConfigProvider_writeInjectDest_jsonUnmarshalerPointerField(t *testing.T) {
	c := &ConfigProvider{}
	var cfg *injectJSONSettings
	raw := `{"Host":"localhost"}`
	c.writeInjectDest("app.settings", &cfg, raw)
	if cfg == nil || cfg.Host != "localhost" {
		t.Fatalf("got %+v", cfg)
	}
}

func TestConfigProvider_writeInjectDest_boolValueField(t *testing.T) {
	c := &ConfigProvider{}
	var debug bool
	src := true
	c.writeInjectDest("app.debug", &debug, &src)
	if !debug {
		t.Fatal("expected true")
	}
}

func TestConfigProvider_InjectValue_fieldAddressAssignable(t *testing.T) {
	type holder struct {
		P *int
	}
	h := holder{}
	rv := reflect.ValueOf(&h.P)
	if !rv.Elem().CanSet() {
		t.Fatal("&h.P should be assignable")
	}
}
