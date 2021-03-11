package api_connection

import (
	"context"
	"encoding/json"
)

// StubTransport provides Transport interface with preconfigured ResponseBody
type StubTransport struct {
	RequestBody  string
	RequestPath  string
	ResponseBody string
}

func (t *StubTransport) Get(ctx context.Context, path string) (json.RawMessage, error) {
	t.RequestPath = path
	return json.RawMessage(t.ResponseBody), nil
}

func (t *StubTransport) Post(ctx context.Context, path string, payload json.RawMessage) (json.RawMessage, error) {
	t.RequestBody = string(payload)
	t.RequestPath = path
	return json.RawMessage(t.ResponseBody), nil
}

func (t *StubTransport) Put(ctx context.Context, path string, payload json.RawMessage) (json.RawMessage, error) {
	t.RequestBody = string(payload)
	t.RequestPath = path
	return json.RawMessage(t.ResponseBody), nil
}

func (t *StubTransport) Delete(ctx context.Context, path string) (responseBody json.RawMessage, err error) {
	t.RequestPath = path
	return json.RawMessage(t.ResponseBody), nil
}
