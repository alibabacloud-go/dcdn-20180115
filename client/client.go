// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type AddDcdnDomainRequest struct {
	CheckUrl        *string `json:"CheckUrl,omitempty" xml:"CheckUrl,omitempty"`
	DomainName      *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	OwnerAccount    *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId         *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Scope           *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
	SecurityToken   *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	Sources         *string `json:"Sources,omitempty" xml:"Sources,omitempty"`
	TopLevelDomain  *string `json:"TopLevelDomain,omitempty" xml:"TopLevelDomain,omitempty"`
}

func (s AddDcdnDomainRequest) String() string {
	return tea.Prettify(s)
}

func (s AddDcdnDomainRequest) GoString() string {
	return s.String()
}

func (s *AddDcdnDomainRequest) SetCheckUrl(v string) *AddDcdnDomainRequest {
	s.CheckUrl = &v
	return s
}

func (s *AddDcdnDomainRequest) SetDomainName(v string) *AddDcdnDomainRequest {
	s.DomainName = &v
	return s
}

func (s *AddDcdnDomainRequest) SetOwnerAccount(v string) *AddDcdnDomainRequest {
	s.OwnerAccount = &v
	return s
}

func (s *AddDcdnDomainRequest) SetOwnerId(v int64) *AddDcdnDomainRequest {
	s.OwnerId = &v
	return s
}

func (s *AddDcdnDomainRequest) SetResourceGroupId(v string) *AddDcdnDomainRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *AddDcdnDomainRequest) SetScope(v string) *AddDcdnDomainRequest {
	s.Scope = &v
	return s
}

func (s *AddDcdnDomainRequest) SetSecurityToken(v string) *AddDcdnDomainRequest {
	s.SecurityToken = &v
	return s
}

func (s *AddDcdnDomainRequest) SetSources(v string) *AddDcdnDomainRequest {
	s.Sources = &v
	return s
}

func (s *AddDcdnDomainRequest) SetTopLevelDomain(v string) *AddDcdnDomainRequest {
	s.TopLevelDomain = &v
	return s
}

type AddDcdnDomainResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddDcdnDomainResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddDcdnDomainResponseBody) GoString() string {
	return s.String()
}

func (s *AddDcdnDomainResponseBody) SetRequestId(v string) *AddDcdnDomainResponseBody {
	s.RequestId = &v
	return s
}

type AddDcdnDomainResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddDcdnDomainResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddDcdnDomainResponse) String() string {
	return tea.Prettify(s)
}

func (s AddDcdnDomainResponse) GoString() string {
	return s.String()
}

func (s *AddDcdnDomainResponse) SetHeaders(v map[string]*string) *AddDcdnDomainResponse {
	s.Headers = v
	return s
}

func (s *AddDcdnDomainResponse) SetBody(v *AddDcdnDomainResponseBody) *AddDcdnDomainResponse {
	s.Body = v
	return s
}

type AddDcdnIpaDomainRequest struct {
	CheckUrl        *string `json:"CheckUrl,omitempty" xml:"CheckUrl,omitempty"`
	DomainName      *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	OwnerAccount    *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId         *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	Protocol        *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Scope           *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
	SecurityToken   *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	Sources         *string `json:"Sources,omitempty" xml:"Sources,omitempty"`
	TopLevelDomain  *string `json:"TopLevelDomain,omitempty" xml:"TopLevelDomain,omitempty"`
}

func (s AddDcdnIpaDomainRequest) String() string {
	return tea.Prettify(s)
}

func (s AddDcdnIpaDomainRequest) GoString() string {
	return s.String()
}

func (s *AddDcdnIpaDomainRequest) SetCheckUrl(v string) *AddDcdnIpaDomainRequest {
	s.CheckUrl = &v
	return s
}

func (s *AddDcdnIpaDomainRequest) SetDomainName(v string) *AddDcdnIpaDomainRequest {
	s.DomainName = &v
	return s
}

func (s *AddDcdnIpaDomainRequest) SetOwnerAccount(v string) *AddDcdnIpaDomainRequest {
	s.OwnerAccount = &v
	return s
}

func (s *AddDcdnIpaDomainRequest) SetOwnerId(v int64) *AddDcdnIpaDomainRequest {
	s.OwnerId = &v
	return s
}

func (s *AddDcdnIpaDomainRequest) SetProtocol(v string) *AddDcdnIpaDomainRequest {
	s.Protocol = &v
	return s
}

func (s *AddDcdnIpaDomainRequest) SetResourceGroupId(v string) *AddDcdnIpaDomainRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *AddDcdnIpaDomainRequest) SetScope(v string) *AddDcdnIpaDomainRequest {
	s.Scope = &v
	return s
}

func (s *AddDcdnIpaDomainRequest) SetSecurityToken(v string) *AddDcdnIpaDomainRequest {
	s.SecurityToken = &v
	return s
}

func (s *AddDcdnIpaDomainRequest) SetSources(v string) *AddDcdnIpaDomainRequest {
	s.Sources = &v
	return s
}

func (s *AddDcdnIpaDomainRequest) SetTopLevelDomain(v string) *AddDcdnIpaDomainRequest {
	s.TopLevelDomain = &v
	return s
}

type AddDcdnIpaDomainResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddDcdnIpaDomainResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddDcdnIpaDomainResponseBody) GoString() string {
	return s.String()
}

func (s *AddDcdnIpaDomainResponseBody) SetRequestId(v string) *AddDcdnIpaDomainResponseBody {
	s.RequestId = &v
	return s
}

type AddDcdnIpaDomainResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddDcdnIpaDomainResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddDcdnIpaDomainResponse) String() string {
	return tea.Prettify(s)
}

func (s AddDcdnIpaDomainResponse) GoString() string {
	return s.String()
}

func (s *AddDcdnIpaDomainResponse) SetHeaders(v map[string]*string) *AddDcdnIpaDomainResponse {
	s.Headers = v
	return s
}

func (s *AddDcdnIpaDomainResponse) SetBody(v *AddDcdnIpaDomainResponseBody) *AddDcdnIpaDomainResponse {
	s.Body = v
	return s
}

type BatchAddDcdnDomainRequest struct {
	CheckUrl        *string `json:"CheckUrl,omitempty" xml:"CheckUrl,omitempty"`
	DomainName      *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	OwnerAccount    *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId         *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Scope           *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
	SecurityToken   *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	Sources         *string `json:"Sources,omitempty" xml:"Sources,omitempty"`
	TopLevelDomain  *string `json:"TopLevelDomain,omitempty" xml:"TopLevelDomain,omitempty"`
}

func (s BatchAddDcdnDomainRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchAddDcdnDomainRequest) GoString() string {
	return s.String()
}

func (s *BatchAddDcdnDomainRequest) SetCheckUrl(v string) *BatchAddDcdnDomainRequest {
	s.CheckUrl = &v
	return s
}

func (s *BatchAddDcdnDomainRequest) SetDomainName(v string) *BatchAddDcdnDomainRequest {
	s.DomainName = &v
	return s
}

func (s *BatchAddDcdnDomainRequest) SetOwnerAccount(v string) *BatchAddDcdnDomainRequest {
	s.OwnerAccount = &v
	return s
}

func (s *BatchAddDcdnDomainRequest) SetOwnerId(v int64) *BatchAddDcdnDomainRequest {
	s.OwnerId = &v
	return s
}

func (s *BatchAddDcdnDomainRequest) SetResourceGroupId(v string) *BatchAddDcdnDomainRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *BatchAddDcdnDomainRequest) SetScope(v string) *BatchAddDcdnDomainRequest {
	s.Scope = &v
	return s
}

func (s *BatchAddDcdnDomainRequest) SetSecurityToken(v string) *BatchAddDcdnDomainRequest {
	s.SecurityToken = &v
	return s
}

func (s *BatchAddDcdnDomainRequest) SetSources(v string) *BatchAddDcdnDomainRequest {
	s.Sources = &v
	return s
}

func (s *BatchAddDcdnDomainRequest) SetTopLevelDomain(v string) *BatchAddDcdnDomainRequest {
	s.TopLevelDomain = &v
	return s
}

type BatchAddDcdnDomainResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BatchAddDcdnDomainResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchAddDcdnDomainResponseBody) GoString() string {
	return s.String()
}

func (s *BatchAddDcdnDomainResponseBody) SetRequestId(v string) *BatchAddDcdnDomainResponseBody {
	s.RequestId = &v
	return s
}

type BatchAddDcdnDomainResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *BatchAddDcdnDomainResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BatchAddDcdnDomainResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchAddDcdnDomainResponse) GoString() string {
	return s.String()
}

func (s *BatchAddDcdnDomainResponse) SetHeaders(v map[string]*string) *BatchAddDcdnDomainResponse {
	s.Headers = v
	return s
}

func (s *BatchAddDcdnDomainResponse) SetBody(v *BatchAddDcdnDomainResponseBody) *BatchAddDcdnDomainResponse {
	s.Body = v
	return s
}

type BatchDeleteDcdnDomainConfigsRequest struct {
	DomainNames   *string `json:"DomainNames,omitempty" xml:"DomainNames,omitempty"`
	FunctionNames *string `json:"FunctionNames,omitempty" xml:"FunctionNames,omitempty"`
	OwnerAccount  *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s BatchDeleteDcdnDomainConfigsRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchDeleteDcdnDomainConfigsRequest) GoString() string {
	return s.String()
}

func (s *BatchDeleteDcdnDomainConfigsRequest) SetDomainNames(v string) *BatchDeleteDcdnDomainConfigsRequest {
	s.DomainNames = &v
	return s
}

func (s *BatchDeleteDcdnDomainConfigsRequest) SetFunctionNames(v string) *BatchDeleteDcdnDomainConfigsRequest {
	s.FunctionNames = &v
	return s
}

func (s *BatchDeleteDcdnDomainConfigsRequest) SetOwnerAccount(v string) *BatchDeleteDcdnDomainConfigsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *BatchDeleteDcdnDomainConfigsRequest) SetOwnerId(v int64) *BatchDeleteDcdnDomainConfigsRequest {
	s.OwnerId = &v
	return s
}

func (s *BatchDeleteDcdnDomainConfigsRequest) SetSecurityToken(v string) *BatchDeleteDcdnDomainConfigsRequest {
	s.SecurityToken = &v
	return s
}

type BatchDeleteDcdnDomainConfigsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BatchDeleteDcdnDomainConfigsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchDeleteDcdnDomainConfigsResponseBody) GoString() string {
	return s.String()
}

func (s *BatchDeleteDcdnDomainConfigsResponseBody) SetRequestId(v string) *BatchDeleteDcdnDomainConfigsResponseBody {
	s.RequestId = &v
	return s
}

type BatchDeleteDcdnDomainConfigsResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *BatchDeleteDcdnDomainConfigsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BatchDeleteDcdnDomainConfigsResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchDeleteDcdnDomainConfigsResponse) GoString() string {
	return s.String()
}

func (s *BatchDeleteDcdnDomainConfigsResponse) SetHeaders(v map[string]*string) *BatchDeleteDcdnDomainConfigsResponse {
	s.Headers = v
	return s
}

func (s *BatchDeleteDcdnDomainConfigsResponse) SetBody(v *BatchDeleteDcdnDomainConfigsResponseBody) *BatchDeleteDcdnDomainConfigsResponse {
	s.Body = v
	return s
}

type BatchSetDcdnDomainCertificateRequest struct {
	CertName      *string `json:"CertName,omitempty" xml:"CertName,omitempty"`
	CertType      *string `json:"CertType,omitempty" xml:"CertType,omitempty"`
	DomainName    *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	Region        *string `json:"Region,omitempty" xml:"Region,omitempty"`
	SSLPri        *string `json:"SSLPri,omitempty" xml:"SSLPri,omitempty"`
	SSLProtocol   *string `json:"SSLProtocol,omitempty" xml:"SSLProtocol,omitempty"`
	SSLPub        *string `json:"SSLPub,omitempty" xml:"SSLPub,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s BatchSetDcdnDomainCertificateRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchSetDcdnDomainCertificateRequest) GoString() string {
	return s.String()
}

func (s *BatchSetDcdnDomainCertificateRequest) SetCertName(v string) *BatchSetDcdnDomainCertificateRequest {
	s.CertName = &v
	return s
}

func (s *BatchSetDcdnDomainCertificateRequest) SetCertType(v string) *BatchSetDcdnDomainCertificateRequest {
	s.CertType = &v
	return s
}

func (s *BatchSetDcdnDomainCertificateRequest) SetDomainName(v string) *BatchSetDcdnDomainCertificateRequest {
	s.DomainName = &v
	return s
}

func (s *BatchSetDcdnDomainCertificateRequest) SetOwnerId(v int64) *BatchSetDcdnDomainCertificateRequest {
	s.OwnerId = &v
	return s
}

func (s *BatchSetDcdnDomainCertificateRequest) SetRegion(v string) *BatchSetDcdnDomainCertificateRequest {
	s.Region = &v
	return s
}

func (s *BatchSetDcdnDomainCertificateRequest) SetSSLPri(v string) *BatchSetDcdnDomainCertificateRequest {
	s.SSLPri = &v
	return s
}

func (s *BatchSetDcdnDomainCertificateRequest) SetSSLProtocol(v string) *BatchSetDcdnDomainCertificateRequest {
	s.SSLProtocol = &v
	return s
}

func (s *BatchSetDcdnDomainCertificateRequest) SetSSLPub(v string) *BatchSetDcdnDomainCertificateRequest {
	s.SSLPub = &v
	return s
}

func (s *BatchSetDcdnDomainCertificateRequest) SetSecurityToken(v string) *BatchSetDcdnDomainCertificateRequest {
	s.SecurityToken = &v
	return s
}

type BatchSetDcdnDomainCertificateResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BatchSetDcdnDomainCertificateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchSetDcdnDomainCertificateResponseBody) GoString() string {
	return s.String()
}

func (s *BatchSetDcdnDomainCertificateResponseBody) SetRequestId(v string) *BatchSetDcdnDomainCertificateResponseBody {
	s.RequestId = &v
	return s
}

type BatchSetDcdnDomainCertificateResponse struct {
	Headers map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *BatchSetDcdnDomainCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BatchSetDcdnDomainCertificateResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchSetDcdnDomainCertificateResponse) GoString() string {
	return s.String()
}

func (s *BatchSetDcdnDomainCertificateResponse) SetHeaders(v map[string]*string) *BatchSetDcdnDomainCertificateResponse {
	s.Headers = v
	return s
}

func (s *BatchSetDcdnDomainCertificateResponse) SetBody(v *BatchSetDcdnDomainCertificateResponseBody) *BatchSetDcdnDomainCertificateResponse {
	s.Body = v
	return s
}

type BatchSetDcdnDomainConfigsRequest struct {
	DomainNames   *string `json:"DomainNames,omitempty" xml:"DomainNames,omitempty"`
	Functions     *string `json:"Functions,omitempty" xml:"Functions,omitempty"`
	OwnerAccount  *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s BatchSetDcdnDomainConfigsRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchSetDcdnDomainConfigsRequest) GoString() string {
	return s.String()
}

func (s *BatchSetDcdnDomainConfigsRequest) SetDomainNames(v string) *BatchSetDcdnDomainConfigsRequest {
	s.DomainNames = &v
	return s
}

func (s *BatchSetDcdnDomainConfigsRequest) SetFunctions(v string) *BatchSetDcdnDomainConfigsRequest {
	s.Functions = &v
	return s
}

func (s *BatchSetDcdnDomainConfigsRequest) SetOwnerAccount(v string) *BatchSetDcdnDomainConfigsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *BatchSetDcdnDomainConfigsRequest) SetOwnerId(v int64) *BatchSetDcdnDomainConfigsRequest {
	s.OwnerId = &v
	return s
}

func (s *BatchSetDcdnDomainConfigsRequest) SetSecurityToken(v string) *BatchSetDcdnDomainConfigsRequest {
	s.SecurityToken = &v
	return s
}

type BatchSetDcdnDomainConfigsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BatchSetDcdnDomainConfigsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchSetDcdnDomainConfigsResponseBody) GoString() string {
	return s.String()
}

func (s *BatchSetDcdnDomainConfigsResponseBody) SetRequestId(v string) *BatchSetDcdnDomainConfigsResponseBody {
	s.RequestId = &v
	return s
}

type BatchSetDcdnDomainConfigsResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *BatchSetDcdnDomainConfigsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BatchSetDcdnDomainConfigsResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchSetDcdnDomainConfigsResponse) GoString() string {
	return s.String()
}

func (s *BatchSetDcdnDomainConfigsResponse) SetHeaders(v map[string]*string) *BatchSetDcdnDomainConfigsResponse {
	s.Headers = v
	return s
}

func (s *BatchSetDcdnDomainConfigsResponse) SetBody(v *BatchSetDcdnDomainConfigsResponseBody) *BatchSetDcdnDomainConfigsResponse {
	s.Body = v
	return s
}

type BatchSetDcdnIpaDomainConfigsRequest struct {
	DomainNames   *string `json:"DomainNames,omitempty" xml:"DomainNames,omitempty"`
	Functions     *string `json:"Functions,omitempty" xml:"Functions,omitempty"`
	OwnerAccount  *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s BatchSetDcdnIpaDomainConfigsRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchSetDcdnIpaDomainConfigsRequest) GoString() string {
	return s.String()
}

func (s *BatchSetDcdnIpaDomainConfigsRequest) SetDomainNames(v string) *BatchSetDcdnIpaDomainConfigsRequest {
	s.DomainNames = &v
	return s
}

func (s *BatchSetDcdnIpaDomainConfigsRequest) SetFunctions(v string) *BatchSetDcdnIpaDomainConfigsRequest {
	s.Functions = &v
	return s
}

func (s *BatchSetDcdnIpaDomainConfigsRequest) SetOwnerAccount(v string) *BatchSetDcdnIpaDomainConfigsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *BatchSetDcdnIpaDomainConfigsRequest) SetOwnerId(v int64) *BatchSetDcdnIpaDomainConfigsRequest {
	s.OwnerId = &v
	return s
}

func (s *BatchSetDcdnIpaDomainConfigsRequest) SetSecurityToken(v string) *BatchSetDcdnIpaDomainConfigsRequest {
	s.SecurityToken = &v
	return s
}

type BatchSetDcdnIpaDomainConfigsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BatchSetDcdnIpaDomainConfigsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchSetDcdnIpaDomainConfigsResponseBody) GoString() string {
	return s.String()
}

func (s *BatchSetDcdnIpaDomainConfigsResponseBody) SetRequestId(v string) *BatchSetDcdnIpaDomainConfigsResponseBody {
	s.RequestId = &v
	return s
}

type BatchSetDcdnIpaDomainConfigsResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *BatchSetDcdnIpaDomainConfigsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BatchSetDcdnIpaDomainConfigsResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchSetDcdnIpaDomainConfigsResponse) GoString() string {
	return s.String()
}

func (s *BatchSetDcdnIpaDomainConfigsResponse) SetHeaders(v map[string]*string) *BatchSetDcdnIpaDomainConfigsResponse {
	s.Headers = v
	return s
}

func (s *BatchSetDcdnIpaDomainConfigsResponse) SetBody(v *BatchSetDcdnIpaDomainConfigsResponseBody) *BatchSetDcdnIpaDomainConfigsResponse {
	s.Body = v
	return s
}

type BatchStartDcdnDomainRequest struct {
	DomainNames   *string `json:"DomainNames,omitempty" xml:"DomainNames,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s BatchStartDcdnDomainRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchStartDcdnDomainRequest) GoString() string {
	return s.String()
}

func (s *BatchStartDcdnDomainRequest) SetDomainNames(v string) *BatchStartDcdnDomainRequest {
	s.DomainNames = &v
	return s
}

func (s *BatchStartDcdnDomainRequest) SetOwnerId(v int64) *BatchStartDcdnDomainRequest {
	s.OwnerId = &v
	return s
}

func (s *BatchStartDcdnDomainRequest) SetSecurityToken(v string) *BatchStartDcdnDomainRequest {
	s.SecurityToken = &v
	return s
}

type BatchStartDcdnDomainResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BatchStartDcdnDomainResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchStartDcdnDomainResponseBody) GoString() string {
	return s.String()
}

func (s *BatchStartDcdnDomainResponseBody) SetRequestId(v string) *BatchStartDcdnDomainResponseBody {
	s.RequestId = &v
	return s
}

type BatchStartDcdnDomainResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *BatchStartDcdnDomainResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BatchStartDcdnDomainResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchStartDcdnDomainResponse) GoString() string {
	return s.String()
}

func (s *BatchStartDcdnDomainResponse) SetHeaders(v map[string]*string) *BatchStartDcdnDomainResponse {
	s.Headers = v
	return s
}

func (s *BatchStartDcdnDomainResponse) SetBody(v *BatchStartDcdnDomainResponseBody) *BatchStartDcdnDomainResponse {
	s.Body = v
	return s
}

type BatchStopDcdnDomainRequest struct {
	DomainNames   *string `json:"DomainNames,omitempty" xml:"DomainNames,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s BatchStopDcdnDomainRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchStopDcdnDomainRequest) GoString() string {
	return s.String()
}

func (s *BatchStopDcdnDomainRequest) SetDomainNames(v string) *BatchStopDcdnDomainRequest {
	s.DomainNames = &v
	return s
}

func (s *BatchStopDcdnDomainRequest) SetOwnerId(v int64) *BatchStopDcdnDomainRequest {
	s.OwnerId = &v
	return s
}

func (s *BatchStopDcdnDomainRequest) SetSecurityToken(v string) *BatchStopDcdnDomainRequest {
	s.SecurityToken = &v
	return s
}

type BatchStopDcdnDomainResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BatchStopDcdnDomainResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchStopDcdnDomainResponseBody) GoString() string {
	return s.String()
}

func (s *BatchStopDcdnDomainResponseBody) SetRequestId(v string) *BatchStopDcdnDomainResponseBody {
	s.RequestId = &v
	return s
}

type BatchStopDcdnDomainResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *BatchStopDcdnDomainResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BatchStopDcdnDomainResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchStopDcdnDomainResponse) GoString() string {
	return s.String()
}

func (s *BatchStopDcdnDomainResponse) SetHeaders(v map[string]*string) *BatchStopDcdnDomainResponse {
	s.Headers = v
	return s
}

func (s *BatchStopDcdnDomainResponse) SetBody(v *BatchStopDcdnDomainResponseBody) *BatchStopDcdnDomainResponse {
	s.Body = v
	return s
}

type CheckDcdnProjectExistRequest struct {
	OwnerId     *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
}

func (s CheckDcdnProjectExistRequest) String() string {
	return tea.Prettify(s)
}

func (s CheckDcdnProjectExistRequest) GoString() string {
	return s.String()
}

func (s *CheckDcdnProjectExistRequest) SetOwnerId(v int64) *CheckDcdnProjectExistRequest {
	s.OwnerId = &v
	return s
}

func (s *CheckDcdnProjectExistRequest) SetProjectName(v string) *CheckDcdnProjectExistRequest {
	s.ProjectName = &v
	return s
}

type CheckDcdnProjectExistResponseBody struct {
	Content   *CheckDcdnProjectExistResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	RequestId *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CheckDcdnProjectExistResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CheckDcdnProjectExistResponseBody) GoString() string {
	return s.String()
}

func (s *CheckDcdnProjectExistResponseBody) SetContent(v *CheckDcdnProjectExistResponseBodyContent) *CheckDcdnProjectExistResponseBody {
	s.Content = v
	return s
}

func (s *CheckDcdnProjectExistResponseBody) SetRequestId(v string) *CheckDcdnProjectExistResponseBody {
	s.RequestId = &v
	return s
}

type CheckDcdnProjectExistResponseBodyContent struct {
	Exist *string `json:"Exist,omitempty" xml:"Exist,omitempty"`
}

func (s CheckDcdnProjectExistResponseBodyContent) String() string {
	return tea.Prettify(s)
}

func (s CheckDcdnProjectExistResponseBodyContent) GoString() string {
	return s.String()
}

func (s *CheckDcdnProjectExistResponseBodyContent) SetExist(v string) *CheckDcdnProjectExistResponseBodyContent {
	s.Exist = &v
	return s
}

type CheckDcdnProjectExistResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CheckDcdnProjectExistResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CheckDcdnProjectExistResponse) String() string {
	return tea.Prettify(s)
}

func (s CheckDcdnProjectExistResponse) GoString() string {
	return s.String()
}

func (s *CheckDcdnProjectExistResponse) SetHeaders(v map[string]*string) *CheckDcdnProjectExistResponse {
	s.Headers = v
	return s
}

func (s *CheckDcdnProjectExistResponse) SetBody(v *CheckDcdnProjectExistResponseBody) *CheckDcdnProjectExistResponse {
	s.Body = v
	return s
}

type CommitStagingRoutineCodeRequest struct {
	CodeDescription *string `json:"CodeDescription,omitempty" xml:"CodeDescription,omitempty"`
	Name            *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerId         *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s CommitStagingRoutineCodeRequest) String() string {
	return tea.Prettify(s)
}

func (s CommitStagingRoutineCodeRequest) GoString() string {
	return s.String()
}

func (s *CommitStagingRoutineCodeRequest) SetCodeDescription(v string) *CommitStagingRoutineCodeRequest {
	s.CodeDescription = &v
	return s
}

func (s *CommitStagingRoutineCodeRequest) SetName(v string) *CommitStagingRoutineCodeRequest {
	s.Name = &v
	return s
}

func (s *CommitStagingRoutineCodeRequest) SetOwnerId(v int64) *CommitStagingRoutineCodeRequest {
	s.OwnerId = &v
	return s
}

type CommitStagingRoutineCodeResponseBody struct {
	Content   map[string]interface{} `json:"Content,omitempty" xml:"Content,omitempty"`
	RequestId *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CommitStagingRoutineCodeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CommitStagingRoutineCodeResponseBody) GoString() string {
	return s.String()
}

func (s *CommitStagingRoutineCodeResponseBody) SetContent(v map[string]interface{}) *CommitStagingRoutineCodeResponseBody {
	s.Content = v
	return s
}

func (s *CommitStagingRoutineCodeResponseBody) SetRequestId(v string) *CommitStagingRoutineCodeResponseBody {
	s.RequestId = &v
	return s
}

type CommitStagingRoutineCodeResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CommitStagingRoutineCodeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CommitStagingRoutineCodeResponse) String() string {
	return tea.Prettify(s)
}

func (s CommitStagingRoutineCodeResponse) GoString() string {
	return s.String()
}

func (s *CommitStagingRoutineCodeResponse) SetHeaders(v map[string]*string) *CommitStagingRoutineCodeResponse {
	s.Headers = v
	return s
}

func (s *CommitStagingRoutineCodeResponse) SetBody(v *CommitStagingRoutineCodeResponseBody) *CommitStagingRoutineCodeResponse {
	s.Body = v
	return s
}

type CreateDcdnCertificateSigningRequestRequest struct {
	City             *string `json:"City,omitempty" xml:"City,omitempty"`
	CommonName       *string `json:"CommonName,omitempty" xml:"CommonName,omitempty"`
	Country          *string `json:"Country,omitempty" xml:"Country,omitempty"`
	Email            *string `json:"Email,omitempty" xml:"Email,omitempty"`
	Organization     *string `json:"Organization,omitempty" xml:"Organization,omitempty"`
	OrganizationUnit *string `json:"OrganizationUnit,omitempty" xml:"OrganizationUnit,omitempty"`
	OwnerId          *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SANs             *string `json:"SANs,omitempty" xml:"SANs,omitempty"`
	State            *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s CreateDcdnCertificateSigningRequestRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDcdnCertificateSigningRequestRequest) GoString() string {
	return s.String()
}

func (s *CreateDcdnCertificateSigningRequestRequest) SetCity(v string) *CreateDcdnCertificateSigningRequestRequest {
	s.City = &v
	return s
}

func (s *CreateDcdnCertificateSigningRequestRequest) SetCommonName(v string) *CreateDcdnCertificateSigningRequestRequest {
	s.CommonName = &v
	return s
}

func (s *CreateDcdnCertificateSigningRequestRequest) SetCountry(v string) *CreateDcdnCertificateSigningRequestRequest {
	s.Country = &v
	return s
}

func (s *CreateDcdnCertificateSigningRequestRequest) SetEmail(v string) *CreateDcdnCertificateSigningRequestRequest {
	s.Email = &v
	return s
}

func (s *CreateDcdnCertificateSigningRequestRequest) SetOrganization(v string) *CreateDcdnCertificateSigningRequestRequest {
	s.Organization = &v
	return s
}

func (s *CreateDcdnCertificateSigningRequestRequest) SetOrganizationUnit(v string) *CreateDcdnCertificateSigningRequestRequest {
	s.OrganizationUnit = &v
	return s
}

func (s *CreateDcdnCertificateSigningRequestRequest) SetOwnerId(v int64) *CreateDcdnCertificateSigningRequestRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateDcdnCertificateSigningRequestRequest) SetSANs(v string) *CreateDcdnCertificateSigningRequestRequest {
	s.SANs = &v
	return s
}

func (s *CreateDcdnCertificateSigningRequestRequest) SetState(v string) *CreateDcdnCertificateSigningRequestRequest {
	s.State = &v
	return s
}

type CreateDcdnCertificateSigningRequestResponseBody struct {
	CommonName *string `json:"CommonName,omitempty" xml:"CommonName,omitempty"`
	Csr        *string `json:"Csr,omitempty" xml:"Csr,omitempty"`
	PubMd5     *string `json:"PubMd5,omitempty" xml:"PubMd5,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDcdnCertificateSigningRequestResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDcdnCertificateSigningRequestResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDcdnCertificateSigningRequestResponseBody) SetCommonName(v string) *CreateDcdnCertificateSigningRequestResponseBody {
	s.CommonName = &v
	return s
}

func (s *CreateDcdnCertificateSigningRequestResponseBody) SetCsr(v string) *CreateDcdnCertificateSigningRequestResponseBody {
	s.Csr = &v
	return s
}

func (s *CreateDcdnCertificateSigningRequestResponseBody) SetPubMd5(v string) *CreateDcdnCertificateSigningRequestResponseBody {
	s.PubMd5 = &v
	return s
}

func (s *CreateDcdnCertificateSigningRequestResponseBody) SetRequestId(v string) *CreateDcdnCertificateSigningRequestResponseBody {
	s.RequestId = &v
	return s
}

type CreateDcdnCertificateSigningRequestResponse struct {
	Headers map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateDcdnCertificateSigningRequestResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateDcdnCertificateSigningRequestResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDcdnCertificateSigningRequestResponse) GoString() string {
	return s.String()
}

func (s *CreateDcdnCertificateSigningRequestResponse) SetHeaders(v map[string]*string) *CreateDcdnCertificateSigningRequestResponse {
	s.Headers = v
	return s
}

func (s *CreateDcdnCertificateSigningRequestResponse) SetBody(v *CreateDcdnCertificateSigningRequestResponseBody) *CreateDcdnCertificateSigningRequestResponse {
	s.Body = v
	return s
}

type CreateDcdnDeliverTaskRequest struct {
	Deliver    *string `json:"Deliver,omitempty" xml:"Deliver,omitempty"`
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	Reports    *string `json:"Reports,omitempty" xml:"Reports,omitempty"`
	Schedule   *string `json:"Schedule,omitempty" xml:"Schedule,omitempty"`
}

func (s CreateDcdnDeliverTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDcdnDeliverTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateDcdnDeliverTaskRequest) SetDeliver(v string) *CreateDcdnDeliverTaskRequest {
	s.Deliver = &v
	return s
}

func (s *CreateDcdnDeliverTaskRequest) SetDomainName(v string) *CreateDcdnDeliverTaskRequest {
	s.DomainName = &v
	return s
}

func (s *CreateDcdnDeliverTaskRequest) SetName(v string) *CreateDcdnDeliverTaskRequest {
	s.Name = &v
	return s
}

func (s *CreateDcdnDeliverTaskRequest) SetOwnerId(v int64) *CreateDcdnDeliverTaskRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateDcdnDeliverTaskRequest) SetReports(v string) *CreateDcdnDeliverTaskRequest {
	s.Reports = &v
	return s
}

func (s *CreateDcdnDeliverTaskRequest) SetSchedule(v string) *CreateDcdnDeliverTaskRequest {
	s.Schedule = &v
	return s
}

type CreateDcdnDeliverTaskResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDcdnDeliverTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDcdnDeliverTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDcdnDeliverTaskResponseBody) SetRequestId(v string) *CreateDcdnDeliverTaskResponseBody {
	s.RequestId = &v
	return s
}

type CreateDcdnDeliverTaskResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateDcdnDeliverTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateDcdnDeliverTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDcdnDeliverTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateDcdnDeliverTaskResponse) SetHeaders(v map[string]*string) *CreateDcdnDeliverTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateDcdnDeliverTaskResponse) SetBody(v *CreateDcdnDeliverTaskResponseBody) *CreateDcdnDeliverTaskResponse {
	s.Body = v
	return s
}

type CreateDcdnSLSRealTimeLogDeliveryRequest struct {
	BusinessType *string `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
	DataCenter   *string `json:"DataCenter,omitempty" xml:"DataCenter,omitempty"`
	DomainName   *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ProjectName  *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	SLSLogStore  *string `json:"SLSLogStore,omitempty" xml:"SLSLogStore,omitempty"`
	SLSProject   *string `json:"SLSProject,omitempty" xml:"SLSProject,omitempty"`
	SLSRegion    *string `json:"SLSRegion,omitempty" xml:"SLSRegion,omitempty"`
	SamplingRate *string `json:"SamplingRate,omitempty" xml:"SamplingRate,omitempty"`
}

func (s CreateDcdnSLSRealTimeLogDeliveryRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDcdnSLSRealTimeLogDeliveryRequest) GoString() string {
	return s.String()
}

func (s *CreateDcdnSLSRealTimeLogDeliveryRequest) SetBusinessType(v string) *CreateDcdnSLSRealTimeLogDeliveryRequest {
	s.BusinessType = &v
	return s
}

func (s *CreateDcdnSLSRealTimeLogDeliveryRequest) SetDataCenter(v string) *CreateDcdnSLSRealTimeLogDeliveryRequest {
	s.DataCenter = &v
	return s
}

func (s *CreateDcdnSLSRealTimeLogDeliveryRequest) SetDomainName(v string) *CreateDcdnSLSRealTimeLogDeliveryRequest {
	s.DomainName = &v
	return s
}

func (s *CreateDcdnSLSRealTimeLogDeliveryRequest) SetOwnerId(v int64) *CreateDcdnSLSRealTimeLogDeliveryRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateDcdnSLSRealTimeLogDeliveryRequest) SetProjectName(v string) *CreateDcdnSLSRealTimeLogDeliveryRequest {
	s.ProjectName = &v
	return s
}

func (s *CreateDcdnSLSRealTimeLogDeliveryRequest) SetSLSLogStore(v string) *CreateDcdnSLSRealTimeLogDeliveryRequest {
	s.SLSLogStore = &v
	return s
}

func (s *CreateDcdnSLSRealTimeLogDeliveryRequest) SetSLSProject(v string) *CreateDcdnSLSRealTimeLogDeliveryRequest {
	s.SLSProject = &v
	return s
}

func (s *CreateDcdnSLSRealTimeLogDeliveryRequest) SetSLSRegion(v string) *CreateDcdnSLSRealTimeLogDeliveryRequest {
	s.SLSRegion = &v
	return s
}

func (s *CreateDcdnSLSRealTimeLogDeliveryRequest) SetSamplingRate(v string) *CreateDcdnSLSRealTimeLogDeliveryRequest {
	s.SamplingRate = &v
	return s
}

type CreateDcdnSLSRealTimeLogDeliveryResponseBody struct {
	Content   *CreateDcdnSLSRealTimeLogDeliveryResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	RequestId *string                                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDcdnSLSRealTimeLogDeliveryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDcdnSLSRealTimeLogDeliveryResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDcdnSLSRealTimeLogDeliveryResponseBody) SetContent(v *CreateDcdnSLSRealTimeLogDeliveryResponseBodyContent) *CreateDcdnSLSRealTimeLogDeliveryResponseBody {
	s.Content = v
	return s
}

func (s *CreateDcdnSLSRealTimeLogDeliveryResponseBody) SetRequestId(v string) *CreateDcdnSLSRealTimeLogDeliveryResponseBody {
	s.RequestId = &v
	return s
}

type CreateDcdnSLSRealTimeLogDeliveryResponseBodyContent struct {
	Domains []*CreateDcdnSLSRealTimeLogDeliveryResponseBodyContentDomains `json:"Domains,omitempty" xml:"Domains,omitempty" type:"Repeated"`
}

func (s CreateDcdnSLSRealTimeLogDeliveryResponseBodyContent) String() string {
	return tea.Prettify(s)
}

func (s CreateDcdnSLSRealTimeLogDeliveryResponseBodyContent) GoString() string {
	return s.String()
}

func (s *CreateDcdnSLSRealTimeLogDeliveryResponseBodyContent) SetDomains(v []*CreateDcdnSLSRealTimeLogDeliveryResponseBodyContentDomains) *CreateDcdnSLSRealTimeLogDeliveryResponseBodyContent {
	s.Domains = v
	return s
}

type CreateDcdnSLSRealTimeLogDeliveryResponseBodyContentDomains struct {
	Desc       *string `json:"Desc,omitempty" xml:"Desc,omitempty"`
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	Region     *string `json:"Region,omitempty" xml:"Region,omitempty"`
	Status     *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s CreateDcdnSLSRealTimeLogDeliveryResponseBodyContentDomains) String() string {
	return tea.Prettify(s)
}

func (s CreateDcdnSLSRealTimeLogDeliveryResponseBodyContentDomains) GoString() string {
	return s.String()
}

func (s *CreateDcdnSLSRealTimeLogDeliveryResponseBodyContentDomains) SetDesc(v string) *CreateDcdnSLSRealTimeLogDeliveryResponseBodyContentDomains {
	s.Desc = &v
	return s
}

func (s *CreateDcdnSLSRealTimeLogDeliveryResponseBodyContentDomains) SetDomainName(v string) *CreateDcdnSLSRealTimeLogDeliveryResponseBodyContentDomains {
	s.DomainName = &v
	return s
}

func (s *CreateDcdnSLSRealTimeLogDeliveryResponseBodyContentDomains) SetRegion(v string) *CreateDcdnSLSRealTimeLogDeliveryResponseBodyContentDomains {
	s.Region = &v
	return s
}

func (s *CreateDcdnSLSRealTimeLogDeliveryResponseBodyContentDomains) SetStatus(v string) *CreateDcdnSLSRealTimeLogDeliveryResponseBodyContentDomains {
	s.Status = &v
	return s
}

type CreateDcdnSLSRealTimeLogDeliveryResponse struct {
	Headers map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateDcdnSLSRealTimeLogDeliveryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateDcdnSLSRealTimeLogDeliveryResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDcdnSLSRealTimeLogDeliveryResponse) GoString() string {
	return s.String()
}

func (s *CreateDcdnSLSRealTimeLogDeliveryResponse) SetHeaders(v map[string]*string) *CreateDcdnSLSRealTimeLogDeliveryResponse {
	s.Headers = v
	return s
}

func (s *CreateDcdnSLSRealTimeLogDeliveryResponse) SetBody(v *CreateDcdnSLSRealTimeLogDeliveryResponseBody) *CreateDcdnSLSRealTimeLogDeliveryResponse {
	s.Body = v
	return s
}

type CreateDcdnSubTaskRequest struct {
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ReportIds  *string `json:"ReportIds,omitempty" xml:"ReportIds,omitempty"`
}

func (s CreateDcdnSubTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDcdnSubTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateDcdnSubTaskRequest) SetDomainName(v string) *CreateDcdnSubTaskRequest {
	s.DomainName = &v
	return s
}

func (s *CreateDcdnSubTaskRequest) SetOwnerId(v int64) *CreateDcdnSubTaskRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateDcdnSubTaskRequest) SetReportIds(v string) *CreateDcdnSubTaskRequest {
	s.ReportIds = &v
	return s
}

type CreateDcdnSubTaskResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDcdnSubTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDcdnSubTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDcdnSubTaskResponseBody) SetRequestId(v string) *CreateDcdnSubTaskResponseBody {
	s.RequestId = &v
	return s
}

type CreateDcdnSubTaskResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateDcdnSubTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateDcdnSubTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDcdnSubTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateDcdnSubTaskResponse) SetHeaders(v map[string]*string) *CreateDcdnSubTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateDcdnSubTaskResponse) SetBody(v *CreateDcdnSubTaskResponseBody) *CreateDcdnSubTaskResponse {
	s.Body = v
	return s
}

type CreateRoutineRequest struct {
	Description *string                `json:"Description,omitempty" xml:"Description,omitempty"`
	EnvConf     map[string]interface{} `json:"EnvConf,omitempty" xml:"EnvConf,omitempty"`
	Name        *string                `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerId     *int64                 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s CreateRoutineRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateRoutineRequest) GoString() string {
	return s.String()
}

func (s *CreateRoutineRequest) SetDescription(v string) *CreateRoutineRequest {
	s.Description = &v
	return s
}

func (s *CreateRoutineRequest) SetEnvConf(v map[string]interface{}) *CreateRoutineRequest {
	s.EnvConf = v
	return s
}

func (s *CreateRoutineRequest) SetName(v string) *CreateRoutineRequest {
	s.Name = &v
	return s
}

func (s *CreateRoutineRequest) SetOwnerId(v int64) *CreateRoutineRequest {
	s.OwnerId = &v
	return s
}

type CreateRoutineShrinkRequest struct {
	Description   *string `json:"Description,omitempty" xml:"Description,omitempty"`
	EnvConfShrink *string `json:"EnvConf,omitempty" xml:"EnvConf,omitempty"`
	Name          *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s CreateRoutineShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateRoutineShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateRoutineShrinkRequest) SetDescription(v string) *CreateRoutineShrinkRequest {
	s.Description = &v
	return s
}

func (s *CreateRoutineShrinkRequest) SetEnvConfShrink(v string) *CreateRoutineShrinkRequest {
	s.EnvConfShrink = &v
	return s
}

func (s *CreateRoutineShrinkRequest) SetName(v string) *CreateRoutineShrinkRequest {
	s.Name = &v
	return s
}

func (s *CreateRoutineShrinkRequest) SetOwnerId(v int64) *CreateRoutineShrinkRequest {
	s.OwnerId = &v
	return s
}

type CreateRoutineResponseBody struct {
	Content   map[string]interface{} `json:"Content,omitempty" xml:"Content,omitempty"`
	RequestId *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateRoutineResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateRoutineResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRoutineResponseBody) SetContent(v map[string]interface{}) *CreateRoutineResponseBody {
	s.Content = v
	return s
}

func (s *CreateRoutineResponseBody) SetRequestId(v string) *CreateRoutineResponseBody {
	s.RequestId = &v
	return s
}

type CreateRoutineResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateRoutineResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateRoutineResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateRoutineResponse) GoString() string {
	return s.String()
}

func (s *CreateRoutineResponse) SetHeaders(v map[string]*string) *CreateRoutineResponse {
	s.Headers = v
	return s
}

func (s *CreateRoutineResponse) SetBody(v *CreateRoutineResponseBody) *CreateRoutineResponse {
	s.Body = v
	return s
}

type CreateSlrAndSlsProjectRequest struct {
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	Region  *string `json:"Region,omitempty" xml:"Region,omitempty"`
}

func (s CreateSlrAndSlsProjectRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateSlrAndSlsProjectRequest) GoString() string {
	return s.String()
}

func (s *CreateSlrAndSlsProjectRequest) SetOwnerId(v int64) *CreateSlrAndSlsProjectRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateSlrAndSlsProjectRequest) SetRegion(v string) *CreateSlrAndSlsProjectRequest {
	s.Region = &v
	return s
}

type CreateSlrAndSlsProjectResponseBody struct {
	RequestId *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SlsInfo   *CreateSlrAndSlsProjectResponseBodySlsInfo `json:"SlsInfo,omitempty" xml:"SlsInfo,omitempty" type:"Struct"`
}

func (s CreateSlrAndSlsProjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateSlrAndSlsProjectResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSlrAndSlsProjectResponseBody) SetRequestId(v string) *CreateSlrAndSlsProjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSlrAndSlsProjectResponseBody) SetSlsInfo(v *CreateSlrAndSlsProjectResponseBodySlsInfo) *CreateSlrAndSlsProjectResponseBody {
	s.SlsInfo = v
	return s
}

type CreateSlrAndSlsProjectResponseBodySlsInfo struct {
	EndPoint *string `json:"EndPoint,omitempty" xml:"EndPoint,omitempty"`
	LogStore *string `json:"LogStore,omitempty" xml:"LogStore,omitempty"`
	Project  *string `json:"Project,omitempty" xml:"Project,omitempty"`
	Region   *string `json:"Region,omitempty" xml:"Region,omitempty"`
}

func (s CreateSlrAndSlsProjectResponseBodySlsInfo) String() string {
	return tea.Prettify(s)
}

func (s CreateSlrAndSlsProjectResponseBodySlsInfo) GoString() string {
	return s.String()
}

func (s *CreateSlrAndSlsProjectResponseBodySlsInfo) SetEndPoint(v string) *CreateSlrAndSlsProjectResponseBodySlsInfo {
	s.EndPoint = &v
	return s
}

func (s *CreateSlrAndSlsProjectResponseBodySlsInfo) SetLogStore(v string) *CreateSlrAndSlsProjectResponseBodySlsInfo {
	s.LogStore = &v
	return s
}

func (s *CreateSlrAndSlsProjectResponseBodySlsInfo) SetProject(v string) *CreateSlrAndSlsProjectResponseBodySlsInfo {
	s.Project = &v
	return s
}

func (s *CreateSlrAndSlsProjectResponseBodySlsInfo) SetRegion(v string) *CreateSlrAndSlsProjectResponseBodySlsInfo {
	s.Region = &v
	return s
}

type CreateSlrAndSlsProjectResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateSlrAndSlsProjectResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateSlrAndSlsProjectResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateSlrAndSlsProjectResponse) GoString() string {
	return s.String()
}

func (s *CreateSlrAndSlsProjectResponse) SetHeaders(v map[string]*string) *CreateSlrAndSlsProjectResponse {
	s.Headers = v
	return s
}

func (s *CreateSlrAndSlsProjectResponse) SetBody(v *CreateSlrAndSlsProjectResponseBody) *CreateSlrAndSlsProjectResponse {
	s.Body = v
	return s
}

type DcdnHttpRequestTestToolRequest struct {
	Args    *string                `json:"Args,omitempty" xml:"Args,omitempty"`
	Body    *string                `json:"Body,omitempty" xml:"Body,omitempty"`
	Header  map[string]interface{} `json:"Header,omitempty" xml:"Header,omitempty"`
	Host    *string                `json:"Host,omitempty" xml:"Host,omitempty"`
	Method  *string                `json:"Method,omitempty" xml:"Method,omitempty"`
	OwnerId *int64                 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ProxyIp *string                `json:"ProxyIp,omitempty" xml:"ProxyIp,omitempty"`
	Scheme  *string                `json:"Scheme,omitempty" xml:"Scheme,omitempty"`
	Uri     *string                `json:"Uri,omitempty" xml:"Uri,omitempty"`
}

func (s DcdnHttpRequestTestToolRequest) String() string {
	return tea.Prettify(s)
}

func (s DcdnHttpRequestTestToolRequest) GoString() string {
	return s.String()
}

func (s *DcdnHttpRequestTestToolRequest) SetArgs(v string) *DcdnHttpRequestTestToolRequest {
	s.Args = &v
	return s
}

func (s *DcdnHttpRequestTestToolRequest) SetBody(v string) *DcdnHttpRequestTestToolRequest {
	s.Body = &v
	return s
}

func (s *DcdnHttpRequestTestToolRequest) SetHeader(v map[string]interface{}) *DcdnHttpRequestTestToolRequest {
	s.Header = v
	return s
}

func (s *DcdnHttpRequestTestToolRequest) SetHost(v string) *DcdnHttpRequestTestToolRequest {
	s.Host = &v
	return s
}

func (s *DcdnHttpRequestTestToolRequest) SetMethod(v string) *DcdnHttpRequestTestToolRequest {
	s.Method = &v
	return s
}

func (s *DcdnHttpRequestTestToolRequest) SetOwnerId(v int64) *DcdnHttpRequestTestToolRequest {
	s.OwnerId = &v
	return s
}

func (s *DcdnHttpRequestTestToolRequest) SetProxyIp(v string) *DcdnHttpRequestTestToolRequest {
	s.ProxyIp = &v
	return s
}

func (s *DcdnHttpRequestTestToolRequest) SetScheme(v string) *DcdnHttpRequestTestToolRequest {
	s.Scheme = &v
	return s
}

func (s *DcdnHttpRequestTestToolRequest) SetUri(v string) *DcdnHttpRequestTestToolRequest {
	s.Uri = &v
	return s
}

type DcdnHttpRequestTestToolShrinkRequest struct {
	Args         *string `json:"Args,omitempty" xml:"Args,omitempty"`
	Body         *string `json:"Body,omitempty" xml:"Body,omitempty"`
	HeaderShrink *string `json:"Header,omitempty" xml:"Header,omitempty"`
	Host         *string `json:"Host,omitempty" xml:"Host,omitempty"`
	Method       *string `json:"Method,omitempty" xml:"Method,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ProxyIp      *string `json:"ProxyIp,omitempty" xml:"ProxyIp,omitempty"`
	Scheme       *string `json:"Scheme,omitempty" xml:"Scheme,omitempty"`
	Uri          *string `json:"Uri,omitempty" xml:"Uri,omitempty"`
}

func (s DcdnHttpRequestTestToolShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s DcdnHttpRequestTestToolShrinkRequest) GoString() string {
	return s.String()
}

func (s *DcdnHttpRequestTestToolShrinkRequest) SetArgs(v string) *DcdnHttpRequestTestToolShrinkRequest {
	s.Args = &v
	return s
}

func (s *DcdnHttpRequestTestToolShrinkRequest) SetBody(v string) *DcdnHttpRequestTestToolShrinkRequest {
	s.Body = &v
	return s
}

func (s *DcdnHttpRequestTestToolShrinkRequest) SetHeaderShrink(v string) *DcdnHttpRequestTestToolShrinkRequest {
	s.HeaderShrink = &v
	return s
}

func (s *DcdnHttpRequestTestToolShrinkRequest) SetHost(v string) *DcdnHttpRequestTestToolShrinkRequest {
	s.Host = &v
	return s
}

func (s *DcdnHttpRequestTestToolShrinkRequest) SetMethod(v string) *DcdnHttpRequestTestToolShrinkRequest {
	s.Method = &v
	return s
}

func (s *DcdnHttpRequestTestToolShrinkRequest) SetOwnerId(v int64) *DcdnHttpRequestTestToolShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *DcdnHttpRequestTestToolShrinkRequest) SetProxyIp(v string) *DcdnHttpRequestTestToolShrinkRequest {
	s.ProxyIp = &v
	return s
}

func (s *DcdnHttpRequestTestToolShrinkRequest) SetScheme(v string) *DcdnHttpRequestTestToolShrinkRequest {
	s.Scheme = &v
	return s
}

func (s *DcdnHttpRequestTestToolShrinkRequest) SetUri(v string) *DcdnHttpRequestTestToolShrinkRequest {
	s.Uri = &v
	return s
}

type DcdnHttpRequestTestToolResponseBody struct {
	Body       *string `json:"Body,omitempty" xml:"Body,omitempty"`
	Header     *string `json:"Header,omitempty" xml:"Header,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StatusCode *int32  `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
}

func (s DcdnHttpRequestTestToolResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DcdnHttpRequestTestToolResponseBody) GoString() string {
	return s.String()
}

func (s *DcdnHttpRequestTestToolResponseBody) SetBody(v string) *DcdnHttpRequestTestToolResponseBody {
	s.Body = &v
	return s
}

func (s *DcdnHttpRequestTestToolResponseBody) SetHeader(v string) *DcdnHttpRequestTestToolResponseBody {
	s.Header = &v
	return s
}

func (s *DcdnHttpRequestTestToolResponseBody) SetRequestId(v string) *DcdnHttpRequestTestToolResponseBody {
	s.RequestId = &v
	return s
}

func (s *DcdnHttpRequestTestToolResponseBody) SetStatusCode(v int32) *DcdnHttpRequestTestToolResponseBody {
	s.StatusCode = &v
	return s
}

type DcdnHttpRequestTestToolResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DcdnHttpRequestTestToolResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DcdnHttpRequestTestToolResponse) String() string {
	return tea.Prettify(s)
}

func (s DcdnHttpRequestTestToolResponse) GoString() string {
	return s.String()
}

func (s *DcdnHttpRequestTestToolResponse) SetHeaders(v map[string]*string) *DcdnHttpRequestTestToolResponse {
	s.Headers = v
	return s
}

func (s *DcdnHttpRequestTestToolResponse) SetBody(v *DcdnHttpRequestTestToolResponseBody) *DcdnHttpRequestTestToolResponse {
	s.Body = v
	return s
}

type DeleteDcdnDeliverTaskRequest struct {
	DeliverId *int64 `json:"DeliverId,omitempty" xml:"DeliverId,omitempty"`
	OwnerId   *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s DeleteDcdnDeliverTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDcdnDeliverTaskRequest) GoString() string {
	return s.String()
}

func (s *DeleteDcdnDeliverTaskRequest) SetDeliverId(v int64) *DeleteDcdnDeliverTaskRequest {
	s.DeliverId = &v
	return s
}

func (s *DeleteDcdnDeliverTaskRequest) SetOwnerId(v int64) *DeleteDcdnDeliverTaskRequest {
	s.OwnerId = &v
	return s
}

type DeleteDcdnDeliverTaskResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDcdnDeliverTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDcdnDeliverTaskResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDcdnDeliverTaskResponseBody) SetRequestId(v string) *DeleteDcdnDeliverTaskResponseBody {
	s.RequestId = &v
	return s
}

type DeleteDcdnDeliverTaskResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteDcdnDeliverTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteDcdnDeliverTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDcdnDeliverTaskResponse) GoString() string {
	return s.String()
}

func (s *DeleteDcdnDeliverTaskResponse) SetHeaders(v map[string]*string) *DeleteDcdnDeliverTaskResponse {
	s.Headers = v
	return s
}

func (s *DeleteDcdnDeliverTaskResponse) SetBody(v *DeleteDcdnDeliverTaskResponseBody) *DeleteDcdnDeliverTaskResponse {
	s.Body = v
	return s
}

type DeleteDcdnDomainRequest struct {
	DomainName    *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	OwnerAccount  *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DeleteDcdnDomainRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDcdnDomainRequest) GoString() string {
	return s.String()
}

func (s *DeleteDcdnDomainRequest) SetDomainName(v string) *DeleteDcdnDomainRequest {
	s.DomainName = &v
	return s
}

func (s *DeleteDcdnDomainRequest) SetOwnerAccount(v string) *DeleteDcdnDomainRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteDcdnDomainRequest) SetOwnerId(v int64) *DeleteDcdnDomainRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteDcdnDomainRequest) SetSecurityToken(v string) *DeleteDcdnDomainRequest {
	s.SecurityToken = &v
	return s
}

type DeleteDcdnDomainResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDcdnDomainResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDcdnDomainResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDcdnDomainResponseBody) SetRequestId(v string) *DeleteDcdnDomainResponseBody {
	s.RequestId = &v
	return s
}

type DeleteDcdnDomainResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteDcdnDomainResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteDcdnDomainResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDcdnDomainResponse) GoString() string {
	return s.String()
}

func (s *DeleteDcdnDomainResponse) SetHeaders(v map[string]*string) *DeleteDcdnDomainResponse {
	s.Headers = v
	return s
}

func (s *DeleteDcdnDomainResponse) SetBody(v *DeleteDcdnDomainResponseBody) *DeleteDcdnDomainResponse {
	s.Body = v
	return s
}

type DeleteDcdnIpaDomainRequest struct {
	DomainName    *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	OwnerAccount  *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DeleteDcdnIpaDomainRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDcdnIpaDomainRequest) GoString() string {
	return s.String()
}

func (s *DeleteDcdnIpaDomainRequest) SetDomainName(v string) *DeleteDcdnIpaDomainRequest {
	s.DomainName = &v
	return s
}

func (s *DeleteDcdnIpaDomainRequest) SetOwnerAccount(v string) *DeleteDcdnIpaDomainRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteDcdnIpaDomainRequest) SetOwnerId(v int64) *DeleteDcdnIpaDomainRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteDcdnIpaDomainRequest) SetSecurityToken(v string) *DeleteDcdnIpaDomainRequest {
	s.SecurityToken = &v
	return s
}

type DeleteDcdnIpaDomainResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDcdnIpaDomainResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDcdnIpaDomainResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDcdnIpaDomainResponseBody) SetRequestId(v string) *DeleteDcdnIpaDomainResponseBody {
	s.RequestId = &v
	return s
}

type DeleteDcdnIpaDomainResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteDcdnIpaDomainResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteDcdnIpaDomainResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDcdnIpaDomainResponse) GoString() string {
	return s.String()
}

func (s *DeleteDcdnIpaDomainResponse) SetHeaders(v map[string]*string) *DeleteDcdnIpaDomainResponse {
	s.Headers = v
	return s
}

func (s *DeleteDcdnIpaDomainResponse) SetBody(v *DeleteDcdnIpaDomainResponseBody) *DeleteDcdnIpaDomainResponse {
	s.Body = v
	return s
}

type DeleteDcdnIpaSpecificConfigRequest struct {
	ConfigId      *string `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	DomainName    *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DeleteDcdnIpaSpecificConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDcdnIpaSpecificConfigRequest) GoString() string {
	return s.String()
}

func (s *DeleteDcdnIpaSpecificConfigRequest) SetConfigId(v string) *DeleteDcdnIpaSpecificConfigRequest {
	s.ConfigId = &v
	return s
}

func (s *DeleteDcdnIpaSpecificConfigRequest) SetDomainName(v string) *DeleteDcdnIpaSpecificConfigRequest {
	s.DomainName = &v
	return s
}

func (s *DeleteDcdnIpaSpecificConfigRequest) SetOwnerId(v int64) *DeleteDcdnIpaSpecificConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteDcdnIpaSpecificConfigRequest) SetSecurityToken(v string) *DeleteDcdnIpaSpecificConfigRequest {
	s.SecurityToken = &v
	return s
}

type DeleteDcdnIpaSpecificConfigResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDcdnIpaSpecificConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDcdnIpaSpecificConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDcdnIpaSpecificConfigResponseBody) SetRequestId(v string) *DeleteDcdnIpaSpecificConfigResponseBody {
	s.RequestId = &v
	return s
}

type DeleteDcdnIpaSpecificConfigResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteDcdnIpaSpecificConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteDcdnIpaSpecificConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDcdnIpaSpecificConfigResponse) GoString() string {
	return s.String()
}

func (s *DeleteDcdnIpaSpecificConfigResponse) SetHeaders(v map[string]*string) *DeleteDcdnIpaSpecificConfigResponse {
	s.Headers = v
	return s
}

func (s *DeleteDcdnIpaSpecificConfigResponse) SetBody(v *DeleteDcdnIpaSpecificConfigResponseBody) *DeleteDcdnIpaSpecificConfigResponse {
	s.Body = v
	return s
}

type DeleteDcdnRealTimeLogProjectRequest struct {
	BusinessType *string `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ProjectName  *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
}

func (s DeleteDcdnRealTimeLogProjectRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDcdnRealTimeLogProjectRequest) GoString() string {
	return s.String()
}

func (s *DeleteDcdnRealTimeLogProjectRequest) SetBusinessType(v string) *DeleteDcdnRealTimeLogProjectRequest {
	s.BusinessType = &v
	return s
}

func (s *DeleteDcdnRealTimeLogProjectRequest) SetOwnerId(v int64) *DeleteDcdnRealTimeLogProjectRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteDcdnRealTimeLogProjectRequest) SetProjectName(v string) *DeleteDcdnRealTimeLogProjectRequest {
	s.ProjectName = &v
	return s
}

type DeleteDcdnRealTimeLogProjectResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDcdnRealTimeLogProjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDcdnRealTimeLogProjectResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDcdnRealTimeLogProjectResponseBody) SetRequestId(v string) *DeleteDcdnRealTimeLogProjectResponseBody {
	s.RequestId = &v
	return s
}

type DeleteDcdnRealTimeLogProjectResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteDcdnRealTimeLogProjectResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteDcdnRealTimeLogProjectResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDcdnRealTimeLogProjectResponse) GoString() string {
	return s.String()
}

func (s *DeleteDcdnRealTimeLogProjectResponse) SetHeaders(v map[string]*string) *DeleteDcdnRealTimeLogProjectResponse {
	s.Headers = v
	return s
}

func (s *DeleteDcdnRealTimeLogProjectResponse) SetBody(v *DeleteDcdnRealTimeLogProjectResponseBody) *DeleteDcdnRealTimeLogProjectResponse {
	s.Body = v
	return s
}

type DeleteDcdnSpecificConfigRequest struct {
	ConfigId      *string `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	DomainName    *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DeleteDcdnSpecificConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDcdnSpecificConfigRequest) GoString() string {
	return s.String()
}

func (s *DeleteDcdnSpecificConfigRequest) SetConfigId(v string) *DeleteDcdnSpecificConfigRequest {
	s.ConfigId = &v
	return s
}

func (s *DeleteDcdnSpecificConfigRequest) SetDomainName(v string) *DeleteDcdnSpecificConfigRequest {
	s.DomainName = &v
	return s
}

func (s *DeleteDcdnSpecificConfigRequest) SetOwnerId(v int64) *DeleteDcdnSpecificConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteDcdnSpecificConfigRequest) SetSecurityToken(v string) *DeleteDcdnSpecificConfigRequest {
	s.SecurityToken = &v
	return s
}

type DeleteDcdnSpecificConfigResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDcdnSpecificConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDcdnSpecificConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDcdnSpecificConfigResponseBody) SetRequestId(v string) *DeleteDcdnSpecificConfigResponseBody {
	s.RequestId = &v
	return s
}

type DeleteDcdnSpecificConfigResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteDcdnSpecificConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteDcdnSpecificConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDcdnSpecificConfigResponse) GoString() string {
	return s.String()
}

func (s *DeleteDcdnSpecificConfigResponse) SetHeaders(v map[string]*string) *DeleteDcdnSpecificConfigResponse {
	s.Headers = v
	return s
}

func (s *DeleteDcdnSpecificConfigResponse) SetBody(v *DeleteDcdnSpecificConfigResponseBody) *DeleteDcdnSpecificConfigResponse {
	s.Body = v
	return s
}

type DeleteDcdnSpecificStagingConfigRequest struct {
	ConfigId      *string `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	DomainName    *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DeleteDcdnSpecificStagingConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDcdnSpecificStagingConfigRequest) GoString() string {
	return s.String()
}

func (s *DeleteDcdnSpecificStagingConfigRequest) SetConfigId(v string) *DeleteDcdnSpecificStagingConfigRequest {
	s.ConfigId = &v
	return s
}

func (s *DeleteDcdnSpecificStagingConfigRequest) SetDomainName(v string) *DeleteDcdnSpecificStagingConfigRequest {
	s.DomainName = &v
	return s
}

func (s *DeleteDcdnSpecificStagingConfigRequest) SetOwnerId(v int64) *DeleteDcdnSpecificStagingConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteDcdnSpecificStagingConfigRequest) SetSecurityToken(v string) *DeleteDcdnSpecificStagingConfigRequest {
	s.SecurityToken = &v
	return s
}

type DeleteDcdnSpecificStagingConfigResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDcdnSpecificStagingConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDcdnSpecificStagingConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDcdnSpecificStagingConfigResponseBody) SetRequestId(v string) *DeleteDcdnSpecificStagingConfigResponseBody {
	s.RequestId = &v
	return s
}

type DeleteDcdnSpecificStagingConfigResponse struct {
	Headers map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteDcdnSpecificStagingConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteDcdnSpecificStagingConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDcdnSpecificStagingConfigResponse) GoString() string {
	return s.String()
}

func (s *DeleteDcdnSpecificStagingConfigResponse) SetHeaders(v map[string]*string) *DeleteDcdnSpecificStagingConfigResponse {
	s.Headers = v
	return s
}

func (s *DeleteDcdnSpecificStagingConfigResponse) SetBody(v *DeleteDcdnSpecificStagingConfigResponseBody) *DeleteDcdnSpecificStagingConfigResponse {
	s.Body = v
	return s
}

type DeleteDcdnSubTaskRequest struct {
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s DeleteDcdnSubTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDcdnSubTaskRequest) GoString() string {
	return s.String()
}

func (s *DeleteDcdnSubTaskRequest) SetOwnerId(v int64) *DeleteDcdnSubTaskRequest {
	s.OwnerId = &v
	return s
}

type DeleteDcdnSubTaskResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDcdnSubTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDcdnSubTaskResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDcdnSubTaskResponseBody) SetRequestId(v string) *DeleteDcdnSubTaskResponseBody {
	s.RequestId = &v
	return s
}

type DeleteDcdnSubTaskResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteDcdnSubTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteDcdnSubTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDcdnSubTaskResponse) GoString() string {
	return s.String()
}

func (s *DeleteDcdnSubTaskResponse) SetHeaders(v map[string]*string) *DeleteDcdnSubTaskResponse {
	s.Headers = v
	return s
}

func (s *DeleteDcdnSubTaskResponse) SetBody(v *DeleteDcdnSubTaskResponseBody) *DeleteDcdnSubTaskResponse {
	s.Body = v
	return s
}

type DeleteRoutineRequest struct {
	Name    *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s DeleteRoutineRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteRoutineRequest) GoString() string {
	return s.String()
}

func (s *DeleteRoutineRequest) SetName(v string) *DeleteRoutineRequest {
	s.Name = &v
	return s
}

func (s *DeleteRoutineRequest) SetOwnerId(v int64) *DeleteRoutineRequest {
	s.OwnerId = &v
	return s
}

type DeleteRoutineResponseBody struct {
	Content   map[string]interface{} `json:"Content,omitempty" xml:"Content,omitempty"`
	RequestId *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteRoutineResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteRoutineResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteRoutineResponseBody) SetContent(v map[string]interface{}) *DeleteRoutineResponseBody {
	s.Content = v
	return s
}

func (s *DeleteRoutineResponseBody) SetRequestId(v string) *DeleteRoutineResponseBody {
	s.RequestId = &v
	return s
}

type DeleteRoutineResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteRoutineResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteRoutineResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteRoutineResponse) GoString() string {
	return s.String()
}

func (s *DeleteRoutineResponse) SetHeaders(v map[string]*string) *DeleteRoutineResponse {
	s.Headers = v
	return s
}

func (s *DeleteRoutineResponse) SetBody(v *DeleteRoutineResponseBody) *DeleteRoutineResponse {
	s.Body = v
	return s
}

type DeleteRoutineCodeRevisionRequest struct {
	Name               *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerId            *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SelectCodeRevision *string `json:"SelectCodeRevision,omitempty" xml:"SelectCodeRevision,omitempty"`
}

func (s DeleteRoutineCodeRevisionRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteRoutineCodeRevisionRequest) GoString() string {
	return s.String()
}

func (s *DeleteRoutineCodeRevisionRequest) SetName(v string) *DeleteRoutineCodeRevisionRequest {
	s.Name = &v
	return s
}

func (s *DeleteRoutineCodeRevisionRequest) SetOwnerId(v int64) *DeleteRoutineCodeRevisionRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteRoutineCodeRevisionRequest) SetSelectCodeRevision(v string) *DeleteRoutineCodeRevisionRequest {
	s.SelectCodeRevision = &v
	return s
}

type DeleteRoutineCodeRevisionResponseBody struct {
	Content   map[string]interface{} `json:"Content,omitempty" xml:"Content,omitempty"`
	RequestId *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteRoutineCodeRevisionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteRoutineCodeRevisionResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteRoutineCodeRevisionResponseBody) SetContent(v map[string]interface{}) *DeleteRoutineCodeRevisionResponseBody {
	s.Content = v
	return s
}

func (s *DeleteRoutineCodeRevisionResponseBody) SetRequestId(v string) *DeleteRoutineCodeRevisionResponseBody {
	s.RequestId = &v
	return s
}

type DeleteRoutineCodeRevisionResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteRoutineCodeRevisionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteRoutineCodeRevisionResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteRoutineCodeRevisionResponse) GoString() string {
	return s.String()
}

func (s *DeleteRoutineCodeRevisionResponse) SetHeaders(v map[string]*string) *DeleteRoutineCodeRevisionResponse {
	s.Headers = v
	return s
}

func (s *DeleteRoutineCodeRevisionResponse) SetBody(v *DeleteRoutineCodeRevisionResponseBody) *DeleteRoutineCodeRevisionResponse {
	s.Body = v
	return s
}

type DeleteRoutineConfEnvsRequest struct {
	Envs    map[string]interface{} `json:"Envs,omitempty" xml:"Envs,omitempty"`
	Name    *string                `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerId *int64                 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s DeleteRoutineConfEnvsRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteRoutineConfEnvsRequest) GoString() string {
	return s.String()
}

func (s *DeleteRoutineConfEnvsRequest) SetEnvs(v map[string]interface{}) *DeleteRoutineConfEnvsRequest {
	s.Envs = v
	return s
}

func (s *DeleteRoutineConfEnvsRequest) SetName(v string) *DeleteRoutineConfEnvsRequest {
	s.Name = &v
	return s
}

func (s *DeleteRoutineConfEnvsRequest) SetOwnerId(v int64) *DeleteRoutineConfEnvsRequest {
	s.OwnerId = &v
	return s
}

type DeleteRoutineConfEnvsShrinkRequest struct {
	EnvsShrink *string `json:"Envs,omitempty" xml:"Envs,omitempty"`
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s DeleteRoutineConfEnvsShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteRoutineConfEnvsShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeleteRoutineConfEnvsShrinkRequest) SetEnvsShrink(v string) *DeleteRoutineConfEnvsShrinkRequest {
	s.EnvsShrink = &v
	return s
}

func (s *DeleteRoutineConfEnvsShrinkRequest) SetName(v string) *DeleteRoutineConfEnvsShrinkRequest {
	s.Name = &v
	return s
}

func (s *DeleteRoutineConfEnvsShrinkRequest) SetOwnerId(v int64) *DeleteRoutineConfEnvsShrinkRequest {
	s.OwnerId = &v
	return s
}

type DeleteRoutineConfEnvsResponseBody struct {
	Content   map[string]interface{} `json:"Content,omitempty" xml:"Content,omitempty"`
	RequestId *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteRoutineConfEnvsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteRoutineConfEnvsResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteRoutineConfEnvsResponseBody) SetContent(v map[string]interface{}) *DeleteRoutineConfEnvsResponseBody {
	s.Content = v
	return s
}

func (s *DeleteRoutineConfEnvsResponseBody) SetRequestId(v string) *DeleteRoutineConfEnvsResponseBody {
	s.RequestId = &v
	return s
}

type DeleteRoutineConfEnvsResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteRoutineConfEnvsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteRoutineConfEnvsResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteRoutineConfEnvsResponse) GoString() string {
	return s.String()
}

func (s *DeleteRoutineConfEnvsResponse) SetHeaders(v map[string]*string) *DeleteRoutineConfEnvsResponse {
	s.Headers = v
	return s
}

func (s *DeleteRoutineConfEnvsResponse) SetBody(v *DeleteRoutineConfEnvsResponseBody) *DeleteRoutineConfEnvsResponse {
	s.Body = v
	return s
}

type DescribeDcdnAclFieldsRequest struct {
	Lang    *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s DescribeDcdnAclFieldsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnAclFieldsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDcdnAclFieldsRequest) SetLang(v string) *DescribeDcdnAclFieldsRequest {
	s.Lang = &v
	return s
}

func (s *DescribeDcdnAclFieldsRequest) SetOwnerId(v int64) *DescribeDcdnAclFieldsRequest {
	s.OwnerId = &v
	return s
}

type DescribeDcdnAclFieldsResponseBody struct {
	Content   []*DescribeDcdnAclFieldsResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Repeated"`
	RequestId *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDcdnAclFieldsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnAclFieldsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnAclFieldsResponseBody) SetContent(v []*DescribeDcdnAclFieldsResponseBodyContent) *DescribeDcdnAclFieldsResponseBody {
	s.Content = v
	return s
}

func (s *DescribeDcdnAclFieldsResponseBody) SetRequestId(v string) *DescribeDcdnAclFieldsResponseBody {
	s.RequestId = &v
	return s
}

type DescribeDcdnAclFieldsResponseBodyContent struct {
	Fields *string `json:"Fields,omitempty" xml:"Fields,omitempty"`
}

func (s DescribeDcdnAclFieldsResponseBodyContent) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnAclFieldsResponseBodyContent) GoString() string {
	return s.String()
}

func (s *DescribeDcdnAclFieldsResponseBodyContent) SetFields(v string) *DescribeDcdnAclFieldsResponseBodyContent {
	s.Fields = &v
	return s
}

type DescribeDcdnAclFieldsResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDcdnAclFieldsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDcdnAclFieldsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnAclFieldsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnAclFieldsResponse) SetHeaders(v map[string]*string) *DescribeDcdnAclFieldsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnAclFieldsResponse) SetBody(v *DescribeDcdnAclFieldsResponseBody) *DescribeDcdnAclFieldsResponse {
	s.Body = v
	return s
}

type DescribeDcdnBgpBpsDataRequest struct {
	EndTime   *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Interval  *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	Isp       *string `json:"Isp,omitempty" xml:"Isp,omitempty"`
	OwnerId   *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDcdnBgpBpsDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnBgpBpsDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeDcdnBgpBpsDataRequest) SetEndTime(v string) *DescribeDcdnBgpBpsDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDcdnBgpBpsDataRequest) SetInterval(v string) *DescribeDcdnBgpBpsDataRequest {
	s.Interval = &v
	return s
}

func (s *DescribeDcdnBgpBpsDataRequest) SetIsp(v string) *DescribeDcdnBgpBpsDataRequest {
	s.Isp = &v
	return s
}

func (s *DescribeDcdnBgpBpsDataRequest) SetOwnerId(v int64) *DescribeDcdnBgpBpsDataRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDcdnBgpBpsDataRequest) SetStartTime(v string) *DescribeDcdnBgpBpsDataRequest {
	s.StartTime = &v
	return s
}

type DescribeDcdnBgpBpsDataResponseBody struct {
	BgpDataInterval []*DescribeDcdnBgpBpsDataResponseBodyBgpDataInterval `json:"BgpDataInterval,omitempty" xml:"BgpDataInterval,omitempty" type:"Repeated"`
	EndTime         *string                                              `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	RequestId       *string                                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StartTime       *string                                              `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDcdnBgpBpsDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnBgpBpsDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnBgpBpsDataResponseBody) SetBgpDataInterval(v []*DescribeDcdnBgpBpsDataResponseBodyBgpDataInterval) *DescribeDcdnBgpBpsDataResponseBody {
	s.BgpDataInterval = v
	return s
}

func (s *DescribeDcdnBgpBpsDataResponseBody) SetEndTime(v string) *DescribeDcdnBgpBpsDataResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeDcdnBgpBpsDataResponseBody) SetRequestId(v string) *DescribeDcdnBgpBpsDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDcdnBgpBpsDataResponseBody) SetStartTime(v string) *DescribeDcdnBgpBpsDataResponseBody {
	s.StartTime = &v
	return s
}

type DescribeDcdnBgpBpsDataResponseBodyBgpDataInterval struct {
	In        *float32 `json:"In,omitempty" xml:"In,omitempty"`
	Out       *float32 `json:"Out,omitempty" xml:"Out,omitempty"`
	TimeStamp *string  `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
}

func (s DescribeDcdnBgpBpsDataResponseBodyBgpDataInterval) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnBgpBpsDataResponseBodyBgpDataInterval) GoString() string {
	return s.String()
}

func (s *DescribeDcdnBgpBpsDataResponseBodyBgpDataInterval) SetIn(v float32) *DescribeDcdnBgpBpsDataResponseBodyBgpDataInterval {
	s.In = &v
	return s
}

func (s *DescribeDcdnBgpBpsDataResponseBodyBgpDataInterval) SetOut(v float32) *DescribeDcdnBgpBpsDataResponseBodyBgpDataInterval {
	s.Out = &v
	return s
}

func (s *DescribeDcdnBgpBpsDataResponseBodyBgpDataInterval) SetTimeStamp(v string) *DescribeDcdnBgpBpsDataResponseBodyBgpDataInterval {
	s.TimeStamp = &v
	return s
}

type DescribeDcdnBgpBpsDataResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDcdnBgpBpsDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDcdnBgpBpsDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnBgpBpsDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnBgpBpsDataResponse) SetHeaders(v map[string]*string) *DescribeDcdnBgpBpsDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnBgpBpsDataResponse) SetBody(v *DescribeDcdnBgpBpsDataResponseBody) *DescribeDcdnBgpBpsDataResponse {
	s.Body = v
	return s
}

type DescribeDcdnBgpTrafficDataRequest struct {
	EndTime   *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Interval  *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	Isp       *string `json:"Isp,omitempty" xml:"Isp,omitempty"`
	OwnerId   *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDcdnBgpTrafficDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnBgpTrafficDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeDcdnBgpTrafficDataRequest) SetEndTime(v string) *DescribeDcdnBgpTrafficDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDcdnBgpTrafficDataRequest) SetInterval(v string) *DescribeDcdnBgpTrafficDataRequest {
	s.Interval = &v
	return s
}

func (s *DescribeDcdnBgpTrafficDataRequest) SetIsp(v string) *DescribeDcdnBgpTrafficDataRequest {
	s.Isp = &v
	return s
}

func (s *DescribeDcdnBgpTrafficDataRequest) SetOwnerId(v int64) *DescribeDcdnBgpTrafficDataRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDcdnBgpTrafficDataRequest) SetStartTime(v string) *DescribeDcdnBgpTrafficDataRequest {
	s.StartTime = &v
	return s
}

type DescribeDcdnBgpTrafficDataResponseBody struct {
	BgpDataInterval []*DescribeDcdnBgpTrafficDataResponseBodyBgpDataInterval `json:"BgpDataInterval,omitempty" xml:"BgpDataInterval,omitempty" type:"Repeated"`
	EndTime         *string                                                  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	RequestId       *string                                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StartTime       *string                                                  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDcdnBgpTrafficDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnBgpTrafficDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnBgpTrafficDataResponseBody) SetBgpDataInterval(v []*DescribeDcdnBgpTrafficDataResponseBodyBgpDataInterval) *DescribeDcdnBgpTrafficDataResponseBody {
	s.BgpDataInterval = v
	return s
}

func (s *DescribeDcdnBgpTrafficDataResponseBody) SetEndTime(v string) *DescribeDcdnBgpTrafficDataResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeDcdnBgpTrafficDataResponseBody) SetRequestId(v string) *DescribeDcdnBgpTrafficDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDcdnBgpTrafficDataResponseBody) SetStartTime(v string) *DescribeDcdnBgpTrafficDataResponseBody {
	s.StartTime = &v
	return s
}

type DescribeDcdnBgpTrafficDataResponseBodyBgpDataInterval struct {
	In        *int64  `json:"In,omitempty" xml:"In,omitempty"`
	Out       *int64  `json:"Out,omitempty" xml:"Out,omitempty"`
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
}

func (s DescribeDcdnBgpTrafficDataResponseBodyBgpDataInterval) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnBgpTrafficDataResponseBodyBgpDataInterval) GoString() string {
	return s.String()
}

func (s *DescribeDcdnBgpTrafficDataResponseBodyBgpDataInterval) SetIn(v int64) *DescribeDcdnBgpTrafficDataResponseBodyBgpDataInterval {
	s.In = &v
	return s
}

func (s *DescribeDcdnBgpTrafficDataResponseBodyBgpDataInterval) SetOut(v int64) *DescribeDcdnBgpTrafficDataResponseBodyBgpDataInterval {
	s.Out = &v
	return s
}

func (s *DescribeDcdnBgpTrafficDataResponseBodyBgpDataInterval) SetTimeStamp(v string) *DescribeDcdnBgpTrafficDataResponseBodyBgpDataInterval {
	s.TimeStamp = &v
	return s
}

type DescribeDcdnBgpTrafficDataResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDcdnBgpTrafficDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDcdnBgpTrafficDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnBgpTrafficDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnBgpTrafficDataResponse) SetHeaders(v map[string]*string) *DescribeDcdnBgpTrafficDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnBgpTrafficDataResponse) SetBody(v *DescribeDcdnBgpTrafficDataResponseBody) *DescribeDcdnBgpTrafficDataResponse {
	s.Body = v
	return s
}

type DescribeDcdnBlockedRegionsRequest struct {
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s DescribeDcdnBlockedRegionsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnBlockedRegionsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDcdnBlockedRegionsRequest) SetLanguage(v string) *DescribeDcdnBlockedRegionsRequest {
	s.Language = &v
	return s
}

func (s *DescribeDcdnBlockedRegionsRequest) SetOwnerId(v int64) *DescribeDcdnBlockedRegionsRequest {
	s.OwnerId = &v
	return s
}

type DescribeDcdnBlockedRegionsResponseBody struct {
	InfoList  *DescribeDcdnBlockedRegionsResponseBodyInfoList `json:"InfoList,omitempty" xml:"InfoList,omitempty" type:"Struct"`
	RequestId *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDcdnBlockedRegionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnBlockedRegionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnBlockedRegionsResponseBody) SetInfoList(v *DescribeDcdnBlockedRegionsResponseBodyInfoList) *DescribeDcdnBlockedRegionsResponseBody {
	s.InfoList = v
	return s
}

func (s *DescribeDcdnBlockedRegionsResponseBody) SetRequestId(v string) *DescribeDcdnBlockedRegionsResponseBody {
	s.RequestId = &v
	return s
}

type DescribeDcdnBlockedRegionsResponseBodyInfoList struct {
	InfoItem []*DescribeDcdnBlockedRegionsResponseBodyInfoListInfoItem `json:"InfoItem,omitempty" xml:"InfoItem,omitempty" type:"Repeated"`
}

func (s DescribeDcdnBlockedRegionsResponseBodyInfoList) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnBlockedRegionsResponseBodyInfoList) GoString() string {
	return s.String()
}

func (s *DescribeDcdnBlockedRegionsResponseBodyInfoList) SetInfoItem(v []*DescribeDcdnBlockedRegionsResponseBodyInfoListInfoItem) *DescribeDcdnBlockedRegionsResponseBodyInfoList {
	s.InfoItem = v
	return s
}

type DescribeDcdnBlockedRegionsResponseBodyInfoListInfoItem struct {
	Continent               *string `json:"Continent,omitempty" xml:"Continent,omitempty"`
	CountriesAndRegions     *string `json:"CountriesAndRegions,omitempty" xml:"CountriesAndRegions,omitempty"`
	CountriesAndRegionsName *string `json:"CountriesAndRegionsName,omitempty" xml:"CountriesAndRegionsName,omitempty"`
}

func (s DescribeDcdnBlockedRegionsResponseBodyInfoListInfoItem) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnBlockedRegionsResponseBodyInfoListInfoItem) GoString() string {
	return s.String()
}

func (s *DescribeDcdnBlockedRegionsResponseBodyInfoListInfoItem) SetContinent(v string) *DescribeDcdnBlockedRegionsResponseBodyInfoListInfoItem {
	s.Continent = &v
	return s
}

func (s *DescribeDcdnBlockedRegionsResponseBodyInfoListInfoItem) SetCountriesAndRegions(v string) *DescribeDcdnBlockedRegionsResponseBodyInfoListInfoItem {
	s.CountriesAndRegions = &v
	return s
}

func (s *DescribeDcdnBlockedRegionsResponseBodyInfoListInfoItem) SetCountriesAndRegionsName(v string) *DescribeDcdnBlockedRegionsResponseBodyInfoListInfoItem {
	s.CountriesAndRegionsName = &v
	return s
}

type DescribeDcdnBlockedRegionsResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDcdnBlockedRegionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDcdnBlockedRegionsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnBlockedRegionsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnBlockedRegionsResponse) SetHeaders(v map[string]*string) *DescribeDcdnBlockedRegionsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnBlockedRegionsResponse) SetBody(v *DescribeDcdnBlockedRegionsResponseBody) *DescribeDcdnBlockedRegionsResponse {
	s.Body = v
	return s
}

type DescribeDcdnCertificateDetailRequest struct {
	CertName      *string `json:"CertName,omitempty" xml:"CertName,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeDcdnCertificateDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnCertificateDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribeDcdnCertificateDetailRequest) SetCertName(v string) *DescribeDcdnCertificateDetailRequest {
	s.CertName = &v
	return s
}

func (s *DescribeDcdnCertificateDetailRequest) SetOwnerId(v int64) *DescribeDcdnCertificateDetailRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDcdnCertificateDetailRequest) SetSecurityToken(v string) *DescribeDcdnCertificateDetailRequest {
	s.SecurityToken = &v
	return s
}

type DescribeDcdnCertificateDetailResponseBody struct {
	Cert      *string `json:"Cert,omitempty" xml:"Cert,omitempty"`
	CertId    *int64  `json:"CertId,omitempty" xml:"CertId,omitempty"`
	CertName  *string `json:"CertName,omitempty" xml:"CertName,omitempty"`
	Key       *string `json:"Key,omitempty" xml:"Key,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDcdnCertificateDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnCertificateDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnCertificateDetailResponseBody) SetCert(v string) *DescribeDcdnCertificateDetailResponseBody {
	s.Cert = &v
	return s
}

func (s *DescribeDcdnCertificateDetailResponseBody) SetCertId(v int64) *DescribeDcdnCertificateDetailResponseBody {
	s.CertId = &v
	return s
}

func (s *DescribeDcdnCertificateDetailResponseBody) SetCertName(v string) *DescribeDcdnCertificateDetailResponseBody {
	s.CertName = &v
	return s
}

func (s *DescribeDcdnCertificateDetailResponseBody) SetKey(v string) *DescribeDcdnCertificateDetailResponseBody {
	s.Key = &v
	return s
}

func (s *DescribeDcdnCertificateDetailResponseBody) SetRequestId(v string) *DescribeDcdnCertificateDetailResponseBody {
	s.RequestId = &v
	return s
}

type DescribeDcdnCertificateDetailResponse struct {
	Headers map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDcdnCertificateDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDcdnCertificateDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnCertificateDetailResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnCertificateDetailResponse) SetHeaders(v map[string]*string) *DescribeDcdnCertificateDetailResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnCertificateDetailResponse) SetBody(v *DescribeDcdnCertificateDetailResponseBody) *DescribeDcdnCertificateDetailResponse {
	s.Body = v
	return s
}

type DescribeDcdnCertificateListRequest struct {
	DomainName    *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeDcdnCertificateListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnCertificateListRequest) GoString() string {
	return s.String()
}

func (s *DescribeDcdnCertificateListRequest) SetDomainName(v string) *DescribeDcdnCertificateListRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDcdnCertificateListRequest) SetOwnerId(v int64) *DescribeDcdnCertificateListRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDcdnCertificateListRequest) SetSecurityToken(v string) *DescribeDcdnCertificateListRequest {
	s.SecurityToken = &v
	return s
}

type DescribeDcdnCertificateListResponseBody struct {
	CertificateListModel *DescribeDcdnCertificateListResponseBodyCertificateListModel `json:"CertificateListModel,omitempty" xml:"CertificateListModel,omitempty" type:"Struct"`
	RequestId            *string                                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDcdnCertificateListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnCertificateListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnCertificateListResponseBody) SetCertificateListModel(v *DescribeDcdnCertificateListResponseBodyCertificateListModel) *DescribeDcdnCertificateListResponseBody {
	s.CertificateListModel = v
	return s
}

func (s *DescribeDcdnCertificateListResponseBody) SetRequestId(v string) *DescribeDcdnCertificateListResponseBody {
	s.RequestId = &v
	return s
}

type DescribeDcdnCertificateListResponseBodyCertificateListModel struct {
	CertList *DescribeDcdnCertificateListResponseBodyCertificateListModelCertList `json:"CertList,omitempty" xml:"CertList,omitempty" type:"Struct"`
	Count    *int32                                                               `json:"Count,omitempty" xml:"Count,omitempty"`
}

func (s DescribeDcdnCertificateListResponseBodyCertificateListModel) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnCertificateListResponseBodyCertificateListModel) GoString() string {
	return s.String()
}

func (s *DescribeDcdnCertificateListResponseBodyCertificateListModel) SetCertList(v *DescribeDcdnCertificateListResponseBodyCertificateListModelCertList) *DescribeDcdnCertificateListResponseBodyCertificateListModel {
	s.CertList = v
	return s
}

func (s *DescribeDcdnCertificateListResponseBodyCertificateListModel) SetCount(v int32) *DescribeDcdnCertificateListResponseBodyCertificateListModel {
	s.Count = &v
	return s
}

type DescribeDcdnCertificateListResponseBodyCertificateListModelCertList struct {
	Cert []*DescribeDcdnCertificateListResponseBodyCertificateListModelCertListCert `json:"Cert,omitempty" xml:"Cert,omitempty" type:"Repeated"`
}

func (s DescribeDcdnCertificateListResponseBodyCertificateListModelCertList) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnCertificateListResponseBodyCertificateListModelCertList) GoString() string {
	return s.String()
}

func (s *DescribeDcdnCertificateListResponseBodyCertificateListModelCertList) SetCert(v []*DescribeDcdnCertificateListResponseBodyCertificateListModelCertListCert) *DescribeDcdnCertificateListResponseBodyCertificateListModelCertList {
	s.Cert = v
	return s
}

type DescribeDcdnCertificateListResponseBodyCertificateListModelCertListCert struct {
	CertId      *int64  `json:"CertId,omitempty" xml:"CertId,omitempty"`
	CertName    *string `json:"CertName,omitempty" xml:"CertName,omitempty"`
	Common      *string `json:"Common,omitempty" xml:"Common,omitempty"`
	Fingerprint *string `json:"Fingerprint,omitempty" xml:"Fingerprint,omitempty"`
	Issuer      *string `json:"Issuer,omitempty" xml:"Issuer,omitempty"`
	LastTime    *int64  `json:"LastTime,omitempty" xml:"LastTime,omitempty"`
}

func (s DescribeDcdnCertificateListResponseBodyCertificateListModelCertListCert) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnCertificateListResponseBodyCertificateListModelCertListCert) GoString() string {
	return s.String()
}

func (s *DescribeDcdnCertificateListResponseBodyCertificateListModelCertListCert) SetCertId(v int64) *DescribeDcdnCertificateListResponseBodyCertificateListModelCertListCert {
	s.CertId = &v
	return s
}

func (s *DescribeDcdnCertificateListResponseBodyCertificateListModelCertListCert) SetCertName(v string) *DescribeDcdnCertificateListResponseBodyCertificateListModelCertListCert {
	s.CertName = &v
	return s
}

func (s *DescribeDcdnCertificateListResponseBodyCertificateListModelCertListCert) SetCommon(v string) *DescribeDcdnCertificateListResponseBodyCertificateListModelCertListCert {
	s.Common = &v
	return s
}

func (s *DescribeDcdnCertificateListResponseBodyCertificateListModelCertListCert) SetFingerprint(v string) *DescribeDcdnCertificateListResponseBodyCertificateListModelCertListCert {
	s.Fingerprint = &v
	return s
}

func (s *DescribeDcdnCertificateListResponseBodyCertificateListModelCertListCert) SetIssuer(v string) *DescribeDcdnCertificateListResponseBodyCertificateListModelCertListCert {
	s.Issuer = &v
	return s
}

func (s *DescribeDcdnCertificateListResponseBodyCertificateListModelCertListCert) SetLastTime(v int64) *DescribeDcdnCertificateListResponseBodyCertificateListModelCertListCert {
	s.LastTime = &v
	return s
}

type DescribeDcdnCertificateListResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDcdnCertificateListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDcdnCertificateListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnCertificateListResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnCertificateListResponse) SetHeaders(v map[string]*string) *DescribeDcdnCertificateListResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnCertificateListResponse) SetBody(v *DescribeDcdnCertificateListResponseBody) *DescribeDcdnCertificateListResponse {
	s.Body = v
	return s
}

type DescribeDcdnConfigGroupDetailRequest struct {
	ConfigGroupId   *string `json:"ConfigGroupId,omitempty" xml:"ConfigGroupId,omitempty"`
	ConfigGroupName *string `json:"ConfigGroupName,omitempty" xml:"ConfigGroupName,omitempty"`
	OwnerId         *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s DescribeDcdnConfigGroupDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnConfigGroupDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribeDcdnConfigGroupDetailRequest) SetConfigGroupId(v string) *DescribeDcdnConfigGroupDetailRequest {
	s.ConfigGroupId = &v
	return s
}

func (s *DescribeDcdnConfigGroupDetailRequest) SetConfigGroupName(v string) *DescribeDcdnConfigGroupDetailRequest {
	s.ConfigGroupName = &v
	return s
}

func (s *DescribeDcdnConfigGroupDetailRequest) SetOwnerId(v int64) *DescribeDcdnConfigGroupDetailRequest {
	s.OwnerId = &v
	return s
}

type DescribeDcdnConfigGroupDetailResponseBody struct {
	BizName         *string `json:"BizName,omitempty" xml:"BizName,omitempty"`
	ConfigGroupId   *string `json:"ConfigGroupId,omitempty" xml:"ConfigGroupId,omitempty"`
	ConfigGroupName *string `json:"ConfigGroupName,omitempty" xml:"ConfigGroupName,omitempty"`
	CreateTime      *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description     *string `json:"Description,omitempty" xml:"Description,omitempty"`
	RequestId       *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	UpdateTime      *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s DescribeDcdnConfigGroupDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnConfigGroupDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnConfigGroupDetailResponseBody) SetBizName(v string) *DescribeDcdnConfigGroupDetailResponseBody {
	s.BizName = &v
	return s
}

func (s *DescribeDcdnConfigGroupDetailResponseBody) SetConfigGroupId(v string) *DescribeDcdnConfigGroupDetailResponseBody {
	s.ConfigGroupId = &v
	return s
}

func (s *DescribeDcdnConfigGroupDetailResponseBody) SetConfigGroupName(v string) *DescribeDcdnConfigGroupDetailResponseBody {
	s.ConfigGroupName = &v
	return s
}

func (s *DescribeDcdnConfigGroupDetailResponseBody) SetCreateTime(v string) *DescribeDcdnConfigGroupDetailResponseBody {
	s.CreateTime = &v
	return s
}

func (s *DescribeDcdnConfigGroupDetailResponseBody) SetDescription(v string) *DescribeDcdnConfigGroupDetailResponseBody {
	s.Description = &v
	return s
}

func (s *DescribeDcdnConfigGroupDetailResponseBody) SetRequestId(v string) *DescribeDcdnConfigGroupDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDcdnConfigGroupDetailResponseBody) SetUpdateTime(v string) *DescribeDcdnConfigGroupDetailResponseBody {
	s.UpdateTime = &v
	return s
}

type DescribeDcdnConfigGroupDetailResponse struct {
	Headers map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDcdnConfigGroupDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDcdnConfigGroupDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnConfigGroupDetailResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnConfigGroupDetailResponse) SetHeaders(v map[string]*string) *DescribeDcdnConfigGroupDetailResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnConfigGroupDetailResponse) SetBody(v *DescribeDcdnConfigGroupDetailResponseBody) *DescribeDcdnConfigGroupDetailResponse {
	s.Body = v
	return s
}

type DescribeDcdnConfigOfVersionRequest struct {
	FunctionId    *int32  `json:"FunctionId,omitempty" xml:"FunctionId,omitempty"`
	FunctionName  *string `json:"FunctionName,omitempty" xml:"FunctionName,omitempty"`
	GroupId       *int64  `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	VersionId     *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
}

func (s DescribeDcdnConfigOfVersionRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnConfigOfVersionRequest) GoString() string {
	return s.String()
}

func (s *DescribeDcdnConfigOfVersionRequest) SetFunctionId(v int32) *DescribeDcdnConfigOfVersionRequest {
	s.FunctionId = &v
	return s
}

func (s *DescribeDcdnConfigOfVersionRequest) SetFunctionName(v string) *DescribeDcdnConfigOfVersionRequest {
	s.FunctionName = &v
	return s
}

func (s *DescribeDcdnConfigOfVersionRequest) SetGroupId(v int64) *DescribeDcdnConfigOfVersionRequest {
	s.GroupId = &v
	return s
}

func (s *DescribeDcdnConfigOfVersionRequest) SetOwnerId(v int64) *DescribeDcdnConfigOfVersionRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDcdnConfigOfVersionRequest) SetSecurityToken(v string) *DescribeDcdnConfigOfVersionRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeDcdnConfigOfVersionRequest) SetVersionId(v string) *DescribeDcdnConfigOfVersionRequest {
	s.VersionId = &v
	return s
}

type DescribeDcdnConfigOfVersionResponseBody struct {
	RequestId      *string                                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	VersionConfigs *DescribeDcdnConfigOfVersionResponseBodyVersionConfigs `json:"VersionConfigs,omitempty" xml:"VersionConfigs,omitempty" type:"Struct"`
}

func (s DescribeDcdnConfigOfVersionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnConfigOfVersionResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnConfigOfVersionResponseBody) SetRequestId(v string) *DescribeDcdnConfigOfVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDcdnConfigOfVersionResponseBody) SetVersionConfigs(v *DescribeDcdnConfigOfVersionResponseBodyVersionConfigs) *DescribeDcdnConfigOfVersionResponseBody {
	s.VersionConfigs = v
	return s
}

type DescribeDcdnConfigOfVersionResponseBodyVersionConfigs struct {
	VersionConfig []*DescribeDcdnConfigOfVersionResponseBodyVersionConfigsVersionConfig `json:"VersionConfig,omitempty" xml:"VersionConfig,omitempty" type:"Repeated"`
}

func (s DescribeDcdnConfigOfVersionResponseBodyVersionConfigs) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnConfigOfVersionResponseBodyVersionConfigs) GoString() string {
	return s.String()
}

func (s *DescribeDcdnConfigOfVersionResponseBodyVersionConfigs) SetVersionConfig(v []*DescribeDcdnConfigOfVersionResponseBodyVersionConfigsVersionConfig) *DescribeDcdnConfigOfVersionResponseBodyVersionConfigs {
	s.VersionConfig = v
	return s
}

type DescribeDcdnConfigOfVersionResponseBodyVersionConfigsVersionConfig struct {
	ConfigId     *string                                                                         `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	FunctionArgs *DescribeDcdnConfigOfVersionResponseBodyVersionConfigsVersionConfigFunctionArgs `json:"FunctionArgs,omitempty" xml:"FunctionArgs,omitempty" type:"Struct"`
	FunctionName *string                                                                         `json:"FunctionName,omitempty" xml:"FunctionName,omitempty"`
	Status       *string                                                                         `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeDcdnConfigOfVersionResponseBodyVersionConfigsVersionConfig) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnConfigOfVersionResponseBodyVersionConfigsVersionConfig) GoString() string {
	return s.String()
}

func (s *DescribeDcdnConfigOfVersionResponseBodyVersionConfigsVersionConfig) SetConfigId(v string) *DescribeDcdnConfigOfVersionResponseBodyVersionConfigsVersionConfig {
	s.ConfigId = &v
	return s
}

func (s *DescribeDcdnConfigOfVersionResponseBodyVersionConfigsVersionConfig) SetFunctionArgs(v *DescribeDcdnConfigOfVersionResponseBodyVersionConfigsVersionConfigFunctionArgs) *DescribeDcdnConfigOfVersionResponseBodyVersionConfigsVersionConfig {
	s.FunctionArgs = v
	return s
}

func (s *DescribeDcdnConfigOfVersionResponseBodyVersionConfigsVersionConfig) SetFunctionName(v string) *DescribeDcdnConfigOfVersionResponseBodyVersionConfigsVersionConfig {
	s.FunctionName = &v
	return s
}

func (s *DescribeDcdnConfigOfVersionResponseBodyVersionConfigsVersionConfig) SetStatus(v string) *DescribeDcdnConfigOfVersionResponseBodyVersionConfigsVersionConfig {
	s.Status = &v
	return s
}

type DescribeDcdnConfigOfVersionResponseBodyVersionConfigsVersionConfigFunctionArgs struct {
	FunctionArg []*DescribeDcdnConfigOfVersionResponseBodyVersionConfigsVersionConfigFunctionArgsFunctionArg `json:"FunctionArg,omitempty" xml:"FunctionArg,omitempty" type:"Repeated"`
}

func (s DescribeDcdnConfigOfVersionResponseBodyVersionConfigsVersionConfigFunctionArgs) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnConfigOfVersionResponseBodyVersionConfigsVersionConfigFunctionArgs) GoString() string {
	return s.String()
}

func (s *DescribeDcdnConfigOfVersionResponseBodyVersionConfigsVersionConfigFunctionArgs) SetFunctionArg(v []*DescribeDcdnConfigOfVersionResponseBodyVersionConfigsVersionConfigFunctionArgsFunctionArg) *DescribeDcdnConfigOfVersionResponseBodyVersionConfigsVersionConfigFunctionArgs {
	s.FunctionArg = v
	return s
}

type DescribeDcdnConfigOfVersionResponseBodyVersionConfigsVersionConfigFunctionArgsFunctionArg struct {
	ArgName  *string `json:"ArgName,omitempty" xml:"ArgName,omitempty"`
	ArgValue *string `json:"ArgValue,omitempty" xml:"ArgValue,omitempty"`
}

func (s DescribeDcdnConfigOfVersionResponseBodyVersionConfigsVersionConfigFunctionArgsFunctionArg) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnConfigOfVersionResponseBodyVersionConfigsVersionConfigFunctionArgsFunctionArg) GoString() string {
	return s.String()
}

func (s *DescribeDcdnConfigOfVersionResponseBodyVersionConfigsVersionConfigFunctionArgsFunctionArg) SetArgName(v string) *DescribeDcdnConfigOfVersionResponseBodyVersionConfigsVersionConfigFunctionArgsFunctionArg {
	s.ArgName = &v
	return s
}

func (s *DescribeDcdnConfigOfVersionResponseBodyVersionConfigsVersionConfigFunctionArgsFunctionArg) SetArgValue(v string) *DescribeDcdnConfigOfVersionResponseBodyVersionConfigsVersionConfigFunctionArgsFunctionArg {
	s.ArgValue = &v
	return s
}

type DescribeDcdnConfigOfVersionResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDcdnConfigOfVersionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDcdnConfigOfVersionResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnConfigOfVersionResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnConfigOfVersionResponse) SetHeaders(v map[string]*string) *DescribeDcdnConfigOfVersionResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnConfigOfVersionResponse) SetBody(v *DescribeDcdnConfigOfVersionResponseBody) *DescribeDcdnConfigOfVersionResponse {
	s.Body = v
	return s
}

type DescribeDcdnDeletedDomainsRequest struct {
	OwnerId    *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeDcdnDeletedDomainsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDeletedDomainsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDeletedDomainsRequest) SetOwnerId(v int64) *DescribeDcdnDeletedDomainsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDcdnDeletedDomainsRequest) SetPageNumber(v int32) *DescribeDcdnDeletedDomainsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDcdnDeletedDomainsRequest) SetPageSize(v int32) *DescribeDcdnDeletedDomainsRequest {
	s.PageSize = &v
	return s
}

type DescribeDcdnDeletedDomainsResponseBody struct {
	Domains    *DescribeDcdnDeletedDomainsResponseBodyDomains `json:"Domains,omitempty" xml:"Domains,omitempty" type:"Struct"`
	PageNumber *int64                                         `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int64                                         `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int64                                         `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDcdnDeletedDomainsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDeletedDomainsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDeletedDomainsResponseBody) SetDomains(v *DescribeDcdnDeletedDomainsResponseBodyDomains) *DescribeDcdnDeletedDomainsResponseBody {
	s.Domains = v
	return s
}

func (s *DescribeDcdnDeletedDomainsResponseBody) SetPageNumber(v int64) *DescribeDcdnDeletedDomainsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeDcdnDeletedDomainsResponseBody) SetPageSize(v int64) *DescribeDcdnDeletedDomainsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeDcdnDeletedDomainsResponseBody) SetRequestId(v string) *DescribeDcdnDeletedDomainsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDcdnDeletedDomainsResponseBody) SetTotalCount(v int64) *DescribeDcdnDeletedDomainsResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeDcdnDeletedDomainsResponseBodyDomains struct {
	PageData []*DescribeDcdnDeletedDomainsResponseBodyDomainsPageData `json:"PageData,omitempty" xml:"PageData,omitempty" type:"Repeated"`
}

func (s DescribeDcdnDeletedDomainsResponseBodyDomains) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDeletedDomainsResponseBodyDomains) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDeletedDomainsResponseBodyDomains) SetPageData(v []*DescribeDcdnDeletedDomainsResponseBodyDomainsPageData) *DescribeDcdnDeletedDomainsResponseBodyDomains {
	s.PageData = v
	return s
}

type DescribeDcdnDeletedDomainsResponseBodyDomainsPageData struct {
	DomainName  *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
}

func (s DescribeDcdnDeletedDomainsResponseBodyDomainsPageData) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDeletedDomainsResponseBodyDomainsPageData) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDeletedDomainsResponseBodyDomainsPageData) SetDomainName(v string) *DescribeDcdnDeletedDomainsResponseBodyDomainsPageData {
	s.DomainName = &v
	return s
}

func (s *DescribeDcdnDeletedDomainsResponseBodyDomainsPageData) SetGmtModified(v string) *DescribeDcdnDeletedDomainsResponseBodyDomainsPageData {
	s.GmtModified = &v
	return s
}

type DescribeDcdnDeletedDomainsResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDcdnDeletedDomainsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDcdnDeletedDomainsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDeletedDomainsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDeletedDomainsResponse) SetHeaders(v map[string]*string) *DescribeDcdnDeletedDomainsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnDeletedDomainsResponse) SetBody(v *DescribeDcdnDeletedDomainsResponseBody) *DescribeDcdnDeletedDomainsResponse {
	s.Body = v
	return s
}

type DescribeDcdnDeliverListRequest struct {
	DeliverId *int64 `json:"DeliverId,omitempty" xml:"DeliverId,omitempty"`
	OwnerId   *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s DescribeDcdnDeliverListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDeliverListRequest) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDeliverListRequest) SetDeliverId(v int64) *DescribeDcdnDeliverListRequest {
	s.DeliverId = &v
	return s
}

func (s *DescribeDcdnDeliverListRequest) SetOwnerId(v int64) *DescribeDcdnDeliverListRequest {
	s.OwnerId = &v
	return s
}

type DescribeDcdnDeliverListResponseBody struct {
	Content   *string `json:"Content,omitempty" xml:"Content,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDcdnDeliverListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDeliverListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDeliverListResponseBody) SetContent(v string) *DescribeDcdnDeliverListResponseBody {
	s.Content = &v
	return s
}

func (s *DescribeDcdnDeliverListResponseBody) SetRequestId(v string) *DescribeDcdnDeliverListResponseBody {
	s.RequestId = &v
	return s
}

type DescribeDcdnDeliverListResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDcdnDeliverListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDcdnDeliverListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDeliverListResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDeliverListResponse) SetHeaders(v map[string]*string) *DescribeDcdnDeliverListResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnDeliverListResponse) SetBody(v *DescribeDcdnDeliverListResponseBody) *DescribeDcdnDeliverListResponse {
	s.Body = v
	return s
}

type DescribeDcdnDomainBpsDataRequest struct {
	DomainName     *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EndTime        *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Interval       *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	IspNameEn      *string `json:"IspNameEn,omitempty" xml:"IspNameEn,omitempty"`
	LocationNameEn *string `json:"LocationNameEn,omitempty" xml:"LocationNameEn,omitempty"`
	OwnerId        *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	StartTime      *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDcdnDomainBpsDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainBpsDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainBpsDataRequest) SetDomainName(v string) *DescribeDcdnDomainBpsDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDcdnDomainBpsDataRequest) SetEndTime(v string) *DescribeDcdnDomainBpsDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDcdnDomainBpsDataRequest) SetInterval(v string) *DescribeDcdnDomainBpsDataRequest {
	s.Interval = &v
	return s
}

func (s *DescribeDcdnDomainBpsDataRequest) SetIspNameEn(v string) *DescribeDcdnDomainBpsDataRequest {
	s.IspNameEn = &v
	return s
}

func (s *DescribeDcdnDomainBpsDataRequest) SetLocationNameEn(v string) *DescribeDcdnDomainBpsDataRequest {
	s.LocationNameEn = &v
	return s
}

func (s *DescribeDcdnDomainBpsDataRequest) SetOwnerId(v int64) *DescribeDcdnDomainBpsDataRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDcdnDomainBpsDataRequest) SetStartTime(v string) *DescribeDcdnDomainBpsDataRequest {
	s.StartTime = &v
	return s
}

type DescribeDcdnDomainBpsDataResponseBody struct {
	BpsDataPerInterval *DescribeDcdnDomainBpsDataResponseBodyBpsDataPerInterval `json:"BpsDataPerInterval,omitempty" xml:"BpsDataPerInterval,omitempty" type:"Struct"`
	DataInterval       *string                                                  `json:"DataInterval,omitempty" xml:"DataInterval,omitempty"`
	DomainName         *string                                                  `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EndTime            *string                                                  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	RequestId          *string                                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StartTime          *string                                                  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDcdnDomainBpsDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainBpsDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainBpsDataResponseBody) SetBpsDataPerInterval(v *DescribeDcdnDomainBpsDataResponseBodyBpsDataPerInterval) *DescribeDcdnDomainBpsDataResponseBody {
	s.BpsDataPerInterval = v
	return s
}

func (s *DescribeDcdnDomainBpsDataResponseBody) SetDataInterval(v string) *DescribeDcdnDomainBpsDataResponseBody {
	s.DataInterval = &v
	return s
}

func (s *DescribeDcdnDomainBpsDataResponseBody) SetDomainName(v string) *DescribeDcdnDomainBpsDataResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeDcdnDomainBpsDataResponseBody) SetEndTime(v string) *DescribeDcdnDomainBpsDataResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeDcdnDomainBpsDataResponseBody) SetRequestId(v string) *DescribeDcdnDomainBpsDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDcdnDomainBpsDataResponseBody) SetStartTime(v string) *DescribeDcdnDomainBpsDataResponseBody {
	s.StartTime = &v
	return s
}

type DescribeDcdnDomainBpsDataResponseBodyBpsDataPerInterval struct {
	DataModule []*DescribeDcdnDomainBpsDataResponseBodyBpsDataPerIntervalDataModule `json:"DataModule,omitempty" xml:"DataModule,omitempty" type:"Repeated"`
}

func (s DescribeDcdnDomainBpsDataResponseBodyBpsDataPerInterval) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainBpsDataResponseBodyBpsDataPerInterval) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainBpsDataResponseBodyBpsDataPerInterval) SetDataModule(v []*DescribeDcdnDomainBpsDataResponseBodyBpsDataPerIntervalDataModule) *DescribeDcdnDomainBpsDataResponseBodyBpsDataPerInterval {
	s.DataModule = v
	return s
}

type DescribeDcdnDomainBpsDataResponseBodyBpsDataPerIntervalDataModule struct {
	Bps             *float32 `json:"Bps,omitempty" xml:"Bps,omitempty"`
	DynamicHttpBps  *float32 `json:"DynamicHttpBps,omitempty" xml:"DynamicHttpBps,omitempty"`
	DynamicHttpsBps *float32 `json:"DynamicHttpsBps,omitempty" xml:"DynamicHttpsBps,omitempty"`
	StaticHttpBps   *float32 `json:"StaticHttpBps,omitempty" xml:"StaticHttpBps,omitempty"`
	StaticHttpsBps  *float32 `json:"StaticHttpsBps,omitempty" xml:"StaticHttpsBps,omitempty"`
	TimeStamp       *string  `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
}

func (s DescribeDcdnDomainBpsDataResponseBodyBpsDataPerIntervalDataModule) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainBpsDataResponseBodyBpsDataPerIntervalDataModule) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainBpsDataResponseBodyBpsDataPerIntervalDataModule) SetBps(v float32) *DescribeDcdnDomainBpsDataResponseBodyBpsDataPerIntervalDataModule {
	s.Bps = &v
	return s
}

func (s *DescribeDcdnDomainBpsDataResponseBodyBpsDataPerIntervalDataModule) SetDynamicHttpBps(v float32) *DescribeDcdnDomainBpsDataResponseBodyBpsDataPerIntervalDataModule {
	s.DynamicHttpBps = &v
	return s
}

func (s *DescribeDcdnDomainBpsDataResponseBodyBpsDataPerIntervalDataModule) SetDynamicHttpsBps(v float32) *DescribeDcdnDomainBpsDataResponseBodyBpsDataPerIntervalDataModule {
	s.DynamicHttpsBps = &v
	return s
}

func (s *DescribeDcdnDomainBpsDataResponseBodyBpsDataPerIntervalDataModule) SetStaticHttpBps(v float32) *DescribeDcdnDomainBpsDataResponseBodyBpsDataPerIntervalDataModule {
	s.StaticHttpBps = &v
	return s
}

func (s *DescribeDcdnDomainBpsDataResponseBodyBpsDataPerIntervalDataModule) SetStaticHttpsBps(v float32) *DescribeDcdnDomainBpsDataResponseBodyBpsDataPerIntervalDataModule {
	s.StaticHttpsBps = &v
	return s
}

func (s *DescribeDcdnDomainBpsDataResponseBodyBpsDataPerIntervalDataModule) SetTimeStamp(v string) *DescribeDcdnDomainBpsDataResponseBodyBpsDataPerIntervalDataModule {
	s.TimeStamp = &v
	return s
}

type DescribeDcdnDomainBpsDataResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDcdnDomainBpsDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDcdnDomainBpsDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainBpsDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainBpsDataResponse) SetHeaders(v map[string]*string) *DescribeDcdnDomainBpsDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnDomainBpsDataResponse) SetBody(v *DescribeDcdnDomainBpsDataResponseBody) *DescribeDcdnDomainBpsDataResponse {
	s.Body = v
	return s
}

type DescribeDcdnDomainByCertificateRequest struct {
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SSLPub  *string `json:"SSLPub,omitempty" xml:"SSLPub,omitempty"`
}

func (s DescribeDcdnDomainByCertificateRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainByCertificateRequest) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainByCertificateRequest) SetOwnerId(v int64) *DescribeDcdnDomainByCertificateRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDcdnDomainByCertificateRequest) SetSSLPub(v string) *DescribeDcdnDomainByCertificateRequest {
	s.SSLPub = &v
	return s
}

type DescribeDcdnDomainByCertificateResponseBody struct {
	CertInfos *DescribeDcdnDomainByCertificateResponseBodyCertInfos `json:"CertInfos,omitempty" xml:"CertInfos,omitempty" type:"Struct"`
	RequestId *string                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDcdnDomainByCertificateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainByCertificateResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainByCertificateResponseBody) SetCertInfos(v *DescribeDcdnDomainByCertificateResponseBodyCertInfos) *DescribeDcdnDomainByCertificateResponseBody {
	s.CertInfos = v
	return s
}

func (s *DescribeDcdnDomainByCertificateResponseBody) SetRequestId(v string) *DescribeDcdnDomainByCertificateResponseBody {
	s.RequestId = &v
	return s
}

type DescribeDcdnDomainByCertificateResponseBodyCertInfos struct {
	CertInfo []*DescribeDcdnDomainByCertificateResponseBodyCertInfosCertInfo `json:"CertInfo,omitempty" xml:"CertInfo,omitempty" type:"Repeated"`
}

func (s DescribeDcdnDomainByCertificateResponseBodyCertInfos) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainByCertificateResponseBodyCertInfos) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainByCertificateResponseBodyCertInfos) SetCertInfo(v []*DescribeDcdnDomainByCertificateResponseBodyCertInfosCertInfo) *DescribeDcdnDomainByCertificateResponseBodyCertInfos {
	s.CertInfo = v
	return s
}

type DescribeDcdnDomainByCertificateResponseBodyCertInfosCertInfo struct {
	CertCaIsLegacy        *string `json:"CertCaIsLegacy,omitempty" xml:"CertCaIsLegacy,omitempty"`
	CertExpireTime        *string `json:"CertExpireTime,omitempty" xml:"CertExpireTime,omitempty"`
	CertExpired           *string `json:"CertExpired,omitempty" xml:"CertExpired,omitempty"`
	CertStartTime         *string `json:"CertStartTime,omitempty" xml:"CertStartTime,omitempty"`
	CertSubjectCommonName *string `json:"CertSubjectCommonName,omitempty" xml:"CertSubjectCommonName,omitempty"`
	CertType              *string `json:"CertType,omitempty" xml:"CertType,omitempty"`
	DomainList            *string `json:"DomainList,omitempty" xml:"DomainList,omitempty"`
	DomainNames           *string `json:"DomainNames,omitempty" xml:"DomainNames,omitempty"`
	Issuer                *string `json:"Issuer,omitempty" xml:"Issuer,omitempty"`
}

func (s DescribeDcdnDomainByCertificateResponseBodyCertInfosCertInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainByCertificateResponseBodyCertInfosCertInfo) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainByCertificateResponseBodyCertInfosCertInfo) SetCertCaIsLegacy(v string) *DescribeDcdnDomainByCertificateResponseBodyCertInfosCertInfo {
	s.CertCaIsLegacy = &v
	return s
}

func (s *DescribeDcdnDomainByCertificateResponseBodyCertInfosCertInfo) SetCertExpireTime(v string) *DescribeDcdnDomainByCertificateResponseBodyCertInfosCertInfo {
	s.CertExpireTime = &v
	return s
}

func (s *DescribeDcdnDomainByCertificateResponseBodyCertInfosCertInfo) SetCertExpired(v string) *DescribeDcdnDomainByCertificateResponseBodyCertInfosCertInfo {
	s.CertExpired = &v
	return s
}

func (s *DescribeDcdnDomainByCertificateResponseBodyCertInfosCertInfo) SetCertStartTime(v string) *DescribeDcdnDomainByCertificateResponseBodyCertInfosCertInfo {
	s.CertStartTime = &v
	return s
}

func (s *DescribeDcdnDomainByCertificateResponseBodyCertInfosCertInfo) SetCertSubjectCommonName(v string) *DescribeDcdnDomainByCertificateResponseBodyCertInfosCertInfo {
	s.CertSubjectCommonName = &v
	return s
}

func (s *DescribeDcdnDomainByCertificateResponseBodyCertInfosCertInfo) SetCertType(v string) *DescribeDcdnDomainByCertificateResponseBodyCertInfosCertInfo {
	s.CertType = &v
	return s
}

func (s *DescribeDcdnDomainByCertificateResponseBodyCertInfosCertInfo) SetDomainList(v string) *DescribeDcdnDomainByCertificateResponseBodyCertInfosCertInfo {
	s.DomainList = &v
	return s
}

func (s *DescribeDcdnDomainByCertificateResponseBodyCertInfosCertInfo) SetDomainNames(v string) *DescribeDcdnDomainByCertificateResponseBodyCertInfosCertInfo {
	s.DomainNames = &v
	return s
}

func (s *DescribeDcdnDomainByCertificateResponseBodyCertInfosCertInfo) SetIssuer(v string) *DescribeDcdnDomainByCertificateResponseBodyCertInfosCertInfo {
	s.Issuer = &v
	return s
}

type DescribeDcdnDomainByCertificateResponse struct {
	Headers map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDcdnDomainByCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDcdnDomainByCertificateResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainByCertificateResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainByCertificateResponse) SetHeaders(v map[string]*string) *DescribeDcdnDomainByCertificateResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnDomainByCertificateResponse) SetBody(v *DescribeDcdnDomainByCertificateResponseBody) *DescribeDcdnDomainByCertificateResponse {
	s.Body = v
	return s
}

type DescribeDcdnDomainCcActivityLogRequest struct {
	DomainName    *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EndTime       *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNumber    *int64  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize      *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RuleName      *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	StartTime     *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	TriggerObject *string `json:"TriggerObject,omitempty" xml:"TriggerObject,omitempty"`
	Value         *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDcdnDomainCcActivityLogRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainCcActivityLogRequest) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainCcActivityLogRequest) SetDomainName(v string) *DescribeDcdnDomainCcActivityLogRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDcdnDomainCcActivityLogRequest) SetEndTime(v string) *DescribeDcdnDomainCcActivityLogRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDcdnDomainCcActivityLogRequest) SetOwnerId(v int64) *DescribeDcdnDomainCcActivityLogRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDcdnDomainCcActivityLogRequest) SetPageNumber(v int64) *DescribeDcdnDomainCcActivityLogRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDcdnDomainCcActivityLogRequest) SetPageSize(v int64) *DescribeDcdnDomainCcActivityLogRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDcdnDomainCcActivityLogRequest) SetRuleName(v string) *DescribeDcdnDomainCcActivityLogRequest {
	s.RuleName = &v
	return s
}

func (s *DescribeDcdnDomainCcActivityLogRequest) SetStartTime(v string) *DescribeDcdnDomainCcActivityLogRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDcdnDomainCcActivityLogRequest) SetTriggerObject(v string) *DescribeDcdnDomainCcActivityLogRequest {
	s.TriggerObject = &v
	return s
}

func (s *DescribeDcdnDomainCcActivityLogRequest) SetValue(v string) *DescribeDcdnDomainCcActivityLogRequest {
	s.Value = &v
	return s
}

type DescribeDcdnDomainCcActivityLogResponseBody struct {
	ActivityLog []*DescribeDcdnDomainCcActivityLogResponseBodyActivityLog `json:"ActivityLog,omitempty" xml:"ActivityLog,omitempty" type:"Repeated"`
	PageIndex   *int64                                                    `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	PageSize    *int64                                                    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId   *string                                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Total       *int64                                                    `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeDcdnDomainCcActivityLogResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainCcActivityLogResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainCcActivityLogResponseBody) SetActivityLog(v []*DescribeDcdnDomainCcActivityLogResponseBodyActivityLog) *DescribeDcdnDomainCcActivityLogResponseBody {
	s.ActivityLog = v
	return s
}

func (s *DescribeDcdnDomainCcActivityLogResponseBody) SetPageIndex(v int64) *DescribeDcdnDomainCcActivityLogResponseBody {
	s.PageIndex = &v
	return s
}

func (s *DescribeDcdnDomainCcActivityLogResponseBody) SetPageSize(v int64) *DescribeDcdnDomainCcActivityLogResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeDcdnDomainCcActivityLogResponseBody) SetRequestId(v string) *DescribeDcdnDomainCcActivityLogResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDcdnDomainCcActivityLogResponseBody) SetTotal(v int64) *DescribeDcdnDomainCcActivityLogResponseBody {
	s.Total = &v
	return s
}

type DescribeDcdnDomainCcActivityLogResponseBodyActivityLog struct {
	Action        *string `json:"Action,omitempty" xml:"Action,omitempty"`
	DomainName    *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	RuleName      *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	TimeStamp     *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
	TriggerObject *string `json:"TriggerObject,omitempty" xml:"TriggerObject,omitempty"`
	Ttl           *int64  `json:"Ttl,omitempty" xml:"Ttl,omitempty"`
	Value         *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDcdnDomainCcActivityLogResponseBodyActivityLog) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainCcActivityLogResponseBodyActivityLog) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainCcActivityLogResponseBodyActivityLog) SetAction(v string) *DescribeDcdnDomainCcActivityLogResponseBodyActivityLog {
	s.Action = &v
	return s
}

func (s *DescribeDcdnDomainCcActivityLogResponseBodyActivityLog) SetDomainName(v string) *DescribeDcdnDomainCcActivityLogResponseBodyActivityLog {
	s.DomainName = &v
	return s
}

func (s *DescribeDcdnDomainCcActivityLogResponseBodyActivityLog) SetRuleName(v string) *DescribeDcdnDomainCcActivityLogResponseBodyActivityLog {
	s.RuleName = &v
	return s
}

func (s *DescribeDcdnDomainCcActivityLogResponseBodyActivityLog) SetTimeStamp(v string) *DescribeDcdnDomainCcActivityLogResponseBodyActivityLog {
	s.TimeStamp = &v
	return s
}

func (s *DescribeDcdnDomainCcActivityLogResponseBodyActivityLog) SetTriggerObject(v string) *DescribeDcdnDomainCcActivityLogResponseBodyActivityLog {
	s.TriggerObject = &v
	return s
}

func (s *DescribeDcdnDomainCcActivityLogResponseBodyActivityLog) SetTtl(v int64) *DescribeDcdnDomainCcActivityLogResponseBodyActivityLog {
	s.Ttl = &v
	return s
}

func (s *DescribeDcdnDomainCcActivityLogResponseBodyActivityLog) SetValue(v string) *DescribeDcdnDomainCcActivityLogResponseBodyActivityLog {
	s.Value = &v
	return s
}

type DescribeDcdnDomainCcActivityLogResponse struct {
	Headers map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDcdnDomainCcActivityLogResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDcdnDomainCcActivityLogResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainCcActivityLogResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainCcActivityLogResponse) SetHeaders(v map[string]*string) *DescribeDcdnDomainCcActivityLogResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnDomainCcActivityLogResponse) SetBody(v *DescribeDcdnDomainCcActivityLogResponseBody) *DescribeDcdnDomainCcActivityLogResponse {
	s.Body = v
	return s
}

type DescribeDcdnDomainCertificateInfoRequest struct {
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s DescribeDcdnDomainCertificateInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainCertificateInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainCertificateInfoRequest) SetDomainName(v string) *DescribeDcdnDomainCertificateInfoRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDcdnDomainCertificateInfoRequest) SetOwnerId(v int64) *DescribeDcdnDomainCertificateInfoRequest {
	s.OwnerId = &v
	return s
}

type DescribeDcdnDomainCertificateInfoResponseBody struct {
	CertInfos *DescribeDcdnDomainCertificateInfoResponseBodyCertInfos `json:"CertInfos,omitempty" xml:"CertInfos,omitempty" type:"Struct"`
	RequestId *string                                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDcdnDomainCertificateInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainCertificateInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainCertificateInfoResponseBody) SetCertInfos(v *DescribeDcdnDomainCertificateInfoResponseBodyCertInfos) *DescribeDcdnDomainCertificateInfoResponseBody {
	s.CertInfos = v
	return s
}

func (s *DescribeDcdnDomainCertificateInfoResponseBody) SetRequestId(v string) *DescribeDcdnDomainCertificateInfoResponseBody {
	s.RequestId = &v
	return s
}

type DescribeDcdnDomainCertificateInfoResponseBodyCertInfos struct {
	CertInfo []*DescribeDcdnDomainCertificateInfoResponseBodyCertInfosCertInfo `json:"CertInfo,omitempty" xml:"CertInfo,omitempty" type:"Repeated"`
}

func (s DescribeDcdnDomainCertificateInfoResponseBodyCertInfos) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainCertificateInfoResponseBodyCertInfos) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainCertificateInfoResponseBodyCertInfos) SetCertInfo(v []*DescribeDcdnDomainCertificateInfoResponseBodyCertInfosCertInfo) *DescribeDcdnDomainCertificateInfoResponseBodyCertInfos {
	s.CertInfo = v
	return s
}

type DescribeDcdnDomainCertificateInfoResponseBodyCertInfosCertInfo struct {
	CertDomainName *string `json:"CertDomainName,omitempty" xml:"CertDomainName,omitempty"`
	CertExpireTime *string `json:"CertExpireTime,omitempty" xml:"CertExpireTime,omitempty"`
	CertLife       *string `json:"CertLife,omitempty" xml:"CertLife,omitempty"`
	CertName       *string `json:"CertName,omitempty" xml:"CertName,omitempty"`
	CertOrg        *string `json:"CertOrg,omitempty" xml:"CertOrg,omitempty"`
	CertType       *string `json:"CertType,omitempty" xml:"CertType,omitempty"`
	DomainName     *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	SSLProtocol    *string `json:"SSLProtocol,omitempty" xml:"SSLProtocol,omitempty"`
	SSLPub         *string `json:"SSLPub,omitempty" xml:"SSLPub,omitempty"`
	Status         *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeDcdnDomainCertificateInfoResponseBodyCertInfosCertInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainCertificateInfoResponseBodyCertInfosCertInfo) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainCertificateInfoResponseBodyCertInfosCertInfo) SetCertDomainName(v string) *DescribeDcdnDomainCertificateInfoResponseBodyCertInfosCertInfo {
	s.CertDomainName = &v
	return s
}

func (s *DescribeDcdnDomainCertificateInfoResponseBodyCertInfosCertInfo) SetCertExpireTime(v string) *DescribeDcdnDomainCertificateInfoResponseBodyCertInfosCertInfo {
	s.CertExpireTime = &v
	return s
}

func (s *DescribeDcdnDomainCertificateInfoResponseBodyCertInfosCertInfo) SetCertLife(v string) *DescribeDcdnDomainCertificateInfoResponseBodyCertInfosCertInfo {
	s.CertLife = &v
	return s
}

func (s *DescribeDcdnDomainCertificateInfoResponseBodyCertInfosCertInfo) SetCertName(v string) *DescribeDcdnDomainCertificateInfoResponseBodyCertInfosCertInfo {
	s.CertName = &v
	return s
}

func (s *DescribeDcdnDomainCertificateInfoResponseBodyCertInfosCertInfo) SetCertOrg(v string) *DescribeDcdnDomainCertificateInfoResponseBodyCertInfosCertInfo {
	s.CertOrg = &v
	return s
}

func (s *DescribeDcdnDomainCertificateInfoResponseBodyCertInfosCertInfo) SetCertType(v string) *DescribeDcdnDomainCertificateInfoResponseBodyCertInfosCertInfo {
	s.CertType = &v
	return s
}

func (s *DescribeDcdnDomainCertificateInfoResponseBodyCertInfosCertInfo) SetDomainName(v string) *DescribeDcdnDomainCertificateInfoResponseBodyCertInfosCertInfo {
	s.DomainName = &v
	return s
}

func (s *DescribeDcdnDomainCertificateInfoResponseBodyCertInfosCertInfo) SetSSLProtocol(v string) *DescribeDcdnDomainCertificateInfoResponseBodyCertInfosCertInfo {
	s.SSLProtocol = &v
	return s
}

func (s *DescribeDcdnDomainCertificateInfoResponseBodyCertInfosCertInfo) SetSSLPub(v string) *DescribeDcdnDomainCertificateInfoResponseBodyCertInfosCertInfo {
	s.SSLPub = &v
	return s
}

func (s *DescribeDcdnDomainCertificateInfoResponseBodyCertInfosCertInfo) SetStatus(v string) *DescribeDcdnDomainCertificateInfoResponseBodyCertInfosCertInfo {
	s.Status = &v
	return s
}

type DescribeDcdnDomainCertificateInfoResponse struct {
	Headers map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDcdnDomainCertificateInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDcdnDomainCertificateInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainCertificateInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainCertificateInfoResponse) SetHeaders(v map[string]*string) *DescribeDcdnDomainCertificateInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnDomainCertificateInfoResponse) SetBody(v *DescribeDcdnDomainCertificateInfoResponseBody) *DescribeDcdnDomainCertificateInfoResponse {
	s.Body = v
	return s
}

type DescribeDcdnDomainCnameRequest struct {
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s DescribeDcdnDomainCnameRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainCnameRequest) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainCnameRequest) SetDomainName(v string) *DescribeDcdnDomainCnameRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDcdnDomainCnameRequest) SetOwnerId(v int64) *DescribeDcdnDomainCnameRequest {
	s.OwnerId = &v
	return s
}

type DescribeDcdnDomainCnameResponseBody struct {
	CnameDatas *DescribeDcdnDomainCnameResponseBodyCnameDatas `json:"CnameDatas,omitempty" xml:"CnameDatas,omitempty" type:"Struct"`
	RequestId  *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDcdnDomainCnameResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainCnameResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainCnameResponseBody) SetCnameDatas(v *DescribeDcdnDomainCnameResponseBodyCnameDatas) *DescribeDcdnDomainCnameResponseBody {
	s.CnameDatas = v
	return s
}

func (s *DescribeDcdnDomainCnameResponseBody) SetRequestId(v string) *DescribeDcdnDomainCnameResponseBody {
	s.RequestId = &v
	return s
}

type DescribeDcdnDomainCnameResponseBodyCnameDatas struct {
	Data []*DescribeDcdnDomainCnameResponseBodyCnameDatasData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
}

func (s DescribeDcdnDomainCnameResponseBodyCnameDatas) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainCnameResponseBodyCnameDatas) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainCnameResponseBodyCnameDatas) SetData(v []*DescribeDcdnDomainCnameResponseBodyCnameDatasData) *DescribeDcdnDomainCnameResponseBodyCnameDatas {
	s.Data = v
	return s
}

type DescribeDcdnDomainCnameResponseBodyCnameDatasData struct {
	Cname  *string `json:"Cname,omitempty" xml:"Cname,omitempty"`
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	Status *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeDcdnDomainCnameResponseBodyCnameDatasData) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainCnameResponseBodyCnameDatasData) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainCnameResponseBodyCnameDatasData) SetCname(v string) *DescribeDcdnDomainCnameResponseBodyCnameDatasData {
	s.Cname = &v
	return s
}

func (s *DescribeDcdnDomainCnameResponseBodyCnameDatasData) SetDomain(v string) *DescribeDcdnDomainCnameResponseBodyCnameDatasData {
	s.Domain = &v
	return s
}

func (s *DescribeDcdnDomainCnameResponseBodyCnameDatasData) SetStatus(v int32) *DescribeDcdnDomainCnameResponseBodyCnameDatasData {
	s.Status = &v
	return s
}

type DescribeDcdnDomainCnameResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDcdnDomainCnameResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDcdnDomainCnameResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainCnameResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainCnameResponse) SetHeaders(v map[string]*string) *DescribeDcdnDomainCnameResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnDomainCnameResponse) SetBody(v *DescribeDcdnDomainCnameResponseBody) *DescribeDcdnDomainCnameResponse {
	s.Body = v
	return s
}

type DescribeDcdnDomainConfigsRequest struct {
	ConfigId      *string `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	DomainName    *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	FunctionNames *string `json:"FunctionNames,omitempty" xml:"FunctionNames,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeDcdnDomainConfigsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainConfigsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainConfigsRequest) SetConfigId(v string) *DescribeDcdnDomainConfigsRequest {
	s.ConfigId = &v
	return s
}

func (s *DescribeDcdnDomainConfigsRequest) SetDomainName(v string) *DescribeDcdnDomainConfigsRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDcdnDomainConfigsRequest) SetFunctionNames(v string) *DescribeDcdnDomainConfigsRequest {
	s.FunctionNames = &v
	return s
}

func (s *DescribeDcdnDomainConfigsRequest) SetOwnerId(v int64) *DescribeDcdnDomainConfigsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDcdnDomainConfigsRequest) SetSecurityToken(v string) *DescribeDcdnDomainConfigsRequest {
	s.SecurityToken = &v
	return s
}

type DescribeDcdnDomainConfigsResponseBody struct {
	DomainConfigs *DescribeDcdnDomainConfigsResponseBodyDomainConfigs `json:"DomainConfigs,omitempty" xml:"DomainConfigs,omitempty" type:"Struct"`
	RequestId     *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDcdnDomainConfigsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainConfigsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainConfigsResponseBody) SetDomainConfigs(v *DescribeDcdnDomainConfigsResponseBodyDomainConfigs) *DescribeDcdnDomainConfigsResponseBody {
	s.DomainConfigs = v
	return s
}

func (s *DescribeDcdnDomainConfigsResponseBody) SetRequestId(v string) *DescribeDcdnDomainConfigsResponseBody {
	s.RequestId = &v
	return s
}

type DescribeDcdnDomainConfigsResponseBodyDomainConfigs struct {
	DomainConfig []*DescribeDcdnDomainConfigsResponseBodyDomainConfigsDomainConfig `json:"DomainConfig,omitempty" xml:"DomainConfig,omitempty" type:"Repeated"`
}

func (s DescribeDcdnDomainConfigsResponseBodyDomainConfigs) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainConfigsResponseBodyDomainConfigs) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainConfigsResponseBodyDomainConfigs) SetDomainConfig(v []*DescribeDcdnDomainConfigsResponseBodyDomainConfigsDomainConfig) *DescribeDcdnDomainConfigsResponseBodyDomainConfigs {
	s.DomainConfig = v
	return s
}

type DescribeDcdnDomainConfigsResponseBodyDomainConfigsDomainConfig struct {
	ConfigId     *string                                                                     `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	FunctionArgs *DescribeDcdnDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgs `json:"FunctionArgs,omitempty" xml:"FunctionArgs,omitempty" type:"Struct"`
	FunctionName *string                                                                     `json:"FunctionName,omitempty" xml:"FunctionName,omitempty"`
	Status       *string                                                                     `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeDcdnDomainConfigsResponseBodyDomainConfigsDomainConfig) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainConfigsResponseBodyDomainConfigsDomainConfig) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainConfigsResponseBodyDomainConfigsDomainConfig) SetConfigId(v string) *DescribeDcdnDomainConfigsResponseBodyDomainConfigsDomainConfig {
	s.ConfigId = &v
	return s
}

func (s *DescribeDcdnDomainConfigsResponseBodyDomainConfigsDomainConfig) SetFunctionArgs(v *DescribeDcdnDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgs) *DescribeDcdnDomainConfigsResponseBodyDomainConfigsDomainConfig {
	s.FunctionArgs = v
	return s
}

func (s *DescribeDcdnDomainConfigsResponseBodyDomainConfigsDomainConfig) SetFunctionName(v string) *DescribeDcdnDomainConfigsResponseBodyDomainConfigsDomainConfig {
	s.FunctionName = &v
	return s
}

func (s *DescribeDcdnDomainConfigsResponseBodyDomainConfigsDomainConfig) SetStatus(v string) *DescribeDcdnDomainConfigsResponseBodyDomainConfigsDomainConfig {
	s.Status = &v
	return s
}

type DescribeDcdnDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgs struct {
	FunctionArg []*DescribeDcdnDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgsFunctionArg `json:"FunctionArg,omitempty" xml:"FunctionArg,omitempty" type:"Repeated"`
}

func (s DescribeDcdnDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgs) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgs) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgs) SetFunctionArg(v []*DescribeDcdnDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgsFunctionArg) *DescribeDcdnDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgs {
	s.FunctionArg = v
	return s
}

type DescribeDcdnDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgsFunctionArg struct {
	ArgName  *string `json:"ArgName,omitempty" xml:"ArgName,omitempty"`
	ArgValue *string `json:"ArgValue,omitempty" xml:"ArgValue,omitempty"`
}

func (s DescribeDcdnDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgsFunctionArg) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgsFunctionArg) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgsFunctionArg) SetArgName(v string) *DescribeDcdnDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgsFunctionArg {
	s.ArgName = &v
	return s
}

func (s *DescribeDcdnDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgsFunctionArg) SetArgValue(v string) *DescribeDcdnDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgsFunctionArg {
	s.ArgValue = &v
	return s
}

type DescribeDcdnDomainConfigsResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDcdnDomainConfigsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDcdnDomainConfigsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainConfigsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainConfigsResponse) SetHeaders(v map[string]*string) *DescribeDcdnDomainConfigsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnDomainConfigsResponse) SetBody(v *DescribeDcdnDomainConfigsResponseBody) *DescribeDcdnDomainConfigsResponse {
	s.Body = v
	return s
}

type DescribeDcdnDomainDetailRequest struct {
	DomainName    *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeDcdnDomainDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainDetailRequest) SetDomainName(v string) *DescribeDcdnDomainDetailRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDcdnDomainDetailRequest) SetOwnerId(v int64) *DescribeDcdnDomainDetailRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDcdnDomainDetailRequest) SetSecurityToken(v string) *DescribeDcdnDomainDetailRequest {
	s.SecurityToken = &v
	return s
}

type DescribeDcdnDomainDetailResponseBody struct {
	DomainDetail *DescribeDcdnDomainDetailResponseBodyDomainDetail `json:"DomainDetail,omitempty" xml:"DomainDetail,omitempty" type:"Struct"`
	RequestId    *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDcdnDomainDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainDetailResponseBody) SetDomainDetail(v *DescribeDcdnDomainDetailResponseBodyDomainDetail) *DescribeDcdnDomainDetailResponseBody {
	s.DomainDetail = v
	return s
}

func (s *DescribeDcdnDomainDetailResponseBody) SetRequestId(v string) *DescribeDcdnDomainDetailResponseBody {
	s.RequestId = &v
	return s
}

type DescribeDcdnDomainDetailResponseBodyDomainDetail struct {
	CertName        *string                                                  `json:"CertName,omitempty" xml:"CertName,omitempty"`
	Cname           *string                                                  `json:"Cname,omitempty" xml:"Cname,omitempty"`
	Description     *string                                                  `json:"Description,omitempty" xml:"Description,omitempty"`
	DomainName      *string                                                  `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	DomainStatus    *string                                                  `json:"DomainStatus,omitempty" xml:"DomainStatus,omitempty"`
	GmtCreated      *string                                                  `json:"GmtCreated,omitempty" xml:"GmtCreated,omitempty"`
	GmtModified     *string                                                  `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	ResourceGroupId *string                                                  `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	SSLProtocol     *string                                                  `json:"SSLProtocol,omitempty" xml:"SSLProtocol,omitempty"`
	SSLPub          *string                                                  `json:"SSLPub,omitempty" xml:"SSLPub,omitempty"`
	Scope           *string                                                  `json:"Scope,omitempty" xml:"Scope,omitempty"`
	Sources         *DescribeDcdnDomainDetailResponseBodyDomainDetailSources `json:"Sources,omitempty" xml:"Sources,omitempty" type:"Struct"`
}

func (s DescribeDcdnDomainDetailResponseBodyDomainDetail) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainDetailResponseBodyDomainDetail) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainDetailResponseBodyDomainDetail) SetCertName(v string) *DescribeDcdnDomainDetailResponseBodyDomainDetail {
	s.CertName = &v
	return s
}

func (s *DescribeDcdnDomainDetailResponseBodyDomainDetail) SetCname(v string) *DescribeDcdnDomainDetailResponseBodyDomainDetail {
	s.Cname = &v
	return s
}

func (s *DescribeDcdnDomainDetailResponseBodyDomainDetail) SetDescription(v string) *DescribeDcdnDomainDetailResponseBodyDomainDetail {
	s.Description = &v
	return s
}

func (s *DescribeDcdnDomainDetailResponseBodyDomainDetail) SetDomainName(v string) *DescribeDcdnDomainDetailResponseBodyDomainDetail {
	s.DomainName = &v
	return s
}

func (s *DescribeDcdnDomainDetailResponseBodyDomainDetail) SetDomainStatus(v string) *DescribeDcdnDomainDetailResponseBodyDomainDetail {
	s.DomainStatus = &v
	return s
}

func (s *DescribeDcdnDomainDetailResponseBodyDomainDetail) SetGmtCreated(v string) *DescribeDcdnDomainDetailResponseBodyDomainDetail {
	s.GmtCreated = &v
	return s
}

func (s *DescribeDcdnDomainDetailResponseBodyDomainDetail) SetGmtModified(v string) *DescribeDcdnDomainDetailResponseBodyDomainDetail {
	s.GmtModified = &v
	return s
}

func (s *DescribeDcdnDomainDetailResponseBodyDomainDetail) SetResourceGroupId(v string) *DescribeDcdnDomainDetailResponseBodyDomainDetail {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDcdnDomainDetailResponseBodyDomainDetail) SetSSLProtocol(v string) *DescribeDcdnDomainDetailResponseBodyDomainDetail {
	s.SSLProtocol = &v
	return s
}

func (s *DescribeDcdnDomainDetailResponseBodyDomainDetail) SetSSLPub(v string) *DescribeDcdnDomainDetailResponseBodyDomainDetail {
	s.SSLPub = &v
	return s
}

func (s *DescribeDcdnDomainDetailResponseBodyDomainDetail) SetScope(v string) *DescribeDcdnDomainDetailResponseBodyDomainDetail {
	s.Scope = &v
	return s
}

func (s *DescribeDcdnDomainDetailResponseBodyDomainDetail) SetSources(v *DescribeDcdnDomainDetailResponseBodyDomainDetailSources) *DescribeDcdnDomainDetailResponseBodyDomainDetail {
	s.Sources = v
	return s
}

type DescribeDcdnDomainDetailResponseBodyDomainDetailSources struct {
	Source []*DescribeDcdnDomainDetailResponseBodyDomainDetailSourcesSource `json:"Source,omitempty" xml:"Source,omitempty" type:"Repeated"`
}

func (s DescribeDcdnDomainDetailResponseBodyDomainDetailSources) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainDetailResponseBodyDomainDetailSources) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainDetailResponseBodyDomainDetailSources) SetSource(v []*DescribeDcdnDomainDetailResponseBodyDomainDetailSourcesSource) *DescribeDcdnDomainDetailResponseBodyDomainDetailSources {
	s.Source = v
	return s
}

type DescribeDcdnDomainDetailResponseBodyDomainDetailSourcesSource struct {
	Content  *string `json:"Content,omitempty" xml:"Content,omitempty"`
	Enabled  *string `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	Port     *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
	Priority *string `json:"Priority,omitempty" xml:"Priority,omitempty"`
	Type     *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Weight   *string `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s DescribeDcdnDomainDetailResponseBodyDomainDetailSourcesSource) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainDetailResponseBodyDomainDetailSourcesSource) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainDetailResponseBodyDomainDetailSourcesSource) SetContent(v string) *DescribeDcdnDomainDetailResponseBodyDomainDetailSourcesSource {
	s.Content = &v
	return s
}

func (s *DescribeDcdnDomainDetailResponseBodyDomainDetailSourcesSource) SetEnabled(v string) *DescribeDcdnDomainDetailResponseBodyDomainDetailSourcesSource {
	s.Enabled = &v
	return s
}

func (s *DescribeDcdnDomainDetailResponseBodyDomainDetailSourcesSource) SetPort(v int32) *DescribeDcdnDomainDetailResponseBodyDomainDetailSourcesSource {
	s.Port = &v
	return s
}

func (s *DescribeDcdnDomainDetailResponseBodyDomainDetailSourcesSource) SetPriority(v string) *DescribeDcdnDomainDetailResponseBodyDomainDetailSourcesSource {
	s.Priority = &v
	return s
}

func (s *DescribeDcdnDomainDetailResponseBodyDomainDetailSourcesSource) SetType(v string) *DescribeDcdnDomainDetailResponseBodyDomainDetailSourcesSource {
	s.Type = &v
	return s
}

func (s *DescribeDcdnDomainDetailResponseBodyDomainDetailSourcesSource) SetWeight(v string) *DescribeDcdnDomainDetailResponseBodyDomainDetailSourcesSource {
	s.Weight = &v
	return s
}

type DescribeDcdnDomainDetailResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDcdnDomainDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDcdnDomainDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainDetailResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainDetailResponse) SetHeaders(v map[string]*string) *DescribeDcdnDomainDetailResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnDomainDetailResponse) SetBody(v *DescribeDcdnDomainDetailResponseBody) *DescribeDcdnDomainDetailResponse {
	s.Body = v
	return s
}

type DescribeDcdnDomainHitRateDataRequest struct {
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EndTime    *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Interval   *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	StartTime  *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDcdnDomainHitRateDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainHitRateDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainHitRateDataRequest) SetDomainName(v string) *DescribeDcdnDomainHitRateDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDcdnDomainHitRateDataRequest) SetEndTime(v string) *DescribeDcdnDomainHitRateDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDcdnDomainHitRateDataRequest) SetInterval(v string) *DescribeDcdnDomainHitRateDataRequest {
	s.Interval = &v
	return s
}

func (s *DescribeDcdnDomainHitRateDataRequest) SetOwnerId(v int64) *DescribeDcdnDomainHitRateDataRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDcdnDomainHitRateDataRequest) SetStartTime(v string) *DescribeDcdnDomainHitRateDataRequest {
	s.StartTime = &v
	return s
}

type DescribeDcdnDomainHitRateDataResponseBody struct {
	DataInterval       *string                                                      `json:"DataInterval,omitempty" xml:"DataInterval,omitempty"`
	DomainName         *string                                                      `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EndTime            *string                                                      `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	HitRatePerInterval *DescribeDcdnDomainHitRateDataResponseBodyHitRatePerInterval `json:"HitRatePerInterval,omitempty" xml:"HitRatePerInterval,omitempty" type:"Struct"`
	RequestId          *string                                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StartTime          *string                                                      `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDcdnDomainHitRateDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainHitRateDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainHitRateDataResponseBody) SetDataInterval(v string) *DescribeDcdnDomainHitRateDataResponseBody {
	s.DataInterval = &v
	return s
}

func (s *DescribeDcdnDomainHitRateDataResponseBody) SetDomainName(v string) *DescribeDcdnDomainHitRateDataResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeDcdnDomainHitRateDataResponseBody) SetEndTime(v string) *DescribeDcdnDomainHitRateDataResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeDcdnDomainHitRateDataResponseBody) SetHitRatePerInterval(v *DescribeDcdnDomainHitRateDataResponseBodyHitRatePerInterval) *DescribeDcdnDomainHitRateDataResponseBody {
	s.HitRatePerInterval = v
	return s
}

func (s *DescribeDcdnDomainHitRateDataResponseBody) SetRequestId(v string) *DescribeDcdnDomainHitRateDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDcdnDomainHitRateDataResponseBody) SetStartTime(v string) *DescribeDcdnDomainHitRateDataResponseBody {
	s.StartTime = &v
	return s
}

type DescribeDcdnDomainHitRateDataResponseBodyHitRatePerInterval struct {
	DataModule []*DescribeDcdnDomainHitRateDataResponseBodyHitRatePerIntervalDataModule `json:"DataModule,omitempty" xml:"DataModule,omitempty" type:"Repeated"`
}

func (s DescribeDcdnDomainHitRateDataResponseBodyHitRatePerInterval) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainHitRateDataResponseBodyHitRatePerInterval) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainHitRateDataResponseBodyHitRatePerInterval) SetDataModule(v []*DescribeDcdnDomainHitRateDataResponseBodyHitRatePerIntervalDataModule) *DescribeDcdnDomainHitRateDataResponseBodyHitRatePerInterval {
	s.DataModule = v
	return s
}

type DescribeDcdnDomainHitRateDataResponseBodyHitRatePerIntervalDataModule struct {
	ByteHitRate *float32 `json:"ByteHitRate,omitempty" xml:"ByteHitRate,omitempty"`
	ReqHitRate  *float32 `json:"ReqHitRate,omitempty" xml:"ReqHitRate,omitempty"`
	TimeStamp   *string  `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
}

func (s DescribeDcdnDomainHitRateDataResponseBodyHitRatePerIntervalDataModule) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainHitRateDataResponseBodyHitRatePerIntervalDataModule) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainHitRateDataResponseBodyHitRatePerIntervalDataModule) SetByteHitRate(v float32) *DescribeDcdnDomainHitRateDataResponseBodyHitRatePerIntervalDataModule {
	s.ByteHitRate = &v
	return s
}

func (s *DescribeDcdnDomainHitRateDataResponseBodyHitRatePerIntervalDataModule) SetReqHitRate(v float32) *DescribeDcdnDomainHitRateDataResponseBodyHitRatePerIntervalDataModule {
	s.ReqHitRate = &v
	return s
}

func (s *DescribeDcdnDomainHitRateDataResponseBodyHitRatePerIntervalDataModule) SetTimeStamp(v string) *DescribeDcdnDomainHitRateDataResponseBodyHitRatePerIntervalDataModule {
	s.TimeStamp = &v
	return s
}

type DescribeDcdnDomainHitRateDataResponse struct {
	Headers map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDcdnDomainHitRateDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDcdnDomainHitRateDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainHitRateDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainHitRateDataResponse) SetHeaders(v map[string]*string) *DescribeDcdnDomainHitRateDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnDomainHitRateDataResponse) SetBody(v *DescribeDcdnDomainHitRateDataResponseBody) *DescribeDcdnDomainHitRateDataResponse {
	s.Body = v
	return s
}

type DescribeDcdnDomainHttpCodeDataRequest struct {
	DomainName     *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EndTime        *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Interval       *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	IspNameEn      *string `json:"IspNameEn,omitempty" xml:"IspNameEn,omitempty"`
	LocationNameEn *string `json:"LocationNameEn,omitempty" xml:"LocationNameEn,omitempty"`
	OwnerId        *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	StartTime      *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDcdnDomainHttpCodeDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainHttpCodeDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainHttpCodeDataRequest) SetDomainName(v string) *DescribeDcdnDomainHttpCodeDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDcdnDomainHttpCodeDataRequest) SetEndTime(v string) *DescribeDcdnDomainHttpCodeDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDcdnDomainHttpCodeDataRequest) SetInterval(v string) *DescribeDcdnDomainHttpCodeDataRequest {
	s.Interval = &v
	return s
}

func (s *DescribeDcdnDomainHttpCodeDataRequest) SetIspNameEn(v string) *DescribeDcdnDomainHttpCodeDataRequest {
	s.IspNameEn = &v
	return s
}

func (s *DescribeDcdnDomainHttpCodeDataRequest) SetLocationNameEn(v string) *DescribeDcdnDomainHttpCodeDataRequest {
	s.LocationNameEn = &v
	return s
}

func (s *DescribeDcdnDomainHttpCodeDataRequest) SetOwnerId(v int64) *DescribeDcdnDomainHttpCodeDataRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDcdnDomainHttpCodeDataRequest) SetStartTime(v string) *DescribeDcdnDomainHttpCodeDataRequest {
	s.StartTime = &v
	return s
}

type DescribeDcdnDomainHttpCodeDataResponseBody struct {
	DataInterval    *string                                                    `json:"DataInterval,omitempty" xml:"DataInterval,omitempty"`
	DataPerInterval *DescribeDcdnDomainHttpCodeDataResponseBodyDataPerInterval `json:"DataPerInterval,omitempty" xml:"DataPerInterval,omitempty" type:"Struct"`
	DomainName      *string                                                    `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EndTime         *string                                                    `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	RequestId       *string                                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StartTime       *string                                                    `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDcdnDomainHttpCodeDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainHttpCodeDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainHttpCodeDataResponseBody) SetDataInterval(v string) *DescribeDcdnDomainHttpCodeDataResponseBody {
	s.DataInterval = &v
	return s
}

func (s *DescribeDcdnDomainHttpCodeDataResponseBody) SetDataPerInterval(v *DescribeDcdnDomainHttpCodeDataResponseBodyDataPerInterval) *DescribeDcdnDomainHttpCodeDataResponseBody {
	s.DataPerInterval = v
	return s
}

func (s *DescribeDcdnDomainHttpCodeDataResponseBody) SetDomainName(v string) *DescribeDcdnDomainHttpCodeDataResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeDcdnDomainHttpCodeDataResponseBody) SetEndTime(v string) *DescribeDcdnDomainHttpCodeDataResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeDcdnDomainHttpCodeDataResponseBody) SetRequestId(v string) *DescribeDcdnDomainHttpCodeDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDcdnDomainHttpCodeDataResponseBody) SetStartTime(v string) *DescribeDcdnDomainHttpCodeDataResponseBody {
	s.StartTime = &v
	return s
}

type DescribeDcdnDomainHttpCodeDataResponseBodyDataPerInterval struct {
	DataModule []*DescribeDcdnDomainHttpCodeDataResponseBodyDataPerIntervalDataModule `json:"DataModule,omitempty" xml:"DataModule,omitempty" type:"Repeated"`
}

func (s DescribeDcdnDomainHttpCodeDataResponseBodyDataPerInterval) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainHttpCodeDataResponseBodyDataPerInterval) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainHttpCodeDataResponseBodyDataPerInterval) SetDataModule(v []*DescribeDcdnDomainHttpCodeDataResponseBodyDataPerIntervalDataModule) *DescribeDcdnDomainHttpCodeDataResponseBodyDataPerInterval {
	s.DataModule = v
	return s
}

type DescribeDcdnDomainHttpCodeDataResponseBodyDataPerIntervalDataModule struct {
	HttpCodeDataPerInterval *DescribeDcdnDomainHttpCodeDataResponseBodyDataPerIntervalDataModuleHttpCodeDataPerInterval `json:"HttpCodeDataPerInterval,omitempty" xml:"HttpCodeDataPerInterval,omitempty" type:"Struct"`
	TimeStamp               *string                                                                                     `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
}

func (s DescribeDcdnDomainHttpCodeDataResponseBodyDataPerIntervalDataModule) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainHttpCodeDataResponseBodyDataPerIntervalDataModule) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainHttpCodeDataResponseBodyDataPerIntervalDataModule) SetHttpCodeDataPerInterval(v *DescribeDcdnDomainHttpCodeDataResponseBodyDataPerIntervalDataModuleHttpCodeDataPerInterval) *DescribeDcdnDomainHttpCodeDataResponseBodyDataPerIntervalDataModule {
	s.HttpCodeDataPerInterval = v
	return s
}

func (s *DescribeDcdnDomainHttpCodeDataResponseBodyDataPerIntervalDataModule) SetTimeStamp(v string) *DescribeDcdnDomainHttpCodeDataResponseBodyDataPerIntervalDataModule {
	s.TimeStamp = &v
	return s
}

type DescribeDcdnDomainHttpCodeDataResponseBodyDataPerIntervalDataModuleHttpCodeDataPerInterval struct {
	HttpCodeDataModule []*DescribeDcdnDomainHttpCodeDataResponseBodyDataPerIntervalDataModuleHttpCodeDataPerIntervalHttpCodeDataModule `json:"HttpCodeDataModule,omitempty" xml:"HttpCodeDataModule,omitempty" type:"Repeated"`
}

func (s DescribeDcdnDomainHttpCodeDataResponseBodyDataPerIntervalDataModuleHttpCodeDataPerInterval) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainHttpCodeDataResponseBodyDataPerIntervalDataModuleHttpCodeDataPerInterval) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainHttpCodeDataResponseBodyDataPerIntervalDataModuleHttpCodeDataPerInterval) SetHttpCodeDataModule(v []*DescribeDcdnDomainHttpCodeDataResponseBodyDataPerIntervalDataModuleHttpCodeDataPerIntervalHttpCodeDataModule) *DescribeDcdnDomainHttpCodeDataResponseBodyDataPerIntervalDataModuleHttpCodeDataPerInterval {
	s.HttpCodeDataModule = v
	return s
}

type DescribeDcdnDomainHttpCodeDataResponseBodyDataPerIntervalDataModuleHttpCodeDataPerIntervalHttpCodeDataModule struct {
	Code       *int32   `json:"Code,omitempty" xml:"Code,omitempty"`
	Count      *float32 `json:"Count,omitempty" xml:"Count,omitempty"`
	Proportion *float32 `json:"Proportion,omitempty" xml:"Proportion,omitempty"`
}

func (s DescribeDcdnDomainHttpCodeDataResponseBodyDataPerIntervalDataModuleHttpCodeDataPerIntervalHttpCodeDataModule) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainHttpCodeDataResponseBodyDataPerIntervalDataModuleHttpCodeDataPerIntervalHttpCodeDataModule) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainHttpCodeDataResponseBodyDataPerIntervalDataModuleHttpCodeDataPerIntervalHttpCodeDataModule) SetCode(v int32) *DescribeDcdnDomainHttpCodeDataResponseBodyDataPerIntervalDataModuleHttpCodeDataPerIntervalHttpCodeDataModule {
	s.Code = &v
	return s
}

func (s *DescribeDcdnDomainHttpCodeDataResponseBodyDataPerIntervalDataModuleHttpCodeDataPerIntervalHttpCodeDataModule) SetCount(v float32) *DescribeDcdnDomainHttpCodeDataResponseBodyDataPerIntervalDataModuleHttpCodeDataPerIntervalHttpCodeDataModule {
	s.Count = &v
	return s
}

func (s *DescribeDcdnDomainHttpCodeDataResponseBodyDataPerIntervalDataModuleHttpCodeDataPerIntervalHttpCodeDataModule) SetProportion(v float32) *DescribeDcdnDomainHttpCodeDataResponseBodyDataPerIntervalDataModuleHttpCodeDataPerIntervalHttpCodeDataModule {
	s.Proportion = &v
	return s
}

type DescribeDcdnDomainHttpCodeDataResponse struct {
	Headers map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDcdnDomainHttpCodeDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDcdnDomainHttpCodeDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainHttpCodeDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainHttpCodeDataResponse) SetHeaders(v map[string]*string) *DescribeDcdnDomainHttpCodeDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnDomainHttpCodeDataResponse) SetBody(v *DescribeDcdnDomainHttpCodeDataResponseBody) *DescribeDcdnDomainHttpCodeDataResponse {
	s.Body = v
	return s
}

type DescribeDcdnDomainIpaBpsDataRequest struct {
	DomainName     *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EndTime        *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	FixTimeGap     *string `json:"FixTimeGap,omitempty" xml:"FixTimeGap,omitempty"`
	Interval       *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	IspNameEn      *string `json:"IspNameEn,omitempty" xml:"IspNameEn,omitempty"`
	LocationNameEn *string `json:"LocationNameEn,omitempty" xml:"LocationNameEn,omitempty"`
	OwnerId        *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	StartTime      *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	TimeMerge      *string `json:"TimeMerge,omitempty" xml:"TimeMerge,omitempty"`
}

func (s DescribeDcdnDomainIpaBpsDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainIpaBpsDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainIpaBpsDataRequest) SetDomainName(v string) *DescribeDcdnDomainIpaBpsDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDcdnDomainIpaBpsDataRequest) SetEndTime(v string) *DescribeDcdnDomainIpaBpsDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDcdnDomainIpaBpsDataRequest) SetFixTimeGap(v string) *DescribeDcdnDomainIpaBpsDataRequest {
	s.FixTimeGap = &v
	return s
}

func (s *DescribeDcdnDomainIpaBpsDataRequest) SetInterval(v string) *DescribeDcdnDomainIpaBpsDataRequest {
	s.Interval = &v
	return s
}

func (s *DescribeDcdnDomainIpaBpsDataRequest) SetIspNameEn(v string) *DescribeDcdnDomainIpaBpsDataRequest {
	s.IspNameEn = &v
	return s
}

func (s *DescribeDcdnDomainIpaBpsDataRequest) SetLocationNameEn(v string) *DescribeDcdnDomainIpaBpsDataRequest {
	s.LocationNameEn = &v
	return s
}

func (s *DescribeDcdnDomainIpaBpsDataRequest) SetOwnerId(v int64) *DescribeDcdnDomainIpaBpsDataRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDcdnDomainIpaBpsDataRequest) SetStartTime(v string) *DescribeDcdnDomainIpaBpsDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDcdnDomainIpaBpsDataRequest) SetTimeMerge(v string) *DescribeDcdnDomainIpaBpsDataRequest {
	s.TimeMerge = &v
	return s
}

type DescribeDcdnDomainIpaBpsDataResponseBody struct {
	BpsDataPerInterval *DescribeDcdnDomainIpaBpsDataResponseBodyBpsDataPerInterval `json:"BpsDataPerInterval,omitempty" xml:"BpsDataPerInterval,omitempty" type:"Struct"`
	DataInterval       *string                                                     `json:"DataInterval,omitempty" xml:"DataInterval,omitempty"`
	DomainName         *string                                                     `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EndTime            *string                                                     `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	RequestId          *string                                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StartTime          *string                                                     `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDcdnDomainIpaBpsDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainIpaBpsDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainIpaBpsDataResponseBody) SetBpsDataPerInterval(v *DescribeDcdnDomainIpaBpsDataResponseBodyBpsDataPerInterval) *DescribeDcdnDomainIpaBpsDataResponseBody {
	s.BpsDataPerInterval = v
	return s
}

func (s *DescribeDcdnDomainIpaBpsDataResponseBody) SetDataInterval(v string) *DescribeDcdnDomainIpaBpsDataResponseBody {
	s.DataInterval = &v
	return s
}

func (s *DescribeDcdnDomainIpaBpsDataResponseBody) SetDomainName(v string) *DescribeDcdnDomainIpaBpsDataResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeDcdnDomainIpaBpsDataResponseBody) SetEndTime(v string) *DescribeDcdnDomainIpaBpsDataResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeDcdnDomainIpaBpsDataResponseBody) SetRequestId(v string) *DescribeDcdnDomainIpaBpsDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDcdnDomainIpaBpsDataResponseBody) SetStartTime(v string) *DescribeDcdnDomainIpaBpsDataResponseBody {
	s.StartTime = &v
	return s
}

type DescribeDcdnDomainIpaBpsDataResponseBodyBpsDataPerInterval struct {
	DataModule []*DescribeDcdnDomainIpaBpsDataResponseBodyBpsDataPerIntervalDataModule `json:"DataModule,omitempty" xml:"DataModule,omitempty" type:"Repeated"`
}

func (s DescribeDcdnDomainIpaBpsDataResponseBodyBpsDataPerInterval) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainIpaBpsDataResponseBodyBpsDataPerInterval) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainIpaBpsDataResponseBodyBpsDataPerInterval) SetDataModule(v []*DescribeDcdnDomainIpaBpsDataResponseBodyBpsDataPerIntervalDataModule) *DescribeDcdnDomainIpaBpsDataResponseBodyBpsDataPerInterval {
	s.DataModule = v
	return s
}

type DescribeDcdnDomainIpaBpsDataResponseBodyBpsDataPerIntervalDataModule struct {
	IpaBps    *float32 `json:"IpaBps,omitempty" xml:"IpaBps,omitempty"`
	TimeStamp *string  `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
}

func (s DescribeDcdnDomainIpaBpsDataResponseBodyBpsDataPerIntervalDataModule) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainIpaBpsDataResponseBodyBpsDataPerIntervalDataModule) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainIpaBpsDataResponseBodyBpsDataPerIntervalDataModule) SetIpaBps(v float32) *DescribeDcdnDomainIpaBpsDataResponseBodyBpsDataPerIntervalDataModule {
	s.IpaBps = &v
	return s
}

func (s *DescribeDcdnDomainIpaBpsDataResponseBodyBpsDataPerIntervalDataModule) SetTimeStamp(v string) *DescribeDcdnDomainIpaBpsDataResponseBodyBpsDataPerIntervalDataModule {
	s.TimeStamp = &v
	return s
}

type DescribeDcdnDomainIpaBpsDataResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDcdnDomainIpaBpsDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDcdnDomainIpaBpsDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainIpaBpsDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainIpaBpsDataResponse) SetHeaders(v map[string]*string) *DescribeDcdnDomainIpaBpsDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnDomainIpaBpsDataResponse) SetBody(v *DescribeDcdnDomainIpaBpsDataResponseBody) *DescribeDcdnDomainIpaBpsDataResponse {
	s.Body = v
	return s
}

type DescribeDcdnDomainIpaTrafficDataRequest struct {
	DomainName     *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EndTime        *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	FixTimeGap     *string `json:"FixTimeGap,omitempty" xml:"FixTimeGap,omitempty"`
	Interval       *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	IspNameEn      *string `json:"IspNameEn,omitempty" xml:"IspNameEn,omitempty"`
	LocationNameEn *string `json:"LocationNameEn,omitempty" xml:"LocationNameEn,omitempty"`
	OwnerId        *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	StartTime      *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	TimeMerge      *string `json:"TimeMerge,omitempty" xml:"TimeMerge,omitempty"`
}

func (s DescribeDcdnDomainIpaTrafficDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainIpaTrafficDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainIpaTrafficDataRequest) SetDomainName(v string) *DescribeDcdnDomainIpaTrafficDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDcdnDomainIpaTrafficDataRequest) SetEndTime(v string) *DescribeDcdnDomainIpaTrafficDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDcdnDomainIpaTrafficDataRequest) SetFixTimeGap(v string) *DescribeDcdnDomainIpaTrafficDataRequest {
	s.FixTimeGap = &v
	return s
}

func (s *DescribeDcdnDomainIpaTrafficDataRequest) SetInterval(v string) *DescribeDcdnDomainIpaTrafficDataRequest {
	s.Interval = &v
	return s
}

func (s *DescribeDcdnDomainIpaTrafficDataRequest) SetIspNameEn(v string) *DescribeDcdnDomainIpaTrafficDataRequest {
	s.IspNameEn = &v
	return s
}

func (s *DescribeDcdnDomainIpaTrafficDataRequest) SetLocationNameEn(v string) *DescribeDcdnDomainIpaTrafficDataRequest {
	s.LocationNameEn = &v
	return s
}

func (s *DescribeDcdnDomainIpaTrafficDataRequest) SetOwnerId(v int64) *DescribeDcdnDomainIpaTrafficDataRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDcdnDomainIpaTrafficDataRequest) SetStartTime(v string) *DescribeDcdnDomainIpaTrafficDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDcdnDomainIpaTrafficDataRequest) SetTimeMerge(v string) *DescribeDcdnDomainIpaTrafficDataRequest {
	s.TimeMerge = &v
	return s
}

type DescribeDcdnDomainIpaTrafficDataResponseBody struct {
	DataInterval           *string                                                             `json:"DataInterval,omitempty" xml:"DataInterval,omitempty"`
	DomainName             *string                                                             `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EndTime                *string                                                             `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	RequestId              *string                                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StartTime              *string                                                             `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	TrafficDataPerInterval *DescribeDcdnDomainIpaTrafficDataResponseBodyTrafficDataPerInterval `json:"TrafficDataPerInterval,omitempty" xml:"TrafficDataPerInterval,omitempty" type:"Struct"`
}

func (s DescribeDcdnDomainIpaTrafficDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainIpaTrafficDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainIpaTrafficDataResponseBody) SetDataInterval(v string) *DescribeDcdnDomainIpaTrafficDataResponseBody {
	s.DataInterval = &v
	return s
}

func (s *DescribeDcdnDomainIpaTrafficDataResponseBody) SetDomainName(v string) *DescribeDcdnDomainIpaTrafficDataResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeDcdnDomainIpaTrafficDataResponseBody) SetEndTime(v string) *DescribeDcdnDomainIpaTrafficDataResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeDcdnDomainIpaTrafficDataResponseBody) SetRequestId(v string) *DescribeDcdnDomainIpaTrafficDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDcdnDomainIpaTrafficDataResponseBody) SetStartTime(v string) *DescribeDcdnDomainIpaTrafficDataResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeDcdnDomainIpaTrafficDataResponseBody) SetTrafficDataPerInterval(v *DescribeDcdnDomainIpaTrafficDataResponseBodyTrafficDataPerInterval) *DescribeDcdnDomainIpaTrafficDataResponseBody {
	s.TrafficDataPerInterval = v
	return s
}

type DescribeDcdnDomainIpaTrafficDataResponseBodyTrafficDataPerInterval struct {
	DataModule []*DescribeDcdnDomainIpaTrafficDataResponseBodyTrafficDataPerIntervalDataModule `json:"DataModule,omitempty" xml:"DataModule,omitempty" type:"Repeated"`
}

func (s DescribeDcdnDomainIpaTrafficDataResponseBodyTrafficDataPerInterval) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainIpaTrafficDataResponseBodyTrafficDataPerInterval) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainIpaTrafficDataResponseBodyTrafficDataPerInterval) SetDataModule(v []*DescribeDcdnDomainIpaTrafficDataResponseBodyTrafficDataPerIntervalDataModule) *DescribeDcdnDomainIpaTrafficDataResponseBodyTrafficDataPerInterval {
	s.DataModule = v
	return s
}

type DescribeDcdnDomainIpaTrafficDataResponseBodyTrafficDataPerIntervalDataModule struct {
	IpaTraffic *float32 `json:"IpaTraffic,omitempty" xml:"IpaTraffic,omitempty"`
	TimeStamp  *string  `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
}

func (s DescribeDcdnDomainIpaTrafficDataResponseBodyTrafficDataPerIntervalDataModule) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainIpaTrafficDataResponseBodyTrafficDataPerIntervalDataModule) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainIpaTrafficDataResponseBodyTrafficDataPerIntervalDataModule) SetIpaTraffic(v float32) *DescribeDcdnDomainIpaTrafficDataResponseBodyTrafficDataPerIntervalDataModule {
	s.IpaTraffic = &v
	return s
}

func (s *DescribeDcdnDomainIpaTrafficDataResponseBodyTrafficDataPerIntervalDataModule) SetTimeStamp(v string) *DescribeDcdnDomainIpaTrafficDataResponseBodyTrafficDataPerIntervalDataModule {
	s.TimeStamp = &v
	return s
}

type DescribeDcdnDomainIpaTrafficDataResponse struct {
	Headers map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDcdnDomainIpaTrafficDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDcdnDomainIpaTrafficDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainIpaTrafficDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainIpaTrafficDataResponse) SetHeaders(v map[string]*string) *DescribeDcdnDomainIpaTrafficDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnDomainIpaTrafficDataResponse) SetBody(v *DescribeDcdnDomainIpaTrafficDataResponseBody) *DescribeDcdnDomainIpaTrafficDataResponse {
	s.Body = v
	return s
}

type DescribeDcdnDomainIspDataRequest struct {
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EndTime    *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	StartTime  *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDcdnDomainIspDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainIspDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainIspDataRequest) SetDomainName(v string) *DescribeDcdnDomainIspDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDcdnDomainIspDataRequest) SetEndTime(v string) *DescribeDcdnDomainIspDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDcdnDomainIspDataRequest) SetOwnerId(v int64) *DescribeDcdnDomainIspDataRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDcdnDomainIspDataRequest) SetStartTime(v string) *DescribeDcdnDomainIspDataRequest {
	s.StartTime = &v
	return s
}

type DescribeDcdnDomainIspDataResponseBody struct {
	DataInterval *string                                     `json:"DataInterval,omitempty" xml:"DataInterval,omitempty"`
	DomainName   *string                                     `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EndTime      *string                                     `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	RequestId    *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StartTime    *string                                     `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Value        *DescribeDcdnDomainIspDataResponseBodyValue `json:"Value,omitempty" xml:"Value,omitempty" type:"Struct"`
}

func (s DescribeDcdnDomainIspDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainIspDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainIspDataResponseBody) SetDataInterval(v string) *DescribeDcdnDomainIspDataResponseBody {
	s.DataInterval = &v
	return s
}

func (s *DescribeDcdnDomainIspDataResponseBody) SetDomainName(v string) *DescribeDcdnDomainIspDataResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeDcdnDomainIspDataResponseBody) SetEndTime(v string) *DescribeDcdnDomainIspDataResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeDcdnDomainIspDataResponseBody) SetRequestId(v string) *DescribeDcdnDomainIspDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDcdnDomainIspDataResponseBody) SetStartTime(v string) *DescribeDcdnDomainIspDataResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeDcdnDomainIspDataResponseBody) SetValue(v *DescribeDcdnDomainIspDataResponseBodyValue) *DescribeDcdnDomainIspDataResponseBody {
	s.Value = v
	return s
}

type DescribeDcdnDomainIspDataResponseBodyValue struct {
	IspProportionData []*DescribeDcdnDomainIspDataResponseBodyValueIspProportionData `json:"IspProportionData,omitempty" xml:"IspProportionData,omitempty" type:"Repeated"`
}

func (s DescribeDcdnDomainIspDataResponseBodyValue) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainIspDataResponseBodyValue) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainIspDataResponseBodyValue) SetIspProportionData(v []*DescribeDcdnDomainIspDataResponseBodyValueIspProportionData) *DescribeDcdnDomainIspDataResponseBodyValue {
	s.IspProportionData = v
	return s
}

type DescribeDcdnDomainIspDataResponseBodyValueIspProportionData struct {
	AvgObjectSize   *string `json:"AvgObjectSize,omitempty" xml:"AvgObjectSize,omitempty"`
	AvgResponseRate *string `json:"AvgResponseRate,omitempty" xml:"AvgResponseRate,omitempty"`
	AvgResponseTime *string `json:"AvgResponseTime,omitempty" xml:"AvgResponseTime,omitempty"`
	Bps             *string `json:"Bps,omitempty" xml:"Bps,omitempty"`
	BytesProportion *string `json:"BytesProportion,omitempty" xml:"BytesProportion,omitempty"`
	Isp             *string `json:"Isp,omitempty" xml:"Isp,omitempty"`
	IspEname        *string `json:"IspEname,omitempty" xml:"IspEname,omitempty"`
	Proportion      *string `json:"Proportion,omitempty" xml:"Proportion,omitempty"`
	Qps             *string `json:"Qps,omitempty" xml:"Qps,omitempty"`
	TotalBytes      *string `json:"TotalBytes,omitempty" xml:"TotalBytes,omitempty"`
	TotalQuery      *string `json:"TotalQuery,omitempty" xml:"TotalQuery,omitempty"`
}

func (s DescribeDcdnDomainIspDataResponseBodyValueIspProportionData) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainIspDataResponseBodyValueIspProportionData) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainIspDataResponseBodyValueIspProportionData) SetAvgObjectSize(v string) *DescribeDcdnDomainIspDataResponseBodyValueIspProportionData {
	s.AvgObjectSize = &v
	return s
}

func (s *DescribeDcdnDomainIspDataResponseBodyValueIspProportionData) SetAvgResponseRate(v string) *DescribeDcdnDomainIspDataResponseBodyValueIspProportionData {
	s.AvgResponseRate = &v
	return s
}

func (s *DescribeDcdnDomainIspDataResponseBodyValueIspProportionData) SetAvgResponseTime(v string) *DescribeDcdnDomainIspDataResponseBodyValueIspProportionData {
	s.AvgResponseTime = &v
	return s
}

func (s *DescribeDcdnDomainIspDataResponseBodyValueIspProportionData) SetBps(v string) *DescribeDcdnDomainIspDataResponseBodyValueIspProportionData {
	s.Bps = &v
	return s
}

func (s *DescribeDcdnDomainIspDataResponseBodyValueIspProportionData) SetBytesProportion(v string) *DescribeDcdnDomainIspDataResponseBodyValueIspProportionData {
	s.BytesProportion = &v
	return s
}

func (s *DescribeDcdnDomainIspDataResponseBodyValueIspProportionData) SetIsp(v string) *DescribeDcdnDomainIspDataResponseBodyValueIspProportionData {
	s.Isp = &v
	return s
}

func (s *DescribeDcdnDomainIspDataResponseBodyValueIspProportionData) SetIspEname(v string) *DescribeDcdnDomainIspDataResponseBodyValueIspProportionData {
	s.IspEname = &v
	return s
}

func (s *DescribeDcdnDomainIspDataResponseBodyValueIspProportionData) SetProportion(v string) *DescribeDcdnDomainIspDataResponseBodyValueIspProportionData {
	s.Proportion = &v
	return s
}

func (s *DescribeDcdnDomainIspDataResponseBodyValueIspProportionData) SetQps(v string) *DescribeDcdnDomainIspDataResponseBodyValueIspProportionData {
	s.Qps = &v
	return s
}

func (s *DescribeDcdnDomainIspDataResponseBodyValueIspProportionData) SetTotalBytes(v string) *DescribeDcdnDomainIspDataResponseBodyValueIspProportionData {
	s.TotalBytes = &v
	return s
}

func (s *DescribeDcdnDomainIspDataResponseBodyValueIspProportionData) SetTotalQuery(v string) *DescribeDcdnDomainIspDataResponseBodyValueIspProportionData {
	s.TotalQuery = &v
	return s
}

type DescribeDcdnDomainIspDataResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDcdnDomainIspDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDcdnDomainIspDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainIspDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainIspDataResponse) SetHeaders(v map[string]*string) *DescribeDcdnDomainIspDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnDomainIspDataResponse) SetBody(v *DescribeDcdnDomainIspDataResponseBody) *DescribeDcdnDomainIspDataResponse {
	s.Body = v
	return s
}

type DescribeDcdnDomainLogRequest struct {
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EndTime    *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNumber *int64  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	StartTime  *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDcdnDomainLogRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainLogRequest) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainLogRequest) SetDomainName(v string) *DescribeDcdnDomainLogRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDcdnDomainLogRequest) SetEndTime(v string) *DescribeDcdnDomainLogRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDcdnDomainLogRequest) SetOwnerId(v int64) *DescribeDcdnDomainLogRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDcdnDomainLogRequest) SetPageNumber(v int64) *DescribeDcdnDomainLogRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDcdnDomainLogRequest) SetPageSize(v int64) *DescribeDcdnDomainLogRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDcdnDomainLogRequest) SetStartTime(v string) *DescribeDcdnDomainLogRequest {
	s.StartTime = &v
	return s
}

type DescribeDcdnDomainLogResponseBody struct {
	DomainLogDetails *DescribeDcdnDomainLogResponseBodyDomainLogDetails `json:"DomainLogDetails,omitempty" xml:"DomainLogDetails,omitempty" type:"Struct"`
	DomainName       *string                                            `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	RequestId        *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDcdnDomainLogResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainLogResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainLogResponseBody) SetDomainLogDetails(v *DescribeDcdnDomainLogResponseBodyDomainLogDetails) *DescribeDcdnDomainLogResponseBody {
	s.DomainLogDetails = v
	return s
}

func (s *DescribeDcdnDomainLogResponseBody) SetDomainName(v string) *DescribeDcdnDomainLogResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeDcdnDomainLogResponseBody) SetRequestId(v string) *DescribeDcdnDomainLogResponseBody {
	s.RequestId = &v
	return s
}

type DescribeDcdnDomainLogResponseBodyDomainLogDetails struct {
	DomainLogDetail []*DescribeDcdnDomainLogResponseBodyDomainLogDetailsDomainLogDetail `json:"DomainLogDetail,omitempty" xml:"DomainLogDetail,omitempty" type:"Repeated"`
}

func (s DescribeDcdnDomainLogResponseBodyDomainLogDetails) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainLogResponseBodyDomainLogDetails) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainLogResponseBodyDomainLogDetails) SetDomainLogDetail(v []*DescribeDcdnDomainLogResponseBodyDomainLogDetailsDomainLogDetail) *DescribeDcdnDomainLogResponseBodyDomainLogDetails {
	s.DomainLogDetail = v
	return s
}

type DescribeDcdnDomainLogResponseBodyDomainLogDetailsDomainLogDetail struct {
	LogCount  *int64                                                                     `json:"LogCount,omitempty" xml:"LogCount,omitempty"`
	LogInfos  *DescribeDcdnDomainLogResponseBodyDomainLogDetailsDomainLogDetailLogInfos  `json:"LogInfos,omitempty" xml:"LogInfos,omitempty" type:"Struct"`
	PageInfos *DescribeDcdnDomainLogResponseBodyDomainLogDetailsDomainLogDetailPageInfos `json:"PageInfos,omitempty" xml:"PageInfos,omitempty" type:"Struct"`
}

func (s DescribeDcdnDomainLogResponseBodyDomainLogDetailsDomainLogDetail) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainLogResponseBodyDomainLogDetailsDomainLogDetail) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainLogResponseBodyDomainLogDetailsDomainLogDetail) SetLogCount(v int64) *DescribeDcdnDomainLogResponseBodyDomainLogDetailsDomainLogDetail {
	s.LogCount = &v
	return s
}

func (s *DescribeDcdnDomainLogResponseBodyDomainLogDetailsDomainLogDetail) SetLogInfos(v *DescribeDcdnDomainLogResponseBodyDomainLogDetailsDomainLogDetailLogInfos) *DescribeDcdnDomainLogResponseBodyDomainLogDetailsDomainLogDetail {
	s.LogInfos = v
	return s
}

func (s *DescribeDcdnDomainLogResponseBodyDomainLogDetailsDomainLogDetail) SetPageInfos(v *DescribeDcdnDomainLogResponseBodyDomainLogDetailsDomainLogDetailPageInfos) *DescribeDcdnDomainLogResponseBodyDomainLogDetailsDomainLogDetail {
	s.PageInfos = v
	return s
}

type DescribeDcdnDomainLogResponseBodyDomainLogDetailsDomainLogDetailLogInfos struct {
	LogInfoDetail []*DescribeDcdnDomainLogResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail `json:"LogInfoDetail,omitempty" xml:"LogInfoDetail,omitempty" type:"Repeated"`
}

func (s DescribeDcdnDomainLogResponseBodyDomainLogDetailsDomainLogDetailLogInfos) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainLogResponseBodyDomainLogDetailsDomainLogDetailLogInfos) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainLogResponseBodyDomainLogDetailsDomainLogDetailLogInfos) SetLogInfoDetail(v []*DescribeDcdnDomainLogResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail) *DescribeDcdnDomainLogResponseBodyDomainLogDetailsDomainLogDetailLogInfos {
	s.LogInfoDetail = v
	return s
}

type DescribeDcdnDomainLogResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail struct {
	EndTime   *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	LogName   *string `json:"LogName,omitempty" xml:"LogName,omitempty"`
	LogPath   *string `json:"LogPath,omitempty" xml:"LogPath,omitempty"`
	LogSize   *int64  `json:"LogSize,omitempty" xml:"LogSize,omitempty"`
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDcdnDomainLogResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainLogResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainLogResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail) SetEndTime(v string) *DescribeDcdnDomainLogResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail {
	s.EndTime = &v
	return s
}

func (s *DescribeDcdnDomainLogResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail) SetLogName(v string) *DescribeDcdnDomainLogResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail {
	s.LogName = &v
	return s
}

func (s *DescribeDcdnDomainLogResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail) SetLogPath(v string) *DescribeDcdnDomainLogResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail {
	s.LogPath = &v
	return s
}

func (s *DescribeDcdnDomainLogResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail) SetLogSize(v int64) *DescribeDcdnDomainLogResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail {
	s.LogSize = &v
	return s
}

func (s *DescribeDcdnDomainLogResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail) SetStartTime(v string) *DescribeDcdnDomainLogResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail {
	s.StartTime = &v
	return s
}

type DescribeDcdnDomainLogResponseBodyDomainLogDetailsDomainLogDetailPageInfos struct {
	PageIndex *int64 `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	PageSize  *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Total     *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeDcdnDomainLogResponseBodyDomainLogDetailsDomainLogDetailPageInfos) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainLogResponseBodyDomainLogDetailsDomainLogDetailPageInfos) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainLogResponseBodyDomainLogDetailsDomainLogDetailPageInfos) SetPageIndex(v int64) *DescribeDcdnDomainLogResponseBodyDomainLogDetailsDomainLogDetailPageInfos {
	s.PageIndex = &v
	return s
}

func (s *DescribeDcdnDomainLogResponseBodyDomainLogDetailsDomainLogDetailPageInfos) SetPageSize(v int64) *DescribeDcdnDomainLogResponseBodyDomainLogDetailsDomainLogDetailPageInfos {
	s.PageSize = &v
	return s
}

func (s *DescribeDcdnDomainLogResponseBodyDomainLogDetailsDomainLogDetailPageInfos) SetTotal(v int64) *DescribeDcdnDomainLogResponseBodyDomainLogDetailsDomainLogDetailPageInfos {
	s.Total = &v
	return s
}

type DescribeDcdnDomainLogResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDcdnDomainLogResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDcdnDomainLogResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainLogResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainLogResponse) SetHeaders(v map[string]*string) *DescribeDcdnDomainLogResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnDomainLogResponse) SetBody(v *DescribeDcdnDomainLogResponseBody) *DescribeDcdnDomainLogResponse {
	s.Body = v
	return s
}

type DescribeDcdnDomainMultiUsageDataRequest struct {
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EndTime    *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	StartTime  *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDcdnDomainMultiUsageDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainMultiUsageDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainMultiUsageDataRequest) SetDomainName(v string) *DescribeDcdnDomainMultiUsageDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDcdnDomainMultiUsageDataRequest) SetEndTime(v string) *DescribeDcdnDomainMultiUsageDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDcdnDomainMultiUsageDataRequest) SetOwnerId(v int64) *DescribeDcdnDomainMultiUsageDataRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDcdnDomainMultiUsageDataRequest) SetStartTime(v string) *DescribeDcdnDomainMultiUsageDataRequest {
	s.StartTime = &v
	return s
}

type DescribeDcdnDomainMultiUsageDataResponseBody struct {
	EndTime            *string                                                         `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	RequestId          *string                                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RequestPerInterval *DescribeDcdnDomainMultiUsageDataResponseBodyRequestPerInterval `json:"RequestPerInterval,omitempty" xml:"RequestPerInterval,omitempty" type:"Struct"`
	StartTime          *string                                                         `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	TrafficPerInterval *DescribeDcdnDomainMultiUsageDataResponseBodyTrafficPerInterval `json:"TrafficPerInterval,omitempty" xml:"TrafficPerInterval,omitempty" type:"Struct"`
}

func (s DescribeDcdnDomainMultiUsageDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainMultiUsageDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainMultiUsageDataResponseBody) SetEndTime(v string) *DescribeDcdnDomainMultiUsageDataResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeDcdnDomainMultiUsageDataResponseBody) SetRequestId(v string) *DescribeDcdnDomainMultiUsageDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDcdnDomainMultiUsageDataResponseBody) SetRequestPerInterval(v *DescribeDcdnDomainMultiUsageDataResponseBodyRequestPerInterval) *DescribeDcdnDomainMultiUsageDataResponseBody {
	s.RequestPerInterval = v
	return s
}

func (s *DescribeDcdnDomainMultiUsageDataResponseBody) SetStartTime(v string) *DescribeDcdnDomainMultiUsageDataResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeDcdnDomainMultiUsageDataResponseBody) SetTrafficPerInterval(v *DescribeDcdnDomainMultiUsageDataResponseBodyTrafficPerInterval) *DescribeDcdnDomainMultiUsageDataResponseBody {
	s.TrafficPerInterval = v
	return s
}

type DescribeDcdnDomainMultiUsageDataResponseBodyRequestPerInterval struct {
	RequestDataModule []*DescribeDcdnDomainMultiUsageDataResponseBodyRequestPerIntervalRequestDataModule `json:"RequestDataModule,omitempty" xml:"RequestDataModule,omitempty" type:"Repeated"`
}

func (s DescribeDcdnDomainMultiUsageDataResponseBodyRequestPerInterval) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainMultiUsageDataResponseBodyRequestPerInterval) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainMultiUsageDataResponseBodyRequestPerInterval) SetRequestDataModule(v []*DescribeDcdnDomainMultiUsageDataResponseBodyRequestPerIntervalRequestDataModule) *DescribeDcdnDomainMultiUsageDataResponseBodyRequestPerInterval {
	s.RequestDataModule = v
	return s
}

type DescribeDcdnDomainMultiUsageDataResponseBodyRequestPerIntervalRequestDataModule struct {
	Domain    *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	Request   *int64  `json:"Request,omitempty" xml:"Request,omitempty"`
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
	Type      *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeDcdnDomainMultiUsageDataResponseBodyRequestPerIntervalRequestDataModule) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainMultiUsageDataResponseBodyRequestPerIntervalRequestDataModule) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainMultiUsageDataResponseBodyRequestPerIntervalRequestDataModule) SetDomain(v string) *DescribeDcdnDomainMultiUsageDataResponseBodyRequestPerIntervalRequestDataModule {
	s.Domain = &v
	return s
}

func (s *DescribeDcdnDomainMultiUsageDataResponseBodyRequestPerIntervalRequestDataModule) SetRequest(v int64) *DescribeDcdnDomainMultiUsageDataResponseBodyRequestPerIntervalRequestDataModule {
	s.Request = &v
	return s
}

func (s *DescribeDcdnDomainMultiUsageDataResponseBodyRequestPerIntervalRequestDataModule) SetTimeStamp(v string) *DescribeDcdnDomainMultiUsageDataResponseBodyRequestPerIntervalRequestDataModule {
	s.TimeStamp = &v
	return s
}

func (s *DescribeDcdnDomainMultiUsageDataResponseBodyRequestPerIntervalRequestDataModule) SetType(v string) *DescribeDcdnDomainMultiUsageDataResponseBodyRequestPerIntervalRequestDataModule {
	s.Type = &v
	return s
}

type DescribeDcdnDomainMultiUsageDataResponseBodyTrafficPerInterval struct {
	TrafficDataModule []*DescribeDcdnDomainMultiUsageDataResponseBodyTrafficPerIntervalTrafficDataModule `json:"TrafficDataModule,omitempty" xml:"TrafficDataModule,omitempty" type:"Repeated"`
}

func (s DescribeDcdnDomainMultiUsageDataResponseBodyTrafficPerInterval) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainMultiUsageDataResponseBodyTrafficPerInterval) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainMultiUsageDataResponseBodyTrafficPerInterval) SetTrafficDataModule(v []*DescribeDcdnDomainMultiUsageDataResponseBodyTrafficPerIntervalTrafficDataModule) *DescribeDcdnDomainMultiUsageDataResponseBodyTrafficPerInterval {
	s.TrafficDataModule = v
	return s
}

type DescribeDcdnDomainMultiUsageDataResponseBodyTrafficPerIntervalTrafficDataModule struct {
	Area      *string  `json:"Area,omitempty" xml:"Area,omitempty"`
	Bps       *float32 `json:"Bps,omitempty" xml:"Bps,omitempty"`
	Domain    *string  `json:"Domain,omitempty" xml:"Domain,omitempty"`
	TimeStamp *string  `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
	Type      *string  `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeDcdnDomainMultiUsageDataResponseBodyTrafficPerIntervalTrafficDataModule) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainMultiUsageDataResponseBodyTrafficPerIntervalTrafficDataModule) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainMultiUsageDataResponseBodyTrafficPerIntervalTrafficDataModule) SetArea(v string) *DescribeDcdnDomainMultiUsageDataResponseBodyTrafficPerIntervalTrafficDataModule {
	s.Area = &v
	return s
}

func (s *DescribeDcdnDomainMultiUsageDataResponseBodyTrafficPerIntervalTrafficDataModule) SetBps(v float32) *DescribeDcdnDomainMultiUsageDataResponseBodyTrafficPerIntervalTrafficDataModule {
	s.Bps = &v
	return s
}

func (s *DescribeDcdnDomainMultiUsageDataResponseBodyTrafficPerIntervalTrafficDataModule) SetDomain(v string) *DescribeDcdnDomainMultiUsageDataResponseBodyTrafficPerIntervalTrafficDataModule {
	s.Domain = &v
	return s
}

func (s *DescribeDcdnDomainMultiUsageDataResponseBodyTrafficPerIntervalTrafficDataModule) SetTimeStamp(v string) *DescribeDcdnDomainMultiUsageDataResponseBodyTrafficPerIntervalTrafficDataModule {
	s.TimeStamp = &v
	return s
}

func (s *DescribeDcdnDomainMultiUsageDataResponseBodyTrafficPerIntervalTrafficDataModule) SetType(v string) *DescribeDcdnDomainMultiUsageDataResponseBodyTrafficPerIntervalTrafficDataModule {
	s.Type = &v
	return s
}

type DescribeDcdnDomainMultiUsageDataResponse struct {
	Headers map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDcdnDomainMultiUsageDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDcdnDomainMultiUsageDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainMultiUsageDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainMultiUsageDataResponse) SetHeaders(v map[string]*string) *DescribeDcdnDomainMultiUsageDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnDomainMultiUsageDataResponse) SetBody(v *DescribeDcdnDomainMultiUsageDataResponseBody) *DescribeDcdnDomainMultiUsageDataResponse {
	s.Body = v
	return s
}

type DescribeDcdnDomainOriginBpsDataRequest struct {
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EndTime    *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Interval   *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	StartTime  *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDcdnDomainOriginBpsDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainOriginBpsDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainOriginBpsDataRequest) SetDomainName(v string) *DescribeDcdnDomainOriginBpsDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDcdnDomainOriginBpsDataRequest) SetEndTime(v string) *DescribeDcdnDomainOriginBpsDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDcdnDomainOriginBpsDataRequest) SetInterval(v string) *DescribeDcdnDomainOriginBpsDataRequest {
	s.Interval = &v
	return s
}

func (s *DescribeDcdnDomainOriginBpsDataRequest) SetOwnerId(v int64) *DescribeDcdnDomainOriginBpsDataRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDcdnDomainOriginBpsDataRequest) SetStartTime(v string) *DescribeDcdnDomainOriginBpsDataRequest {
	s.StartTime = &v
	return s
}

type DescribeDcdnDomainOriginBpsDataResponseBody struct {
	DataInterval             *string                                                              `json:"DataInterval,omitempty" xml:"DataInterval,omitempty"`
	DomainName               *string                                                              `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EndTime                  *string                                                              `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	OriginBpsDataPerInterval *DescribeDcdnDomainOriginBpsDataResponseBodyOriginBpsDataPerInterval `json:"OriginBpsDataPerInterval,omitempty" xml:"OriginBpsDataPerInterval,omitempty" type:"Struct"`
	RequestId                *string                                                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StartTime                *string                                                              `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDcdnDomainOriginBpsDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainOriginBpsDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainOriginBpsDataResponseBody) SetDataInterval(v string) *DescribeDcdnDomainOriginBpsDataResponseBody {
	s.DataInterval = &v
	return s
}

func (s *DescribeDcdnDomainOriginBpsDataResponseBody) SetDomainName(v string) *DescribeDcdnDomainOriginBpsDataResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeDcdnDomainOriginBpsDataResponseBody) SetEndTime(v string) *DescribeDcdnDomainOriginBpsDataResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeDcdnDomainOriginBpsDataResponseBody) SetOriginBpsDataPerInterval(v *DescribeDcdnDomainOriginBpsDataResponseBodyOriginBpsDataPerInterval) *DescribeDcdnDomainOriginBpsDataResponseBody {
	s.OriginBpsDataPerInterval = v
	return s
}

func (s *DescribeDcdnDomainOriginBpsDataResponseBody) SetRequestId(v string) *DescribeDcdnDomainOriginBpsDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDcdnDomainOriginBpsDataResponseBody) SetStartTime(v string) *DescribeDcdnDomainOriginBpsDataResponseBody {
	s.StartTime = &v
	return s
}

type DescribeDcdnDomainOriginBpsDataResponseBodyOriginBpsDataPerInterval struct {
	DataModule []*DescribeDcdnDomainOriginBpsDataResponseBodyOriginBpsDataPerIntervalDataModule `json:"DataModule,omitempty" xml:"DataModule,omitempty" type:"Repeated"`
}

func (s DescribeDcdnDomainOriginBpsDataResponseBodyOriginBpsDataPerInterval) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainOriginBpsDataResponseBodyOriginBpsDataPerInterval) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainOriginBpsDataResponseBodyOriginBpsDataPerInterval) SetDataModule(v []*DescribeDcdnDomainOriginBpsDataResponseBodyOriginBpsDataPerIntervalDataModule) *DescribeDcdnDomainOriginBpsDataResponseBodyOriginBpsDataPerInterval {
	s.DataModule = v
	return s
}

type DescribeDcdnDomainOriginBpsDataResponseBodyOriginBpsDataPerIntervalDataModule struct {
	DynamicHttpOriginBps  *float32 `json:"DynamicHttpOriginBps,omitempty" xml:"DynamicHttpOriginBps,omitempty"`
	DynamicHttpsOriginBps *float32 `json:"DynamicHttpsOriginBps,omitempty" xml:"DynamicHttpsOriginBps,omitempty"`
	OriginBps             *float32 `json:"OriginBps,omitempty" xml:"OriginBps,omitempty"`
	StaticHttpOriginBps   *float32 `json:"StaticHttpOriginBps,omitempty" xml:"StaticHttpOriginBps,omitempty"`
	StaticHttpsOriginBps  *float32 `json:"StaticHttpsOriginBps,omitempty" xml:"StaticHttpsOriginBps,omitempty"`
	TimeStamp             *string  `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
}

func (s DescribeDcdnDomainOriginBpsDataResponseBodyOriginBpsDataPerIntervalDataModule) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainOriginBpsDataResponseBodyOriginBpsDataPerIntervalDataModule) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainOriginBpsDataResponseBodyOriginBpsDataPerIntervalDataModule) SetDynamicHttpOriginBps(v float32) *DescribeDcdnDomainOriginBpsDataResponseBodyOriginBpsDataPerIntervalDataModule {
	s.DynamicHttpOriginBps = &v
	return s
}

func (s *DescribeDcdnDomainOriginBpsDataResponseBodyOriginBpsDataPerIntervalDataModule) SetDynamicHttpsOriginBps(v float32) *DescribeDcdnDomainOriginBpsDataResponseBodyOriginBpsDataPerIntervalDataModule {
	s.DynamicHttpsOriginBps = &v
	return s
}

func (s *DescribeDcdnDomainOriginBpsDataResponseBodyOriginBpsDataPerIntervalDataModule) SetOriginBps(v float32) *DescribeDcdnDomainOriginBpsDataResponseBodyOriginBpsDataPerIntervalDataModule {
	s.OriginBps = &v
	return s
}

func (s *DescribeDcdnDomainOriginBpsDataResponseBodyOriginBpsDataPerIntervalDataModule) SetStaticHttpOriginBps(v float32) *DescribeDcdnDomainOriginBpsDataResponseBodyOriginBpsDataPerIntervalDataModule {
	s.StaticHttpOriginBps = &v
	return s
}

func (s *DescribeDcdnDomainOriginBpsDataResponseBodyOriginBpsDataPerIntervalDataModule) SetStaticHttpsOriginBps(v float32) *DescribeDcdnDomainOriginBpsDataResponseBodyOriginBpsDataPerIntervalDataModule {
	s.StaticHttpsOriginBps = &v
	return s
}

func (s *DescribeDcdnDomainOriginBpsDataResponseBodyOriginBpsDataPerIntervalDataModule) SetTimeStamp(v string) *DescribeDcdnDomainOriginBpsDataResponseBodyOriginBpsDataPerIntervalDataModule {
	s.TimeStamp = &v
	return s
}

type DescribeDcdnDomainOriginBpsDataResponse struct {
	Headers map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDcdnDomainOriginBpsDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDcdnDomainOriginBpsDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainOriginBpsDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainOriginBpsDataResponse) SetHeaders(v map[string]*string) *DescribeDcdnDomainOriginBpsDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnDomainOriginBpsDataResponse) SetBody(v *DescribeDcdnDomainOriginBpsDataResponseBody) *DescribeDcdnDomainOriginBpsDataResponse {
	s.Body = v
	return s
}

type DescribeDcdnDomainOriginTrafficDataRequest struct {
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EndTime    *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Interval   *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	StartTime  *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDcdnDomainOriginTrafficDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainOriginTrafficDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainOriginTrafficDataRequest) SetDomainName(v string) *DescribeDcdnDomainOriginTrafficDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDcdnDomainOriginTrafficDataRequest) SetEndTime(v string) *DescribeDcdnDomainOriginTrafficDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDcdnDomainOriginTrafficDataRequest) SetInterval(v string) *DescribeDcdnDomainOriginTrafficDataRequest {
	s.Interval = &v
	return s
}

func (s *DescribeDcdnDomainOriginTrafficDataRequest) SetOwnerId(v int64) *DescribeDcdnDomainOriginTrafficDataRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDcdnDomainOriginTrafficDataRequest) SetStartTime(v string) *DescribeDcdnDomainOriginTrafficDataRequest {
	s.StartTime = &v
	return s
}

type DescribeDcdnDomainOriginTrafficDataResponseBody struct {
	DataInterval                 *string                                                                      `json:"DataInterval,omitempty" xml:"DataInterval,omitempty"`
	DomainName                   *string                                                                      `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EndTime                      *string                                                                      `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	OriginTrafficDataPerInterval *DescribeDcdnDomainOriginTrafficDataResponseBodyOriginTrafficDataPerInterval `json:"OriginTrafficDataPerInterval,omitempty" xml:"OriginTrafficDataPerInterval,omitempty" type:"Struct"`
	RequestId                    *string                                                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StartTime                    *string                                                                      `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDcdnDomainOriginTrafficDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainOriginTrafficDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainOriginTrafficDataResponseBody) SetDataInterval(v string) *DescribeDcdnDomainOriginTrafficDataResponseBody {
	s.DataInterval = &v
	return s
}

func (s *DescribeDcdnDomainOriginTrafficDataResponseBody) SetDomainName(v string) *DescribeDcdnDomainOriginTrafficDataResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeDcdnDomainOriginTrafficDataResponseBody) SetEndTime(v string) *DescribeDcdnDomainOriginTrafficDataResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeDcdnDomainOriginTrafficDataResponseBody) SetOriginTrafficDataPerInterval(v *DescribeDcdnDomainOriginTrafficDataResponseBodyOriginTrafficDataPerInterval) *DescribeDcdnDomainOriginTrafficDataResponseBody {
	s.OriginTrafficDataPerInterval = v
	return s
}

func (s *DescribeDcdnDomainOriginTrafficDataResponseBody) SetRequestId(v string) *DescribeDcdnDomainOriginTrafficDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDcdnDomainOriginTrafficDataResponseBody) SetStartTime(v string) *DescribeDcdnDomainOriginTrafficDataResponseBody {
	s.StartTime = &v
	return s
}

type DescribeDcdnDomainOriginTrafficDataResponseBodyOriginTrafficDataPerInterval struct {
	DataModule []*DescribeDcdnDomainOriginTrafficDataResponseBodyOriginTrafficDataPerIntervalDataModule `json:"DataModule,omitempty" xml:"DataModule,omitempty" type:"Repeated"`
}

func (s DescribeDcdnDomainOriginTrafficDataResponseBodyOriginTrafficDataPerInterval) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainOriginTrafficDataResponseBodyOriginTrafficDataPerInterval) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainOriginTrafficDataResponseBodyOriginTrafficDataPerInterval) SetDataModule(v []*DescribeDcdnDomainOriginTrafficDataResponseBodyOriginTrafficDataPerIntervalDataModule) *DescribeDcdnDomainOriginTrafficDataResponseBodyOriginTrafficDataPerInterval {
	s.DataModule = v
	return s
}

type DescribeDcdnDomainOriginTrafficDataResponseBodyOriginTrafficDataPerIntervalDataModule struct {
	DynamicHttpOriginTraffic  *float32 `json:"DynamicHttpOriginTraffic,omitempty" xml:"DynamicHttpOriginTraffic,omitempty"`
	DynamicHttpsOriginTraffic *float32 `json:"DynamicHttpsOriginTraffic,omitempty" xml:"DynamicHttpsOriginTraffic,omitempty"`
	OriginTraffic             *float32 `json:"OriginTraffic,omitempty" xml:"OriginTraffic,omitempty"`
	StaticHttpOriginTraffic   *float32 `json:"StaticHttpOriginTraffic,omitempty" xml:"StaticHttpOriginTraffic,omitempty"`
	StaticHttpsOriginTraffic  *float32 `json:"StaticHttpsOriginTraffic,omitempty" xml:"StaticHttpsOriginTraffic,omitempty"`
	TimeStamp                 *string  `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
}

func (s DescribeDcdnDomainOriginTrafficDataResponseBodyOriginTrafficDataPerIntervalDataModule) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainOriginTrafficDataResponseBodyOriginTrafficDataPerIntervalDataModule) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainOriginTrafficDataResponseBodyOriginTrafficDataPerIntervalDataModule) SetDynamicHttpOriginTraffic(v float32) *DescribeDcdnDomainOriginTrafficDataResponseBodyOriginTrafficDataPerIntervalDataModule {
	s.DynamicHttpOriginTraffic = &v
	return s
}

func (s *DescribeDcdnDomainOriginTrafficDataResponseBodyOriginTrafficDataPerIntervalDataModule) SetDynamicHttpsOriginTraffic(v float32) *DescribeDcdnDomainOriginTrafficDataResponseBodyOriginTrafficDataPerIntervalDataModule {
	s.DynamicHttpsOriginTraffic = &v
	return s
}

func (s *DescribeDcdnDomainOriginTrafficDataResponseBodyOriginTrafficDataPerIntervalDataModule) SetOriginTraffic(v float32) *DescribeDcdnDomainOriginTrafficDataResponseBodyOriginTrafficDataPerIntervalDataModule {
	s.OriginTraffic = &v
	return s
}

func (s *DescribeDcdnDomainOriginTrafficDataResponseBodyOriginTrafficDataPerIntervalDataModule) SetStaticHttpOriginTraffic(v float32) *DescribeDcdnDomainOriginTrafficDataResponseBodyOriginTrafficDataPerIntervalDataModule {
	s.StaticHttpOriginTraffic = &v
	return s
}

func (s *DescribeDcdnDomainOriginTrafficDataResponseBodyOriginTrafficDataPerIntervalDataModule) SetStaticHttpsOriginTraffic(v float32) *DescribeDcdnDomainOriginTrafficDataResponseBodyOriginTrafficDataPerIntervalDataModule {
	s.StaticHttpsOriginTraffic = &v
	return s
}

func (s *DescribeDcdnDomainOriginTrafficDataResponseBodyOriginTrafficDataPerIntervalDataModule) SetTimeStamp(v string) *DescribeDcdnDomainOriginTrafficDataResponseBodyOriginTrafficDataPerIntervalDataModule {
	s.TimeStamp = &v
	return s
}

type DescribeDcdnDomainOriginTrafficDataResponse struct {
	Headers map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDcdnDomainOriginTrafficDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDcdnDomainOriginTrafficDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainOriginTrafficDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainOriginTrafficDataResponse) SetHeaders(v map[string]*string) *DescribeDcdnDomainOriginTrafficDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnDomainOriginTrafficDataResponse) SetBody(v *DescribeDcdnDomainOriginTrafficDataResponseBody) *DescribeDcdnDomainOriginTrafficDataResponse {
	s.Body = v
	return s
}

type DescribeDcdnDomainPropertyRequest struct {
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s DescribeDcdnDomainPropertyRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainPropertyRequest) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainPropertyRequest) SetDomainName(v string) *DescribeDcdnDomainPropertyRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDcdnDomainPropertyRequest) SetOwnerId(v int64) *DescribeDcdnDomainPropertyRequest {
	s.OwnerId = &v
	return s
}

type DescribeDcdnDomainPropertyResponseBody struct {
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	Protocol   *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDcdnDomainPropertyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainPropertyResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainPropertyResponseBody) SetDomainName(v string) *DescribeDcdnDomainPropertyResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeDcdnDomainPropertyResponseBody) SetProtocol(v string) *DescribeDcdnDomainPropertyResponseBody {
	s.Protocol = &v
	return s
}

func (s *DescribeDcdnDomainPropertyResponseBody) SetRequestId(v string) *DescribeDcdnDomainPropertyResponseBody {
	s.RequestId = &v
	return s
}

type DescribeDcdnDomainPropertyResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDcdnDomainPropertyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDcdnDomainPropertyResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainPropertyResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainPropertyResponse) SetHeaders(v map[string]*string) *DescribeDcdnDomainPropertyResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnDomainPropertyResponse) SetBody(v *DescribeDcdnDomainPropertyResponseBody) *DescribeDcdnDomainPropertyResponse {
	s.Body = v
	return s
}

type DescribeDcdnDomainPvDataRequest struct {
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EndTime    *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	StartTime  *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDcdnDomainPvDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainPvDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainPvDataRequest) SetDomainName(v string) *DescribeDcdnDomainPvDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDcdnDomainPvDataRequest) SetEndTime(v string) *DescribeDcdnDomainPvDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDcdnDomainPvDataRequest) SetOwnerId(v int64) *DescribeDcdnDomainPvDataRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDcdnDomainPvDataRequest) SetStartTime(v string) *DescribeDcdnDomainPvDataRequest {
	s.StartTime = &v
	return s
}

type DescribeDcdnDomainPvDataResponseBody struct {
	DataInterval   *string                                             `json:"DataInterval,omitempty" xml:"DataInterval,omitempty"`
	DomainName     *string                                             `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EndTime        *string                                             `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	PvDataInterval *DescribeDcdnDomainPvDataResponseBodyPvDataInterval `json:"PvDataInterval,omitempty" xml:"PvDataInterval,omitempty" type:"Struct"`
	RequestId      *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StartTime      *string                                             `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDcdnDomainPvDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainPvDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainPvDataResponseBody) SetDataInterval(v string) *DescribeDcdnDomainPvDataResponseBody {
	s.DataInterval = &v
	return s
}

func (s *DescribeDcdnDomainPvDataResponseBody) SetDomainName(v string) *DescribeDcdnDomainPvDataResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeDcdnDomainPvDataResponseBody) SetEndTime(v string) *DescribeDcdnDomainPvDataResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeDcdnDomainPvDataResponseBody) SetPvDataInterval(v *DescribeDcdnDomainPvDataResponseBodyPvDataInterval) *DescribeDcdnDomainPvDataResponseBody {
	s.PvDataInterval = v
	return s
}

func (s *DescribeDcdnDomainPvDataResponseBody) SetRequestId(v string) *DescribeDcdnDomainPvDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDcdnDomainPvDataResponseBody) SetStartTime(v string) *DescribeDcdnDomainPvDataResponseBody {
	s.StartTime = &v
	return s
}

type DescribeDcdnDomainPvDataResponseBodyPvDataInterval struct {
	UsageData []*DescribeDcdnDomainPvDataResponseBodyPvDataIntervalUsageData `json:"UsageData,omitempty" xml:"UsageData,omitempty" type:"Repeated"`
}

func (s DescribeDcdnDomainPvDataResponseBodyPvDataInterval) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainPvDataResponseBodyPvDataInterval) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainPvDataResponseBodyPvDataInterval) SetUsageData(v []*DescribeDcdnDomainPvDataResponseBodyPvDataIntervalUsageData) *DescribeDcdnDomainPvDataResponseBodyPvDataInterval {
	s.UsageData = v
	return s
}

type DescribeDcdnDomainPvDataResponseBodyPvDataIntervalUsageData struct {
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
	Value     *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDcdnDomainPvDataResponseBodyPvDataIntervalUsageData) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainPvDataResponseBodyPvDataIntervalUsageData) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainPvDataResponseBodyPvDataIntervalUsageData) SetTimeStamp(v string) *DescribeDcdnDomainPvDataResponseBodyPvDataIntervalUsageData {
	s.TimeStamp = &v
	return s
}

func (s *DescribeDcdnDomainPvDataResponseBodyPvDataIntervalUsageData) SetValue(v string) *DescribeDcdnDomainPvDataResponseBodyPvDataIntervalUsageData {
	s.Value = &v
	return s
}

type DescribeDcdnDomainPvDataResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDcdnDomainPvDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDcdnDomainPvDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainPvDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainPvDataResponse) SetHeaders(v map[string]*string) *DescribeDcdnDomainPvDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnDomainPvDataResponse) SetBody(v *DescribeDcdnDomainPvDataResponseBody) *DescribeDcdnDomainPvDataResponse {
	s.Body = v
	return s
}

type DescribeDcdnDomainQpsDataRequest struct {
	DomainName     *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EndTime        *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Interval       *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	IspNameEn      *string `json:"IspNameEn,omitempty" xml:"IspNameEn,omitempty"`
	LocationNameEn *string `json:"LocationNameEn,omitempty" xml:"LocationNameEn,omitempty"`
	OwnerId        *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	StartTime      *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDcdnDomainQpsDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainQpsDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainQpsDataRequest) SetDomainName(v string) *DescribeDcdnDomainQpsDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDcdnDomainQpsDataRequest) SetEndTime(v string) *DescribeDcdnDomainQpsDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDcdnDomainQpsDataRequest) SetInterval(v string) *DescribeDcdnDomainQpsDataRequest {
	s.Interval = &v
	return s
}

func (s *DescribeDcdnDomainQpsDataRequest) SetIspNameEn(v string) *DescribeDcdnDomainQpsDataRequest {
	s.IspNameEn = &v
	return s
}

func (s *DescribeDcdnDomainQpsDataRequest) SetLocationNameEn(v string) *DescribeDcdnDomainQpsDataRequest {
	s.LocationNameEn = &v
	return s
}

func (s *DescribeDcdnDomainQpsDataRequest) SetOwnerId(v int64) *DescribeDcdnDomainQpsDataRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDcdnDomainQpsDataRequest) SetStartTime(v string) *DescribeDcdnDomainQpsDataRequest {
	s.StartTime = &v
	return s
}

type DescribeDcdnDomainQpsDataResponseBody struct {
	DataInterval       *string                                                  `json:"DataInterval,omitempty" xml:"DataInterval,omitempty"`
	DomainName         *string                                                  `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EndTime            *string                                                  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	QpsDataPerInterval *DescribeDcdnDomainQpsDataResponseBodyQpsDataPerInterval `json:"QpsDataPerInterval,omitempty" xml:"QpsDataPerInterval,omitempty" type:"Struct"`
	RequestId          *string                                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StartTime          *string                                                  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDcdnDomainQpsDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainQpsDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainQpsDataResponseBody) SetDataInterval(v string) *DescribeDcdnDomainQpsDataResponseBody {
	s.DataInterval = &v
	return s
}

func (s *DescribeDcdnDomainQpsDataResponseBody) SetDomainName(v string) *DescribeDcdnDomainQpsDataResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeDcdnDomainQpsDataResponseBody) SetEndTime(v string) *DescribeDcdnDomainQpsDataResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeDcdnDomainQpsDataResponseBody) SetQpsDataPerInterval(v *DescribeDcdnDomainQpsDataResponseBodyQpsDataPerInterval) *DescribeDcdnDomainQpsDataResponseBody {
	s.QpsDataPerInterval = v
	return s
}

func (s *DescribeDcdnDomainQpsDataResponseBody) SetRequestId(v string) *DescribeDcdnDomainQpsDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDcdnDomainQpsDataResponseBody) SetStartTime(v string) *DescribeDcdnDomainQpsDataResponseBody {
	s.StartTime = &v
	return s
}

type DescribeDcdnDomainQpsDataResponseBodyQpsDataPerInterval struct {
	DataModule []*DescribeDcdnDomainQpsDataResponseBodyQpsDataPerIntervalDataModule `json:"DataModule,omitempty" xml:"DataModule,omitempty" type:"Repeated"`
}

func (s DescribeDcdnDomainQpsDataResponseBodyQpsDataPerInterval) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainQpsDataResponseBodyQpsDataPerInterval) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainQpsDataResponseBodyQpsDataPerInterval) SetDataModule(v []*DescribeDcdnDomainQpsDataResponseBodyQpsDataPerIntervalDataModule) *DescribeDcdnDomainQpsDataResponseBodyQpsDataPerInterval {
	s.DataModule = v
	return s
}

type DescribeDcdnDomainQpsDataResponseBodyQpsDataPerIntervalDataModule struct {
	Acc             *float32 `json:"Acc,omitempty" xml:"Acc,omitempty"`
	DynamicHttpAcc  *float32 `json:"DynamicHttpAcc,omitempty" xml:"DynamicHttpAcc,omitempty"`
	DynamicHttpQps  *float32 `json:"DynamicHttpQps,omitempty" xml:"DynamicHttpQps,omitempty"`
	DynamicHttpsAcc *float32 `json:"DynamicHttpsAcc,omitempty" xml:"DynamicHttpsAcc,omitempty"`
	DynamicHttpsQps *float32 `json:"DynamicHttpsQps,omitempty" xml:"DynamicHttpsQps,omitempty"`
	Qps             *float32 `json:"Qps,omitempty" xml:"Qps,omitempty"`
	StaticHttpAcc   *float32 `json:"StaticHttpAcc,omitempty" xml:"StaticHttpAcc,omitempty"`
	StaticHttpQps   *float32 `json:"StaticHttpQps,omitempty" xml:"StaticHttpQps,omitempty"`
	StaticHttpsAcc  *float32 `json:"StaticHttpsAcc,omitempty" xml:"StaticHttpsAcc,omitempty"`
	StaticHttpsQps  *float32 `json:"StaticHttpsQps,omitempty" xml:"StaticHttpsQps,omitempty"`
	TimeStamp       *string  `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
}

func (s DescribeDcdnDomainQpsDataResponseBodyQpsDataPerIntervalDataModule) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainQpsDataResponseBodyQpsDataPerIntervalDataModule) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainQpsDataResponseBodyQpsDataPerIntervalDataModule) SetAcc(v float32) *DescribeDcdnDomainQpsDataResponseBodyQpsDataPerIntervalDataModule {
	s.Acc = &v
	return s
}

func (s *DescribeDcdnDomainQpsDataResponseBodyQpsDataPerIntervalDataModule) SetDynamicHttpAcc(v float32) *DescribeDcdnDomainQpsDataResponseBodyQpsDataPerIntervalDataModule {
	s.DynamicHttpAcc = &v
	return s
}

func (s *DescribeDcdnDomainQpsDataResponseBodyQpsDataPerIntervalDataModule) SetDynamicHttpQps(v float32) *DescribeDcdnDomainQpsDataResponseBodyQpsDataPerIntervalDataModule {
	s.DynamicHttpQps = &v
	return s
}

func (s *DescribeDcdnDomainQpsDataResponseBodyQpsDataPerIntervalDataModule) SetDynamicHttpsAcc(v float32) *DescribeDcdnDomainQpsDataResponseBodyQpsDataPerIntervalDataModule {
	s.DynamicHttpsAcc = &v
	return s
}

func (s *DescribeDcdnDomainQpsDataResponseBodyQpsDataPerIntervalDataModule) SetDynamicHttpsQps(v float32) *DescribeDcdnDomainQpsDataResponseBodyQpsDataPerIntervalDataModule {
	s.DynamicHttpsQps = &v
	return s
}

func (s *DescribeDcdnDomainQpsDataResponseBodyQpsDataPerIntervalDataModule) SetQps(v float32) *DescribeDcdnDomainQpsDataResponseBodyQpsDataPerIntervalDataModule {
	s.Qps = &v
	return s
}

func (s *DescribeDcdnDomainQpsDataResponseBodyQpsDataPerIntervalDataModule) SetStaticHttpAcc(v float32) *DescribeDcdnDomainQpsDataResponseBodyQpsDataPerIntervalDataModule {
	s.StaticHttpAcc = &v
	return s
}

func (s *DescribeDcdnDomainQpsDataResponseBodyQpsDataPerIntervalDataModule) SetStaticHttpQps(v float32) *DescribeDcdnDomainQpsDataResponseBodyQpsDataPerIntervalDataModule {
	s.StaticHttpQps = &v
	return s
}

func (s *DescribeDcdnDomainQpsDataResponseBodyQpsDataPerIntervalDataModule) SetStaticHttpsAcc(v float32) *DescribeDcdnDomainQpsDataResponseBodyQpsDataPerIntervalDataModule {
	s.StaticHttpsAcc = &v
	return s
}

func (s *DescribeDcdnDomainQpsDataResponseBodyQpsDataPerIntervalDataModule) SetStaticHttpsQps(v float32) *DescribeDcdnDomainQpsDataResponseBodyQpsDataPerIntervalDataModule {
	s.StaticHttpsQps = &v
	return s
}

func (s *DescribeDcdnDomainQpsDataResponseBodyQpsDataPerIntervalDataModule) SetTimeStamp(v string) *DescribeDcdnDomainQpsDataResponseBodyQpsDataPerIntervalDataModule {
	s.TimeStamp = &v
	return s
}

type DescribeDcdnDomainQpsDataResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDcdnDomainQpsDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDcdnDomainQpsDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainQpsDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainQpsDataResponse) SetHeaders(v map[string]*string) *DescribeDcdnDomainQpsDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnDomainQpsDataResponse) SetBody(v *DescribeDcdnDomainQpsDataResponseBody) *DescribeDcdnDomainQpsDataResponse {
	s.Body = v
	return s
}

type DescribeDcdnDomainRealTimeBpsDataRequest struct {
	DomainName     *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EndTime        *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	IspNameEn      *string `json:"IspNameEn,omitempty" xml:"IspNameEn,omitempty"`
	LocationNameEn *string `json:"LocationNameEn,omitempty" xml:"LocationNameEn,omitempty"`
	OwnerId        *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	StartTime      *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDcdnDomainRealTimeBpsDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainRealTimeBpsDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainRealTimeBpsDataRequest) SetDomainName(v string) *DescribeDcdnDomainRealTimeBpsDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDcdnDomainRealTimeBpsDataRequest) SetEndTime(v string) *DescribeDcdnDomainRealTimeBpsDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDcdnDomainRealTimeBpsDataRequest) SetIspNameEn(v string) *DescribeDcdnDomainRealTimeBpsDataRequest {
	s.IspNameEn = &v
	return s
}

func (s *DescribeDcdnDomainRealTimeBpsDataRequest) SetLocationNameEn(v string) *DescribeDcdnDomainRealTimeBpsDataRequest {
	s.LocationNameEn = &v
	return s
}

func (s *DescribeDcdnDomainRealTimeBpsDataRequest) SetOwnerId(v int64) *DescribeDcdnDomainRealTimeBpsDataRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDcdnDomainRealTimeBpsDataRequest) SetStartTime(v string) *DescribeDcdnDomainRealTimeBpsDataRequest {
	s.StartTime = &v
	return s
}

type DescribeDcdnDomainRealTimeBpsDataResponseBody struct {
	Data      *DescribeDcdnDomainRealTimeBpsDataResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDcdnDomainRealTimeBpsDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainRealTimeBpsDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainRealTimeBpsDataResponseBody) SetData(v *DescribeDcdnDomainRealTimeBpsDataResponseBodyData) *DescribeDcdnDomainRealTimeBpsDataResponseBody {
	s.Data = v
	return s
}

func (s *DescribeDcdnDomainRealTimeBpsDataResponseBody) SetRequestId(v string) *DescribeDcdnDomainRealTimeBpsDataResponseBody {
	s.RequestId = &v
	return s
}

type DescribeDcdnDomainRealTimeBpsDataResponseBodyData struct {
	BpsModel []*DescribeDcdnDomainRealTimeBpsDataResponseBodyDataBpsModel `json:"BpsModel,omitempty" xml:"BpsModel,omitempty" type:"Repeated"`
}

func (s DescribeDcdnDomainRealTimeBpsDataResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainRealTimeBpsDataResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainRealTimeBpsDataResponseBodyData) SetBpsModel(v []*DescribeDcdnDomainRealTimeBpsDataResponseBodyDataBpsModel) *DescribeDcdnDomainRealTimeBpsDataResponseBodyData {
	s.BpsModel = v
	return s
}

type DescribeDcdnDomainRealTimeBpsDataResponseBodyDataBpsModel struct {
	Bps       *float32 `json:"Bps,omitempty" xml:"Bps,omitempty"`
	TimeStamp *string  `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
}

func (s DescribeDcdnDomainRealTimeBpsDataResponseBodyDataBpsModel) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainRealTimeBpsDataResponseBodyDataBpsModel) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainRealTimeBpsDataResponseBodyDataBpsModel) SetBps(v float32) *DescribeDcdnDomainRealTimeBpsDataResponseBodyDataBpsModel {
	s.Bps = &v
	return s
}

func (s *DescribeDcdnDomainRealTimeBpsDataResponseBodyDataBpsModel) SetTimeStamp(v string) *DescribeDcdnDomainRealTimeBpsDataResponseBodyDataBpsModel {
	s.TimeStamp = &v
	return s
}

type DescribeDcdnDomainRealTimeBpsDataResponse struct {
	Headers map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDcdnDomainRealTimeBpsDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDcdnDomainRealTimeBpsDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainRealTimeBpsDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainRealTimeBpsDataResponse) SetHeaders(v map[string]*string) *DescribeDcdnDomainRealTimeBpsDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnDomainRealTimeBpsDataResponse) SetBody(v *DescribeDcdnDomainRealTimeBpsDataResponseBody) *DescribeDcdnDomainRealTimeBpsDataResponse {
	s.Body = v
	return s
}

type DescribeDcdnDomainRealTimeByteHitRateDataRequest struct {
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EndTime    *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	StartTime  *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDcdnDomainRealTimeByteHitRateDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainRealTimeByteHitRateDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainRealTimeByteHitRateDataRequest) SetDomainName(v string) *DescribeDcdnDomainRealTimeByteHitRateDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDcdnDomainRealTimeByteHitRateDataRequest) SetEndTime(v string) *DescribeDcdnDomainRealTimeByteHitRateDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDcdnDomainRealTimeByteHitRateDataRequest) SetOwnerId(v int64) *DescribeDcdnDomainRealTimeByteHitRateDataRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDcdnDomainRealTimeByteHitRateDataRequest) SetStartTime(v string) *DescribeDcdnDomainRealTimeByteHitRateDataRequest {
	s.StartTime = &v
	return s
}

type DescribeDcdnDomainRealTimeByteHitRateDataResponseBody struct {
	Data      *DescribeDcdnDomainRealTimeByteHitRateDataResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDcdnDomainRealTimeByteHitRateDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainRealTimeByteHitRateDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainRealTimeByteHitRateDataResponseBody) SetData(v *DescribeDcdnDomainRealTimeByteHitRateDataResponseBodyData) *DescribeDcdnDomainRealTimeByteHitRateDataResponseBody {
	s.Data = v
	return s
}

func (s *DescribeDcdnDomainRealTimeByteHitRateDataResponseBody) SetRequestId(v string) *DescribeDcdnDomainRealTimeByteHitRateDataResponseBody {
	s.RequestId = &v
	return s
}

type DescribeDcdnDomainRealTimeByteHitRateDataResponseBodyData struct {
	ByteHitRateDataModel []*DescribeDcdnDomainRealTimeByteHitRateDataResponseBodyDataByteHitRateDataModel `json:"ByteHitRateDataModel,omitempty" xml:"ByteHitRateDataModel,omitempty" type:"Repeated"`
}

func (s DescribeDcdnDomainRealTimeByteHitRateDataResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainRealTimeByteHitRateDataResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainRealTimeByteHitRateDataResponseBodyData) SetByteHitRateDataModel(v []*DescribeDcdnDomainRealTimeByteHitRateDataResponseBodyDataByteHitRateDataModel) *DescribeDcdnDomainRealTimeByteHitRateDataResponseBodyData {
	s.ByteHitRateDataModel = v
	return s
}

type DescribeDcdnDomainRealTimeByteHitRateDataResponseBodyDataByteHitRateDataModel struct {
	ByteHitRate *float32 `json:"ByteHitRate,omitempty" xml:"ByteHitRate,omitempty"`
	TimeStamp   *string  `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
}

func (s DescribeDcdnDomainRealTimeByteHitRateDataResponseBodyDataByteHitRateDataModel) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainRealTimeByteHitRateDataResponseBodyDataByteHitRateDataModel) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainRealTimeByteHitRateDataResponseBodyDataByteHitRateDataModel) SetByteHitRate(v float32) *DescribeDcdnDomainRealTimeByteHitRateDataResponseBodyDataByteHitRateDataModel {
	s.ByteHitRate = &v
	return s
}

func (s *DescribeDcdnDomainRealTimeByteHitRateDataResponseBodyDataByteHitRateDataModel) SetTimeStamp(v string) *DescribeDcdnDomainRealTimeByteHitRateDataResponseBodyDataByteHitRateDataModel {
	s.TimeStamp = &v
	return s
}

type DescribeDcdnDomainRealTimeByteHitRateDataResponse struct {
	Headers map[string]*string                                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDcdnDomainRealTimeByteHitRateDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDcdnDomainRealTimeByteHitRateDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainRealTimeByteHitRateDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainRealTimeByteHitRateDataResponse) SetHeaders(v map[string]*string) *DescribeDcdnDomainRealTimeByteHitRateDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnDomainRealTimeByteHitRateDataResponse) SetBody(v *DescribeDcdnDomainRealTimeByteHitRateDataResponseBody) *DescribeDcdnDomainRealTimeByteHitRateDataResponse {
	s.Body = v
	return s
}

type DescribeDcdnDomainRealTimeDetailDataRequest struct {
	DomainName     *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EndTime        *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Field          *string `json:"Field,omitempty" xml:"Field,omitempty"`
	IspNameEn      *string `json:"IspNameEn,omitempty" xml:"IspNameEn,omitempty"`
	LocationNameEn *string `json:"LocationNameEn,omitempty" xml:"LocationNameEn,omitempty"`
	Merge          *string `json:"Merge,omitempty" xml:"Merge,omitempty"`
	MergeLocIsp    *string `json:"MergeLocIsp,omitempty" xml:"MergeLocIsp,omitempty"`
	OwnerId        *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	StartTime      *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDcdnDomainRealTimeDetailDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainRealTimeDetailDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainRealTimeDetailDataRequest) SetDomainName(v string) *DescribeDcdnDomainRealTimeDetailDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDcdnDomainRealTimeDetailDataRequest) SetEndTime(v string) *DescribeDcdnDomainRealTimeDetailDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDcdnDomainRealTimeDetailDataRequest) SetField(v string) *DescribeDcdnDomainRealTimeDetailDataRequest {
	s.Field = &v
	return s
}

func (s *DescribeDcdnDomainRealTimeDetailDataRequest) SetIspNameEn(v string) *DescribeDcdnDomainRealTimeDetailDataRequest {
	s.IspNameEn = &v
	return s
}

func (s *DescribeDcdnDomainRealTimeDetailDataRequest) SetLocationNameEn(v string) *DescribeDcdnDomainRealTimeDetailDataRequest {
	s.LocationNameEn = &v
	return s
}

func (s *DescribeDcdnDomainRealTimeDetailDataRequest) SetMerge(v string) *DescribeDcdnDomainRealTimeDetailDataRequest {
	s.Merge = &v
	return s
}

func (s *DescribeDcdnDomainRealTimeDetailDataRequest) SetMergeLocIsp(v string) *DescribeDcdnDomainRealTimeDetailDataRequest {
	s.MergeLocIsp = &v
	return s
}

func (s *DescribeDcdnDomainRealTimeDetailDataRequest) SetOwnerId(v int64) *DescribeDcdnDomainRealTimeDetailDataRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDcdnDomainRealTimeDetailDataRequest) SetStartTime(v string) *DescribeDcdnDomainRealTimeDetailDataRequest {
	s.StartTime = &v
	return s
}

type DescribeDcdnDomainRealTimeDetailDataResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDcdnDomainRealTimeDetailDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainRealTimeDetailDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainRealTimeDetailDataResponseBody) SetData(v string) *DescribeDcdnDomainRealTimeDetailDataResponseBody {
	s.Data = &v
	return s
}

func (s *DescribeDcdnDomainRealTimeDetailDataResponseBody) SetRequestId(v string) *DescribeDcdnDomainRealTimeDetailDataResponseBody {
	s.RequestId = &v
	return s
}

type DescribeDcdnDomainRealTimeDetailDataResponse struct {
	Headers map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDcdnDomainRealTimeDetailDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDcdnDomainRealTimeDetailDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainRealTimeDetailDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainRealTimeDetailDataResponse) SetHeaders(v map[string]*string) *DescribeDcdnDomainRealTimeDetailDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnDomainRealTimeDetailDataResponse) SetBody(v *DescribeDcdnDomainRealTimeDetailDataResponseBody) *DescribeDcdnDomainRealTimeDetailDataResponse {
	s.Body = v
	return s
}

type DescribeDcdnDomainRealTimeHttpCodeDataRequest struct {
	DomainName     *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EndTime        *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	IspNameEn      *string `json:"IspNameEn,omitempty" xml:"IspNameEn,omitempty"`
	LocationNameEn *string `json:"LocationNameEn,omitempty" xml:"LocationNameEn,omitempty"`
	OwnerId        *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	StartTime      *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDcdnDomainRealTimeHttpCodeDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainRealTimeHttpCodeDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainRealTimeHttpCodeDataRequest) SetDomainName(v string) *DescribeDcdnDomainRealTimeHttpCodeDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDcdnDomainRealTimeHttpCodeDataRequest) SetEndTime(v string) *DescribeDcdnDomainRealTimeHttpCodeDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDcdnDomainRealTimeHttpCodeDataRequest) SetIspNameEn(v string) *DescribeDcdnDomainRealTimeHttpCodeDataRequest {
	s.IspNameEn = &v
	return s
}

func (s *DescribeDcdnDomainRealTimeHttpCodeDataRequest) SetLocationNameEn(v string) *DescribeDcdnDomainRealTimeHttpCodeDataRequest {
	s.LocationNameEn = &v
	return s
}

func (s *DescribeDcdnDomainRealTimeHttpCodeDataRequest) SetOwnerId(v int64) *DescribeDcdnDomainRealTimeHttpCodeDataRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDcdnDomainRealTimeHttpCodeDataRequest) SetStartTime(v string) *DescribeDcdnDomainRealTimeHttpCodeDataRequest {
	s.StartTime = &v
	return s
}

type DescribeDcdnDomainRealTimeHttpCodeDataResponseBody struct {
	DataInterval         *string                                                                 `json:"DataInterval,omitempty" xml:"DataInterval,omitempty"`
	DomainName           *string                                                                 `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EndTime              *string                                                                 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	RealTimeHttpCodeData *DescribeDcdnDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeData `json:"RealTimeHttpCodeData,omitempty" xml:"RealTimeHttpCodeData,omitempty" type:"Struct"`
	RequestId            *string                                                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StartTime            *string                                                                 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDcdnDomainRealTimeHttpCodeDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainRealTimeHttpCodeDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainRealTimeHttpCodeDataResponseBody) SetDataInterval(v string) *DescribeDcdnDomainRealTimeHttpCodeDataResponseBody {
	s.DataInterval = &v
	return s
}

func (s *DescribeDcdnDomainRealTimeHttpCodeDataResponseBody) SetDomainName(v string) *DescribeDcdnDomainRealTimeHttpCodeDataResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeDcdnDomainRealTimeHttpCodeDataResponseBody) SetEndTime(v string) *DescribeDcdnDomainRealTimeHttpCodeDataResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeDcdnDomainRealTimeHttpCodeDataResponseBody) SetRealTimeHttpCodeData(v *DescribeDcdnDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeData) *DescribeDcdnDomainRealTimeHttpCodeDataResponseBody {
	s.RealTimeHttpCodeData = v
	return s
}

func (s *DescribeDcdnDomainRealTimeHttpCodeDataResponseBody) SetRequestId(v string) *DescribeDcdnDomainRealTimeHttpCodeDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDcdnDomainRealTimeHttpCodeDataResponseBody) SetStartTime(v string) *DescribeDcdnDomainRealTimeHttpCodeDataResponseBody {
	s.StartTime = &v
	return s
}

type DescribeDcdnDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeData struct {
	UsageData []*DescribeDcdnDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeDataUsageData `json:"UsageData,omitempty" xml:"UsageData,omitempty" type:"Repeated"`
}

func (s DescribeDcdnDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeData) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeData) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeData) SetUsageData(v []*DescribeDcdnDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeDataUsageData) *DescribeDcdnDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeData {
	s.UsageData = v
	return s
}

type DescribeDcdnDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeDataUsageData struct {
	TimeStamp *string                                                                               `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
	Value     *DescribeDcdnDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeDataUsageDataValue `json:"Value,omitempty" xml:"Value,omitempty" type:"Struct"`
}

func (s DescribeDcdnDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeDataUsageData) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeDataUsageData) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeDataUsageData) SetTimeStamp(v string) *DescribeDcdnDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeDataUsageData {
	s.TimeStamp = &v
	return s
}

func (s *DescribeDcdnDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeDataUsageData) SetValue(v *DescribeDcdnDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeDataUsageDataValue) *DescribeDcdnDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeDataUsageData {
	s.Value = v
	return s
}

type DescribeDcdnDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeDataUsageDataValue struct {
	RealTimeCodeProportionData []*DescribeDcdnDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeDataUsageDataValueRealTimeCodeProportionData `json:"RealTimeCodeProportionData,omitempty" xml:"RealTimeCodeProportionData,omitempty" type:"Repeated"`
}

func (s DescribeDcdnDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeDataUsageDataValue) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeDataUsageDataValue) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeDataUsageDataValue) SetRealTimeCodeProportionData(v []*DescribeDcdnDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeDataUsageDataValueRealTimeCodeProportionData) *DescribeDcdnDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeDataUsageDataValue {
	s.RealTimeCodeProportionData = v
	return s
}

type DescribeDcdnDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeDataUsageDataValueRealTimeCodeProportionData struct {
	Code       *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Count      *string `json:"Count,omitempty" xml:"Count,omitempty"`
	Proportion *string `json:"Proportion,omitempty" xml:"Proportion,omitempty"`
}

func (s DescribeDcdnDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeDataUsageDataValueRealTimeCodeProportionData) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeDataUsageDataValueRealTimeCodeProportionData) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeDataUsageDataValueRealTimeCodeProportionData) SetCode(v string) *DescribeDcdnDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeDataUsageDataValueRealTimeCodeProportionData {
	s.Code = &v
	return s
}

func (s *DescribeDcdnDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeDataUsageDataValueRealTimeCodeProportionData) SetCount(v string) *DescribeDcdnDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeDataUsageDataValueRealTimeCodeProportionData {
	s.Count = &v
	return s
}

func (s *DescribeDcdnDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeDataUsageDataValueRealTimeCodeProportionData) SetProportion(v string) *DescribeDcdnDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeDataUsageDataValueRealTimeCodeProportionData {
	s.Proportion = &v
	return s
}

type DescribeDcdnDomainRealTimeHttpCodeDataResponse struct {
	Headers map[string]*string                                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDcdnDomainRealTimeHttpCodeDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDcdnDomainRealTimeHttpCodeDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainRealTimeHttpCodeDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainRealTimeHttpCodeDataResponse) SetHeaders(v map[string]*string) *DescribeDcdnDomainRealTimeHttpCodeDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnDomainRealTimeHttpCodeDataResponse) SetBody(v *DescribeDcdnDomainRealTimeHttpCodeDataResponseBody) *DescribeDcdnDomainRealTimeHttpCodeDataResponse {
	s.Body = v
	return s
}

type DescribeDcdnDomainRealTimeQpsDataRequest struct {
	DomainName     *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EndTime        *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	IspNameEn      *string `json:"IspNameEn,omitempty" xml:"IspNameEn,omitempty"`
	LocationNameEn *string `json:"LocationNameEn,omitempty" xml:"LocationNameEn,omitempty"`
	OwnerId        *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	StartTime      *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDcdnDomainRealTimeQpsDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainRealTimeQpsDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainRealTimeQpsDataRequest) SetDomainName(v string) *DescribeDcdnDomainRealTimeQpsDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDcdnDomainRealTimeQpsDataRequest) SetEndTime(v string) *DescribeDcdnDomainRealTimeQpsDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDcdnDomainRealTimeQpsDataRequest) SetIspNameEn(v string) *DescribeDcdnDomainRealTimeQpsDataRequest {
	s.IspNameEn = &v
	return s
}

func (s *DescribeDcdnDomainRealTimeQpsDataRequest) SetLocationNameEn(v string) *DescribeDcdnDomainRealTimeQpsDataRequest {
	s.LocationNameEn = &v
	return s
}

func (s *DescribeDcdnDomainRealTimeQpsDataRequest) SetOwnerId(v int64) *DescribeDcdnDomainRealTimeQpsDataRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDcdnDomainRealTimeQpsDataRequest) SetStartTime(v string) *DescribeDcdnDomainRealTimeQpsDataRequest {
	s.StartTime = &v
	return s
}

type DescribeDcdnDomainRealTimeQpsDataResponseBody struct {
	Data      *DescribeDcdnDomainRealTimeQpsDataResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDcdnDomainRealTimeQpsDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainRealTimeQpsDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainRealTimeQpsDataResponseBody) SetData(v *DescribeDcdnDomainRealTimeQpsDataResponseBodyData) *DescribeDcdnDomainRealTimeQpsDataResponseBody {
	s.Data = v
	return s
}

func (s *DescribeDcdnDomainRealTimeQpsDataResponseBody) SetRequestId(v string) *DescribeDcdnDomainRealTimeQpsDataResponseBody {
	s.RequestId = &v
	return s
}

type DescribeDcdnDomainRealTimeQpsDataResponseBodyData struct {
	QpsModel []*DescribeDcdnDomainRealTimeQpsDataResponseBodyDataQpsModel `json:"QpsModel,omitempty" xml:"QpsModel,omitempty" type:"Repeated"`
}

func (s DescribeDcdnDomainRealTimeQpsDataResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainRealTimeQpsDataResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainRealTimeQpsDataResponseBodyData) SetQpsModel(v []*DescribeDcdnDomainRealTimeQpsDataResponseBodyDataQpsModel) *DescribeDcdnDomainRealTimeQpsDataResponseBodyData {
	s.QpsModel = v
	return s
}

type DescribeDcdnDomainRealTimeQpsDataResponseBodyDataQpsModel struct {
	Qps       *float32 `json:"Qps,omitempty" xml:"Qps,omitempty"`
	TimeStamp *string  `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
}

func (s DescribeDcdnDomainRealTimeQpsDataResponseBodyDataQpsModel) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainRealTimeQpsDataResponseBodyDataQpsModel) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainRealTimeQpsDataResponseBodyDataQpsModel) SetQps(v float32) *DescribeDcdnDomainRealTimeQpsDataResponseBodyDataQpsModel {
	s.Qps = &v
	return s
}

func (s *DescribeDcdnDomainRealTimeQpsDataResponseBodyDataQpsModel) SetTimeStamp(v string) *DescribeDcdnDomainRealTimeQpsDataResponseBodyDataQpsModel {
	s.TimeStamp = &v
	return s
}

type DescribeDcdnDomainRealTimeQpsDataResponse struct {
	Headers map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDcdnDomainRealTimeQpsDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDcdnDomainRealTimeQpsDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainRealTimeQpsDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainRealTimeQpsDataResponse) SetHeaders(v map[string]*string) *DescribeDcdnDomainRealTimeQpsDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnDomainRealTimeQpsDataResponse) SetBody(v *DescribeDcdnDomainRealTimeQpsDataResponseBody) *DescribeDcdnDomainRealTimeQpsDataResponse {
	s.Body = v
	return s
}

type DescribeDcdnDomainRealTimeReqHitRateDataRequest struct {
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EndTime    *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	StartTime  *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDcdnDomainRealTimeReqHitRateDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainRealTimeReqHitRateDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainRealTimeReqHitRateDataRequest) SetDomainName(v string) *DescribeDcdnDomainRealTimeReqHitRateDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDcdnDomainRealTimeReqHitRateDataRequest) SetEndTime(v string) *DescribeDcdnDomainRealTimeReqHitRateDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDcdnDomainRealTimeReqHitRateDataRequest) SetOwnerId(v int64) *DescribeDcdnDomainRealTimeReqHitRateDataRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDcdnDomainRealTimeReqHitRateDataRequest) SetStartTime(v string) *DescribeDcdnDomainRealTimeReqHitRateDataRequest {
	s.StartTime = &v
	return s
}

type DescribeDcdnDomainRealTimeReqHitRateDataResponseBody struct {
	Data      *DescribeDcdnDomainRealTimeReqHitRateDataResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDcdnDomainRealTimeReqHitRateDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainRealTimeReqHitRateDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainRealTimeReqHitRateDataResponseBody) SetData(v *DescribeDcdnDomainRealTimeReqHitRateDataResponseBodyData) *DescribeDcdnDomainRealTimeReqHitRateDataResponseBody {
	s.Data = v
	return s
}

func (s *DescribeDcdnDomainRealTimeReqHitRateDataResponseBody) SetRequestId(v string) *DescribeDcdnDomainRealTimeReqHitRateDataResponseBody {
	s.RequestId = &v
	return s
}

type DescribeDcdnDomainRealTimeReqHitRateDataResponseBodyData struct {
	ReqHitRateDataModel []*DescribeDcdnDomainRealTimeReqHitRateDataResponseBodyDataReqHitRateDataModel `json:"ReqHitRateDataModel,omitempty" xml:"ReqHitRateDataModel,omitempty" type:"Repeated"`
}

func (s DescribeDcdnDomainRealTimeReqHitRateDataResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainRealTimeReqHitRateDataResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainRealTimeReqHitRateDataResponseBodyData) SetReqHitRateDataModel(v []*DescribeDcdnDomainRealTimeReqHitRateDataResponseBodyDataReqHitRateDataModel) *DescribeDcdnDomainRealTimeReqHitRateDataResponseBodyData {
	s.ReqHitRateDataModel = v
	return s
}

type DescribeDcdnDomainRealTimeReqHitRateDataResponseBodyDataReqHitRateDataModel struct {
	ReqHitRate *float32 `json:"ReqHitRate,omitempty" xml:"ReqHitRate,omitempty"`
	TimeStamp  *string  `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
}

func (s DescribeDcdnDomainRealTimeReqHitRateDataResponseBodyDataReqHitRateDataModel) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainRealTimeReqHitRateDataResponseBodyDataReqHitRateDataModel) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainRealTimeReqHitRateDataResponseBodyDataReqHitRateDataModel) SetReqHitRate(v float32) *DescribeDcdnDomainRealTimeReqHitRateDataResponseBodyDataReqHitRateDataModel {
	s.ReqHitRate = &v
	return s
}

func (s *DescribeDcdnDomainRealTimeReqHitRateDataResponseBodyDataReqHitRateDataModel) SetTimeStamp(v string) *DescribeDcdnDomainRealTimeReqHitRateDataResponseBodyDataReqHitRateDataModel {
	s.TimeStamp = &v
	return s
}

type DescribeDcdnDomainRealTimeReqHitRateDataResponse struct {
	Headers map[string]*string                                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDcdnDomainRealTimeReqHitRateDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDcdnDomainRealTimeReqHitRateDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainRealTimeReqHitRateDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainRealTimeReqHitRateDataResponse) SetHeaders(v map[string]*string) *DescribeDcdnDomainRealTimeReqHitRateDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnDomainRealTimeReqHitRateDataResponse) SetBody(v *DescribeDcdnDomainRealTimeReqHitRateDataResponseBody) *DescribeDcdnDomainRealTimeReqHitRateDataResponse {
	s.Body = v
	return s
}

type DescribeDcdnDomainRealTimeSrcBpsDataRequest struct {
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EndTime    *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	StartTime  *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDcdnDomainRealTimeSrcBpsDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainRealTimeSrcBpsDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainRealTimeSrcBpsDataRequest) SetDomainName(v string) *DescribeDcdnDomainRealTimeSrcBpsDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDcdnDomainRealTimeSrcBpsDataRequest) SetEndTime(v string) *DescribeDcdnDomainRealTimeSrcBpsDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDcdnDomainRealTimeSrcBpsDataRequest) SetOwnerId(v int64) *DescribeDcdnDomainRealTimeSrcBpsDataRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDcdnDomainRealTimeSrcBpsDataRequest) SetStartTime(v string) *DescribeDcdnDomainRealTimeSrcBpsDataRequest {
	s.StartTime = &v
	return s
}

type DescribeDcdnDomainRealTimeSrcBpsDataResponseBody struct {
	DataInterval                  *string                                                                        `json:"DataInterval,omitempty" xml:"DataInterval,omitempty"`
	DomainName                    *string                                                                        `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EndTime                       *string                                                                        `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	RealTimeSrcBpsDataPerInterval *DescribeDcdnDomainRealTimeSrcBpsDataResponseBodyRealTimeSrcBpsDataPerInterval `json:"RealTimeSrcBpsDataPerInterval,omitempty" xml:"RealTimeSrcBpsDataPerInterval,omitempty" type:"Struct"`
	RequestId                     *string                                                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StartTime                     *string                                                                        `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDcdnDomainRealTimeSrcBpsDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainRealTimeSrcBpsDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainRealTimeSrcBpsDataResponseBody) SetDataInterval(v string) *DescribeDcdnDomainRealTimeSrcBpsDataResponseBody {
	s.DataInterval = &v
	return s
}

func (s *DescribeDcdnDomainRealTimeSrcBpsDataResponseBody) SetDomainName(v string) *DescribeDcdnDomainRealTimeSrcBpsDataResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeDcdnDomainRealTimeSrcBpsDataResponseBody) SetEndTime(v string) *DescribeDcdnDomainRealTimeSrcBpsDataResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeDcdnDomainRealTimeSrcBpsDataResponseBody) SetRealTimeSrcBpsDataPerInterval(v *DescribeDcdnDomainRealTimeSrcBpsDataResponseBodyRealTimeSrcBpsDataPerInterval) *DescribeDcdnDomainRealTimeSrcBpsDataResponseBody {
	s.RealTimeSrcBpsDataPerInterval = v
	return s
}

func (s *DescribeDcdnDomainRealTimeSrcBpsDataResponseBody) SetRequestId(v string) *DescribeDcdnDomainRealTimeSrcBpsDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDcdnDomainRealTimeSrcBpsDataResponseBody) SetStartTime(v string) *DescribeDcdnDomainRealTimeSrcBpsDataResponseBody {
	s.StartTime = &v
	return s
}

type DescribeDcdnDomainRealTimeSrcBpsDataResponseBodyRealTimeSrcBpsDataPerInterval struct {
	DataModule []*DescribeDcdnDomainRealTimeSrcBpsDataResponseBodyRealTimeSrcBpsDataPerIntervalDataModule `json:"DataModule,omitempty" xml:"DataModule,omitempty" type:"Repeated"`
}

func (s DescribeDcdnDomainRealTimeSrcBpsDataResponseBodyRealTimeSrcBpsDataPerInterval) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainRealTimeSrcBpsDataResponseBodyRealTimeSrcBpsDataPerInterval) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainRealTimeSrcBpsDataResponseBodyRealTimeSrcBpsDataPerInterval) SetDataModule(v []*DescribeDcdnDomainRealTimeSrcBpsDataResponseBodyRealTimeSrcBpsDataPerIntervalDataModule) *DescribeDcdnDomainRealTimeSrcBpsDataResponseBodyRealTimeSrcBpsDataPerInterval {
	s.DataModule = v
	return s
}

type DescribeDcdnDomainRealTimeSrcBpsDataResponseBodyRealTimeSrcBpsDataPerIntervalDataModule struct {
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
	Value     *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDcdnDomainRealTimeSrcBpsDataResponseBodyRealTimeSrcBpsDataPerIntervalDataModule) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainRealTimeSrcBpsDataResponseBodyRealTimeSrcBpsDataPerIntervalDataModule) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainRealTimeSrcBpsDataResponseBodyRealTimeSrcBpsDataPerIntervalDataModule) SetTimeStamp(v string) *DescribeDcdnDomainRealTimeSrcBpsDataResponseBodyRealTimeSrcBpsDataPerIntervalDataModule {
	s.TimeStamp = &v
	return s
}

func (s *DescribeDcdnDomainRealTimeSrcBpsDataResponseBodyRealTimeSrcBpsDataPerIntervalDataModule) SetValue(v string) *DescribeDcdnDomainRealTimeSrcBpsDataResponseBodyRealTimeSrcBpsDataPerIntervalDataModule {
	s.Value = &v
	return s
}

type DescribeDcdnDomainRealTimeSrcBpsDataResponse struct {
	Headers map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDcdnDomainRealTimeSrcBpsDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDcdnDomainRealTimeSrcBpsDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainRealTimeSrcBpsDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainRealTimeSrcBpsDataResponse) SetHeaders(v map[string]*string) *DescribeDcdnDomainRealTimeSrcBpsDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnDomainRealTimeSrcBpsDataResponse) SetBody(v *DescribeDcdnDomainRealTimeSrcBpsDataResponseBody) *DescribeDcdnDomainRealTimeSrcBpsDataResponse {
	s.Body = v
	return s
}

type DescribeDcdnDomainRealTimeSrcHttpCodeDataRequest struct {
	DomainName     *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EndTime        *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	IspNameEn      *string `json:"IspNameEn,omitempty" xml:"IspNameEn,omitempty"`
	LocationNameEn *string `json:"LocationNameEn,omitempty" xml:"LocationNameEn,omitempty"`
	OwnerId        *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	StartTime      *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDcdnDomainRealTimeSrcHttpCodeDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainRealTimeSrcHttpCodeDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainRealTimeSrcHttpCodeDataRequest) SetDomainName(v string) *DescribeDcdnDomainRealTimeSrcHttpCodeDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDcdnDomainRealTimeSrcHttpCodeDataRequest) SetEndTime(v string) *DescribeDcdnDomainRealTimeSrcHttpCodeDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDcdnDomainRealTimeSrcHttpCodeDataRequest) SetIspNameEn(v string) *DescribeDcdnDomainRealTimeSrcHttpCodeDataRequest {
	s.IspNameEn = &v
	return s
}

func (s *DescribeDcdnDomainRealTimeSrcHttpCodeDataRequest) SetLocationNameEn(v string) *DescribeDcdnDomainRealTimeSrcHttpCodeDataRequest {
	s.LocationNameEn = &v
	return s
}

func (s *DescribeDcdnDomainRealTimeSrcHttpCodeDataRequest) SetOwnerId(v int64) *DescribeDcdnDomainRealTimeSrcHttpCodeDataRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDcdnDomainRealTimeSrcHttpCodeDataRequest) SetStartTime(v string) *DescribeDcdnDomainRealTimeSrcHttpCodeDataRequest {
	s.StartTime = &v
	return s
}

type DescribeDcdnDomainRealTimeSrcHttpCodeDataResponseBody struct {
	DataInterval            *string                                                                       `json:"DataInterval,omitempty" xml:"DataInterval,omitempty"`
	DomainName              *string                                                                       `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EndTime                 *string                                                                       `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	RealTimeSrcHttpCodeData *DescribeDcdnDomainRealTimeSrcHttpCodeDataResponseBodyRealTimeSrcHttpCodeData `json:"RealTimeSrcHttpCodeData,omitempty" xml:"RealTimeSrcHttpCodeData,omitempty" type:"Struct"`
	RequestId               *string                                                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StartTime               *string                                                                       `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDcdnDomainRealTimeSrcHttpCodeDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainRealTimeSrcHttpCodeDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainRealTimeSrcHttpCodeDataResponseBody) SetDataInterval(v string) *DescribeDcdnDomainRealTimeSrcHttpCodeDataResponseBody {
	s.DataInterval = &v
	return s
}

func (s *DescribeDcdnDomainRealTimeSrcHttpCodeDataResponseBody) SetDomainName(v string) *DescribeDcdnDomainRealTimeSrcHttpCodeDataResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeDcdnDomainRealTimeSrcHttpCodeDataResponseBody) SetEndTime(v string) *DescribeDcdnDomainRealTimeSrcHttpCodeDataResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeDcdnDomainRealTimeSrcHttpCodeDataResponseBody) SetRealTimeSrcHttpCodeData(v *DescribeDcdnDomainRealTimeSrcHttpCodeDataResponseBodyRealTimeSrcHttpCodeData) *DescribeDcdnDomainRealTimeSrcHttpCodeDataResponseBody {
	s.RealTimeSrcHttpCodeData = v
	return s
}

func (s *DescribeDcdnDomainRealTimeSrcHttpCodeDataResponseBody) SetRequestId(v string) *DescribeDcdnDomainRealTimeSrcHttpCodeDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDcdnDomainRealTimeSrcHttpCodeDataResponseBody) SetStartTime(v string) *DescribeDcdnDomainRealTimeSrcHttpCodeDataResponseBody {
	s.StartTime = &v
	return s
}

type DescribeDcdnDomainRealTimeSrcHttpCodeDataResponseBodyRealTimeSrcHttpCodeData struct {
	UsageData []*DescribeDcdnDomainRealTimeSrcHttpCodeDataResponseBodyRealTimeSrcHttpCodeDataUsageData `json:"UsageData,omitempty" xml:"UsageData,omitempty" type:"Repeated"`
}

func (s DescribeDcdnDomainRealTimeSrcHttpCodeDataResponseBodyRealTimeSrcHttpCodeData) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainRealTimeSrcHttpCodeDataResponseBodyRealTimeSrcHttpCodeData) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainRealTimeSrcHttpCodeDataResponseBodyRealTimeSrcHttpCodeData) SetUsageData(v []*DescribeDcdnDomainRealTimeSrcHttpCodeDataResponseBodyRealTimeSrcHttpCodeDataUsageData) *DescribeDcdnDomainRealTimeSrcHttpCodeDataResponseBodyRealTimeSrcHttpCodeData {
	s.UsageData = v
	return s
}

type DescribeDcdnDomainRealTimeSrcHttpCodeDataResponseBodyRealTimeSrcHttpCodeDataUsageData struct {
	TimeStamp *string                                                                                     `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
	Value     *DescribeDcdnDomainRealTimeSrcHttpCodeDataResponseBodyRealTimeSrcHttpCodeDataUsageDataValue `json:"Value,omitempty" xml:"Value,omitempty" type:"Struct"`
}

func (s DescribeDcdnDomainRealTimeSrcHttpCodeDataResponseBodyRealTimeSrcHttpCodeDataUsageData) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainRealTimeSrcHttpCodeDataResponseBodyRealTimeSrcHttpCodeDataUsageData) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainRealTimeSrcHttpCodeDataResponseBodyRealTimeSrcHttpCodeDataUsageData) SetTimeStamp(v string) *DescribeDcdnDomainRealTimeSrcHttpCodeDataResponseBodyRealTimeSrcHttpCodeDataUsageData {
	s.TimeStamp = &v
	return s
}

func (s *DescribeDcdnDomainRealTimeSrcHttpCodeDataResponseBodyRealTimeSrcHttpCodeDataUsageData) SetValue(v *DescribeDcdnDomainRealTimeSrcHttpCodeDataResponseBodyRealTimeSrcHttpCodeDataUsageDataValue) *DescribeDcdnDomainRealTimeSrcHttpCodeDataResponseBodyRealTimeSrcHttpCodeDataUsageData {
	s.Value = v
	return s
}

type DescribeDcdnDomainRealTimeSrcHttpCodeDataResponseBodyRealTimeSrcHttpCodeDataUsageDataValue struct {
	RealTimeSrcCodeProportionData []*DescribeDcdnDomainRealTimeSrcHttpCodeDataResponseBodyRealTimeSrcHttpCodeDataUsageDataValueRealTimeSrcCodeProportionData `json:"RealTimeSrcCodeProportionData,omitempty" xml:"RealTimeSrcCodeProportionData,omitempty" type:"Repeated"`
}

func (s DescribeDcdnDomainRealTimeSrcHttpCodeDataResponseBodyRealTimeSrcHttpCodeDataUsageDataValue) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainRealTimeSrcHttpCodeDataResponseBodyRealTimeSrcHttpCodeDataUsageDataValue) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainRealTimeSrcHttpCodeDataResponseBodyRealTimeSrcHttpCodeDataUsageDataValue) SetRealTimeSrcCodeProportionData(v []*DescribeDcdnDomainRealTimeSrcHttpCodeDataResponseBodyRealTimeSrcHttpCodeDataUsageDataValueRealTimeSrcCodeProportionData) *DescribeDcdnDomainRealTimeSrcHttpCodeDataResponseBodyRealTimeSrcHttpCodeDataUsageDataValue {
	s.RealTimeSrcCodeProportionData = v
	return s
}

type DescribeDcdnDomainRealTimeSrcHttpCodeDataResponseBodyRealTimeSrcHttpCodeDataUsageDataValueRealTimeSrcCodeProportionData struct {
	Code       *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Count      *string `json:"Count,omitempty" xml:"Count,omitempty"`
	Proportion *string `json:"Proportion,omitempty" xml:"Proportion,omitempty"`
}

func (s DescribeDcdnDomainRealTimeSrcHttpCodeDataResponseBodyRealTimeSrcHttpCodeDataUsageDataValueRealTimeSrcCodeProportionData) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainRealTimeSrcHttpCodeDataResponseBodyRealTimeSrcHttpCodeDataUsageDataValueRealTimeSrcCodeProportionData) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainRealTimeSrcHttpCodeDataResponseBodyRealTimeSrcHttpCodeDataUsageDataValueRealTimeSrcCodeProportionData) SetCode(v string) *DescribeDcdnDomainRealTimeSrcHttpCodeDataResponseBodyRealTimeSrcHttpCodeDataUsageDataValueRealTimeSrcCodeProportionData {
	s.Code = &v
	return s
}

func (s *DescribeDcdnDomainRealTimeSrcHttpCodeDataResponseBodyRealTimeSrcHttpCodeDataUsageDataValueRealTimeSrcCodeProportionData) SetCount(v string) *DescribeDcdnDomainRealTimeSrcHttpCodeDataResponseBodyRealTimeSrcHttpCodeDataUsageDataValueRealTimeSrcCodeProportionData {
	s.Count = &v
	return s
}

func (s *DescribeDcdnDomainRealTimeSrcHttpCodeDataResponseBodyRealTimeSrcHttpCodeDataUsageDataValueRealTimeSrcCodeProportionData) SetProportion(v string) *DescribeDcdnDomainRealTimeSrcHttpCodeDataResponseBodyRealTimeSrcHttpCodeDataUsageDataValueRealTimeSrcCodeProportionData {
	s.Proportion = &v
	return s
}

type DescribeDcdnDomainRealTimeSrcHttpCodeDataResponse struct {
	Headers map[string]*string                                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDcdnDomainRealTimeSrcHttpCodeDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDcdnDomainRealTimeSrcHttpCodeDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainRealTimeSrcHttpCodeDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainRealTimeSrcHttpCodeDataResponse) SetHeaders(v map[string]*string) *DescribeDcdnDomainRealTimeSrcHttpCodeDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnDomainRealTimeSrcHttpCodeDataResponse) SetBody(v *DescribeDcdnDomainRealTimeSrcHttpCodeDataResponseBody) *DescribeDcdnDomainRealTimeSrcHttpCodeDataResponse {
	s.Body = v
	return s
}

type DescribeDcdnDomainRealTimeSrcTrafficDataRequest struct {
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EndTime    *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	StartTime  *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDcdnDomainRealTimeSrcTrafficDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainRealTimeSrcTrafficDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainRealTimeSrcTrafficDataRequest) SetDomainName(v string) *DescribeDcdnDomainRealTimeSrcTrafficDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDcdnDomainRealTimeSrcTrafficDataRequest) SetEndTime(v string) *DescribeDcdnDomainRealTimeSrcTrafficDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDcdnDomainRealTimeSrcTrafficDataRequest) SetOwnerId(v int64) *DescribeDcdnDomainRealTimeSrcTrafficDataRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDcdnDomainRealTimeSrcTrafficDataRequest) SetStartTime(v string) *DescribeDcdnDomainRealTimeSrcTrafficDataRequest {
	s.StartTime = &v
	return s
}

type DescribeDcdnDomainRealTimeSrcTrafficDataResponseBody struct {
	DataInterval                      *string                                                                                `json:"DataInterval,omitempty" xml:"DataInterval,omitempty"`
	DomainName                        *string                                                                                `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EndTime                           *string                                                                                `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	RealTimeSrcTrafficDataPerInterval *DescribeDcdnDomainRealTimeSrcTrafficDataResponseBodyRealTimeSrcTrafficDataPerInterval `json:"RealTimeSrcTrafficDataPerInterval,omitempty" xml:"RealTimeSrcTrafficDataPerInterval,omitempty" type:"Struct"`
	RequestId                         *string                                                                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StartTime                         *string                                                                                `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDcdnDomainRealTimeSrcTrafficDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainRealTimeSrcTrafficDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainRealTimeSrcTrafficDataResponseBody) SetDataInterval(v string) *DescribeDcdnDomainRealTimeSrcTrafficDataResponseBody {
	s.DataInterval = &v
	return s
}

func (s *DescribeDcdnDomainRealTimeSrcTrafficDataResponseBody) SetDomainName(v string) *DescribeDcdnDomainRealTimeSrcTrafficDataResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeDcdnDomainRealTimeSrcTrafficDataResponseBody) SetEndTime(v string) *DescribeDcdnDomainRealTimeSrcTrafficDataResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeDcdnDomainRealTimeSrcTrafficDataResponseBody) SetRealTimeSrcTrafficDataPerInterval(v *DescribeDcdnDomainRealTimeSrcTrafficDataResponseBodyRealTimeSrcTrafficDataPerInterval) *DescribeDcdnDomainRealTimeSrcTrafficDataResponseBody {
	s.RealTimeSrcTrafficDataPerInterval = v
	return s
}

func (s *DescribeDcdnDomainRealTimeSrcTrafficDataResponseBody) SetRequestId(v string) *DescribeDcdnDomainRealTimeSrcTrafficDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDcdnDomainRealTimeSrcTrafficDataResponseBody) SetStartTime(v string) *DescribeDcdnDomainRealTimeSrcTrafficDataResponseBody {
	s.StartTime = &v
	return s
}

type DescribeDcdnDomainRealTimeSrcTrafficDataResponseBodyRealTimeSrcTrafficDataPerInterval struct {
	DataModule []*DescribeDcdnDomainRealTimeSrcTrafficDataResponseBodyRealTimeSrcTrafficDataPerIntervalDataModule `json:"DataModule,omitempty" xml:"DataModule,omitempty" type:"Repeated"`
}

func (s DescribeDcdnDomainRealTimeSrcTrafficDataResponseBodyRealTimeSrcTrafficDataPerInterval) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainRealTimeSrcTrafficDataResponseBodyRealTimeSrcTrafficDataPerInterval) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainRealTimeSrcTrafficDataResponseBodyRealTimeSrcTrafficDataPerInterval) SetDataModule(v []*DescribeDcdnDomainRealTimeSrcTrafficDataResponseBodyRealTimeSrcTrafficDataPerIntervalDataModule) *DescribeDcdnDomainRealTimeSrcTrafficDataResponseBodyRealTimeSrcTrafficDataPerInterval {
	s.DataModule = v
	return s
}

type DescribeDcdnDomainRealTimeSrcTrafficDataResponseBodyRealTimeSrcTrafficDataPerIntervalDataModule struct {
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
	Value     *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDcdnDomainRealTimeSrcTrafficDataResponseBodyRealTimeSrcTrafficDataPerIntervalDataModule) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainRealTimeSrcTrafficDataResponseBodyRealTimeSrcTrafficDataPerIntervalDataModule) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainRealTimeSrcTrafficDataResponseBodyRealTimeSrcTrafficDataPerIntervalDataModule) SetTimeStamp(v string) *DescribeDcdnDomainRealTimeSrcTrafficDataResponseBodyRealTimeSrcTrafficDataPerIntervalDataModule {
	s.TimeStamp = &v
	return s
}

func (s *DescribeDcdnDomainRealTimeSrcTrafficDataResponseBodyRealTimeSrcTrafficDataPerIntervalDataModule) SetValue(v string) *DescribeDcdnDomainRealTimeSrcTrafficDataResponseBodyRealTimeSrcTrafficDataPerIntervalDataModule {
	s.Value = &v
	return s
}

type DescribeDcdnDomainRealTimeSrcTrafficDataResponse struct {
	Headers map[string]*string                                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDcdnDomainRealTimeSrcTrafficDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDcdnDomainRealTimeSrcTrafficDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainRealTimeSrcTrafficDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainRealTimeSrcTrafficDataResponse) SetHeaders(v map[string]*string) *DescribeDcdnDomainRealTimeSrcTrafficDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnDomainRealTimeSrcTrafficDataResponse) SetBody(v *DescribeDcdnDomainRealTimeSrcTrafficDataResponseBody) *DescribeDcdnDomainRealTimeSrcTrafficDataResponse {
	s.Body = v
	return s
}

type DescribeDcdnDomainRealTimeTrafficDataRequest struct {
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EndTime    *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	StartTime  *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDcdnDomainRealTimeTrafficDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainRealTimeTrafficDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainRealTimeTrafficDataRequest) SetDomainName(v string) *DescribeDcdnDomainRealTimeTrafficDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDcdnDomainRealTimeTrafficDataRequest) SetEndTime(v string) *DescribeDcdnDomainRealTimeTrafficDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDcdnDomainRealTimeTrafficDataRequest) SetOwnerId(v int64) *DescribeDcdnDomainRealTimeTrafficDataRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDcdnDomainRealTimeTrafficDataRequest) SetStartTime(v string) *DescribeDcdnDomainRealTimeTrafficDataRequest {
	s.StartTime = &v
	return s
}

type DescribeDcdnDomainRealTimeTrafficDataResponseBody struct {
	DataInterval                   *string                                                                          `json:"DataInterval,omitempty" xml:"DataInterval,omitempty"`
	DomainName                     *string                                                                          `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EndTime                        *string                                                                          `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	RealTimeTrafficDataPerInterval *DescribeDcdnDomainRealTimeTrafficDataResponseBodyRealTimeTrafficDataPerInterval `json:"RealTimeTrafficDataPerInterval,omitempty" xml:"RealTimeTrafficDataPerInterval,omitempty" type:"Struct"`
	RequestId                      *string                                                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StartTime                      *string                                                                          `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDcdnDomainRealTimeTrafficDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainRealTimeTrafficDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainRealTimeTrafficDataResponseBody) SetDataInterval(v string) *DescribeDcdnDomainRealTimeTrafficDataResponseBody {
	s.DataInterval = &v
	return s
}

func (s *DescribeDcdnDomainRealTimeTrafficDataResponseBody) SetDomainName(v string) *DescribeDcdnDomainRealTimeTrafficDataResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeDcdnDomainRealTimeTrafficDataResponseBody) SetEndTime(v string) *DescribeDcdnDomainRealTimeTrafficDataResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeDcdnDomainRealTimeTrafficDataResponseBody) SetRealTimeTrafficDataPerInterval(v *DescribeDcdnDomainRealTimeTrafficDataResponseBodyRealTimeTrafficDataPerInterval) *DescribeDcdnDomainRealTimeTrafficDataResponseBody {
	s.RealTimeTrafficDataPerInterval = v
	return s
}

func (s *DescribeDcdnDomainRealTimeTrafficDataResponseBody) SetRequestId(v string) *DescribeDcdnDomainRealTimeTrafficDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDcdnDomainRealTimeTrafficDataResponseBody) SetStartTime(v string) *DescribeDcdnDomainRealTimeTrafficDataResponseBody {
	s.StartTime = &v
	return s
}

type DescribeDcdnDomainRealTimeTrafficDataResponseBodyRealTimeTrafficDataPerInterval struct {
	DataModule []*DescribeDcdnDomainRealTimeTrafficDataResponseBodyRealTimeTrafficDataPerIntervalDataModule `json:"DataModule,omitempty" xml:"DataModule,omitempty" type:"Repeated"`
}

func (s DescribeDcdnDomainRealTimeTrafficDataResponseBodyRealTimeTrafficDataPerInterval) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainRealTimeTrafficDataResponseBodyRealTimeTrafficDataPerInterval) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainRealTimeTrafficDataResponseBodyRealTimeTrafficDataPerInterval) SetDataModule(v []*DescribeDcdnDomainRealTimeTrafficDataResponseBodyRealTimeTrafficDataPerIntervalDataModule) *DescribeDcdnDomainRealTimeTrafficDataResponseBodyRealTimeTrafficDataPerInterval {
	s.DataModule = v
	return s
}

type DescribeDcdnDomainRealTimeTrafficDataResponseBodyRealTimeTrafficDataPerIntervalDataModule struct {
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
	Value     *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDcdnDomainRealTimeTrafficDataResponseBodyRealTimeTrafficDataPerIntervalDataModule) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainRealTimeTrafficDataResponseBodyRealTimeTrafficDataPerIntervalDataModule) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainRealTimeTrafficDataResponseBodyRealTimeTrafficDataPerIntervalDataModule) SetTimeStamp(v string) *DescribeDcdnDomainRealTimeTrafficDataResponseBodyRealTimeTrafficDataPerIntervalDataModule {
	s.TimeStamp = &v
	return s
}

func (s *DescribeDcdnDomainRealTimeTrafficDataResponseBodyRealTimeTrafficDataPerIntervalDataModule) SetValue(v string) *DescribeDcdnDomainRealTimeTrafficDataResponseBodyRealTimeTrafficDataPerIntervalDataModule {
	s.Value = &v
	return s
}

type DescribeDcdnDomainRealTimeTrafficDataResponse struct {
	Headers map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDcdnDomainRealTimeTrafficDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDcdnDomainRealTimeTrafficDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainRealTimeTrafficDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainRealTimeTrafficDataResponse) SetHeaders(v map[string]*string) *DescribeDcdnDomainRealTimeTrafficDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnDomainRealTimeTrafficDataResponse) SetBody(v *DescribeDcdnDomainRealTimeTrafficDataResponseBody) *DescribeDcdnDomainRealTimeTrafficDataResponse {
	s.Body = v
	return s
}

type DescribeDcdnDomainRegionDataRequest struct {
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EndTime    *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	StartTime  *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDcdnDomainRegionDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainRegionDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainRegionDataRequest) SetDomainName(v string) *DescribeDcdnDomainRegionDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDcdnDomainRegionDataRequest) SetEndTime(v string) *DescribeDcdnDomainRegionDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDcdnDomainRegionDataRequest) SetOwnerId(v int64) *DescribeDcdnDomainRegionDataRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDcdnDomainRegionDataRequest) SetStartTime(v string) *DescribeDcdnDomainRegionDataRequest {
	s.StartTime = &v
	return s
}

type DescribeDcdnDomainRegionDataResponseBody struct {
	DataInterval *string                                        `json:"DataInterval,omitempty" xml:"DataInterval,omitempty"`
	DomainName   *string                                        `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EndTime      *string                                        `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	RequestId    *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StartTime    *string                                        `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Value        *DescribeDcdnDomainRegionDataResponseBodyValue `json:"Value,omitempty" xml:"Value,omitempty" type:"Struct"`
}

func (s DescribeDcdnDomainRegionDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainRegionDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainRegionDataResponseBody) SetDataInterval(v string) *DescribeDcdnDomainRegionDataResponseBody {
	s.DataInterval = &v
	return s
}

func (s *DescribeDcdnDomainRegionDataResponseBody) SetDomainName(v string) *DescribeDcdnDomainRegionDataResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeDcdnDomainRegionDataResponseBody) SetEndTime(v string) *DescribeDcdnDomainRegionDataResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeDcdnDomainRegionDataResponseBody) SetRequestId(v string) *DescribeDcdnDomainRegionDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDcdnDomainRegionDataResponseBody) SetStartTime(v string) *DescribeDcdnDomainRegionDataResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeDcdnDomainRegionDataResponseBody) SetValue(v *DescribeDcdnDomainRegionDataResponseBodyValue) *DescribeDcdnDomainRegionDataResponseBody {
	s.Value = v
	return s
}

type DescribeDcdnDomainRegionDataResponseBodyValue struct {
	RegionProportionData []*DescribeDcdnDomainRegionDataResponseBodyValueRegionProportionData `json:"RegionProportionData,omitempty" xml:"RegionProportionData,omitempty" type:"Repeated"`
}

func (s DescribeDcdnDomainRegionDataResponseBodyValue) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainRegionDataResponseBodyValue) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainRegionDataResponseBodyValue) SetRegionProportionData(v []*DescribeDcdnDomainRegionDataResponseBodyValueRegionProportionData) *DescribeDcdnDomainRegionDataResponseBodyValue {
	s.RegionProportionData = v
	return s
}

type DescribeDcdnDomainRegionDataResponseBodyValueRegionProportionData struct {
	AvgObjectSize   *string `json:"AvgObjectSize,omitempty" xml:"AvgObjectSize,omitempty"`
	AvgResponseRate *string `json:"AvgResponseRate,omitempty" xml:"AvgResponseRate,omitempty"`
	AvgResponseTime *string `json:"AvgResponseTime,omitempty" xml:"AvgResponseTime,omitempty"`
	Bps             *string `json:"Bps,omitempty" xml:"Bps,omitempty"`
	BytesProportion *string `json:"BytesProportion,omitempty" xml:"BytesProportion,omitempty"`
	Proportion      *string `json:"Proportion,omitempty" xml:"Proportion,omitempty"`
	Qps             *string `json:"Qps,omitempty" xml:"Qps,omitempty"`
	Region          *string `json:"Region,omitempty" xml:"Region,omitempty"`
	RegionEname     *string `json:"RegionEname,omitempty" xml:"RegionEname,omitempty"`
	TotalBytes      *string `json:"TotalBytes,omitempty" xml:"TotalBytes,omitempty"`
	TotalQuery      *string `json:"TotalQuery,omitempty" xml:"TotalQuery,omitempty"`
}

func (s DescribeDcdnDomainRegionDataResponseBodyValueRegionProportionData) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainRegionDataResponseBodyValueRegionProportionData) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainRegionDataResponseBodyValueRegionProportionData) SetAvgObjectSize(v string) *DescribeDcdnDomainRegionDataResponseBodyValueRegionProportionData {
	s.AvgObjectSize = &v
	return s
}

func (s *DescribeDcdnDomainRegionDataResponseBodyValueRegionProportionData) SetAvgResponseRate(v string) *DescribeDcdnDomainRegionDataResponseBodyValueRegionProportionData {
	s.AvgResponseRate = &v
	return s
}

func (s *DescribeDcdnDomainRegionDataResponseBodyValueRegionProportionData) SetAvgResponseTime(v string) *DescribeDcdnDomainRegionDataResponseBodyValueRegionProportionData {
	s.AvgResponseTime = &v
	return s
}

func (s *DescribeDcdnDomainRegionDataResponseBodyValueRegionProportionData) SetBps(v string) *DescribeDcdnDomainRegionDataResponseBodyValueRegionProportionData {
	s.Bps = &v
	return s
}

func (s *DescribeDcdnDomainRegionDataResponseBodyValueRegionProportionData) SetBytesProportion(v string) *DescribeDcdnDomainRegionDataResponseBodyValueRegionProportionData {
	s.BytesProportion = &v
	return s
}

func (s *DescribeDcdnDomainRegionDataResponseBodyValueRegionProportionData) SetProportion(v string) *DescribeDcdnDomainRegionDataResponseBodyValueRegionProportionData {
	s.Proportion = &v
	return s
}

func (s *DescribeDcdnDomainRegionDataResponseBodyValueRegionProportionData) SetQps(v string) *DescribeDcdnDomainRegionDataResponseBodyValueRegionProportionData {
	s.Qps = &v
	return s
}

func (s *DescribeDcdnDomainRegionDataResponseBodyValueRegionProportionData) SetRegion(v string) *DescribeDcdnDomainRegionDataResponseBodyValueRegionProportionData {
	s.Region = &v
	return s
}

func (s *DescribeDcdnDomainRegionDataResponseBodyValueRegionProportionData) SetRegionEname(v string) *DescribeDcdnDomainRegionDataResponseBodyValueRegionProportionData {
	s.RegionEname = &v
	return s
}

func (s *DescribeDcdnDomainRegionDataResponseBodyValueRegionProportionData) SetTotalBytes(v string) *DescribeDcdnDomainRegionDataResponseBodyValueRegionProportionData {
	s.TotalBytes = &v
	return s
}

func (s *DescribeDcdnDomainRegionDataResponseBodyValueRegionProportionData) SetTotalQuery(v string) *DescribeDcdnDomainRegionDataResponseBodyValueRegionProportionData {
	s.TotalQuery = &v
	return s
}

type DescribeDcdnDomainRegionDataResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDcdnDomainRegionDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDcdnDomainRegionDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainRegionDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainRegionDataResponse) SetHeaders(v map[string]*string) *DescribeDcdnDomainRegionDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnDomainRegionDataResponse) SetBody(v *DescribeDcdnDomainRegionDataResponseBody) *DescribeDcdnDomainRegionDataResponse {
	s.Body = v
	return s
}

type DescribeDcdnDomainStagingConfigRequest struct {
	DomainName    *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	FunctionNames *string `json:"FunctionNames,omitempty" xml:"FunctionNames,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s DescribeDcdnDomainStagingConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainStagingConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainStagingConfigRequest) SetDomainName(v string) *DescribeDcdnDomainStagingConfigRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDcdnDomainStagingConfigRequest) SetFunctionNames(v string) *DescribeDcdnDomainStagingConfigRequest {
	s.FunctionNames = &v
	return s
}

func (s *DescribeDcdnDomainStagingConfigRequest) SetOwnerId(v int64) *DescribeDcdnDomainStagingConfigRequest {
	s.OwnerId = &v
	return s
}

type DescribeDcdnDomainStagingConfigResponseBody struct {
	DomainConfigs []*DescribeDcdnDomainStagingConfigResponseBodyDomainConfigs `json:"DomainConfigs,omitempty" xml:"DomainConfigs,omitempty" type:"Repeated"`
	RequestId     *string                                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDcdnDomainStagingConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainStagingConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainStagingConfigResponseBody) SetDomainConfigs(v []*DescribeDcdnDomainStagingConfigResponseBodyDomainConfigs) *DescribeDcdnDomainStagingConfigResponseBody {
	s.DomainConfigs = v
	return s
}

func (s *DescribeDcdnDomainStagingConfigResponseBody) SetRequestId(v string) *DescribeDcdnDomainStagingConfigResponseBody {
	s.RequestId = &v
	return s
}

type DescribeDcdnDomainStagingConfigResponseBodyDomainConfigs struct {
	ConfigId     *string                                                                 `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	FunctionArgs []*DescribeDcdnDomainStagingConfigResponseBodyDomainConfigsFunctionArgs `json:"FunctionArgs,omitempty" xml:"FunctionArgs,omitempty" type:"Repeated"`
	FunctionName *string                                                                 `json:"FunctionName,omitempty" xml:"FunctionName,omitempty"`
	Status       *string                                                                 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeDcdnDomainStagingConfigResponseBodyDomainConfigs) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainStagingConfigResponseBodyDomainConfigs) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainStagingConfigResponseBodyDomainConfigs) SetConfigId(v string) *DescribeDcdnDomainStagingConfigResponseBodyDomainConfigs {
	s.ConfigId = &v
	return s
}

func (s *DescribeDcdnDomainStagingConfigResponseBodyDomainConfigs) SetFunctionArgs(v []*DescribeDcdnDomainStagingConfigResponseBodyDomainConfigsFunctionArgs) *DescribeDcdnDomainStagingConfigResponseBodyDomainConfigs {
	s.FunctionArgs = v
	return s
}

func (s *DescribeDcdnDomainStagingConfigResponseBodyDomainConfigs) SetFunctionName(v string) *DescribeDcdnDomainStagingConfigResponseBodyDomainConfigs {
	s.FunctionName = &v
	return s
}

func (s *DescribeDcdnDomainStagingConfigResponseBodyDomainConfigs) SetStatus(v string) *DescribeDcdnDomainStagingConfigResponseBodyDomainConfigs {
	s.Status = &v
	return s
}

type DescribeDcdnDomainStagingConfigResponseBodyDomainConfigsFunctionArgs struct {
	ArgName  *string `json:"ArgName,omitempty" xml:"ArgName,omitempty"`
	ArgValue *string `json:"ArgValue,omitempty" xml:"ArgValue,omitempty"`
}

func (s DescribeDcdnDomainStagingConfigResponseBodyDomainConfigsFunctionArgs) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainStagingConfigResponseBodyDomainConfigsFunctionArgs) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainStagingConfigResponseBodyDomainConfigsFunctionArgs) SetArgName(v string) *DescribeDcdnDomainStagingConfigResponseBodyDomainConfigsFunctionArgs {
	s.ArgName = &v
	return s
}

func (s *DescribeDcdnDomainStagingConfigResponseBodyDomainConfigsFunctionArgs) SetArgValue(v string) *DescribeDcdnDomainStagingConfigResponseBodyDomainConfigsFunctionArgs {
	s.ArgValue = &v
	return s
}

type DescribeDcdnDomainStagingConfigResponse struct {
	Headers map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDcdnDomainStagingConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDcdnDomainStagingConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainStagingConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainStagingConfigResponse) SetHeaders(v map[string]*string) *DescribeDcdnDomainStagingConfigResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnDomainStagingConfigResponse) SetBody(v *DescribeDcdnDomainStagingConfigResponseBody) *DescribeDcdnDomainStagingConfigResponse {
	s.Body = v
	return s
}

type DescribeDcdnDomainTopReferVisitRequest struct {
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SortBy     *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	StartTime  *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDcdnDomainTopReferVisitRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainTopReferVisitRequest) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainTopReferVisitRequest) SetDomainName(v string) *DescribeDcdnDomainTopReferVisitRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDcdnDomainTopReferVisitRequest) SetOwnerId(v int64) *DescribeDcdnDomainTopReferVisitRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDcdnDomainTopReferVisitRequest) SetSortBy(v string) *DescribeDcdnDomainTopReferVisitRequest {
	s.SortBy = &v
	return s
}

func (s *DescribeDcdnDomainTopReferVisitRequest) SetStartTime(v string) *DescribeDcdnDomainTopReferVisitRequest {
	s.StartTime = &v
	return s
}

type DescribeDcdnDomainTopReferVisitResponseBody struct {
	DomainName   *string                                                  `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	RequestId    *string                                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StartTime    *string                                                  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	TopReferList *DescribeDcdnDomainTopReferVisitResponseBodyTopReferList `json:"TopReferList,omitempty" xml:"TopReferList,omitempty" type:"Struct"`
}

func (s DescribeDcdnDomainTopReferVisitResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainTopReferVisitResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainTopReferVisitResponseBody) SetDomainName(v string) *DescribeDcdnDomainTopReferVisitResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeDcdnDomainTopReferVisitResponseBody) SetRequestId(v string) *DescribeDcdnDomainTopReferVisitResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDcdnDomainTopReferVisitResponseBody) SetStartTime(v string) *DescribeDcdnDomainTopReferVisitResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeDcdnDomainTopReferVisitResponseBody) SetTopReferList(v *DescribeDcdnDomainTopReferVisitResponseBodyTopReferList) *DescribeDcdnDomainTopReferVisitResponseBody {
	s.TopReferList = v
	return s
}

type DescribeDcdnDomainTopReferVisitResponseBodyTopReferList struct {
	ReferList []*DescribeDcdnDomainTopReferVisitResponseBodyTopReferListReferList `json:"ReferList,omitempty" xml:"ReferList,omitempty" type:"Repeated"`
}

func (s DescribeDcdnDomainTopReferVisitResponseBodyTopReferList) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainTopReferVisitResponseBodyTopReferList) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainTopReferVisitResponseBodyTopReferList) SetReferList(v []*DescribeDcdnDomainTopReferVisitResponseBodyTopReferListReferList) *DescribeDcdnDomainTopReferVisitResponseBodyTopReferList {
	s.ReferList = v
	return s
}

type DescribeDcdnDomainTopReferVisitResponseBodyTopReferListReferList struct {
	Flow            *string  `json:"Flow,omitempty" xml:"Flow,omitempty"`
	FlowProportion  *float32 `json:"FlowProportion,omitempty" xml:"FlowProportion,omitempty"`
	ReferDetail     *string  `json:"ReferDetail,omitempty" xml:"ReferDetail,omitempty"`
	VisitData       *string  `json:"VisitData,omitempty" xml:"VisitData,omitempty"`
	VisitProportion *float32 `json:"VisitProportion,omitempty" xml:"VisitProportion,omitempty"`
}

func (s DescribeDcdnDomainTopReferVisitResponseBodyTopReferListReferList) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainTopReferVisitResponseBodyTopReferListReferList) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainTopReferVisitResponseBodyTopReferListReferList) SetFlow(v string) *DescribeDcdnDomainTopReferVisitResponseBodyTopReferListReferList {
	s.Flow = &v
	return s
}

func (s *DescribeDcdnDomainTopReferVisitResponseBodyTopReferListReferList) SetFlowProportion(v float32) *DescribeDcdnDomainTopReferVisitResponseBodyTopReferListReferList {
	s.FlowProportion = &v
	return s
}

func (s *DescribeDcdnDomainTopReferVisitResponseBodyTopReferListReferList) SetReferDetail(v string) *DescribeDcdnDomainTopReferVisitResponseBodyTopReferListReferList {
	s.ReferDetail = &v
	return s
}

func (s *DescribeDcdnDomainTopReferVisitResponseBodyTopReferListReferList) SetVisitData(v string) *DescribeDcdnDomainTopReferVisitResponseBodyTopReferListReferList {
	s.VisitData = &v
	return s
}

func (s *DescribeDcdnDomainTopReferVisitResponseBodyTopReferListReferList) SetVisitProportion(v float32) *DescribeDcdnDomainTopReferVisitResponseBodyTopReferListReferList {
	s.VisitProportion = &v
	return s
}

type DescribeDcdnDomainTopReferVisitResponse struct {
	Headers map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDcdnDomainTopReferVisitResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDcdnDomainTopReferVisitResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainTopReferVisitResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainTopReferVisitResponse) SetHeaders(v map[string]*string) *DescribeDcdnDomainTopReferVisitResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnDomainTopReferVisitResponse) SetBody(v *DescribeDcdnDomainTopReferVisitResponseBody) *DescribeDcdnDomainTopReferVisitResponse {
	s.Body = v
	return s
}

type DescribeDcdnDomainTopUrlVisitRequest struct {
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SortBy     *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	StartTime  *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDcdnDomainTopUrlVisitRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainTopUrlVisitRequest) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainTopUrlVisitRequest) SetDomainName(v string) *DescribeDcdnDomainTopUrlVisitRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDcdnDomainTopUrlVisitRequest) SetOwnerId(v int64) *DescribeDcdnDomainTopUrlVisitRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDcdnDomainTopUrlVisitRequest) SetSortBy(v string) *DescribeDcdnDomainTopUrlVisitRequest {
	s.SortBy = &v
	return s
}

func (s *DescribeDcdnDomainTopUrlVisitRequest) SetStartTime(v string) *DescribeDcdnDomainTopUrlVisitRequest {
	s.StartTime = &v
	return s
}

type DescribeDcdnDomainTopUrlVisitResponseBody struct {
	AllUrlList *DescribeDcdnDomainTopUrlVisitResponseBodyAllUrlList `json:"AllUrlList,omitempty" xml:"AllUrlList,omitempty" type:"Struct"`
	DomainName *string                                              `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	RequestId  *string                                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StartTime  *string                                              `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Url200List *DescribeDcdnDomainTopUrlVisitResponseBodyUrl200List `json:"Url200List,omitempty" xml:"Url200List,omitempty" type:"Struct"`
	Url300List *DescribeDcdnDomainTopUrlVisitResponseBodyUrl300List `json:"Url300List,omitempty" xml:"Url300List,omitempty" type:"Struct"`
	Url400List *DescribeDcdnDomainTopUrlVisitResponseBodyUrl400List `json:"Url400List,omitempty" xml:"Url400List,omitempty" type:"Struct"`
	Url500List *DescribeDcdnDomainTopUrlVisitResponseBodyUrl500List `json:"Url500List,omitempty" xml:"Url500List,omitempty" type:"Struct"`
}

func (s DescribeDcdnDomainTopUrlVisitResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainTopUrlVisitResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainTopUrlVisitResponseBody) SetAllUrlList(v *DescribeDcdnDomainTopUrlVisitResponseBodyAllUrlList) *DescribeDcdnDomainTopUrlVisitResponseBody {
	s.AllUrlList = v
	return s
}

func (s *DescribeDcdnDomainTopUrlVisitResponseBody) SetDomainName(v string) *DescribeDcdnDomainTopUrlVisitResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeDcdnDomainTopUrlVisitResponseBody) SetRequestId(v string) *DescribeDcdnDomainTopUrlVisitResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDcdnDomainTopUrlVisitResponseBody) SetStartTime(v string) *DescribeDcdnDomainTopUrlVisitResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeDcdnDomainTopUrlVisitResponseBody) SetUrl200List(v *DescribeDcdnDomainTopUrlVisitResponseBodyUrl200List) *DescribeDcdnDomainTopUrlVisitResponseBody {
	s.Url200List = v
	return s
}

func (s *DescribeDcdnDomainTopUrlVisitResponseBody) SetUrl300List(v *DescribeDcdnDomainTopUrlVisitResponseBodyUrl300List) *DescribeDcdnDomainTopUrlVisitResponseBody {
	s.Url300List = v
	return s
}

func (s *DescribeDcdnDomainTopUrlVisitResponseBody) SetUrl400List(v *DescribeDcdnDomainTopUrlVisitResponseBodyUrl400List) *DescribeDcdnDomainTopUrlVisitResponseBody {
	s.Url400List = v
	return s
}

func (s *DescribeDcdnDomainTopUrlVisitResponseBody) SetUrl500List(v *DescribeDcdnDomainTopUrlVisitResponseBodyUrl500List) *DescribeDcdnDomainTopUrlVisitResponseBody {
	s.Url500List = v
	return s
}

type DescribeDcdnDomainTopUrlVisitResponseBodyAllUrlList struct {
	UrlList []*DescribeDcdnDomainTopUrlVisitResponseBodyAllUrlListUrlList `json:"UrlList,omitempty" xml:"UrlList,omitempty" type:"Repeated"`
}

func (s DescribeDcdnDomainTopUrlVisitResponseBodyAllUrlList) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainTopUrlVisitResponseBodyAllUrlList) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainTopUrlVisitResponseBodyAllUrlList) SetUrlList(v []*DescribeDcdnDomainTopUrlVisitResponseBodyAllUrlListUrlList) *DescribeDcdnDomainTopUrlVisitResponseBodyAllUrlList {
	s.UrlList = v
	return s
}

type DescribeDcdnDomainTopUrlVisitResponseBodyAllUrlListUrlList struct {
	Flow            *string  `json:"Flow,omitempty" xml:"Flow,omitempty"`
	FlowProportion  *float32 `json:"FlowProportion,omitempty" xml:"FlowProportion,omitempty"`
	UrlDetail       *string  `json:"UrlDetail,omitempty" xml:"UrlDetail,omitempty"`
	VisitData       *string  `json:"VisitData,omitempty" xml:"VisitData,omitempty"`
	VisitProportion *float32 `json:"VisitProportion,omitempty" xml:"VisitProportion,omitempty"`
}

func (s DescribeDcdnDomainTopUrlVisitResponseBodyAllUrlListUrlList) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainTopUrlVisitResponseBodyAllUrlListUrlList) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainTopUrlVisitResponseBodyAllUrlListUrlList) SetFlow(v string) *DescribeDcdnDomainTopUrlVisitResponseBodyAllUrlListUrlList {
	s.Flow = &v
	return s
}

func (s *DescribeDcdnDomainTopUrlVisitResponseBodyAllUrlListUrlList) SetFlowProportion(v float32) *DescribeDcdnDomainTopUrlVisitResponseBodyAllUrlListUrlList {
	s.FlowProportion = &v
	return s
}

func (s *DescribeDcdnDomainTopUrlVisitResponseBodyAllUrlListUrlList) SetUrlDetail(v string) *DescribeDcdnDomainTopUrlVisitResponseBodyAllUrlListUrlList {
	s.UrlDetail = &v
	return s
}

func (s *DescribeDcdnDomainTopUrlVisitResponseBodyAllUrlListUrlList) SetVisitData(v string) *DescribeDcdnDomainTopUrlVisitResponseBodyAllUrlListUrlList {
	s.VisitData = &v
	return s
}

func (s *DescribeDcdnDomainTopUrlVisitResponseBodyAllUrlListUrlList) SetVisitProportion(v float32) *DescribeDcdnDomainTopUrlVisitResponseBodyAllUrlListUrlList {
	s.VisitProportion = &v
	return s
}

type DescribeDcdnDomainTopUrlVisitResponseBodyUrl200List struct {
	UrlList []*DescribeDcdnDomainTopUrlVisitResponseBodyUrl200ListUrlList `json:"UrlList,omitempty" xml:"UrlList,omitempty" type:"Repeated"`
}

func (s DescribeDcdnDomainTopUrlVisitResponseBodyUrl200List) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainTopUrlVisitResponseBodyUrl200List) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainTopUrlVisitResponseBodyUrl200List) SetUrlList(v []*DescribeDcdnDomainTopUrlVisitResponseBodyUrl200ListUrlList) *DescribeDcdnDomainTopUrlVisitResponseBodyUrl200List {
	s.UrlList = v
	return s
}

type DescribeDcdnDomainTopUrlVisitResponseBodyUrl200ListUrlList struct {
	Flow            *string  `json:"Flow,omitempty" xml:"Flow,omitempty"`
	FlowProportion  *float32 `json:"FlowProportion,omitempty" xml:"FlowProportion,omitempty"`
	UrlDetail       *string  `json:"UrlDetail,omitempty" xml:"UrlDetail,omitempty"`
	VisitData       *string  `json:"VisitData,omitempty" xml:"VisitData,omitempty"`
	VisitProportion *float32 `json:"VisitProportion,omitempty" xml:"VisitProportion,omitempty"`
}

func (s DescribeDcdnDomainTopUrlVisitResponseBodyUrl200ListUrlList) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainTopUrlVisitResponseBodyUrl200ListUrlList) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainTopUrlVisitResponseBodyUrl200ListUrlList) SetFlow(v string) *DescribeDcdnDomainTopUrlVisitResponseBodyUrl200ListUrlList {
	s.Flow = &v
	return s
}

func (s *DescribeDcdnDomainTopUrlVisitResponseBodyUrl200ListUrlList) SetFlowProportion(v float32) *DescribeDcdnDomainTopUrlVisitResponseBodyUrl200ListUrlList {
	s.FlowProportion = &v
	return s
}

func (s *DescribeDcdnDomainTopUrlVisitResponseBodyUrl200ListUrlList) SetUrlDetail(v string) *DescribeDcdnDomainTopUrlVisitResponseBodyUrl200ListUrlList {
	s.UrlDetail = &v
	return s
}

func (s *DescribeDcdnDomainTopUrlVisitResponseBodyUrl200ListUrlList) SetVisitData(v string) *DescribeDcdnDomainTopUrlVisitResponseBodyUrl200ListUrlList {
	s.VisitData = &v
	return s
}

func (s *DescribeDcdnDomainTopUrlVisitResponseBodyUrl200ListUrlList) SetVisitProportion(v float32) *DescribeDcdnDomainTopUrlVisitResponseBodyUrl200ListUrlList {
	s.VisitProportion = &v
	return s
}

type DescribeDcdnDomainTopUrlVisitResponseBodyUrl300List struct {
	UrlList []*DescribeDcdnDomainTopUrlVisitResponseBodyUrl300ListUrlList `json:"UrlList,omitempty" xml:"UrlList,omitempty" type:"Repeated"`
}

func (s DescribeDcdnDomainTopUrlVisitResponseBodyUrl300List) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainTopUrlVisitResponseBodyUrl300List) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainTopUrlVisitResponseBodyUrl300List) SetUrlList(v []*DescribeDcdnDomainTopUrlVisitResponseBodyUrl300ListUrlList) *DescribeDcdnDomainTopUrlVisitResponseBodyUrl300List {
	s.UrlList = v
	return s
}

type DescribeDcdnDomainTopUrlVisitResponseBodyUrl300ListUrlList struct {
	Flow            *string  `json:"Flow,omitempty" xml:"Flow,omitempty"`
	FlowProportion  *float32 `json:"FlowProportion,omitempty" xml:"FlowProportion,omitempty"`
	UrlDetail       *string  `json:"UrlDetail,omitempty" xml:"UrlDetail,omitempty"`
	VisitData       *string  `json:"VisitData,omitempty" xml:"VisitData,omitempty"`
	VisitProportion *float32 `json:"VisitProportion,omitempty" xml:"VisitProportion,omitempty"`
}

func (s DescribeDcdnDomainTopUrlVisitResponseBodyUrl300ListUrlList) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainTopUrlVisitResponseBodyUrl300ListUrlList) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainTopUrlVisitResponseBodyUrl300ListUrlList) SetFlow(v string) *DescribeDcdnDomainTopUrlVisitResponseBodyUrl300ListUrlList {
	s.Flow = &v
	return s
}

func (s *DescribeDcdnDomainTopUrlVisitResponseBodyUrl300ListUrlList) SetFlowProportion(v float32) *DescribeDcdnDomainTopUrlVisitResponseBodyUrl300ListUrlList {
	s.FlowProportion = &v
	return s
}

func (s *DescribeDcdnDomainTopUrlVisitResponseBodyUrl300ListUrlList) SetUrlDetail(v string) *DescribeDcdnDomainTopUrlVisitResponseBodyUrl300ListUrlList {
	s.UrlDetail = &v
	return s
}

func (s *DescribeDcdnDomainTopUrlVisitResponseBodyUrl300ListUrlList) SetVisitData(v string) *DescribeDcdnDomainTopUrlVisitResponseBodyUrl300ListUrlList {
	s.VisitData = &v
	return s
}

func (s *DescribeDcdnDomainTopUrlVisitResponseBodyUrl300ListUrlList) SetVisitProportion(v float32) *DescribeDcdnDomainTopUrlVisitResponseBodyUrl300ListUrlList {
	s.VisitProportion = &v
	return s
}

type DescribeDcdnDomainTopUrlVisitResponseBodyUrl400List struct {
	UrlList []*DescribeDcdnDomainTopUrlVisitResponseBodyUrl400ListUrlList `json:"UrlList,omitempty" xml:"UrlList,omitempty" type:"Repeated"`
}

func (s DescribeDcdnDomainTopUrlVisitResponseBodyUrl400List) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainTopUrlVisitResponseBodyUrl400List) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainTopUrlVisitResponseBodyUrl400List) SetUrlList(v []*DescribeDcdnDomainTopUrlVisitResponseBodyUrl400ListUrlList) *DescribeDcdnDomainTopUrlVisitResponseBodyUrl400List {
	s.UrlList = v
	return s
}

type DescribeDcdnDomainTopUrlVisitResponseBodyUrl400ListUrlList struct {
	Flow            *string  `json:"Flow,omitempty" xml:"Flow,omitempty"`
	FlowProportion  *float32 `json:"FlowProportion,omitempty" xml:"FlowProportion,omitempty"`
	UrlDetail       *string  `json:"UrlDetail,omitempty" xml:"UrlDetail,omitempty"`
	VisitData       *string  `json:"VisitData,omitempty" xml:"VisitData,omitempty"`
	VisitProportion *float32 `json:"VisitProportion,omitempty" xml:"VisitProportion,omitempty"`
}

func (s DescribeDcdnDomainTopUrlVisitResponseBodyUrl400ListUrlList) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainTopUrlVisitResponseBodyUrl400ListUrlList) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainTopUrlVisitResponseBodyUrl400ListUrlList) SetFlow(v string) *DescribeDcdnDomainTopUrlVisitResponseBodyUrl400ListUrlList {
	s.Flow = &v
	return s
}

func (s *DescribeDcdnDomainTopUrlVisitResponseBodyUrl400ListUrlList) SetFlowProportion(v float32) *DescribeDcdnDomainTopUrlVisitResponseBodyUrl400ListUrlList {
	s.FlowProportion = &v
	return s
}

func (s *DescribeDcdnDomainTopUrlVisitResponseBodyUrl400ListUrlList) SetUrlDetail(v string) *DescribeDcdnDomainTopUrlVisitResponseBodyUrl400ListUrlList {
	s.UrlDetail = &v
	return s
}

func (s *DescribeDcdnDomainTopUrlVisitResponseBodyUrl400ListUrlList) SetVisitData(v string) *DescribeDcdnDomainTopUrlVisitResponseBodyUrl400ListUrlList {
	s.VisitData = &v
	return s
}

func (s *DescribeDcdnDomainTopUrlVisitResponseBodyUrl400ListUrlList) SetVisitProportion(v float32) *DescribeDcdnDomainTopUrlVisitResponseBodyUrl400ListUrlList {
	s.VisitProportion = &v
	return s
}

type DescribeDcdnDomainTopUrlVisitResponseBodyUrl500List struct {
	UrlList []*DescribeDcdnDomainTopUrlVisitResponseBodyUrl500ListUrlList `json:"UrlList,omitempty" xml:"UrlList,omitempty" type:"Repeated"`
}

func (s DescribeDcdnDomainTopUrlVisitResponseBodyUrl500List) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainTopUrlVisitResponseBodyUrl500List) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainTopUrlVisitResponseBodyUrl500List) SetUrlList(v []*DescribeDcdnDomainTopUrlVisitResponseBodyUrl500ListUrlList) *DescribeDcdnDomainTopUrlVisitResponseBodyUrl500List {
	s.UrlList = v
	return s
}

type DescribeDcdnDomainTopUrlVisitResponseBodyUrl500ListUrlList struct {
	Flow            *string  `json:"Flow,omitempty" xml:"Flow,omitempty"`
	FlowProportion  *float32 `json:"FlowProportion,omitempty" xml:"FlowProportion,omitempty"`
	UrlDetail       *string  `json:"UrlDetail,omitempty" xml:"UrlDetail,omitempty"`
	VisitData       *string  `json:"VisitData,omitempty" xml:"VisitData,omitempty"`
	VisitProportion *float32 `json:"VisitProportion,omitempty" xml:"VisitProportion,omitempty"`
}

func (s DescribeDcdnDomainTopUrlVisitResponseBodyUrl500ListUrlList) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainTopUrlVisitResponseBodyUrl500ListUrlList) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainTopUrlVisitResponseBodyUrl500ListUrlList) SetFlow(v string) *DescribeDcdnDomainTopUrlVisitResponseBodyUrl500ListUrlList {
	s.Flow = &v
	return s
}

func (s *DescribeDcdnDomainTopUrlVisitResponseBodyUrl500ListUrlList) SetFlowProportion(v float32) *DescribeDcdnDomainTopUrlVisitResponseBodyUrl500ListUrlList {
	s.FlowProportion = &v
	return s
}

func (s *DescribeDcdnDomainTopUrlVisitResponseBodyUrl500ListUrlList) SetUrlDetail(v string) *DescribeDcdnDomainTopUrlVisitResponseBodyUrl500ListUrlList {
	s.UrlDetail = &v
	return s
}

func (s *DescribeDcdnDomainTopUrlVisitResponseBodyUrl500ListUrlList) SetVisitData(v string) *DescribeDcdnDomainTopUrlVisitResponseBodyUrl500ListUrlList {
	s.VisitData = &v
	return s
}

func (s *DescribeDcdnDomainTopUrlVisitResponseBodyUrl500ListUrlList) SetVisitProportion(v float32) *DescribeDcdnDomainTopUrlVisitResponseBodyUrl500ListUrlList {
	s.VisitProportion = &v
	return s
}

type DescribeDcdnDomainTopUrlVisitResponse struct {
	Headers map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDcdnDomainTopUrlVisitResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDcdnDomainTopUrlVisitResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainTopUrlVisitResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainTopUrlVisitResponse) SetHeaders(v map[string]*string) *DescribeDcdnDomainTopUrlVisitResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnDomainTopUrlVisitResponse) SetBody(v *DescribeDcdnDomainTopUrlVisitResponseBody) *DescribeDcdnDomainTopUrlVisitResponse {
	s.Body = v
	return s
}

type DescribeDcdnDomainTrafficDataRequest struct {
	DomainName     *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EndTime        *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Interval       *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	IspNameEn      *string `json:"IspNameEn,omitempty" xml:"IspNameEn,omitempty"`
	LocationNameEn *string `json:"LocationNameEn,omitempty" xml:"LocationNameEn,omitempty"`
	OwnerId        *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	StartTime      *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDcdnDomainTrafficDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainTrafficDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainTrafficDataRequest) SetDomainName(v string) *DescribeDcdnDomainTrafficDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDcdnDomainTrafficDataRequest) SetEndTime(v string) *DescribeDcdnDomainTrafficDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDcdnDomainTrafficDataRequest) SetInterval(v string) *DescribeDcdnDomainTrafficDataRequest {
	s.Interval = &v
	return s
}

func (s *DescribeDcdnDomainTrafficDataRequest) SetIspNameEn(v string) *DescribeDcdnDomainTrafficDataRequest {
	s.IspNameEn = &v
	return s
}

func (s *DescribeDcdnDomainTrafficDataRequest) SetLocationNameEn(v string) *DescribeDcdnDomainTrafficDataRequest {
	s.LocationNameEn = &v
	return s
}

func (s *DescribeDcdnDomainTrafficDataRequest) SetOwnerId(v int64) *DescribeDcdnDomainTrafficDataRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDcdnDomainTrafficDataRequest) SetStartTime(v string) *DescribeDcdnDomainTrafficDataRequest {
	s.StartTime = &v
	return s
}

type DescribeDcdnDomainTrafficDataResponseBody struct {
	DataInterval           *string                                                          `json:"DataInterval,omitempty" xml:"DataInterval,omitempty"`
	DomainName             *string                                                          `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EndTime                *string                                                          `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	RequestId              *string                                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StartTime              *string                                                          `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	TrafficDataPerInterval *DescribeDcdnDomainTrafficDataResponseBodyTrafficDataPerInterval `json:"TrafficDataPerInterval,omitempty" xml:"TrafficDataPerInterval,omitempty" type:"Struct"`
}

func (s DescribeDcdnDomainTrafficDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainTrafficDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainTrafficDataResponseBody) SetDataInterval(v string) *DescribeDcdnDomainTrafficDataResponseBody {
	s.DataInterval = &v
	return s
}

func (s *DescribeDcdnDomainTrafficDataResponseBody) SetDomainName(v string) *DescribeDcdnDomainTrafficDataResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeDcdnDomainTrafficDataResponseBody) SetEndTime(v string) *DescribeDcdnDomainTrafficDataResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeDcdnDomainTrafficDataResponseBody) SetRequestId(v string) *DescribeDcdnDomainTrafficDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDcdnDomainTrafficDataResponseBody) SetStartTime(v string) *DescribeDcdnDomainTrafficDataResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeDcdnDomainTrafficDataResponseBody) SetTrafficDataPerInterval(v *DescribeDcdnDomainTrafficDataResponseBodyTrafficDataPerInterval) *DescribeDcdnDomainTrafficDataResponseBody {
	s.TrafficDataPerInterval = v
	return s
}

type DescribeDcdnDomainTrafficDataResponseBodyTrafficDataPerInterval struct {
	DataModule []*DescribeDcdnDomainTrafficDataResponseBodyTrafficDataPerIntervalDataModule `json:"DataModule,omitempty" xml:"DataModule,omitempty" type:"Repeated"`
}

func (s DescribeDcdnDomainTrafficDataResponseBodyTrafficDataPerInterval) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainTrafficDataResponseBodyTrafficDataPerInterval) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainTrafficDataResponseBodyTrafficDataPerInterval) SetDataModule(v []*DescribeDcdnDomainTrafficDataResponseBodyTrafficDataPerIntervalDataModule) *DescribeDcdnDomainTrafficDataResponseBodyTrafficDataPerInterval {
	s.DataModule = v
	return s
}

type DescribeDcdnDomainTrafficDataResponseBodyTrafficDataPerIntervalDataModule struct {
	DynamicHttpTraffic  *float32 `json:"DynamicHttpTraffic,omitempty" xml:"DynamicHttpTraffic,omitempty"`
	DynamicHttpsTraffic *float32 `json:"DynamicHttpsTraffic,omitempty" xml:"DynamicHttpsTraffic,omitempty"`
	StaticHttpTraffic   *float32 `json:"StaticHttpTraffic,omitempty" xml:"StaticHttpTraffic,omitempty"`
	StaticHttpsTraffic  *float32 `json:"StaticHttpsTraffic,omitempty" xml:"StaticHttpsTraffic,omitempty"`
	TimeStamp           *string  `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
	Traffic             *float32 `json:"Traffic,omitempty" xml:"Traffic,omitempty"`
}

func (s DescribeDcdnDomainTrafficDataResponseBodyTrafficDataPerIntervalDataModule) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainTrafficDataResponseBodyTrafficDataPerIntervalDataModule) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainTrafficDataResponseBodyTrafficDataPerIntervalDataModule) SetDynamicHttpTraffic(v float32) *DescribeDcdnDomainTrafficDataResponseBodyTrafficDataPerIntervalDataModule {
	s.DynamicHttpTraffic = &v
	return s
}

func (s *DescribeDcdnDomainTrafficDataResponseBodyTrafficDataPerIntervalDataModule) SetDynamicHttpsTraffic(v float32) *DescribeDcdnDomainTrafficDataResponseBodyTrafficDataPerIntervalDataModule {
	s.DynamicHttpsTraffic = &v
	return s
}

func (s *DescribeDcdnDomainTrafficDataResponseBodyTrafficDataPerIntervalDataModule) SetStaticHttpTraffic(v float32) *DescribeDcdnDomainTrafficDataResponseBodyTrafficDataPerIntervalDataModule {
	s.StaticHttpTraffic = &v
	return s
}

func (s *DescribeDcdnDomainTrafficDataResponseBodyTrafficDataPerIntervalDataModule) SetStaticHttpsTraffic(v float32) *DescribeDcdnDomainTrafficDataResponseBodyTrafficDataPerIntervalDataModule {
	s.StaticHttpsTraffic = &v
	return s
}

func (s *DescribeDcdnDomainTrafficDataResponseBodyTrafficDataPerIntervalDataModule) SetTimeStamp(v string) *DescribeDcdnDomainTrafficDataResponseBodyTrafficDataPerIntervalDataModule {
	s.TimeStamp = &v
	return s
}

func (s *DescribeDcdnDomainTrafficDataResponseBodyTrafficDataPerIntervalDataModule) SetTraffic(v float32) *DescribeDcdnDomainTrafficDataResponseBodyTrafficDataPerIntervalDataModule {
	s.Traffic = &v
	return s
}

type DescribeDcdnDomainTrafficDataResponse struct {
	Headers map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDcdnDomainTrafficDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDcdnDomainTrafficDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainTrafficDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainTrafficDataResponse) SetHeaders(v map[string]*string) *DescribeDcdnDomainTrafficDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnDomainTrafficDataResponse) SetBody(v *DescribeDcdnDomainTrafficDataResponseBody) *DescribeDcdnDomainTrafficDataResponse {
	s.Body = v
	return s
}

type DescribeDcdnDomainUsageDataRequest struct {
	Area         *string `json:"Area,omitempty" xml:"Area,omitempty"`
	DataProtocol *string `json:"DataProtocol,omitempty" xml:"DataProtocol,omitempty"`
	DomainName   *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EndTime      *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Field        *string `json:"Field,omitempty" xml:"Field,omitempty"`
	Interval     *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	StartTime    *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDcdnDomainUsageDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainUsageDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainUsageDataRequest) SetArea(v string) *DescribeDcdnDomainUsageDataRequest {
	s.Area = &v
	return s
}

func (s *DescribeDcdnDomainUsageDataRequest) SetDataProtocol(v string) *DescribeDcdnDomainUsageDataRequest {
	s.DataProtocol = &v
	return s
}

func (s *DescribeDcdnDomainUsageDataRequest) SetDomainName(v string) *DescribeDcdnDomainUsageDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDcdnDomainUsageDataRequest) SetEndTime(v string) *DescribeDcdnDomainUsageDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDcdnDomainUsageDataRequest) SetField(v string) *DescribeDcdnDomainUsageDataRequest {
	s.Field = &v
	return s
}

func (s *DescribeDcdnDomainUsageDataRequest) SetInterval(v string) *DescribeDcdnDomainUsageDataRequest {
	s.Interval = &v
	return s
}

func (s *DescribeDcdnDomainUsageDataRequest) SetOwnerId(v int64) *DescribeDcdnDomainUsageDataRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDcdnDomainUsageDataRequest) SetStartTime(v string) *DescribeDcdnDomainUsageDataRequest {
	s.StartTime = &v
	return s
}

type DescribeDcdnDomainUsageDataResponseBody struct {
	Area                 *string                                                      `json:"Area,omitempty" xml:"Area,omitempty"`
	DataInterval         *string                                                      `json:"DataInterval,omitempty" xml:"DataInterval,omitempty"`
	DomainName           *string                                                      `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EndTime              *string                                                      `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	RequestId            *string                                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StartTime            *string                                                      `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Type                 *string                                                      `json:"Type,omitempty" xml:"Type,omitempty"`
	UsageDataPerInterval *DescribeDcdnDomainUsageDataResponseBodyUsageDataPerInterval `json:"UsageDataPerInterval,omitempty" xml:"UsageDataPerInterval,omitempty" type:"Struct"`
}

func (s DescribeDcdnDomainUsageDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainUsageDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainUsageDataResponseBody) SetArea(v string) *DescribeDcdnDomainUsageDataResponseBody {
	s.Area = &v
	return s
}

func (s *DescribeDcdnDomainUsageDataResponseBody) SetDataInterval(v string) *DescribeDcdnDomainUsageDataResponseBody {
	s.DataInterval = &v
	return s
}

func (s *DescribeDcdnDomainUsageDataResponseBody) SetDomainName(v string) *DescribeDcdnDomainUsageDataResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeDcdnDomainUsageDataResponseBody) SetEndTime(v string) *DescribeDcdnDomainUsageDataResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeDcdnDomainUsageDataResponseBody) SetRequestId(v string) *DescribeDcdnDomainUsageDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDcdnDomainUsageDataResponseBody) SetStartTime(v string) *DescribeDcdnDomainUsageDataResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeDcdnDomainUsageDataResponseBody) SetType(v string) *DescribeDcdnDomainUsageDataResponseBody {
	s.Type = &v
	return s
}

func (s *DescribeDcdnDomainUsageDataResponseBody) SetUsageDataPerInterval(v *DescribeDcdnDomainUsageDataResponseBodyUsageDataPerInterval) *DescribeDcdnDomainUsageDataResponseBody {
	s.UsageDataPerInterval = v
	return s
}

type DescribeDcdnDomainUsageDataResponseBodyUsageDataPerInterval struct {
	DataModule []*DescribeDcdnDomainUsageDataResponseBodyUsageDataPerIntervalDataModule `json:"DataModule,omitempty" xml:"DataModule,omitempty" type:"Repeated"`
}

func (s DescribeDcdnDomainUsageDataResponseBodyUsageDataPerInterval) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainUsageDataResponseBodyUsageDataPerInterval) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainUsageDataResponseBodyUsageDataPerInterval) SetDataModule(v []*DescribeDcdnDomainUsageDataResponseBodyUsageDataPerIntervalDataModule) *DescribeDcdnDomainUsageDataResponseBodyUsageDataPerInterval {
	s.DataModule = v
	return s
}

type DescribeDcdnDomainUsageDataResponseBodyUsageDataPerIntervalDataModule struct {
	PeakTime     *string `json:"PeakTime,omitempty" xml:"PeakTime,omitempty"`
	SpecialValue *string `json:"SpecialValue,omitempty" xml:"SpecialValue,omitempty"`
	TimeStamp    *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
	Value        *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDcdnDomainUsageDataResponseBodyUsageDataPerIntervalDataModule) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainUsageDataResponseBodyUsageDataPerIntervalDataModule) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainUsageDataResponseBodyUsageDataPerIntervalDataModule) SetPeakTime(v string) *DescribeDcdnDomainUsageDataResponseBodyUsageDataPerIntervalDataModule {
	s.PeakTime = &v
	return s
}

func (s *DescribeDcdnDomainUsageDataResponseBodyUsageDataPerIntervalDataModule) SetSpecialValue(v string) *DescribeDcdnDomainUsageDataResponseBodyUsageDataPerIntervalDataModule {
	s.SpecialValue = &v
	return s
}

func (s *DescribeDcdnDomainUsageDataResponseBodyUsageDataPerIntervalDataModule) SetTimeStamp(v string) *DescribeDcdnDomainUsageDataResponseBodyUsageDataPerIntervalDataModule {
	s.TimeStamp = &v
	return s
}

func (s *DescribeDcdnDomainUsageDataResponseBodyUsageDataPerIntervalDataModule) SetValue(v string) *DescribeDcdnDomainUsageDataResponseBodyUsageDataPerIntervalDataModule {
	s.Value = &v
	return s
}

type DescribeDcdnDomainUsageDataResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDcdnDomainUsageDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDcdnDomainUsageDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainUsageDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainUsageDataResponse) SetHeaders(v map[string]*string) *DescribeDcdnDomainUsageDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnDomainUsageDataResponse) SetBody(v *DescribeDcdnDomainUsageDataResponseBody) *DescribeDcdnDomainUsageDataResponse {
	s.Body = v
	return s
}

type DescribeDcdnDomainUvDataRequest struct {
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EndTime    *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	StartTime  *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDcdnDomainUvDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainUvDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainUvDataRequest) SetDomainName(v string) *DescribeDcdnDomainUvDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDcdnDomainUvDataRequest) SetEndTime(v string) *DescribeDcdnDomainUvDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDcdnDomainUvDataRequest) SetOwnerId(v int64) *DescribeDcdnDomainUvDataRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDcdnDomainUvDataRequest) SetStartTime(v string) *DescribeDcdnDomainUvDataRequest {
	s.StartTime = &v
	return s
}

type DescribeDcdnDomainUvDataResponseBody struct {
	DataInterval   *string                                             `json:"DataInterval,omitempty" xml:"DataInterval,omitempty"`
	DomainName     *string                                             `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EndTime        *string                                             `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	RequestId      *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StartTime      *string                                             `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	UvDataInterval *DescribeDcdnDomainUvDataResponseBodyUvDataInterval `json:"UvDataInterval,omitempty" xml:"UvDataInterval,omitempty" type:"Struct"`
}

func (s DescribeDcdnDomainUvDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainUvDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainUvDataResponseBody) SetDataInterval(v string) *DescribeDcdnDomainUvDataResponseBody {
	s.DataInterval = &v
	return s
}

func (s *DescribeDcdnDomainUvDataResponseBody) SetDomainName(v string) *DescribeDcdnDomainUvDataResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeDcdnDomainUvDataResponseBody) SetEndTime(v string) *DescribeDcdnDomainUvDataResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeDcdnDomainUvDataResponseBody) SetRequestId(v string) *DescribeDcdnDomainUvDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDcdnDomainUvDataResponseBody) SetStartTime(v string) *DescribeDcdnDomainUvDataResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeDcdnDomainUvDataResponseBody) SetUvDataInterval(v *DescribeDcdnDomainUvDataResponseBodyUvDataInterval) *DescribeDcdnDomainUvDataResponseBody {
	s.UvDataInterval = v
	return s
}

type DescribeDcdnDomainUvDataResponseBodyUvDataInterval struct {
	UsageData []*DescribeDcdnDomainUvDataResponseBodyUvDataIntervalUsageData `json:"UsageData,omitempty" xml:"UsageData,omitempty" type:"Repeated"`
}

func (s DescribeDcdnDomainUvDataResponseBodyUvDataInterval) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainUvDataResponseBodyUvDataInterval) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainUvDataResponseBodyUvDataInterval) SetUsageData(v []*DescribeDcdnDomainUvDataResponseBodyUvDataIntervalUsageData) *DescribeDcdnDomainUvDataResponseBodyUvDataInterval {
	s.UsageData = v
	return s
}

type DescribeDcdnDomainUvDataResponseBodyUvDataIntervalUsageData struct {
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
	Value     *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDcdnDomainUvDataResponseBodyUvDataIntervalUsageData) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainUvDataResponseBodyUvDataIntervalUsageData) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainUvDataResponseBodyUvDataIntervalUsageData) SetTimeStamp(v string) *DescribeDcdnDomainUvDataResponseBodyUvDataIntervalUsageData {
	s.TimeStamp = &v
	return s
}

func (s *DescribeDcdnDomainUvDataResponseBodyUvDataIntervalUsageData) SetValue(v string) *DescribeDcdnDomainUvDataResponseBodyUvDataIntervalUsageData {
	s.Value = &v
	return s
}

type DescribeDcdnDomainUvDataResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDcdnDomainUvDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDcdnDomainUvDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainUvDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainUvDataResponse) SetHeaders(v map[string]*string) *DescribeDcdnDomainUvDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnDomainUvDataResponse) SetBody(v *DescribeDcdnDomainUvDataResponseBody) *DescribeDcdnDomainUvDataResponse {
	s.Body = v
	return s
}

type DescribeDcdnDomainWebsocketBpsDataRequest struct {
	DomainName     *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EndTime        *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Interval       *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	IspNameEn      *string `json:"IspNameEn,omitempty" xml:"IspNameEn,omitempty"`
	LocationNameEn *string `json:"LocationNameEn,omitempty" xml:"LocationNameEn,omitempty"`
	OwnerId        *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	StartTime      *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDcdnDomainWebsocketBpsDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainWebsocketBpsDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainWebsocketBpsDataRequest) SetDomainName(v string) *DescribeDcdnDomainWebsocketBpsDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDcdnDomainWebsocketBpsDataRequest) SetEndTime(v string) *DescribeDcdnDomainWebsocketBpsDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDcdnDomainWebsocketBpsDataRequest) SetInterval(v string) *DescribeDcdnDomainWebsocketBpsDataRequest {
	s.Interval = &v
	return s
}

func (s *DescribeDcdnDomainWebsocketBpsDataRequest) SetIspNameEn(v string) *DescribeDcdnDomainWebsocketBpsDataRequest {
	s.IspNameEn = &v
	return s
}

func (s *DescribeDcdnDomainWebsocketBpsDataRequest) SetLocationNameEn(v string) *DescribeDcdnDomainWebsocketBpsDataRequest {
	s.LocationNameEn = &v
	return s
}

func (s *DescribeDcdnDomainWebsocketBpsDataRequest) SetOwnerId(v int64) *DescribeDcdnDomainWebsocketBpsDataRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDcdnDomainWebsocketBpsDataRequest) SetStartTime(v string) *DescribeDcdnDomainWebsocketBpsDataRequest {
	s.StartTime = &v
	return s
}

type DescribeDcdnDomainWebsocketBpsDataResponseBody struct {
	BpsDataPerInterval *DescribeDcdnDomainWebsocketBpsDataResponseBodyBpsDataPerInterval `json:"BpsDataPerInterval,omitempty" xml:"BpsDataPerInterval,omitempty" type:"Struct"`
	DataInterval       *string                                                           `json:"DataInterval,omitempty" xml:"DataInterval,omitempty"`
	DomainName         *string                                                           `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EndTime            *string                                                           `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	RequestId          *string                                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StartTime          *string                                                           `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDcdnDomainWebsocketBpsDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainWebsocketBpsDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainWebsocketBpsDataResponseBody) SetBpsDataPerInterval(v *DescribeDcdnDomainWebsocketBpsDataResponseBodyBpsDataPerInterval) *DescribeDcdnDomainWebsocketBpsDataResponseBody {
	s.BpsDataPerInterval = v
	return s
}

func (s *DescribeDcdnDomainWebsocketBpsDataResponseBody) SetDataInterval(v string) *DescribeDcdnDomainWebsocketBpsDataResponseBody {
	s.DataInterval = &v
	return s
}

func (s *DescribeDcdnDomainWebsocketBpsDataResponseBody) SetDomainName(v string) *DescribeDcdnDomainWebsocketBpsDataResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeDcdnDomainWebsocketBpsDataResponseBody) SetEndTime(v string) *DescribeDcdnDomainWebsocketBpsDataResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeDcdnDomainWebsocketBpsDataResponseBody) SetRequestId(v string) *DescribeDcdnDomainWebsocketBpsDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDcdnDomainWebsocketBpsDataResponseBody) SetStartTime(v string) *DescribeDcdnDomainWebsocketBpsDataResponseBody {
	s.StartTime = &v
	return s
}

type DescribeDcdnDomainWebsocketBpsDataResponseBodyBpsDataPerInterval struct {
	DataModule []*DescribeDcdnDomainWebsocketBpsDataResponseBodyBpsDataPerIntervalDataModule `json:"DataModule,omitempty" xml:"DataModule,omitempty" type:"Repeated"`
}

func (s DescribeDcdnDomainWebsocketBpsDataResponseBodyBpsDataPerInterval) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainWebsocketBpsDataResponseBodyBpsDataPerInterval) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainWebsocketBpsDataResponseBodyBpsDataPerInterval) SetDataModule(v []*DescribeDcdnDomainWebsocketBpsDataResponseBodyBpsDataPerIntervalDataModule) *DescribeDcdnDomainWebsocketBpsDataResponseBodyBpsDataPerInterval {
	s.DataModule = v
	return s
}

type DescribeDcdnDomainWebsocketBpsDataResponseBodyBpsDataPerIntervalDataModule struct {
	TimeStamp    *string  `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
	WebsocketBps *float32 `json:"WebsocketBps,omitempty" xml:"WebsocketBps,omitempty"`
}

func (s DescribeDcdnDomainWebsocketBpsDataResponseBodyBpsDataPerIntervalDataModule) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainWebsocketBpsDataResponseBodyBpsDataPerIntervalDataModule) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainWebsocketBpsDataResponseBodyBpsDataPerIntervalDataModule) SetTimeStamp(v string) *DescribeDcdnDomainWebsocketBpsDataResponseBodyBpsDataPerIntervalDataModule {
	s.TimeStamp = &v
	return s
}

func (s *DescribeDcdnDomainWebsocketBpsDataResponseBodyBpsDataPerIntervalDataModule) SetWebsocketBps(v float32) *DescribeDcdnDomainWebsocketBpsDataResponseBodyBpsDataPerIntervalDataModule {
	s.WebsocketBps = &v
	return s
}

type DescribeDcdnDomainWebsocketBpsDataResponse struct {
	Headers map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDcdnDomainWebsocketBpsDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDcdnDomainWebsocketBpsDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainWebsocketBpsDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainWebsocketBpsDataResponse) SetHeaders(v map[string]*string) *DescribeDcdnDomainWebsocketBpsDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnDomainWebsocketBpsDataResponse) SetBody(v *DescribeDcdnDomainWebsocketBpsDataResponseBody) *DescribeDcdnDomainWebsocketBpsDataResponse {
	s.Body = v
	return s
}

type DescribeDcdnDomainWebsocketHttpCodeDataRequest struct {
	DomainName     *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EndTime        *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Interval       *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	IspNameEn      *string `json:"IspNameEn,omitempty" xml:"IspNameEn,omitempty"`
	LocationNameEn *string `json:"LocationNameEn,omitempty" xml:"LocationNameEn,omitempty"`
	OwnerId        *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	StartTime      *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDcdnDomainWebsocketHttpCodeDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainWebsocketHttpCodeDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainWebsocketHttpCodeDataRequest) SetDomainName(v string) *DescribeDcdnDomainWebsocketHttpCodeDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDcdnDomainWebsocketHttpCodeDataRequest) SetEndTime(v string) *DescribeDcdnDomainWebsocketHttpCodeDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDcdnDomainWebsocketHttpCodeDataRequest) SetInterval(v string) *DescribeDcdnDomainWebsocketHttpCodeDataRequest {
	s.Interval = &v
	return s
}

func (s *DescribeDcdnDomainWebsocketHttpCodeDataRequest) SetIspNameEn(v string) *DescribeDcdnDomainWebsocketHttpCodeDataRequest {
	s.IspNameEn = &v
	return s
}

func (s *DescribeDcdnDomainWebsocketHttpCodeDataRequest) SetLocationNameEn(v string) *DescribeDcdnDomainWebsocketHttpCodeDataRequest {
	s.LocationNameEn = &v
	return s
}

func (s *DescribeDcdnDomainWebsocketHttpCodeDataRequest) SetOwnerId(v int64) *DescribeDcdnDomainWebsocketHttpCodeDataRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDcdnDomainWebsocketHttpCodeDataRequest) SetStartTime(v string) *DescribeDcdnDomainWebsocketHttpCodeDataRequest {
	s.StartTime = &v
	return s
}

type DescribeDcdnDomainWebsocketHttpCodeDataResponseBody struct {
	DataInterval            *string                                                                     `json:"DataInterval,omitempty" xml:"DataInterval,omitempty"`
	DomainName              *string                                                                     `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EndTime                 *string                                                                     `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	HttpCodeDataPerInterval *DescribeDcdnDomainWebsocketHttpCodeDataResponseBodyHttpCodeDataPerInterval `json:"HttpCodeDataPerInterval,omitempty" xml:"HttpCodeDataPerInterval,omitempty" type:"Struct"`
	RequestId               *string                                                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StartTime               *string                                                                     `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDcdnDomainWebsocketHttpCodeDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainWebsocketHttpCodeDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainWebsocketHttpCodeDataResponseBody) SetDataInterval(v string) *DescribeDcdnDomainWebsocketHttpCodeDataResponseBody {
	s.DataInterval = &v
	return s
}

func (s *DescribeDcdnDomainWebsocketHttpCodeDataResponseBody) SetDomainName(v string) *DescribeDcdnDomainWebsocketHttpCodeDataResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeDcdnDomainWebsocketHttpCodeDataResponseBody) SetEndTime(v string) *DescribeDcdnDomainWebsocketHttpCodeDataResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeDcdnDomainWebsocketHttpCodeDataResponseBody) SetHttpCodeDataPerInterval(v *DescribeDcdnDomainWebsocketHttpCodeDataResponseBodyHttpCodeDataPerInterval) *DescribeDcdnDomainWebsocketHttpCodeDataResponseBody {
	s.HttpCodeDataPerInterval = v
	return s
}

func (s *DescribeDcdnDomainWebsocketHttpCodeDataResponseBody) SetRequestId(v string) *DescribeDcdnDomainWebsocketHttpCodeDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDcdnDomainWebsocketHttpCodeDataResponseBody) SetStartTime(v string) *DescribeDcdnDomainWebsocketHttpCodeDataResponseBody {
	s.StartTime = &v
	return s
}

type DescribeDcdnDomainWebsocketHttpCodeDataResponseBodyHttpCodeDataPerInterval struct {
	DataModule []*DescribeDcdnDomainWebsocketHttpCodeDataResponseBodyHttpCodeDataPerIntervalDataModule `json:"DataModule,omitempty" xml:"DataModule,omitempty" type:"Repeated"`
}

func (s DescribeDcdnDomainWebsocketHttpCodeDataResponseBodyHttpCodeDataPerInterval) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainWebsocketHttpCodeDataResponseBodyHttpCodeDataPerInterval) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainWebsocketHttpCodeDataResponseBodyHttpCodeDataPerInterval) SetDataModule(v []*DescribeDcdnDomainWebsocketHttpCodeDataResponseBodyHttpCodeDataPerIntervalDataModule) *DescribeDcdnDomainWebsocketHttpCodeDataResponseBodyHttpCodeDataPerInterval {
	s.DataModule = v
	return s
}

type DescribeDcdnDomainWebsocketHttpCodeDataResponseBodyHttpCodeDataPerIntervalDataModule struct {
	TimeStamp         *string                                                                                                `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
	WebsocketHttpCode *DescribeDcdnDomainWebsocketHttpCodeDataResponseBodyHttpCodeDataPerIntervalDataModuleWebsocketHttpCode `json:"WebsocketHttpCode,omitempty" xml:"WebsocketHttpCode,omitempty" type:"Struct"`
}

func (s DescribeDcdnDomainWebsocketHttpCodeDataResponseBodyHttpCodeDataPerIntervalDataModule) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainWebsocketHttpCodeDataResponseBodyHttpCodeDataPerIntervalDataModule) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainWebsocketHttpCodeDataResponseBodyHttpCodeDataPerIntervalDataModule) SetTimeStamp(v string) *DescribeDcdnDomainWebsocketHttpCodeDataResponseBodyHttpCodeDataPerIntervalDataModule {
	s.TimeStamp = &v
	return s
}

func (s *DescribeDcdnDomainWebsocketHttpCodeDataResponseBodyHttpCodeDataPerIntervalDataModule) SetWebsocketHttpCode(v *DescribeDcdnDomainWebsocketHttpCodeDataResponseBodyHttpCodeDataPerIntervalDataModuleWebsocketHttpCode) *DescribeDcdnDomainWebsocketHttpCodeDataResponseBodyHttpCodeDataPerIntervalDataModule {
	s.WebsocketHttpCode = v
	return s
}

type DescribeDcdnDomainWebsocketHttpCodeDataResponseBodyHttpCodeDataPerIntervalDataModuleWebsocketHttpCode struct {
	HttpCodeDataModule []*DescribeDcdnDomainWebsocketHttpCodeDataResponseBodyHttpCodeDataPerIntervalDataModuleWebsocketHttpCodeHttpCodeDataModule `json:"HttpCodeDataModule,omitempty" xml:"HttpCodeDataModule,omitempty" type:"Repeated"`
}

func (s DescribeDcdnDomainWebsocketHttpCodeDataResponseBodyHttpCodeDataPerIntervalDataModuleWebsocketHttpCode) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainWebsocketHttpCodeDataResponseBodyHttpCodeDataPerIntervalDataModuleWebsocketHttpCode) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainWebsocketHttpCodeDataResponseBodyHttpCodeDataPerIntervalDataModuleWebsocketHttpCode) SetHttpCodeDataModule(v []*DescribeDcdnDomainWebsocketHttpCodeDataResponseBodyHttpCodeDataPerIntervalDataModuleWebsocketHttpCodeHttpCodeDataModule) *DescribeDcdnDomainWebsocketHttpCodeDataResponseBodyHttpCodeDataPerIntervalDataModuleWebsocketHttpCode {
	s.HttpCodeDataModule = v
	return s
}

type DescribeDcdnDomainWebsocketHttpCodeDataResponseBodyHttpCodeDataPerIntervalDataModuleWebsocketHttpCodeHttpCodeDataModule struct {
	Code       *int32   `json:"Code,omitempty" xml:"Code,omitempty"`
	Count      *float32 `json:"Count,omitempty" xml:"Count,omitempty"`
	Proportion *float32 `json:"Proportion,omitempty" xml:"Proportion,omitempty"`
}

func (s DescribeDcdnDomainWebsocketHttpCodeDataResponseBodyHttpCodeDataPerIntervalDataModuleWebsocketHttpCodeHttpCodeDataModule) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainWebsocketHttpCodeDataResponseBodyHttpCodeDataPerIntervalDataModuleWebsocketHttpCodeHttpCodeDataModule) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainWebsocketHttpCodeDataResponseBodyHttpCodeDataPerIntervalDataModuleWebsocketHttpCodeHttpCodeDataModule) SetCode(v int32) *DescribeDcdnDomainWebsocketHttpCodeDataResponseBodyHttpCodeDataPerIntervalDataModuleWebsocketHttpCodeHttpCodeDataModule {
	s.Code = &v
	return s
}

func (s *DescribeDcdnDomainWebsocketHttpCodeDataResponseBodyHttpCodeDataPerIntervalDataModuleWebsocketHttpCodeHttpCodeDataModule) SetCount(v float32) *DescribeDcdnDomainWebsocketHttpCodeDataResponseBodyHttpCodeDataPerIntervalDataModuleWebsocketHttpCodeHttpCodeDataModule {
	s.Count = &v
	return s
}

func (s *DescribeDcdnDomainWebsocketHttpCodeDataResponseBodyHttpCodeDataPerIntervalDataModuleWebsocketHttpCodeHttpCodeDataModule) SetProportion(v float32) *DescribeDcdnDomainWebsocketHttpCodeDataResponseBodyHttpCodeDataPerIntervalDataModuleWebsocketHttpCodeHttpCodeDataModule {
	s.Proportion = &v
	return s
}

type DescribeDcdnDomainWebsocketHttpCodeDataResponse struct {
	Headers map[string]*string                                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDcdnDomainWebsocketHttpCodeDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDcdnDomainWebsocketHttpCodeDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainWebsocketHttpCodeDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainWebsocketHttpCodeDataResponse) SetHeaders(v map[string]*string) *DescribeDcdnDomainWebsocketHttpCodeDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnDomainWebsocketHttpCodeDataResponse) SetBody(v *DescribeDcdnDomainWebsocketHttpCodeDataResponseBody) *DescribeDcdnDomainWebsocketHttpCodeDataResponse {
	s.Body = v
	return s
}

type DescribeDcdnDomainWebsocketTrafficDataRequest struct {
	DomainName     *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EndTime        *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Interval       *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	IspNameEn      *string `json:"IspNameEn,omitempty" xml:"IspNameEn,omitempty"`
	LocationNameEn *string `json:"LocationNameEn,omitempty" xml:"LocationNameEn,omitempty"`
	OwnerId        *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	StartTime      *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDcdnDomainWebsocketTrafficDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainWebsocketTrafficDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainWebsocketTrafficDataRequest) SetDomainName(v string) *DescribeDcdnDomainWebsocketTrafficDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDcdnDomainWebsocketTrafficDataRequest) SetEndTime(v string) *DescribeDcdnDomainWebsocketTrafficDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDcdnDomainWebsocketTrafficDataRequest) SetInterval(v string) *DescribeDcdnDomainWebsocketTrafficDataRequest {
	s.Interval = &v
	return s
}

func (s *DescribeDcdnDomainWebsocketTrafficDataRequest) SetIspNameEn(v string) *DescribeDcdnDomainWebsocketTrafficDataRequest {
	s.IspNameEn = &v
	return s
}

func (s *DescribeDcdnDomainWebsocketTrafficDataRequest) SetLocationNameEn(v string) *DescribeDcdnDomainWebsocketTrafficDataRequest {
	s.LocationNameEn = &v
	return s
}

func (s *DescribeDcdnDomainWebsocketTrafficDataRequest) SetOwnerId(v int64) *DescribeDcdnDomainWebsocketTrafficDataRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDcdnDomainWebsocketTrafficDataRequest) SetStartTime(v string) *DescribeDcdnDomainWebsocketTrafficDataRequest {
	s.StartTime = &v
	return s
}

type DescribeDcdnDomainWebsocketTrafficDataResponseBody struct {
	DataInterval           *string                                                                   `json:"DataInterval,omitempty" xml:"DataInterval,omitempty"`
	DomainName             *string                                                                   `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EndTime                *string                                                                   `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	RequestId              *string                                                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StartTime              *string                                                                   `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	TrafficDataPerInterval *DescribeDcdnDomainWebsocketTrafficDataResponseBodyTrafficDataPerInterval `json:"TrafficDataPerInterval,omitempty" xml:"TrafficDataPerInterval,omitempty" type:"Struct"`
}

func (s DescribeDcdnDomainWebsocketTrafficDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainWebsocketTrafficDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainWebsocketTrafficDataResponseBody) SetDataInterval(v string) *DescribeDcdnDomainWebsocketTrafficDataResponseBody {
	s.DataInterval = &v
	return s
}

func (s *DescribeDcdnDomainWebsocketTrafficDataResponseBody) SetDomainName(v string) *DescribeDcdnDomainWebsocketTrafficDataResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeDcdnDomainWebsocketTrafficDataResponseBody) SetEndTime(v string) *DescribeDcdnDomainWebsocketTrafficDataResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeDcdnDomainWebsocketTrafficDataResponseBody) SetRequestId(v string) *DescribeDcdnDomainWebsocketTrafficDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDcdnDomainWebsocketTrafficDataResponseBody) SetStartTime(v string) *DescribeDcdnDomainWebsocketTrafficDataResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeDcdnDomainWebsocketTrafficDataResponseBody) SetTrafficDataPerInterval(v *DescribeDcdnDomainWebsocketTrafficDataResponseBodyTrafficDataPerInterval) *DescribeDcdnDomainWebsocketTrafficDataResponseBody {
	s.TrafficDataPerInterval = v
	return s
}

type DescribeDcdnDomainWebsocketTrafficDataResponseBodyTrafficDataPerInterval struct {
	DataModule []*DescribeDcdnDomainWebsocketTrafficDataResponseBodyTrafficDataPerIntervalDataModule `json:"DataModule,omitempty" xml:"DataModule,omitempty" type:"Repeated"`
}

func (s DescribeDcdnDomainWebsocketTrafficDataResponseBodyTrafficDataPerInterval) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainWebsocketTrafficDataResponseBodyTrafficDataPerInterval) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainWebsocketTrafficDataResponseBodyTrafficDataPerInterval) SetDataModule(v []*DescribeDcdnDomainWebsocketTrafficDataResponseBodyTrafficDataPerIntervalDataModule) *DescribeDcdnDomainWebsocketTrafficDataResponseBodyTrafficDataPerInterval {
	s.DataModule = v
	return s
}

type DescribeDcdnDomainWebsocketTrafficDataResponseBodyTrafficDataPerIntervalDataModule struct {
	TimeStamp        *string  `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
	WebsocketTraffic *float32 `json:"WebsocketTraffic,omitempty" xml:"WebsocketTraffic,omitempty"`
}

func (s DescribeDcdnDomainWebsocketTrafficDataResponseBodyTrafficDataPerIntervalDataModule) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainWebsocketTrafficDataResponseBodyTrafficDataPerIntervalDataModule) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainWebsocketTrafficDataResponseBodyTrafficDataPerIntervalDataModule) SetTimeStamp(v string) *DescribeDcdnDomainWebsocketTrafficDataResponseBodyTrafficDataPerIntervalDataModule {
	s.TimeStamp = &v
	return s
}

func (s *DescribeDcdnDomainWebsocketTrafficDataResponseBodyTrafficDataPerIntervalDataModule) SetWebsocketTraffic(v float32) *DescribeDcdnDomainWebsocketTrafficDataResponseBodyTrafficDataPerIntervalDataModule {
	s.WebsocketTraffic = &v
	return s
}

type DescribeDcdnDomainWebsocketTrafficDataResponse struct {
	Headers map[string]*string                                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDcdnDomainWebsocketTrafficDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDcdnDomainWebsocketTrafficDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnDomainWebsocketTrafficDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainWebsocketTrafficDataResponse) SetHeaders(v map[string]*string) *DescribeDcdnDomainWebsocketTrafficDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnDomainWebsocketTrafficDataResponse) SetBody(v *DescribeDcdnDomainWebsocketTrafficDataResponseBody) *DescribeDcdnDomainWebsocketTrafficDataResponse {
	s.Body = v
	return s
}

type DescribeDcdnEsExceptionDataRequest struct {
	EndTime   *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	OwnerId   *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RuleId    *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDcdnEsExceptionDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnEsExceptionDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeDcdnEsExceptionDataRequest) SetEndTime(v string) *DescribeDcdnEsExceptionDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDcdnEsExceptionDataRequest) SetOwnerId(v int64) *DescribeDcdnEsExceptionDataRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDcdnEsExceptionDataRequest) SetRuleId(v string) *DescribeDcdnEsExceptionDataRequest {
	s.RuleId = &v
	return s
}

func (s *DescribeDcdnEsExceptionDataRequest) SetStartTime(v string) *DescribeDcdnEsExceptionDataRequest {
	s.StartTime = &v
	return s
}

type DescribeDcdnEsExceptionDataResponseBody struct {
	Contents  []*DescribeDcdnEsExceptionDataResponseBodyContents `json:"Contents,omitempty" xml:"Contents,omitempty" type:"Repeated"`
	RequestId *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDcdnEsExceptionDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnEsExceptionDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnEsExceptionDataResponseBody) SetContents(v []*DescribeDcdnEsExceptionDataResponseBodyContents) *DescribeDcdnEsExceptionDataResponseBody {
	s.Contents = v
	return s
}

func (s *DescribeDcdnEsExceptionDataResponseBody) SetRequestId(v string) *DescribeDcdnEsExceptionDataResponseBody {
	s.RequestId = &v
	return s
}

type DescribeDcdnEsExceptionDataResponseBodyContents struct {
	Columns []*string `json:"Columns,omitempty" xml:"Columns,omitempty" type:"Repeated"`
	Name    *string   `json:"Name,omitempty" xml:"Name,omitempty"`
	Points  []*string `json:"Points,omitempty" xml:"Points,omitempty" type:"Repeated"`
}

func (s DescribeDcdnEsExceptionDataResponseBodyContents) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnEsExceptionDataResponseBodyContents) GoString() string {
	return s.String()
}

func (s *DescribeDcdnEsExceptionDataResponseBodyContents) SetColumns(v []*string) *DescribeDcdnEsExceptionDataResponseBodyContents {
	s.Columns = v
	return s
}

func (s *DescribeDcdnEsExceptionDataResponseBodyContents) SetName(v string) *DescribeDcdnEsExceptionDataResponseBodyContents {
	s.Name = &v
	return s
}

func (s *DescribeDcdnEsExceptionDataResponseBodyContents) SetPoints(v []*string) *DescribeDcdnEsExceptionDataResponseBodyContents {
	s.Points = v
	return s
}

type DescribeDcdnEsExceptionDataResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDcdnEsExceptionDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDcdnEsExceptionDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnEsExceptionDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnEsExceptionDataResponse) SetHeaders(v map[string]*string) *DescribeDcdnEsExceptionDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnEsExceptionDataResponse) SetBody(v *DescribeDcdnEsExceptionDataResponseBody) *DescribeDcdnEsExceptionDataResponse {
	s.Body = v
	return s
}

type DescribeDcdnEsExecuteDataRequest struct {
	EndTime   *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	OwnerId   *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RuleId    *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDcdnEsExecuteDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnEsExecuteDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeDcdnEsExecuteDataRequest) SetEndTime(v string) *DescribeDcdnEsExecuteDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDcdnEsExecuteDataRequest) SetOwnerId(v int64) *DescribeDcdnEsExecuteDataRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDcdnEsExecuteDataRequest) SetRuleId(v string) *DescribeDcdnEsExecuteDataRequest {
	s.RuleId = &v
	return s
}

func (s *DescribeDcdnEsExecuteDataRequest) SetStartTime(v string) *DescribeDcdnEsExecuteDataRequest {
	s.StartTime = &v
	return s
}

type DescribeDcdnEsExecuteDataResponseBody struct {
	Contents  []*DescribeDcdnEsExecuteDataResponseBodyContents `json:"Contents,omitempty" xml:"Contents,omitempty" type:"Repeated"`
	RequestId *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDcdnEsExecuteDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnEsExecuteDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnEsExecuteDataResponseBody) SetContents(v []*DescribeDcdnEsExecuteDataResponseBodyContents) *DescribeDcdnEsExecuteDataResponseBody {
	s.Contents = v
	return s
}

func (s *DescribeDcdnEsExecuteDataResponseBody) SetRequestId(v string) *DescribeDcdnEsExecuteDataResponseBody {
	s.RequestId = &v
	return s
}

type DescribeDcdnEsExecuteDataResponseBodyContents struct {
	Columns []*string `json:"Columns,omitempty" xml:"Columns,omitempty" type:"Repeated"`
	Name    *string   `json:"Name,omitempty" xml:"Name,omitempty"`
	Points  []*string `json:"Points,omitempty" xml:"Points,omitempty" type:"Repeated"`
}

func (s DescribeDcdnEsExecuteDataResponseBodyContents) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnEsExecuteDataResponseBodyContents) GoString() string {
	return s.String()
}

func (s *DescribeDcdnEsExecuteDataResponseBodyContents) SetColumns(v []*string) *DescribeDcdnEsExecuteDataResponseBodyContents {
	s.Columns = v
	return s
}

func (s *DescribeDcdnEsExecuteDataResponseBodyContents) SetName(v string) *DescribeDcdnEsExecuteDataResponseBodyContents {
	s.Name = &v
	return s
}

func (s *DescribeDcdnEsExecuteDataResponseBodyContents) SetPoints(v []*string) *DescribeDcdnEsExecuteDataResponseBodyContents {
	s.Points = v
	return s
}

type DescribeDcdnEsExecuteDataResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDcdnEsExecuteDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDcdnEsExecuteDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnEsExecuteDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnEsExecuteDataResponse) SetHeaders(v map[string]*string) *DescribeDcdnEsExecuteDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnEsExecuteDataResponse) SetBody(v *DescribeDcdnEsExecuteDataResponseBody) *DescribeDcdnEsExecuteDataResponse {
	s.Body = v
	return s
}

type DescribeDcdnHttpsDomainListRequest struct {
	Keyword    *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeDcdnHttpsDomainListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnHttpsDomainListRequest) GoString() string {
	return s.String()
}

func (s *DescribeDcdnHttpsDomainListRequest) SetKeyword(v string) *DescribeDcdnHttpsDomainListRequest {
	s.Keyword = &v
	return s
}

func (s *DescribeDcdnHttpsDomainListRequest) SetOwnerId(v int64) *DescribeDcdnHttpsDomainListRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDcdnHttpsDomainListRequest) SetPageNumber(v int32) *DescribeDcdnHttpsDomainListRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDcdnHttpsDomainListRequest) SetPageSize(v int32) *DescribeDcdnHttpsDomainListRequest {
	s.PageSize = &v
	return s
}

type DescribeDcdnHttpsDomainListResponseBody struct {
	CertInfos  *DescribeDcdnHttpsDomainListResponseBodyCertInfos `json:"CertInfos,omitempty" xml:"CertInfos,omitempty" type:"Struct"`
	RequestId  *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32                                            `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDcdnHttpsDomainListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnHttpsDomainListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnHttpsDomainListResponseBody) SetCertInfos(v *DescribeDcdnHttpsDomainListResponseBodyCertInfos) *DescribeDcdnHttpsDomainListResponseBody {
	s.CertInfos = v
	return s
}

func (s *DescribeDcdnHttpsDomainListResponseBody) SetRequestId(v string) *DescribeDcdnHttpsDomainListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDcdnHttpsDomainListResponseBody) SetTotalCount(v int32) *DescribeDcdnHttpsDomainListResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeDcdnHttpsDomainListResponseBodyCertInfos struct {
	CertInfo []*DescribeDcdnHttpsDomainListResponseBodyCertInfosCertInfo `json:"CertInfo,omitempty" xml:"CertInfo,omitempty" type:"Repeated"`
}

func (s DescribeDcdnHttpsDomainListResponseBodyCertInfos) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnHttpsDomainListResponseBodyCertInfos) GoString() string {
	return s.String()
}

func (s *DescribeDcdnHttpsDomainListResponseBodyCertInfos) SetCertInfo(v []*DescribeDcdnHttpsDomainListResponseBodyCertInfosCertInfo) *DescribeDcdnHttpsDomainListResponseBodyCertInfos {
	s.CertInfo = v
	return s
}

type DescribeDcdnHttpsDomainListResponseBodyCertInfosCertInfo struct {
	CertCommonName *string `json:"CertCommonName,omitempty" xml:"CertCommonName,omitempty"`
	CertExpireTime *string `json:"CertExpireTime,omitempty" xml:"CertExpireTime,omitempty"`
	CertName       *string `json:"CertName,omitempty" xml:"CertName,omitempty"`
	CertStartTime  *string `json:"CertStartTime,omitempty" xml:"CertStartTime,omitempty"`
	CertStatus     *string `json:"CertStatus,omitempty" xml:"CertStatus,omitempty"`
	CertType       *string `json:"CertType,omitempty" xml:"CertType,omitempty"`
	CertUpdateTime *string `json:"CertUpdateTime,omitempty" xml:"CertUpdateTime,omitempty"`
	DomainName     *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
}

func (s DescribeDcdnHttpsDomainListResponseBodyCertInfosCertInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnHttpsDomainListResponseBodyCertInfosCertInfo) GoString() string {
	return s.String()
}

func (s *DescribeDcdnHttpsDomainListResponseBodyCertInfosCertInfo) SetCertCommonName(v string) *DescribeDcdnHttpsDomainListResponseBodyCertInfosCertInfo {
	s.CertCommonName = &v
	return s
}

func (s *DescribeDcdnHttpsDomainListResponseBodyCertInfosCertInfo) SetCertExpireTime(v string) *DescribeDcdnHttpsDomainListResponseBodyCertInfosCertInfo {
	s.CertExpireTime = &v
	return s
}

func (s *DescribeDcdnHttpsDomainListResponseBodyCertInfosCertInfo) SetCertName(v string) *DescribeDcdnHttpsDomainListResponseBodyCertInfosCertInfo {
	s.CertName = &v
	return s
}

func (s *DescribeDcdnHttpsDomainListResponseBodyCertInfosCertInfo) SetCertStartTime(v string) *DescribeDcdnHttpsDomainListResponseBodyCertInfosCertInfo {
	s.CertStartTime = &v
	return s
}

func (s *DescribeDcdnHttpsDomainListResponseBodyCertInfosCertInfo) SetCertStatus(v string) *DescribeDcdnHttpsDomainListResponseBodyCertInfosCertInfo {
	s.CertStatus = &v
	return s
}

func (s *DescribeDcdnHttpsDomainListResponseBodyCertInfosCertInfo) SetCertType(v string) *DescribeDcdnHttpsDomainListResponseBodyCertInfosCertInfo {
	s.CertType = &v
	return s
}

func (s *DescribeDcdnHttpsDomainListResponseBodyCertInfosCertInfo) SetCertUpdateTime(v string) *DescribeDcdnHttpsDomainListResponseBodyCertInfosCertInfo {
	s.CertUpdateTime = &v
	return s
}

func (s *DescribeDcdnHttpsDomainListResponseBodyCertInfosCertInfo) SetDomainName(v string) *DescribeDcdnHttpsDomainListResponseBodyCertInfosCertInfo {
	s.DomainName = &v
	return s
}

type DescribeDcdnHttpsDomainListResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDcdnHttpsDomainListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDcdnHttpsDomainListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnHttpsDomainListResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnHttpsDomainListResponse) SetHeaders(v map[string]*string) *DescribeDcdnHttpsDomainListResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnHttpsDomainListResponse) SetBody(v *DescribeDcdnHttpsDomainListResponseBody) *DescribeDcdnHttpsDomainListResponse {
	s.Body = v
	return s
}

type DescribeDcdnIpInfoRequest struct {
	IP            *string `json:"IP,omitempty" xml:"IP,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeDcdnIpInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnIpInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeDcdnIpInfoRequest) SetIP(v string) *DescribeDcdnIpInfoRequest {
	s.IP = &v
	return s
}

func (s *DescribeDcdnIpInfoRequest) SetOwnerId(v int64) *DescribeDcdnIpInfoRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDcdnIpInfoRequest) SetSecurityToken(v string) *DescribeDcdnIpInfoRequest {
	s.SecurityToken = &v
	return s
}

type DescribeDcdnIpInfoResponseBody struct {
	DcdnIp      *string `json:"DcdnIp,omitempty" xml:"DcdnIp,omitempty"`
	ISP         *string `json:"ISP,omitempty" xml:"ISP,omitempty"`
	IspEname    *string `json:"IspEname,omitempty" xml:"IspEname,omitempty"`
	Region      *string `json:"Region,omitempty" xml:"Region,omitempty"`
	RegionEname *string `json:"RegionEname,omitempty" xml:"RegionEname,omitempty"`
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDcdnIpInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnIpInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnIpInfoResponseBody) SetDcdnIp(v string) *DescribeDcdnIpInfoResponseBody {
	s.DcdnIp = &v
	return s
}

func (s *DescribeDcdnIpInfoResponseBody) SetISP(v string) *DescribeDcdnIpInfoResponseBody {
	s.ISP = &v
	return s
}

func (s *DescribeDcdnIpInfoResponseBody) SetIspEname(v string) *DescribeDcdnIpInfoResponseBody {
	s.IspEname = &v
	return s
}

func (s *DescribeDcdnIpInfoResponseBody) SetRegion(v string) *DescribeDcdnIpInfoResponseBody {
	s.Region = &v
	return s
}

func (s *DescribeDcdnIpInfoResponseBody) SetRegionEname(v string) *DescribeDcdnIpInfoResponseBody {
	s.RegionEname = &v
	return s
}

func (s *DescribeDcdnIpInfoResponseBody) SetRequestId(v string) *DescribeDcdnIpInfoResponseBody {
	s.RequestId = &v
	return s
}

type DescribeDcdnIpInfoResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDcdnIpInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDcdnIpInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnIpInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnIpInfoResponse) SetHeaders(v map[string]*string) *DescribeDcdnIpInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnIpInfoResponse) SetBody(v *DescribeDcdnIpInfoResponseBody) *DescribeDcdnIpInfoResponse {
	s.Body = v
	return s
}

type DescribeDcdnIpaDomainConfigsRequest struct {
	DomainName    *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	FunctionNames *string `json:"FunctionNames,omitempty" xml:"FunctionNames,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeDcdnIpaDomainConfigsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnIpaDomainConfigsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDcdnIpaDomainConfigsRequest) SetDomainName(v string) *DescribeDcdnIpaDomainConfigsRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDcdnIpaDomainConfigsRequest) SetFunctionNames(v string) *DescribeDcdnIpaDomainConfigsRequest {
	s.FunctionNames = &v
	return s
}

func (s *DescribeDcdnIpaDomainConfigsRequest) SetOwnerId(v int64) *DescribeDcdnIpaDomainConfigsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDcdnIpaDomainConfigsRequest) SetSecurityToken(v string) *DescribeDcdnIpaDomainConfigsRequest {
	s.SecurityToken = &v
	return s
}

type DescribeDcdnIpaDomainConfigsResponseBody struct {
	DomainConfigs *DescribeDcdnIpaDomainConfigsResponseBodyDomainConfigs `json:"DomainConfigs,omitempty" xml:"DomainConfigs,omitempty" type:"Struct"`
	RequestId     *string                                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDcdnIpaDomainConfigsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnIpaDomainConfigsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnIpaDomainConfigsResponseBody) SetDomainConfigs(v *DescribeDcdnIpaDomainConfigsResponseBodyDomainConfigs) *DescribeDcdnIpaDomainConfigsResponseBody {
	s.DomainConfigs = v
	return s
}

func (s *DescribeDcdnIpaDomainConfigsResponseBody) SetRequestId(v string) *DescribeDcdnIpaDomainConfigsResponseBody {
	s.RequestId = &v
	return s
}

type DescribeDcdnIpaDomainConfigsResponseBodyDomainConfigs struct {
	DomainConfig []*DescribeDcdnIpaDomainConfigsResponseBodyDomainConfigsDomainConfig `json:"DomainConfig,omitempty" xml:"DomainConfig,omitempty" type:"Repeated"`
}

func (s DescribeDcdnIpaDomainConfigsResponseBodyDomainConfigs) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnIpaDomainConfigsResponseBodyDomainConfigs) GoString() string {
	return s.String()
}

func (s *DescribeDcdnIpaDomainConfigsResponseBodyDomainConfigs) SetDomainConfig(v []*DescribeDcdnIpaDomainConfigsResponseBodyDomainConfigsDomainConfig) *DescribeDcdnIpaDomainConfigsResponseBodyDomainConfigs {
	s.DomainConfig = v
	return s
}

type DescribeDcdnIpaDomainConfigsResponseBodyDomainConfigsDomainConfig struct {
	ConfigId     *string                                                                        `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	FunctionArgs *DescribeDcdnIpaDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgs `json:"FunctionArgs,omitempty" xml:"FunctionArgs,omitempty" type:"Struct"`
	FunctionName *string                                                                        `json:"FunctionName,omitempty" xml:"FunctionName,omitempty"`
	Status       *string                                                                        `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeDcdnIpaDomainConfigsResponseBodyDomainConfigsDomainConfig) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnIpaDomainConfigsResponseBodyDomainConfigsDomainConfig) GoString() string {
	return s.String()
}

func (s *DescribeDcdnIpaDomainConfigsResponseBodyDomainConfigsDomainConfig) SetConfigId(v string) *DescribeDcdnIpaDomainConfigsResponseBodyDomainConfigsDomainConfig {
	s.ConfigId = &v
	return s
}

func (s *DescribeDcdnIpaDomainConfigsResponseBodyDomainConfigsDomainConfig) SetFunctionArgs(v *DescribeDcdnIpaDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgs) *DescribeDcdnIpaDomainConfigsResponseBodyDomainConfigsDomainConfig {
	s.FunctionArgs = v
	return s
}

func (s *DescribeDcdnIpaDomainConfigsResponseBodyDomainConfigsDomainConfig) SetFunctionName(v string) *DescribeDcdnIpaDomainConfigsResponseBodyDomainConfigsDomainConfig {
	s.FunctionName = &v
	return s
}

func (s *DescribeDcdnIpaDomainConfigsResponseBodyDomainConfigsDomainConfig) SetStatus(v string) *DescribeDcdnIpaDomainConfigsResponseBodyDomainConfigsDomainConfig {
	s.Status = &v
	return s
}

type DescribeDcdnIpaDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgs struct {
	FunctionArg []*DescribeDcdnIpaDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgsFunctionArg `json:"FunctionArg,omitempty" xml:"FunctionArg,omitempty" type:"Repeated"`
}

func (s DescribeDcdnIpaDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgs) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnIpaDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgs) GoString() string {
	return s.String()
}

func (s *DescribeDcdnIpaDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgs) SetFunctionArg(v []*DescribeDcdnIpaDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgsFunctionArg) *DescribeDcdnIpaDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgs {
	s.FunctionArg = v
	return s
}

type DescribeDcdnIpaDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgsFunctionArg struct {
	ArgName  *string `json:"ArgName,omitempty" xml:"ArgName,omitempty"`
	ArgValue *string `json:"ArgValue,omitempty" xml:"ArgValue,omitempty"`
}

func (s DescribeDcdnIpaDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgsFunctionArg) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnIpaDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgsFunctionArg) GoString() string {
	return s.String()
}

func (s *DescribeDcdnIpaDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgsFunctionArg) SetArgName(v string) *DescribeDcdnIpaDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgsFunctionArg {
	s.ArgName = &v
	return s
}

func (s *DescribeDcdnIpaDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgsFunctionArg) SetArgValue(v string) *DescribeDcdnIpaDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgsFunctionArg {
	s.ArgValue = &v
	return s
}

type DescribeDcdnIpaDomainConfigsResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDcdnIpaDomainConfigsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDcdnIpaDomainConfigsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnIpaDomainConfigsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnIpaDomainConfigsResponse) SetHeaders(v map[string]*string) *DescribeDcdnIpaDomainConfigsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnIpaDomainConfigsResponse) SetBody(v *DescribeDcdnIpaDomainConfigsResponseBody) *DescribeDcdnIpaDomainConfigsResponse {
	s.Body = v
	return s
}

type DescribeDcdnIpaDomainDetailRequest struct {
	DomainName    *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeDcdnIpaDomainDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnIpaDomainDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribeDcdnIpaDomainDetailRequest) SetDomainName(v string) *DescribeDcdnIpaDomainDetailRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDcdnIpaDomainDetailRequest) SetOwnerId(v int64) *DescribeDcdnIpaDomainDetailRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDcdnIpaDomainDetailRequest) SetSecurityToken(v string) *DescribeDcdnIpaDomainDetailRequest {
	s.SecurityToken = &v
	return s
}

type DescribeDcdnIpaDomainDetailResponseBody struct {
	DomainDetail *DescribeDcdnIpaDomainDetailResponseBodyDomainDetail `json:"DomainDetail,omitempty" xml:"DomainDetail,omitempty" type:"Struct"`
	RequestId    *string                                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDcdnIpaDomainDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnIpaDomainDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnIpaDomainDetailResponseBody) SetDomainDetail(v *DescribeDcdnIpaDomainDetailResponseBodyDomainDetail) *DescribeDcdnIpaDomainDetailResponseBody {
	s.DomainDetail = v
	return s
}

func (s *DescribeDcdnIpaDomainDetailResponseBody) SetRequestId(v string) *DescribeDcdnIpaDomainDetailResponseBody {
	s.RequestId = &v
	return s
}

type DescribeDcdnIpaDomainDetailResponseBodyDomainDetail struct {
	CertName        *string                                                     `json:"CertName,omitempty" xml:"CertName,omitempty"`
	Cname           *string                                                     `json:"Cname,omitempty" xml:"Cname,omitempty"`
	Description     *string                                                     `json:"Description,omitempty" xml:"Description,omitempty"`
	DomainName      *string                                                     `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	DomainStatus    *string                                                     `json:"DomainStatus,omitempty" xml:"DomainStatus,omitempty"`
	GmtCreated      *string                                                     `json:"GmtCreated,omitempty" xml:"GmtCreated,omitempty"`
	GmtModified     *string                                                     `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	ResourceGroupId *string                                                     `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	SSLProtocol     *string                                                     `json:"SSLProtocol,omitempty" xml:"SSLProtocol,omitempty"`
	SSLPub          *string                                                     `json:"SSLPub,omitempty" xml:"SSLPub,omitempty"`
	Scope           *string                                                     `json:"Scope,omitempty" xml:"Scope,omitempty"`
	Sources         *DescribeDcdnIpaDomainDetailResponseBodyDomainDetailSources `json:"Sources,omitempty" xml:"Sources,omitempty" type:"Struct"`
}

func (s DescribeDcdnIpaDomainDetailResponseBodyDomainDetail) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnIpaDomainDetailResponseBodyDomainDetail) GoString() string {
	return s.String()
}

func (s *DescribeDcdnIpaDomainDetailResponseBodyDomainDetail) SetCertName(v string) *DescribeDcdnIpaDomainDetailResponseBodyDomainDetail {
	s.CertName = &v
	return s
}

func (s *DescribeDcdnIpaDomainDetailResponseBodyDomainDetail) SetCname(v string) *DescribeDcdnIpaDomainDetailResponseBodyDomainDetail {
	s.Cname = &v
	return s
}

func (s *DescribeDcdnIpaDomainDetailResponseBodyDomainDetail) SetDescription(v string) *DescribeDcdnIpaDomainDetailResponseBodyDomainDetail {
	s.Description = &v
	return s
}

func (s *DescribeDcdnIpaDomainDetailResponseBodyDomainDetail) SetDomainName(v string) *DescribeDcdnIpaDomainDetailResponseBodyDomainDetail {
	s.DomainName = &v
	return s
}

func (s *DescribeDcdnIpaDomainDetailResponseBodyDomainDetail) SetDomainStatus(v string) *DescribeDcdnIpaDomainDetailResponseBodyDomainDetail {
	s.DomainStatus = &v
	return s
}

func (s *DescribeDcdnIpaDomainDetailResponseBodyDomainDetail) SetGmtCreated(v string) *DescribeDcdnIpaDomainDetailResponseBodyDomainDetail {
	s.GmtCreated = &v
	return s
}

func (s *DescribeDcdnIpaDomainDetailResponseBodyDomainDetail) SetGmtModified(v string) *DescribeDcdnIpaDomainDetailResponseBodyDomainDetail {
	s.GmtModified = &v
	return s
}

func (s *DescribeDcdnIpaDomainDetailResponseBodyDomainDetail) SetResourceGroupId(v string) *DescribeDcdnIpaDomainDetailResponseBodyDomainDetail {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDcdnIpaDomainDetailResponseBodyDomainDetail) SetSSLProtocol(v string) *DescribeDcdnIpaDomainDetailResponseBodyDomainDetail {
	s.SSLProtocol = &v
	return s
}

func (s *DescribeDcdnIpaDomainDetailResponseBodyDomainDetail) SetSSLPub(v string) *DescribeDcdnIpaDomainDetailResponseBodyDomainDetail {
	s.SSLPub = &v
	return s
}

func (s *DescribeDcdnIpaDomainDetailResponseBodyDomainDetail) SetScope(v string) *DescribeDcdnIpaDomainDetailResponseBodyDomainDetail {
	s.Scope = &v
	return s
}

func (s *DescribeDcdnIpaDomainDetailResponseBodyDomainDetail) SetSources(v *DescribeDcdnIpaDomainDetailResponseBodyDomainDetailSources) *DescribeDcdnIpaDomainDetailResponseBodyDomainDetail {
	s.Sources = v
	return s
}

type DescribeDcdnIpaDomainDetailResponseBodyDomainDetailSources struct {
	Source []*DescribeDcdnIpaDomainDetailResponseBodyDomainDetailSourcesSource `json:"Source,omitempty" xml:"Source,omitempty" type:"Repeated"`
}

func (s DescribeDcdnIpaDomainDetailResponseBodyDomainDetailSources) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnIpaDomainDetailResponseBodyDomainDetailSources) GoString() string {
	return s.String()
}

func (s *DescribeDcdnIpaDomainDetailResponseBodyDomainDetailSources) SetSource(v []*DescribeDcdnIpaDomainDetailResponseBodyDomainDetailSourcesSource) *DescribeDcdnIpaDomainDetailResponseBodyDomainDetailSources {
	s.Source = v
	return s
}

type DescribeDcdnIpaDomainDetailResponseBodyDomainDetailSourcesSource struct {
	Content  *string `json:"Content,omitempty" xml:"Content,omitempty"`
	Enabled  *string `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	Port     *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
	Priority *string `json:"Priority,omitempty" xml:"Priority,omitempty"`
	Type     *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Weight   *string `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s DescribeDcdnIpaDomainDetailResponseBodyDomainDetailSourcesSource) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnIpaDomainDetailResponseBodyDomainDetailSourcesSource) GoString() string {
	return s.String()
}

func (s *DescribeDcdnIpaDomainDetailResponseBodyDomainDetailSourcesSource) SetContent(v string) *DescribeDcdnIpaDomainDetailResponseBodyDomainDetailSourcesSource {
	s.Content = &v
	return s
}

func (s *DescribeDcdnIpaDomainDetailResponseBodyDomainDetailSourcesSource) SetEnabled(v string) *DescribeDcdnIpaDomainDetailResponseBodyDomainDetailSourcesSource {
	s.Enabled = &v
	return s
}

func (s *DescribeDcdnIpaDomainDetailResponseBodyDomainDetailSourcesSource) SetPort(v int32) *DescribeDcdnIpaDomainDetailResponseBodyDomainDetailSourcesSource {
	s.Port = &v
	return s
}

func (s *DescribeDcdnIpaDomainDetailResponseBodyDomainDetailSourcesSource) SetPriority(v string) *DescribeDcdnIpaDomainDetailResponseBodyDomainDetailSourcesSource {
	s.Priority = &v
	return s
}

func (s *DescribeDcdnIpaDomainDetailResponseBodyDomainDetailSourcesSource) SetType(v string) *DescribeDcdnIpaDomainDetailResponseBodyDomainDetailSourcesSource {
	s.Type = &v
	return s
}

func (s *DescribeDcdnIpaDomainDetailResponseBodyDomainDetailSourcesSource) SetWeight(v string) *DescribeDcdnIpaDomainDetailResponseBodyDomainDetailSourcesSource {
	s.Weight = &v
	return s
}

type DescribeDcdnIpaDomainDetailResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDcdnIpaDomainDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDcdnIpaDomainDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnIpaDomainDetailResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnIpaDomainDetailResponse) SetHeaders(v map[string]*string) *DescribeDcdnIpaDomainDetailResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnIpaDomainDetailResponse) SetBody(v *DescribeDcdnIpaDomainDetailResponseBody) *DescribeDcdnIpaDomainDetailResponse {
	s.Body = v
	return s
}

type DescribeDcdnIpaServiceRequest struct {
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeDcdnIpaServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnIpaServiceRequest) GoString() string {
	return s.String()
}

func (s *DescribeDcdnIpaServiceRequest) SetOwnerId(v int64) *DescribeDcdnIpaServiceRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDcdnIpaServiceRequest) SetSecurityToken(v string) *DescribeDcdnIpaServiceRequest {
	s.SecurityToken = &v
	return s
}

type DescribeDcdnIpaServiceResponseBody struct {
	ChangingAffectTime *string                                           `json:"ChangingAffectTime,omitempty" xml:"ChangingAffectTime,omitempty"`
	ChangingChargeType *string                                           `json:"ChangingChargeType,omitempty" xml:"ChangingChargeType,omitempty"`
	InstanceId         *string                                           `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InternetChargeType *string                                           `json:"InternetChargeType,omitempty" xml:"InternetChargeType,omitempty"`
	OpeningTime        *string                                           `json:"OpeningTime,omitempty" xml:"OpeningTime,omitempty"`
	OperationLocks     *DescribeDcdnIpaServiceResponseBodyOperationLocks `json:"OperationLocks,omitempty" xml:"OperationLocks,omitempty" type:"Struct"`
	RequestId          *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDcdnIpaServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnIpaServiceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnIpaServiceResponseBody) SetChangingAffectTime(v string) *DescribeDcdnIpaServiceResponseBody {
	s.ChangingAffectTime = &v
	return s
}

func (s *DescribeDcdnIpaServiceResponseBody) SetChangingChargeType(v string) *DescribeDcdnIpaServiceResponseBody {
	s.ChangingChargeType = &v
	return s
}

func (s *DescribeDcdnIpaServiceResponseBody) SetInstanceId(v string) *DescribeDcdnIpaServiceResponseBody {
	s.InstanceId = &v
	return s
}

func (s *DescribeDcdnIpaServiceResponseBody) SetInternetChargeType(v string) *DescribeDcdnIpaServiceResponseBody {
	s.InternetChargeType = &v
	return s
}

func (s *DescribeDcdnIpaServiceResponseBody) SetOpeningTime(v string) *DescribeDcdnIpaServiceResponseBody {
	s.OpeningTime = &v
	return s
}

func (s *DescribeDcdnIpaServiceResponseBody) SetOperationLocks(v *DescribeDcdnIpaServiceResponseBodyOperationLocks) *DescribeDcdnIpaServiceResponseBody {
	s.OperationLocks = v
	return s
}

func (s *DescribeDcdnIpaServiceResponseBody) SetRequestId(v string) *DescribeDcdnIpaServiceResponseBody {
	s.RequestId = &v
	return s
}

type DescribeDcdnIpaServiceResponseBodyOperationLocks struct {
	LockReason []*DescribeDcdnIpaServiceResponseBodyOperationLocksLockReason `json:"LockReason,omitempty" xml:"LockReason,omitempty" type:"Repeated"`
}

func (s DescribeDcdnIpaServiceResponseBodyOperationLocks) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnIpaServiceResponseBodyOperationLocks) GoString() string {
	return s.String()
}

func (s *DescribeDcdnIpaServiceResponseBodyOperationLocks) SetLockReason(v []*DescribeDcdnIpaServiceResponseBodyOperationLocksLockReason) *DescribeDcdnIpaServiceResponseBodyOperationLocks {
	s.LockReason = v
	return s
}

type DescribeDcdnIpaServiceResponseBodyOperationLocksLockReason struct {
	LockReason *string `json:"LockReason,omitempty" xml:"LockReason,omitempty"`
}

func (s DescribeDcdnIpaServiceResponseBodyOperationLocksLockReason) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnIpaServiceResponseBodyOperationLocksLockReason) GoString() string {
	return s.String()
}

func (s *DescribeDcdnIpaServiceResponseBodyOperationLocksLockReason) SetLockReason(v string) *DescribeDcdnIpaServiceResponseBodyOperationLocksLockReason {
	s.LockReason = &v
	return s
}

type DescribeDcdnIpaServiceResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDcdnIpaServiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDcdnIpaServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnIpaServiceResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnIpaServiceResponse) SetHeaders(v map[string]*string) *DescribeDcdnIpaServiceResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnIpaServiceResponse) SetBody(v *DescribeDcdnIpaServiceResponseBody) *DescribeDcdnIpaServiceResponse {
	s.Body = v
	return s
}

type DescribeDcdnIpaUserDomainsRequest struct {
	CheckDomainShow  *bool                                   `json:"CheckDomainShow,omitempty" xml:"CheckDomainShow,omitempty"`
	DomainName       *string                                 `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	DomainSearchType *string                                 `json:"DomainSearchType,omitempty" xml:"DomainSearchType,omitempty"`
	DomainStatus     *string                                 `json:"DomainStatus,omitempty" xml:"DomainStatus,omitempty"`
	FuncFilter       *string                                 `json:"FuncFilter,omitempty" xml:"FuncFilter,omitempty"`
	FuncId           *string                                 `json:"FuncId,omitempty" xml:"FuncId,omitempty"`
	OwnerId          *int64                                  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNumber       *int32                                  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize         *int32                                  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceGroupId  *string                                 `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	SecurityToken    *string                                 `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	Tag              []*DescribeDcdnIpaUserDomainsRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeDcdnIpaUserDomainsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnIpaUserDomainsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDcdnIpaUserDomainsRequest) SetCheckDomainShow(v bool) *DescribeDcdnIpaUserDomainsRequest {
	s.CheckDomainShow = &v
	return s
}

func (s *DescribeDcdnIpaUserDomainsRequest) SetDomainName(v string) *DescribeDcdnIpaUserDomainsRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDcdnIpaUserDomainsRequest) SetDomainSearchType(v string) *DescribeDcdnIpaUserDomainsRequest {
	s.DomainSearchType = &v
	return s
}

func (s *DescribeDcdnIpaUserDomainsRequest) SetDomainStatus(v string) *DescribeDcdnIpaUserDomainsRequest {
	s.DomainStatus = &v
	return s
}

func (s *DescribeDcdnIpaUserDomainsRequest) SetFuncFilter(v string) *DescribeDcdnIpaUserDomainsRequest {
	s.FuncFilter = &v
	return s
}

func (s *DescribeDcdnIpaUserDomainsRequest) SetFuncId(v string) *DescribeDcdnIpaUserDomainsRequest {
	s.FuncId = &v
	return s
}

func (s *DescribeDcdnIpaUserDomainsRequest) SetOwnerId(v int64) *DescribeDcdnIpaUserDomainsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDcdnIpaUserDomainsRequest) SetPageNumber(v int32) *DescribeDcdnIpaUserDomainsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDcdnIpaUserDomainsRequest) SetPageSize(v int32) *DescribeDcdnIpaUserDomainsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDcdnIpaUserDomainsRequest) SetResourceGroupId(v string) *DescribeDcdnIpaUserDomainsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDcdnIpaUserDomainsRequest) SetSecurityToken(v string) *DescribeDcdnIpaUserDomainsRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeDcdnIpaUserDomainsRequest) SetTag(v []*DescribeDcdnIpaUserDomainsRequestTag) *DescribeDcdnIpaUserDomainsRequest {
	s.Tag = v
	return s
}

type DescribeDcdnIpaUserDomainsRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDcdnIpaUserDomainsRequestTag) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnIpaUserDomainsRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeDcdnIpaUserDomainsRequestTag) SetKey(v string) *DescribeDcdnIpaUserDomainsRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeDcdnIpaUserDomainsRequestTag) SetValue(v string) *DescribeDcdnIpaUserDomainsRequestTag {
	s.Value = &v
	return s
}

type DescribeDcdnIpaUserDomainsResponseBody struct {
	Domains    *DescribeDcdnIpaUserDomainsResponseBodyDomains `json:"Domains,omitempty" xml:"Domains,omitempty" type:"Struct"`
	PageNumber *int64                                         `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int64                                         `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int64                                         `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDcdnIpaUserDomainsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnIpaUserDomainsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnIpaUserDomainsResponseBody) SetDomains(v *DescribeDcdnIpaUserDomainsResponseBodyDomains) *DescribeDcdnIpaUserDomainsResponseBody {
	s.Domains = v
	return s
}

func (s *DescribeDcdnIpaUserDomainsResponseBody) SetPageNumber(v int64) *DescribeDcdnIpaUserDomainsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeDcdnIpaUserDomainsResponseBody) SetPageSize(v int64) *DescribeDcdnIpaUserDomainsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeDcdnIpaUserDomainsResponseBody) SetRequestId(v string) *DescribeDcdnIpaUserDomainsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDcdnIpaUserDomainsResponseBody) SetTotalCount(v int64) *DescribeDcdnIpaUserDomainsResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeDcdnIpaUserDomainsResponseBodyDomains struct {
	PageData []*DescribeDcdnIpaUserDomainsResponseBodyDomainsPageData `json:"PageData,omitempty" xml:"PageData,omitempty" type:"Repeated"`
}

func (s DescribeDcdnIpaUserDomainsResponseBodyDomains) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnIpaUserDomainsResponseBodyDomains) GoString() string {
	return s.String()
}

func (s *DescribeDcdnIpaUserDomainsResponseBodyDomains) SetPageData(v []*DescribeDcdnIpaUserDomainsResponseBodyDomainsPageData) *DescribeDcdnIpaUserDomainsResponseBodyDomains {
	s.PageData = v
	return s
}

type DescribeDcdnIpaUserDomainsResponseBodyDomainsPageData struct {
	Cname           *string                                                       `json:"Cname,omitempty" xml:"Cname,omitempty"`
	Description     *string                                                       `json:"Description,omitempty" xml:"Description,omitempty"`
	DomainName      *string                                                       `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	DomainStatus    *string                                                       `json:"DomainStatus,omitempty" xml:"DomainStatus,omitempty"`
	GmtCreated      *string                                                       `json:"GmtCreated,omitempty" xml:"GmtCreated,omitempty"`
	GmtModified     *string                                                       `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	ResourceGroupId *string                                                       `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	SSLProtocol     *string                                                       `json:"SSLProtocol,omitempty" xml:"SSLProtocol,omitempty"`
	Sandbox         *string                                                       `json:"Sandbox,omitempty" xml:"Sandbox,omitempty"`
	Sources         *DescribeDcdnIpaUserDomainsResponseBodyDomainsPageDataSources `json:"Sources,omitempty" xml:"Sources,omitempty" type:"Struct"`
}

func (s DescribeDcdnIpaUserDomainsResponseBodyDomainsPageData) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnIpaUserDomainsResponseBodyDomainsPageData) GoString() string {
	return s.String()
}

func (s *DescribeDcdnIpaUserDomainsResponseBodyDomainsPageData) SetCname(v string) *DescribeDcdnIpaUserDomainsResponseBodyDomainsPageData {
	s.Cname = &v
	return s
}

func (s *DescribeDcdnIpaUserDomainsResponseBodyDomainsPageData) SetDescription(v string) *DescribeDcdnIpaUserDomainsResponseBodyDomainsPageData {
	s.Description = &v
	return s
}

func (s *DescribeDcdnIpaUserDomainsResponseBodyDomainsPageData) SetDomainName(v string) *DescribeDcdnIpaUserDomainsResponseBodyDomainsPageData {
	s.DomainName = &v
	return s
}

func (s *DescribeDcdnIpaUserDomainsResponseBodyDomainsPageData) SetDomainStatus(v string) *DescribeDcdnIpaUserDomainsResponseBodyDomainsPageData {
	s.DomainStatus = &v
	return s
}

func (s *DescribeDcdnIpaUserDomainsResponseBodyDomainsPageData) SetGmtCreated(v string) *DescribeDcdnIpaUserDomainsResponseBodyDomainsPageData {
	s.GmtCreated = &v
	return s
}

func (s *DescribeDcdnIpaUserDomainsResponseBodyDomainsPageData) SetGmtModified(v string) *DescribeDcdnIpaUserDomainsResponseBodyDomainsPageData {
	s.GmtModified = &v
	return s
}

func (s *DescribeDcdnIpaUserDomainsResponseBodyDomainsPageData) SetResourceGroupId(v string) *DescribeDcdnIpaUserDomainsResponseBodyDomainsPageData {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDcdnIpaUserDomainsResponseBodyDomainsPageData) SetSSLProtocol(v string) *DescribeDcdnIpaUserDomainsResponseBodyDomainsPageData {
	s.SSLProtocol = &v
	return s
}

func (s *DescribeDcdnIpaUserDomainsResponseBodyDomainsPageData) SetSandbox(v string) *DescribeDcdnIpaUserDomainsResponseBodyDomainsPageData {
	s.Sandbox = &v
	return s
}

func (s *DescribeDcdnIpaUserDomainsResponseBodyDomainsPageData) SetSources(v *DescribeDcdnIpaUserDomainsResponseBodyDomainsPageDataSources) *DescribeDcdnIpaUserDomainsResponseBodyDomainsPageData {
	s.Sources = v
	return s
}

type DescribeDcdnIpaUserDomainsResponseBodyDomainsPageDataSources struct {
	Source []*DescribeDcdnIpaUserDomainsResponseBodyDomainsPageDataSourcesSource `json:"Source,omitempty" xml:"Source,omitempty" type:"Repeated"`
}

func (s DescribeDcdnIpaUserDomainsResponseBodyDomainsPageDataSources) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnIpaUserDomainsResponseBodyDomainsPageDataSources) GoString() string {
	return s.String()
}

func (s *DescribeDcdnIpaUserDomainsResponseBodyDomainsPageDataSources) SetSource(v []*DescribeDcdnIpaUserDomainsResponseBodyDomainsPageDataSourcesSource) *DescribeDcdnIpaUserDomainsResponseBodyDomainsPageDataSources {
	s.Source = v
	return s
}

type DescribeDcdnIpaUserDomainsResponseBodyDomainsPageDataSourcesSource struct {
	Content  *string `json:"Content,omitempty" xml:"Content,omitempty"`
	Port     *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
	Priority *string `json:"Priority,omitempty" xml:"Priority,omitempty"`
	Type     *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Weight   *string `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s DescribeDcdnIpaUserDomainsResponseBodyDomainsPageDataSourcesSource) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnIpaUserDomainsResponseBodyDomainsPageDataSourcesSource) GoString() string {
	return s.String()
}

func (s *DescribeDcdnIpaUserDomainsResponseBodyDomainsPageDataSourcesSource) SetContent(v string) *DescribeDcdnIpaUserDomainsResponseBodyDomainsPageDataSourcesSource {
	s.Content = &v
	return s
}

func (s *DescribeDcdnIpaUserDomainsResponseBodyDomainsPageDataSourcesSource) SetPort(v int32) *DescribeDcdnIpaUserDomainsResponseBodyDomainsPageDataSourcesSource {
	s.Port = &v
	return s
}

func (s *DescribeDcdnIpaUserDomainsResponseBodyDomainsPageDataSourcesSource) SetPriority(v string) *DescribeDcdnIpaUserDomainsResponseBodyDomainsPageDataSourcesSource {
	s.Priority = &v
	return s
}

func (s *DescribeDcdnIpaUserDomainsResponseBodyDomainsPageDataSourcesSource) SetType(v string) *DescribeDcdnIpaUserDomainsResponseBodyDomainsPageDataSourcesSource {
	s.Type = &v
	return s
}

func (s *DescribeDcdnIpaUserDomainsResponseBodyDomainsPageDataSourcesSource) SetWeight(v string) *DescribeDcdnIpaUserDomainsResponseBodyDomainsPageDataSourcesSource {
	s.Weight = &v
	return s
}

type DescribeDcdnIpaUserDomainsResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDcdnIpaUserDomainsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDcdnIpaUserDomainsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnIpaUserDomainsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnIpaUserDomainsResponse) SetHeaders(v map[string]*string) *DescribeDcdnIpaUserDomainsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnIpaUserDomainsResponse) SetBody(v *DescribeDcdnIpaUserDomainsResponseBody) *DescribeDcdnIpaUserDomainsResponse {
	s.Body = v
	return s
}

type DescribeDcdnRealTimeDeliveryFieldRequest struct {
	BusinessType *string `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s DescribeDcdnRealTimeDeliveryFieldRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnRealTimeDeliveryFieldRequest) GoString() string {
	return s.String()
}

func (s *DescribeDcdnRealTimeDeliveryFieldRequest) SetBusinessType(v string) *DescribeDcdnRealTimeDeliveryFieldRequest {
	s.BusinessType = &v
	return s
}

func (s *DescribeDcdnRealTimeDeliveryFieldRequest) SetOwnerId(v int64) *DescribeDcdnRealTimeDeliveryFieldRequest {
	s.OwnerId = &v
	return s
}

type DescribeDcdnRealTimeDeliveryFieldResponseBody struct {
	Content   *DescribeDcdnRealTimeDeliveryFieldResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	RequestId *string                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDcdnRealTimeDeliveryFieldResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnRealTimeDeliveryFieldResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnRealTimeDeliveryFieldResponseBody) SetContent(v *DescribeDcdnRealTimeDeliveryFieldResponseBodyContent) *DescribeDcdnRealTimeDeliveryFieldResponseBody {
	s.Content = v
	return s
}

func (s *DescribeDcdnRealTimeDeliveryFieldResponseBody) SetRequestId(v string) *DescribeDcdnRealTimeDeliveryFieldResponseBody {
	s.RequestId = &v
	return s
}

type DescribeDcdnRealTimeDeliveryFieldResponseBodyContent struct {
	Fields []*DescribeDcdnRealTimeDeliveryFieldResponseBodyContentFields `json:"Fields,omitempty" xml:"Fields,omitempty" type:"Repeated"`
}

func (s DescribeDcdnRealTimeDeliveryFieldResponseBodyContent) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnRealTimeDeliveryFieldResponseBodyContent) GoString() string {
	return s.String()
}

func (s *DescribeDcdnRealTimeDeliveryFieldResponseBodyContent) SetFields(v []*DescribeDcdnRealTimeDeliveryFieldResponseBodyContentFields) *DescribeDcdnRealTimeDeliveryFieldResponseBodyContent {
	s.Fields = v
	return s
}

type DescribeDcdnRealTimeDeliveryFieldResponseBodyContentFields struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	FieldName   *string `json:"FieldName,omitempty" xml:"FieldName,omitempty"`
}

func (s DescribeDcdnRealTimeDeliveryFieldResponseBodyContentFields) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnRealTimeDeliveryFieldResponseBodyContentFields) GoString() string {
	return s.String()
}

func (s *DescribeDcdnRealTimeDeliveryFieldResponseBodyContentFields) SetDescription(v string) *DescribeDcdnRealTimeDeliveryFieldResponseBodyContentFields {
	s.Description = &v
	return s
}

func (s *DescribeDcdnRealTimeDeliveryFieldResponseBodyContentFields) SetFieldName(v string) *DescribeDcdnRealTimeDeliveryFieldResponseBodyContentFields {
	s.FieldName = &v
	return s
}

type DescribeDcdnRealTimeDeliveryFieldResponse struct {
	Headers map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDcdnRealTimeDeliveryFieldResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDcdnRealTimeDeliveryFieldResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnRealTimeDeliveryFieldResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnRealTimeDeliveryFieldResponse) SetHeaders(v map[string]*string) *DescribeDcdnRealTimeDeliveryFieldResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnRealTimeDeliveryFieldResponse) SetBody(v *DescribeDcdnRealTimeDeliveryFieldResponseBody) *DescribeDcdnRealTimeDeliveryFieldResponse {
	s.Body = v
	return s
}

type DescribeDcdnRefreshQuotaRequest struct {
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeDcdnRefreshQuotaRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnRefreshQuotaRequest) GoString() string {
	return s.String()
}

func (s *DescribeDcdnRefreshQuotaRequest) SetOwnerId(v int64) *DescribeDcdnRefreshQuotaRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDcdnRefreshQuotaRequest) SetSecurityToken(v string) *DescribeDcdnRefreshQuotaRequest {
	s.SecurityToken = &v
	return s
}

type DescribeDcdnRefreshQuotaResponseBody struct {
	BlockQuota    *string `json:"BlockQuota,omitempty" xml:"BlockQuota,omitempty"`
	BlockRemain   *string `json:"BlockRemain,omitempty" xml:"BlockRemain,omitempty"`
	DirQuota      *string `json:"DirQuota,omitempty" xml:"DirQuota,omitempty"`
	DirRemain     *string `json:"DirRemain,omitempty" xml:"DirRemain,omitempty"`
	PreloadQuota  *string `json:"PreloadQuota,omitempty" xml:"PreloadQuota,omitempty"`
	PreloadRemain *string `json:"PreloadRemain,omitempty" xml:"PreloadRemain,omitempty"`
	RegexQuota    *string `json:"RegexQuota,omitempty" xml:"RegexQuota,omitempty"`
	RegexRemain   *string `json:"RegexRemain,omitempty" xml:"RegexRemain,omitempty"`
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	UrlQuota      *string `json:"UrlQuota,omitempty" xml:"UrlQuota,omitempty"`
	UrlRemain     *string `json:"UrlRemain,omitempty" xml:"UrlRemain,omitempty"`
}

func (s DescribeDcdnRefreshQuotaResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnRefreshQuotaResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnRefreshQuotaResponseBody) SetBlockQuota(v string) *DescribeDcdnRefreshQuotaResponseBody {
	s.BlockQuota = &v
	return s
}

func (s *DescribeDcdnRefreshQuotaResponseBody) SetBlockRemain(v string) *DescribeDcdnRefreshQuotaResponseBody {
	s.BlockRemain = &v
	return s
}

func (s *DescribeDcdnRefreshQuotaResponseBody) SetDirQuota(v string) *DescribeDcdnRefreshQuotaResponseBody {
	s.DirQuota = &v
	return s
}

func (s *DescribeDcdnRefreshQuotaResponseBody) SetDirRemain(v string) *DescribeDcdnRefreshQuotaResponseBody {
	s.DirRemain = &v
	return s
}

func (s *DescribeDcdnRefreshQuotaResponseBody) SetPreloadQuota(v string) *DescribeDcdnRefreshQuotaResponseBody {
	s.PreloadQuota = &v
	return s
}

func (s *DescribeDcdnRefreshQuotaResponseBody) SetPreloadRemain(v string) *DescribeDcdnRefreshQuotaResponseBody {
	s.PreloadRemain = &v
	return s
}

func (s *DescribeDcdnRefreshQuotaResponseBody) SetRegexQuota(v string) *DescribeDcdnRefreshQuotaResponseBody {
	s.RegexQuota = &v
	return s
}

func (s *DescribeDcdnRefreshQuotaResponseBody) SetRegexRemain(v string) *DescribeDcdnRefreshQuotaResponseBody {
	s.RegexRemain = &v
	return s
}

func (s *DescribeDcdnRefreshQuotaResponseBody) SetRequestId(v string) *DescribeDcdnRefreshQuotaResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDcdnRefreshQuotaResponseBody) SetUrlQuota(v string) *DescribeDcdnRefreshQuotaResponseBody {
	s.UrlQuota = &v
	return s
}

func (s *DescribeDcdnRefreshQuotaResponseBody) SetUrlRemain(v string) *DescribeDcdnRefreshQuotaResponseBody {
	s.UrlRemain = &v
	return s
}

type DescribeDcdnRefreshQuotaResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDcdnRefreshQuotaResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDcdnRefreshQuotaResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnRefreshQuotaResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnRefreshQuotaResponse) SetHeaders(v map[string]*string) *DescribeDcdnRefreshQuotaResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnRefreshQuotaResponse) SetBody(v *DescribeDcdnRefreshQuotaResponseBody) *DescribeDcdnRefreshQuotaResponse {
	s.Body = v
	return s
}

type DescribeDcdnRefreshTaskByIdRequest struct {
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	TaskId  *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DescribeDcdnRefreshTaskByIdRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnRefreshTaskByIdRequest) GoString() string {
	return s.String()
}

func (s *DescribeDcdnRefreshTaskByIdRequest) SetOwnerId(v int64) *DescribeDcdnRefreshTaskByIdRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDcdnRefreshTaskByIdRequest) SetTaskId(v string) *DescribeDcdnRefreshTaskByIdRequest {
	s.TaskId = &v
	return s
}

type DescribeDcdnRefreshTaskByIdResponseBody struct {
	RequestId  *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Tasks      []*DescribeDcdnRefreshTaskByIdResponseBodyTasks `json:"Tasks,omitempty" xml:"Tasks,omitempty" type:"Repeated"`
	TotalCount *int64                                          `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDcdnRefreshTaskByIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnRefreshTaskByIdResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnRefreshTaskByIdResponseBody) SetRequestId(v string) *DescribeDcdnRefreshTaskByIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDcdnRefreshTaskByIdResponseBody) SetTasks(v []*DescribeDcdnRefreshTaskByIdResponseBodyTasks) *DescribeDcdnRefreshTaskByIdResponseBody {
	s.Tasks = v
	return s
}

func (s *DescribeDcdnRefreshTaskByIdResponseBody) SetTotalCount(v int64) *DescribeDcdnRefreshTaskByIdResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeDcdnRefreshTaskByIdResponseBodyTasks struct {
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	Description  *string `json:"Description,omitempty" xml:"Description,omitempty"`
	ObjectPath   *string `json:"ObjectPath,omitempty" xml:"ObjectPath,omitempty"`
	ObjectType   *string `json:"ObjectType,omitempty" xml:"ObjectType,omitempty"`
	Process      *string `json:"Process,omitempty" xml:"Process,omitempty"`
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
	TaskId       *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DescribeDcdnRefreshTaskByIdResponseBodyTasks) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnRefreshTaskByIdResponseBodyTasks) GoString() string {
	return s.String()
}

func (s *DescribeDcdnRefreshTaskByIdResponseBodyTasks) SetCreationTime(v string) *DescribeDcdnRefreshTaskByIdResponseBodyTasks {
	s.CreationTime = &v
	return s
}

func (s *DescribeDcdnRefreshTaskByIdResponseBodyTasks) SetDescription(v string) *DescribeDcdnRefreshTaskByIdResponseBodyTasks {
	s.Description = &v
	return s
}

func (s *DescribeDcdnRefreshTaskByIdResponseBodyTasks) SetObjectPath(v string) *DescribeDcdnRefreshTaskByIdResponseBodyTasks {
	s.ObjectPath = &v
	return s
}

func (s *DescribeDcdnRefreshTaskByIdResponseBodyTasks) SetObjectType(v string) *DescribeDcdnRefreshTaskByIdResponseBodyTasks {
	s.ObjectType = &v
	return s
}

func (s *DescribeDcdnRefreshTaskByIdResponseBodyTasks) SetProcess(v string) *DescribeDcdnRefreshTaskByIdResponseBodyTasks {
	s.Process = &v
	return s
}

func (s *DescribeDcdnRefreshTaskByIdResponseBodyTasks) SetStatus(v string) *DescribeDcdnRefreshTaskByIdResponseBodyTasks {
	s.Status = &v
	return s
}

func (s *DescribeDcdnRefreshTaskByIdResponseBodyTasks) SetTaskId(v string) *DescribeDcdnRefreshTaskByIdResponseBodyTasks {
	s.TaskId = &v
	return s
}

type DescribeDcdnRefreshTaskByIdResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDcdnRefreshTaskByIdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDcdnRefreshTaskByIdResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnRefreshTaskByIdResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnRefreshTaskByIdResponse) SetHeaders(v map[string]*string) *DescribeDcdnRefreshTaskByIdResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnRefreshTaskByIdResponse) SetBody(v *DescribeDcdnRefreshTaskByIdResponseBody) *DescribeDcdnRefreshTaskByIdResponse {
	s.Body = v
	return s
}

type DescribeDcdnRefreshTasksRequest struct {
	DomainName    *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EndTime       *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	ObjectPath    *string `json:"ObjectPath,omitempty" xml:"ObjectPath,omitempty"`
	ObjectType    *string `json:"ObjectType,omitempty" xml:"ObjectType,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNumber    *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize      *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	StartTime     *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Status        *string `json:"Status,omitempty" xml:"Status,omitempty"`
	TaskId        *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DescribeDcdnRefreshTasksRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnRefreshTasksRequest) GoString() string {
	return s.String()
}

func (s *DescribeDcdnRefreshTasksRequest) SetDomainName(v string) *DescribeDcdnRefreshTasksRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDcdnRefreshTasksRequest) SetEndTime(v string) *DescribeDcdnRefreshTasksRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDcdnRefreshTasksRequest) SetObjectPath(v string) *DescribeDcdnRefreshTasksRequest {
	s.ObjectPath = &v
	return s
}

func (s *DescribeDcdnRefreshTasksRequest) SetObjectType(v string) *DescribeDcdnRefreshTasksRequest {
	s.ObjectType = &v
	return s
}

func (s *DescribeDcdnRefreshTasksRequest) SetOwnerId(v int64) *DescribeDcdnRefreshTasksRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDcdnRefreshTasksRequest) SetPageNumber(v int32) *DescribeDcdnRefreshTasksRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDcdnRefreshTasksRequest) SetPageSize(v int32) *DescribeDcdnRefreshTasksRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDcdnRefreshTasksRequest) SetSecurityToken(v string) *DescribeDcdnRefreshTasksRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeDcdnRefreshTasksRequest) SetStartTime(v string) *DescribeDcdnRefreshTasksRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDcdnRefreshTasksRequest) SetStatus(v string) *DescribeDcdnRefreshTasksRequest {
	s.Status = &v
	return s
}

func (s *DescribeDcdnRefreshTasksRequest) SetTaskId(v string) *DescribeDcdnRefreshTasksRequest {
	s.TaskId = &v
	return s
}

type DescribeDcdnRefreshTasksResponseBody struct {
	PageNumber *int64                                     `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int64                                     `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Tasks      *DescribeDcdnRefreshTasksResponseBodyTasks `json:"Tasks,omitempty" xml:"Tasks,omitempty" type:"Struct"`
	TotalCount *int64                                     `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDcdnRefreshTasksResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnRefreshTasksResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnRefreshTasksResponseBody) SetPageNumber(v int64) *DescribeDcdnRefreshTasksResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeDcdnRefreshTasksResponseBody) SetPageSize(v int64) *DescribeDcdnRefreshTasksResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeDcdnRefreshTasksResponseBody) SetRequestId(v string) *DescribeDcdnRefreshTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDcdnRefreshTasksResponseBody) SetTasks(v *DescribeDcdnRefreshTasksResponseBodyTasks) *DescribeDcdnRefreshTasksResponseBody {
	s.Tasks = v
	return s
}

func (s *DescribeDcdnRefreshTasksResponseBody) SetTotalCount(v int64) *DescribeDcdnRefreshTasksResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeDcdnRefreshTasksResponseBodyTasks struct {
	Task []*DescribeDcdnRefreshTasksResponseBodyTasksTask `json:"Task,omitempty" xml:"Task,omitempty" type:"Repeated"`
}

func (s DescribeDcdnRefreshTasksResponseBodyTasks) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnRefreshTasksResponseBodyTasks) GoString() string {
	return s.String()
}

func (s *DescribeDcdnRefreshTasksResponseBodyTasks) SetTask(v []*DescribeDcdnRefreshTasksResponseBodyTasksTask) *DescribeDcdnRefreshTasksResponseBodyTasks {
	s.Task = v
	return s
}

type DescribeDcdnRefreshTasksResponseBodyTasksTask struct {
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	Description  *string `json:"Description,omitempty" xml:"Description,omitempty"`
	ObjectPath   *string `json:"ObjectPath,omitempty" xml:"ObjectPath,omitempty"`
	ObjectType   *string `json:"ObjectType,omitempty" xml:"ObjectType,omitempty"`
	Process      *string `json:"Process,omitempty" xml:"Process,omitempty"`
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
	TaskId       *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DescribeDcdnRefreshTasksResponseBodyTasksTask) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnRefreshTasksResponseBodyTasksTask) GoString() string {
	return s.String()
}

func (s *DescribeDcdnRefreshTasksResponseBodyTasksTask) SetCreationTime(v string) *DescribeDcdnRefreshTasksResponseBodyTasksTask {
	s.CreationTime = &v
	return s
}

func (s *DescribeDcdnRefreshTasksResponseBodyTasksTask) SetDescription(v string) *DescribeDcdnRefreshTasksResponseBodyTasksTask {
	s.Description = &v
	return s
}

func (s *DescribeDcdnRefreshTasksResponseBodyTasksTask) SetObjectPath(v string) *DescribeDcdnRefreshTasksResponseBodyTasksTask {
	s.ObjectPath = &v
	return s
}

func (s *DescribeDcdnRefreshTasksResponseBodyTasksTask) SetObjectType(v string) *DescribeDcdnRefreshTasksResponseBodyTasksTask {
	s.ObjectType = &v
	return s
}

func (s *DescribeDcdnRefreshTasksResponseBodyTasksTask) SetProcess(v string) *DescribeDcdnRefreshTasksResponseBodyTasksTask {
	s.Process = &v
	return s
}

func (s *DescribeDcdnRefreshTasksResponseBodyTasksTask) SetStatus(v string) *DescribeDcdnRefreshTasksResponseBodyTasksTask {
	s.Status = &v
	return s
}

func (s *DescribeDcdnRefreshTasksResponseBodyTasksTask) SetTaskId(v string) *DescribeDcdnRefreshTasksResponseBodyTasksTask {
	s.TaskId = &v
	return s
}

type DescribeDcdnRefreshTasksResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDcdnRefreshTasksResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDcdnRefreshTasksResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnRefreshTasksResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnRefreshTasksResponse) SetHeaders(v map[string]*string) *DescribeDcdnRefreshTasksResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnRefreshTasksResponse) SetBody(v *DescribeDcdnRefreshTasksResponseBody) *DescribeDcdnRefreshTasksResponse {
	s.Body = v
	return s
}

type DescribeDcdnRegionAndIspRequest struct {
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeDcdnRegionAndIspRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnRegionAndIspRequest) GoString() string {
	return s.String()
}

func (s *DescribeDcdnRegionAndIspRequest) SetOwnerId(v int64) *DescribeDcdnRegionAndIspRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDcdnRegionAndIspRequest) SetSecurityToken(v string) *DescribeDcdnRegionAndIspRequest {
	s.SecurityToken = &v
	return s
}

type DescribeDcdnRegionAndIspResponseBody struct {
	Isps      *DescribeDcdnRegionAndIspResponseBodyIsps    `json:"Isps,omitempty" xml:"Isps,omitempty" type:"Struct"`
	Regions   *DescribeDcdnRegionAndIspResponseBodyRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Struct"`
	RequestId *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDcdnRegionAndIspResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnRegionAndIspResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnRegionAndIspResponseBody) SetIsps(v *DescribeDcdnRegionAndIspResponseBodyIsps) *DescribeDcdnRegionAndIspResponseBody {
	s.Isps = v
	return s
}

func (s *DescribeDcdnRegionAndIspResponseBody) SetRegions(v *DescribeDcdnRegionAndIspResponseBodyRegions) *DescribeDcdnRegionAndIspResponseBody {
	s.Regions = v
	return s
}

func (s *DescribeDcdnRegionAndIspResponseBody) SetRequestId(v string) *DescribeDcdnRegionAndIspResponseBody {
	s.RequestId = &v
	return s
}

type DescribeDcdnRegionAndIspResponseBodyIsps struct {
	Isp []*DescribeDcdnRegionAndIspResponseBodyIspsIsp `json:"Isp,omitempty" xml:"Isp,omitempty" type:"Repeated"`
}

func (s DescribeDcdnRegionAndIspResponseBodyIsps) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnRegionAndIspResponseBodyIsps) GoString() string {
	return s.String()
}

func (s *DescribeDcdnRegionAndIspResponseBodyIsps) SetIsp(v []*DescribeDcdnRegionAndIspResponseBodyIspsIsp) *DescribeDcdnRegionAndIspResponseBodyIsps {
	s.Isp = v
	return s
}

type DescribeDcdnRegionAndIspResponseBodyIspsIsp struct {
	NameEn *string `json:"NameEn,omitempty" xml:"NameEn,omitempty"`
	NameZh *string `json:"NameZh,omitempty" xml:"NameZh,omitempty"`
}

func (s DescribeDcdnRegionAndIspResponseBodyIspsIsp) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnRegionAndIspResponseBodyIspsIsp) GoString() string {
	return s.String()
}

func (s *DescribeDcdnRegionAndIspResponseBodyIspsIsp) SetNameEn(v string) *DescribeDcdnRegionAndIspResponseBodyIspsIsp {
	s.NameEn = &v
	return s
}

func (s *DescribeDcdnRegionAndIspResponseBodyIspsIsp) SetNameZh(v string) *DescribeDcdnRegionAndIspResponseBodyIspsIsp {
	s.NameZh = &v
	return s
}

type DescribeDcdnRegionAndIspResponseBodyRegions struct {
	Region []*DescribeDcdnRegionAndIspResponseBodyRegionsRegion `json:"Region,omitempty" xml:"Region,omitempty" type:"Repeated"`
}

func (s DescribeDcdnRegionAndIspResponseBodyRegions) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnRegionAndIspResponseBodyRegions) GoString() string {
	return s.String()
}

func (s *DescribeDcdnRegionAndIspResponseBodyRegions) SetRegion(v []*DescribeDcdnRegionAndIspResponseBodyRegionsRegion) *DescribeDcdnRegionAndIspResponseBodyRegions {
	s.Region = v
	return s
}

type DescribeDcdnRegionAndIspResponseBodyRegionsRegion struct {
	NameEn *string `json:"NameEn,omitempty" xml:"NameEn,omitempty"`
	NameZh *string `json:"NameZh,omitempty" xml:"NameZh,omitempty"`
}

func (s DescribeDcdnRegionAndIspResponseBodyRegionsRegion) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnRegionAndIspResponseBodyRegionsRegion) GoString() string {
	return s.String()
}

func (s *DescribeDcdnRegionAndIspResponseBodyRegionsRegion) SetNameEn(v string) *DescribeDcdnRegionAndIspResponseBodyRegionsRegion {
	s.NameEn = &v
	return s
}

func (s *DescribeDcdnRegionAndIspResponseBodyRegionsRegion) SetNameZh(v string) *DescribeDcdnRegionAndIspResponseBodyRegionsRegion {
	s.NameZh = &v
	return s
}

type DescribeDcdnRegionAndIspResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDcdnRegionAndIspResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDcdnRegionAndIspResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnRegionAndIspResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnRegionAndIspResponse) SetHeaders(v map[string]*string) *DescribeDcdnRegionAndIspResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnRegionAndIspResponse) SetBody(v *DescribeDcdnRegionAndIspResponseBody) *DescribeDcdnRegionAndIspResponse {
	s.Body = v
	return s
}

type DescribeDcdnReportRequest struct {
	Area       *string `json:"Area,omitempty" xml:"Area,omitempty"`
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EndTime    *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	HttpCode   *string `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	IsOverseas *string `json:"IsOverseas,omitempty" xml:"IsOverseas,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ReportId   *int64  `json:"ReportId,omitempty" xml:"ReportId,omitempty"`
	StartTime  *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDcdnReportRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnReportRequest) GoString() string {
	return s.String()
}

func (s *DescribeDcdnReportRequest) SetArea(v string) *DescribeDcdnReportRequest {
	s.Area = &v
	return s
}

func (s *DescribeDcdnReportRequest) SetDomainName(v string) *DescribeDcdnReportRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDcdnReportRequest) SetEndTime(v string) *DescribeDcdnReportRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDcdnReportRequest) SetHttpCode(v string) *DescribeDcdnReportRequest {
	s.HttpCode = &v
	return s
}

func (s *DescribeDcdnReportRequest) SetIsOverseas(v string) *DescribeDcdnReportRequest {
	s.IsOverseas = &v
	return s
}

func (s *DescribeDcdnReportRequest) SetOwnerId(v int64) *DescribeDcdnReportRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDcdnReportRequest) SetReportId(v int64) *DescribeDcdnReportRequest {
	s.ReportId = &v
	return s
}

func (s *DescribeDcdnReportRequest) SetStartTime(v string) *DescribeDcdnReportRequest {
	s.StartTime = &v
	return s
}

type DescribeDcdnReportResponseBody struct {
	Content   *string `json:"Content,omitempty" xml:"Content,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDcdnReportResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnReportResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnReportResponseBody) SetContent(v string) *DescribeDcdnReportResponseBody {
	s.Content = &v
	return s
}

func (s *DescribeDcdnReportResponseBody) SetRequestId(v string) *DescribeDcdnReportResponseBody {
	s.RequestId = &v
	return s
}

type DescribeDcdnReportResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDcdnReportResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDcdnReportResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnReportResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnReportResponse) SetHeaders(v map[string]*string) *DescribeDcdnReportResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnReportResponse) SetBody(v *DescribeDcdnReportResponseBody) *DescribeDcdnReportResponse {
	s.Body = v
	return s
}

type DescribeDcdnReportListRequest struct {
	OwnerId  *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ReportId *int64 `json:"ReportId,omitempty" xml:"ReportId,omitempty"`
}

func (s DescribeDcdnReportListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnReportListRequest) GoString() string {
	return s.String()
}

func (s *DescribeDcdnReportListRequest) SetOwnerId(v int64) *DescribeDcdnReportListRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDcdnReportListRequest) SetReportId(v int64) *DescribeDcdnReportListRequest {
	s.ReportId = &v
	return s
}

type DescribeDcdnReportListResponseBody struct {
	Content   *string `json:"Content,omitempty" xml:"Content,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDcdnReportListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnReportListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnReportListResponseBody) SetContent(v string) *DescribeDcdnReportListResponseBody {
	s.Content = &v
	return s
}

func (s *DescribeDcdnReportListResponseBody) SetRequestId(v string) *DescribeDcdnReportListResponseBody {
	s.RequestId = &v
	return s
}

type DescribeDcdnReportListResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDcdnReportListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDcdnReportListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnReportListResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnReportListResponse) SetHeaders(v map[string]*string) *DescribeDcdnReportListResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnReportListResponse) SetBody(v *DescribeDcdnReportListResponseBody) *DescribeDcdnReportListResponse {
	s.Body = v
	return s
}

type DescribeDcdnSLSRealtimeLogDeliveryRequest struct {
	BusinessType *string `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ProjectName  *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
}

func (s DescribeDcdnSLSRealtimeLogDeliveryRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnSLSRealtimeLogDeliveryRequest) GoString() string {
	return s.String()
}

func (s *DescribeDcdnSLSRealtimeLogDeliveryRequest) SetBusinessType(v string) *DescribeDcdnSLSRealtimeLogDeliveryRequest {
	s.BusinessType = &v
	return s
}

func (s *DescribeDcdnSLSRealtimeLogDeliveryRequest) SetOwnerId(v int64) *DescribeDcdnSLSRealtimeLogDeliveryRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDcdnSLSRealtimeLogDeliveryRequest) SetProjectName(v string) *DescribeDcdnSLSRealtimeLogDeliveryRequest {
	s.ProjectName = &v
	return s
}

type DescribeDcdnSLSRealtimeLogDeliveryResponseBody struct {
	Content   *DescribeDcdnSLSRealtimeLogDeliveryResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	RequestId *string                                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDcdnSLSRealtimeLogDeliveryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnSLSRealtimeLogDeliveryResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnSLSRealtimeLogDeliveryResponseBody) SetContent(v *DescribeDcdnSLSRealtimeLogDeliveryResponseBodyContent) *DescribeDcdnSLSRealtimeLogDeliveryResponseBody {
	s.Content = v
	return s
}

func (s *DescribeDcdnSLSRealtimeLogDeliveryResponseBody) SetRequestId(v string) *DescribeDcdnSLSRealtimeLogDeliveryResponseBody {
	s.RequestId = &v
	return s
}

type DescribeDcdnSLSRealtimeLogDeliveryResponseBodyContent struct {
	BusinessType *string `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
	DataCenter   *string `json:"DataCenter,omitempty" xml:"DataCenter,omitempty"`
	DomainName   *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	FieldName    *string `json:"FieldName,omitempty" xml:"FieldName,omitempty"`
	ProjectName  *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	SLSLogStore  *string `json:"SLSLogStore,omitempty" xml:"SLSLogStore,omitempty"`
	SLSProject   *string `json:"SLSProject,omitempty" xml:"SLSProject,omitempty"`
	SLSRegion    *string `json:"SLSRegion,omitempty" xml:"SLSRegion,omitempty"`
	SamplingRate *string `json:"SamplingRate,omitempty" xml:"SamplingRate,omitempty"`
	Type         *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeDcdnSLSRealtimeLogDeliveryResponseBodyContent) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnSLSRealtimeLogDeliveryResponseBodyContent) GoString() string {
	return s.String()
}

func (s *DescribeDcdnSLSRealtimeLogDeliveryResponseBodyContent) SetBusinessType(v string) *DescribeDcdnSLSRealtimeLogDeliveryResponseBodyContent {
	s.BusinessType = &v
	return s
}

func (s *DescribeDcdnSLSRealtimeLogDeliveryResponseBodyContent) SetDataCenter(v string) *DescribeDcdnSLSRealtimeLogDeliveryResponseBodyContent {
	s.DataCenter = &v
	return s
}

func (s *DescribeDcdnSLSRealtimeLogDeliveryResponseBodyContent) SetDomainName(v string) *DescribeDcdnSLSRealtimeLogDeliveryResponseBodyContent {
	s.DomainName = &v
	return s
}

func (s *DescribeDcdnSLSRealtimeLogDeliveryResponseBodyContent) SetFieldName(v string) *DescribeDcdnSLSRealtimeLogDeliveryResponseBodyContent {
	s.FieldName = &v
	return s
}

func (s *DescribeDcdnSLSRealtimeLogDeliveryResponseBodyContent) SetProjectName(v string) *DescribeDcdnSLSRealtimeLogDeliveryResponseBodyContent {
	s.ProjectName = &v
	return s
}

func (s *DescribeDcdnSLSRealtimeLogDeliveryResponseBodyContent) SetSLSLogStore(v string) *DescribeDcdnSLSRealtimeLogDeliveryResponseBodyContent {
	s.SLSLogStore = &v
	return s
}

func (s *DescribeDcdnSLSRealtimeLogDeliveryResponseBodyContent) SetSLSProject(v string) *DescribeDcdnSLSRealtimeLogDeliveryResponseBodyContent {
	s.SLSProject = &v
	return s
}

func (s *DescribeDcdnSLSRealtimeLogDeliveryResponseBodyContent) SetSLSRegion(v string) *DescribeDcdnSLSRealtimeLogDeliveryResponseBodyContent {
	s.SLSRegion = &v
	return s
}

func (s *DescribeDcdnSLSRealtimeLogDeliveryResponseBodyContent) SetSamplingRate(v string) *DescribeDcdnSLSRealtimeLogDeliveryResponseBodyContent {
	s.SamplingRate = &v
	return s
}

func (s *DescribeDcdnSLSRealtimeLogDeliveryResponseBodyContent) SetType(v string) *DescribeDcdnSLSRealtimeLogDeliveryResponseBodyContent {
	s.Type = &v
	return s
}

type DescribeDcdnSLSRealtimeLogDeliveryResponse struct {
	Headers map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDcdnSLSRealtimeLogDeliveryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDcdnSLSRealtimeLogDeliveryResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnSLSRealtimeLogDeliveryResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnSLSRealtimeLogDeliveryResponse) SetHeaders(v map[string]*string) *DescribeDcdnSLSRealtimeLogDeliveryResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnSLSRealtimeLogDeliveryResponse) SetBody(v *DescribeDcdnSLSRealtimeLogDeliveryResponseBody) *DescribeDcdnSLSRealtimeLogDeliveryResponse {
	s.Body = v
	return s
}

type DescribeDcdnSMCertificateDetailRequest struct {
	CertIdentifier *string `json:"CertIdentifier,omitempty" xml:"CertIdentifier,omitempty"`
	OwnerId        *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken  *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeDcdnSMCertificateDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnSMCertificateDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribeDcdnSMCertificateDetailRequest) SetCertIdentifier(v string) *DescribeDcdnSMCertificateDetailRequest {
	s.CertIdentifier = &v
	return s
}

func (s *DescribeDcdnSMCertificateDetailRequest) SetOwnerId(v int64) *DescribeDcdnSMCertificateDetailRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDcdnSMCertificateDetailRequest) SetSecurityToken(v string) *DescribeDcdnSMCertificateDetailRequest {
	s.SecurityToken = &v
	return s
}

type DescribeDcdnSMCertificateDetailResponseBody struct {
	CertExpireTime     *string `json:"CertExpireTime,omitempty" xml:"CertExpireTime,omitempty"`
	CertIdentifier     *string `json:"CertIdentifier,omitempty" xml:"CertIdentifier,omitempty"`
	CertName           *string `json:"CertName,omitempty" xml:"CertName,omitempty"`
	CertOrg            *string `json:"CertOrg,omitempty" xml:"CertOrg,omitempty"`
	CommonName         *string `json:"CommonName,omitempty" xml:"CommonName,omitempty"`
	EncryptCertificate *string `json:"EncryptCertificate,omitempty" xml:"EncryptCertificate,omitempty"`
	RequestId          *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Sans               *string `json:"Sans,omitempty" xml:"Sans,omitempty"`
	SignCertificate    *string `json:"SignCertificate,omitempty" xml:"SignCertificate,omitempty"`
}

func (s DescribeDcdnSMCertificateDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnSMCertificateDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnSMCertificateDetailResponseBody) SetCertExpireTime(v string) *DescribeDcdnSMCertificateDetailResponseBody {
	s.CertExpireTime = &v
	return s
}

func (s *DescribeDcdnSMCertificateDetailResponseBody) SetCertIdentifier(v string) *DescribeDcdnSMCertificateDetailResponseBody {
	s.CertIdentifier = &v
	return s
}

func (s *DescribeDcdnSMCertificateDetailResponseBody) SetCertName(v string) *DescribeDcdnSMCertificateDetailResponseBody {
	s.CertName = &v
	return s
}

func (s *DescribeDcdnSMCertificateDetailResponseBody) SetCertOrg(v string) *DescribeDcdnSMCertificateDetailResponseBody {
	s.CertOrg = &v
	return s
}

func (s *DescribeDcdnSMCertificateDetailResponseBody) SetCommonName(v string) *DescribeDcdnSMCertificateDetailResponseBody {
	s.CommonName = &v
	return s
}

func (s *DescribeDcdnSMCertificateDetailResponseBody) SetEncryptCertificate(v string) *DescribeDcdnSMCertificateDetailResponseBody {
	s.EncryptCertificate = &v
	return s
}

func (s *DescribeDcdnSMCertificateDetailResponseBody) SetRequestId(v string) *DescribeDcdnSMCertificateDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDcdnSMCertificateDetailResponseBody) SetSans(v string) *DescribeDcdnSMCertificateDetailResponseBody {
	s.Sans = &v
	return s
}

func (s *DescribeDcdnSMCertificateDetailResponseBody) SetSignCertificate(v string) *DescribeDcdnSMCertificateDetailResponseBody {
	s.SignCertificate = &v
	return s
}

type DescribeDcdnSMCertificateDetailResponse struct {
	Headers map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDcdnSMCertificateDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDcdnSMCertificateDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnSMCertificateDetailResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnSMCertificateDetailResponse) SetHeaders(v map[string]*string) *DescribeDcdnSMCertificateDetailResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnSMCertificateDetailResponse) SetBody(v *DescribeDcdnSMCertificateDetailResponseBody) *DescribeDcdnSMCertificateDetailResponse {
	s.Body = v
	return s
}

type DescribeDcdnSMCertificateListRequest struct {
	DomainName    *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeDcdnSMCertificateListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnSMCertificateListRequest) GoString() string {
	return s.String()
}

func (s *DescribeDcdnSMCertificateListRequest) SetDomainName(v string) *DescribeDcdnSMCertificateListRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDcdnSMCertificateListRequest) SetOwnerId(v int64) *DescribeDcdnSMCertificateListRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDcdnSMCertificateListRequest) SetSecurityToken(v string) *DescribeDcdnSMCertificateListRequest {
	s.SecurityToken = &v
	return s
}

type DescribeDcdnSMCertificateListResponseBody struct {
	CertificateListModel *DescribeDcdnSMCertificateListResponseBodyCertificateListModel `json:"CertificateListModel,omitempty" xml:"CertificateListModel,omitempty" type:"Struct"`
	RequestId            *string                                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDcdnSMCertificateListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnSMCertificateListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnSMCertificateListResponseBody) SetCertificateListModel(v *DescribeDcdnSMCertificateListResponseBodyCertificateListModel) *DescribeDcdnSMCertificateListResponseBody {
	s.CertificateListModel = v
	return s
}

func (s *DescribeDcdnSMCertificateListResponseBody) SetRequestId(v string) *DescribeDcdnSMCertificateListResponseBody {
	s.RequestId = &v
	return s
}

type DescribeDcdnSMCertificateListResponseBodyCertificateListModel struct {
	CertList []*DescribeDcdnSMCertificateListResponseBodyCertificateListModelCertList `json:"CertList,omitempty" xml:"CertList,omitempty" type:"Repeated"`
	Count    *int32                                                                   `json:"Count,omitempty" xml:"Count,omitempty"`
}

func (s DescribeDcdnSMCertificateListResponseBodyCertificateListModel) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnSMCertificateListResponseBodyCertificateListModel) GoString() string {
	return s.String()
}

func (s *DescribeDcdnSMCertificateListResponseBodyCertificateListModel) SetCertList(v []*DescribeDcdnSMCertificateListResponseBodyCertificateListModelCertList) *DescribeDcdnSMCertificateListResponseBodyCertificateListModel {
	s.CertList = v
	return s
}

func (s *DescribeDcdnSMCertificateListResponseBodyCertificateListModel) SetCount(v int32) *DescribeDcdnSMCertificateListResponseBodyCertificateListModel {
	s.Count = &v
	return s
}

type DescribeDcdnSMCertificateListResponseBodyCertificateListModelCertList struct {
	CertIdentifier *string `json:"CertIdentifier,omitempty" xml:"CertIdentifier,omitempty"`
	CertName       *string `json:"CertName,omitempty" xml:"CertName,omitempty"`
	Common         *string `json:"Common,omitempty" xml:"Common,omitempty"`
	Issuer         *string `json:"Issuer,omitempty" xml:"Issuer,omitempty"`
}

func (s DescribeDcdnSMCertificateListResponseBodyCertificateListModelCertList) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnSMCertificateListResponseBodyCertificateListModelCertList) GoString() string {
	return s.String()
}

func (s *DescribeDcdnSMCertificateListResponseBodyCertificateListModelCertList) SetCertIdentifier(v string) *DescribeDcdnSMCertificateListResponseBodyCertificateListModelCertList {
	s.CertIdentifier = &v
	return s
}

func (s *DescribeDcdnSMCertificateListResponseBodyCertificateListModelCertList) SetCertName(v string) *DescribeDcdnSMCertificateListResponseBodyCertificateListModelCertList {
	s.CertName = &v
	return s
}

func (s *DescribeDcdnSMCertificateListResponseBodyCertificateListModelCertList) SetCommon(v string) *DescribeDcdnSMCertificateListResponseBodyCertificateListModelCertList {
	s.Common = &v
	return s
}

func (s *DescribeDcdnSMCertificateListResponseBodyCertificateListModelCertList) SetIssuer(v string) *DescribeDcdnSMCertificateListResponseBodyCertificateListModelCertList {
	s.Issuer = &v
	return s
}

type DescribeDcdnSMCertificateListResponse struct {
	Headers map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDcdnSMCertificateListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDcdnSMCertificateListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnSMCertificateListResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnSMCertificateListResponse) SetHeaders(v map[string]*string) *DescribeDcdnSMCertificateListResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnSMCertificateListResponse) SetBody(v *DescribeDcdnSMCertificateListResponseBody) *DescribeDcdnSMCertificateListResponse {
	s.Body = v
	return s
}

type DescribeDcdnSecFuncInfoRequest struct {
	Lang        *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	OwnerId     *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecFuncType *string `json:"SecFuncType,omitempty" xml:"SecFuncType,omitempty"`
}

func (s DescribeDcdnSecFuncInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnSecFuncInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeDcdnSecFuncInfoRequest) SetLang(v string) *DescribeDcdnSecFuncInfoRequest {
	s.Lang = &v
	return s
}

func (s *DescribeDcdnSecFuncInfoRequest) SetOwnerId(v int64) *DescribeDcdnSecFuncInfoRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDcdnSecFuncInfoRequest) SetSecFuncType(v string) *DescribeDcdnSecFuncInfoRequest {
	s.SecFuncType = &v
	return s
}

type DescribeDcdnSecFuncInfoResponseBody struct {
	Content     []*DescribeDcdnSecFuncInfoResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Repeated"`
	Description *string                                       `json:"Description,omitempty" xml:"Description,omitempty"`
	HttpStatus  *string                                       `json:"HttpStatus,omitempty" xml:"HttpStatus,omitempty"`
	RequestId   *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RetCode     *string                                       `json:"RetCode,omitempty" xml:"RetCode,omitempty"`
}

func (s DescribeDcdnSecFuncInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnSecFuncInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnSecFuncInfoResponseBody) SetContent(v []*DescribeDcdnSecFuncInfoResponseBodyContent) *DescribeDcdnSecFuncInfoResponseBody {
	s.Content = v
	return s
}

func (s *DescribeDcdnSecFuncInfoResponseBody) SetDescription(v string) *DescribeDcdnSecFuncInfoResponseBody {
	s.Description = &v
	return s
}

func (s *DescribeDcdnSecFuncInfoResponseBody) SetHttpStatus(v string) *DescribeDcdnSecFuncInfoResponseBody {
	s.HttpStatus = &v
	return s
}

func (s *DescribeDcdnSecFuncInfoResponseBody) SetRequestId(v string) *DescribeDcdnSecFuncInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDcdnSecFuncInfoResponseBody) SetRetCode(v string) *DescribeDcdnSecFuncInfoResponseBody {
	s.RetCode = &v
	return s
}

type DescribeDcdnSecFuncInfoResponseBodyContent struct {
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDcdnSecFuncInfoResponseBodyContent) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnSecFuncInfoResponseBodyContent) GoString() string {
	return s.String()
}

func (s *DescribeDcdnSecFuncInfoResponseBodyContent) SetLabel(v string) *DescribeDcdnSecFuncInfoResponseBodyContent {
	s.Label = &v
	return s
}

func (s *DescribeDcdnSecFuncInfoResponseBodyContent) SetValue(v string) *DescribeDcdnSecFuncInfoResponseBodyContent {
	s.Value = &v
	return s
}

type DescribeDcdnSecFuncInfoResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDcdnSecFuncInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDcdnSecFuncInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnSecFuncInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnSecFuncInfoResponse) SetHeaders(v map[string]*string) *DescribeDcdnSecFuncInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnSecFuncInfoResponse) SetBody(v *DescribeDcdnSecFuncInfoResponseBody) *DescribeDcdnSecFuncInfoResponse {
	s.Body = v
	return s
}

type DescribeDcdnSecSpecInfoRequest struct {
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s DescribeDcdnSecSpecInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnSecSpecInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeDcdnSecSpecInfoRequest) SetOwnerId(v int64) *DescribeDcdnSecSpecInfoRequest {
	s.OwnerId = &v
	return s
}

type DescribeDcdnSecSpecInfoResponseBody struct {
	RequestId *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SpecInfos []*DescribeDcdnSecSpecInfoResponseBodySpecInfos `json:"SpecInfos,omitempty" xml:"SpecInfos,omitempty" type:"Repeated"`
	Version   *string                                         `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s DescribeDcdnSecSpecInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnSecSpecInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnSecSpecInfoResponseBody) SetRequestId(v string) *DescribeDcdnSecSpecInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDcdnSecSpecInfoResponseBody) SetSpecInfos(v []*DescribeDcdnSecSpecInfoResponseBodySpecInfos) *DescribeDcdnSecSpecInfoResponseBody {
	s.SpecInfos = v
	return s
}

func (s *DescribeDcdnSecSpecInfoResponseBody) SetVersion(v string) *DescribeDcdnSecSpecInfoResponseBody {
	s.Version = &v
	return s
}

type DescribeDcdnSecSpecInfoResponseBodySpecInfos struct {
	RuleCode    *string                                                    `json:"RuleCode,omitempty" xml:"RuleCode,omitempty"`
	RuleConfigs []*DescribeDcdnSecSpecInfoResponseBodySpecInfosRuleConfigs `json:"RuleConfigs,omitempty" xml:"RuleConfigs,omitempty" type:"Repeated"`
}

func (s DescribeDcdnSecSpecInfoResponseBodySpecInfos) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnSecSpecInfoResponseBodySpecInfos) GoString() string {
	return s.String()
}

func (s *DescribeDcdnSecSpecInfoResponseBodySpecInfos) SetRuleCode(v string) *DescribeDcdnSecSpecInfoResponseBodySpecInfos {
	s.RuleCode = &v
	return s
}

func (s *DescribeDcdnSecSpecInfoResponseBodySpecInfos) SetRuleConfigs(v []*DescribeDcdnSecSpecInfoResponseBodySpecInfosRuleConfigs) *DescribeDcdnSecSpecInfoResponseBodySpecInfos {
	s.RuleConfigs = v
	return s
}

type DescribeDcdnSecSpecInfoResponseBodySpecInfosRuleConfigs struct {
	Code  *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Expr  *string `json:"Expr,omitempty" xml:"Expr,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDcdnSecSpecInfoResponseBodySpecInfosRuleConfigs) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnSecSpecInfoResponseBodySpecInfosRuleConfigs) GoString() string {
	return s.String()
}

func (s *DescribeDcdnSecSpecInfoResponseBodySpecInfosRuleConfigs) SetCode(v string) *DescribeDcdnSecSpecInfoResponseBodySpecInfosRuleConfigs {
	s.Code = &v
	return s
}

func (s *DescribeDcdnSecSpecInfoResponseBodySpecInfosRuleConfigs) SetExpr(v string) *DescribeDcdnSecSpecInfoResponseBodySpecInfosRuleConfigs {
	s.Expr = &v
	return s
}

func (s *DescribeDcdnSecSpecInfoResponseBodySpecInfosRuleConfigs) SetValue(v string) *DescribeDcdnSecSpecInfoResponseBodySpecInfosRuleConfigs {
	s.Value = &v
	return s
}

type DescribeDcdnSecSpecInfoResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDcdnSecSpecInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDcdnSecSpecInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnSecSpecInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnSecSpecInfoResponse) SetHeaders(v map[string]*string) *DescribeDcdnSecSpecInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnSecSpecInfoResponse) SetBody(v *DescribeDcdnSecSpecInfoResponseBody) *DescribeDcdnSecSpecInfoResponse {
	s.Body = v
	return s
}

type DescribeDcdnServiceRequest struct {
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeDcdnServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnServiceRequest) GoString() string {
	return s.String()
}

func (s *DescribeDcdnServiceRequest) SetOwnerId(v int64) *DescribeDcdnServiceRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDcdnServiceRequest) SetSecurityToken(v string) *DescribeDcdnServiceRequest {
	s.SecurityToken = &v
	return s
}

type DescribeDcdnServiceResponseBody struct {
	ChangingAffectTime    *string                                        `json:"ChangingAffectTime,omitempty" xml:"ChangingAffectTime,omitempty"`
	ChangingChargeType    *string                                        `json:"ChangingChargeType,omitempty" xml:"ChangingChargeType,omitempty"`
	InstanceId            *string                                        `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InternetChargeType    *string                                        `json:"InternetChargeType,omitempty" xml:"InternetChargeType,omitempty"`
	OpeningTime           *string                                        `json:"OpeningTime,omitempty" xml:"OpeningTime,omitempty"`
	OperationLocks        *DescribeDcdnServiceResponseBodyOperationLocks `json:"OperationLocks,omitempty" xml:"OperationLocks,omitempty" type:"Struct"`
	RequestId             *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	WebsocketChangingTime *string                                        `json:"WebsocketChangingTime,omitempty" xml:"WebsocketChangingTime,omitempty"`
	WebsocketChangingType *string                                        `json:"WebsocketChangingType,omitempty" xml:"WebsocketChangingType,omitempty"`
	WebsocketType         *string                                        `json:"WebsocketType,omitempty" xml:"WebsocketType,omitempty"`
}

func (s DescribeDcdnServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnServiceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnServiceResponseBody) SetChangingAffectTime(v string) *DescribeDcdnServiceResponseBody {
	s.ChangingAffectTime = &v
	return s
}

func (s *DescribeDcdnServiceResponseBody) SetChangingChargeType(v string) *DescribeDcdnServiceResponseBody {
	s.ChangingChargeType = &v
	return s
}

func (s *DescribeDcdnServiceResponseBody) SetInstanceId(v string) *DescribeDcdnServiceResponseBody {
	s.InstanceId = &v
	return s
}

func (s *DescribeDcdnServiceResponseBody) SetInternetChargeType(v string) *DescribeDcdnServiceResponseBody {
	s.InternetChargeType = &v
	return s
}

func (s *DescribeDcdnServiceResponseBody) SetOpeningTime(v string) *DescribeDcdnServiceResponseBody {
	s.OpeningTime = &v
	return s
}

func (s *DescribeDcdnServiceResponseBody) SetOperationLocks(v *DescribeDcdnServiceResponseBodyOperationLocks) *DescribeDcdnServiceResponseBody {
	s.OperationLocks = v
	return s
}

func (s *DescribeDcdnServiceResponseBody) SetRequestId(v string) *DescribeDcdnServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDcdnServiceResponseBody) SetWebsocketChangingTime(v string) *DescribeDcdnServiceResponseBody {
	s.WebsocketChangingTime = &v
	return s
}

func (s *DescribeDcdnServiceResponseBody) SetWebsocketChangingType(v string) *DescribeDcdnServiceResponseBody {
	s.WebsocketChangingType = &v
	return s
}

func (s *DescribeDcdnServiceResponseBody) SetWebsocketType(v string) *DescribeDcdnServiceResponseBody {
	s.WebsocketType = &v
	return s
}

type DescribeDcdnServiceResponseBodyOperationLocks struct {
	LockReason []*DescribeDcdnServiceResponseBodyOperationLocksLockReason `json:"LockReason,omitempty" xml:"LockReason,omitempty" type:"Repeated"`
}

func (s DescribeDcdnServiceResponseBodyOperationLocks) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnServiceResponseBodyOperationLocks) GoString() string {
	return s.String()
}

func (s *DescribeDcdnServiceResponseBodyOperationLocks) SetLockReason(v []*DescribeDcdnServiceResponseBodyOperationLocksLockReason) *DescribeDcdnServiceResponseBodyOperationLocks {
	s.LockReason = v
	return s
}

type DescribeDcdnServiceResponseBodyOperationLocksLockReason struct {
	LockReason *string `json:"LockReason,omitempty" xml:"LockReason,omitempty"`
}

func (s DescribeDcdnServiceResponseBodyOperationLocksLockReason) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnServiceResponseBodyOperationLocksLockReason) GoString() string {
	return s.String()
}

func (s *DescribeDcdnServiceResponseBodyOperationLocksLockReason) SetLockReason(v string) *DescribeDcdnServiceResponseBodyOperationLocksLockReason {
	s.LockReason = &v
	return s
}

type DescribeDcdnServiceResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDcdnServiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDcdnServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnServiceResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnServiceResponse) SetHeaders(v map[string]*string) *DescribeDcdnServiceResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnServiceResponse) SetBody(v *DescribeDcdnServiceResponseBody) *DescribeDcdnServiceResponse {
	s.Body = v
	return s
}

type DescribeDcdnStagingIpRequest struct {
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s DescribeDcdnStagingIpRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnStagingIpRequest) GoString() string {
	return s.String()
}

func (s *DescribeDcdnStagingIpRequest) SetOwnerId(v int64) *DescribeDcdnStagingIpRequest {
	s.OwnerId = &v
	return s
}

type DescribeDcdnStagingIpResponseBody struct {
	IPV4s     *DescribeDcdnStagingIpResponseBodyIPV4s `json:"IPV4s,omitempty" xml:"IPV4s,omitempty" type:"Struct"`
	RequestId *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDcdnStagingIpResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnStagingIpResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnStagingIpResponseBody) SetIPV4s(v *DescribeDcdnStagingIpResponseBodyIPV4s) *DescribeDcdnStagingIpResponseBody {
	s.IPV4s = v
	return s
}

func (s *DescribeDcdnStagingIpResponseBody) SetRequestId(v string) *DescribeDcdnStagingIpResponseBody {
	s.RequestId = &v
	return s
}

type DescribeDcdnStagingIpResponseBodyIPV4s struct {
	IPV4 []*string `json:"IPV4,omitempty" xml:"IPV4,omitempty" type:"Repeated"`
}

func (s DescribeDcdnStagingIpResponseBodyIPV4s) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnStagingIpResponseBodyIPV4s) GoString() string {
	return s.String()
}

func (s *DescribeDcdnStagingIpResponseBodyIPV4s) SetIPV4(v []*string) *DescribeDcdnStagingIpResponseBodyIPV4s {
	s.IPV4 = v
	return s
}

type DescribeDcdnStagingIpResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDcdnStagingIpResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDcdnStagingIpResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnStagingIpResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnStagingIpResponse) SetHeaders(v map[string]*string) *DescribeDcdnStagingIpResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnStagingIpResponse) SetBody(v *DescribeDcdnStagingIpResponseBody) *DescribeDcdnStagingIpResponse {
	s.Body = v
	return s
}

type DescribeDcdnSubListRequest struct {
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s DescribeDcdnSubListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnSubListRequest) GoString() string {
	return s.String()
}

func (s *DescribeDcdnSubListRequest) SetOwnerId(v int64) *DescribeDcdnSubListRequest {
	s.OwnerId = &v
	return s
}

type DescribeDcdnSubListResponseBody struct {
	Content   *string `json:"Content,omitempty" xml:"Content,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDcdnSubListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnSubListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnSubListResponseBody) SetContent(v string) *DescribeDcdnSubListResponseBody {
	s.Content = &v
	return s
}

func (s *DescribeDcdnSubListResponseBody) SetRequestId(v string) *DescribeDcdnSubListResponseBody {
	s.RequestId = &v
	return s
}

type DescribeDcdnSubListResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDcdnSubListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDcdnSubListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnSubListResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnSubListResponse) SetHeaders(v map[string]*string) *DescribeDcdnSubListResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnSubListResponse) SetBody(v *DescribeDcdnSubListResponseBody) *DescribeDcdnSubListResponse {
	s.Body = v
	return s
}

type DescribeDcdnTagResourcesRequest struct {
	OwnerId      *int64                                `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceId   []*string                             `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	ResourceType *string                               `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	Tag          []*DescribeDcdnTagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeDcdnTagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnTagResourcesRequest) GoString() string {
	return s.String()
}

func (s *DescribeDcdnTagResourcesRequest) SetOwnerId(v int64) *DescribeDcdnTagResourcesRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDcdnTagResourcesRequest) SetResourceId(v []*string) *DescribeDcdnTagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *DescribeDcdnTagResourcesRequest) SetResourceType(v string) *DescribeDcdnTagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *DescribeDcdnTagResourcesRequest) SetTag(v []*DescribeDcdnTagResourcesRequestTag) *DescribeDcdnTagResourcesRequest {
	s.Tag = v
	return s
}

type DescribeDcdnTagResourcesRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDcdnTagResourcesRequestTag) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnTagResourcesRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeDcdnTagResourcesRequestTag) SetKey(v string) *DescribeDcdnTagResourcesRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeDcdnTagResourcesRequestTag) SetValue(v string) *DescribeDcdnTagResourcesRequestTag {
	s.Value = &v
	return s
}

type DescribeDcdnTagResourcesResponseBody struct {
	RequestId    *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TagResources []*DescribeDcdnTagResourcesResponseBodyTagResources `json:"TagResources,omitempty" xml:"TagResources,omitempty" type:"Repeated"`
}

func (s DescribeDcdnTagResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnTagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnTagResourcesResponseBody) SetRequestId(v string) *DescribeDcdnTagResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDcdnTagResourcesResponseBody) SetTagResources(v []*DescribeDcdnTagResourcesResponseBodyTagResources) *DescribeDcdnTagResourcesResponseBody {
	s.TagResources = v
	return s
}

type DescribeDcdnTagResourcesResponseBodyTagResources struct {
	ResourceId *string                                                `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	Tag        []*DescribeDcdnTagResourcesResponseBodyTagResourcesTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeDcdnTagResourcesResponseBodyTagResources) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnTagResourcesResponseBodyTagResources) GoString() string {
	return s.String()
}

func (s *DescribeDcdnTagResourcesResponseBodyTagResources) SetResourceId(v string) *DescribeDcdnTagResourcesResponseBodyTagResources {
	s.ResourceId = &v
	return s
}

func (s *DescribeDcdnTagResourcesResponseBodyTagResources) SetTag(v []*DescribeDcdnTagResourcesResponseBodyTagResourcesTag) *DescribeDcdnTagResourcesResponseBodyTagResources {
	s.Tag = v
	return s
}

type DescribeDcdnTagResourcesResponseBodyTagResourcesTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDcdnTagResourcesResponseBodyTagResourcesTag) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnTagResourcesResponseBodyTagResourcesTag) GoString() string {
	return s.String()
}

func (s *DescribeDcdnTagResourcesResponseBodyTagResourcesTag) SetKey(v string) *DescribeDcdnTagResourcesResponseBodyTagResourcesTag {
	s.Key = &v
	return s
}

func (s *DescribeDcdnTagResourcesResponseBodyTagResourcesTag) SetValue(v string) *DescribeDcdnTagResourcesResponseBodyTagResourcesTag {
	s.Value = &v
	return s
}

type DescribeDcdnTagResourcesResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDcdnTagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDcdnTagResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnTagResourcesResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnTagResourcesResponse) SetHeaders(v map[string]*string) *DescribeDcdnTagResourcesResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnTagResourcesResponse) SetBody(v *DescribeDcdnTagResourcesResponseBody) *DescribeDcdnTagResourcesResponse {
	s.Body = v
	return s
}

type DescribeDcdnTopDomainsByFlowRequest struct {
	EndTime   *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Limit     *int64  `json:"Limit,omitempty" xml:"Limit,omitempty"`
	OwnerId   *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDcdnTopDomainsByFlowRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnTopDomainsByFlowRequest) GoString() string {
	return s.String()
}

func (s *DescribeDcdnTopDomainsByFlowRequest) SetEndTime(v string) *DescribeDcdnTopDomainsByFlowRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDcdnTopDomainsByFlowRequest) SetLimit(v int64) *DescribeDcdnTopDomainsByFlowRequest {
	s.Limit = &v
	return s
}

func (s *DescribeDcdnTopDomainsByFlowRequest) SetOwnerId(v int64) *DescribeDcdnTopDomainsByFlowRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDcdnTopDomainsByFlowRequest) SetStartTime(v string) *DescribeDcdnTopDomainsByFlowRequest {
	s.StartTime = &v
	return s
}

type DescribeDcdnTopDomainsByFlowResponseBody struct {
	DomainCount       *int64                                              `json:"DomainCount,omitempty" xml:"DomainCount,omitempty"`
	DomainOnlineCount *int64                                              `json:"DomainOnlineCount,omitempty" xml:"DomainOnlineCount,omitempty"`
	EndTime           *string                                             `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	RequestId         *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StartTime         *string                                             `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	TopDomains        *DescribeDcdnTopDomainsByFlowResponseBodyTopDomains `json:"TopDomains,omitempty" xml:"TopDomains,omitempty" type:"Struct"`
}

func (s DescribeDcdnTopDomainsByFlowResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnTopDomainsByFlowResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnTopDomainsByFlowResponseBody) SetDomainCount(v int64) *DescribeDcdnTopDomainsByFlowResponseBody {
	s.DomainCount = &v
	return s
}

func (s *DescribeDcdnTopDomainsByFlowResponseBody) SetDomainOnlineCount(v int64) *DescribeDcdnTopDomainsByFlowResponseBody {
	s.DomainOnlineCount = &v
	return s
}

func (s *DescribeDcdnTopDomainsByFlowResponseBody) SetEndTime(v string) *DescribeDcdnTopDomainsByFlowResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeDcdnTopDomainsByFlowResponseBody) SetRequestId(v string) *DescribeDcdnTopDomainsByFlowResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDcdnTopDomainsByFlowResponseBody) SetStartTime(v string) *DescribeDcdnTopDomainsByFlowResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeDcdnTopDomainsByFlowResponseBody) SetTopDomains(v *DescribeDcdnTopDomainsByFlowResponseBodyTopDomains) *DescribeDcdnTopDomainsByFlowResponseBody {
	s.TopDomains = v
	return s
}

type DescribeDcdnTopDomainsByFlowResponseBodyTopDomains struct {
	TopDomain []*DescribeDcdnTopDomainsByFlowResponseBodyTopDomainsTopDomain `json:"TopDomain,omitempty" xml:"TopDomain,omitempty" type:"Repeated"`
}

func (s DescribeDcdnTopDomainsByFlowResponseBodyTopDomains) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnTopDomainsByFlowResponseBodyTopDomains) GoString() string {
	return s.String()
}

func (s *DescribeDcdnTopDomainsByFlowResponseBodyTopDomains) SetTopDomain(v []*DescribeDcdnTopDomainsByFlowResponseBodyTopDomainsTopDomain) *DescribeDcdnTopDomainsByFlowResponseBodyTopDomains {
	s.TopDomain = v
	return s
}

type DescribeDcdnTopDomainsByFlowResponseBodyTopDomainsTopDomain struct {
	DomainName     *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	MaxBps         *int64  `json:"MaxBps,omitempty" xml:"MaxBps,omitempty"`
	MaxBpsTime     *string `json:"MaxBpsTime,omitempty" xml:"MaxBpsTime,omitempty"`
	Rank           *int64  `json:"Rank,omitempty" xml:"Rank,omitempty"`
	TotalAccess    *int64  `json:"TotalAccess,omitempty" xml:"TotalAccess,omitempty"`
	TotalTraffic   *string `json:"TotalTraffic,omitempty" xml:"TotalTraffic,omitempty"`
	TrafficPercent *string `json:"TrafficPercent,omitempty" xml:"TrafficPercent,omitempty"`
}

func (s DescribeDcdnTopDomainsByFlowResponseBodyTopDomainsTopDomain) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnTopDomainsByFlowResponseBodyTopDomainsTopDomain) GoString() string {
	return s.String()
}

func (s *DescribeDcdnTopDomainsByFlowResponseBodyTopDomainsTopDomain) SetDomainName(v string) *DescribeDcdnTopDomainsByFlowResponseBodyTopDomainsTopDomain {
	s.DomainName = &v
	return s
}

func (s *DescribeDcdnTopDomainsByFlowResponseBodyTopDomainsTopDomain) SetMaxBps(v int64) *DescribeDcdnTopDomainsByFlowResponseBodyTopDomainsTopDomain {
	s.MaxBps = &v
	return s
}

func (s *DescribeDcdnTopDomainsByFlowResponseBodyTopDomainsTopDomain) SetMaxBpsTime(v string) *DescribeDcdnTopDomainsByFlowResponseBodyTopDomainsTopDomain {
	s.MaxBpsTime = &v
	return s
}

func (s *DescribeDcdnTopDomainsByFlowResponseBodyTopDomainsTopDomain) SetRank(v int64) *DescribeDcdnTopDomainsByFlowResponseBodyTopDomainsTopDomain {
	s.Rank = &v
	return s
}

func (s *DescribeDcdnTopDomainsByFlowResponseBodyTopDomainsTopDomain) SetTotalAccess(v int64) *DescribeDcdnTopDomainsByFlowResponseBodyTopDomainsTopDomain {
	s.TotalAccess = &v
	return s
}

func (s *DescribeDcdnTopDomainsByFlowResponseBodyTopDomainsTopDomain) SetTotalTraffic(v string) *DescribeDcdnTopDomainsByFlowResponseBodyTopDomainsTopDomain {
	s.TotalTraffic = &v
	return s
}

func (s *DescribeDcdnTopDomainsByFlowResponseBodyTopDomainsTopDomain) SetTrafficPercent(v string) *DescribeDcdnTopDomainsByFlowResponseBodyTopDomainsTopDomain {
	s.TrafficPercent = &v
	return s
}

type DescribeDcdnTopDomainsByFlowResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDcdnTopDomainsByFlowResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDcdnTopDomainsByFlowResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnTopDomainsByFlowResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnTopDomainsByFlowResponse) SetHeaders(v map[string]*string) *DescribeDcdnTopDomainsByFlowResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnTopDomainsByFlowResponse) SetBody(v *DescribeDcdnTopDomainsByFlowResponseBody) *DescribeDcdnTopDomainsByFlowResponse {
	s.Body = v
	return s
}

type DescribeDcdnUserBillHistoryRequest struct {
	EndTime   *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	OwnerId   *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDcdnUserBillHistoryRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnUserBillHistoryRequest) GoString() string {
	return s.String()
}

func (s *DescribeDcdnUserBillHistoryRequest) SetEndTime(v string) *DescribeDcdnUserBillHistoryRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDcdnUserBillHistoryRequest) SetOwnerId(v int64) *DescribeDcdnUserBillHistoryRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDcdnUserBillHistoryRequest) SetStartTime(v string) *DescribeDcdnUserBillHistoryRequest {
	s.StartTime = &v
	return s
}

type DescribeDcdnUserBillHistoryResponseBody struct {
	BillHistoryData *DescribeDcdnUserBillHistoryResponseBodyBillHistoryData `json:"BillHistoryData,omitempty" xml:"BillHistoryData,omitempty" type:"Struct"`
	RequestId       *string                                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDcdnUserBillHistoryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnUserBillHistoryResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnUserBillHistoryResponseBody) SetBillHistoryData(v *DescribeDcdnUserBillHistoryResponseBodyBillHistoryData) *DescribeDcdnUserBillHistoryResponseBody {
	s.BillHistoryData = v
	return s
}

func (s *DescribeDcdnUserBillHistoryResponseBody) SetRequestId(v string) *DescribeDcdnUserBillHistoryResponseBody {
	s.RequestId = &v
	return s
}

type DescribeDcdnUserBillHistoryResponseBodyBillHistoryData struct {
	BillHistoryDataItem []*DescribeDcdnUserBillHistoryResponseBodyBillHistoryDataBillHistoryDataItem `json:"BillHistoryDataItem,omitempty" xml:"BillHistoryDataItem,omitempty" type:"Repeated"`
}

func (s DescribeDcdnUserBillHistoryResponseBodyBillHistoryData) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnUserBillHistoryResponseBodyBillHistoryData) GoString() string {
	return s.String()
}

func (s *DescribeDcdnUserBillHistoryResponseBodyBillHistoryData) SetBillHistoryDataItem(v []*DescribeDcdnUserBillHistoryResponseBodyBillHistoryDataBillHistoryDataItem) *DescribeDcdnUserBillHistoryResponseBodyBillHistoryData {
	s.BillHistoryDataItem = v
	return s
}

type DescribeDcdnUserBillHistoryResponseBodyBillHistoryDataBillHistoryDataItem struct {
	BillTime    *string                                                                               `json:"BillTime,omitempty" xml:"BillTime,omitempty"`
	BillType    *string                                                                               `json:"BillType,omitempty" xml:"BillType,omitempty"`
	BillingData *DescribeDcdnUserBillHistoryResponseBodyBillHistoryDataBillHistoryDataItemBillingData `json:"BillingData,omitempty" xml:"BillingData,omitempty" type:"Struct"`
	Dimension   *string                                                                               `json:"Dimension,omitempty" xml:"Dimension,omitempty"`
}

func (s DescribeDcdnUserBillHistoryResponseBodyBillHistoryDataBillHistoryDataItem) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnUserBillHistoryResponseBodyBillHistoryDataBillHistoryDataItem) GoString() string {
	return s.String()
}

func (s *DescribeDcdnUserBillHistoryResponseBodyBillHistoryDataBillHistoryDataItem) SetBillTime(v string) *DescribeDcdnUserBillHistoryResponseBodyBillHistoryDataBillHistoryDataItem {
	s.BillTime = &v
	return s
}

func (s *DescribeDcdnUserBillHistoryResponseBodyBillHistoryDataBillHistoryDataItem) SetBillType(v string) *DescribeDcdnUserBillHistoryResponseBodyBillHistoryDataBillHistoryDataItem {
	s.BillType = &v
	return s
}

func (s *DescribeDcdnUserBillHistoryResponseBodyBillHistoryDataBillHistoryDataItem) SetBillingData(v *DescribeDcdnUserBillHistoryResponseBodyBillHistoryDataBillHistoryDataItemBillingData) *DescribeDcdnUserBillHistoryResponseBodyBillHistoryDataBillHistoryDataItem {
	s.BillingData = v
	return s
}

func (s *DescribeDcdnUserBillHistoryResponseBodyBillHistoryDataBillHistoryDataItem) SetDimension(v string) *DescribeDcdnUserBillHistoryResponseBodyBillHistoryDataBillHistoryDataItem {
	s.Dimension = &v
	return s
}

type DescribeDcdnUserBillHistoryResponseBodyBillHistoryDataBillHistoryDataItemBillingData struct {
	BillingDataItem []*DescribeDcdnUserBillHistoryResponseBodyBillHistoryDataBillHistoryDataItemBillingDataBillingDataItem `json:"BillingDataItem,omitempty" xml:"BillingDataItem,omitempty" type:"Repeated"`
}

func (s DescribeDcdnUserBillHistoryResponseBodyBillHistoryDataBillHistoryDataItemBillingData) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnUserBillHistoryResponseBodyBillHistoryDataBillHistoryDataItemBillingData) GoString() string {
	return s.String()
}

func (s *DescribeDcdnUserBillHistoryResponseBodyBillHistoryDataBillHistoryDataItemBillingData) SetBillingDataItem(v []*DescribeDcdnUserBillHistoryResponseBodyBillHistoryDataBillHistoryDataItemBillingDataBillingDataItem) *DescribeDcdnUserBillHistoryResponseBodyBillHistoryDataBillHistoryDataItemBillingData {
	s.BillingDataItem = v
	return s
}

type DescribeDcdnUserBillHistoryResponseBodyBillHistoryDataBillHistoryDataItemBillingDataBillingDataItem struct {
	Bandwidth  *float32 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	CdnRegion  *string  `json:"CdnRegion,omitempty" xml:"CdnRegion,omitempty"`
	ChargeType *string  `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	Count      *float32 `json:"Count,omitempty" xml:"Count,omitempty"`
	Flow       *float32 `json:"Flow,omitempty" xml:"Flow,omitempty"`
}

func (s DescribeDcdnUserBillHistoryResponseBodyBillHistoryDataBillHistoryDataItemBillingDataBillingDataItem) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnUserBillHistoryResponseBodyBillHistoryDataBillHistoryDataItemBillingDataBillingDataItem) GoString() string {
	return s.String()
}

func (s *DescribeDcdnUserBillHistoryResponseBodyBillHistoryDataBillHistoryDataItemBillingDataBillingDataItem) SetBandwidth(v float32) *DescribeDcdnUserBillHistoryResponseBodyBillHistoryDataBillHistoryDataItemBillingDataBillingDataItem {
	s.Bandwidth = &v
	return s
}

func (s *DescribeDcdnUserBillHistoryResponseBodyBillHistoryDataBillHistoryDataItemBillingDataBillingDataItem) SetCdnRegion(v string) *DescribeDcdnUserBillHistoryResponseBodyBillHistoryDataBillHistoryDataItemBillingDataBillingDataItem {
	s.CdnRegion = &v
	return s
}

func (s *DescribeDcdnUserBillHistoryResponseBodyBillHistoryDataBillHistoryDataItemBillingDataBillingDataItem) SetChargeType(v string) *DescribeDcdnUserBillHistoryResponseBodyBillHistoryDataBillHistoryDataItemBillingDataBillingDataItem {
	s.ChargeType = &v
	return s
}

func (s *DescribeDcdnUserBillHistoryResponseBodyBillHistoryDataBillHistoryDataItemBillingDataBillingDataItem) SetCount(v float32) *DescribeDcdnUserBillHistoryResponseBodyBillHistoryDataBillHistoryDataItemBillingDataBillingDataItem {
	s.Count = &v
	return s
}

func (s *DescribeDcdnUserBillHistoryResponseBodyBillHistoryDataBillHistoryDataItemBillingDataBillingDataItem) SetFlow(v float32) *DescribeDcdnUserBillHistoryResponseBodyBillHistoryDataBillHistoryDataItemBillingDataBillingDataItem {
	s.Flow = &v
	return s
}

type DescribeDcdnUserBillHistoryResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDcdnUserBillHistoryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDcdnUserBillHistoryResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnUserBillHistoryResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnUserBillHistoryResponse) SetHeaders(v map[string]*string) *DescribeDcdnUserBillHistoryResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnUserBillHistoryResponse) SetBody(v *DescribeDcdnUserBillHistoryResponseBody) *DescribeDcdnUserBillHistoryResponse {
	s.Body = v
	return s
}

type DescribeDcdnUserBillTypeRequest struct {
	EndTime   *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	OwnerId   *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDcdnUserBillTypeRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnUserBillTypeRequest) GoString() string {
	return s.String()
}

func (s *DescribeDcdnUserBillTypeRequest) SetEndTime(v string) *DescribeDcdnUserBillTypeRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDcdnUserBillTypeRequest) SetOwnerId(v int64) *DescribeDcdnUserBillTypeRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDcdnUserBillTypeRequest) SetStartTime(v string) *DescribeDcdnUserBillTypeRequest {
	s.StartTime = &v
	return s
}

type DescribeDcdnUserBillTypeResponseBody struct {
	BillTypeData *DescribeDcdnUserBillTypeResponseBodyBillTypeData `json:"BillTypeData,omitempty" xml:"BillTypeData,omitempty" type:"Struct"`
	RequestId    *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDcdnUserBillTypeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnUserBillTypeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnUserBillTypeResponseBody) SetBillTypeData(v *DescribeDcdnUserBillTypeResponseBodyBillTypeData) *DescribeDcdnUserBillTypeResponseBody {
	s.BillTypeData = v
	return s
}

func (s *DescribeDcdnUserBillTypeResponseBody) SetRequestId(v string) *DescribeDcdnUserBillTypeResponseBody {
	s.RequestId = &v
	return s
}

type DescribeDcdnUserBillTypeResponseBodyBillTypeData struct {
	BillTypeDataItem []*DescribeDcdnUserBillTypeResponseBodyBillTypeDataBillTypeDataItem `json:"BillTypeDataItem,omitempty" xml:"BillTypeDataItem,omitempty" type:"Repeated"`
}

func (s DescribeDcdnUserBillTypeResponseBodyBillTypeData) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnUserBillTypeResponseBodyBillTypeData) GoString() string {
	return s.String()
}

func (s *DescribeDcdnUserBillTypeResponseBodyBillTypeData) SetBillTypeDataItem(v []*DescribeDcdnUserBillTypeResponseBodyBillTypeDataBillTypeDataItem) *DescribeDcdnUserBillTypeResponseBodyBillTypeData {
	s.BillTypeDataItem = v
	return s
}

type DescribeDcdnUserBillTypeResponseBodyBillTypeDataBillTypeDataItem struct {
	BillType     *string `json:"BillType,omitempty" xml:"BillType,omitempty"`
	BillingCycle *string `json:"BillingCycle,omitempty" xml:"BillingCycle,omitempty"`
	Dimension    *string `json:"Dimension,omitempty" xml:"Dimension,omitempty"`
	EndTime      *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Product      *string `json:"Product,omitempty" xml:"Product,omitempty"`
	StartTime    *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDcdnUserBillTypeResponseBodyBillTypeDataBillTypeDataItem) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnUserBillTypeResponseBodyBillTypeDataBillTypeDataItem) GoString() string {
	return s.String()
}

func (s *DescribeDcdnUserBillTypeResponseBodyBillTypeDataBillTypeDataItem) SetBillType(v string) *DescribeDcdnUserBillTypeResponseBodyBillTypeDataBillTypeDataItem {
	s.BillType = &v
	return s
}

func (s *DescribeDcdnUserBillTypeResponseBodyBillTypeDataBillTypeDataItem) SetBillingCycle(v string) *DescribeDcdnUserBillTypeResponseBodyBillTypeDataBillTypeDataItem {
	s.BillingCycle = &v
	return s
}

func (s *DescribeDcdnUserBillTypeResponseBodyBillTypeDataBillTypeDataItem) SetDimension(v string) *DescribeDcdnUserBillTypeResponseBodyBillTypeDataBillTypeDataItem {
	s.Dimension = &v
	return s
}

func (s *DescribeDcdnUserBillTypeResponseBodyBillTypeDataBillTypeDataItem) SetEndTime(v string) *DescribeDcdnUserBillTypeResponseBodyBillTypeDataBillTypeDataItem {
	s.EndTime = &v
	return s
}

func (s *DescribeDcdnUserBillTypeResponseBodyBillTypeDataBillTypeDataItem) SetProduct(v string) *DescribeDcdnUserBillTypeResponseBodyBillTypeDataBillTypeDataItem {
	s.Product = &v
	return s
}

func (s *DescribeDcdnUserBillTypeResponseBodyBillTypeDataBillTypeDataItem) SetStartTime(v string) *DescribeDcdnUserBillTypeResponseBodyBillTypeDataBillTypeDataItem {
	s.StartTime = &v
	return s
}

type DescribeDcdnUserBillTypeResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDcdnUserBillTypeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDcdnUserBillTypeResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnUserBillTypeResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnUserBillTypeResponse) SetHeaders(v map[string]*string) *DescribeDcdnUserBillTypeResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnUserBillTypeResponse) SetBody(v *DescribeDcdnUserBillTypeResponseBody) *DescribeDcdnUserBillTypeResponse {
	s.Body = v
	return s
}

type DescribeDcdnUserDomainsRequest struct {
	ChangeEndTime    *string                              `json:"ChangeEndTime,omitempty" xml:"ChangeEndTime,omitempty"`
	ChangeStartTime  *string                              `json:"ChangeStartTime,omitempty" xml:"ChangeStartTime,omitempty"`
	CheckDomainShow  *bool                                `json:"CheckDomainShow,omitempty" xml:"CheckDomainShow,omitempty"`
	Coverage         *string                              `json:"Coverage,omitempty" xml:"Coverage,omitempty"`
	DomainName       *string                              `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	DomainSearchType *string                              `json:"DomainSearchType,omitempty" xml:"DomainSearchType,omitempty"`
	DomainStatus     *string                              `json:"DomainStatus,omitempty" xml:"DomainStatus,omitempty"`
	OwnerId          *int64                               `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNumber       *int32                               `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize         *int32                               `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceGroupId  *string                              `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	SecurityToken    *string                              `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	Tag              []*DescribeDcdnUserDomainsRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeDcdnUserDomainsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnUserDomainsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDcdnUserDomainsRequest) SetChangeEndTime(v string) *DescribeDcdnUserDomainsRequest {
	s.ChangeEndTime = &v
	return s
}

func (s *DescribeDcdnUserDomainsRequest) SetChangeStartTime(v string) *DescribeDcdnUserDomainsRequest {
	s.ChangeStartTime = &v
	return s
}

func (s *DescribeDcdnUserDomainsRequest) SetCheckDomainShow(v bool) *DescribeDcdnUserDomainsRequest {
	s.CheckDomainShow = &v
	return s
}

func (s *DescribeDcdnUserDomainsRequest) SetCoverage(v string) *DescribeDcdnUserDomainsRequest {
	s.Coverage = &v
	return s
}

func (s *DescribeDcdnUserDomainsRequest) SetDomainName(v string) *DescribeDcdnUserDomainsRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDcdnUserDomainsRequest) SetDomainSearchType(v string) *DescribeDcdnUserDomainsRequest {
	s.DomainSearchType = &v
	return s
}

func (s *DescribeDcdnUserDomainsRequest) SetDomainStatus(v string) *DescribeDcdnUserDomainsRequest {
	s.DomainStatus = &v
	return s
}

func (s *DescribeDcdnUserDomainsRequest) SetOwnerId(v int64) *DescribeDcdnUserDomainsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDcdnUserDomainsRequest) SetPageNumber(v int32) *DescribeDcdnUserDomainsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDcdnUserDomainsRequest) SetPageSize(v int32) *DescribeDcdnUserDomainsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDcdnUserDomainsRequest) SetResourceGroupId(v string) *DescribeDcdnUserDomainsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDcdnUserDomainsRequest) SetSecurityToken(v string) *DescribeDcdnUserDomainsRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeDcdnUserDomainsRequest) SetTag(v []*DescribeDcdnUserDomainsRequestTag) *DescribeDcdnUserDomainsRequest {
	s.Tag = v
	return s
}

type DescribeDcdnUserDomainsRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDcdnUserDomainsRequestTag) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnUserDomainsRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeDcdnUserDomainsRequestTag) SetKey(v string) *DescribeDcdnUserDomainsRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeDcdnUserDomainsRequestTag) SetValue(v string) *DescribeDcdnUserDomainsRequestTag {
	s.Value = &v
	return s
}

type DescribeDcdnUserDomainsResponseBody struct {
	Domains    *DescribeDcdnUserDomainsResponseBodyDomains `json:"Domains,omitempty" xml:"Domains,omitempty" type:"Struct"`
	PageNumber *int64                                      `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int64                                      `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int64                                      `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDcdnUserDomainsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnUserDomainsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnUserDomainsResponseBody) SetDomains(v *DescribeDcdnUserDomainsResponseBodyDomains) *DescribeDcdnUserDomainsResponseBody {
	s.Domains = v
	return s
}

func (s *DescribeDcdnUserDomainsResponseBody) SetPageNumber(v int64) *DescribeDcdnUserDomainsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeDcdnUserDomainsResponseBody) SetPageSize(v int64) *DescribeDcdnUserDomainsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeDcdnUserDomainsResponseBody) SetRequestId(v string) *DescribeDcdnUserDomainsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDcdnUserDomainsResponseBody) SetTotalCount(v int64) *DescribeDcdnUserDomainsResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeDcdnUserDomainsResponseBodyDomains struct {
	PageData []*DescribeDcdnUserDomainsResponseBodyDomainsPageData `json:"PageData,omitempty" xml:"PageData,omitempty" type:"Repeated"`
}

func (s DescribeDcdnUserDomainsResponseBodyDomains) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnUserDomainsResponseBodyDomains) GoString() string {
	return s.String()
}

func (s *DescribeDcdnUserDomainsResponseBodyDomains) SetPageData(v []*DescribeDcdnUserDomainsResponseBodyDomainsPageData) *DescribeDcdnUserDomainsResponseBodyDomains {
	s.PageData = v
	return s
}

type DescribeDcdnUserDomainsResponseBodyDomainsPageData struct {
	Cname           *string                                                    `json:"Cname,omitempty" xml:"Cname,omitempty"`
	Description     *string                                                    `json:"Description,omitempty" xml:"Description,omitempty"`
	DomainName      *string                                                    `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	DomainStatus    *string                                                    `json:"DomainStatus,omitempty" xml:"DomainStatus,omitempty"`
	GmtCreated      *string                                                    `json:"GmtCreated,omitempty" xml:"GmtCreated,omitempty"`
	GmtModified     *string                                                    `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	ResourceGroupId *string                                                    `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	SSLProtocol     *string                                                    `json:"SSLProtocol,omitempty" xml:"SSLProtocol,omitempty"`
	Sandbox         *string                                                    `json:"Sandbox,omitempty" xml:"Sandbox,omitempty"`
	Sources         *DescribeDcdnUserDomainsResponseBodyDomainsPageDataSources `json:"Sources,omitempty" xml:"Sources,omitempty" type:"Struct"`
}

func (s DescribeDcdnUserDomainsResponseBodyDomainsPageData) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnUserDomainsResponseBodyDomainsPageData) GoString() string {
	return s.String()
}

func (s *DescribeDcdnUserDomainsResponseBodyDomainsPageData) SetCname(v string) *DescribeDcdnUserDomainsResponseBodyDomainsPageData {
	s.Cname = &v
	return s
}

func (s *DescribeDcdnUserDomainsResponseBodyDomainsPageData) SetDescription(v string) *DescribeDcdnUserDomainsResponseBodyDomainsPageData {
	s.Description = &v
	return s
}

func (s *DescribeDcdnUserDomainsResponseBodyDomainsPageData) SetDomainName(v string) *DescribeDcdnUserDomainsResponseBodyDomainsPageData {
	s.DomainName = &v
	return s
}

func (s *DescribeDcdnUserDomainsResponseBodyDomainsPageData) SetDomainStatus(v string) *DescribeDcdnUserDomainsResponseBodyDomainsPageData {
	s.DomainStatus = &v
	return s
}

func (s *DescribeDcdnUserDomainsResponseBodyDomainsPageData) SetGmtCreated(v string) *DescribeDcdnUserDomainsResponseBodyDomainsPageData {
	s.GmtCreated = &v
	return s
}

func (s *DescribeDcdnUserDomainsResponseBodyDomainsPageData) SetGmtModified(v string) *DescribeDcdnUserDomainsResponseBodyDomainsPageData {
	s.GmtModified = &v
	return s
}

func (s *DescribeDcdnUserDomainsResponseBodyDomainsPageData) SetResourceGroupId(v string) *DescribeDcdnUserDomainsResponseBodyDomainsPageData {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDcdnUserDomainsResponseBodyDomainsPageData) SetSSLProtocol(v string) *DescribeDcdnUserDomainsResponseBodyDomainsPageData {
	s.SSLProtocol = &v
	return s
}

func (s *DescribeDcdnUserDomainsResponseBodyDomainsPageData) SetSandbox(v string) *DescribeDcdnUserDomainsResponseBodyDomainsPageData {
	s.Sandbox = &v
	return s
}

func (s *DescribeDcdnUserDomainsResponseBodyDomainsPageData) SetSources(v *DescribeDcdnUserDomainsResponseBodyDomainsPageDataSources) *DescribeDcdnUserDomainsResponseBodyDomainsPageData {
	s.Sources = v
	return s
}

type DescribeDcdnUserDomainsResponseBodyDomainsPageDataSources struct {
	Source []*DescribeDcdnUserDomainsResponseBodyDomainsPageDataSourcesSource `json:"Source,omitempty" xml:"Source,omitempty" type:"Repeated"`
}

func (s DescribeDcdnUserDomainsResponseBodyDomainsPageDataSources) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnUserDomainsResponseBodyDomainsPageDataSources) GoString() string {
	return s.String()
}

func (s *DescribeDcdnUserDomainsResponseBodyDomainsPageDataSources) SetSource(v []*DescribeDcdnUserDomainsResponseBodyDomainsPageDataSourcesSource) *DescribeDcdnUserDomainsResponseBodyDomainsPageDataSources {
	s.Source = v
	return s
}

type DescribeDcdnUserDomainsResponseBodyDomainsPageDataSourcesSource struct {
	Content  *string `json:"Content,omitempty" xml:"Content,omitempty"`
	Port     *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
	Priority *string `json:"Priority,omitempty" xml:"Priority,omitempty"`
	Type     *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Weight   *string `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s DescribeDcdnUserDomainsResponseBodyDomainsPageDataSourcesSource) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnUserDomainsResponseBodyDomainsPageDataSourcesSource) GoString() string {
	return s.String()
}

func (s *DescribeDcdnUserDomainsResponseBodyDomainsPageDataSourcesSource) SetContent(v string) *DescribeDcdnUserDomainsResponseBodyDomainsPageDataSourcesSource {
	s.Content = &v
	return s
}

func (s *DescribeDcdnUserDomainsResponseBodyDomainsPageDataSourcesSource) SetPort(v int32) *DescribeDcdnUserDomainsResponseBodyDomainsPageDataSourcesSource {
	s.Port = &v
	return s
}

func (s *DescribeDcdnUserDomainsResponseBodyDomainsPageDataSourcesSource) SetPriority(v string) *DescribeDcdnUserDomainsResponseBodyDomainsPageDataSourcesSource {
	s.Priority = &v
	return s
}

func (s *DescribeDcdnUserDomainsResponseBodyDomainsPageDataSourcesSource) SetType(v string) *DescribeDcdnUserDomainsResponseBodyDomainsPageDataSourcesSource {
	s.Type = &v
	return s
}

func (s *DescribeDcdnUserDomainsResponseBodyDomainsPageDataSourcesSource) SetWeight(v string) *DescribeDcdnUserDomainsResponseBodyDomainsPageDataSourcesSource {
	s.Weight = &v
	return s
}

type DescribeDcdnUserDomainsResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDcdnUserDomainsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDcdnUserDomainsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnUserDomainsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnUserDomainsResponse) SetHeaders(v map[string]*string) *DescribeDcdnUserDomainsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnUserDomainsResponse) SetBody(v *DescribeDcdnUserDomainsResponseBody) *DescribeDcdnUserDomainsResponse {
	s.Body = v
	return s
}

type DescribeDcdnUserDomainsByFuncRequest struct {
	DomainName      *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	FuncFilter      *string `json:"FuncFilter,omitempty" xml:"FuncFilter,omitempty"`
	FuncId          *int32  `json:"FuncId,omitempty" xml:"FuncId,omitempty"`
	OwnerId         *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNumber      *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize        *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s DescribeDcdnUserDomainsByFuncRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnUserDomainsByFuncRequest) GoString() string {
	return s.String()
}

func (s *DescribeDcdnUserDomainsByFuncRequest) SetDomainName(v string) *DescribeDcdnUserDomainsByFuncRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDcdnUserDomainsByFuncRequest) SetFuncFilter(v string) *DescribeDcdnUserDomainsByFuncRequest {
	s.FuncFilter = &v
	return s
}

func (s *DescribeDcdnUserDomainsByFuncRequest) SetFuncId(v int32) *DescribeDcdnUserDomainsByFuncRequest {
	s.FuncId = &v
	return s
}

func (s *DescribeDcdnUserDomainsByFuncRequest) SetOwnerId(v int64) *DescribeDcdnUserDomainsByFuncRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDcdnUserDomainsByFuncRequest) SetPageNumber(v int32) *DescribeDcdnUserDomainsByFuncRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDcdnUserDomainsByFuncRequest) SetPageSize(v int32) *DescribeDcdnUserDomainsByFuncRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDcdnUserDomainsByFuncRequest) SetResourceGroupId(v string) *DescribeDcdnUserDomainsByFuncRequest {
	s.ResourceGroupId = &v
	return s
}

type DescribeDcdnUserDomainsByFuncResponseBody struct {
	Domains    *DescribeDcdnUserDomainsByFuncResponseBodyDomains `json:"Domains,omitempty" xml:"Domains,omitempty" type:"Struct"`
	PageNumber *int64                                            `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int64                                            `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int64                                            `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDcdnUserDomainsByFuncResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnUserDomainsByFuncResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnUserDomainsByFuncResponseBody) SetDomains(v *DescribeDcdnUserDomainsByFuncResponseBodyDomains) *DescribeDcdnUserDomainsByFuncResponseBody {
	s.Domains = v
	return s
}

func (s *DescribeDcdnUserDomainsByFuncResponseBody) SetPageNumber(v int64) *DescribeDcdnUserDomainsByFuncResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeDcdnUserDomainsByFuncResponseBody) SetPageSize(v int64) *DescribeDcdnUserDomainsByFuncResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeDcdnUserDomainsByFuncResponseBody) SetRequestId(v string) *DescribeDcdnUserDomainsByFuncResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDcdnUserDomainsByFuncResponseBody) SetTotalCount(v int64) *DescribeDcdnUserDomainsByFuncResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeDcdnUserDomainsByFuncResponseBodyDomains struct {
	PageData []*DescribeDcdnUserDomainsByFuncResponseBodyDomainsPageData `json:"PageData,omitempty" xml:"PageData,omitempty" type:"Repeated"`
}

func (s DescribeDcdnUserDomainsByFuncResponseBodyDomains) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnUserDomainsByFuncResponseBodyDomains) GoString() string {
	return s.String()
}

func (s *DescribeDcdnUserDomainsByFuncResponseBodyDomains) SetPageData(v []*DescribeDcdnUserDomainsByFuncResponseBodyDomainsPageData) *DescribeDcdnUserDomainsByFuncResponseBodyDomains {
	s.PageData = v
	return s
}

type DescribeDcdnUserDomainsByFuncResponseBodyDomainsPageData struct {
	Cname           *string                                                          `json:"Cname,omitempty" xml:"Cname,omitempty"`
	Description     *string                                                          `json:"Description,omitempty" xml:"Description,omitempty"`
	DomainName      *string                                                          `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	DomainStatus    *string                                                          `json:"DomainStatus,omitempty" xml:"DomainStatus,omitempty"`
	GmtCreated      *string                                                          `json:"GmtCreated,omitempty" xml:"GmtCreated,omitempty"`
	GmtModified     *string                                                          `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	ResourceGroupId *string                                                          `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Sandbox         *string                                                          `json:"Sandbox,omitempty" xml:"Sandbox,omitempty"`
	Sources         *DescribeDcdnUserDomainsByFuncResponseBodyDomainsPageDataSources `json:"Sources,omitempty" xml:"Sources,omitempty" type:"Struct"`
	SslProtocol     *string                                                          `json:"SslProtocol,omitempty" xml:"SslProtocol,omitempty"`
}

func (s DescribeDcdnUserDomainsByFuncResponseBodyDomainsPageData) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnUserDomainsByFuncResponseBodyDomainsPageData) GoString() string {
	return s.String()
}

func (s *DescribeDcdnUserDomainsByFuncResponseBodyDomainsPageData) SetCname(v string) *DescribeDcdnUserDomainsByFuncResponseBodyDomainsPageData {
	s.Cname = &v
	return s
}

func (s *DescribeDcdnUserDomainsByFuncResponseBodyDomainsPageData) SetDescription(v string) *DescribeDcdnUserDomainsByFuncResponseBodyDomainsPageData {
	s.Description = &v
	return s
}

func (s *DescribeDcdnUserDomainsByFuncResponseBodyDomainsPageData) SetDomainName(v string) *DescribeDcdnUserDomainsByFuncResponseBodyDomainsPageData {
	s.DomainName = &v
	return s
}

func (s *DescribeDcdnUserDomainsByFuncResponseBodyDomainsPageData) SetDomainStatus(v string) *DescribeDcdnUserDomainsByFuncResponseBodyDomainsPageData {
	s.DomainStatus = &v
	return s
}

func (s *DescribeDcdnUserDomainsByFuncResponseBodyDomainsPageData) SetGmtCreated(v string) *DescribeDcdnUserDomainsByFuncResponseBodyDomainsPageData {
	s.GmtCreated = &v
	return s
}

func (s *DescribeDcdnUserDomainsByFuncResponseBodyDomainsPageData) SetGmtModified(v string) *DescribeDcdnUserDomainsByFuncResponseBodyDomainsPageData {
	s.GmtModified = &v
	return s
}

func (s *DescribeDcdnUserDomainsByFuncResponseBodyDomainsPageData) SetResourceGroupId(v string) *DescribeDcdnUserDomainsByFuncResponseBodyDomainsPageData {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDcdnUserDomainsByFuncResponseBodyDomainsPageData) SetSandbox(v string) *DescribeDcdnUserDomainsByFuncResponseBodyDomainsPageData {
	s.Sandbox = &v
	return s
}

func (s *DescribeDcdnUserDomainsByFuncResponseBodyDomainsPageData) SetSources(v *DescribeDcdnUserDomainsByFuncResponseBodyDomainsPageDataSources) *DescribeDcdnUserDomainsByFuncResponseBodyDomainsPageData {
	s.Sources = v
	return s
}

func (s *DescribeDcdnUserDomainsByFuncResponseBodyDomainsPageData) SetSslProtocol(v string) *DescribeDcdnUserDomainsByFuncResponseBodyDomainsPageData {
	s.SslProtocol = &v
	return s
}

type DescribeDcdnUserDomainsByFuncResponseBodyDomainsPageDataSources struct {
	Source []*DescribeDcdnUserDomainsByFuncResponseBodyDomainsPageDataSourcesSource `json:"Source,omitempty" xml:"Source,omitempty" type:"Repeated"`
}

func (s DescribeDcdnUserDomainsByFuncResponseBodyDomainsPageDataSources) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnUserDomainsByFuncResponseBodyDomainsPageDataSources) GoString() string {
	return s.String()
}

func (s *DescribeDcdnUserDomainsByFuncResponseBodyDomainsPageDataSources) SetSource(v []*DescribeDcdnUserDomainsByFuncResponseBodyDomainsPageDataSourcesSource) *DescribeDcdnUserDomainsByFuncResponseBodyDomainsPageDataSources {
	s.Source = v
	return s
}

type DescribeDcdnUserDomainsByFuncResponseBodyDomainsPageDataSourcesSource struct {
	Content  *string `json:"Content,omitempty" xml:"Content,omitempty"`
	Port     *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
	Priority *string `json:"Priority,omitempty" xml:"Priority,omitempty"`
	Type     *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Weight   *string `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s DescribeDcdnUserDomainsByFuncResponseBodyDomainsPageDataSourcesSource) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnUserDomainsByFuncResponseBodyDomainsPageDataSourcesSource) GoString() string {
	return s.String()
}

func (s *DescribeDcdnUserDomainsByFuncResponseBodyDomainsPageDataSourcesSource) SetContent(v string) *DescribeDcdnUserDomainsByFuncResponseBodyDomainsPageDataSourcesSource {
	s.Content = &v
	return s
}

func (s *DescribeDcdnUserDomainsByFuncResponseBodyDomainsPageDataSourcesSource) SetPort(v int32) *DescribeDcdnUserDomainsByFuncResponseBodyDomainsPageDataSourcesSource {
	s.Port = &v
	return s
}

func (s *DescribeDcdnUserDomainsByFuncResponseBodyDomainsPageDataSourcesSource) SetPriority(v string) *DescribeDcdnUserDomainsByFuncResponseBodyDomainsPageDataSourcesSource {
	s.Priority = &v
	return s
}

func (s *DescribeDcdnUserDomainsByFuncResponseBodyDomainsPageDataSourcesSource) SetType(v string) *DescribeDcdnUserDomainsByFuncResponseBodyDomainsPageDataSourcesSource {
	s.Type = &v
	return s
}

func (s *DescribeDcdnUserDomainsByFuncResponseBodyDomainsPageDataSourcesSource) SetWeight(v string) *DescribeDcdnUserDomainsByFuncResponseBodyDomainsPageDataSourcesSource {
	s.Weight = &v
	return s
}

type DescribeDcdnUserDomainsByFuncResponse struct {
	Headers map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDcdnUserDomainsByFuncResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDcdnUserDomainsByFuncResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnUserDomainsByFuncResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnUserDomainsByFuncResponse) SetHeaders(v map[string]*string) *DescribeDcdnUserDomainsByFuncResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnUserDomainsByFuncResponse) SetBody(v *DescribeDcdnUserDomainsByFuncResponseBody) *DescribeDcdnUserDomainsByFuncResponse {
	s.Body = v
	return s
}

type DescribeDcdnUserQuotaRequest struct {
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeDcdnUserQuotaRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnUserQuotaRequest) GoString() string {
	return s.String()
}

func (s *DescribeDcdnUserQuotaRequest) SetOwnerId(v int64) *DescribeDcdnUserQuotaRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDcdnUserQuotaRequest) SetSecurityToken(v string) *DescribeDcdnUserQuotaRequest {
	s.SecurityToken = &v
	return s
}

type DescribeDcdnUserQuotaResponseBody struct {
	BlockQuota       *int32  `json:"BlockQuota,omitempty" xml:"BlockQuota,omitempty"`
	BlockRemain      *int32  `json:"BlockRemain,omitempty" xml:"BlockRemain,omitempty"`
	DomainQuota      *int32  `json:"DomainQuota,omitempty" xml:"DomainQuota,omitempty"`
	PreloadQuota     *int32  `json:"PreloadQuota,omitempty" xml:"PreloadQuota,omitempty"`
	PreloadRemain    *int32  `json:"PreloadRemain,omitempty" xml:"PreloadRemain,omitempty"`
	RefreshDirQuota  *int32  `json:"RefreshDirQuota,omitempty" xml:"RefreshDirQuota,omitempty"`
	RefreshDirRemain *int32  `json:"RefreshDirRemain,omitempty" xml:"RefreshDirRemain,omitempty"`
	RefreshUrlQuota  *int32  `json:"RefreshUrlQuota,omitempty" xml:"RefreshUrlQuota,omitempty"`
	RefreshUrlRemain *int32  `json:"RefreshUrlRemain,omitempty" xml:"RefreshUrlRemain,omitempty"`
	RequestId        *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDcdnUserQuotaResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnUserQuotaResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnUserQuotaResponseBody) SetBlockQuota(v int32) *DescribeDcdnUserQuotaResponseBody {
	s.BlockQuota = &v
	return s
}

func (s *DescribeDcdnUserQuotaResponseBody) SetBlockRemain(v int32) *DescribeDcdnUserQuotaResponseBody {
	s.BlockRemain = &v
	return s
}

func (s *DescribeDcdnUserQuotaResponseBody) SetDomainQuota(v int32) *DescribeDcdnUserQuotaResponseBody {
	s.DomainQuota = &v
	return s
}

func (s *DescribeDcdnUserQuotaResponseBody) SetPreloadQuota(v int32) *DescribeDcdnUserQuotaResponseBody {
	s.PreloadQuota = &v
	return s
}

func (s *DescribeDcdnUserQuotaResponseBody) SetPreloadRemain(v int32) *DescribeDcdnUserQuotaResponseBody {
	s.PreloadRemain = &v
	return s
}

func (s *DescribeDcdnUserQuotaResponseBody) SetRefreshDirQuota(v int32) *DescribeDcdnUserQuotaResponseBody {
	s.RefreshDirQuota = &v
	return s
}

func (s *DescribeDcdnUserQuotaResponseBody) SetRefreshDirRemain(v int32) *DescribeDcdnUserQuotaResponseBody {
	s.RefreshDirRemain = &v
	return s
}

func (s *DescribeDcdnUserQuotaResponseBody) SetRefreshUrlQuota(v int32) *DescribeDcdnUserQuotaResponseBody {
	s.RefreshUrlQuota = &v
	return s
}

func (s *DescribeDcdnUserQuotaResponseBody) SetRefreshUrlRemain(v int32) *DescribeDcdnUserQuotaResponseBody {
	s.RefreshUrlRemain = &v
	return s
}

func (s *DescribeDcdnUserQuotaResponseBody) SetRequestId(v string) *DescribeDcdnUserQuotaResponseBody {
	s.RequestId = &v
	return s
}

type DescribeDcdnUserQuotaResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDcdnUserQuotaResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDcdnUserQuotaResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnUserQuotaResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnUserQuotaResponse) SetHeaders(v map[string]*string) *DescribeDcdnUserQuotaResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnUserQuotaResponse) SetBody(v *DescribeDcdnUserQuotaResponseBody) *DescribeDcdnUserQuotaResponse {
	s.Body = v
	return s
}

type DescribeDcdnUserRealTimeDeliveryFieldRequest struct {
	BusinessType *string `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s DescribeDcdnUserRealTimeDeliveryFieldRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnUserRealTimeDeliveryFieldRequest) GoString() string {
	return s.String()
}

func (s *DescribeDcdnUserRealTimeDeliveryFieldRequest) SetBusinessType(v string) *DescribeDcdnUserRealTimeDeliveryFieldRequest {
	s.BusinessType = &v
	return s
}

func (s *DescribeDcdnUserRealTimeDeliveryFieldRequest) SetOwnerId(v int64) *DescribeDcdnUserRealTimeDeliveryFieldRequest {
	s.OwnerId = &v
	return s
}

type DescribeDcdnUserRealTimeDeliveryFieldResponseBody struct {
	Content   *DescribeDcdnUserRealTimeDeliveryFieldResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	RequestId *string                                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDcdnUserRealTimeDeliveryFieldResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnUserRealTimeDeliveryFieldResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnUserRealTimeDeliveryFieldResponseBody) SetContent(v *DescribeDcdnUserRealTimeDeliveryFieldResponseBodyContent) *DescribeDcdnUserRealTimeDeliveryFieldResponseBody {
	s.Content = v
	return s
}

func (s *DescribeDcdnUserRealTimeDeliveryFieldResponseBody) SetRequestId(v string) *DescribeDcdnUserRealTimeDeliveryFieldResponseBody {
	s.RequestId = &v
	return s
}

type DescribeDcdnUserRealTimeDeliveryFieldResponseBodyContent struct {
	Fields []*DescribeDcdnUserRealTimeDeliveryFieldResponseBodyContentFields `json:"Fields,omitempty" xml:"Fields,omitempty" type:"Repeated"`
}

func (s DescribeDcdnUserRealTimeDeliveryFieldResponseBodyContent) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnUserRealTimeDeliveryFieldResponseBodyContent) GoString() string {
	return s.String()
}

func (s *DescribeDcdnUserRealTimeDeliveryFieldResponseBodyContent) SetFields(v []*DescribeDcdnUserRealTimeDeliveryFieldResponseBodyContentFields) *DescribeDcdnUserRealTimeDeliveryFieldResponseBodyContent {
	s.Fields = v
	return s
}

type DescribeDcdnUserRealTimeDeliveryFieldResponseBodyContentFields struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	FieldName   *string `json:"FieldName,omitempty" xml:"FieldName,omitempty"`
	Selected    *bool   `json:"Selected,omitempty" xml:"Selected,omitempty"`
}

func (s DescribeDcdnUserRealTimeDeliveryFieldResponseBodyContentFields) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnUserRealTimeDeliveryFieldResponseBodyContentFields) GoString() string {
	return s.String()
}

func (s *DescribeDcdnUserRealTimeDeliveryFieldResponseBodyContentFields) SetDescription(v string) *DescribeDcdnUserRealTimeDeliveryFieldResponseBodyContentFields {
	s.Description = &v
	return s
}

func (s *DescribeDcdnUserRealTimeDeliveryFieldResponseBodyContentFields) SetFieldName(v string) *DescribeDcdnUserRealTimeDeliveryFieldResponseBodyContentFields {
	s.FieldName = &v
	return s
}

func (s *DescribeDcdnUserRealTimeDeliveryFieldResponseBodyContentFields) SetSelected(v bool) *DescribeDcdnUserRealTimeDeliveryFieldResponseBodyContentFields {
	s.Selected = &v
	return s
}

type DescribeDcdnUserRealTimeDeliveryFieldResponse struct {
	Headers map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDcdnUserRealTimeDeliveryFieldResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDcdnUserRealTimeDeliveryFieldResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnUserRealTimeDeliveryFieldResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnUserRealTimeDeliveryFieldResponse) SetHeaders(v map[string]*string) *DescribeDcdnUserRealTimeDeliveryFieldResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnUserRealTimeDeliveryFieldResponse) SetBody(v *DescribeDcdnUserRealTimeDeliveryFieldResponseBody) *DescribeDcdnUserRealTimeDeliveryFieldResponse {
	s.Body = v
	return s
}

type DescribeDcdnUserResourcePackageRequest struct {
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	Status        *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeDcdnUserResourcePackageRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnUserResourcePackageRequest) GoString() string {
	return s.String()
}

func (s *DescribeDcdnUserResourcePackageRequest) SetOwnerId(v int64) *DescribeDcdnUserResourcePackageRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDcdnUserResourcePackageRequest) SetSecurityToken(v string) *DescribeDcdnUserResourcePackageRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeDcdnUserResourcePackageRequest) SetStatus(v string) *DescribeDcdnUserResourcePackageRequest {
	s.Status = &v
	return s
}

type DescribeDcdnUserResourcePackageResponseBody struct {
	RequestId            *string                                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResourcePackageInfos *DescribeDcdnUserResourcePackageResponseBodyResourcePackageInfos `json:"ResourcePackageInfos,omitempty" xml:"ResourcePackageInfos,omitempty" type:"Struct"`
}

func (s DescribeDcdnUserResourcePackageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnUserResourcePackageResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnUserResourcePackageResponseBody) SetRequestId(v string) *DescribeDcdnUserResourcePackageResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDcdnUserResourcePackageResponseBody) SetResourcePackageInfos(v *DescribeDcdnUserResourcePackageResponseBodyResourcePackageInfos) *DescribeDcdnUserResourcePackageResponseBody {
	s.ResourcePackageInfos = v
	return s
}

type DescribeDcdnUserResourcePackageResponseBodyResourcePackageInfos struct {
	ResourcePackageInfo []*DescribeDcdnUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo `json:"ResourcePackageInfo,omitempty" xml:"ResourcePackageInfo,omitempty" type:"Repeated"`
}

func (s DescribeDcdnUserResourcePackageResponseBodyResourcePackageInfos) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnUserResourcePackageResponseBodyResourcePackageInfos) GoString() string {
	return s.String()
}

func (s *DescribeDcdnUserResourcePackageResponseBodyResourcePackageInfos) SetResourcePackageInfo(v []*DescribeDcdnUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo) *DescribeDcdnUserResourcePackageResponseBodyResourcePackageInfos {
	s.ResourcePackageInfo = v
	return s
}

type DescribeDcdnUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo struct {
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	CurrCapacity  *string `json:"CurrCapacity,omitempty" xml:"CurrCapacity,omitempty"`
	DisplayName   *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	EndTime       *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	InitCapacity  *string `json:"InitCapacity,omitempty" xml:"InitCapacity,omitempty"`
	InstanceId    *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	StartTime     *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Status        *string `json:"Status,omitempty" xml:"Status,omitempty"`
	TemplateName  *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
}

func (s DescribeDcdnUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo) GoString() string {
	return s.String()
}

func (s *DescribeDcdnUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo) SetCommodityCode(v string) *DescribeDcdnUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo {
	s.CommodityCode = &v
	return s
}

func (s *DescribeDcdnUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo) SetCurrCapacity(v string) *DescribeDcdnUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo {
	s.CurrCapacity = &v
	return s
}

func (s *DescribeDcdnUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo) SetDisplayName(v string) *DescribeDcdnUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo {
	s.DisplayName = &v
	return s
}

func (s *DescribeDcdnUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo) SetEndTime(v string) *DescribeDcdnUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo {
	s.EndTime = &v
	return s
}

func (s *DescribeDcdnUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo) SetInitCapacity(v string) *DescribeDcdnUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo {
	s.InitCapacity = &v
	return s
}

func (s *DescribeDcdnUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo) SetInstanceId(v string) *DescribeDcdnUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo {
	s.InstanceId = &v
	return s
}

func (s *DescribeDcdnUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo) SetStartTime(v string) *DescribeDcdnUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo {
	s.StartTime = &v
	return s
}

func (s *DescribeDcdnUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo) SetStatus(v string) *DescribeDcdnUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo {
	s.Status = &v
	return s
}

func (s *DescribeDcdnUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo) SetTemplateName(v string) *DescribeDcdnUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo {
	s.TemplateName = &v
	return s
}

type DescribeDcdnUserResourcePackageResponse struct {
	Headers map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDcdnUserResourcePackageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDcdnUserResourcePackageResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnUserResourcePackageResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnUserResourcePackageResponse) SetHeaders(v map[string]*string) *DescribeDcdnUserResourcePackageResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnUserResourcePackageResponse) SetBody(v *DescribeDcdnUserResourcePackageResponseBody) *DescribeDcdnUserResourcePackageResponse {
	s.Body = v
	return s
}

type DescribeDcdnUserSecDropRequest struct {
	Data    *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Metric  *string `json:"Metric,omitempty" xml:"Metric,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecFunc *string `json:"SecFunc,omitempty" xml:"SecFunc,omitempty"`
}

func (s DescribeDcdnUserSecDropRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnUserSecDropRequest) GoString() string {
	return s.String()
}

func (s *DescribeDcdnUserSecDropRequest) SetData(v string) *DescribeDcdnUserSecDropRequest {
	s.Data = &v
	return s
}

func (s *DescribeDcdnUserSecDropRequest) SetMetric(v string) *DescribeDcdnUserSecDropRequest {
	s.Metric = &v
	return s
}

func (s *DescribeDcdnUserSecDropRequest) SetOwnerId(v int64) *DescribeDcdnUserSecDropRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDcdnUserSecDropRequest) SetSecFunc(v string) *DescribeDcdnUserSecDropRequest {
	s.SecFunc = &v
	return s
}

type DescribeDcdnUserSecDropResponseBody struct {
	Drops     *int32  `json:"Drops,omitempty" xml:"Drops,omitempty"`
	Msg       *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	UuidStr   *string `json:"UuidStr,omitempty" xml:"UuidStr,omitempty"`
}

func (s DescribeDcdnUserSecDropResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnUserSecDropResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnUserSecDropResponseBody) SetDrops(v int32) *DescribeDcdnUserSecDropResponseBody {
	s.Drops = &v
	return s
}

func (s *DescribeDcdnUserSecDropResponseBody) SetMsg(v string) *DescribeDcdnUserSecDropResponseBody {
	s.Msg = &v
	return s
}

func (s *DescribeDcdnUserSecDropResponseBody) SetRequestId(v string) *DescribeDcdnUserSecDropResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDcdnUserSecDropResponseBody) SetUuidStr(v string) *DescribeDcdnUserSecDropResponseBody {
	s.UuidStr = &v
	return s
}

type DescribeDcdnUserSecDropResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDcdnUserSecDropResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDcdnUserSecDropResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnUserSecDropResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnUserSecDropResponse) SetHeaders(v map[string]*string) *DescribeDcdnUserSecDropResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnUserSecDropResponse) SetBody(v *DescribeDcdnUserSecDropResponseBody) *DescribeDcdnUserSecDropResponse {
	s.Body = v
	return s
}

type DescribeDcdnUserSecDropByMinuteRequest struct {
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EndTime    *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Lang       *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	Object     *string `json:"Object,omitempty" xml:"Object,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNumber *int64  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RuleName   *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	SecFunc    *string `json:"SecFunc,omitempty" xml:"SecFunc,omitempty"`
	StartTime  *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDcdnUserSecDropByMinuteRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnUserSecDropByMinuteRequest) GoString() string {
	return s.String()
}

func (s *DescribeDcdnUserSecDropByMinuteRequest) SetDomainName(v string) *DescribeDcdnUserSecDropByMinuteRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDcdnUserSecDropByMinuteRequest) SetEndTime(v string) *DescribeDcdnUserSecDropByMinuteRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDcdnUserSecDropByMinuteRequest) SetLang(v string) *DescribeDcdnUserSecDropByMinuteRequest {
	s.Lang = &v
	return s
}

func (s *DescribeDcdnUserSecDropByMinuteRequest) SetObject(v string) *DescribeDcdnUserSecDropByMinuteRequest {
	s.Object = &v
	return s
}

func (s *DescribeDcdnUserSecDropByMinuteRequest) SetOwnerId(v int64) *DescribeDcdnUserSecDropByMinuteRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDcdnUserSecDropByMinuteRequest) SetPageNumber(v int64) *DescribeDcdnUserSecDropByMinuteRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDcdnUserSecDropByMinuteRequest) SetPageSize(v int64) *DescribeDcdnUserSecDropByMinuteRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDcdnUserSecDropByMinuteRequest) SetRuleName(v string) *DescribeDcdnUserSecDropByMinuteRequest {
	s.RuleName = &v
	return s
}

func (s *DescribeDcdnUserSecDropByMinuteRequest) SetSecFunc(v string) *DescribeDcdnUserSecDropByMinuteRequest {
	s.SecFunc = &v
	return s
}

func (s *DescribeDcdnUserSecDropByMinuteRequest) SetStartTime(v string) *DescribeDcdnUserSecDropByMinuteRequest {
	s.StartTime = &v
	return s
}

type DescribeDcdnUserSecDropByMinuteResponseBody struct {
	Description *string                                            `json:"Description,omitempty" xml:"Description,omitempty"`
	Len         *int32                                             `json:"Len,omitempty" xml:"Len,omitempty"`
	PageNumber  *int32                                             `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize    *int32                                             `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId   *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Rows        []*DescribeDcdnUserSecDropByMinuteResponseBodyRows `json:"Rows,omitempty" xml:"Rows,omitempty" type:"Repeated"`
	TotalCount  *int32                                             `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDcdnUserSecDropByMinuteResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnUserSecDropByMinuteResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnUserSecDropByMinuteResponseBody) SetDescription(v string) *DescribeDcdnUserSecDropByMinuteResponseBody {
	s.Description = &v
	return s
}

func (s *DescribeDcdnUserSecDropByMinuteResponseBody) SetLen(v int32) *DescribeDcdnUserSecDropByMinuteResponseBody {
	s.Len = &v
	return s
}

func (s *DescribeDcdnUserSecDropByMinuteResponseBody) SetPageNumber(v int32) *DescribeDcdnUserSecDropByMinuteResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeDcdnUserSecDropByMinuteResponseBody) SetPageSize(v int32) *DescribeDcdnUserSecDropByMinuteResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeDcdnUserSecDropByMinuteResponseBody) SetRequestId(v string) *DescribeDcdnUserSecDropByMinuteResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDcdnUserSecDropByMinuteResponseBody) SetRows(v []*DescribeDcdnUserSecDropByMinuteResponseBodyRows) *DescribeDcdnUserSecDropByMinuteResponseBody {
	s.Rows = v
	return s
}

func (s *DescribeDcdnUserSecDropByMinuteResponseBody) SetTotalCount(v int32) *DescribeDcdnUserSecDropByMinuteResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeDcdnUserSecDropByMinuteResponseBodyRows struct {
	Domain   *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	Drops    *int32  `json:"Drops,omitempty" xml:"Drops,omitempty"`
	Object   *string `json:"Object,omitempty" xml:"Object,omitempty"`
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	SecFunc  *string `json:"SecFunc,omitempty" xml:"SecFunc,omitempty"`
	TmStr    *string `json:"TmStr,omitempty" xml:"TmStr,omitempty"`
}

func (s DescribeDcdnUserSecDropByMinuteResponseBodyRows) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnUserSecDropByMinuteResponseBodyRows) GoString() string {
	return s.String()
}

func (s *DescribeDcdnUserSecDropByMinuteResponseBodyRows) SetDomain(v string) *DescribeDcdnUserSecDropByMinuteResponseBodyRows {
	s.Domain = &v
	return s
}

func (s *DescribeDcdnUserSecDropByMinuteResponseBodyRows) SetDrops(v int32) *DescribeDcdnUserSecDropByMinuteResponseBodyRows {
	s.Drops = &v
	return s
}

func (s *DescribeDcdnUserSecDropByMinuteResponseBodyRows) SetObject(v string) *DescribeDcdnUserSecDropByMinuteResponseBodyRows {
	s.Object = &v
	return s
}

func (s *DescribeDcdnUserSecDropByMinuteResponseBodyRows) SetRuleName(v string) *DescribeDcdnUserSecDropByMinuteResponseBodyRows {
	s.RuleName = &v
	return s
}

func (s *DescribeDcdnUserSecDropByMinuteResponseBodyRows) SetSecFunc(v string) *DescribeDcdnUserSecDropByMinuteResponseBodyRows {
	s.SecFunc = &v
	return s
}

func (s *DescribeDcdnUserSecDropByMinuteResponseBodyRows) SetTmStr(v string) *DescribeDcdnUserSecDropByMinuteResponseBodyRows {
	s.TmStr = &v
	return s
}

type DescribeDcdnUserSecDropByMinuteResponse struct {
	Headers map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDcdnUserSecDropByMinuteResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDcdnUserSecDropByMinuteResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnUserSecDropByMinuteResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnUserSecDropByMinuteResponse) SetHeaders(v map[string]*string) *DescribeDcdnUserSecDropByMinuteResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnUserSecDropByMinuteResponse) SetBody(v *DescribeDcdnUserSecDropByMinuteResponseBody) *DescribeDcdnUserSecDropByMinuteResponse {
	s.Body = v
	return s
}

type DescribeDcdnUserTagsRequest struct {
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s DescribeDcdnUserTagsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnUserTagsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDcdnUserTagsRequest) SetOwnerId(v int64) *DescribeDcdnUserTagsRequest {
	s.OwnerId = &v
	return s
}

type DescribeDcdnUserTagsResponseBody struct {
	RequestId *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Tags      []*DescribeDcdnUserTagsResponseBodyTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s DescribeDcdnUserTagsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnUserTagsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnUserTagsResponseBody) SetRequestId(v string) *DescribeDcdnUserTagsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDcdnUserTagsResponseBody) SetTags(v []*DescribeDcdnUserTagsResponseBodyTags) *DescribeDcdnUserTagsResponseBody {
	s.Tags = v
	return s
}

type DescribeDcdnUserTagsResponseBodyTags struct {
	Key   *string   `json:"Key,omitempty" xml:"Key,omitempty"`
	Value []*string `json:"Value,omitempty" xml:"Value,omitempty" type:"Repeated"`
}

func (s DescribeDcdnUserTagsResponseBodyTags) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnUserTagsResponseBodyTags) GoString() string {
	return s.String()
}

func (s *DescribeDcdnUserTagsResponseBodyTags) SetKey(v string) *DescribeDcdnUserTagsResponseBodyTags {
	s.Key = &v
	return s
}

func (s *DescribeDcdnUserTagsResponseBodyTags) SetValue(v []*string) *DescribeDcdnUserTagsResponseBodyTags {
	s.Value = v
	return s
}

type DescribeDcdnUserTagsResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDcdnUserTagsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDcdnUserTagsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnUserTagsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnUserTagsResponse) SetHeaders(v map[string]*string) *DescribeDcdnUserTagsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnUserTagsResponse) SetBody(v *DescribeDcdnUserTagsResponseBody) *DescribeDcdnUserTagsResponse {
	s.Body = v
	return s
}

type DescribeDcdnVerifyContentRequest struct {
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s DescribeDcdnVerifyContentRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnVerifyContentRequest) GoString() string {
	return s.String()
}

func (s *DescribeDcdnVerifyContentRequest) SetDomainName(v string) *DescribeDcdnVerifyContentRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDcdnVerifyContentRequest) SetOwnerId(v int64) *DescribeDcdnVerifyContentRequest {
	s.OwnerId = &v
	return s
}

type DescribeDcdnVerifyContentResponseBody struct {
	Content   *string `json:"Content,omitempty" xml:"Content,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDcdnVerifyContentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnVerifyContentResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnVerifyContentResponseBody) SetContent(v string) *DescribeDcdnVerifyContentResponseBody {
	s.Content = &v
	return s
}

func (s *DescribeDcdnVerifyContentResponseBody) SetRequestId(v string) *DescribeDcdnVerifyContentResponseBody {
	s.RequestId = &v
	return s
}

type DescribeDcdnVerifyContentResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDcdnVerifyContentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDcdnVerifyContentResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnVerifyContentResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnVerifyContentResponse) SetHeaders(v map[string]*string) *DescribeDcdnVerifyContentResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnVerifyContentResponse) SetBody(v *DescribeDcdnVerifyContentResponseBody) *DescribeDcdnVerifyContentResponse {
	s.Body = v
	return s
}

type DescribeDcdnWafDomainRequest struct {
	DomainName      *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	OwnerId         *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s DescribeDcdnWafDomainRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnWafDomainRequest) GoString() string {
	return s.String()
}

func (s *DescribeDcdnWafDomainRequest) SetDomainName(v string) *DescribeDcdnWafDomainRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDcdnWafDomainRequest) SetOwnerId(v int64) *DescribeDcdnWafDomainRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDcdnWafDomainRequest) SetRegionId(v string) *DescribeDcdnWafDomainRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDcdnWafDomainRequest) SetResourceGroupId(v string) *DescribeDcdnWafDomainRequest {
	s.ResourceGroupId = &v
	return s
}

type DescribeDcdnWafDomainResponseBody struct {
	OutPutDomains []*DescribeDcdnWafDomainResponseBodyOutPutDomains `json:"OutPutDomains,omitempty" xml:"OutPutDomains,omitempty" type:"Repeated"`
	RequestId     *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount    *int32                                            `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDcdnWafDomainResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnWafDomainResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnWafDomainResponseBody) SetOutPutDomains(v []*DescribeDcdnWafDomainResponseBodyOutPutDomains) *DescribeDcdnWafDomainResponseBody {
	s.OutPutDomains = v
	return s
}

func (s *DescribeDcdnWafDomainResponseBody) SetRequestId(v string) *DescribeDcdnWafDomainResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDcdnWafDomainResponseBody) SetTotalCount(v int32) *DescribeDcdnWafDomainResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeDcdnWafDomainResponseBodyOutPutDomains struct {
	AclStatus *int32  `json:"AclStatus,omitempty" xml:"AclStatus,omitempty"`
	CcStatus  *int32  `json:"CcStatus,omitempty" xml:"CcStatus,omitempty"`
	Domain    *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	Status    *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	WafStatus *int32  `json:"WafStatus,omitempty" xml:"WafStatus,omitempty"`
}

func (s DescribeDcdnWafDomainResponseBodyOutPutDomains) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnWafDomainResponseBodyOutPutDomains) GoString() string {
	return s.String()
}

func (s *DescribeDcdnWafDomainResponseBodyOutPutDomains) SetAclStatus(v int32) *DescribeDcdnWafDomainResponseBodyOutPutDomains {
	s.AclStatus = &v
	return s
}

func (s *DescribeDcdnWafDomainResponseBodyOutPutDomains) SetCcStatus(v int32) *DescribeDcdnWafDomainResponseBodyOutPutDomains {
	s.CcStatus = &v
	return s
}

func (s *DescribeDcdnWafDomainResponseBodyOutPutDomains) SetDomain(v string) *DescribeDcdnWafDomainResponseBodyOutPutDomains {
	s.Domain = &v
	return s
}

func (s *DescribeDcdnWafDomainResponseBodyOutPutDomains) SetStatus(v int32) *DescribeDcdnWafDomainResponseBodyOutPutDomains {
	s.Status = &v
	return s
}

func (s *DescribeDcdnWafDomainResponseBodyOutPutDomains) SetWafStatus(v int32) *DescribeDcdnWafDomainResponseBodyOutPutDomains {
	s.WafStatus = &v
	return s
}

type DescribeDcdnWafDomainResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDcdnWafDomainResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDcdnWafDomainResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnWafDomainResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnWafDomainResponse) SetHeaders(v map[string]*string) *DescribeDcdnWafDomainResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnWafDomainResponse) SetBody(v *DescribeDcdnWafDomainResponseBody) *DescribeDcdnWafDomainResponse {
	s.Body = v
	return s
}

type DescribeDcdnsecServiceRequest struct {
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeDcdnsecServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnsecServiceRequest) GoString() string {
	return s.String()
}

func (s *DescribeDcdnsecServiceRequest) SetOwnerId(v int64) *DescribeDcdnsecServiceRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDcdnsecServiceRequest) SetSecurityToken(v string) *DescribeDcdnsecServiceRequest {
	s.SecurityToken = &v
	return s
}

type DescribeDcdnsecServiceResponseBody struct {
	ChangingAffectTime *string                                           `json:"ChangingAffectTime,omitempty" xml:"ChangingAffectTime,omitempty"`
	ChangingChargeType *string                                           `json:"ChangingChargeType,omitempty" xml:"ChangingChargeType,omitempty"`
	DomainNum          *string                                           `json:"DomainNum,omitempty" xml:"DomainNum,omitempty"`
	EndTime            *string                                           `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	FlowType           *string                                           `json:"FlowType,omitempty" xml:"FlowType,omitempty"`
	InstanceId         *string                                           `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InternetChargeType *string                                           `json:"InternetChargeType,omitempty" xml:"InternetChargeType,omitempty"`
	OperationLocks     *DescribeDcdnsecServiceResponseBodyOperationLocks `json:"OperationLocks,omitempty" xml:"OperationLocks,omitempty" type:"Struct"`
	RequestId          *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RequestType        *string                                           `json:"RequestType,omitempty" xml:"RequestType,omitempty"`
	StartTime          *string                                           `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Version            *string                                           `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s DescribeDcdnsecServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnsecServiceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnsecServiceResponseBody) SetChangingAffectTime(v string) *DescribeDcdnsecServiceResponseBody {
	s.ChangingAffectTime = &v
	return s
}

func (s *DescribeDcdnsecServiceResponseBody) SetChangingChargeType(v string) *DescribeDcdnsecServiceResponseBody {
	s.ChangingChargeType = &v
	return s
}

func (s *DescribeDcdnsecServiceResponseBody) SetDomainNum(v string) *DescribeDcdnsecServiceResponseBody {
	s.DomainNum = &v
	return s
}

func (s *DescribeDcdnsecServiceResponseBody) SetEndTime(v string) *DescribeDcdnsecServiceResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeDcdnsecServiceResponseBody) SetFlowType(v string) *DescribeDcdnsecServiceResponseBody {
	s.FlowType = &v
	return s
}

func (s *DescribeDcdnsecServiceResponseBody) SetInstanceId(v string) *DescribeDcdnsecServiceResponseBody {
	s.InstanceId = &v
	return s
}

func (s *DescribeDcdnsecServiceResponseBody) SetInternetChargeType(v string) *DescribeDcdnsecServiceResponseBody {
	s.InternetChargeType = &v
	return s
}

func (s *DescribeDcdnsecServiceResponseBody) SetOperationLocks(v *DescribeDcdnsecServiceResponseBodyOperationLocks) *DescribeDcdnsecServiceResponseBody {
	s.OperationLocks = v
	return s
}

func (s *DescribeDcdnsecServiceResponseBody) SetRequestId(v string) *DescribeDcdnsecServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDcdnsecServiceResponseBody) SetRequestType(v string) *DescribeDcdnsecServiceResponseBody {
	s.RequestType = &v
	return s
}

func (s *DescribeDcdnsecServiceResponseBody) SetStartTime(v string) *DescribeDcdnsecServiceResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeDcdnsecServiceResponseBody) SetVersion(v string) *DescribeDcdnsecServiceResponseBody {
	s.Version = &v
	return s
}

type DescribeDcdnsecServiceResponseBodyOperationLocks struct {
	LockReason []*DescribeDcdnsecServiceResponseBodyOperationLocksLockReason `json:"LockReason,omitempty" xml:"LockReason,omitempty" type:"Repeated"`
}

func (s DescribeDcdnsecServiceResponseBodyOperationLocks) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnsecServiceResponseBodyOperationLocks) GoString() string {
	return s.String()
}

func (s *DescribeDcdnsecServiceResponseBodyOperationLocks) SetLockReason(v []*DescribeDcdnsecServiceResponseBodyOperationLocksLockReason) *DescribeDcdnsecServiceResponseBodyOperationLocks {
	s.LockReason = v
	return s
}

type DescribeDcdnsecServiceResponseBodyOperationLocksLockReason struct {
	LockReason *string `json:"LockReason,omitempty" xml:"LockReason,omitempty"`
}

func (s DescribeDcdnsecServiceResponseBodyOperationLocksLockReason) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnsecServiceResponseBodyOperationLocksLockReason) GoString() string {
	return s.String()
}

func (s *DescribeDcdnsecServiceResponseBodyOperationLocksLockReason) SetLockReason(v string) *DescribeDcdnsecServiceResponseBodyOperationLocksLockReason {
	s.LockReason = &v
	return s
}

type DescribeDcdnsecServiceResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDcdnsecServiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDcdnsecServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDcdnsecServiceResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnsecServiceResponse) SetHeaders(v map[string]*string) *DescribeDcdnsecServiceResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnsecServiceResponse) SetBody(v *DescribeDcdnsecServiceResponseBody) *DescribeDcdnsecServiceResponse {
	s.Body = v
	return s
}

type DescribeRoutineRequest struct {
	Name    *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s DescribeRoutineRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRoutineRequest) GoString() string {
	return s.String()
}

func (s *DescribeRoutineRequest) SetName(v string) *DescribeRoutineRequest {
	s.Name = &v
	return s
}

func (s *DescribeRoutineRequest) SetOwnerId(v int64) *DescribeRoutineRequest {
	s.OwnerId = &v
	return s
}

type DescribeRoutineResponseBody struct {
	Content   map[string]interface{} `json:"Content,omitempty" xml:"Content,omitempty"`
	RequestId *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeRoutineResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRoutineResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRoutineResponseBody) SetContent(v map[string]interface{}) *DescribeRoutineResponseBody {
	s.Content = v
	return s
}

func (s *DescribeRoutineResponseBody) SetRequestId(v string) *DescribeRoutineResponseBody {
	s.RequestId = &v
	return s
}

type DescribeRoutineResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeRoutineResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeRoutineResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRoutineResponse) GoString() string {
	return s.String()
}

func (s *DescribeRoutineResponse) SetHeaders(v map[string]*string) *DescribeRoutineResponse {
	s.Headers = v
	return s
}

func (s *DescribeRoutineResponse) SetBody(v *DescribeRoutineResponseBody) *DescribeRoutineResponse {
	s.Body = v
	return s
}

type DescribeRoutineCanaryEnvsRequest struct {
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s DescribeRoutineCanaryEnvsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRoutineCanaryEnvsRequest) GoString() string {
	return s.String()
}

func (s *DescribeRoutineCanaryEnvsRequest) SetOwnerId(v int64) *DescribeRoutineCanaryEnvsRequest {
	s.OwnerId = &v
	return s
}

type DescribeRoutineCanaryEnvsResponseBody struct {
	Content   map[string]interface{} `json:"Content,omitempty" xml:"Content,omitempty"`
	RequestId *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeRoutineCanaryEnvsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRoutineCanaryEnvsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRoutineCanaryEnvsResponseBody) SetContent(v map[string]interface{}) *DescribeRoutineCanaryEnvsResponseBody {
	s.Content = v
	return s
}

func (s *DescribeRoutineCanaryEnvsResponseBody) SetRequestId(v string) *DescribeRoutineCanaryEnvsResponseBody {
	s.RequestId = &v
	return s
}

type DescribeRoutineCanaryEnvsResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeRoutineCanaryEnvsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeRoutineCanaryEnvsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRoutineCanaryEnvsResponse) GoString() string {
	return s.String()
}

func (s *DescribeRoutineCanaryEnvsResponse) SetHeaders(v map[string]*string) *DescribeRoutineCanaryEnvsResponse {
	s.Headers = v
	return s
}

func (s *DescribeRoutineCanaryEnvsResponse) SetBody(v *DescribeRoutineCanaryEnvsResponseBody) *DescribeRoutineCanaryEnvsResponse {
	s.Body = v
	return s
}

type DescribeRoutineCodeRevisionRequest struct {
	Name               *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerId            *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SelectCodeRevision *string `json:"SelectCodeRevision,omitempty" xml:"SelectCodeRevision,omitempty"`
}

func (s DescribeRoutineCodeRevisionRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRoutineCodeRevisionRequest) GoString() string {
	return s.String()
}

func (s *DescribeRoutineCodeRevisionRequest) SetName(v string) *DescribeRoutineCodeRevisionRequest {
	s.Name = &v
	return s
}

func (s *DescribeRoutineCodeRevisionRequest) SetOwnerId(v int64) *DescribeRoutineCodeRevisionRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeRoutineCodeRevisionRequest) SetSelectCodeRevision(v string) *DescribeRoutineCodeRevisionRequest {
	s.SelectCodeRevision = &v
	return s
}

type DescribeRoutineCodeRevisionResponseBody struct {
	Content   map[string]interface{} `json:"Content,omitempty" xml:"Content,omitempty"`
	RequestId *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeRoutineCodeRevisionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRoutineCodeRevisionResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRoutineCodeRevisionResponseBody) SetContent(v map[string]interface{}) *DescribeRoutineCodeRevisionResponseBody {
	s.Content = v
	return s
}

func (s *DescribeRoutineCodeRevisionResponseBody) SetRequestId(v string) *DescribeRoutineCodeRevisionResponseBody {
	s.RequestId = &v
	return s
}

type DescribeRoutineCodeRevisionResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeRoutineCodeRevisionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeRoutineCodeRevisionResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRoutineCodeRevisionResponse) GoString() string {
	return s.String()
}

func (s *DescribeRoutineCodeRevisionResponse) SetHeaders(v map[string]*string) *DescribeRoutineCodeRevisionResponse {
	s.Headers = v
	return s
}

func (s *DescribeRoutineCodeRevisionResponse) SetBody(v *DescribeRoutineCodeRevisionResponseBody) *DescribeRoutineCodeRevisionResponse {
	s.Body = v
	return s
}

type DescribeRoutineSpecRequest struct {
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s DescribeRoutineSpecRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRoutineSpecRequest) GoString() string {
	return s.String()
}

func (s *DescribeRoutineSpecRequest) SetOwnerId(v int64) *DescribeRoutineSpecRequest {
	s.OwnerId = &v
	return s
}

type DescribeRoutineSpecResponseBody struct {
	Content   map[string]interface{} `json:"Content,omitempty" xml:"Content,omitempty"`
	RequestId *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeRoutineSpecResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRoutineSpecResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRoutineSpecResponseBody) SetContent(v map[string]interface{}) *DescribeRoutineSpecResponseBody {
	s.Content = v
	return s
}

func (s *DescribeRoutineSpecResponseBody) SetRequestId(v string) *DescribeRoutineSpecResponseBody {
	s.RequestId = &v
	return s
}

type DescribeRoutineSpecResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeRoutineSpecResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeRoutineSpecResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRoutineSpecResponse) GoString() string {
	return s.String()
}

func (s *DescribeRoutineSpecResponse) SetHeaders(v map[string]*string) *DescribeRoutineSpecResponse {
	s.Headers = v
	return s
}

func (s *DescribeRoutineSpecResponse) SetBody(v *DescribeRoutineSpecResponseBody) *DescribeRoutineSpecResponse {
	s.Body = v
	return s
}

type DescribeRoutineUserInfoRequest struct {
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s DescribeRoutineUserInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRoutineUserInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeRoutineUserInfoRequest) SetOwnerId(v int64) *DescribeRoutineUserInfoRequest {
	s.OwnerId = &v
	return s
}

type DescribeRoutineUserInfoResponseBody struct {
	Content   map[string]interface{} `json:"Content,omitempty" xml:"Content,omitempty"`
	RequestId *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeRoutineUserInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRoutineUserInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRoutineUserInfoResponseBody) SetContent(v map[string]interface{}) *DescribeRoutineUserInfoResponseBody {
	s.Content = v
	return s
}

func (s *DescribeRoutineUserInfoResponseBody) SetRequestId(v string) *DescribeRoutineUserInfoResponseBody {
	s.RequestId = &v
	return s
}

type DescribeRoutineUserInfoResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeRoutineUserInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeRoutineUserInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRoutineUserInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeRoutineUserInfoResponse) SetHeaders(v map[string]*string) *DescribeRoutineUserInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeRoutineUserInfoResponse) SetBody(v *DescribeRoutineUserInfoResponseBody) *DescribeRoutineUserInfoResponse {
	s.Body = v
	return s
}

type DescribeUserDcdnIpaStatusRequest struct {
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeUserDcdnIpaStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserDcdnIpaStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeUserDcdnIpaStatusRequest) SetOwnerId(v int64) *DescribeUserDcdnIpaStatusRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeUserDcdnIpaStatusRequest) SetSecurityToken(v string) *DescribeUserDcdnIpaStatusRequest {
	s.SecurityToken = &v
	return s
}

type DescribeUserDcdnIpaStatusResponseBody struct {
	Enabled       *bool   `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	InDebt        *bool   `json:"InDebt,omitempty" xml:"InDebt,omitempty"`
	InDebtOverdue *bool   `json:"InDebtOverdue,omitempty" xml:"InDebtOverdue,omitempty"`
	OnService     *bool   `json:"OnService,omitempty" xml:"OnService,omitempty"`
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeUserDcdnIpaStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserDcdnIpaStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUserDcdnIpaStatusResponseBody) SetEnabled(v bool) *DescribeUserDcdnIpaStatusResponseBody {
	s.Enabled = &v
	return s
}

func (s *DescribeUserDcdnIpaStatusResponseBody) SetInDebt(v bool) *DescribeUserDcdnIpaStatusResponseBody {
	s.InDebt = &v
	return s
}

func (s *DescribeUserDcdnIpaStatusResponseBody) SetInDebtOverdue(v bool) *DescribeUserDcdnIpaStatusResponseBody {
	s.InDebtOverdue = &v
	return s
}

func (s *DescribeUserDcdnIpaStatusResponseBody) SetOnService(v bool) *DescribeUserDcdnIpaStatusResponseBody {
	s.OnService = &v
	return s
}

func (s *DescribeUserDcdnIpaStatusResponseBody) SetRequestId(v string) *DescribeUserDcdnIpaStatusResponseBody {
	s.RequestId = &v
	return s
}

type DescribeUserDcdnIpaStatusResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeUserDcdnIpaStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeUserDcdnIpaStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserDcdnIpaStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeUserDcdnIpaStatusResponse) SetHeaders(v map[string]*string) *DescribeUserDcdnIpaStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeUserDcdnIpaStatusResponse) SetBody(v *DescribeUserDcdnIpaStatusResponseBody) *DescribeUserDcdnIpaStatusResponse {
	s.Body = v
	return s
}

type DescribeUserDcdnStatusRequest struct {
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeUserDcdnStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserDcdnStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeUserDcdnStatusRequest) SetOwnerId(v int64) *DescribeUserDcdnStatusRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeUserDcdnStatusRequest) SetSecurityToken(v string) *DescribeUserDcdnStatusRequest {
	s.SecurityToken = &v
	return s
}

type DescribeUserDcdnStatusResponseBody struct {
	Enabled       *bool   `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	InDebt        *bool   `json:"InDebt,omitempty" xml:"InDebt,omitempty"`
	InDebtOverdue *bool   `json:"InDebtOverdue,omitempty" xml:"InDebtOverdue,omitempty"`
	OnService     *bool   `json:"OnService,omitempty" xml:"OnService,omitempty"`
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeUserDcdnStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserDcdnStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUserDcdnStatusResponseBody) SetEnabled(v bool) *DescribeUserDcdnStatusResponseBody {
	s.Enabled = &v
	return s
}

func (s *DescribeUserDcdnStatusResponseBody) SetInDebt(v bool) *DescribeUserDcdnStatusResponseBody {
	s.InDebt = &v
	return s
}

func (s *DescribeUserDcdnStatusResponseBody) SetInDebtOverdue(v bool) *DescribeUserDcdnStatusResponseBody {
	s.InDebtOverdue = &v
	return s
}

func (s *DescribeUserDcdnStatusResponseBody) SetOnService(v bool) *DescribeUserDcdnStatusResponseBody {
	s.OnService = &v
	return s
}

func (s *DescribeUserDcdnStatusResponseBody) SetRequestId(v string) *DescribeUserDcdnStatusResponseBody {
	s.RequestId = &v
	return s
}

type DescribeUserDcdnStatusResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeUserDcdnStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeUserDcdnStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserDcdnStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeUserDcdnStatusResponse) SetHeaders(v map[string]*string) *DescribeUserDcdnStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeUserDcdnStatusResponse) SetBody(v *DescribeUserDcdnStatusResponseBody) *DescribeUserDcdnStatusResponse {
	s.Body = v
	return s
}

type DescribeUserErStatusRequest struct {
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeUserErStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserErStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeUserErStatusRequest) SetOwnerId(v int64) *DescribeUserErStatusRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeUserErStatusRequest) SetSecurityToken(v string) *DescribeUserErStatusRequest {
	s.SecurityToken = &v
	return s
}

type DescribeUserErStatusResponseBody struct {
	Enabled       *bool   `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	InDebt        *bool   `json:"InDebt,omitempty" xml:"InDebt,omitempty"`
	InDebtOverdue *bool   `json:"InDebtOverdue,omitempty" xml:"InDebtOverdue,omitempty"`
	OnService     *bool   `json:"OnService,omitempty" xml:"OnService,omitempty"`
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeUserErStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserErStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUserErStatusResponseBody) SetEnabled(v bool) *DescribeUserErStatusResponseBody {
	s.Enabled = &v
	return s
}

func (s *DescribeUserErStatusResponseBody) SetInDebt(v bool) *DescribeUserErStatusResponseBody {
	s.InDebt = &v
	return s
}

func (s *DescribeUserErStatusResponseBody) SetInDebtOverdue(v bool) *DescribeUserErStatusResponseBody {
	s.InDebtOverdue = &v
	return s
}

func (s *DescribeUserErStatusResponseBody) SetOnService(v bool) *DescribeUserErStatusResponseBody {
	s.OnService = &v
	return s
}

func (s *DescribeUserErStatusResponseBody) SetRequestId(v string) *DescribeUserErStatusResponseBody {
	s.RequestId = &v
	return s
}

type DescribeUserErStatusResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeUserErStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeUserErStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserErStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeUserErStatusResponse) SetHeaders(v map[string]*string) *DescribeUserErStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeUserErStatusResponse) SetBody(v *DescribeUserErStatusResponseBody) *DescribeUserErStatusResponse {
	s.Body = v
	return s
}

type DescribeUserLogserviceStatusRequest struct {
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeUserLogserviceStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserLogserviceStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeUserLogserviceStatusRequest) SetOwnerId(v int64) *DescribeUserLogserviceStatusRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeUserLogserviceStatusRequest) SetSecurityToken(v string) *DescribeUserLogserviceStatusRequest {
	s.SecurityToken = &v
	return s
}

type DescribeUserLogserviceStatusResponseBody struct {
	Enabled       *bool   `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	InDebt        *bool   `json:"InDebt,omitempty" xml:"InDebt,omitempty"`
	InDebtOverdue *bool   `json:"InDebtOverdue,omitempty" xml:"InDebtOverdue,omitempty"`
	OnService     *bool   `json:"OnService,omitempty" xml:"OnService,omitempty"`
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeUserLogserviceStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserLogserviceStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUserLogserviceStatusResponseBody) SetEnabled(v bool) *DescribeUserLogserviceStatusResponseBody {
	s.Enabled = &v
	return s
}

func (s *DescribeUserLogserviceStatusResponseBody) SetInDebt(v bool) *DescribeUserLogserviceStatusResponseBody {
	s.InDebt = &v
	return s
}

func (s *DescribeUserLogserviceStatusResponseBody) SetInDebtOverdue(v bool) *DescribeUserLogserviceStatusResponseBody {
	s.InDebtOverdue = &v
	return s
}

func (s *DescribeUserLogserviceStatusResponseBody) SetOnService(v bool) *DescribeUserLogserviceStatusResponseBody {
	s.OnService = &v
	return s
}

func (s *DescribeUserLogserviceStatusResponseBody) SetRequestId(v string) *DescribeUserLogserviceStatusResponseBody {
	s.RequestId = &v
	return s
}

type DescribeUserLogserviceStatusResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeUserLogserviceStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeUserLogserviceStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserLogserviceStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeUserLogserviceStatusResponse) SetHeaders(v map[string]*string) *DescribeUserLogserviceStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeUserLogserviceStatusResponse) SetBody(v *DescribeUserLogserviceStatusResponseBody) *DescribeUserLogserviceStatusResponse {
	s.Body = v
	return s
}

type EditRoutineConfRequest struct {
	Description *string                `json:"Description,omitempty" xml:"Description,omitempty"`
	EnvConf     map[string]interface{} `json:"EnvConf,omitempty" xml:"EnvConf,omitempty"`
	Name        *string                `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerId     *int64                 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s EditRoutineConfRequest) String() string {
	return tea.Prettify(s)
}

func (s EditRoutineConfRequest) GoString() string {
	return s.String()
}

func (s *EditRoutineConfRequest) SetDescription(v string) *EditRoutineConfRequest {
	s.Description = &v
	return s
}

func (s *EditRoutineConfRequest) SetEnvConf(v map[string]interface{}) *EditRoutineConfRequest {
	s.EnvConf = v
	return s
}

func (s *EditRoutineConfRequest) SetName(v string) *EditRoutineConfRequest {
	s.Name = &v
	return s
}

func (s *EditRoutineConfRequest) SetOwnerId(v int64) *EditRoutineConfRequest {
	s.OwnerId = &v
	return s
}

type EditRoutineConfShrinkRequest struct {
	Description   *string `json:"Description,omitempty" xml:"Description,omitempty"`
	EnvConfShrink *string `json:"EnvConf,omitempty" xml:"EnvConf,omitempty"`
	Name          *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s EditRoutineConfShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s EditRoutineConfShrinkRequest) GoString() string {
	return s.String()
}

func (s *EditRoutineConfShrinkRequest) SetDescription(v string) *EditRoutineConfShrinkRequest {
	s.Description = &v
	return s
}

func (s *EditRoutineConfShrinkRequest) SetEnvConfShrink(v string) *EditRoutineConfShrinkRequest {
	s.EnvConfShrink = &v
	return s
}

func (s *EditRoutineConfShrinkRequest) SetName(v string) *EditRoutineConfShrinkRequest {
	s.Name = &v
	return s
}

func (s *EditRoutineConfShrinkRequest) SetOwnerId(v int64) *EditRoutineConfShrinkRequest {
	s.OwnerId = &v
	return s
}

type EditRoutineConfResponseBody struct {
	Content   map[string]interface{} `json:"Content,omitempty" xml:"Content,omitempty"`
	RequestId *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EditRoutineConfResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EditRoutineConfResponseBody) GoString() string {
	return s.String()
}

func (s *EditRoutineConfResponseBody) SetContent(v map[string]interface{}) *EditRoutineConfResponseBody {
	s.Content = v
	return s
}

func (s *EditRoutineConfResponseBody) SetRequestId(v string) *EditRoutineConfResponseBody {
	s.RequestId = &v
	return s
}

type EditRoutineConfResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *EditRoutineConfResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s EditRoutineConfResponse) String() string {
	return tea.Prettify(s)
}

func (s EditRoutineConfResponse) GoString() string {
	return s.String()
}

func (s *EditRoutineConfResponse) SetHeaders(v map[string]*string) *EditRoutineConfResponse {
	s.Headers = v
	return s
}

func (s *EditRoutineConfResponse) SetBody(v *EditRoutineConfResponseBody) *EditRoutineConfResponse {
	s.Body = v
	return s
}

type ListDcdnEsTemplateInfoRequest struct {
	Language   *string `json:"Language,omitempty" xml:"Language,omitempty"`
	Op         *string `json:"Op,omitempty" xml:"Op,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListDcdnEsTemplateInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDcdnEsTemplateInfoRequest) GoString() string {
	return s.String()
}

func (s *ListDcdnEsTemplateInfoRequest) SetLanguage(v string) *ListDcdnEsTemplateInfoRequest {
	s.Language = &v
	return s
}

func (s *ListDcdnEsTemplateInfoRequest) SetOp(v string) *ListDcdnEsTemplateInfoRequest {
	s.Op = &v
	return s
}

func (s *ListDcdnEsTemplateInfoRequest) SetOwnerId(v int64) *ListDcdnEsTemplateInfoRequest {
	s.OwnerId = &v
	return s
}

func (s *ListDcdnEsTemplateInfoRequest) SetPageNumber(v int32) *ListDcdnEsTemplateInfoRequest {
	s.PageNumber = &v
	return s
}

func (s *ListDcdnEsTemplateInfoRequest) SetPageSize(v int32) *ListDcdnEsTemplateInfoRequest {
	s.PageSize = &v
	return s
}

type ListDcdnEsTemplateInfoResponseBody struct {
	DataItems  *ListDcdnEsTemplateInfoResponseBodyDataItems `json:"DataItems,omitempty" xml:"DataItems,omitempty" type:"Struct"`
	PageNumber *int32                                       `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                       `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32                                       `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListDcdnEsTemplateInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDcdnEsTemplateInfoResponseBody) GoString() string {
	return s.String()
}

func (s *ListDcdnEsTemplateInfoResponseBody) SetDataItems(v *ListDcdnEsTemplateInfoResponseBodyDataItems) *ListDcdnEsTemplateInfoResponseBody {
	s.DataItems = v
	return s
}

func (s *ListDcdnEsTemplateInfoResponseBody) SetPageNumber(v int32) *ListDcdnEsTemplateInfoResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListDcdnEsTemplateInfoResponseBody) SetPageSize(v int32) *ListDcdnEsTemplateInfoResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListDcdnEsTemplateInfoResponseBody) SetRequestId(v string) *ListDcdnEsTemplateInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDcdnEsTemplateInfoResponseBody) SetTotalCount(v int32) *ListDcdnEsTemplateInfoResponseBody {
	s.TotalCount = &v
	return s
}

type ListDcdnEsTemplateInfoResponseBodyDataItems struct {
	DataItem []*ListDcdnEsTemplateInfoResponseBodyDataItemsDataItem `json:"DataItem,omitempty" xml:"DataItem,omitempty" type:"Repeated"`
}

func (s ListDcdnEsTemplateInfoResponseBodyDataItems) String() string {
	return tea.Prettify(s)
}

func (s ListDcdnEsTemplateInfoResponseBodyDataItems) GoString() string {
	return s.String()
}

func (s *ListDcdnEsTemplateInfoResponseBodyDataItems) SetDataItem(v []*ListDcdnEsTemplateInfoResponseBodyDataItemsDataItem) *ListDcdnEsTemplateInfoResponseBodyDataItems {
	s.DataItem = v
	return s
}

type ListDcdnEsTemplateInfoResponseBodyDataItemsDataItem struct {
	Id       *int32  `json:"Id,omitempty" xml:"Id,omitempty"`
	TmplDef  *string `json:"TmplDef,omitempty" xml:"TmplDef,omitempty"`
	TmplDesc *string `json:"TmplDesc,omitempty" xml:"TmplDesc,omitempty"`
	TmplName *string `json:"TmplName,omitempty" xml:"TmplName,omitempty"`
}

func (s ListDcdnEsTemplateInfoResponseBodyDataItemsDataItem) String() string {
	return tea.Prettify(s)
}

func (s ListDcdnEsTemplateInfoResponseBodyDataItemsDataItem) GoString() string {
	return s.String()
}

func (s *ListDcdnEsTemplateInfoResponseBodyDataItemsDataItem) SetId(v int32) *ListDcdnEsTemplateInfoResponseBodyDataItemsDataItem {
	s.Id = &v
	return s
}

func (s *ListDcdnEsTemplateInfoResponseBodyDataItemsDataItem) SetTmplDef(v string) *ListDcdnEsTemplateInfoResponseBodyDataItemsDataItem {
	s.TmplDef = &v
	return s
}

func (s *ListDcdnEsTemplateInfoResponseBodyDataItemsDataItem) SetTmplDesc(v string) *ListDcdnEsTemplateInfoResponseBodyDataItemsDataItem {
	s.TmplDesc = &v
	return s
}

func (s *ListDcdnEsTemplateInfoResponseBodyDataItemsDataItem) SetTmplName(v string) *ListDcdnEsTemplateInfoResponseBodyDataItemsDataItem {
	s.TmplName = &v
	return s
}

type ListDcdnEsTemplateInfoResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListDcdnEsTemplateInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListDcdnEsTemplateInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDcdnEsTemplateInfoResponse) GoString() string {
	return s.String()
}

func (s *ListDcdnEsTemplateInfoResponse) SetHeaders(v map[string]*string) *ListDcdnEsTemplateInfoResponse {
	s.Headers = v
	return s
}

func (s *ListDcdnEsTemplateInfoResponse) SetBody(v *ListDcdnEsTemplateInfoResponseBody) *ListDcdnEsTemplateInfoResponse {
	s.Body = v
	return s
}

type ListDcdnRealTimeDeliveryProjectRequest struct {
	BusinessType *string `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
	DomainName   *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNumber   *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize     *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListDcdnRealTimeDeliveryProjectRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDcdnRealTimeDeliveryProjectRequest) GoString() string {
	return s.String()
}

func (s *ListDcdnRealTimeDeliveryProjectRequest) SetBusinessType(v string) *ListDcdnRealTimeDeliveryProjectRequest {
	s.BusinessType = &v
	return s
}

func (s *ListDcdnRealTimeDeliveryProjectRequest) SetDomainName(v string) *ListDcdnRealTimeDeliveryProjectRequest {
	s.DomainName = &v
	return s
}

func (s *ListDcdnRealTimeDeliveryProjectRequest) SetOwnerId(v int64) *ListDcdnRealTimeDeliveryProjectRequest {
	s.OwnerId = &v
	return s
}

func (s *ListDcdnRealTimeDeliveryProjectRequest) SetPageNumber(v int32) *ListDcdnRealTimeDeliveryProjectRequest {
	s.PageNumber = &v
	return s
}

func (s *ListDcdnRealTimeDeliveryProjectRequest) SetPageSize(v int32) *ListDcdnRealTimeDeliveryProjectRequest {
	s.PageSize = &v
	return s
}

type ListDcdnRealTimeDeliveryProjectResponseBody struct {
	Content    *ListDcdnRealTimeDeliveryProjectResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	RequestId  *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32                                              `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListDcdnRealTimeDeliveryProjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDcdnRealTimeDeliveryProjectResponseBody) GoString() string {
	return s.String()
}

func (s *ListDcdnRealTimeDeliveryProjectResponseBody) SetContent(v *ListDcdnRealTimeDeliveryProjectResponseBodyContent) *ListDcdnRealTimeDeliveryProjectResponseBody {
	s.Content = v
	return s
}

func (s *ListDcdnRealTimeDeliveryProjectResponseBody) SetRequestId(v string) *ListDcdnRealTimeDeliveryProjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDcdnRealTimeDeliveryProjectResponseBody) SetTotalCount(v int32) *ListDcdnRealTimeDeliveryProjectResponseBody {
	s.TotalCount = &v
	return s
}

type ListDcdnRealTimeDeliveryProjectResponseBodyContent struct {
	Projects []*ListDcdnRealTimeDeliveryProjectResponseBodyContentProjects `json:"Projects,omitempty" xml:"Projects,omitempty" type:"Repeated"`
}

func (s ListDcdnRealTimeDeliveryProjectResponseBodyContent) String() string {
	return tea.Prettify(s)
}

func (s ListDcdnRealTimeDeliveryProjectResponseBodyContent) GoString() string {
	return s.String()
}

func (s *ListDcdnRealTimeDeliveryProjectResponseBodyContent) SetProjects(v []*ListDcdnRealTimeDeliveryProjectResponseBodyContentProjects) *ListDcdnRealTimeDeliveryProjectResponseBodyContent {
	s.Projects = v
	return s
}

type ListDcdnRealTimeDeliveryProjectResponseBodyContentProjects struct {
	BusinessType *string  `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
	DataCenter   *string  `json:"DataCenter,omitempty" xml:"DataCenter,omitempty"`
	DomainName   *string  `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	FieldName    *string  `json:"FieldName,omitempty" xml:"FieldName,omitempty"`
	ProjectName  *string  `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	SLSLogStore  *string  `json:"SLSLogStore,omitempty" xml:"SLSLogStore,omitempty"`
	SLSProject   *string  `json:"SLSProject,omitempty" xml:"SLSProject,omitempty"`
	SLSRegion    *string  `json:"SLSRegion,omitempty" xml:"SLSRegion,omitempty"`
	SamplingRate *float32 `json:"SamplingRate,omitempty" xml:"SamplingRate,omitempty"`
	Type         *string  `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListDcdnRealTimeDeliveryProjectResponseBodyContentProjects) String() string {
	return tea.Prettify(s)
}

func (s ListDcdnRealTimeDeliveryProjectResponseBodyContentProjects) GoString() string {
	return s.String()
}

func (s *ListDcdnRealTimeDeliveryProjectResponseBodyContentProjects) SetBusinessType(v string) *ListDcdnRealTimeDeliveryProjectResponseBodyContentProjects {
	s.BusinessType = &v
	return s
}

func (s *ListDcdnRealTimeDeliveryProjectResponseBodyContentProjects) SetDataCenter(v string) *ListDcdnRealTimeDeliveryProjectResponseBodyContentProjects {
	s.DataCenter = &v
	return s
}

func (s *ListDcdnRealTimeDeliveryProjectResponseBodyContentProjects) SetDomainName(v string) *ListDcdnRealTimeDeliveryProjectResponseBodyContentProjects {
	s.DomainName = &v
	return s
}

func (s *ListDcdnRealTimeDeliveryProjectResponseBodyContentProjects) SetFieldName(v string) *ListDcdnRealTimeDeliveryProjectResponseBodyContentProjects {
	s.FieldName = &v
	return s
}

func (s *ListDcdnRealTimeDeliveryProjectResponseBodyContentProjects) SetProjectName(v string) *ListDcdnRealTimeDeliveryProjectResponseBodyContentProjects {
	s.ProjectName = &v
	return s
}

func (s *ListDcdnRealTimeDeliveryProjectResponseBodyContentProjects) SetSLSLogStore(v string) *ListDcdnRealTimeDeliveryProjectResponseBodyContentProjects {
	s.SLSLogStore = &v
	return s
}

func (s *ListDcdnRealTimeDeliveryProjectResponseBodyContentProjects) SetSLSProject(v string) *ListDcdnRealTimeDeliveryProjectResponseBodyContentProjects {
	s.SLSProject = &v
	return s
}

func (s *ListDcdnRealTimeDeliveryProjectResponseBodyContentProjects) SetSLSRegion(v string) *ListDcdnRealTimeDeliveryProjectResponseBodyContentProjects {
	s.SLSRegion = &v
	return s
}

func (s *ListDcdnRealTimeDeliveryProjectResponseBodyContentProjects) SetSamplingRate(v float32) *ListDcdnRealTimeDeliveryProjectResponseBodyContentProjects {
	s.SamplingRate = &v
	return s
}

func (s *ListDcdnRealTimeDeliveryProjectResponseBodyContentProjects) SetType(v string) *ListDcdnRealTimeDeliveryProjectResponseBodyContentProjects {
	s.Type = &v
	return s
}

type ListDcdnRealTimeDeliveryProjectResponse struct {
	Headers map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListDcdnRealTimeDeliveryProjectResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListDcdnRealTimeDeliveryProjectResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDcdnRealTimeDeliveryProjectResponse) GoString() string {
	return s.String()
}

func (s *ListDcdnRealTimeDeliveryProjectResponse) SetHeaders(v map[string]*string) *ListDcdnRealTimeDeliveryProjectResponse {
	s.Headers = v
	return s
}

func (s *ListDcdnRealTimeDeliveryProjectResponse) SetBody(v *ListDcdnRealTimeDeliveryProjectResponseBody) *ListDcdnRealTimeDeliveryProjectResponse {
	s.Body = v
	return s
}

type ModifyDCdnDomainSchdmByPropertyRequest struct {
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	Property   *string `json:"Property,omitempty" xml:"Property,omitempty"`
}

func (s ModifyDCdnDomainSchdmByPropertyRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDCdnDomainSchdmByPropertyRequest) GoString() string {
	return s.String()
}

func (s *ModifyDCdnDomainSchdmByPropertyRequest) SetDomainName(v string) *ModifyDCdnDomainSchdmByPropertyRequest {
	s.DomainName = &v
	return s
}

func (s *ModifyDCdnDomainSchdmByPropertyRequest) SetOwnerId(v int64) *ModifyDCdnDomainSchdmByPropertyRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyDCdnDomainSchdmByPropertyRequest) SetProperty(v string) *ModifyDCdnDomainSchdmByPropertyRequest {
	s.Property = &v
	return s
}

type ModifyDCdnDomainSchdmByPropertyResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDCdnDomainSchdmByPropertyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyDCdnDomainSchdmByPropertyResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDCdnDomainSchdmByPropertyResponseBody) SetRequestId(v string) *ModifyDCdnDomainSchdmByPropertyResponseBody {
	s.RequestId = &v
	return s
}

type ModifyDCdnDomainSchdmByPropertyResponse struct {
	Headers map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyDCdnDomainSchdmByPropertyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyDCdnDomainSchdmByPropertyResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyDCdnDomainSchdmByPropertyResponse) GoString() string {
	return s.String()
}

func (s *ModifyDCdnDomainSchdmByPropertyResponse) SetHeaders(v map[string]*string) *ModifyDCdnDomainSchdmByPropertyResponse {
	s.Headers = v
	return s
}

func (s *ModifyDCdnDomainSchdmByPropertyResponse) SetBody(v *ModifyDCdnDomainSchdmByPropertyResponseBody) *ModifyDCdnDomainSchdmByPropertyResponse {
	s.Body = v
	return s
}

type OpenDcdnServiceRequest struct {
	BillType          *string `json:"BillType,omitempty" xml:"BillType,omitempty"`
	OwnerId           *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken     *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	WebsocketBillType *string `json:"WebsocketBillType,omitempty" xml:"WebsocketBillType,omitempty"`
}

func (s OpenDcdnServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s OpenDcdnServiceRequest) GoString() string {
	return s.String()
}

func (s *OpenDcdnServiceRequest) SetBillType(v string) *OpenDcdnServiceRequest {
	s.BillType = &v
	return s
}

func (s *OpenDcdnServiceRequest) SetOwnerId(v int64) *OpenDcdnServiceRequest {
	s.OwnerId = &v
	return s
}

func (s *OpenDcdnServiceRequest) SetSecurityToken(v string) *OpenDcdnServiceRequest {
	s.SecurityToken = &v
	return s
}

func (s *OpenDcdnServiceRequest) SetWebsocketBillType(v string) *OpenDcdnServiceRequest {
	s.WebsocketBillType = &v
	return s
}

type OpenDcdnServiceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OpenDcdnServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OpenDcdnServiceResponseBody) GoString() string {
	return s.String()
}

func (s *OpenDcdnServiceResponseBody) SetRequestId(v string) *OpenDcdnServiceResponseBody {
	s.RequestId = &v
	return s
}

type OpenDcdnServiceResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *OpenDcdnServiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s OpenDcdnServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s OpenDcdnServiceResponse) GoString() string {
	return s.String()
}

func (s *OpenDcdnServiceResponse) SetHeaders(v map[string]*string) *OpenDcdnServiceResponse {
	s.Headers = v
	return s
}

func (s *OpenDcdnServiceResponse) SetBody(v *OpenDcdnServiceResponseBody) *OpenDcdnServiceResponse {
	s.Body = v
	return s
}

type PreloadDcdnObjectCachesRequest struct {
	Area          *string `json:"Area,omitempty" xml:"Area,omitempty"`
	ObjectPath    *string `json:"ObjectPath,omitempty" xml:"ObjectPath,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s PreloadDcdnObjectCachesRequest) String() string {
	return tea.Prettify(s)
}

func (s PreloadDcdnObjectCachesRequest) GoString() string {
	return s.String()
}

func (s *PreloadDcdnObjectCachesRequest) SetArea(v string) *PreloadDcdnObjectCachesRequest {
	s.Area = &v
	return s
}

func (s *PreloadDcdnObjectCachesRequest) SetObjectPath(v string) *PreloadDcdnObjectCachesRequest {
	s.ObjectPath = &v
	return s
}

func (s *PreloadDcdnObjectCachesRequest) SetOwnerId(v int64) *PreloadDcdnObjectCachesRequest {
	s.OwnerId = &v
	return s
}

func (s *PreloadDcdnObjectCachesRequest) SetSecurityToken(v string) *PreloadDcdnObjectCachesRequest {
	s.SecurityToken = &v
	return s
}

type PreloadDcdnObjectCachesResponseBody struct {
	PreloadTaskId *string `json:"PreloadTaskId,omitempty" xml:"PreloadTaskId,omitempty"`
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s PreloadDcdnObjectCachesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PreloadDcdnObjectCachesResponseBody) GoString() string {
	return s.String()
}

func (s *PreloadDcdnObjectCachesResponseBody) SetPreloadTaskId(v string) *PreloadDcdnObjectCachesResponseBody {
	s.PreloadTaskId = &v
	return s
}

func (s *PreloadDcdnObjectCachesResponseBody) SetRequestId(v string) *PreloadDcdnObjectCachesResponseBody {
	s.RequestId = &v
	return s
}

type PreloadDcdnObjectCachesResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *PreloadDcdnObjectCachesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PreloadDcdnObjectCachesResponse) String() string {
	return tea.Prettify(s)
}

func (s PreloadDcdnObjectCachesResponse) GoString() string {
	return s.String()
}

func (s *PreloadDcdnObjectCachesResponse) SetHeaders(v map[string]*string) *PreloadDcdnObjectCachesResponse {
	s.Headers = v
	return s
}

func (s *PreloadDcdnObjectCachesResponse) SetBody(v *PreloadDcdnObjectCachesResponseBody) *PreloadDcdnObjectCachesResponse {
	s.Body = v
	return s
}

type PublishDcdnStagingConfigToProductionRequest struct {
	DomainName   *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	FunctionName *string `json:"FunctionName,omitempty" xml:"FunctionName,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s PublishDcdnStagingConfigToProductionRequest) String() string {
	return tea.Prettify(s)
}

func (s PublishDcdnStagingConfigToProductionRequest) GoString() string {
	return s.String()
}

func (s *PublishDcdnStagingConfigToProductionRequest) SetDomainName(v string) *PublishDcdnStagingConfigToProductionRequest {
	s.DomainName = &v
	return s
}

func (s *PublishDcdnStagingConfigToProductionRequest) SetFunctionName(v string) *PublishDcdnStagingConfigToProductionRequest {
	s.FunctionName = &v
	return s
}

func (s *PublishDcdnStagingConfigToProductionRequest) SetOwnerId(v int64) *PublishDcdnStagingConfigToProductionRequest {
	s.OwnerId = &v
	return s
}

type PublishDcdnStagingConfigToProductionResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s PublishDcdnStagingConfigToProductionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PublishDcdnStagingConfigToProductionResponseBody) GoString() string {
	return s.String()
}

func (s *PublishDcdnStagingConfigToProductionResponseBody) SetRequestId(v string) *PublishDcdnStagingConfigToProductionResponseBody {
	s.RequestId = &v
	return s
}

type PublishDcdnStagingConfigToProductionResponse struct {
	Headers map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *PublishDcdnStagingConfigToProductionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PublishDcdnStagingConfigToProductionResponse) String() string {
	return tea.Prettify(s)
}

func (s PublishDcdnStagingConfigToProductionResponse) GoString() string {
	return s.String()
}

func (s *PublishDcdnStagingConfigToProductionResponse) SetHeaders(v map[string]*string) *PublishDcdnStagingConfigToProductionResponse {
	s.Headers = v
	return s
}

func (s *PublishDcdnStagingConfigToProductionResponse) SetBody(v *PublishDcdnStagingConfigToProductionResponseBody) *PublishDcdnStagingConfigToProductionResponse {
	s.Body = v
	return s
}

type PublishRoutineCodeRevisionRequest struct {
	Envs               map[string]interface{} `json:"Envs,omitempty" xml:"Envs,omitempty"`
	Name               *string                `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerId            *int64                 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SelectCodeRevision *string                `json:"SelectCodeRevision,omitempty" xml:"SelectCodeRevision,omitempty"`
}

func (s PublishRoutineCodeRevisionRequest) String() string {
	return tea.Prettify(s)
}

func (s PublishRoutineCodeRevisionRequest) GoString() string {
	return s.String()
}

func (s *PublishRoutineCodeRevisionRequest) SetEnvs(v map[string]interface{}) *PublishRoutineCodeRevisionRequest {
	s.Envs = v
	return s
}

func (s *PublishRoutineCodeRevisionRequest) SetName(v string) *PublishRoutineCodeRevisionRequest {
	s.Name = &v
	return s
}

func (s *PublishRoutineCodeRevisionRequest) SetOwnerId(v int64) *PublishRoutineCodeRevisionRequest {
	s.OwnerId = &v
	return s
}

func (s *PublishRoutineCodeRevisionRequest) SetSelectCodeRevision(v string) *PublishRoutineCodeRevisionRequest {
	s.SelectCodeRevision = &v
	return s
}

type PublishRoutineCodeRevisionShrinkRequest struct {
	EnvsShrink         *string `json:"Envs,omitempty" xml:"Envs,omitempty"`
	Name               *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerId            *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SelectCodeRevision *string `json:"SelectCodeRevision,omitempty" xml:"SelectCodeRevision,omitempty"`
}

func (s PublishRoutineCodeRevisionShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s PublishRoutineCodeRevisionShrinkRequest) GoString() string {
	return s.String()
}

func (s *PublishRoutineCodeRevisionShrinkRequest) SetEnvsShrink(v string) *PublishRoutineCodeRevisionShrinkRequest {
	s.EnvsShrink = &v
	return s
}

func (s *PublishRoutineCodeRevisionShrinkRequest) SetName(v string) *PublishRoutineCodeRevisionShrinkRequest {
	s.Name = &v
	return s
}

func (s *PublishRoutineCodeRevisionShrinkRequest) SetOwnerId(v int64) *PublishRoutineCodeRevisionShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *PublishRoutineCodeRevisionShrinkRequest) SetSelectCodeRevision(v string) *PublishRoutineCodeRevisionShrinkRequest {
	s.SelectCodeRevision = &v
	return s
}

type PublishRoutineCodeRevisionResponseBody struct {
	Content   map[string]interface{} `json:"Content,omitempty" xml:"Content,omitempty"`
	RequestId *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s PublishRoutineCodeRevisionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PublishRoutineCodeRevisionResponseBody) GoString() string {
	return s.String()
}

func (s *PublishRoutineCodeRevisionResponseBody) SetContent(v map[string]interface{}) *PublishRoutineCodeRevisionResponseBody {
	s.Content = v
	return s
}

func (s *PublishRoutineCodeRevisionResponseBody) SetRequestId(v string) *PublishRoutineCodeRevisionResponseBody {
	s.RequestId = &v
	return s
}

type PublishRoutineCodeRevisionResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *PublishRoutineCodeRevisionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PublishRoutineCodeRevisionResponse) String() string {
	return tea.Prettify(s)
}

func (s PublishRoutineCodeRevisionResponse) GoString() string {
	return s.String()
}

func (s *PublishRoutineCodeRevisionResponse) SetHeaders(v map[string]*string) *PublishRoutineCodeRevisionResponse {
	s.Headers = v
	return s
}

func (s *PublishRoutineCodeRevisionResponse) SetBody(v *PublishRoutineCodeRevisionResponseBody) *PublishRoutineCodeRevisionResponse {
	s.Body = v
	return s
}

type RefreshDcdnObjectCachesRequest struct {
	ObjectPath    *string `json:"ObjectPath,omitempty" xml:"ObjectPath,omitempty"`
	ObjectType    *string `json:"ObjectType,omitempty" xml:"ObjectType,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s RefreshDcdnObjectCachesRequest) String() string {
	return tea.Prettify(s)
}

func (s RefreshDcdnObjectCachesRequest) GoString() string {
	return s.String()
}

func (s *RefreshDcdnObjectCachesRequest) SetObjectPath(v string) *RefreshDcdnObjectCachesRequest {
	s.ObjectPath = &v
	return s
}

func (s *RefreshDcdnObjectCachesRequest) SetObjectType(v string) *RefreshDcdnObjectCachesRequest {
	s.ObjectType = &v
	return s
}

func (s *RefreshDcdnObjectCachesRequest) SetOwnerId(v int64) *RefreshDcdnObjectCachesRequest {
	s.OwnerId = &v
	return s
}

func (s *RefreshDcdnObjectCachesRequest) SetSecurityToken(v string) *RefreshDcdnObjectCachesRequest {
	s.SecurityToken = &v
	return s
}

type RefreshDcdnObjectCachesResponseBody struct {
	RefreshTaskId *string `json:"RefreshTaskId,omitempty" xml:"RefreshTaskId,omitempty"`
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RefreshDcdnObjectCachesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RefreshDcdnObjectCachesResponseBody) GoString() string {
	return s.String()
}

func (s *RefreshDcdnObjectCachesResponseBody) SetRefreshTaskId(v string) *RefreshDcdnObjectCachesResponseBody {
	s.RefreshTaskId = &v
	return s
}

func (s *RefreshDcdnObjectCachesResponseBody) SetRequestId(v string) *RefreshDcdnObjectCachesResponseBody {
	s.RequestId = &v
	return s
}

type RefreshDcdnObjectCachesResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RefreshDcdnObjectCachesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RefreshDcdnObjectCachesResponse) String() string {
	return tea.Prettify(s)
}

func (s RefreshDcdnObjectCachesResponse) GoString() string {
	return s.String()
}

func (s *RefreshDcdnObjectCachesResponse) SetHeaders(v map[string]*string) *RefreshDcdnObjectCachesResponse {
	s.Headers = v
	return s
}

func (s *RefreshDcdnObjectCachesResponse) SetBody(v *RefreshDcdnObjectCachesResponseBody) *RefreshDcdnObjectCachesResponse {
	s.Body = v
	return s
}

type RollbackDcdnStagingConfigRequest struct {
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s RollbackDcdnStagingConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s RollbackDcdnStagingConfigRequest) GoString() string {
	return s.String()
}

func (s *RollbackDcdnStagingConfigRequest) SetDomainName(v string) *RollbackDcdnStagingConfigRequest {
	s.DomainName = &v
	return s
}

func (s *RollbackDcdnStagingConfigRequest) SetOwnerId(v int64) *RollbackDcdnStagingConfigRequest {
	s.OwnerId = &v
	return s
}

type RollbackDcdnStagingConfigResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RollbackDcdnStagingConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RollbackDcdnStagingConfigResponseBody) GoString() string {
	return s.String()
}

func (s *RollbackDcdnStagingConfigResponseBody) SetRequestId(v string) *RollbackDcdnStagingConfigResponseBody {
	s.RequestId = &v
	return s
}

type RollbackDcdnStagingConfigResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RollbackDcdnStagingConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RollbackDcdnStagingConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s RollbackDcdnStagingConfigResponse) GoString() string {
	return s.String()
}

func (s *RollbackDcdnStagingConfigResponse) SetHeaders(v map[string]*string) *RollbackDcdnStagingConfigResponse {
	s.Headers = v
	return s
}

func (s *RollbackDcdnStagingConfigResponse) SetBody(v *RollbackDcdnStagingConfigResponseBody) *RollbackDcdnStagingConfigResponse {
	s.Body = v
	return s
}

type SetDcdnConfigOfVersionRequest struct {
	ConfigId      *string `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	FunctionArgs  *string `json:"FunctionArgs,omitempty" xml:"FunctionArgs,omitempty"`
	FunctionId    *int64  `json:"FunctionId,omitempty" xml:"FunctionId,omitempty"`
	FunctionName  *string `json:"FunctionName,omitempty" xml:"FunctionName,omitempty"`
	OwnerAccount  *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	VersionId     *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
}

func (s SetDcdnConfigOfVersionRequest) String() string {
	return tea.Prettify(s)
}

func (s SetDcdnConfigOfVersionRequest) GoString() string {
	return s.String()
}

func (s *SetDcdnConfigOfVersionRequest) SetConfigId(v string) *SetDcdnConfigOfVersionRequest {
	s.ConfigId = &v
	return s
}

func (s *SetDcdnConfigOfVersionRequest) SetFunctionArgs(v string) *SetDcdnConfigOfVersionRequest {
	s.FunctionArgs = &v
	return s
}

func (s *SetDcdnConfigOfVersionRequest) SetFunctionId(v int64) *SetDcdnConfigOfVersionRequest {
	s.FunctionId = &v
	return s
}

func (s *SetDcdnConfigOfVersionRequest) SetFunctionName(v string) *SetDcdnConfigOfVersionRequest {
	s.FunctionName = &v
	return s
}

func (s *SetDcdnConfigOfVersionRequest) SetOwnerAccount(v string) *SetDcdnConfigOfVersionRequest {
	s.OwnerAccount = &v
	return s
}

func (s *SetDcdnConfigOfVersionRequest) SetOwnerId(v int64) *SetDcdnConfigOfVersionRequest {
	s.OwnerId = &v
	return s
}

func (s *SetDcdnConfigOfVersionRequest) SetSecurityToken(v string) *SetDcdnConfigOfVersionRequest {
	s.SecurityToken = &v
	return s
}

func (s *SetDcdnConfigOfVersionRequest) SetVersionId(v string) *SetDcdnConfigOfVersionRequest {
	s.VersionId = &v
	return s
}

type SetDcdnConfigOfVersionResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetDcdnConfigOfVersionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetDcdnConfigOfVersionResponseBody) GoString() string {
	return s.String()
}

func (s *SetDcdnConfigOfVersionResponseBody) SetRequestId(v string) *SetDcdnConfigOfVersionResponseBody {
	s.RequestId = &v
	return s
}

type SetDcdnConfigOfVersionResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SetDcdnConfigOfVersionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetDcdnConfigOfVersionResponse) String() string {
	return tea.Prettify(s)
}

func (s SetDcdnConfigOfVersionResponse) GoString() string {
	return s.String()
}

func (s *SetDcdnConfigOfVersionResponse) SetHeaders(v map[string]*string) *SetDcdnConfigOfVersionResponse {
	s.Headers = v
	return s
}

func (s *SetDcdnConfigOfVersionResponse) SetBody(v *SetDcdnConfigOfVersionResponseBody) *SetDcdnConfigOfVersionResponse {
	s.Body = v
	return s
}

type SetDcdnDomainCSRCertificateRequest struct {
	DomainName        *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	OwnerId           *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ServerCertificate *string `json:"ServerCertificate,omitempty" xml:"ServerCertificate,omitempty"`
}

func (s SetDcdnDomainCSRCertificateRequest) String() string {
	return tea.Prettify(s)
}

func (s SetDcdnDomainCSRCertificateRequest) GoString() string {
	return s.String()
}

func (s *SetDcdnDomainCSRCertificateRequest) SetDomainName(v string) *SetDcdnDomainCSRCertificateRequest {
	s.DomainName = &v
	return s
}

func (s *SetDcdnDomainCSRCertificateRequest) SetOwnerId(v int64) *SetDcdnDomainCSRCertificateRequest {
	s.OwnerId = &v
	return s
}

func (s *SetDcdnDomainCSRCertificateRequest) SetServerCertificate(v string) *SetDcdnDomainCSRCertificateRequest {
	s.ServerCertificate = &v
	return s
}

type SetDcdnDomainCSRCertificateResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetDcdnDomainCSRCertificateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetDcdnDomainCSRCertificateResponseBody) GoString() string {
	return s.String()
}

func (s *SetDcdnDomainCSRCertificateResponseBody) SetRequestId(v string) *SetDcdnDomainCSRCertificateResponseBody {
	s.RequestId = &v
	return s
}

type SetDcdnDomainCSRCertificateResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SetDcdnDomainCSRCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetDcdnDomainCSRCertificateResponse) String() string {
	return tea.Prettify(s)
}

func (s SetDcdnDomainCSRCertificateResponse) GoString() string {
	return s.String()
}

func (s *SetDcdnDomainCSRCertificateResponse) SetHeaders(v map[string]*string) *SetDcdnDomainCSRCertificateResponse {
	s.Headers = v
	return s
}

func (s *SetDcdnDomainCSRCertificateResponse) SetBody(v *SetDcdnDomainCSRCertificateResponseBody) *SetDcdnDomainCSRCertificateResponse {
	s.Body = v
	return s
}

type SetDcdnDomainCertificateRequest struct {
	CertName      *string `json:"CertName,omitempty" xml:"CertName,omitempty"`
	CertType      *string `json:"CertType,omitempty" xml:"CertType,omitempty"`
	DomainName    *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	ForceSet      *string `json:"ForceSet,omitempty" xml:"ForceSet,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	Region        *string `json:"Region,omitempty" xml:"Region,omitempty"`
	SSLPri        *string `json:"SSLPri,omitempty" xml:"SSLPri,omitempty"`
	SSLProtocol   *string `json:"SSLProtocol,omitempty" xml:"SSLProtocol,omitempty"`
	SSLPub        *string `json:"SSLPub,omitempty" xml:"SSLPub,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s SetDcdnDomainCertificateRequest) String() string {
	return tea.Prettify(s)
}

func (s SetDcdnDomainCertificateRequest) GoString() string {
	return s.String()
}

func (s *SetDcdnDomainCertificateRequest) SetCertName(v string) *SetDcdnDomainCertificateRequest {
	s.CertName = &v
	return s
}

func (s *SetDcdnDomainCertificateRequest) SetCertType(v string) *SetDcdnDomainCertificateRequest {
	s.CertType = &v
	return s
}

func (s *SetDcdnDomainCertificateRequest) SetDomainName(v string) *SetDcdnDomainCertificateRequest {
	s.DomainName = &v
	return s
}

func (s *SetDcdnDomainCertificateRequest) SetForceSet(v string) *SetDcdnDomainCertificateRequest {
	s.ForceSet = &v
	return s
}

func (s *SetDcdnDomainCertificateRequest) SetOwnerId(v int64) *SetDcdnDomainCertificateRequest {
	s.OwnerId = &v
	return s
}

func (s *SetDcdnDomainCertificateRequest) SetRegion(v string) *SetDcdnDomainCertificateRequest {
	s.Region = &v
	return s
}

func (s *SetDcdnDomainCertificateRequest) SetSSLPri(v string) *SetDcdnDomainCertificateRequest {
	s.SSLPri = &v
	return s
}

func (s *SetDcdnDomainCertificateRequest) SetSSLProtocol(v string) *SetDcdnDomainCertificateRequest {
	s.SSLProtocol = &v
	return s
}

func (s *SetDcdnDomainCertificateRequest) SetSSLPub(v string) *SetDcdnDomainCertificateRequest {
	s.SSLPub = &v
	return s
}

func (s *SetDcdnDomainCertificateRequest) SetSecurityToken(v string) *SetDcdnDomainCertificateRequest {
	s.SecurityToken = &v
	return s
}

type SetDcdnDomainCertificateResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetDcdnDomainCertificateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetDcdnDomainCertificateResponseBody) GoString() string {
	return s.String()
}

func (s *SetDcdnDomainCertificateResponseBody) SetRequestId(v string) *SetDcdnDomainCertificateResponseBody {
	s.RequestId = &v
	return s
}

type SetDcdnDomainCertificateResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SetDcdnDomainCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetDcdnDomainCertificateResponse) String() string {
	return tea.Prettify(s)
}

func (s SetDcdnDomainCertificateResponse) GoString() string {
	return s.String()
}

func (s *SetDcdnDomainCertificateResponse) SetHeaders(v map[string]*string) *SetDcdnDomainCertificateResponse {
	s.Headers = v
	return s
}

func (s *SetDcdnDomainCertificateResponse) SetBody(v *SetDcdnDomainCertificateResponseBody) *SetDcdnDomainCertificateResponse {
	s.Body = v
	return s
}

type SetDcdnDomainSMCertificateRequest struct {
	CertIdentifier *string `json:"CertIdentifier,omitempty" xml:"CertIdentifier,omitempty"`
	DomainName     *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	OwnerId        *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SSLProtocol    *string `json:"SSLProtocol,omitempty" xml:"SSLProtocol,omitempty"`
	SecurityToken  *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s SetDcdnDomainSMCertificateRequest) String() string {
	return tea.Prettify(s)
}

func (s SetDcdnDomainSMCertificateRequest) GoString() string {
	return s.String()
}

func (s *SetDcdnDomainSMCertificateRequest) SetCertIdentifier(v string) *SetDcdnDomainSMCertificateRequest {
	s.CertIdentifier = &v
	return s
}

func (s *SetDcdnDomainSMCertificateRequest) SetDomainName(v string) *SetDcdnDomainSMCertificateRequest {
	s.DomainName = &v
	return s
}

func (s *SetDcdnDomainSMCertificateRequest) SetOwnerId(v int64) *SetDcdnDomainSMCertificateRequest {
	s.OwnerId = &v
	return s
}

func (s *SetDcdnDomainSMCertificateRequest) SetSSLProtocol(v string) *SetDcdnDomainSMCertificateRequest {
	s.SSLProtocol = &v
	return s
}

func (s *SetDcdnDomainSMCertificateRequest) SetSecurityToken(v string) *SetDcdnDomainSMCertificateRequest {
	s.SecurityToken = &v
	return s
}

type SetDcdnDomainSMCertificateResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetDcdnDomainSMCertificateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetDcdnDomainSMCertificateResponseBody) GoString() string {
	return s.String()
}

func (s *SetDcdnDomainSMCertificateResponseBody) SetRequestId(v string) *SetDcdnDomainSMCertificateResponseBody {
	s.RequestId = &v
	return s
}

type SetDcdnDomainSMCertificateResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SetDcdnDomainSMCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetDcdnDomainSMCertificateResponse) String() string {
	return tea.Prettify(s)
}

func (s SetDcdnDomainSMCertificateResponse) GoString() string {
	return s.String()
}

func (s *SetDcdnDomainSMCertificateResponse) SetHeaders(v map[string]*string) *SetDcdnDomainSMCertificateResponse {
	s.Headers = v
	return s
}

func (s *SetDcdnDomainSMCertificateResponse) SetBody(v *SetDcdnDomainSMCertificateResponseBody) *SetDcdnDomainSMCertificateResponse {
	s.Body = v
	return s
}

type SetDcdnDomainStagingConfigRequest struct {
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	Functions  *string `json:"Functions,omitempty" xml:"Functions,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s SetDcdnDomainStagingConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s SetDcdnDomainStagingConfigRequest) GoString() string {
	return s.String()
}

func (s *SetDcdnDomainStagingConfigRequest) SetDomainName(v string) *SetDcdnDomainStagingConfigRequest {
	s.DomainName = &v
	return s
}

func (s *SetDcdnDomainStagingConfigRequest) SetFunctions(v string) *SetDcdnDomainStagingConfigRequest {
	s.Functions = &v
	return s
}

func (s *SetDcdnDomainStagingConfigRequest) SetOwnerId(v int64) *SetDcdnDomainStagingConfigRequest {
	s.OwnerId = &v
	return s
}

type SetDcdnDomainStagingConfigResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetDcdnDomainStagingConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetDcdnDomainStagingConfigResponseBody) GoString() string {
	return s.String()
}

func (s *SetDcdnDomainStagingConfigResponseBody) SetRequestId(v string) *SetDcdnDomainStagingConfigResponseBody {
	s.RequestId = &v
	return s
}

type SetDcdnDomainStagingConfigResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SetDcdnDomainStagingConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetDcdnDomainStagingConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s SetDcdnDomainStagingConfigResponse) GoString() string {
	return s.String()
}

func (s *SetDcdnDomainStagingConfigResponse) SetHeaders(v map[string]*string) *SetDcdnDomainStagingConfigResponse {
	s.Headers = v
	return s
}

func (s *SetDcdnDomainStagingConfigResponse) SetBody(v *SetDcdnDomainStagingConfigResponseBody) *SetDcdnDomainStagingConfigResponse {
	s.Body = v
	return s
}

type SetDcdnFullDomainsBlockIPRequest struct {
	BlockInterval *int32  `json:"BlockInterval,omitempty" xml:"BlockInterval,omitempty"`
	IPList        *string `json:"IPList,omitempty" xml:"IPList,omitempty"`
	OperationType *string `json:"OperationType,omitempty" xml:"OperationType,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s SetDcdnFullDomainsBlockIPRequest) String() string {
	return tea.Prettify(s)
}

func (s SetDcdnFullDomainsBlockIPRequest) GoString() string {
	return s.String()
}

func (s *SetDcdnFullDomainsBlockIPRequest) SetBlockInterval(v int32) *SetDcdnFullDomainsBlockIPRequest {
	s.BlockInterval = &v
	return s
}

func (s *SetDcdnFullDomainsBlockIPRequest) SetIPList(v string) *SetDcdnFullDomainsBlockIPRequest {
	s.IPList = &v
	return s
}

func (s *SetDcdnFullDomainsBlockIPRequest) SetOperationType(v string) *SetDcdnFullDomainsBlockIPRequest {
	s.OperationType = &v
	return s
}

func (s *SetDcdnFullDomainsBlockIPRequest) SetOwnerId(v int64) *SetDcdnFullDomainsBlockIPRequest {
	s.OwnerId = &v
	return s
}

type SetDcdnFullDomainsBlockIPResponseBody struct {
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetDcdnFullDomainsBlockIPResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetDcdnFullDomainsBlockIPResponseBody) GoString() string {
	return s.String()
}

func (s *SetDcdnFullDomainsBlockIPResponseBody) SetCode(v int32) *SetDcdnFullDomainsBlockIPResponseBody {
	s.Code = &v
	return s
}

func (s *SetDcdnFullDomainsBlockIPResponseBody) SetMessage(v string) *SetDcdnFullDomainsBlockIPResponseBody {
	s.Message = &v
	return s
}

func (s *SetDcdnFullDomainsBlockIPResponseBody) SetRequestId(v string) *SetDcdnFullDomainsBlockIPResponseBody {
	s.RequestId = &v
	return s
}

type SetDcdnFullDomainsBlockIPResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SetDcdnFullDomainsBlockIPResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetDcdnFullDomainsBlockIPResponse) String() string {
	return tea.Prettify(s)
}

func (s SetDcdnFullDomainsBlockIPResponse) GoString() string {
	return s.String()
}

func (s *SetDcdnFullDomainsBlockIPResponse) SetHeaders(v map[string]*string) *SetDcdnFullDomainsBlockIPResponse {
	s.Headers = v
	return s
}

func (s *SetDcdnFullDomainsBlockIPResponse) SetBody(v *SetDcdnFullDomainsBlockIPResponseBody) *SetDcdnFullDomainsBlockIPResponse {
	s.Body = v
	return s
}

type SetDcdnUserConfigRequest struct {
	Configs       *string `json:"Configs,omitempty" xml:"Configs,omitempty"`
	FunctionId    *int32  `json:"FunctionId,omitempty" xml:"FunctionId,omitempty"`
	OwnerAccount  *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s SetDcdnUserConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s SetDcdnUserConfigRequest) GoString() string {
	return s.String()
}

func (s *SetDcdnUserConfigRequest) SetConfigs(v string) *SetDcdnUserConfigRequest {
	s.Configs = &v
	return s
}

func (s *SetDcdnUserConfigRequest) SetFunctionId(v int32) *SetDcdnUserConfigRequest {
	s.FunctionId = &v
	return s
}

func (s *SetDcdnUserConfigRequest) SetOwnerAccount(v string) *SetDcdnUserConfigRequest {
	s.OwnerAccount = &v
	return s
}

func (s *SetDcdnUserConfigRequest) SetOwnerId(v int64) *SetDcdnUserConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *SetDcdnUserConfigRequest) SetSecurityToken(v string) *SetDcdnUserConfigRequest {
	s.SecurityToken = &v
	return s
}

type SetDcdnUserConfigResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetDcdnUserConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetDcdnUserConfigResponseBody) GoString() string {
	return s.String()
}

func (s *SetDcdnUserConfigResponseBody) SetRequestId(v string) *SetDcdnUserConfigResponseBody {
	s.RequestId = &v
	return s
}

type SetDcdnUserConfigResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SetDcdnUserConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetDcdnUserConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s SetDcdnUserConfigResponse) GoString() string {
	return s.String()
}

func (s *SetDcdnUserConfigResponse) SetHeaders(v map[string]*string) *SetDcdnUserConfigResponse {
	s.Headers = v
	return s
}

func (s *SetDcdnUserConfigResponse) SetBody(v *SetDcdnUserConfigResponseBody) *SetDcdnUserConfigResponse {
	s.Body = v
	return s
}

type SetRoutineSubdomainRequest struct {
	OwnerId    *int64                 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	Subdomains map[string]interface{} `json:"Subdomains,omitempty" xml:"Subdomains,omitempty"`
}

func (s SetRoutineSubdomainRequest) String() string {
	return tea.Prettify(s)
}

func (s SetRoutineSubdomainRequest) GoString() string {
	return s.String()
}

func (s *SetRoutineSubdomainRequest) SetOwnerId(v int64) *SetRoutineSubdomainRequest {
	s.OwnerId = &v
	return s
}

func (s *SetRoutineSubdomainRequest) SetSubdomains(v map[string]interface{}) *SetRoutineSubdomainRequest {
	s.Subdomains = v
	return s
}

type SetRoutineSubdomainShrinkRequest struct {
	OwnerId          *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SubdomainsShrink *string `json:"Subdomains,omitempty" xml:"Subdomains,omitempty"`
}

func (s SetRoutineSubdomainShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s SetRoutineSubdomainShrinkRequest) GoString() string {
	return s.String()
}

func (s *SetRoutineSubdomainShrinkRequest) SetOwnerId(v int64) *SetRoutineSubdomainShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *SetRoutineSubdomainShrinkRequest) SetSubdomainsShrink(v string) *SetRoutineSubdomainShrinkRequest {
	s.SubdomainsShrink = &v
	return s
}

type SetRoutineSubdomainResponseBody struct {
	Content   map[string]interface{} `json:"Content,omitempty" xml:"Content,omitempty"`
	RequestId *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetRoutineSubdomainResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetRoutineSubdomainResponseBody) GoString() string {
	return s.String()
}

func (s *SetRoutineSubdomainResponseBody) SetContent(v map[string]interface{}) *SetRoutineSubdomainResponseBody {
	s.Content = v
	return s
}

func (s *SetRoutineSubdomainResponseBody) SetRequestId(v string) *SetRoutineSubdomainResponseBody {
	s.RequestId = &v
	return s
}

type SetRoutineSubdomainResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SetRoutineSubdomainResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetRoutineSubdomainResponse) String() string {
	return tea.Prettify(s)
}

func (s SetRoutineSubdomainResponse) GoString() string {
	return s.String()
}

func (s *SetRoutineSubdomainResponse) SetHeaders(v map[string]*string) *SetRoutineSubdomainResponse {
	s.Headers = v
	return s
}

func (s *SetRoutineSubdomainResponse) SetBody(v *SetRoutineSubdomainResponseBody) *SetRoutineSubdomainResponse {
	s.Body = v
	return s
}

type StartDcdnDomainRequest struct {
	DomainName    *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s StartDcdnDomainRequest) String() string {
	return tea.Prettify(s)
}

func (s StartDcdnDomainRequest) GoString() string {
	return s.String()
}

func (s *StartDcdnDomainRequest) SetDomainName(v string) *StartDcdnDomainRequest {
	s.DomainName = &v
	return s
}

func (s *StartDcdnDomainRequest) SetOwnerId(v int64) *StartDcdnDomainRequest {
	s.OwnerId = &v
	return s
}

func (s *StartDcdnDomainRequest) SetSecurityToken(v string) *StartDcdnDomainRequest {
	s.SecurityToken = &v
	return s
}

type StartDcdnDomainResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StartDcdnDomainResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartDcdnDomainResponseBody) GoString() string {
	return s.String()
}

func (s *StartDcdnDomainResponseBody) SetRequestId(v string) *StartDcdnDomainResponseBody {
	s.RequestId = &v
	return s
}

type StartDcdnDomainResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *StartDcdnDomainResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StartDcdnDomainResponse) String() string {
	return tea.Prettify(s)
}

func (s StartDcdnDomainResponse) GoString() string {
	return s.String()
}

func (s *StartDcdnDomainResponse) SetHeaders(v map[string]*string) *StartDcdnDomainResponse {
	s.Headers = v
	return s
}

func (s *StartDcdnDomainResponse) SetBody(v *StartDcdnDomainResponseBody) *StartDcdnDomainResponse {
	s.Body = v
	return s
}

type StartDcdnIpaDomainRequest struct {
	DomainName    *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s StartDcdnIpaDomainRequest) String() string {
	return tea.Prettify(s)
}

func (s StartDcdnIpaDomainRequest) GoString() string {
	return s.String()
}

func (s *StartDcdnIpaDomainRequest) SetDomainName(v string) *StartDcdnIpaDomainRequest {
	s.DomainName = &v
	return s
}

func (s *StartDcdnIpaDomainRequest) SetOwnerId(v int64) *StartDcdnIpaDomainRequest {
	s.OwnerId = &v
	return s
}

func (s *StartDcdnIpaDomainRequest) SetSecurityToken(v string) *StartDcdnIpaDomainRequest {
	s.SecurityToken = &v
	return s
}

type StartDcdnIpaDomainResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StartDcdnIpaDomainResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartDcdnIpaDomainResponseBody) GoString() string {
	return s.String()
}

func (s *StartDcdnIpaDomainResponseBody) SetRequestId(v string) *StartDcdnIpaDomainResponseBody {
	s.RequestId = &v
	return s
}

type StartDcdnIpaDomainResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *StartDcdnIpaDomainResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StartDcdnIpaDomainResponse) String() string {
	return tea.Prettify(s)
}

func (s StartDcdnIpaDomainResponse) GoString() string {
	return s.String()
}

func (s *StartDcdnIpaDomainResponse) SetHeaders(v map[string]*string) *StartDcdnIpaDomainResponse {
	s.Headers = v
	return s
}

func (s *StartDcdnIpaDomainResponse) SetBody(v *StartDcdnIpaDomainResponseBody) *StartDcdnIpaDomainResponse {
	s.Body = v
	return s
}

type StopDcdnDomainRequest struct {
	DomainName    *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s StopDcdnDomainRequest) String() string {
	return tea.Prettify(s)
}

func (s StopDcdnDomainRequest) GoString() string {
	return s.String()
}

func (s *StopDcdnDomainRequest) SetDomainName(v string) *StopDcdnDomainRequest {
	s.DomainName = &v
	return s
}

func (s *StopDcdnDomainRequest) SetOwnerId(v int64) *StopDcdnDomainRequest {
	s.OwnerId = &v
	return s
}

func (s *StopDcdnDomainRequest) SetSecurityToken(v string) *StopDcdnDomainRequest {
	s.SecurityToken = &v
	return s
}

type StopDcdnDomainResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StopDcdnDomainResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopDcdnDomainResponseBody) GoString() string {
	return s.String()
}

func (s *StopDcdnDomainResponseBody) SetRequestId(v string) *StopDcdnDomainResponseBody {
	s.RequestId = &v
	return s
}

type StopDcdnDomainResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *StopDcdnDomainResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StopDcdnDomainResponse) String() string {
	return tea.Prettify(s)
}

func (s StopDcdnDomainResponse) GoString() string {
	return s.String()
}

func (s *StopDcdnDomainResponse) SetHeaders(v map[string]*string) *StopDcdnDomainResponse {
	s.Headers = v
	return s
}

func (s *StopDcdnDomainResponse) SetBody(v *StopDcdnDomainResponseBody) *StopDcdnDomainResponse {
	s.Body = v
	return s
}

type StopDcdnIpaDomainRequest struct {
	DomainName    *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s StopDcdnIpaDomainRequest) String() string {
	return tea.Prettify(s)
}

func (s StopDcdnIpaDomainRequest) GoString() string {
	return s.String()
}

func (s *StopDcdnIpaDomainRequest) SetDomainName(v string) *StopDcdnIpaDomainRequest {
	s.DomainName = &v
	return s
}

func (s *StopDcdnIpaDomainRequest) SetOwnerId(v int64) *StopDcdnIpaDomainRequest {
	s.OwnerId = &v
	return s
}

func (s *StopDcdnIpaDomainRequest) SetSecurityToken(v string) *StopDcdnIpaDomainRequest {
	s.SecurityToken = &v
	return s
}

type StopDcdnIpaDomainResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StopDcdnIpaDomainResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopDcdnIpaDomainResponseBody) GoString() string {
	return s.String()
}

func (s *StopDcdnIpaDomainResponseBody) SetRequestId(v string) *StopDcdnIpaDomainResponseBody {
	s.RequestId = &v
	return s
}

type StopDcdnIpaDomainResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *StopDcdnIpaDomainResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StopDcdnIpaDomainResponse) String() string {
	return tea.Prettify(s)
}

func (s StopDcdnIpaDomainResponse) GoString() string {
	return s.String()
}

func (s *StopDcdnIpaDomainResponse) SetHeaders(v map[string]*string) *StopDcdnIpaDomainResponse {
	s.Headers = v
	return s
}

func (s *StopDcdnIpaDomainResponse) SetBody(v *StopDcdnIpaDomainResponseBody) *StopDcdnIpaDomainResponse {
	s.Body = v
	return s
}

type TagDcdnResourcesRequest struct {
	OwnerId      *int64                        `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceId   []*string                     `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	ResourceType *string                       `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	Tag          []*TagDcdnResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s TagDcdnResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s TagDcdnResourcesRequest) GoString() string {
	return s.String()
}

func (s *TagDcdnResourcesRequest) SetOwnerId(v int64) *TagDcdnResourcesRequest {
	s.OwnerId = &v
	return s
}

func (s *TagDcdnResourcesRequest) SetResourceId(v []*string) *TagDcdnResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *TagDcdnResourcesRequest) SetResourceType(v string) *TagDcdnResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *TagDcdnResourcesRequest) SetTag(v []*TagDcdnResourcesRequestTag) *TagDcdnResourcesRequest {
	s.Tag = v
	return s
}

type TagDcdnResourcesRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s TagDcdnResourcesRequestTag) String() string {
	return tea.Prettify(s)
}

func (s TagDcdnResourcesRequestTag) GoString() string {
	return s.String()
}

func (s *TagDcdnResourcesRequestTag) SetKey(v string) *TagDcdnResourcesRequestTag {
	s.Key = &v
	return s
}

func (s *TagDcdnResourcesRequestTag) SetValue(v string) *TagDcdnResourcesRequestTag {
	s.Value = &v
	return s
}

type TagDcdnResourcesResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s TagDcdnResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TagDcdnResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *TagDcdnResourcesResponseBody) SetRequestId(v string) *TagDcdnResourcesResponseBody {
	s.RequestId = &v
	return s
}

type TagDcdnResourcesResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *TagDcdnResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s TagDcdnResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s TagDcdnResourcesResponse) GoString() string {
	return s.String()
}

func (s *TagDcdnResourcesResponse) SetHeaders(v map[string]*string) *TagDcdnResourcesResponse {
	s.Headers = v
	return s
}

func (s *TagDcdnResourcesResponse) SetBody(v *TagDcdnResourcesResponseBody) *TagDcdnResourcesResponse {
	s.Body = v
	return s
}

type UntagDcdnResourcesRequest struct {
	All          *bool     `json:"All,omitempty" xml:"All,omitempty"`
	OwnerId      *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceId   []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	ResourceType *string   `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	TagKey       []*string `json:"TagKey,omitempty" xml:"TagKey,omitempty" type:"Repeated"`
}

func (s UntagDcdnResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s UntagDcdnResourcesRequest) GoString() string {
	return s.String()
}

func (s *UntagDcdnResourcesRequest) SetAll(v bool) *UntagDcdnResourcesRequest {
	s.All = &v
	return s
}

func (s *UntagDcdnResourcesRequest) SetOwnerId(v int64) *UntagDcdnResourcesRequest {
	s.OwnerId = &v
	return s
}

func (s *UntagDcdnResourcesRequest) SetResourceId(v []*string) *UntagDcdnResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *UntagDcdnResourcesRequest) SetResourceType(v string) *UntagDcdnResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *UntagDcdnResourcesRequest) SetTagKey(v []*string) *UntagDcdnResourcesRequest {
	s.TagKey = v
	return s
}

type UntagDcdnResourcesResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UntagDcdnResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UntagDcdnResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *UntagDcdnResourcesResponseBody) SetRequestId(v string) *UntagDcdnResourcesResponseBody {
	s.RequestId = &v
	return s
}

type UntagDcdnResourcesResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UntagDcdnResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UntagDcdnResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s UntagDcdnResourcesResponse) GoString() string {
	return s.String()
}

func (s *UntagDcdnResourcesResponse) SetHeaders(v map[string]*string) *UntagDcdnResourcesResponse {
	s.Headers = v
	return s
}

func (s *UntagDcdnResourcesResponse) SetBody(v *UntagDcdnResourcesResponseBody) *UntagDcdnResourcesResponse {
	s.Body = v
	return s
}

type UpdateDcdnDeliverTaskRequest struct {
	Deliver    *string `json:"Deliver,omitempty" xml:"Deliver,omitempty"`
	DeliverId  *int64  `json:"DeliverId,omitempty" xml:"DeliverId,omitempty"`
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	Reports    *string `json:"Reports,omitempty" xml:"Reports,omitempty"`
	Schedule   *string `json:"Schedule,omitempty" xml:"Schedule,omitempty"`
}

func (s UpdateDcdnDeliverTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateDcdnDeliverTaskRequest) GoString() string {
	return s.String()
}

func (s *UpdateDcdnDeliverTaskRequest) SetDeliver(v string) *UpdateDcdnDeliverTaskRequest {
	s.Deliver = &v
	return s
}

func (s *UpdateDcdnDeliverTaskRequest) SetDeliverId(v int64) *UpdateDcdnDeliverTaskRequest {
	s.DeliverId = &v
	return s
}

func (s *UpdateDcdnDeliverTaskRequest) SetDomainName(v string) *UpdateDcdnDeliverTaskRequest {
	s.DomainName = &v
	return s
}

func (s *UpdateDcdnDeliverTaskRequest) SetName(v string) *UpdateDcdnDeliverTaskRequest {
	s.Name = &v
	return s
}

func (s *UpdateDcdnDeliverTaskRequest) SetOwnerId(v int64) *UpdateDcdnDeliverTaskRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateDcdnDeliverTaskRequest) SetReports(v string) *UpdateDcdnDeliverTaskRequest {
	s.Reports = &v
	return s
}

func (s *UpdateDcdnDeliverTaskRequest) SetSchedule(v string) *UpdateDcdnDeliverTaskRequest {
	s.Schedule = &v
	return s
}

type UpdateDcdnDeliverTaskResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateDcdnDeliverTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateDcdnDeliverTaskResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDcdnDeliverTaskResponseBody) SetRequestId(v string) *UpdateDcdnDeliverTaskResponseBody {
	s.RequestId = &v
	return s
}

type UpdateDcdnDeliverTaskResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateDcdnDeliverTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateDcdnDeliverTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateDcdnDeliverTaskResponse) GoString() string {
	return s.String()
}

func (s *UpdateDcdnDeliverTaskResponse) SetHeaders(v map[string]*string) *UpdateDcdnDeliverTaskResponse {
	s.Headers = v
	return s
}

func (s *UpdateDcdnDeliverTaskResponse) SetBody(v *UpdateDcdnDeliverTaskResponseBody) *UpdateDcdnDeliverTaskResponse {
	s.Body = v
	return s
}

type UpdateDcdnDomainRequest struct {
	DomainName      *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	OwnerId         *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	SecurityToken   *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	Sources         *string `json:"Sources,omitempty" xml:"Sources,omitempty"`
	TopLevelDomain  *string `json:"TopLevelDomain,omitempty" xml:"TopLevelDomain,omitempty"`
}

func (s UpdateDcdnDomainRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateDcdnDomainRequest) GoString() string {
	return s.String()
}

func (s *UpdateDcdnDomainRequest) SetDomainName(v string) *UpdateDcdnDomainRequest {
	s.DomainName = &v
	return s
}

func (s *UpdateDcdnDomainRequest) SetOwnerId(v int64) *UpdateDcdnDomainRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateDcdnDomainRequest) SetResourceGroupId(v string) *UpdateDcdnDomainRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *UpdateDcdnDomainRequest) SetSecurityToken(v string) *UpdateDcdnDomainRequest {
	s.SecurityToken = &v
	return s
}

func (s *UpdateDcdnDomainRequest) SetSources(v string) *UpdateDcdnDomainRequest {
	s.Sources = &v
	return s
}

func (s *UpdateDcdnDomainRequest) SetTopLevelDomain(v string) *UpdateDcdnDomainRequest {
	s.TopLevelDomain = &v
	return s
}

type UpdateDcdnDomainResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateDcdnDomainResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateDcdnDomainResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDcdnDomainResponseBody) SetRequestId(v string) *UpdateDcdnDomainResponseBody {
	s.RequestId = &v
	return s
}

type UpdateDcdnDomainResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateDcdnDomainResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateDcdnDomainResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateDcdnDomainResponse) GoString() string {
	return s.String()
}

func (s *UpdateDcdnDomainResponse) SetHeaders(v map[string]*string) *UpdateDcdnDomainResponse {
	s.Headers = v
	return s
}

func (s *UpdateDcdnDomainResponse) SetBody(v *UpdateDcdnDomainResponseBody) *UpdateDcdnDomainResponse {
	s.Body = v
	return s
}

type UpdateDcdnIpaDomainRequest struct {
	DomainName      *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	OwnerId         *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	SecurityToken   *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	Sources         *string `json:"Sources,omitempty" xml:"Sources,omitempty"`
	TopLevelDomain  *string `json:"TopLevelDomain,omitempty" xml:"TopLevelDomain,omitempty"`
}

func (s UpdateDcdnIpaDomainRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateDcdnIpaDomainRequest) GoString() string {
	return s.String()
}

func (s *UpdateDcdnIpaDomainRequest) SetDomainName(v string) *UpdateDcdnIpaDomainRequest {
	s.DomainName = &v
	return s
}

func (s *UpdateDcdnIpaDomainRequest) SetOwnerId(v int64) *UpdateDcdnIpaDomainRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateDcdnIpaDomainRequest) SetResourceGroupId(v string) *UpdateDcdnIpaDomainRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *UpdateDcdnIpaDomainRequest) SetSecurityToken(v string) *UpdateDcdnIpaDomainRequest {
	s.SecurityToken = &v
	return s
}

func (s *UpdateDcdnIpaDomainRequest) SetSources(v string) *UpdateDcdnIpaDomainRequest {
	s.Sources = &v
	return s
}

func (s *UpdateDcdnIpaDomainRequest) SetTopLevelDomain(v string) *UpdateDcdnIpaDomainRequest {
	s.TopLevelDomain = &v
	return s
}

type UpdateDcdnIpaDomainResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateDcdnIpaDomainResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateDcdnIpaDomainResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDcdnIpaDomainResponseBody) SetRequestId(v string) *UpdateDcdnIpaDomainResponseBody {
	s.RequestId = &v
	return s
}

type UpdateDcdnIpaDomainResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateDcdnIpaDomainResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateDcdnIpaDomainResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateDcdnIpaDomainResponse) GoString() string {
	return s.String()
}

func (s *UpdateDcdnIpaDomainResponse) SetHeaders(v map[string]*string) *UpdateDcdnIpaDomainResponse {
	s.Headers = v
	return s
}

func (s *UpdateDcdnIpaDomainResponse) SetBody(v *UpdateDcdnIpaDomainResponseBody) *UpdateDcdnIpaDomainResponse {
	s.Body = v
	return s
}

type UpdateDcdnSLSRealtimeLogDeliveryRequest struct {
	DataCenter   *string `json:"DataCenter,omitempty" xml:"DataCenter,omitempty"`
	DomainName   *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ProjectName  *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	SLSLogStore  *string `json:"SLSLogStore,omitempty" xml:"SLSLogStore,omitempty"`
	SLSProject   *string `json:"SLSProject,omitempty" xml:"SLSProject,omitempty"`
	SLSRegion    *string `json:"SLSRegion,omitempty" xml:"SLSRegion,omitempty"`
	SamplingRate *string `json:"SamplingRate,omitempty" xml:"SamplingRate,omitempty"`
}

func (s UpdateDcdnSLSRealtimeLogDeliveryRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateDcdnSLSRealtimeLogDeliveryRequest) GoString() string {
	return s.String()
}

func (s *UpdateDcdnSLSRealtimeLogDeliveryRequest) SetDataCenter(v string) *UpdateDcdnSLSRealtimeLogDeliveryRequest {
	s.DataCenter = &v
	return s
}

func (s *UpdateDcdnSLSRealtimeLogDeliveryRequest) SetDomainName(v string) *UpdateDcdnSLSRealtimeLogDeliveryRequest {
	s.DomainName = &v
	return s
}

func (s *UpdateDcdnSLSRealtimeLogDeliveryRequest) SetOwnerId(v int64) *UpdateDcdnSLSRealtimeLogDeliveryRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateDcdnSLSRealtimeLogDeliveryRequest) SetProjectName(v string) *UpdateDcdnSLSRealtimeLogDeliveryRequest {
	s.ProjectName = &v
	return s
}

func (s *UpdateDcdnSLSRealtimeLogDeliveryRequest) SetSLSLogStore(v string) *UpdateDcdnSLSRealtimeLogDeliveryRequest {
	s.SLSLogStore = &v
	return s
}

func (s *UpdateDcdnSLSRealtimeLogDeliveryRequest) SetSLSProject(v string) *UpdateDcdnSLSRealtimeLogDeliveryRequest {
	s.SLSProject = &v
	return s
}

func (s *UpdateDcdnSLSRealtimeLogDeliveryRequest) SetSLSRegion(v string) *UpdateDcdnSLSRealtimeLogDeliveryRequest {
	s.SLSRegion = &v
	return s
}

func (s *UpdateDcdnSLSRealtimeLogDeliveryRequest) SetSamplingRate(v string) *UpdateDcdnSLSRealtimeLogDeliveryRequest {
	s.SamplingRate = &v
	return s
}

type UpdateDcdnSLSRealtimeLogDeliveryResponseBody struct {
	Content   *UpdateDcdnSLSRealtimeLogDeliveryResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	RequestId *string                                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateDcdnSLSRealtimeLogDeliveryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateDcdnSLSRealtimeLogDeliveryResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDcdnSLSRealtimeLogDeliveryResponseBody) SetContent(v *UpdateDcdnSLSRealtimeLogDeliveryResponseBodyContent) *UpdateDcdnSLSRealtimeLogDeliveryResponseBody {
	s.Content = v
	return s
}

func (s *UpdateDcdnSLSRealtimeLogDeliveryResponseBody) SetRequestId(v string) *UpdateDcdnSLSRealtimeLogDeliveryResponseBody {
	s.RequestId = &v
	return s
}

type UpdateDcdnSLSRealtimeLogDeliveryResponseBodyContent struct {
	Domains []*UpdateDcdnSLSRealtimeLogDeliveryResponseBodyContentDomains `json:"Domains,omitempty" xml:"Domains,omitempty" type:"Repeated"`
}

func (s UpdateDcdnSLSRealtimeLogDeliveryResponseBodyContent) String() string {
	return tea.Prettify(s)
}

func (s UpdateDcdnSLSRealtimeLogDeliveryResponseBodyContent) GoString() string {
	return s.String()
}

func (s *UpdateDcdnSLSRealtimeLogDeliveryResponseBodyContent) SetDomains(v []*UpdateDcdnSLSRealtimeLogDeliveryResponseBodyContentDomains) *UpdateDcdnSLSRealtimeLogDeliveryResponseBodyContent {
	s.Domains = v
	return s
}

type UpdateDcdnSLSRealtimeLogDeliveryResponseBodyContentDomains struct {
	Desc       *string `json:"Desc,omitempty" xml:"Desc,omitempty"`
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	Region     *string `json:"Region,omitempty" xml:"Region,omitempty"`
	Status     *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s UpdateDcdnSLSRealtimeLogDeliveryResponseBodyContentDomains) String() string {
	return tea.Prettify(s)
}

func (s UpdateDcdnSLSRealtimeLogDeliveryResponseBodyContentDomains) GoString() string {
	return s.String()
}

func (s *UpdateDcdnSLSRealtimeLogDeliveryResponseBodyContentDomains) SetDesc(v string) *UpdateDcdnSLSRealtimeLogDeliveryResponseBodyContentDomains {
	s.Desc = &v
	return s
}

func (s *UpdateDcdnSLSRealtimeLogDeliveryResponseBodyContentDomains) SetDomainName(v string) *UpdateDcdnSLSRealtimeLogDeliveryResponseBodyContentDomains {
	s.DomainName = &v
	return s
}

func (s *UpdateDcdnSLSRealtimeLogDeliveryResponseBodyContentDomains) SetRegion(v string) *UpdateDcdnSLSRealtimeLogDeliveryResponseBodyContentDomains {
	s.Region = &v
	return s
}

func (s *UpdateDcdnSLSRealtimeLogDeliveryResponseBodyContentDomains) SetStatus(v string) *UpdateDcdnSLSRealtimeLogDeliveryResponseBodyContentDomains {
	s.Status = &v
	return s
}

type UpdateDcdnSLSRealtimeLogDeliveryResponse struct {
	Headers map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateDcdnSLSRealtimeLogDeliveryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateDcdnSLSRealtimeLogDeliveryResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateDcdnSLSRealtimeLogDeliveryResponse) GoString() string {
	return s.String()
}

func (s *UpdateDcdnSLSRealtimeLogDeliveryResponse) SetHeaders(v map[string]*string) *UpdateDcdnSLSRealtimeLogDeliveryResponse {
	s.Headers = v
	return s
}

func (s *UpdateDcdnSLSRealtimeLogDeliveryResponse) SetBody(v *UpdateDcdnSLSRealtimeLogDeliveryResponseBody) *UpdateDcdnSLSRealtimeLogDeliveryResponse {
	s.Body = v
	return s
}

type UpdateDcdnSubTaskRequest struct {
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EndTime    *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ReportIds  *string `json:"ReportIds,omitempty" xml:"ReportIds,omitempty"`
	StartTime  *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s UpdateDcdnSubTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateDcdnSubTaskRequest) GoString() string {
	return s.String()
}

func (s *UpdateDcdnSubTaskRequest) SetDomainName(v string) *UpdateDcdnSubTaskRequest {
	s.DomainName = &v
	return s
}

func (s *UpdateDcdnSubTaskRequest) SetEndTime(v string) *UpdateDcdnSubTaskRequest {
	s.EndTime = &v
	return s
}

func (s *UpdateDcdnSubTaskRequest) SetOwnerId(v int64) *UpdateDcdnSubTaskRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateDcdnSubTaskRequest) SetReportIds(v string) *UpdateDcdnSubTaskRequest {
	s.ReportIds = &v
	return s
}

func (s *UpdateDcdnSubTaskRequest) SetStartTime(v string) *UpdateDcdnSubTaskRequest {
	s.StartTime = &v
	return s
}

type UpdateDcdnSubTaskResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateDcdnSubTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateDcdnSubTaskResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDcdnSubTaskResponseBody) SetRequestId(v string) *UpdateDcdnSubTaskResponseBody {
	s.RequestId = &v
	return s
}

type UpdateDcdnSubTaskResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateDcdnSubTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateDcdnSubTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateDcdnSubTaskResponse) GoString() string {
	return s.String()
}

func (s *UpdateDcdnSubTaskResponse) SetHeaders(v map[string]*string) *UpdateDcdnSubTaskResponse {
	s.Headers = v
	return s
}

func (s *UpdateDcdnSubTaskResponse) SetBody(v *UpdateDcdnSubTaskResponseBody) *UpdateDcdnSubTaskResponse {
	s.Body = v
	return s
}

type UpdateDcdnUserRealTimeDeliveryFieldRequest struct {
	BusinessType *string `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
	Fields       *string `json:"Fields,omitempty" xml:"Fields,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s UpdateDcdnUserRealTimeDeliveryFieldRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateDcdnUserRealTimeDeliveryFieldRequest) GoString() string {
	return s.String()
}

func (s *UpdateDcdnUserRealTimeDeliveryFieldRequest) SetBusinessType(v string) *UpdateDcdnUserRealTimeDeliveryFieldRequest {
	s.BusinessType = &v
	return s
}

func (s *UpdateDcdnUserRealTimeDeliveryFieldRequest) SetFields(v string) *UpdateDcdnUserRealTimeDeliveryFieldRequest {
	s.Fields = &v
	return s
}

func (s *UpdateDcdnUserRealTimeDeliveryFieldRequest) SetOwnerId(v int64) *UpdateDcdnUserRealTimeDeliveryFieldRequest {
	s.OwnerId = &v
	return s
}

type UpdateDcdnUserRealTimeDeliveryFieldResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateDcdnUserRealTimeDeliveryFieldResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateDcdnUserRealTimeDeliveryFieldResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDcdnUserRealTimeDeliveryFieldResponseBody) SetRequestId(v string) *UpdateDcdnUserRealTimeDeliveryFieldResponseBody {
	s.RequestId = &v
	return s
}

type UpdateDcdnUserRealTimeDeliveryFieldResponse struct {
	Headers map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateDcdnUserRealTimeDeliveryFieldResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateDcdnUserRealTimeDeliveryFieldResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateDcdnUserRealTimeDeliveryFieldResponse) GoString() string {
	return s.String()
}

func (s *UpdateDcdnUserRealTimeDeliveryFieldResponse) SetHeaders(v map[string]*string) *UpdateDcdnUserRealTimeDeliveryFieldResponse {
	s.Headers = v
	return s
}

func (s *UpdateDcdnUserRealTimeDeliveryFieldResponse) SetBody(v *UpdateDcdnUserRealTimeDeliveryFieldResponseBody) *UpdateDcdnUserRealTimeDeliveryFieldResponse {
	s.Body = v
	return s
}

type UploadRoutineCodeRequest struct {
	CodeDescription *string `json:"CodeDescription,omitempty" xml:"CodeDescription,omitempty"`
	Name            *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerId         *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s UploadRoutineCodeRequest) String() string {
	return tea.Prettify(s)
}

func (s UploadRoutineCodeRequest) GoString() string {
	return s.String()
}

func (s *UploadRoutineCodeRequest) SetCodeDescription(v string) *UploadRoutineCodeRequest {
	s.CodeDescription = &v
	return s
}

func (s *UploadRoutineCodeRequest) SetName(v string) *UploadRoutineCodeRequest {
	s.Name = &v
	return s
}

func (s *UploadRoutineCodeRequest) SetOwnerId(v int64) *UploadRoutineCodeRequest {
	s.OwnerId = &v
	return s
}

type UploadRoutineCodeResponseBody struct {
	Content   map[string]interface{} `json:"Content,omitempty" xml:"Content,omitempty"`
	RequestId *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UploadRoutineCodeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UploadRoutineCodeResponseBody) GoString() string {
	return s.String()
}

func (s *UploadRoutineCodeResponseBody) SetContent(v map[string]interface{}) *UploadRoutineCodeResponseBody {
	s.Content = v
	return s
}

func (s *UploadRoutineCodeResponseBody) SetRequestId(v string) *UploadRoutineCodeResponseBody {
	s.RequestId = &v
	return s
}

type UploadRoutineCodeResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UploadRoutineCodeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UploadRoutineCodeResponse) String() string {
	return tea.Prettify(s)
}

func (s UploadRoutineCodeResponse) GoString() string {
	return s.String()
}

func (s *UploadRoutineCodeResponse) SetHeaders(v map[string]*string) *UploadRoutineCodeResponse {
	s.Headers = v
	return s
}

func (s *UploadRoutineCodeResponse) SetBody(v *UploadRoutineCodeResponseBody) *UploadRoutineCodeResponse {
	s.Body = v
	return s
}

type UploadStagingRoutineCodeRequest struct {
	CodeDescription *string `json:"CodeDescription,omitempty" xml:"CodeDescription,omitempty"`
	Name            *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerId         *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s UploadStagingRoutineCodeRequest) String() string {
	return tea.Prettify(s)
}

func (s UploadStagingRoutineCodeRequest) GoString() string {
	return s.String()
}

func (s *UploadStagingRoutineCodeRequest) SetCodeDescription(v string) *UploadStagingRoutineCodeRequest {
	s.CodeDescription = &v
	return s
}

func (s *UploadStagingRoutineCodeRequest) SetName(v string) *UploadStagingRoutineCodeRequest {
	s.Name = &v
	return s
}

func (s *UploadStagingRoutineCodeRequest) SetOwnerId(v int64) *UploadStagingRoutineCodeRequest {
	s.OwnerId = &v
	return s
}

type UploadStagingRoutineCodeResponseBody struct {
	Content   map[string]interface{} `json:"Content,omitempty" xml:"Content,omitempty"`
	RequestId *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UploadStagingRoutineCodeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UploadStagingRoutineCodeResponseBody) GoString() string {
	return s.String()
}

func (s *UploadStagingRoutineCodeResponseBody) SetContent(v map[string]interface{}) *UploadStagingRoutineCodeResponseBody {
	s.Content = v
	return s
}

func (s *UploadStagingRoutineCodeResponseBody) SetRequestId(v string) *UploadStagingRoutineCodeResponseBody {
	s.RequestId = &v
	return s
}

type UploadStagingRoutineCodeResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UploadStagingRoutineCodeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UploadStagingRoutineCodeResponse) String() string {
	return tea.Prettify(s)
}

func (s UploadStagingRoutineCodeResponse) GoString() string {
	return s.String()
}

func (s *UploadStagingRoutineCodeResponse) SetHeaders(v map[string]*string) *UploadStagingRoutineCodeResponse {
	s.Headers = v
	return s
}

func (s *UploadStagingRoutineCodeResponse) SetBody(v *UploadStagingRoutineCodeResponseBody) *UploadStagingRoutineCodeResponse {
	s.Body = v
	return s
}

type VerifyDcdnDomainOwnerRequest struct {
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	VerifyType *string `json:"VerifyType,omitempty" xml:"VerifyType,omitempty"`
}

func (s VerifyDcdnDomainOwnerRequest) String() string {
	return tea.Prettify(s)
}

func (s VerifyDcdnDomainOwnerRequest) GoString() string {
	return s.String()
}

func (s *VerifyDcdnDomainOwnerRequest) SetDomainName(v string) *VerifyDcdnDomainOwnerRequest {
	s.DomainName = &v
	return s
}

func (s *VerifyDcdnDomainOwnerRequest) SetOwnerId(v int64) *VerifyDcdnDomainOwnerRequest {
	s.OwnerId = &v
	return s
}

func (s *VerifyDcdnDomainOwnerRequest) SetVerifyType(v string) *VerifyDcdnDomainOwnerRequest {
	s.VerifyType = &v
	return s
}

type VerifyDcdnDomainOwnerResponseBody struct {
	Content   *string `json:"Content,omitempty" xml:"Content,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s VerifyDcdnDomainOwnerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s VerifyDcdnDomainOwnerResponseBody) GoString() string {
	return s.String()
}

func (s *VerifyDcdnDomainOwnerResponseBody) SetContent(v string) *VerifyDcdnDomainOwnerResponseBody {
	s.Content = &v
	return s
}

func (s *VerifyDcdnDomainOwnerResponseBody) SetRequestId(v string) *VerifyDcdnDomainOwnerResponseBody {
	s.RequestId = &v
	return s
}

type VerifyDcdnDomainOwnerResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *VerifyDcdnDomainOwnerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s VerifyDcdnDomainOwnerResponse) String() string {
	return tea.Prettify(s)
}

func (s VerifyDcdnDomainOwnerResponse) GoString() string {
	return s.String()
}

func (s *VerifyDcdnDomainOwnerResponse) SetHeaders(v map[string]*string) *VerifyDcdnDomainOwnerResponse {
	s.Headers = v
	return s
}

func (s *VerifyDcdnDomainOwnerResponse) SetBody(v *VerifyDcdnDomainOwnerResponseBody) *VerifyDcdnDomainOwnerResponse {
	s.Body = v
	return s
}

type Client struct {
	openapi.Client
}

func NewClient(config *openapi.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapi.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = tea.String("regional")
	client.EndpointMap = map[string]*string{
		"ap-northeast-1":              tea.String("dcdn.aliyuncs.com"),
		"ap-northeast-2-pop":          tea.String("dcdn.aliyuncs.com"),
		"ap-south-1":                  tea.String("dcdn.aliyuncs.com"),
		"ap-southeast-1":              tea.String("dcdn.aliyuncs.com"),
		"ap-southeast-2":              tea.String("dcdn.aliyuncs.com"),
		"ap-southeast-3":              tea.String("dcdn.aliyuncs.com"),
		"ap-southeast-5":              tea.String("dcdn.aliyuncs.com"),
		"cn-beijing":                  tea.String("dcdn.aliyuncs.com"),
		"cn-beijing-finance-1":        tea.String("dcdn.aliyuncs.com"),
		"cn-beijing-finance-pop":      tea.String("dcdn.aliyuncs.com"),
		"cn-beijing-gov-1":            tea.String("dcdn.aliyuncs.com"),
		"cn-beijing-nu16-b01":         tea.String("dcdn.aliyuncs.com"),
		"cn-chengdu":                  tea.String("dcdn.aliyuncs.com"),
		"cn-edge-1":                   tea.String("dcdn.aliyuncs.com"),
		"cn-fujian":                   tea.String("dcdn.aliyuncs.com"),
		"cn-haidian-cm12-c01":         tea.String("dcdn.aliyuncs.com"),
		"cn-hangzhou":                 tea.String("dcdn.aliyuncs.com"),
		"cn-hangzhou-bj-b01":          tea.String("dcdn.aliyuncs.com"),
		"cn-hangzhou-finance":         tea.String("dcdn.aliyuncs.com"),
		"cn-hangzhou-internal-prod-1": tea.String("dcdn.aliyuncs.com"),
		"cn-hangzhou-internal-test-1": tea.String("dcdn.aliyuncs.com"),
		"cn-hangzhou-internal-test-2": tea.String("dcdn.aliyuncs.com"),
		"cn-hangzhou-internal-test-3": tea.String("dcdn.aliyuncs.com"),
		"cn-hangzhou-test-306":        tea.String("dcdn.aliyuncs.com"),
		"cn-hongkong":                 tea.String("dcdn.aliyuncs.com"),
		"cn-hongkong-finance-pop":     tea.String("dcdn.aliyuncs.com"),
		"cn-huhehaote":                tea.String("dcdn.aliyuncs.com"),
		"cn-north-2-gov-1":            tea.String("dcdn.aliyuncs.com"),
		"cn-qingdao":                  tea.String("dcdn.aliyuncs.com"),
		"cn-qingdao-nebula":           tea.String("dcdn.aliyuncs.com"),
		"cn-shanghai":                 tea.String("dcdn.aliyuncs.com"),
		"cn-shanghai-et15-b01":        tea.String("dcdn.aliyuncs.com"),
		"cn-shanghai-et2-b01":         tea.String("dcdn.aliyuncs.com"),
		"cn-shanghai-finance-1":       tea.String("dcdn.aliyuncs.com"),
		"cn-shanghai-inner":           tea.String("dcdn.aliyuncs.com"),
		"cn-shanghai-internal-test-1": tea.String("dcdn.aliyuncs.com"),
		"cn-shenzhen":                 tea.String("dcdn.aliyuncs.com"),
		"cn-shenzhen-finance-1":       tea.String("dcdn.aliyuncs.com"),
		"cn-shenzhen-inner":           tea.String("dcdn.aliyuncs.com"),
		"cn-shenzhen-st4-d01":         tea.String("dcdn.aliyuncs.com"),
		"cn-shenzhen-su18-b01":        tea.String("dcdn.aliyuncs.com"),
		"cn-wuhan":                    tea.String("dcdn.aliyuncs.com"),
		"cn-yushanfang":               tea.String("dcdn.aliyuncs.com"),
		"cn-zhangbei-na61-b01":        tea.String("dcdn.aliyuncs.com"),
		"cn-zhangjiakou":              tea.String("dcdn.aliyuncs.com"),
		"cn-zhangjiakou-na62-a01":     tea.String("dcdn.aliyuncs.com"),
		"cn-zhengzhou-nebula-1":       tea.String("dcdn.aliyuncs.com"),
		"eu-central-1":                tea.String("dcdn.aliyuncs.com"),
		"eu-west-1":                   tea.String("dcdn.aliyuncs.com"),
		"eu-west-1-oxs":               tea.String("dcdn.aliyuncs.com"),
		"me-east-1":                   tea.String("dcdn.aliyuncs.com"),
		"rus-west-1-pop":              tea.String("dcdn.aliyuncs.com"),
		"us-east-1":                   tea.String("dcdn.aliyuncs.com"),
		"us-west-1":                   tea.String("dcdn.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("dcdn"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !tea.BoolValue(util.Empty(endpoint)) {
		_result = endpoint
		return _result, _err
	}

	if !tea.BoolValue(util.IsUnset(endpointMap)) && !tea.BoolValue(util.Empty(endpointMap[tea.StringValue(regionId)])) {
		_result = endpointMap[tea.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := endpointutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddDcdnDomainWithOptions(request *AddDcdnDomainRequest, runtime *util.RuntimeOptions) (_result *AddDcdnDomainResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CheckUrl)) {
		query["CheckUrl"] = request.CheckUrl
	}

	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.Scope)) {
		query["Scope"] = request.Scope
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !tea.BoolValue(util.IsUnset(request.Sources)) {
		query["Sources"] = request.Sources
	}

	if !tea.BoolValue(util.IsUnset(request.TopLevelDomain)) {
		query["TopLevelDomain"] = request.TopLevelDomain
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AddDcdnDomain"),
		Version:     tea.String("2018-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddDcdnDomainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddDcdnDomain(request *AddDcdnDomainRequest) (_result *AddDcdnDomainResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddDcdnDomainResponse{}
	_body, _err := client.AddDcdnDomainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddDcdnIpaDomainWithOptions(request *AddDcdnIpaDomainRequest, runtime *util.RuntimeOptions) (_result *AddDcdnIpaDomainResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CheckUrl)) {
		query["CheckUrl"] = request.CheckUrl
	}

	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.Protocol)) {
		query["Protocol"] = request.Protocol
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.Scope)) {
		query["Scope"] = request.Scope
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !tea.BoolValue(util.IsUnset(request.Sources)) {
		query["Sources"] = request.Sources
	}

	if !tea.BoolValue(util.IsUnset(request.TopLevelDomain)) {
		query["TopLevelDomain"] = request.TopLevelDomain
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AddDcdnIpaDomain"),
		Version:     tea.String("2018-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddDcdnIpaDomainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddDcdnIpaDomain(request *AddDcdnIpaDomainRequest) (_result *AddDcdnIpaDomainResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddDcdnIpaDomainResponse{}
	_body, _err := client.AddDcdnIpaDomainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BatchAddDcdnDomainWithOptions(request *BatchAddDcdnDomainRequest, runtime *util.RuntimeOptions) (_result *BatchAddDcdnDomainResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CheckUrl)) {
		query["CheckUrl"] = request.CheckUrl
	}

	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.Scope)) {
		query["Scope"] = request.Scope
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !tea.BoolValue(util.IsUnset(request.Sources)) {
		query["Sources"] = request.Sources
	}

	if !tea.BoolValue(util.IsUnset(request.TopLevelDomain)) {
		query["TopLevelDomain"] = request.TopLevelDomain
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("BatchAddDcdnDomain"),
		Version:     tea.String("2018-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &BatchAddDcdnDomainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BatchAddDcdnDomain(request *BatchAddDcdnDomainRequest) (_result *BatchAddDcdnDomainResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BatchAddDcdnDomainResponse{}
	_body, _err := client.BatchAddDcdnDomainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BatchDeleteDcdnDomainConfigsWithOptions(request *BatchDeleteDcdnDomainConfigsRequest, runtime *util.RuntimeOptions) (_result *BatchDeleteDcdnDomainConfigsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainNames)) {
		query["DomainNames"] = request.DomainNames
	}

	if !tea.BoolValue(util.IsUnset(request.FunctionNames)) {
		query["FunctionNames"] = request.FunctionNames
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("BatchDeleteDcdnDomainConfigs"),
		Version:     tea.String("2018-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &BatchDeleteDcdnDomainConfigsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BatchDeleteDcdnDomainConfigs(request *BatchDeleteDcdnDomainConfigsRequest) (_result *BatchDeleteDcdnDomainConfigsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BatchDeleteDcdnDomainConfigsResponse{}
	_body, _err := client.BatchDeleteDcdnDomainConfigsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BatchSetDcdnDomainCertificateWithOptions(request *BatchSetDcdnDomainCertificateRequest, runtime *util.RuntimeOptions) (_result *BatchSetDcdnDomainCertificateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CertName)) {
		query["CertName"] = request.CertName
	}

	if !tea.BoolValue(util.IsUnset(request.CertType)) {
		query["CertType"] = request.CertType
	}

	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.Region)) {
		query["Region"] = request.Region
	}

	if !tea.BoolValue(util.IsUnset(request.SSLPri)) {
		query["SSLPri"] = request.SSLPri
	}

	if !tea.BoolValue(util.IsUnset(request.SSLProtocol)) {
		query["SSLProtocol"] = request.SSLProtocol
	}

	if !tea.BoolValue(util.IsUnset(request.SSLPub)) {
		query["SSLPub"] = request.SSLPub
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("BatchSetDcdnDomainCertificate"),
		Version:     tea.String("2018-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &BatchSetDcdnDomainCertificateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BatchSetDcdnDomainCertificate(request *BatchSetDcdnDomainCertificateRequest) (_result *BatchSetDcdnDomainCertificateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BatchSetDcdnDomainCertificateResponse{}
	_body, _err := client.BatchSetDcdnDomainCertificateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BatchSetDcdnDomainConfigsWithOptions(request *BatchSetDcdnDomainConfigsRequest, runtime *util.RuntimeOptions) (_result *BatchSetDcdnDomainConfigsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainNames)) {
		query["DomainNames"] = request.DomainNames
	}

	if !tea.BoolValue(util.IsUnset(request.Functions)) {
		query["Functions"] = request.Functions
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("BatchSetDcdnDomainConfigs"),
		Version:     tea.String("2018-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &BatchSetDcdnDomainConfigsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BatchSetDcdnDomainConfigs(request *BatchSetDcdnDomainConfigsRequest) (_result *BatchSetDcdnDomainConfigsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BatchSetDcdnDomainConfigsResponse{}
	_body, _err := client.BatchSetDcdnDomainConfigsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BatchSetDcdnIpaDomainConfigsWithOptions(request *BatchSetDcdnIpaDomainConfigsRequest, runtime *util.RuntimeOptions) (_result *BatchSetDcdnIpaDomainConfigsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainNames)) {
		query["DomainNames"] = request.DomainNames
	}

	if !tea.BoolValue(util.IsUnset(request.Functions)) {
		query["Functions"] = request.Functions
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("BatchSetDcdnIpaDomainConfigs"),
		Version:     tea.String("2018-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &BatchSetDcdnIpaDomainConfigsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BatchSetDcdnIpaDomainConfigs(request *BatchSetDcdnIpaDomainConfigsRequest) (_result *BatchSetDcdnIpaDomainConfigsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BatchSetDcdnIpaDomainConfigsResponse{}
	_body, _err := client.BatchSetDcdnIpaDomainConfigsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BatchStartDcdnDomainWithOptions(request *BatchStartDcdnDomainRequest, runtime *util.RuntimeOptions) (_result *BatchStartDcdnDomainResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainNames)) {
		query["DomainNames"] = request.DomainNames
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("BatchStartDcdnDomain"),
		Version:     tea.String("2018-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &BatchStartDcdnDomainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BatchStartDcdnDomain(request *BatchStartDcdnDomainRequest) (_result *BatchStartDcdnDomainResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BatchStartDcdnDomainResponse{}
	_body, _err := client.BatchStartDcdnDomainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BatchStopDcdnDomainWithOptions(request *BatchStopDcdnDomainRequest, runtime *util.RuntimeOptions) (_result *BatchStopDcdnDomainResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainNames)) {
		query["DomainNames"] = request.DomainNames
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("BatchStopDcdnDomain"),
		Version:     tea.String("2018-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &BatchStopDcdnDomainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BatchStopDcdnDomain(request *BatchStopDcdnDomainRequest) (_result *BatchStopDcdnDomainResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BatchStopDcdnDomainResponse{}
	_body, _err := client.BatchStopDcdnDomainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CheckDcdnProjectExistWithOptions(request *CheckDcdnProjectExistRequest, runtime *util.RuntimeOptions) (_result *CheckDcdnProjectExistResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CheckDcdnProjectExist"),
		Version:     tea.String("2018-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CheckDcdnProjectExistResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CheckDcdnProjectExist(request *CheckDcdnProjectExistRequest) (_result *CheckDcdnProjectExistResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CheckDcdnProjectExistResponse{}
	_body, _err := client.CheckDcdnProjectExistWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CommitStagingRoutineCodeWithOptions(request *CommitStagingRoutineCodeRequest, runtime *util.RuntimeOptions) (_result *CommitStagingRoutineCodeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CodeDescription)) {
		body["CodeDescription"] = request.CodeDescription
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CommitStagingRoutineCode"),
		Version:     tea.String("2018-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CommitStagingRoutineCodeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CommitStagingRoutineCode(request *CommitStagingRoutineCodeRequest) (_result *CommitStagingRoutineCodeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CommitStagingRoutineCodeResponse{}
	_body, _err := client.CommitStagingRoutineCodeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateDcdnCertificateSigningRequestWithOptions(request *CreateDcdnCertificateSigningRequestRequest, runtime *util.RuntimeOptions) (_result *CreateDcdnCertificateSigningRequestResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.City)) {
		query["City"] = request.City
	}

	if !tea.BoolValue(util.IsUnset(request.CommonName)) {
		query["CommonName"] = request.CommonName
	}

	if !tea.BoolValue(util.IsUnset(request.Country)) {
		query["Country"] = request.Country
	}

	if !tea.BoolValue(util.IsUnset(request.Email)) {
		query["Email"] = request.Email
	}

	if !tea.BoolValue(util.IsUnset(request.Organization)) {
		query["Organization"] = request.Organization
	}

	if !tea.BoolValue(util.IsUnset(request.OrganizationUnit)) {
		query["OrganizationUnit"] = request.OrganizationUnit
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SANs)) {
		query["SANs"] = request.SANs
	}

	if !tea.BoolValue(util.IsUnset(request.State)) {
		query["State"] = request.State
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateDcdnCertificateSigningRequest"),
		Version:     tea.String("2018-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateDcdnCertificateSigningRequestResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateDcdnCertificateSigningRequest(request *CreateDcdnCertificateSigningRequestRequest) (_result *CreateDcdnCertificateSigningRequestResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateDcdnCertificateSigningRequestResponse{}
	_body, _err := client.CreateDcdnCertificateSigningRequestWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateDcdnDeliverTaskWithOptions(request *CreateDcdnDeliverTaskRequest, runtime *util.RuntimeOptions) (_result *CreateDcdnDeliverTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Deliver)) {
		body["Deliver"] = request.Deliver
	}

	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		body["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Reports)) {
		body["Reports"] = request.Reports
	}

	if !tea.BoolValue(util.IsUnset(request.Schedule)) {
		body["Schedule"] = request.Schedule
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateDcdnDeliverTask"),
		Version:     tea.String("2018-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateDcdnDeliverTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateDcdnDeliverTask(request *CreateDcdnDeliverTaskRequest) (_result *CreateDcdnDeliverTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateDcdnDeliverTaskResponse{}
	_body, _err := client.CreateDcdnDeliverTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateDcdnSLSRealTimeLogDeliveryWithOptions(request *CreateDcdnSLSRealTimeLogDeliveryRequest, runtime *util.RuntimeOptions) (_result *CreateDcdnSLSRealTimeLogDeliveryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BusinessType)) {
		body["BusinessType"] = request.BusinessType
	}

	if !tea.BoolValue(util.IsUnset(request.DataCenter)) {
		body["DataCenter"] = request.DataCenter
	}

	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		body["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectName)) {
		body["ProjectName"] = request.ProjectName
	}

	if !tea.BoolValue(util.IsUnset(request.SLSLogStore)) {
		body["SLSLogStore"] = request.SLSLogStore
	}

	if !tea.BoolValue(util.IsUnset(request.SLSProject)) {
		body["SLSProject"] = request.SLSProject
	}

	if !tea.BoolValue(util.IsUnset(request.SLSRegion)) {
		body["SLSRegion"] = request.SLSRegion
	}

	if !tea.BoolValue(util.IsUnset(request.SamplingRate)) {
		body["SamplingRate"] = request.SamplingRate
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateDcdnSLSRealTimeLogDelivery"),
		Version:     tea.String("2018-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateDcdnSLSRealTimeLogDeliveryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateDcdnSLSRealTimeLogDelivery(request *CreateDcdnSLSRealTimeLogDeliveryRequest) (_result *CreateDcdnSLSRealTimeLogDeliveryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateDcdnSLSRealTimeLogDeliveryResponse{}
	_body, _err := client.CreateDcdnSLSRealTimeLogDeliveryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateDcdnSubTaskWithOptions(request *CreateDcdnSubTaskRequest, runtime *util.RuntimeOptions) (_result *CreateDcdnSubTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		body["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.ReportIds)) {
		body["ReportIds"] = request.ReportIds
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateDcdnSubTask"),
		Version:     tea.String("2018-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateDcdnSubTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateDcdnSubTask(request *CreateDcdnSubTaskRequest) (_result *CreateDcdnSubTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateDcdnSubTaskResponse{}
	_body, _err := client.CreateDcdnSubTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateRoutineWithOptions(tmpReq *CreateRoutineRequest, runtime *util.RuntimeOptions) (_result *CreateRoutineResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateRoutineShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.EnvConf)) {
		request.EnvConfShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.EnvConf, tea.String("EnvConf"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.EnvConfShrink)) {
		body["EnvConf"] = request.EnvConfShrink
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateRoutine"),
		Version:     tea.String("2018-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateRoutineResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateRoutine(request *CreateRoutineRequest) (_result *CreateRoutineResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateRoutineResponse{}
	_body, _err := client.CreateRoutineWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateSlrAndSlsProjectWithOptions(request *CreateSlrAndSlsProjectRequest, runtime *util.RuntimeOptions) (_result *CreateSlrAndSlsProjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Region)) {
		body["Region"] = request.Region
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateSlrAndSlsProject"),
		Version:     tea.String("2018-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateSlrAndSlsProjectResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateSlrAndSlsProject(request *CreateSlrAndSlsProjectRequest) (_result *CreateSlrAndSlsProjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateSlrAndSlsProjectResponse{}
	_body, _err := client.CreateSlrAndSlsProjectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DcdnHttpRequestTestToolWithOptions(tmpReq *DcdnHttpRequestTestToolRequest, runtime *util.RuntimeOptions) (_result *DcdnHttpRequestTestToolResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &DcdnHttpRequestTestToolShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Header)) {
		request.HeaderShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Header, tea.String("Header"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Args)) {
		body["Args"] = request.Args
	}

	if !tea.BoolValue(util.IsUnset(request.Body)) {
		body["Body"] = request.Body
	}

	if !tea.BoolValue(util.IsUnset(request.HeaderShrink)) {
		body["Header"] = request.HeaderShrink
	}

	if !tea.BoolValue(util.IsUnset(request.Host)) {
		body["Host"] = request.Host
	}

	if !tea.BoolValue(util.IsUnset(request.Method)) {
		body["Method"] = request.Method
	}

	if !tea.BoolValue(util.IsUnset(request.ProxyIp)) {
		body["ProxyIp"] = request.ProxyIp
	}

	if !tea.BoolValue(util.IsUnset(request.Scheme)) {
		body["Scheme"] = request.Scheme
	}

	if !tea.BoolValue(util.IsUnset(request.Uri)) {
		body["Uri"] = request.Uri
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DcdnHttpRequestTestTool"),
		Version:     tea.String("2018-01-15"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DcdnHttpRequestTestToolResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DcdnHttpRequestTestTool(request *DcdnHttpRequestTestToolRequest) (_result *DcdnHttpRequestTestToolResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DcdnHttpRequestTestToolResponse{}
	_body, _err := client.DcdnHttpRequestTestToolWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteDcdnDeliverTaskWithOptions(request *DeleteDcdnDeliverTaskRequest, runtime *util.RuntimeOptions) (_result *DeleteDcdnDeliverTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeliverId)) {
		query["DeliverId"] = request.DeliverId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteDcdnDeliverTask"),
		Version:     tea.String("2018-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteDcdnDeliverTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteDcdnDeliverTask(request *DeleteDcdnDeliverTaskRequest) (_result *DeleteDcdnDeliverTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteDcdnDeliverTaskResponse{}
	_body, _err := client.DeleteDcdnDeliverTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteDcdnDomainWithOptions(request *DeleteDcdnDomainRequest, runtime *util.RuntimeOptions) (_result *DeleteDcdnDomainResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteDcdnDomain"),
		Version:     tea.String("2018-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteDcdnDomainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteDcdnDomain(request *DeleteDcdnDomainRequest) (_result *DeleteDcdnDomainResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteDcdnDomainResponse{}
	_body, _err := client.DeleteDcdnDomainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteDcdnIpaDomainWithOptions(request *DeleteDcdnIpaDomainRequest, runtime *util.RuntimeOptions) (_result *DeleteDcdnIpaDomainResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteDcdnIpaDomain"),
		Version:     tea.String("2018-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteDcdnIpaDomainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteDcdnIpaDomain(request *DeleteDcdnIpaDomainRequest) (_result *DeleteDcdnIpaDomainResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteDcdnIpaDomainResponse{}
	_body, _err := client.DeleteDcdnIpaDomainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteDcdnIpaSpecificConfigWithOptions(request *DeleteDcdnIpaSpecificConfigRequest, runtime *util.RuntimeOptions) (_result *DeleteDcdnIpaSpecificConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConfigId)) {
		query["ConfigId"] = request.ConfigId
	}

	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteDcdnIpaSpecificConfig"),
		Version:     tea.String("2018-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteDcdnIpaSpecificConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteDcdnIpaSpecificConfig(request *DeleteDcdnIpaSpecificConfigRequest) (_result *DeleteDcdnIpaSpecificConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteDcdnIpaSpecificConfigResponse{}
	_body, _err := client.DeleteDcdnIpaSpecificConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteDcdnRealTimeLogProjectWithOptions(request *DeleteDcdnRealTimeLogProjectRequest, runtime *util.RuntimeOptions) (_result *DeleteDcdnRealTimeLogProjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BusinessType)) {
		query["BusinessType"] = request.BusinessType
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectName)) {
		query["ProjectName"] = request.ProjectName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteDcdnRealTimeLogProject"),
		Version:     tea.String("2018-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteDcdnRealTimeLogProjectResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteDcdnRealTimeLogProject(request *DeleteDcdnRealTimeLogProjectRequest) (_result *DeleteDcdnRealTimeLogProjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteDcdnRealTimeLogProjectResponse{}
	_body, _err := client.DeleteDcdnRealTimeLogProjectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteDcdnSpecificConfigWithOptions(request *DeleteDcdnSpecificConfigRequest, runtime *util.RuntimeOptions) (_result *DeleteDcdnSpecificConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConfigId)) {
		query["ConfigId"] = request.ConfigId
	}

	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteDcdnSpecificConfig"),
		Version:     tea.String("2018-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteDcdnSpecificConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteDcdnSpecificConfig(request *DeleteDcdnSpecificConfigRequest) (_result *DeleteDcdnSpecificConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteDcdnSpecificConfigResponse{}
	_body, _err := client.DeleteDcdnSpecificConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteDcdnSpecificStagingConfigWithOptions(request *DeleteDcdnSpecificStagingConfigRequest, runtime *util.RuntimeOptions) (_result *DeleteDcdnSpecificStagingConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConfigId)) {
		query["ConfigId"] = request.ConfigId
	}

	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteDcdnSpecificStagingConfig"),
		Version:     tea.String("2018-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteDcdnSpecificStagingConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteDcdnSpecificStagingConfig(request *DeleteDcdnSpecificStagingConfigRequest) (_result *DeleteDcdnSpecificStagingConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteDcdnSpecificStagingConfigResponse{}
	_body, _err := client.DeleteDcdnSpecificStagingConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteDcdnSubTaskWithOptions(request *DeleteDcdnSubTaskRequest, runtime *util.RuntimeOptions) (_result *DeleteDcdnSubTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteDcdnSubTask"),
		Version:     tea.String("2018-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteDcdnSubTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteDcdnSubTask(request *DeleteDcdnSubTaskRequest) (_result *DeleteDcdnSubTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteDcdnSubTaskResponse{}
	_body, _err := client.DeleteDcdnSubTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteRoutineWithOptions(request *DeleteRoutineRequest, runtime *util.RuntimeOptions) (_result *DeleteRoutineResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteRoutine"),
		Version:     tea.String("2018-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteRoutineResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteRoutine(request *DeleteRoutineRequest) (_result *DeleteRoutineResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteRoutineResponse{}
	_body, _err := client.DeleteRoutineWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteRoutineCodeRevisionWithOptions(request *DeleteRoutineCodeRevisionRequest, runtime *util.RuntimeOptions) (_result *DeleteRoutineCodeRevisionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.SelectCodeRevision)) {
		body["SelectCodeRevision"] = request.SelectCodeRevision
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteRoutineCodeRevision"),
		Version:     tea.String("2018-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteRoutineCodeRevisionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteRoutineCodeRevision(request *DeleteRoutineCodeRevisionRequest) (_result *DeleteRoutineCodeRevisionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteRoutineCodeRevisionResponse{}
	_body, _err := client.DeleteRoutineCodeRevisionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteRoutineConfEnvsWithOptions(tmpReq *DeleteRoutineConfEnvsRequest, runtime *util.RuntimeOptions) (_result *DeleteRoutineConfEnvsResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &DeleteRoutineConfEnvsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Envs)) {
		request.EnvsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Envs, tea.String("Envs"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EnvsShrink)) {
		body["Envs"] = request.EnvsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteRoutineConfEnvs"),
		Version:     tea.String("2018-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteRoutineConfEnvsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteRoutineConfEnvs(request *DeleteRoutineConfEnvsRequest) (_result *DeleteRoutineConfEnvsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteRoutineConfEnvsResponse{}
	_body, _err := client.DeleteRoutineConfEnvsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDcdnAclFieldsWithOptions(request *DescribeDcdnAclFieldsRequest, runtime *util.RuntimeOptions) (_result *DescribeDcdnAclFieldsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDcdnAclFields"),
		Version:     tea.String("2018-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDcdnAclFieldsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDcdnAclFields(request *DescribeDcdnAclFieldsRequest) (_result *DescribeDcdnAclFieldsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDcdnAclFieldsResponse{}
	_body, _err := client.DescribeDcdnAclFieldsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDcdnBgpBpsDataWithOptions(request *DescribeDcdnBgpBpsDataRequest, runtime *util.RuntimeOptions) (_result *DescribeDcdnBgpBpsDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Interval)) {
		query["Interval"] = request.Interval
	}

	if !tea.BoolValue(util.IsUnset(request.Isp)) {
		query["Isp"] = request.Isp
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDcdnBgpBpsData"),
		Version:     tea.String("2018-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDcdnBgpBpsDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDcdnBgpBpsData(request *DescribeDcdnBgpBpsDataRequest) (_result *DescribeDcdnBgpBpsDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDcdnBgpBpsDataResponse{}
	_body, _err := client.DescribeDcdnBgpBpsDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDcdnBgpTrafficDataWithOptions(request *DescribeDcdnBgpTrafficDataRequest, runtime *util.RuntimeOptions) (_result *DescribeDcdnBgpTrafficDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Interval)) {
		query["Interval"] = request.Interval
	}

	if !tea.BoolValue(util.IsUnset(request.Isp)) {
		query["Isp"] = request.Isp
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDcdnBgpTrafficData"),
		Version:     tea.String("2018-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDcdnBgpTrafficDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDcdnBgpTrafficData(request *DescribeDcdnBgpTrafficDataRequest) (_result *DescribeDcdnBgpTrafficDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDcdnBgpTrafficDataResponse{}
	_body, _err := client.DescribeDcdnBgpTrafficDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDcdnBlockedRegionsWithOptions(request *DescribeDcdnBlockedRegionsRequest, runtime *util.RuntimeOptions) (_result *DescribeDcdnBlockedRegionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDcdnBlockedRegions"),
		Version:     tea.String("2018-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDcdnBlockedRegionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDcdnBlockedRegions(request *DescribeDcdnBlockedRegionsRequest) (_result *DescribeDcdnBlockedRegionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDcdnBlockedRegionsResponse{}
	_body, _err := client.DescribeDcdnBlockedRegionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDcdnCertificateDetailWithOptions(request *DescribeDcdnCertificateDetailRequest, runtime *util.RuntimeOptions) (_result *DescribeDcdnCertificateDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CertName)) {
		query["CertName"] = request.CertName
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDcdnCertificateDetail"),
		Version:     tea.String("2018-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDcdnCertificateDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDcdnCertificateDetail(request *DescribeDcdnCertificateDetailRequest) (_result *DescribeDcdnCertificateDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDcdnCertificateDetailResponse{}
	_body, _err := client.DescribeDcdnCertificateDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDcdnCertificateListWithOptions(request *DescribeDcdnCertificateListRequest, runtime *util.RuntimeOptions) (_result *DescribeDcdnCertificateListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDcdnCertificateList"),
		Version:     tea.String("2018-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDcdnCertificateListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDcdnCertificateList(request *DescribeDcdnCertificateListRequest) (_result *DescribeDcdnCertificateListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDcdnCertificateListResponse{}
	_body, _err := client.DescribeDcdnCertificateListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDcdnConfigGroupDetailWithOptions(request *DescribeDcdnConfigGroupDetailRequest, runtime *util.RuntimeOptions) (_result *DescribeDcdnConfigGroupDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConfigGroupId)) {
		query["ConfigGroupId"] = request.ConfigGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ConfigGroupName)) {
		query["ConfigGroupName"] = request.ConfigGroupName
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDcdnConfigGroupDetail"),
		Version:     tea.String("2018-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDcdnConfigGroupDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDcdnConfigGroupDetail(request *DescribeDcdnConfigGroupDetailRequest) (_result *DescribeDcdnConfigGroupDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDcdnConfigGroupDetailResponse{}
	_body, _err := client.DescribeDcdnConfigGroupDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDcdnConfigOfVersionWithOptions(request *DescribeDcdnConfigOfVersionRequest, runtime *util.RuntimeOptions) (_result *DescribeDcdnConfigOfVersionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FunctionId)) {
		query["FunctionId"] = request.FunctionId
	}

	if !tea.BoolValue(util.IsUnset(request.FunctionName)) {
		query["FunctionName"] = request.FunctionName
	}

	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		query["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !tea.BoolValue(util.IsUnset(request.VersionId)) {
		query["VersionId"] = request.VersionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDcdnConfigOfVersion"),
		Version:     tea.String("2018-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDcdnConfigOfVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDcdnConfigOfVersion(request *DescribeDcdnConfigOfVersionRequest) (_result *DescribeDcdnConfigOfVersionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDcdnConfigOfVersionResponse{}
	_body, _err := client.DescribeDcdnConfigOfVersionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDcdnDeletedDomainsWithOptions(request *DescribeDcdnDeletedDomainsRequest, runtime *util.RuntimeOptions) (_result *DescribeDcdnDeletedDomainsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDcdnDeletedDomains"),
		Version:     tea.String("2018-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDcdnDeletedDomainsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDcdnDeletedDomains(request *DescribeDcdnDeletedDomainsRequest) (_result *DescribeDcdnDeletedDomainsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDcdnDeletedDomainsResponse{}
	_body, _err := client.DescribeDcdnDeletedDomainsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDcdnDeliverListWithOptions(request *DescribeDcdnDeliverListRequest, runtime *util.RuntimeOptions) (_result *DescribeDcdnDeliverListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeliverId)) {
		query["DeliverId"] = request.DeliverId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDcdnDeliverList"),
		Version:     tea.String("2018-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDcdnDeliverListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDcdnDeliverList(request *DescribeDcdnDeliverListRequest) (_result *DescribeDcdnDeliverListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDcdnDeliverListResponse{}
	_body, _err := client.DescribeDcdnDeliverListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDcdnDomainBpsDataWithOptions(request *DescribeDcdnDomainBpsDataRequest, runtime *util.RuntimeOptions) (_result *DescribeDcdnDomainBpsDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Interval)) {
		query["Interval"] = request.Interval
	}

	if !tea.BoolValue(util.IsUnset(request.IspNameEn)) {
		query["IspNameEn"] = request.IspNameEn
	}

	if !tea.BoolValue(util.IsUnset(request.LocationNameEn)) {
		query["LocationNameEn"] = request.LocationNameEn
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDcdnDomainBpsData"),
		Version:     tea.String("2018-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDcdnDomainBpsDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDcdnDomainBpsData(request *DescribeDcdnDomainBpsDataRequest) (_result *DescribeDcdnDomainBpsDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDcdnDomainBpsDataResponse{}
	_body, _err := client.DescribeDcdnDomainBpsDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDcdnDomainByCertificateWithOptions(request *DescribeDcdnDomainByCertificateRequest, runtime *util.RuntimeOptions) (_result *DescribeDcdnDomainByCertificateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SSLPub)) {
		query["SSLPub"] = request.SSLPub
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDcdnDomainByCertificate"),
		Version:     tea.String("2018-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDcdnDomainByCertificateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDcdnDomainByCertificate(request *DescribeDcdnDomainByCertificateRequest) (_result *DescribeDcdnDomainByCertificateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDcdnDomainByCertificateResponse{}
	_body, _err := client.DescribeDcdnDomainByCertificateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDcdnDomainCcActivityLogWithOptions(request *DescribeDcdnDomainCcActivityLogRequest, runtime *util.RuntimeOptions) (_result *DescribeDcdnDomainCcActivityLogResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RuleName)) {
		query["RuleName"] = request.RuleName
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.TriggerObject)) {
		query["TriggerObject"] = request.TriggerObject
	}

	if !tea.BoolValue(util.IsUnset(request.Value)) {
		query["Value"] = request.Value
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDcdnDomainCcActivityLog"),
		Version:     tea.String("2018-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDcdnDomainCcActivityLogResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDcdnDomainCcActivityLog(request *DescribeDcdnDomainCcActivityLogRequest) (_result *DescribeDcdnDomainCcActivityLogResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDcdnDomainCcActivityLogResponse{}
	_body, _err := client.DescribeDcdnDomainCcActivityLogWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDcdnDomainCertificateInfoWithOptions(request *DescribeDcdnDomainCertificateInfoRequest, runtime *util.RuntimeOptions) (_result *DescribeDcdnDomainCertificateInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDcdnDomainCertificateInfo"),
		Version:     tea.String("2018-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDcdnDomainCertificateInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDcdnDomainCertificateInfo(request *DescribeDcdnDomainCertificateInfoRequest) (_result *DescribeDcdnDomainCertificateInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDcdnDomainCertificateInfoResponse{}
	_body, _err := client.DescribeDcdnDomainCertificateInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDcdnDomainCnameWithOptions(request *DescribeDcdnDomainCnameRequest, runtime *util.RuntimeOptions) (_result *DescribeDcdnDomainCnameResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDcdnDomainCname"),
		Version:     tea.String("2018-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDcdnDomainCnameResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDcdnDomainCname(request *DescribeDcdnDomainCnameRequest) (_result *DescribeDcdnDomainCnameResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDcdnDomainCnameResponse{}
	_body, _err := client.DescribeDcdnDomainCnameWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDcdnDomainConfigsWithOptions(request *DescribeDcdnDomainConfigsRequest, runtime *util.RuntimeOptions) (_result *DescribeDcdnDomainConfigsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConfigId)) {
		query["ConfigId"] = request.ConfigId
	}

	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.FunctionNames)) {
		query["FunctionNames"] = request.FunctionNames
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDcdnDomainConfigs"),
		Version:     tea.String("2018-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDcdnDomainConfigsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDcdnDomainConfigs(request *DescribeDcdnDomainConfigsRequest) (_result *DescribeDcdnDomainConfigsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDcdnDomainConfigsResponse{}
	_body, _err := client.DescribeDcdnDomainConfigsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDcdnDomainDetailWithOptions(request *DescribeDcdnDomainDetailRequest, runtime *util.RuntimeOptions) (_result *DescribeDcdnDomainDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDcdnDomainDetail"),
		Version:     tea.String("2018-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDcdnDomainDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDcdnDomainDetail(request *DescribeDcdnDomainDetailRequest) (_result *DescribeDcdnDomainDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDcdnDomainDetailResponse{}
	_body, _err := client.DescribeDcdnDomainDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDcdnDomainHitRateDataWithOptions(request *DescribeDcdnDomainHitRateDataRequest, runtime *util.RuntimeOptions) (_result *DescribeDcdnDomainHitRateDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Interval)) {
		query["Interval"] = request.Interval
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDcdnDomainHitRateData"),
		Version:     tea.String("2018-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDcdnDomainHitRateDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDcdnDomainHitRateData(request *DescribeDcdnDomainHitRateDataRequest) (_result *DescribeDcdnDomainHitRateDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDcdnDomainHitRateDataResponse{}
	_body, _err := client.DescribeDcdnDomainHitRateDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDcdnDomainHttpCodeDataWithOptions(request *DescribeDcdnDomainHttpCodeDataRequest, runtime *util.RuntimeOptions) (_result *DescribeDcdnDomainHttpCodeDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Interval)) {
		query["Interval"] = request.Interval
	}

	if !tea.BoolValue(util.IsUnset(request.IspNameEn)) {
		query["IspNameEn"] = request.IspNameEn
	}

	if !tea.BoolValue(util.IsUnset(request.LocationNameEn)) {
		query["LocationNameEn"] = request.LocationNameEn
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDcdnDomainHttpCodeData"),
		Version:     tea.String("2018-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDcdnDomainHttpCodeDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDcdnDomainHttpCodeData(request *DescribeDcdnDomainHttpCodeDataRequest) (_result *DescribeDcdnDomainHttpCodeDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDcdnDomainHttpCodeDataResponse{}
	_body, _err := client.DescribeDcdnDomainHttpCodeDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDcdnDomainIpaBpsDataWithOptions(request *DescribeDcdnDomainIpaBpsDataRequest, runtime *util.RuntimeOptions) (_result *DescribeDcdnDomainIpaBpsDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.FixTimeGap)) {
		query["FixTimeGap"] = request.FixTimeGap
	}

	if !tea.BoolValue(util.IsUnset(request.Interval)) {
		query["Interval"] = request.Interval
	}

	if !tea.BoolValue(util.IsUnset(request.IspNameEn)) {
		query["IspNameEn"] = request.IspNameEn
	}

	if !tea.BoolValue(util.IsUnset(request.LocationNameEn)) {
		query["LocationNameEn"] = request.LocationNameEn
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.TimeMerge)) {
		query["TimeMerge"] = request.TimeMerge
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDcdnDomainIpaBpsData"),
		Version:     tea.String("2018-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDcdnDomainIpaBpsDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDcdnDomainIpaBpsData(request *DescribeDcdnDomainIpaBpsDataRequest) (_result *DescribeDcdnDomainIpaBpsDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDcdnDomainIpaBpsDataResponse{}
	_body, _err := client.DescribeDcdnDomainIpaBpsDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDcdnDomainIpaTrafficDataWithOptions(request *DescribeDcdnDomainIpaTrafficDataRequest, runtime *util.RuntimeOptions) (_result *DescribeDcdnDomainIpaTrafficDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.FixTimeGap)) {
		query["FixTimeGap"] = request.FixTimeGap
	}

	if !tea.BoolValue(util.IsUnset(request.Interval)) {
		query["Interval"] = request.Interval
	}

	if !tea.BoolValue(util.IsUnset(request.IspNameEn)) {
		query["IspNameEn"] = request.IspNameEn
	}

	if !tea.BoolValue(util.IsUnset(request.LocationNameEn)) {
		query["LocationNameEn"] = request.LocationNameEn
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.TimeMerge)) {
		query["TimeMerge"] = request.TimeMerge
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDcdnDomainIpaTrafficData"),
		Version:     tea.String("2018-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDcdnDomainIpaTrafficDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDcdnDomainIpaTrafficData(request *DescribeDcdnDomainIpaTrafficDataRequest) (_result *DescribeDcdnDomainIpaTrafficDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDcdnDomainIpaTrafficDataResponse{}
	_body, _err := client.DescribeDcdnDomainIpaTrafficDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDcdnDomainIspDataWithOptions(request *DescribeDcdnDomainIspDataRequest, runtime *util.RuntimeOptions) (_result *DescribeDcdnDomainIspDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDcdnDomainIspData"),
		Version:     tea.String("2018-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDcdnDomainIspDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDcdnDomainIspData(request *DescribeDcdnDomainIspDataRequest) (_result *DescribeDcdnDomainIspDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDcdnDomainIspDataResponse{}
	_body, _err := client.DescribeDcdnDomainIspDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDcdnDomainLogWithOptions(request *DescribeDcdnDomainLogRequest, runtime *util.RuntimeOptions) (_result *DescribeDcdnDomainLogResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDcdnDomainLog"),
		Version:     tea.String("2018-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDcdnDomainLogResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDcdnDomainLog(request *DescribeDcdnDomainLogRequest) (_result *DescribeDcdnDomainLogResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDcdnDomainLogResponse{}
	_body, _err := client.DescribeDcdnDomainLogWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDcdnDomainMultiUsageDataWithOptions(request *DescribeDcdnDomainMultiUsageDataRequest, runtime *util.RuntimeOptions) (_result *DescribeDcdnDomainMultiUsageDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDcdnDomainMultiUsageData"),
		Version:     tea.String("2018-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDcdnDomainMultiUsageDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDcdnDomainMultiUsageData(request *DescribeDcdnDomainMultiUsageDataRequest) (_result *DescribeDcdnDomainMultiUsageDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDcdnDomainMultiUsageDataResponse{}
	_body, _err := client.DescribeDcdnDomainMultiUsageDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDcdnDomainOriginBpsDataWithOptions(request *DescribeDcdnDomainOriginBpsDataRequest, runtime *util.RuntimeOptions) (_result *DescribeDcdnDomainOriginBpsDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Interval)) {
		query["Interval"] = request.Interval
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDcdnDomainOriginBpsData"),
		Version:     tea.String("2018-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDcdnDomainOriginBpsDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDcdnDomainOriginBpsData(request *DescribeDcdnDomainOriginBpsDataRequest) (_result *DescribeDcdnDomainOriginBpsDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDcdnDomainOriginBpsDataResponse{}
	_body, _err := client.DescribeDcdnDomainOriginBpsDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDcdnDomainOriginTrafficDataWithOptions(request *DescribeDcdnDomainOriginTrafficDataRequest, runtime *util.RuntimeOptions) (_result *DescribeDcdnDomainOriginTrafficDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Interval)) {
		query["Interval"] = request.Interval
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDcdnDomainOriginTrafficData"),
		Version:     tea.String("2018-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDcdnDomainOriginTrafficDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDcdnDomainOriginTrafficData(request *DescribeDcdnDomainOriginTrafficDataRequest) (_result *DescribeDcdnDomainOriginTrafficDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDcdnDomainOriginTrafficDataResponse{}
	_body, _err := client.DescribeDcdnDomainOriginTrafficDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDcdnDomainPropertyWithOptions(request *DescribeDcdnDomainPropertyRequest, runtime *util.RuntimeOptions) (_result *DescribeDcdnDomainPropertyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDcdnDomainProperty"),
		Version:     tea.String("2018-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDcdnDomainPropertyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDcdnDomainProperty(request *DescribeDcdnDomainPropertyRequest) (_result *DescribeDcdnDomainPropertyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDcdnDomainPropertyResponse{}
	_body, _err := client.DescribeDcdnDomainPropertyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDcdnDomainPvDataWithOptions(request *DescribeDcdnDomainPvDataRequest, runtime *util.RuntimeOptions) (_result *DescribeDcdnDomainPvDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDcdnDomainPvData"),
		Version:     tea.String("2018-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDcdnDomainPvDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDcdnDomainPvData(request *DescribeDcdnDomainPvDataRequest) (_result *DescribeDcdnDomainPvDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDcdnDomainPvDataResponse{}
	_body, _err := client.DescribeDcdnDomainPvDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDcdnDomainQpsDataWithOptions(request *DescribeDcdnDomainQpsDataRequest, runtime *util.RuntimeOptions) (_result *DescribeDcdnDomainQpsDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Interval)) {
		query["Interval"] = request.Interval
	}

	if !tea.BoolValue(util.IsUnset(request.IspNameEn)) {
		query["IspNameEn"] = request.IspNameEn
	}

	if !tea.BoolValue(util.IsUnset(request.LocationNameEn)) {
		query["LocationNameEn"] = request.LocationNameEn
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDcdnDomainQpsData"),
		Version:     tea.String("2018-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDcdnDomainQpsDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDcdnDomainQpsData(request *DescribeDcdnDomainQpsDataRequest) (_result *DescribeDcdnDomainQpsDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDcdnDomainQpsDataResponse{}
	_body, _err := client.DescribeDcdnDomainQpsDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDcdnDomainRealTimeBpsDataWithOptions(request *DescribeDcdnDomainRealTimeBpsDataRequest, runtime *util.RuntimeOptions) (_result *DescribeDcdnDomainRealTimeBpsDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDcdnDomainRealTimeBpsData"),
		Version:     tea.String("2018-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDcdnDomainRealTimeBpsDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDcdnDomainRealTimeBpsData(request *DescribeDcdnDomainRealTimeBpsDataRequest) (_result *DescribeDcdnDomainRealTimeBpsDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDcdnDomainRealTimeBpsDataResponse{}
	_body, _err := client.DescribeDcdnDomainRealTimeBpsDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDcdnDomainRealTimeByteHitRateDataWithOptions(request *DescribeDcdnDomainRealTimeByteHitRateDataRequest, runtime *util.RuntimeOptions) (_result *DescribeDcdnDomainRealTimeByteHitRateDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDcdnDomainRealTimeByteHitRateData"),
		Version:     tea.String("2018-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDcdnDomainRealTimeByteHitRateDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDcdnDomainRealTimeByteHitRateData(request *DescribeDcdnDomainRealTimeByteHitRateDataRequest) (_result *DescribeDcdnDomainRealTimeByteHitRateDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDcdnDomainRealTimeByteHitRateDataResponse{}
	_body, _err := client.DescribeDcdnDomainRealTimeByteHitRateDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDcdnDomainRealTimeDetailDataWithOptions(request *DescribeDcdnDomainRealTimeDetailDataRequest, runtime *util.RuntimeOptions) (_result *DescribeDcdnDomainRealTimeDetailDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDcdnDomainRealTimeDetailData"),
		Version:     tea.String("2018-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDcdnDomainRealTimeDetailDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDcdnDomainRealTimeDetailData(request *DescribeDcdnDomainRealTimeDetailDataRequest) (_result *DescribeDcdnDomainRealTimeDetailDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDcdnDomainRealTimeDetailDataResponse{}
	_body, _err := client.DescribeDcdnDomainRealTimeDetailDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDcdnDomainRealTimeHttpCodeDataWithOptions(request *DescribeDcdnDomainRealTimeHttpCodeDataRequest, runtime *util.RuntimeOptions) (_result *DescribeDcdnDomainRealTimeHttpCodeDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.IspNameEn)) {
		query["IspNameEn"] = request.IspNameEn
	}

	if !tea.BoolValue(util.IsUnset(request.LocationNameEn)) {
		query["LocationNameEn"] = request.LocationNameEn
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDcdnDomainRealTimeHttpCodeData"),
		Version:     tea.String("2018-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDcdnDomainRealTimeHttpCodeDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDcdnDomainRealTimeHttpCodeData(request *DescribeDcdnDomainRealTimeHttpCodeDataRequest) (_result *DescribeDcdnDomainRealTimeHttpCodeDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDcdnDomainRealTimeHttpCodeDataResponse{}
	_body, _err := client.DescribeDcdnDomainRealTimeHttpCodeDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDcdnDomainRealTimeQpsDataWithOptions(request *DescribeDcdnDomainRealTimeQpsDataRequest, runtime *util.RuntimeOptions) (_result *DescribeDcdnDomainRealTimeQpsDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDcdnDomainRealTimeQpsData"),
		Version:     tea.String("2018-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDcdnDomainRealTimeQpsDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDcdnDomainRealTimeQpsData(request *DescribeDcdnDomainRealTimeQpsDataRequest) (_result *DescribeDcdnDomainRealTimeQpsDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDcdnDomainRealTimeQpsDataResponse{}
	_body, _err := client.DescribeDcdnDomainRealTimeQpsDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDcdnDomainRealTimeReqHitRateDataWithOptions(request *DescribeDcdnDomainRealTimeReqHitRateDataRequest, runtime *util.RuntimeOptions) (_result *DescribeDcdnDomainRealTimeReqHitRateDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDcdnDomainRealTimeReqHitRateData"),
		Version:     tea.String("2018-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDcdnDomainRealTimeReqHitRateDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDcdnDomainRealTimeReqHitRateData(request *DescribeDcdnDomainRealTimeReqHitRateDataRequest) (_result *DescribeDcdnDomainRealTimeReqHitRateDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDcdnDomainRealTimeReqHitRateDataResponse{}
	_body, _err := client.DescribeDcdnDomainRealTimeReqHitRateDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDcdnDomainRealTimeSrcBpsDataWithOptions(request *DescribeDcdnDomainRealTimeSrcBpsDataRequest, runtime *util.RuntimeOptions) (_result *DescribeDcdnDomainRealTimeSrcBpsDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDcdnDomainRealTimeSrcBpsData"),
		Version:     tea.String("2018-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDcdnDomainRealTimeSrcBpsDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDcdnDomainRealTimeSrcBpsData(request *DescribeDcdnDomainRealTimeSrcBpsDataRequest) (_result *DescribeDcdnDomainRealTimeSrcBpsDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDcdnDomainRealTimeSrcBpsDataResponse{}
	_body, _err := client.DescribeDcdnDomainRealTimeSrcBpsDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDcdnDomainRealTimeSrcHttpCodeDataWithOptions(request *DescribeDcdnDomainRealTimeSrcHttpCodeDataRequest, runtime *util.RuntimeOptions) (_result *DescribeDcdnDomainRealTimeSrcHttpCodeDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.IspNameEn)) {
		query["IspNameEn"] = request.IspNameEn
	}

	if !tea.BoolValue(util.IsUnset(request.LocationNameEn)) {
		query["LocationNameEn"] = request.LocationNameEn
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDcdnDomainRealTimeSrcHttpCodeData"),
		Version:     tea.String("2018-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDcdnDomainRealTimeSrcHttpCodeDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDcdnDomainRealTimeSrcHttpCodeData(request *DescribeDcdnDomainRealTimeSrcHttpCodeDataRequest) (_result *DescribeDcdnDomainRealTimeSrcHttpCodeDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDcdnDomainRealTimeSrcHttpCodeDataResponse{}
	_body, _err := client.DescribeDcdnDomainRealTimeSrcHttpCodeDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDcdnDomainRealTimeSrcTrafficDataWithOptions(request *DescribeDcdnDomainRealTimeSrcTrafficDataRequest, runtime *util.RuntimeOptions) (_result *DescribeDcdnDomainRealTimeSrcTrafficDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDcdnDomainRealTimeSrcTrafficData"),
		Version:     tea.String("2018-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDcdnDomainRealTimeSrcTrafficDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDcdnDomainRealTimeSrcTrafficData(request *DescribeDcdnDomainRealTimeSrcTrafficDataRequest) (_result *DescribeDcdnDomainRealTimeSrcTrafficDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDcdnDomainRealTimeSrcTrafficDataResponse{}
	_body, _err := client.DescribeDcdnDomainRealTimeSrcTrafficDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDcdnDomainRealTimeTrafficDataWithOptions(request *DescribeDcdnDomainRealTimeTrafficDataRequest, runtime *util.RuntimeOptions) (_result *DescribeDcdnDomainRealTimeTrafficDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDcdnDomainRealTimeTrafficData"),
		Version:     tea.String("2018-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDcdnDomainRealTimeTrafficDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDcdnDomainRealTimeTrafficData(request *DescribeDcdnDomainRealTimeTrafficDataRequest) (_result *DescribeDcdnDomainRealTimeTrafficDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDcdnDomainRealTimeTrafficDataResponse{}
	_body, _err := client.DescribeDcdnDomainRealTimeTrafficDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDcdnDomainRegionDataWithOptions(request *DescribeDcdnDomainRegionDataRequest, runtime *util.RuntimeOptions) (_result *DescribeDcdnDomainRegionDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDcdnDomainRegionData"),
		Version:     tea.String("2018-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDcdnDomainRegionDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDcdnDomainRegionData(request *DescribeDcdnDomainRegionDataRequest) (_result *DescribeDcdnDomainRegionDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDcdnDomainRegionDataResponse{}
	_body, _err := client.DescribeDcdnDomainRegionDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDcdnDomainStagingConfigWithOptions(request *DescribeDcdnDomainStagingConfigRequest, runtime *util.RuntimeOptions) (_result *DescribeDcdnDomainStagingConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.FunctionNames)) {
		query["FunctionNames"] = request.FunctionNames
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDcdnDomainStagingConfig"),
		Version:     tea.String("2018-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDcdnDomainStagingConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDcdnDomainStagingConfig(request *DescribeDcdnDomainStagingConfigRequest) (_result *DescribeDcdnDomainStagingConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDcdnDomainStagingConfigResponse{}
	_body, _err := client.DescribeDcdnDomainStagingConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDcdnDomainTopReferVisitWithOptions(request *DescribeDcdnDomainTopReferVisitRequest, runtime *util.RuntimeOptions) (_result *DescribeDcdnDomainTopReferVisitResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SortBy)) {
		query["SortBy"] = request.SortBy
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDcdnDomainTopReferVisit"),
		Version:     tea.String("2018-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDcdnDomainTopReferVisitResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDcdnDomainTopReferVisit(request *DescribeDcdnDomainTopReferVisitRequest) (_result *DescribeDcdnDomainTopReferVisitResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDcdnDomainTopReferVisitResponse{}
	_body, _err := client.DescribeDcdnDomainTopReferVisitWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDcdnDomainTopUrlVisitWithOptions(request *DescribeDcdnDomainTopUrlVisitRequest, runtime *util.RuntimeOptions) (_result *DescribeDcdnDomainTopUrlVisitResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SortBy)) {
		query["SortBy"] = request.SortBy
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDcdnDomainTopUrlVisit"),
		Version:     tea.String("2018-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDcdnDomainTopUrlVisitResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDcdnDomainTopUrlVisit(request *DescribeDcdnDomainTopUrlVisitRequest) (_result *DescribeDcdnDomainTopUrlVisitResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDcdnDomainTopUrlVisitResponse{}
	_body, _err := client.DescribeDcdnDomainTopUrlVisitWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDcdnDomainTrafficDataWithOptions(request *DescribeDcdnDomainTrafficDataRequest, runtime *util.RuntimeOptions) (_result *DescribeDcdnDomainTrafficDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Interval)) {
		query["Interval"] = request.Interval
	}

	if !tea.BoolValue(util.IsUnset(request.IspNameEn)) {
		query["IspNameEn"] = request.IspNameEn
	}

	if !tea.BoolValue(util.IsUnset(request.LocationNameEn)) {
		query["LocationNameEn"] = request.LocationNameEn
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDcdnDomainTrafficData"),
		Version:     tea.String("2018-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDcdnDomainTrafficDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDcdnDomainTrafficData(request *DescribeDcdnDomainTrafficDataRequest) (_result *DescribeDcdnDomainTrafficDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDcdnDomainTrafficDataResponse{}
	_body, _err := client.DescribeDcdnDomainTrafficDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDcdnDomainUsageDataWithOptions(request *DescribeDcdnDomainUsageDataRequest, runtime *util.RuntimeOptions) (_result *DescribeDcdnDomainUsageDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Area)) {
		query["Area"] = request.Area
	}

	if !tea.BoolValue(util.IsUnset(request.DataProtocol)) {
		query["DataProtocol"] = request.DataProtocol
	}

	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Field)) {
		query["Field"] = request.Field
	}

	if !tea.BoolValue(util.IsUnset(request.Interval)) {
		query["Interval"] = request.Interval
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDcdnDomainUsageData"),
		Version:     tea.String("2018-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDcdnDomainUsageDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDcdnDomainUsageData(request *DescribeDcdnDomainUsageDataRequest) (_result *DescribeDcdnDomainUsageDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDcdnDomainUsageDataResponse{}
	_body, _err := client.DescribeDcdnDomainUsageDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDcdnDomainUvDataWithOptions(request *DescribeDcdnDomainUvDataRequest, runtime *util.RuntimeOptions) (_result *DescribeDcdnDomainUvDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDcdnDomainUvData"),
		Version:     tea.String("2018-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDcdnDomainUvDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDcdnDomainUvData(request *DescribeDcdnDomainUvDataRequest) (_result *DescribeDcdnDomainUvDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDcdnDomainUvDataResponse{}
	_body, _err := client.DescribeDcdnDomainUvDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDcdnDomainWebsocketBpsDataWithOptions(request *DescribeDcdnDomainWebsocketBpsDataRequest, runtime *util.RuntimeOptions) (_result *DescribeDcdnDomainWebsocketBpsDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Interval)) {
		query["Interval"] = request.Interval
	}

	if !tea.BoolValue(util.IsUnset(request.IspNameEn)) {
		query["IspNameEn"] = request.IspNameEn
	}

	if !tea.BoolValue(util.IsUnset(request.LocationNameEn)) {
		query["LocationNameEn"] = request.LocationNameEn
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDcdnDomainWebsocketBpsData"),
		Version:     tea.String("2018-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDcdnDomainWebsocketBpsDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDcdnDomainWebsocketBpsData(request *DescribeDcdnDomainWebsocketBpsDataRequest) (_result *DescribeDcdnDomainWebsocketBpsDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDcdnDomainWebsocketBpsDataResponse{}
	_body, _err := client.DescribeDcdnDomainWebsocketBpsDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDcdnDomainWebsocketHttpCodeDataWithOptions(request *DescribeDcdnDomainWebsocketHttpCodeDataRequest, runtime *util.RuntimeOptions) (_result *DescribeDcdnDomainWebsocketHttpCodeDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Interval)) {
		query["Interval"] = request.Interval
	}

	if !tea.BoolValue(util.IsUnset(request.IspNameEn)) {
		query["IspNameEn"] = request.IspNameEn
	}

	if !tea.BoolValue(util.IsUnset(request.LocationNameEn)) {
		query["LocationNameEn"] = request.LocationNameEn
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDcdnDomainWebsocketHttpCodeData"),
		Version:     tea.String("2018-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDcdnDomainWebsocketHttpCodeDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDcdnDomainWebsocketHttpCodeData(request *DescribeDcdnDomainWebsocketHttpCodeDataRequest) (_result *DescribeDcdnDomainWebsocketHttpCodeDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDcdnDomainWebsocketHttpCodeDataResponse{}
	_body, _err := client.DescribeDcdnDomainWebsocketHttpCodeDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDcdnDomainWebsocketTrafficDataWithOptions(request *DescribeDcdnDomainWebsocketTrafficDataRequest, runtime *util.RuntimeOptions) (_result *DescribeDcdnDomainWebsocketTrafficDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Interval)) {
		query["Interval"] = request.Interval
	}

	if !tea.BoolValue(util.IsUnset(request.IspNameEn)) {
		query["IspNameEn"] = request.IspNameEn
	}

	if !tea.BoolValue(util.IsUnset(request.LocationNameEn)) {
		query["LocationNameEn"] = request.LocationNameEn
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDcdnDomainWebsocketTrafficData"),
		Version:     tea.String("2018-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDcdnDomainWebsocketTrafficDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDcdnDomainWebsocketTrafficData(request *DescribeDcdnDomainWebsocketTrafficDataRequest) (_result *DescribeDcdnDomainWebsocketTrafficDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDcdnDomainWebsocketTrafficDataResponse{}
	_body, _err := client.DescribeDcdnDomainWebsocketTrafficDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDcdnEsExceptionDataWithOptions(request *DescribeDcdnEsExceptionDataRequest, runtime *util.RuntimeOptions) (_result *DescribeDcdnEsExceptionDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RuleId)) {
		query["RuleId"] = request.RuleId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDcdnEsExceptionData"),
		Version:     tea.String("2018-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDcdnEsExceptionDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDcdnEsExceptionData(request *DescribeDcdnEsExceptionDataRequest) (_result *DescribeDcdnEsExceptionDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDcdnEsExceptionDataResponse{}
	_body, _err := client.DescribeDcdnEsExceptionDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDcdnEsExecuteDataWithOptions(request *DescribeDcdnEsExecuteDataRequest, runtime *util.RuntimeOptions) (_result *DescribeDcdnEsExecuteDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RuleId)) {
		query["RuleId"] = request.RuleId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDcdnEsExecuteData"),
		Version:     tea.String("2018-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDcdnEsExecuteDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDcdnEsExecuteData(request *DescribeDcdnEsExecuteDataRequest) (_result *DescribeDcdnEsExecuteDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDcdnEsExecuteDataResponse{}
	_body, _err := client.DescribeDcdnEsExecuteDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDcdnHttpsDomainListWithOptions(request *DescribeDcdnHttpsDomainListRequest, runtime *util.RuntimeOptions) (_result *DescribeDcdnHttpsDomainListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Keyword)) {
		query["Keyword"] = request.Keyword
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDcdnHttpsDomainList"),
		Version:     tea.String("2018-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDcdnHttpsDomainListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDcdnHttpsDomainList(request *DescribeDcdnHttpsDomainListRequest) (_result *DescribeDcdnHttpsDomainListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDcdnHttpsDomainListResponse{}
	_body, _err := client.DescribeDcdnHttpsDomainListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDcdnIpInfoWithOptions(request *DescribeDcdnIpInfoRequest, runtime *util.RuntimeOptions) (_result *DescribeDcdnIpInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IP)) {
		query["IP"] = request.IP
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDcdnIpInfo"),
		Version:     tea.String("2018-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDcdnIpInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDcdnIpInfo(request *DescribeDcdnIpInfoRequest) (_result *DescribeDcdnIpInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDcdnIpInfoResponse{}
	_body, _err := client.DescribeDcdnIpInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDcdnIpaDomainConfigsWithOptions(request *DescribeDcdnIpaDomainConfigsRequest, runtime *util.RuntimeOptions) (_result *DescribeDcdnIpaDomainConfigsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.FunctionNames)) {
		query["FunctionNames"] = request.FunctionNames
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDcdnIpaDomainConfigs"),
		Version:     tea.String("2018-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDcdnIpaDomainConfigsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDcdnIpaDomainConfigs(request *DescribeDcdnIpaDomainConfigsRequest) (_result *DescribeDcdnIpaDomainConfigsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDcdnIpaDomainConfigsResponse{}
	_body, _err := client.DescribeDcdnIpaDomainConfigsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDcdnIpaDomainDetailWithOptions(request *DescribeDcdnIpaDomainDetailRequest, runtime *util.RuntimeOptions) (_result *DescribeDcdnIpaDomainDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDcdnIpaDomainDetail"),
		Version:     tea.String("2018-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDcdnIpaDomainDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDcdnIpaDomainDetail(request *DescribeDcdnIpaDomainDetailRequest) (_result *DescribeDcdnIpaDomainDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDcdnIpaDomainDetailResponse{}
	_body, _err := client.DescribeDcdnIpaDomainDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDcdnIpaServiceWithOptions(request *DescribeDcdnIpaServiceRequest, runtime *util.RuntimeOptions) (_result *DescribeDcdnIpaServiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDcdnIpaService"),
		Version:     tea.String("2018-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDcdnIpaServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDcdnIpaService(request *DescribeDcdnIpaServiceRequest) (_result *DescribeDcdnIpaServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDcdnIpaServiceResponse{}
	_body, _err := client.DescribeDcdnIpaServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDcdnIpaUserDomainsWithOptions(request *DescribeDcdnIpaUserDomainsRequest, runtime *util.RuntimeOptions) (_result *DescribeDcdnIpaUserDomainsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CheckDomainShow)) {
		query["CheckDomainShow"] = request.CheckDomainShow
	}

	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.DomainSearchType)) {
		query["DomainSearchType"] = request.DomainSearchType
	}

	if !tea.BoolValue(util.IsUnset(request.DomainStatus)) {
		query["DomainStatus"] = request.DomainStatus
	}

	if !tea.BoolValue(util.IsUnset(request.FuncFilter)) {
		query["FuncFilter"] = request.FuncFilter
	}

	if !tea.BoolValue(util.IsUnset(request.FuncId)) {
		query["FuncId"] = request.FuncId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDcdnIpaUserDomains"),
		Version:     tea.String("2018-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDcdnIpaUserDomainsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDcdnIpaUserDomains(request *DescribeDcdnIpaUserDomainsRequest) (_result *DescribeDcdnIpaUserDomainsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDcdnIpaUserDomainsResponse{}
	_body, _err := client.DescribeDcdnIpaUserDomainsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDcdnRealTimeDeliveryFieldWithOptions(request *DescribeDcdnRealTimeDeliveryFieldRequest, runtime *util.RuntimeOptions) (_result *DescribeDcdnRealTimeDeliveryFieldResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BusinessType)) {
		query["BusinessType"] = request.BusinessType
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDcdnRealTimeDeliveryField"),
		Version:     tea.String("2018-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDcdnRealTimeDeliveryFieldResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDcdnRealTimeDeliveryField(request *DescribeDcdnRealTimeDeliveryFieldRequest) (_result *DescribeDcdnRealTimeDeliveryFieldResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDcdnRealTimeDeliveryFieldResponse{}
	_body, _err := client.DescribeDcdnRealTimeDeliveryFieldWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDcdnRefreshQuotaWithOptions(request *DescribeDcdnRefreshQuotaRequest, runtime *util.RuntimeOptions) (_result *DescribeDcdnRefreshQuotaResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDcdnRefreshQuota"),
		Version:     tea.String("2018-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDcdnRefreshQuotaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDcdnRefreshQuota(request *DescribeDcdnRefreshQuotaRequest) (_result *DescribeDcdnRefreshQuotaResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDcdnRefreshQuotaResponse{}
	_body, _err := client.DescribeDcdnRefreshQuotaWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDcdnRefreshTaskByIdWithOptions(request *DescribeDcdnRefreshTaskByIdRequest, runtime *util.RuntimeOptions) (_result *DescribeDcdnRefreshTaskByIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		query["TaskId"] = request.TaskId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDcdnRefreshTaskById"),
		Version:     tea.String("2018-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDcdnRefreshTaskByIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDcdnRefreshTaskById(request *DescribeDcdnRefreshTaskByIdRequest) (_result *DescribeDcdnRefreshTaskByIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDcdnRefreshTaskByIdResponse{}
	_body, _err := client.DescribeDcdnRefreshTaskByIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDcdnRefreshTasksWithOptions(request *DescribeDcdnRefreshTasksRequest, runtime *util.RuntimeOptions) (_result *DescribeDcdnRefreshTasksResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.ObjectPath)) {
		query["ObjectPath"] = request.ObjectPath
	}

	if !tea.BoolValue(util.IsUnset(request.ObjectType)) {
		query["ObjectType"] = request.ObjectType
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		query["TaskId"] = request.TaskId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDcdnRefreshTasks"),
		Version:     tea.String("2018-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDcdnRefreshTasksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDcdnRefreshTasks(request *DescribeDcdnRefreshTasksRequest) (_result *DescribeDcdnRefreshTasksResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDcdnRefreshTasksResponse{}
	_body, _err := client.DescribeDcdnRefreshTasksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDcdnRegionAndIspWithOptions(request *DescribeDcdnRegionAndIspRequest, runtime *util.RuntimeOptions) (_result *DescribeDcdnRegionAndIspResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDcdnRegionAndIsp"),
		Version:     tea.String("2018-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDcdnRegionAndIspResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDcdnRegionAndIsp(request *DescribeDcdnRegionAndIspRequest) (_result *DescribeDcdnRegionAndIspResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDcdnRegionAndIspResponse{}
	_body, _err := client.DescribeDcdnRegionAndIspWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDcdnReportWithOptions(request *DescribeDcdnReportRequest, runtime *util.RuntimeOptions) (_result *DescribeDcdnReportResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Area)) {
		query["Area"] = request.Area
	}

	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.HttpCode)) {
		query["HttpCode"] = request.HttpCode
	}

	if !tea.BoolValue(util.IsUnset(request.IsOverseas)) {
		query["IsOverseas"] = request.IsOverseas
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ReportId)) {
		query["ReportId"] = request.ReportId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDcdnReport"),
		Version:     tea.String("2018-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDcdnReportResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDcdnReport(request *DescribeDcdnReportRequest) (_result *DescribeDcdnReportResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDcdnReportResponse{}
	_body, _err := client.DescribeDcdnReportWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDcdnReportListWithOptions(request *DescribeDcdnReportListRequest, runtime *util.RuntimeOptions) (_result *DescribeDcdnReportListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ReportId)) {
		query["ReportId"] = request.ReportId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDcdnReportList"),
		Version:     tea.String("2018-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDcdnReportListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDcdnReportList(request *DescribeDcdnReportListRequest) (_result *DescribeDcdnReportListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDcdnReportListResponse{}
	_body, _err := client.DescribeDcdnReportListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDcdnSLSRealtimeLogDeliveryWithOptions(request *DescribeDcdnSLSRealtimeLogDeliveryRequest, runtime *util.RuntimeOptions) (_result *DescribeDcdnSLSRealtimeLogDeliveryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BusinessType)) {
		query["BusinessType"] = request.BusinessType
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectName)) {
		query["ProjectName"] = request.ProjectName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDcdnSLSRealtimeLogDelivery"),
		Version:     tea.String("2018-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDcdnSLSRealtimeLogDeliveryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDcdnSLSRealtimeLogDelivery(request *DescribeDcdnSLSRealtimeLogDeliveryRequest) (_result *DescribeDcdnSLSRealtimeLogDeliveryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDcdnSLSRealtimeLogDeliveryResponse{}
	_body, _err := client.DescribeDcdnSLSRealtimeLogDeliveryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDcdnSMCertificateDetailWithOptions(request *DescribeDcdnSMCertificateDetailRequest, runtime *util.RuntimeOptions) (_result *DescribeDcdnSMCertificateDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CertIdentifier)) {
		query["CertIdentifier"] = request.CertIdentifier
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDcdnSMCertificateDetail"),
		Version:     tea.String("2018-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDcdnSMCertificateDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDcdnSMCertificateDetail(request *DescribeDcdnSMCertificateDetailRequest) (_result *DescribeDcdnSMCertificateDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDcdnSMCertificateDetailResponse{}
	_body, _err := client.DescribeDcdnSMCertificateDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDcdnSMCertificateListWithOptions(request *DescribeDcdnSMCertificateListRequest, runtime *util.RuntimeOptions) (_result *DescribeDcdnSMCertificateListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDcdnSMCertificateList"),
		Version:     tea.String("2018-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDcdnSMCertificateListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDcdnSMCertificateList(request *DescribeDcdnSMCertificateListRequest) (_result *DescribeDcdnSMCertificateListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDcdnSMCertificateListResponse{}
	_body, _err := client.DescribeDcdnSMCertificateListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDcdnSecFuncInfoWithOptions(request *DescribeDcdnSecFuncInfoRequest, runtime *util.RuntimeOptions) (_result *DescribeDcdnSecFuncInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SecFuncType)) {
		query["SecFuncType"] = request.SecFuncType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDcdnSecFuncInfo"),
		Version:     tea.String("2018-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDcdnSecFuncInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDcdnSecFuncInfo(request *DescribeDcdnSecFuncInfoRequest) (_result *DescribeDcdnSecFuncInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDcdnSecFuncInfoResponse{}
	_body, _err := client.DescribeDcdnSecFuncInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDcdnSecSpecInfoWithOptions(request *DescribeDcdnSecSpecInfoRequest, runtime *util.RuntimeOptions) (_result *DescribeDcdnSecSpecInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDcdnSecSpecInfo"),
		Version:     tea.String("2018-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDcdnSecSpecInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDcdnSecSpecInfo(request *DescribeDcdnSecSpecInfoRequest) (_result *DescribeDcdnSecSpecInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDcdnSecSpecInfoResponse{}
	_body, _err := client.DescribeDcdnSecSpecInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDcdnServiceWithOptions(request *DescribeDcdnServiceRequest, runtime *util.RuntimeOptions) (_result *DescribeDcdnServiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDcdnService"),
		Version:     tea.String("2018-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDcdnServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDcdnService(request *DescribeDcdnServiceRequest) (_result *DescribeDcdnServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDcdnServiceResponse{}
	_body, _err := client.DescribeDcdnServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDcdnStagingIpWithOptions(request *DescribeDcdnStagingIpRequest, runtime *util.RuntimeOptions) (_result *DescribeDcdnStagingIpResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDcdnStagingIp"),
		Version:     tea.String("2018-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDcdnStagingIpResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDcdnStagingIp(request *DescribeDcdnStagingIpRequest) (_result *DescribeDcdnStagingIpResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDcdnStagingIpResponse{}
	_body, _err := client.DescribeDcdnStagingIpWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDcdnSubListWithOptions(request *DescribeDcdnSubListRequest, runtime *util.RuntimeOptions) (_result *DescribeDcdnSubListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDcdnSubList"),
		Version:     tea.String("2018-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDcdnSubListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDcdnSubList(request *DescribeDcdnSubListRequest) (_result *DescribeDcdnSubListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDcdnSubListResponse{}
	_body, _err := client.DescribeDcdnSubListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDcdnTagResourcesWithOptions(request *DescribeDcdnTagResourcesRequest, runtime *util.RuntimeOptions) (_result *DescribeDcdnTagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDcdnTagResources"),
		Version:     tea.String("2018-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDcdnTagResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDcdnTagResources(request *DescribeDcdnTagResourcesRequest) (_result *DescribeDcdnTagResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDcdnTagResourcesResponse{}
	_body, _err := client.DescribeDcdnTagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDcdnTopDomainsByFlowWithOptions(request *DescribeDcdnTopDomainsByFlowRequest, runtime *util.RuntimeOptions) (_result *DescribeDcdnTopDomainsByFlowResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Limit)) {
		query["Limit"] = request.Limit
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDcdnTopDomainsByFlow"),
		Version:     tea.String("2018-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDcdnTopDomainsByFlowResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDcdnTopDomainsByFlow(request *DescribeDcdnTopDomainsByFlowRequest) (_result *DescribeDcdnTopDomainsByFlowResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDcdnTopDomainsByFlowResponse{}
	_body, _err := client.DescribeDcdnTopDomainsByFlowWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDcdnUserBillHistoryWithOptions(request *DescribeDcdnUserBillHistoryRequest, runtime *util.RuntimeOptions) (_result *DescribeDcdnUserBillHistoryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDcdnUserBillHistory"),
		Version:     tea.String("2018-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDcdnUserBillHistoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDcdnUserBillHistory(request *DescribeDcdnUserBillHistoryRequest) (_result *DescribeDcdnUserBillHistoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDcdnUserBillHistoryResponse{}
	_body, _err := client.DescribeDcdnUserBillHistoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDcdnUserBillTypeWithOptions(request *DescribeDcdnUserBillTypeRequest, runtime *util.RuntimeOptions) (_result *DescribeDcdnUserBillTypeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDcdnUserBillType"),
		Version:     tea.String("2018-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDcdnUserBillTypeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDcdnUserBillType(request *DescribeDcdnUserBillTypeRequest) (_result *DescribeDcdnUserBillTypeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDcdnUserBillTypeResponse{}
	_body, _err := client.DescribeDcdnUserBillTypeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDcdnUserDomainsWithOptions(request *DescribeDcdnUserDomainsRequest, runtime *util.RuntimeOptions) (_result *DescribeDcdnUserDomainsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ChangeEndTime)) {
		query["ChangeEndTime"] = request.ChangeEndTime
	}

	if !tea.BoolValue(util.IsUnset(request.ChangeStartTime)) {
		query["ChangeStartTime"] = request.ChangeStartTime
	}

	if !tea.BoolValue(util.IsUnset(request.CheckDomainShow)) {
		query["CheckDomainShow"] = request.CheckDomainShow
	}

	if !tea.BoolValue(util.IsUnset(request.Coverage)) {
		query["Coverage"] = request.Coverage
	}

	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.DomainSearchType)) {
		query["DomainSearchType"] = request.DomainSearchType
	}

	if !tea.BoolValue(util.IsUnset(request.DomainStatus)) {
		query["DomainStatus"] = request.DomainStatus
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDcdnUserDomains"),
		Version:     tea.String("2018-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDcdnUserDomainsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDcdnUserDomains(request *DescribeDcdnUserDomainsRequest) (_result *DescribeDcdnUserDomainsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDcdnUserDomainsResponse{}
	_body, _err := client.DescribeDcdnUserDomainsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDcdnUserDomainsByFuncWithOptions(request *DescribeDcdnUserDomainsByFuncRequest, runtime *util.RuntimeOptions) (_result *DescribeDcdnUserDomainsByFuncResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.FuncFilter)) {
		query["FuncFilter"] = request.FuncFilter
	}

	if !tea.BoolValue(util.IsUnset(request.FuncId)) {
		query["FuncId"] = request.FuncId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDcdnUserDomainsByFunc"),
		Version:     tea.String("2018-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDcdnUserDomainsByFuncResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDcdnUserDomainsByFunc(request *DescribeDcdnUserDomainsByFuncRequest) (_result *DescribeDcdnUserDomainsByFuncResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDcdnUserDomainsByFuncResponse{}
	_body, _err := client.DescribeDcdnUserDomainsByFuncWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDcdnUserQuotaWithOptions(request *DescribeDcdnUserQuotaRequest, runtime *util.RuntimeOptions) (_result *DescribeDcdnUserQuotaResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDcdnUserQuota"),
		Version:     tea.String("2018-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDcdnUserQuotaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDcdnUserQuota(request *DescribeDcdnUserQuotaRequest) (_result *DescribeDcdnUserQuotaResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDcdnUserQuotaResponse{}
	_body, _err := client.DescribeDcdnUserQuotaWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDcdnUserRealTimeDeliveryFieldWithOptions(request *DescribeDcdnUserRealTimeDeliveryFieldRequest, runtime *util.RuntimeOptions) (_result *DescribeDcdnUserRealTimeDeliveryFieldResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDcdnUserRealTimeDeliveryField"),
		Version:     tea.String("2018-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDcdnUserRealTimeDeliveryFieldResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDcdnUserRealTimeDeliveryField(request *DescribeDcdnUserRealTimeDeliveryFieldRequest) (_result *DescribeDcdnUserRealTimeDeliveryFieldResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDcdnUserRealTimeDeliveryFieldResponse{}
	_body, _err := client.DescribeDcdnUserRealTimeDeliveryFieldWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDcdnUserResourcePackageWithOptions(request *DescribeDcdnUserResourcePackageRequest, runtime *util.RuntimeOptions) (_result *DescribeDcdnUserResourcePackageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDcdnUserResourcePackage"),
		Version:     tea.String("2018-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDcdnUserResourcePackageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDcdnUserResourcePackage(request *DescribeDcdnUserResourcePackageRequest) (_result *DescribeDcdnUserResourcePackageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDcdnUserResourcePackageResponse{}
	_body, _err := client.DescribeDcdnUserResourcePackageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDcdnUserSecDropWithOptions(request *DescribeDcdnUserSecDropRequest, runtime *util.RuntimeOptions) (_result *DescribeDcdnUserSecDropResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Data)) {
		query["Data"] = request.Data
	}

	if !tea.BoolValue(util.IsUnset(request.Metric)) {
		query["Metric"] = request.Metric
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SecFunc)) {
		query["SecFunc"] = request.SecFunc
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDcdnUserSecDrop"),
		Version:     tea.String("2018-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDcdnUserSecDropResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDcdnUserSecDrop(request *DescribeDcdnUserSecDropRequest) (_result *DescribeDcdnUserSecDropResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDcdnUserSecDropResponse{}
	_body, _err := client.DescribeDcdnUserSecDropWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDcdnUserSecDropByMinuteWithOptions(request *DescribeDcdnUserSecDropByMinuteRequest, runtime *util.RuntimeOptions) (_result *DescribeDcdnUserSecDropByMinuteResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.Object)) {
		query["Object"] = request.Object
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RuleName)) {
		query["RuleName"] = request.RuleName
	}

	if !tea.BoolValue(util.IsUnset(request.SecFunc)) {
		query["SecFunc"] = request.SecFunc
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDcdnUserSecDropByMinute"),
		Version:     tea.String("2018-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDcdnUserSecDropByMinuteResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDcdnUserSecDropByMinute(request *DescribeDcdnUserSecDropByMinuteRequest) (_result *DescribeDcdnUserSecDropByMinuteResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDcdnUserSecDropByMinuteResponse{}
	_body, _err := client.DescribeDcdnUserSecDropByMinuteWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDcdnUserTagsWithOptions(request *DescribeDcdnUserTagsRequest, runtime *util.RuntimeOptions) (_result *DescribeDcdnUserTagsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDcdnUserTags"),
		Version:     tea.String("2018-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDcdnUserTagsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDcdnUserTags(request *DescribeDcdnUserTagsRequest) (_result *DescribeDcdnUserTagsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDcdnUserTagsResponse{}
	_body, _err := client.DescribeDcdnUserTagsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDcdnVerifyContentWithOptions(request *DescribeDcdnVerifyContentRequest, runtime *util.RuntimeOptions) (_result *DescribeDcdnVerifyContentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDcdnVerifyContent"),
		Version:     tea.String("2018-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDcdnVerifyContentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDcdnVerifyContent(request *DescribeDcdnVerifyContentRequest) (_result *DescribeDcdnVerifyContentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDcdnVerifyContentResponse{}
	_body, _err := client.DescribeDcdnVerifyContentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDcdnWafDomainWithOptions(request *DescribeDcdnWafDomainRequest, runtime *util.RuntimeOptions) (_result *DescribeDcdnWafDomainResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDcdnWafDomain"),
		Version:     tea.String("2018-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDcdnWafDomainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDcdnWafDomain(request *DescribeDcdnWafDomainRequest) (_result *DescribeDcdnWafDomainResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDcdnWafDomainResponse{}
	_body, _err := client.DescribeDcdnWafDomainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDcdnsecServiceWithOptions(request *DescribeDcdnsecServiceRequest, runtime *util.RuntimeOptions) (_result *DescribeDcdnsecServiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDcdnsecService"),
		Version:     tea.String("2018-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDcdnsecServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDcdnsecService(request *DescribeDcdnsecServiceRequest) (_result *DescribeDcdnsecServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDcdnsecServiceResponse{}
	_body, _err := client.DescribeDcdnsecServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeRoutineWithOptions(request *DescribeRoutineRequest, runtime *util.RuntimeOptions) (_result *DescribeRoutineResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeRoutine"),
		Version:     tea.String("2018-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeRoutineResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeRoutine(request *DescribeRoutineRequest) (_result *DescribeRoutineResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRoutineResponse{}
	_body, _err := client.DescribeRoutineWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeRoutineCanaryEnvsWithOptions(request *DescribeRoutineCanaryEnvsRequest, runtime *util.RuntimeOptions) (_result *DescribeRoutineCanaryEnvsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeRoutineCanaryEnvs"),
		Version:     tea.String("2018-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeRoutineCanaryEnvsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeRoutineCanaryEnvs(request *DescribeRoutineCanaryEnvsRequest) (_result *DescribeRoutineCanaryEnvsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRoutineCanaryEnvsResponse{}
	_body, _err := client.DescribeRoutineCanaryEnvsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeRoutineCodeRevisionWithOptions(request *DescribeRoutineCodeRevisionRequest, runtime *util.RuntimeOptions) (_result *DescribeRoutineCodeRevisionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.SelectCodeRevision)) {
		body["SelectCodeRevision"] = request.SelectCodeRevision
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeRoutineCodeRevision"),
		Version:     tea.String("2018-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeRoutineCodeRevisionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeRoutineCodeRevision(request *DescribeRoutineCodeRevisionRequest) (_result *DescribeRoutineCodeRevisionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRoutineCodeRevisionResponse{}
	_body, _err := client.DescribeRoutineCodeRevisionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeRoutineSpecWithOptions(request *DescribeRoutineSpecRequest, runtime *util.RuntimeOptions) (_result *DescribeRoutineSpecResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeRoutineSpec"),
		Version:     tea.String("2018-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeRoutineSpecResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeRoutineSpec(request *DescribeRoutineSpecRequest) (_result *DescribeRoutineSpecResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRoutineSpecResponse{}
	_body, _err := client.DescribeRoutineSpecWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeRoutineUserInfoWithOptions(request *DescribeRoutineUserInfoRequest, runtime *util.RuntimeOptions) (_result *DescribeRoutineUserInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeRoutineUserInfo"),
		Version:     tea.String("2018-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeRoutineUserInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeRoutineUserInfo(request *DescribeRoutineUserInfoRequest) (_result *DescribeRoutineUserInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRoutineUserInfoResponse{}
	_body, _err := client.DescribeRoutineUserInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeUserDcdnIpaStatusWithOptions(request *DescribeUserDcdnIpaStatusRequest, runtime *util.RuntimeOptions) (_result *DescribeUserDcdnIpaStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeUserDcdnIpaStatus"),
		Version:     tea.String("2018-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeUserDcdnIpaStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeUserDcdnIpaStatus(request *DescribeUserDcdnIpaStatusRequest) (_result *DescribeUserDcdnIpaStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeUserDcdnIpaStatusResponse{}
	_body, _err := client.DescribeUserDcdnIpaStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeUserDcdnStatusWithOptions(request *DescribeUserDcdnStatusRequest, runtime *util.RuntimeOptions) (_result *DescribeUserDcdnStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeUserDcdnStatus"),
		Version:     tea.String("2018-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeUserDcdnStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeUserDcdnStatus(request *DescribeUserDcdnStatusRequest) (_result *DescribeUserDcdnStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeUserDcdnStatusResponse{}
	_body, _err := client.DescribeUserDcdnStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeUserErStatusWithOptions(request *DescribeUserErStatusRequest, runtime *util.RuntimeOptions) (_result *DescribeUserErStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeUserErStatus"),
		Version:     tea.String("2018-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeUserErStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeUserErStatus(request *DescribeUserErStatusRequest) (_result *DescribeUserErStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeUserErStatusResponse{}
	_body, _err := client.DescribeUserErStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeUserLogserviceStatusWithOptions(request *DescribeUserLogserviceStatusRequest, runtime *util.RuntimeOptions) (_result *DescribeUserLogserviceStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeUserLogserviceStatus"),
		Version:     tea.String("2018-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeUserLogserviceStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeUserLogserviceStatus(request *DescribeUserLogserviceStatusRequest) (_result *DescribeUserLogserviceStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeUserLogserviceStatusResponse{}
	_body, _err := client.DescribeUserLogserviceStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) EditRoutineConfWithOptions(tmpReq *EditRoutineConfRequest, runtime *util.RuntimeOptions) (_result *EditRoutineConfResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &EditRoutineConfShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.EnvConf)) {
		request.EnvConfShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.EnvConf, tea.String("EnvConf"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.EnvConfShrink)) {
		body["EnvConf"] = request.EnvConfShrink
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("EditRoutineConf"),
		Version:     tea.String("2018-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &EditRoutineConfResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) EditRoutineConf(request *EditRoutineConfRequest) (_result *EditRoutineConfResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EditRoutineConfResponse{}
	_body, _err := client.EditRoutineConfWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListDcdnEsTemplateInfoWithOptions(request *ListDcdnEsTemplateInfoRequest, runtime *util.RuntimeOptions) (_result *ListDcdnEsTemplateInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListDcdnEsTemplateInfo"),
		Version:     tea.String("2018-01-15"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListDcdnEsTemplateInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListDcdnEsTemplateInfo(request *ListDcdnEsTemplateInfoRequest) (_result *ListDcdnEsTemplateInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListDcdnEsTemplateInfoResponse{}
	_body, _err := client.ListDcdnEsTemplateInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListDcdnRealTimeDeliveryProjectWithOptions(request *ListDcdnRealTimeDeliveryProjectRequest, runtime *util.RuntimeOptions) (_result *ListDcdnRealTimeDeliveryProjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BusinessType)) {
		query["BusinessType"] = request.BusinessType
	}

	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListDcdnRealTimeDeliveryProject"),
		Version:     tea.String("2018-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListDcdnRealTimeDeliveryProjectResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListDcdnRealTimeDeliveryProject(request *ListDcdnRealTimeDeliveryProjectRequest) (_result *ListDcdnRealTimeDeliveryProjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListDcdnRealTimeDeliveryProjectResponse{}
	_body, _err := client.ListDcdnRealTimeDeliveryProjectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyDCdnDomainSchdmByPropertyWithOptions(request *ModifyDCdnDomainSchdmByPropertyRequest, runtime *util.RuntimeOptions) (_result *ModifyDCdnDomainSchdmByPropertyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.Property)) {
		query["Property"] = request.Property
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyDCdnDomainSchdmByProperty"),
		Version:     tea.String("2018-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyDCdnDomainSchdmByPropertyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyDCdnDomainSchdmByProperty(request *ModifyDCdnDomainSchdmByPropertyRequest) (_result *ModifyDCdnDomainSchdmByPropertyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyDCdnDomainSchdmByPropertyResponse{}
	_body, _err := client.ModifyDCdnDomainSchdmByPropertyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) OpenDcdnServiceWithOptions(request *OpenDcdnServiceRequest, runtime *util.RuntimeOptions) (_result *OpenDcdnServiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BillType)) {
		query["BillType"] = request.BillType
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !tea.BoolValue(util.IsUnset(request.WebsocketBillType)) {
		query["WebsocketBillType"] = request.WebsocketBillType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("OpenDcdnService"),
		Version:     tea.String("2018-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &OpenDcdnServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) OpenDcdnService(request *OpenDcdnServiceRequest) (_result *OpenDcdnServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &OpenDcdnServiceResponse{}
	_body, _err := client.OpenDcdnServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PreloadDcdnObjectCachesWithOptions(request *PreloadDcdnObjectCachesRequest, runtime *util.RuntimeOptions) (_result *PreloadDcdnObjectCachesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Area)) {
		query["Area"] = request.Area
	}

	if !tea.BoolValue(util.IsUnset(request.ObjectPath)) {
		query["ObjectPath"] = request.ObjectPath
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("PreloadDcdnObjectCaches"),
		Version:     tea.String("2018-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PreloadDcdnObjectCachesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PreloadDcdnObjectCaches(request *PreloadDcdnObjectCachesRequest) (_result *PreloadDcdnObjectCachesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PreloadDcdnObjectCachesResponse{}
	_body, _err := client.PreloadDcdnObjectCachesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PublishDcdnStagingConfigToProductionWithOptions(request *PublishDcdnStagingConfigToProductionRequest, runtime *util.RuntimeOptions) (_result *PublishDcdnStagingConfigToProductionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.FunctionName)) {
		query["FunctionName"] = request.FunctionName
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("PublishDcdnStagingConfigToProduction"),
		Version:     tea.String("2018-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PublishDcdnStagingConfigToProductionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PublishDcdnStagingConfigToProduction(request *PublishDcdnStagingConfigToProductionRequest) (_result *PublishDcdnStagingConfigToProductionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PublishDcdnStagingConfigToProductionResponse{}
	_body, _err := client.PublishDcdnStagingConfigToProductionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PublishRoutineCodeRevisionWithOptions(tmpReq *PublishRoutineCodeRevisionRequest, runtime *util.RuntimeOptions) (_result *PublishRoutineCodeRevisionResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &PublishRoutineCodeRevisionShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Envs)) {
		request.EnvsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Envs, tea.String("Envs"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EnvsShrink)) {
		body["Envs"] = request.EnvsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.SelectCodeRevision)) {
		body["SelectCodeRevision"] = request.SelectCodeRevision
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("PublishRoutineCodeRevision"),
		Version:     tea.String("2018-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PublishRoutineCodeRevisionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PublishRoutineCodeRevision(request *PublishRoutineCodeRevisionRequest) (_result *PublishRoutineCodeRevisionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PublishRoutineCodeRevisionResponse{}
	_body, _err := client.PublishRoutineCodeRevisionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RefreshDcdnObjectCachesWithOptions(request *RefreshDcdnObjectCachesRequest, runtime *util.RuntimeOptions) (_result *RefreshDcdnObjectCachesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ObjectPath)) {
		query["ObjectPath"] = request.ObjectPath
	}

	if !tea.BoolValue(util.IsUnset(request.ObjectType)) {
		query["ObjectType"] = request.ObjectType
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RefreshDcdnObjectCaches"),
		Version:     tea.String("2018-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RefreshDcdnObjectCachesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RefreshDcdnObjectCaches(request *RefreshDcdnObjectCachesRequest) (_result *RefreshDcdnObjectCachesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RefreshDcdnObjectCachesResponse{}
	_body, _err := client.RefreshDcdnObjectCachesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RollbackDcdnStagingConfigWithOptions(request *RollbackDcdnStagingConfigRequest, runtime *util.RuntimeOptions) (_result *RollbackDcdnStagingConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RollbackDcdnStagingConfig"),
		Version:     tea.String("2018-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RollbackDcdnStagingConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RollbackDcdnStagingConfig(request *RollbackDcdnStagingConfigRequest) (_result *RollbackDcdnStagingConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RollbackDcdnStagingConfigResponse{}
	_body, _err := client.RollbackDcdnStagingConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetDcdnConfigOfVersionWithOptions(request *SetDcdnConfigOfVersionRequest, runtime *util.RuntimeOptions) (_result *SetDcdnConfigOfVersionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConfigId)) {
		query["ConfigId"] = request.ConfigId
	}

	if !tea.BoolValue(util.IsUnset(request.FunctionArgs)) {
		query["FunctionArgs"] = request.FunctionArgs
	}

	if !tea.BoolValue(util.IsUnset(request.FunctionId)) {
		query["FunctionId"] = request.FunctionId
	}

	if !tea.BoolValue(util.IsUnset(request.FunctionName)) {
		query["FunctionName"] = request.FunctionName
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !tea.BoolValue(util.IsUnset(request.VersionId)) {
		query["VersionId"] = request.VersionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SetDcdnConfigOfVersion"),
		Version:     tea.String("2018-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SetDcdnConfigOfVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetDcdnConfigOfVersion(request *SetDcdnConfigOfVersionRequest) (_result *SetDcdnConfigOfVersionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetDcdnConfigOfVersionResponse{}
	_body, _err := client.SetDcdnConfigOfVersionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetDcdnDomainCSRCertificateWithOptions(request *SetDcdnDomainCSRCertificateRequest, runtime *util.RuntimeOptions) (_result *SetDcdnDomainCSRCertificateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ServerCertificate)) {
		query["ServerCertificate"] = request.ServerCertificate
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SetDcdnDomainCSRCertificate"),
		Version:     tea.String("2018-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SetDcdnDomainCSRCertificateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetDcdnDomainCSRCertificate(request *SetDcdnDomainCSRCertificateRequest) (_result *SetDcdnDomainCSRCertificateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetDcdnDomainCSRCertificateResponse{}
	_body, _err := client.SetDcdnDomainCSRCertificateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetDcdnDomainCertificateWithOptions(request *SetDcdnDomainCertificateRequest, runtime *util.RuntimeOptions) (_result *SetDcdnDomainCertificateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CertName)) {
		query["CertName"] = request.CertName
	}

	if !tea.BoolValue(util.IsUnset(request.CertType)) {
		query["CertType"] = request.CertType
	}

	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.ForceSet)) {
		query["ForceSet"] = request.ForceSet
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.Region)) {
		query["Region"] = request.Region
	}

	if !tea.BoolValue(util.IsUnset(request.SSLPri)) {
		query["SSLPri"] = request.SSLPri
	}

	if !tea.BoolValue(util.IsUnset(request.SSLProtocol)) {
		query["SSLProtocol"] = request.SSLProtocol
	}

	if !tea.BoolValue(util.IsUnset(request.SSLPub)) {
		query["SSLPub"] = request.SSLPub
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SetDcdnDomainCertificate"),
		Version:     tea.String("2018-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SetDcdnDomainCertificateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetDcdnDomainCertificate(request *SetDcdnDomainCertificateRequest) (_result *SetDcdnDomainCertificateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetDcdnDomainCertificateResponse{}
	_body, _err := client.SetDcdnDomainCertificateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetDcdnDomainSMCertificateWithOptions(request *SetDcdnDomainSMCertificateRequest, runtime *util.RuntimeOptions) (_result *SetDcdnDomainSMCertificateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CertIdentifier)) {
		query["CertIdentifier"] = request.CertIdentifier
	}

	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SSLProtocol)) {
		query["SSLProtocol"] = request.SSLProtocol
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SetDcdnDomainSMCertificate"),
		Version:     tea.String("2018-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SetDcdnDomainSMCertificateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetDcdnDomainSMCertificate(request *SetDcdnDomainSMCertificateRequest) (_result *SetDcdnDomainSMCertificateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetDcdnDomainSMCertificateResponse{}
	_body, _err := client.SetDcdnDomainSMCertificateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetDcdnDomainStagingConfigWithOptions(request *SetDcdnDomainStagingConfigRequest, runtime *util.RuntimeOptions) (_result *SetDcdnDomainStagingConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.Functions)) {
		query["Functions"] = request.Functions
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SetDcdnDomainStagingConfig"),
		Version:     tea.String("2018-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SetDcdnDomainStagingConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetDcdnDomainStagingConfig(request *SetDcdnDomainStagingConfigRequest) (_result *SetDcdnDomainStagingConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetDcdnDomainStagingConfigResponse{}
	_body, _err := client.SetDcdnDomainStagingConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetDcdnFullDomainsBlockIPWithOptions(request *SetDcdnFullDomainsBlockIPRequest, runtime *util.RuntimeOptions) (_result *SetDcdnFullDomainsBlockIPResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BlockInterval)) {
		body["BlockInterval"] = request.BlockInterval
	}

	if !tea.BoolValue(util.IsUnset(request.IPList)) {
		body["IPList"] = request.IPList
	}

	if !tea.BoolValue(util.IsUnset(request.OperationType)) {
		body["OperationType"] = request.OperationType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("SetDcdnFullDomainsBlockIP"),
		Version:     tea.String("2018-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SetDcdnFullDomainsBlockIPResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetDcdnFullDomainsBlockIP(request *SetDcdnFullDomainsBlockIPRequest) (_result *SetDcdnFullDomainsBlockIPResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetDcdnFullDomainsBlockIPResponse{}
	_body, _err := client.SetDcdnFullDomainsBlockIPWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetDcdnUserConfigWithOptions(request *SetDcdnUserConfigRequest, runtime *util.RuntimeOptions) (_result *SetDcdnUserConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Configs)) {
		query["Configs"] = request.Configs
	}

	if !tea.BoolValue(util.IsUnset(request.FunctionId)) {
		query["FunctionId"] = request.FunctionId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SetDcdnUserConfig"),
		Version:     tea.String("2018-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SetDcdnUserConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetDcdnUserConfig(request *SetDcdnUserConfigRequest) (_result *SetDcdnUserConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetDcdnUserConfigResponse{}
	_body, _err := client.SetDcdnUserConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetRoutineSubdomainWithOptions(tmpReq *SetRoutineSubdomainRequest, runtime *util.RuntimeOptions) (_result *SetRoutineSubdomainResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &SetRoutineSubdomainShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Subdomains)) {
		request.SubdomainsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Subdomains, tea.String("Subdomains"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SubdomainsShrink)) {
		body["Subdomains"] = request.SubdomainsShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("SetRoutineSubdomain"),
		Version:     tea.String("2018-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SetRoutineSubdomainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetRoutineSubdomain(request *SetRoutineSubdomainRequest) (_result *SetRoutineSubdomainResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetRoutineSubdomainResponse{}
	_body, _err := client.SetRoutineSubdomainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StartDcdnDomainWithOptions(request *StartDcdnDomainRequest, runtime *util.RuntimeOptions) (_result *StartDcdnDomainResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("StartDcdnDomain"),
		Version:     tea.String("2018-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StartDcdnDomainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StartDcdnDomain(request *StartDcdnDomainRequest) (_result *StartDcdnDomainResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StartDcdnDomainResponse{}
	_body, _err := client.StartDcdnDomainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StartDcdnIpaDomainWithOptions(request *StartDcdnIpaDomainRequest, runtime *util.RuntimeOptions) (_result *StartDcdnIpaDomainResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("StartDcdnIpaDomain"),
		Version:     tea.String("2018-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StartDcdnIpaDomainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StartDcdnIpaDomain(request *StartDcdnIpaDomainRequest) (_result *StartDcdnIpaDomainResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StartDcdnIpaDomainResponse{}
	_body, _err := client.StartDcdnIpaDomainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StopDcdnDomainWithOptions(request *StopDcdnDomainRequest, runtime *util.RuntimeOptions) (_result *StopDcdnDomainResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("StopDcdnDomain"),
		Version:     tea.String("2018-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StopDcdnDomainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StopDcdnDomain(request *StopDcdnDomainRequest) (_result *StopDcdnDomainResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StopDcdnDomainResponse{}
	_body, _err := client.StopDcdnDomainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StopDcdnIpaDomainWithOptions(request *StopDcdnIpaDomainRequest, runtime *util.RuntimeOptions) (_result *StopDcdnIpaDomainResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("StopDcdnIpaDomain"),
		Version:     tea.String("2018-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StopDcdnIpaDomainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StopDcdnIpaDomain(request *StopDcdnIpaDomainRequest) (_result *StopDcdnIpaDomainResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StopDcdnIpaDomainResponse{}
	_body, _err := client.StopDcdnIpaDomainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) TagDcdnResourcesWithOptions(request *TagDcdnResourcesRequest, runtime *util.RuntimeOptions) (_result *TagDcdnResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("TagDcdnResources"),
		Version:     tea.String("2018-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &TagDcdnResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) TagDcdnResources(request *TagDcdnResourcesRequest) (_result *TagDcdnResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &TagDcdnResourcesResponse{}
	_body, _err := client.TagDcdnResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UntagDcdnResourcesWithOptions(request *UntagDcdnResourcesRequest, runtime *util.RuntimeOptions) (_result *UntagDcdnResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.All)) {
		query["All"] = request.All
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.TagKey)) {
		query["TagKey"] = request.TagKey
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UntagDcdnResources"),
		Version:     tea.String("2018-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UntagDcdnResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UntagDcdnResources(request *UntagDcdnResourcesRequest) (_result *UntagDcdnResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UntagDcdnResourcesResponse{}
	_body, _err := client.UntagDcdnResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateDcdnDeliverTaskWithOptions(request *UpdateDcdnDeliverTaskRequest, runtime *util.RuntimeOptions) (_result *UpdateDcdnDeliverTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Deliver)) {
		body["Deliver"] = request.Deliver
	}

	if !tea.BoolValue(util.IsUnset(request.DeliverId)) {
		body["DeliverId"] = request.DeliverId
	}

	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		body["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Reports)) {
		body["Reports"] = request.Reports
	}

	if !tea.BoolValue(util.IsUnset(request.Schedule)) {
		body["Schedule"] = request.Schedule
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateDcdnDeliverTask"),
		Version:     tea.String("2018-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateDcdnDeliverTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateDcdnDeliverTask(request *UpdateDcdnDeliverTaskRequest) (_result *UpdateDcdnDeliverTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateDcdnDeliverTaskResponse{}
	_body, _err := client.UpdateDcdnDeliverTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateDcdnDomainWithOptions(request *UpdateDcdnDomainRequest, runtime *util.RuntimeOptions) (_result *UpdateDcdnDomainResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !tea.BoolValue(util.IsUnset(request.Sources)) {
		query["Sources"] = request.Sources
	}

	if !tea.BoolValue(util.IsUnset(request.TopLevelDomain)) {
		query["TopLevelDomain"] = request.TopLevelDomain
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateDcdnDomain"),
		Version:     tea.String("2018-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateDcdnDomainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateDcdnDomain(request *UpdateDcdnDomainRequest) (_result *UpdateDcdnDomainResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateDcdnDomainResponse{}
	_body, _err := client.UpdateDcdnDomainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateDcdnIpaDomainWithOptions(request *UpdateDcdnIpaDomainRequest, runtime *util.RuntimeOptions) (_result *UpdateDcdnIpaDomainResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !tea.BoolValue(util.IsUnset(request.Sources)) {
		query["Sources"] = request.Sources
	}

	if !tea.BoolValue(util.IsUnset(request.TopLevelDomain)) {
		query["TopLevelDomain"] = request.TopLevelDomain
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateDcdnIpaDomain"),
		Version:     tea.String("2018-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateDcdnIpaDomainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateDcdnIpaDomain(request *UpdateDcdnIpaDomainRequest) (_result *UpdateDcdnIpaDomainResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateDcdnIpaDomainResponse{}
	_body, _err := client.UpdateDcdnIpaDomainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateDcdnSLSRealtimeLogDeliveryWithOptions(request *UpdateDcdnSLSRealtimeLogDeliveryRequest, runtime *util.RuntimeOptions) (_result *UpdateDcdnSLSRealtimeLogDeliveryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DataCenter)) {
		body["DataCenter"] = request.DataCenter
	}

	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		body["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectName)) {
		body["ProjectName"] = request.ProjectName
	}

	if !tea.BoolValue(util.IsUnset(request.SLSLogStore)) {
		body["SLSLogStore"] = request.SLSLogStore
	}

	if !tea.BoolValue(util.IsUnset(request.SLSProject)) {
		body["SLSProject"] = request.SLSProject
	}

	if !tea.BoolValue(util.IsUnset(request.SLSRegion)) {
		body["SLSRegion"] = request.SLSRegion
	}

	if !tea.BoolValue(util.IsUnset(request.SamplingRate)) {
		body["SamplingRate"] = request.SamplingRate
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateDcdnSLSRealtimeLogDelivery"),
		Version:     tea.String("2018-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateDcdnSLSRealtimeLogDeliveryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateDcdnSLSRealtimeLogDelivery(request *UpdateDcdnSLSRealtimeLogDeliveryRequest) (_result *UpdateDcdnSLSRealtimeLogDeliveryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateDcdnSLSRealtimeLogDeliveryResponse{}
	_body, _err := client.UpdateDcdnSLSRealtimeLogDeliveryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateDcdnSubTaskWithOptions(request *UpdateDcdnSubTaskRequest, runtime *util.RuntimeOptions) (_result *UpdateDcdnSubTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		body["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		body["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.ReportIds)) {
		body["ReportIds"] = request.ReportIds
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		body["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateDcdnSubTask"),
		Version:     tea.String("2018-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateDcdnSubTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateDcdnSubTask(request *UpdateDcdnSubTaskRequest) (_result *UpdateDcdnSubTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateDcdnSubTaskResponse{}
	_body, _err := client.UpdateDcdnSubTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateDcdnUserRealTimeDeliveryFieldWithOptions(request *UpdateDcdnUserRealTimeDeliveryFieldRequest, runtime *util.RuntimeOptions) (_result *UpdateDcdnUserRealTimeDeliveryFieldResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateDcdnUserRealTimeDeliveryField"),
		Version:     tea.String("2018-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateDcdnUserRealTimeDeliveryFieldResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateDcdnUserRealTimeDeliveryField(request *UpdateDcdnUserRealTimeDeliveryFieldRequest) (_result *UpdateDcdnUserRealTimeDeliveryFieldResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateDcdnUserRealTimeDeliveryFieldResponse{}
	_body, _err := client.UpdateDcdnUserRealTimeDeliveryFieldWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UploadRoutineCodeWithOptions(request *UploadRoutineCodeRequest, runtime *util.RuntimeOptions) (_result *UploadRoutineCodeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CodeDescription)) {
		body["CodeDescription"] = request.CodeDescription
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UploadRoutineCode"),
		Version:     tea.String("2018-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UploadRoutineCodeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UploadRoutineCode(request *UploadRoutineCodeRequest) (_result *UploadRoutineCodeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UploadRoutineCodeResponse{}
	_body, _err := client.UploadRoutineCodeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UploadStagingRoutineCodeWithOptions(request *UploadStagingRoutineCodeRequest, runtime *util.RuntimeOptions) (_result *UploadStagingRoutineCodeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CodeDescription)) {
		body["CodeDescription"] = request.CodeDescription
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UploadStagingRoutineCode"),
		Version:     tea.String("2018-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UploadStagingRoutineCodeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UploadStagingRoutineCode(request *UploadStagingRoutineCodeRequest) (_result *UploadStagingRoutineCodeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UploadStagingRoutineCodeResponse{}
	_body, _err := client.UploadStagingRoutineCodeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) VerifyDcdnDomainOwnerWithOptions(request *VerifyDcdnDomainOwnerRequest, runtime *util.RuntimeOptions) (_result *VerifyDcdnDomainOwnerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.VerifyType)) {
		query["VerifyType"] = request.VerifyType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("VerifyDcdnDomainOwner"),
		Version:     tea.String("2018-01-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &VerifyDcdnDomainOwnerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) VerifyDcdnDomainOwner(request *VerifyDcdnDomainOwnerRequest) (_result *VerifyDcdnDomainOwnerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &VerifyDcdnDomainOwnerResponse{}
	_body, _err := client.VerifyDcdnDomainOwnerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
