// Package apigin provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.4.1 DO NOT EDIT.
package apigin

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"
	"strings"
	"time"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/gin-gonic/gin"
)

// POFPCreateReq defines model for POFPCreateReq.
type POFPCreateReq struct {
	Pofp *POPFInfo `json:"pofp,omitempty"`
}

// POFPCreateResp defines model for POFPCreateResp.
type POFPCreateResp = map[string]interface{}

// POPFInfo defines model for POPFInfo.
type POPFInfo struct {
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	Id        *uint32    `json:"id,omitempty"`
	Name      *string    `json:"name,omitempty"`
	Ssss      *uint32    `json:"ssss,omitempty"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
}

// PetInfo defines model for PetInfo.
type PetInfo struct {
	Avatar    *string    `json:"avatar,omitempty"`
	BirthDate *string    `json:"birthDate,omitempty"`
	Breed     *string    `json:"breed,omitempty"`
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	Gender    *uint32    `json:"gender,omitempty"`
	Id        *uint32    `json:"id,omitempty"`
	Name      *string    `json:"name,omitempty"`
	Specie    *string    `json:"specie,omitempty"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
	Weight    *int32     `json:"weight,omitempty"`
}

// PetInfoReg 宠物注册信息
type PetInfoReg struct {
	// AvatarData base64 data
	AvatarData *string `json:"avatarData,omitempty"`

	// Name 名字
	Name *string `json:"name,omitempty"`
}

// UserInfo defines model for UserInfo.
type UserInfo struct {
	Id   *uint32    `json:"id,omitempty"`
	Pets *[]PetInfo `json:"pets,omitempty"`
}

// WeChatLoginReq 登录请求
type WeChatLoginReq struct {
	WxCode *string `json:"wxCode,omitempty"`
}

// WeChatLoginResp defines model for WeChatLoginResp.
type WeChatLoginResp struct {
	Data *WeChatLoginRespData `json:"data,omitempty"`
}

// WeChatLoginRespData defines model for WeChatLoginRespData.
type WeChatLoginRespData struct {
	Token    *string   `json:"token,omitempty"`
	UserInfo *UserInfo `json:"userInfo,omitempty"`
}

// WeChatRegisterFastReq defines model for WeChatRegisterFastReq.
type WeChatRegisterFastReq struct {
	// Pet 宠物注册信息
	Pet    *PetInfoReg `json:"pet,omitempty"`
	WxCode *string     `json:"wxCode,omitempty"`
}

// WeChatRegisterFastResp defines model for WeChatRegisterFastResp.
type WeChatRegisterFastResp struct {
	Data *WeChatLoginRespData `json:"data,omitempty"`
}

// BaseServicePOFPCreateJSONRequestBody defines body for BaseServicePOFPCreate for application/json ContentType.
type BaseServicePOFPCreateJSONRequestBody = POFPCreateReq

// BaseServiceWeChatLoginJSONRequestBody defines body for BaseServiceWeChatLogin for application/json ContentType.
type BaseServiceWeChatLoginJSONRequestBody = WeChatLoginReq

// BaseServiceWeChatRegisterFastJSONRequestBody defines body for BaseServiceWeChatRegisterFast for application/json ContentType.
type BaseServiceWeChatRegisterFastJSONRequestBody = WeChatRegisterFastReq

// RequestEditorFn  is the function signature for the RequestEditor callback function
type RequestEditorFn func(ctx context.Context, req *http.Request) error

// Doer performs HTTP requests.
//
// The standard http.Client implements this interface.
type HttpRequestDoer interface {
	Do(req *http.Request) (*http.Response, error)
}

// Client which conforms to the OpenAPI3 specification for this service.
type Client struct {
	// The endpoint of the server conforming to this interface, with scheme,
	// https://api.deepmap.com for example. This can contain a path relative
	// to the server, such as https://api.deepmap.com/dev-test, and all the
	// paths in the swagger spec will be appended to the server.
	Server string

	// Doer for performing requests, typically a *http.Client with any
	// customized settings, such as certificate chains.
	Client HttpRequestDoer

	// A list of callbacks for modifying requests which are generated before sending over
	// the network.
	RequestEditors []RequestEditorFn
}

// ClientOption allows setting custom parameters during construction
type ClientOption func(*Client) error

// Creates a new Client, with reasonable defaults
func NewClient(server string, opts ...ClientOption) (*Client, error) {
	// create a client with sane default values
	client := Client{
		Server: server,
	}
	// mutate client and add all optional params
	for _, o := range opts {
		if err := o(&client); err != nil {
			return nil, err
		}
	}
	// ensure the server URL always has a trailing slash
	if !strings.HasSuffix(client.Server, "/") {
		client.Server += "/"
	}
	// create httpClient, if not already present
	if client.Client == nil {
		client.Client = &http.Client{}
	}
	return &client, nil
}

// WithHTTPClient allows overriding the default Doer, which is
// automatically created using http.Client. This is useful for tests.
func WithHTTPClient(doer HttpRequestDoer) ClientOption {
	return func(c *Client) error {
		c.Client = doer
		return nil
	}
}

// WithRequestEditorFn allows setting up a callback function, which will be
// called right before sending the request. This can be used to mutate the request.
func WithRequestEditorFn(fn RequestEditorFn) ClientOption {
	return func(c *Client) error {
		c.RequestEditors = append(c.RequestEditors, fn)
		return nil
	}
}

// The interface specification for the client above.
type ClientInterface interface {
	// BaseServicePOFPCreateWithBody request with any body
	BaseServicePOFPCreateWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	BaseServicePOFPCreate(ctx context.Context, body BaseServicePOFPCreateJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// BaseServiceWeChatLoginWithBody request with any body
	BaseServiceWeChatLoginWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	BaseServiceWeChatLogin(ctx context.Context, body BaseServiceWeChatLoginJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// BaseServiceWeChatRegisterFastWithBody request with any body
	BaseServiceWeChatRegisterFastWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	BaseServiceWeChatRegisterFast(ctx context.Context, body BaseServiceWeChatRegisterFastJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)
}

func (c *Client) BaseServicePOFPCreateWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewBaseServicePOFPCreateRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) BaseServicePOFPCreate(ctx context.Context, body BaseServicePOFPCreateJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewBaseServicePOFPCreateRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) BaseServiceWeChatLoginWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewBaseServiceWeChatLoginRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) BaseServiceWeChatLogin(ctx context.Context, body BaseServiceWeChatLoginJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewBaseServiceWeChatLoginRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) BaseServiceWeChatRegisterFastWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewBaseServiceWeChatRegisterFastRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) BaseServiceWeChatRegisterFast(ctx context.Context, body BaseServiceWeChatRegisterFastJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewBaseServiceWeChatRegisterFastRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// NewBaseServicePOFPCreateRequest calls the generic BaseServicePOFPCreate builder with application/json body
func NewBaseServicePOFPCreateRequest(server string, body BaseServicePOFPCreateJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewBaseServicePOFPCreateRequestWithBody(server, "application/json", bodyReader)
}

// NewBaseServicePOFPCreateRequestWithBody generates requests for BaseServicePOFPCreate with any type of body
func NewBaseServicePOFPCreateRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/v1/popf/create")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

// NewBaseServiceWeChatLoginRequest calls the generic BaseServiceWeChatLogin builder with application/json body
func NewBaseServiceWeChatLoginRequest(server string, body BaseServiceWeChatLoginJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewBaseServiceWeChatLoginRequestWithBody(server, "application/json", bodyReader)
}

// NewBaseServiceWeChatLoginRequestWithBody generates requests for BaseServiceWeChatLogin with any type of body
func NewBaseServiceWeChatLoginRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/v1/user/wx/login")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

// NewBaseServiceWeChatRegisterFastRequest calls the generic BaseServiceWeChatRegisterFast builder with application/json body
func NewBaseServiceWeChatRegisterFastRequest(server string, body BaseServiceWeChatRegisterFastJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewBaseServiceWeChatRegisterFastRequestWithBody(server, "application/json", bodyReader)
}

// NewBaseServiceWeChatRegisterFastRequestWithBody generates requests for BaseServiceWeChatRegisterFast with any type of body
func NewBaseServiceWeChatRegisterFastRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/v1/user/wx/reg/fast")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

func (c *Client) applyEditors(ctx context.Context, req *http.Request, additionalEditors []RequestEditorFn) error {
	for _, r := range c.RequestEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	for _, r := range additionalEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	return nil
}

// ClientWithResponses builds on ClientInterface to offer response payloads
type ClientWithResponses struct {
	ClientInterface
}

// NewClientWithResponses creates a new ClientWithResponses, which wraps
// Client with return type handling
func NewClientWithResponses(server string, opts ...ClientOption) (*ClientWithResponses, error) {
	client, err := NewClient(server, opts...)
	if err != nil {
		return nil, err
	}
	return &ClientWithResponses{client}, nil
}

// WithBaseURL overrides the baseURL.
func WithBaseURL(baseURL string) ClientOption {
	return func(c *Client) error {
		newBaseURL, err := url.Parse(baseURL)
		if err != nil {
			return err
		}
		c.Server = newBaseURL.String()
		return nil
	}
}

// ClientWithResponsesInterface is the interface specification for the client with responses above.
type ClientWithResponsesInterface interface {
	// BaseServicePOFPCreateWithBodyWithResponse request with any body
	BaseServicePOFPCreateWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*BaseServicePOFPCreateResponse, error)

	BaseServicePOFPCreateWithResponse(ctx context.Context, body BaseServicePOFPCreateJSONRequestBody, reqEditors ...RequestEditorFn) (*BaseServicePOFPCreateResponse, error)

	// BaseServiceWeChatLoginWithBodyWithResponse request with any body
	BaseServiceWeChatLoginWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*BaseServiceWeChatLoginResponse, error)

	BaseServiceWeChatLoginWithResponse(ctx context.Context, body BaseServiceWeChatLoginJSONRequestBody, reqEditors ...RequestEditorFn) (*BaseServiceWeChatLoginResponse, error)

	// BaseServiceWeChatRegisterFastWithBodyWithResponse request with any body
	BaseServiceWeChatRegisterFastWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*BaseServiceWeChatRegisterFastResponse, error)

	BaseServiceWeChatRegisterFastWithResponse(ctx context.Context, body BaseServiceWeChatRegisterFastJSONRequestBody, reqEditors ...RequestEditorFn) (*BaseServiceWeChatRegisterFastResponse, error)
}

type BaseServicePOFPCreateResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *POFPCreateResp
}

// Status returns HTTPResponse.Status
func (r BaseServicePOFPCreateResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r BaseServicePOFPCreateResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type BaseServiceWeChatLoginResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *WeChatLoginResp
}

// Status returns HTTPResponse.Status
func (r BaseServiceWeChatLoginResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r BaseServiceWeChatLoginResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type BaseServiceWeChatRegisterFastResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *WeChatRegisterFastResp
}

// Status returns HTTPResponse.Status
func (r BaseServiceWeChatRegisterFastResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r BaseServiceWeChatRegisterFastResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// BaseServicePOFPCreateWithBodyWithResponse request with arbitrary body returning *BaseServicePOFPCreateResponse
func (c *ClientWithResponses) BaseServicePOFPCreateWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*BaseServicePOFPCreateResponse, error) {
	rsp, err := c.BaseServicePOFPCreateWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseBaseServicePOFPCreateResponse(rsp)
}

func (c *ClientWithResponses) BaseServicePOFPCreateWithResponse(ctx context.Context, body BaseServicePOFPCreateJSONRequestBody, reqEditors ...RequestEditorFn) (*BaseServicePOFPCreateResponse, error) {
	rsp, err := c.BaseServicePOFPCreate(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseBaseServicePOFPCreateResponse(rsp)
}

// BaseServiceWeChatLoginWithBodyWithResponse request with arbitrary body returning *BaseServiceWeChatLoginResponse
func (c *ClientWithResponses) BaseServiceWeChatLoginWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*BaseServiceWeChatLoginResponse, error) {
	rsp, err := c.BaseServiceWeChatLoginWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseBaseServiceWeChatLoginResponse(rsp)
}

func (c *ClientWithResponses) BaseServiceWeChatLoginWithResponse(ctx context.Context, body BaseServiceWeChatLoginJSONRequestBody, reqEditors ...RequestEditorFn) (*BaseServiceWeChatLoginResponse, error) {
	rsp, err := c.BaseServiceWeChatLogin(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseBaseServiceWeChatLoginResponse(rsp)
}

// BaseServiceWeChatRegisterFastWithBodyWithResponse request with arbitrary body returning *BaseServiceWeChatRegisterFastResponse
func (c *ClientWithResponses) BaseServiceWeChatRegisterFastWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*BaseServiceWeChatRegisterFastResponse, error) {
	rsp, err := c.BaseServiceWeChatRegisterFastWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseBaseServiceWeChatRegisterFastResponse(rsp)
}

func (c *ClientWithResponses) BaseServiceWeChatRegisterFastWithResponse(ctx context.Context, body BaseServiceWeChatRegisterFastJSONRequestBody, reqEditors ...RequestEditorFn) (*BaseServiceWeChatRegisterFastResponse, error) {
	rsp, err := c.BaseServiceWeChatRegisterFast(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseBaseServiceWeChatRegisterFastResponse(rsp)
}

// ParseBaseServicePOFPCreateResponse parses an HTTP response from a BaseServicePOFPCreateWithResponse call
func ParseBaseServicePOFPCreateResponse(rsp *http.Response) (*BaseServicePOFPCreateResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &BaseServicePOFPCreateResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest POFPCreateResp
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}

// ParseBaseServiceWeChatLoginResponse parses an HTTP response from a BaseServiceWeChatLoginWithResponse call
func ParseBaseServiceWeChatLoginResponse(rsp *http.Response) (*BaseServiceWeChatLoginResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &BaseServiceWeChatLoginResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest WeChatLoginResp
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}

// ParseBaseServiceWeChatRegisterFastResponse parses an HTTP response from a BaseServiceWeChatRegisterFastWithResponse call
func ParseBaseServiceWeChatRegisterFastResponse(rsp *http.Response) (*BaseServiceWeChatRegisterFastResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &BaseServiceWeChatRegisterFastResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest WeChatRegisterFastResp
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}

// ServerInterface represents all server handlers.
type ServerInterface interface {

	// (POST /v1/popf/create)
	BaseServicePOFPCreate(c *gin.Context)

	// (POST /v1/user/wx/login)
	BaseServiceWeChatLogin(c *gin.Context)

	// (POST /v1/user/wx/reg/fast)
	BaseServiceWeChatRegisterFast(c *gin.Context)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandler       func(*gin.Context, error, int)
}

type MiddlewareFunc func(c *gin.Context)

// BaseServicePOFPCreate operation middleware
func (siw *ServerInterfaceWrapper) BaseServicePOFPCreate(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.BaseServicePOFPCreate(c)
}

// BaseServiceWeChatLogin operation middleware
func (siw *ServerInterfaceWrapper) BaseServiceWeChatLogin(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.BaseServiceWeChatLogin(c)
}

// BaseServiceWeChatRegisterFast operation middleware
func (siw *ServerInterfaceWrapper) BaseServiceWeChatRegisterFast(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.BaseServiceWeChatRegisterFast(c)
}

// GinServerOptions provides options for the Gin server.
type GinServerOptions struct {
	BaseURL      string
	Middlewares  []MiddlewareFunc
	ErrorHandler func(*gin.Context, error, int)
}

// RegisterHandlers creates http.Handler with routing matching OpenAPI spec.
func RegisterHandlers(router gin.IRouter, si ServerInterface) {
	RegisterHandlersWithOptions(router, si, GinServerOptions{})
}

// RegisterHandlersWithOptions creates http.Handler with additional options
func RegisterHandlersWithOptions(router gin.IRouter, si ServerInterface, options GinServerOptions) {
	errorHandler := options.ErrorHandler
	if errorHandler == nil {
		errorHandler = func(c *gin.Context, err error, statusCode int) {
			c.JSON(statusCode, gin.H{"msg": err.Error()})
		}
	}

	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
		ErrorHandler:       errorHandler,
	}

	router.POST(options.BaseURL+"/v1/popf/create", wrapper.BaseServicePOFPCreate)
	router.POST(options.BaseURL+"/v1/user/wx/login", wrapper.BaseServiceWeChatLogin)
	router.POST(options.BaseURL+"/v1/user/wx/reg/fast", wrapper.BaseServiceWeChatRegisterFast)
}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/8RWX08bRxD/Kta2j67PQNUHv/GnSKhVsaiqPlAe1nfj81K8u+yODRayRHmoSltKHxJF",
	"IZESFEUiigKJkJIASvxlfDZ8i2jvDs5nX/wniZMX/9mZ+83M7zc7c1vEFmUpOHDUJLdFtF2CMvV/5hfn",
	"87MKKMISrJsDqYQEhQx8sxRFab6/VlAkOfKVFQFZIYqVX8zPL/CiIPV6mmBNAskRUVgFG0k9HQugfawE",
	"lxCgJ7ztP+hMo/lTFKpMkeSIQxG+QVYGchNPo2LcNWDMiflWGMepyciRcQQXlPHktAwd+UQQWms9LEhF",
	"OqNlmMgRYHL9tEqRqsQkC0xhaY5icgkFBeAkWj6AUhe4A2pYRj6FABJslmwame802QDmluL+78upjzhL",
	"4BoIB7StmEQmOMkR7/hhe/dJ6/TI+/PfZuOw9ccJSSdKOEeR9j5eoBq++zblGGM6yq5QQ9BJlVwT1pXE",
	"/3veszvDddovGlRyqw0vm4RgijCEsh44HMLWjpKhStFacna/wmyJ4o/CZTycRvFK23cvvDe3L09etV7s",
	"9BC9sTkrnKSmGRgpGEtxNCcUrF9tXSC+xsNEu26GeEQUvwNP7vkO0frlcyNunySWwGUaQc1TjckTH3BI",
	"Tc2NMNdrdN7jSYyZfnPEQvbi7XS1fdC+ddR8feD9fdi6v2c+/3vs7T/yjg+aZ7uZ33hq+fufVlLT+YWU",
	"KKZ+BlVlNmTMnWC4ZkLMUA3hsfEiaVIFpQPwbCabmTAlCwmcSkZyZCqTzUyZvqVY8qu0qhOWFLJoBUM5",
	"WLgaE674X/e8i/PLl6eXjbP2zhnxURU11gUnnki0b0maKFivgMYZ4dT8dSo4AvfxqZRrzPYRrFUtePRa",
	"MHjdd74x+PSaMEyZhYOqAv6BloLrQMvJbHYswbUMosepWvwhaALqapJb7mSGrBiD4dxcKGtj01ozvdOH",
	"9bfHzcah93y/ffSPd74fTJ+gRfop0NGXY5Kga05+Zg26Z+dHiqDAtYo04H8oHbzG06vtB6Ea/urtvLaD",
	"lemcPmMVqHvWfhGdembtyHJFlq3wDSTmUV+pvwsAAP//aIKAJGAMAAA=",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %w", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	res := make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	resolvePath := PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		pathToFile := url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
