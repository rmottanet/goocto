// tests/pkg/mock/mock_http_client.go
package mock

import (
    "net/http"
    "net/http/httptest"
    "bytes"
    "io"
)


// MockHTTPClient é um cliente HTTP de mock para uso em testes.
type MockHTTPClient struct {
    DoFunc        func(req *http.Request) (*http.Response, error)
    RoundTripFunc func(req *http.Request) (*http.Response, error)
}

// Do executa a função DoFunc do cliente mock ao receber uma solicitação HTTP.
func (m *MockHTTPClient) Do(req *http.Request) (*http.Response, error) {
    return m.DoFunc(req)
}

// RoundTrip executa a função RoundTripFunc do cliente mock.
func (m *MockHTTPClient) RoundTrip(req *http.Request) (*http.Response, error) {
    return m.RoundTripFunc(req)
}


// Estrutura para armazenar os dados do corpo do mock.
type mockBody struct {
    data []byte
}

// NewMockBody cria um corpo de resposta mock com base nos dados fornecidos.
func NewMockBody(data []byte) io.ReadCloser {
    return &mockBody{data: data}
}


// NewMockResponse cria um novo objeto *http.Response a partir dos dados fornecidos.
func NewMockResponse(statusCode int, body []byte) *http.Response {
    return &http.Response{
        StatusCode: statusCode,
        Body:       NewMockBody(body),
    }
}


// Read lê os dados do corpo do mock e os escreve no slice de bytes fornecido.
func (m *mockBody) Read(p []byte) (n int, err error) {
    return bytes.NewBuffer(m.data).Read(p)
}


// Close fecha o corpo do mock.
func (m *mockBody) Close() error {
    return nil
}


// NewTestServer cria um servidor de teste para uso nos testes.
func NewTestServer(handler http.Handler) *httptest.Server {
    return httptest.NewServer(handler)
}


// Client é uma implementação de um cliente HTTP para uso em testes.
type Client struct {
    DoFunc        func(req *http.Request) (*http.Response, error)
    RoundTripFunc func(req *http.Request) (*http.Response, error)
}

// Do executa a função DoFunc do cliente mock ao receber uma solicitação HTTP.
func (c *Client) Do(req *http.Request) (*http.Response, error) {
    return c.DoFunc(req)
}

// RoundTrip executa a função RoundTripFunc do cliente mock.
func (c *Client) RoundTrip(req *http.Request) (*http.Response, error) {
    return c.RoundTripFunc(req)
}

