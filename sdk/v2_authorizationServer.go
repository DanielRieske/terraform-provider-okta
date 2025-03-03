package sdk

import (
	"context"
	"fmt"
	"time"

	"github.com/okta/terraform-provider-okta/sdk/query"
)

type AuthorizationServerResource resource

type AuthorizationServer struct {
	Links       interface{}                     `json:"_links,omitempty"`
	Audiences   []string                        `json:"audiences,omitempty"`
	Created     *time.Time                      `json:"created,omitempty"`
	Credentials *AuthorizationServerCredentials `json:"credentials,omitempty"`
	Default     *bool                           `json:"default,omitempty"`
	Description string                          `json:"description,omitempty"`
	Id          string                          `json:"id,omitempty"`
	Issuer      string                          `json:"issuer,omitempty"`
	IssuerMode  string                          `json:"issuerMode,omitempty"`
	LastUpdated *time.Time                      `json:"lastUpdated,omitempty"`
	Name        string                          `json:"name,omitempty"`
	Status      string                          `json:"status,omitempty"`
}

func (m *AuthorizationServerResource) CreateAuthorizationServer(ctx context.Context, body AuthorizationServer) (*AuthorizationServer, *Response, error) {
	url := "/api/v1/authorizationServers"

	rq := m.client.CloneRequestExecutor()

	req, err := rq.WithAccept("application/json").WithContentType("application/json").NewRequest("POST", url, body)
	if err != nil {
		return nil, nil, err
	}

	var authorizationServer *AuthorizationServer

	resp, err := rq.Do(ctx, req, &authorizationServer)
	if err != nil {
		return nil, resp, err
	}

	return authorizationServer, resp, nil
}

func (m *AuthorizationServerResource) GetAuthorizationServer(ctx context.Context, authServerId string) (*AuthorizationServer, *Response, error) {
	url := fmt.Sprintf("/api/v1/authorizationServers/%v", authServerId)

	rq := m.client.CloneRequestExecutor()

	req, err := rq.WithAccept("application/json").WithContentType("application/json").NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	var authorizationServer *AuthorizationServer

	resp, err := rq.Do(ctx, req, &authorizationServer)
	if err != nil {
		return nil, resp, err
	}

	return authorizationServer, resp, nil
}

func (m *AuthorizationServerResource) UpdateAuthorizationServer(ctx context.Context, authServerId string, body AuthorizationServer) (*AuthorizationServer, *Response, error) {
	url := fmt.Sprintf("/api/v1/authorizationServers/%v", authServerId)

	rq := m.client.CloneRequestExecutor()

	req, err := rq.WithAccept("application/json").WithContentType("application/json").NewRequest("PUT", url, body)
	if err != nil {
		return nil, nil, err
	}

	var authorizationServer *AuthorizationServer

	resp, err := rq.Do(ctx, req, &authorizationServer)
	if err != nil {
		return nil, resp, err
	}

	return authorizationServer, resp, nil
}

func (m *AuthorizationServerResource) DeleteAuthorizationServer(ctx context.Context, authServerId string) (*Response, error) {
	url := fmt.Sprintf("/api/v1/authorizationServers/%v", authServerId)

	rq := m.client.CloneRequestExecutor()

	req, err := rq.WithAccept("application/json").WithContentType("application/json").NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := m.client.requestExecutor.Do(ctx, req, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

func (m *AuthorizationServerResource) ListAuthorizationServers(ctx context.Context, qp *query.Params) ([]*AuthorizationServer, *Response, error) {
	url := "/api/v1/authorizationServers"
	if qp != nil {
		url = url + qp.String()
	}

	rq := m.client.CloneRequestExecutor()

	req, err := rq.WithAccept("application/json").WithContentType("application/json").NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	var authorizationServer []*AuthorizationServer

	resp, err := rq.Do(ctx, req, &authorizationServer)
	if err != nil {
		return nil, resp, err
	}

	return authorizationServer, resp, nil
}

func (m *AuthorizationServerResource) ListOAuth2Claims(ctx context.Context, authServerId string) ([]*OAuth2Claim, *Response, error) {
	url := fmt.Sprintf("/api/v1/authorizationServers/%v/claims", authServerId)

	rq := m.client.CloneRequestExecutor()

	req, err := rq.WithAccept("application/json").WithContentType("application/json").NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	var oAuth2Claim []*OAuth2Claim

	resp, err := rq.Do(ctx, req, &oAuth2Claim)
	if err != nil {
		return nil, resp, err
	}

	return oAuth2Claim, resp, nil
}

func (m *AuthorizationServerResource) CreateOAuth2Claim(ctx context.Context, authServerId string, body OAuth2Claim) (*OAuth2Claim, *Response, error) {
	url := fmt.Sprintf("/api/v1/authorizationServers/%v/claims", authServerId)

	rq := m.client.CloneRequestExecutor()

	req, err := rq.WithAccept("application/json").WithContentType("application/json").NewRequest("POST", url, body)
	if err != nil {
		return nil, nil, err
	}

	var oAuth2Claim *OAuth2Claim

	resp, err := rq.Do(ctx, req, &oAuth2Claim)
	if err != nil {
		return nil, resp, err
	}

	return oAuth2Claim, resp, nil
}

func (m *AuthorizationServerResource) DeleteOAuth2Claim(ctx context.Context, authServerId, claimId string) (*Response, error) {
	url := fmt.Sprintf("/api/v1/authorizationServers/%v/claims/%v", authServerId, claimId)

	rq := m.client.CloneRequestExecutor()

	req, err := rq.WithAccept("application/json").WithContentType("application/json").NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := m.client.requestExecutor.Do(ctx, req, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

func (m *AuthorizationServerResource) GetOAuth2Claim(ctx context.Context, authServerId, claimId string) (*OAuth2Claim, *Response, error) {
	url := fmt.Sprintf("/api/v1/authorizationServers/%v/claims/%v", authServerId, claimId)

	rq := m.client.CloneRequestExecutor()

	req, err := rq.WithAccept("application/json").WithContentType("application/json").NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	var oAuth2Claim *OAuth2Claim

	resp, err := rq.Do(ctx, req, &oAuth2Claim)
	if err != nil {
		return nil, resp, err
	}

	return oAuth2Claim, resp, nil
}

func (m *AuthorizationServerResource) UpdateOAuth2Claim(ctx context.Context, authServerId, claimId string, body OAuth2Claim) (*OAuth2Claim, *Response, error) {
	url := fmt.Sprintf("/api/v1/authorizationServers/%v/claims/%v", authServerId, claimId)

	rq := m.client.CloneRequestExecutor()

	req, err := rq.WithAccept("application/json").WithContentType("application/json").NewRequest("PUT", url, body)
	if err != nil {
		return nil, nil, err
	}

	var oAuth2Claim *OAuth2Claim

	resp, err := rq.Do(ctx, req, &oAuth2Claim)
	if err != nil {
		return nil, resp, err
	}

	return oAuth2Claim, resp, nil
}

func (m *AuthorizationServerResource) ListOAuth2ClientsForAuthorizationServer(ctx context.Context, authServerId string) ([]*OAuth2Client, *Response, error) {
	url := fmt.Sprintf("/api/v1/authorizationServers/%v/clients", authServerId)

	rq := m.client.CloneRequestExecutor()

	req, err := rq.WithAccept("application/json").WithContentType("application/json").NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	var oAuth2Client []*OAuth2Client

	resp, err := rq.Do(ctx, req, &oAuth2Client)
	if err != nil {
		return nil, resp, err
	}

	return oAuth2Client, resp, nil
}

func (m *AuthorizationServerResource) RevokeRefreshTokensForAuthorizationServerAndClient(ctx context.Context, authServerId, clientId string) (*Response, error) {
	url := fmt.Sprintf("/api/v1/authorizationServers/%v/clients/%v/tokens", authServerId, clientId)

	rq := m.client.CloneRequestExecutor()

	req, err := rq.WithAccept("application/json").WithContentType("application/json").NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := m.client.requestExecutor.Do(ctx, req, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

func (m *AuthorizationServerResource) ListRefreshTokensForAuthorizationServerAndClient(ctx context.Context, authServerId, clientId string, qp *query.Params) ([]*OAuth2RefreshToken, *Response, error) {
	url := fmt.Sprintf("/api/v1/authorizationServers/%v/clients/%v/tokens", authServerId, clientId)
	if qp != nil {
		url = url + qp.String()
	}

	rq := m.client.CloneRequestExecutor()

	req, err := rq.WithAccept("application/json").WithContentType("application/json").NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	var oAuth2RefreshToken []*OAuth2RefreshToken

	resp, err := rq.Do(ctx, req, &oAuth2RefreshToken)
	if err != nil {
		return nil, resp, err
	}

	return oAuth2RefreshToken, resp, nil
}

func (m *AuthorizationServerResource) RevokeRefreshTokenForAuthorizationServerAndClient(ctx context.Context, authServerId, clientId, tokenId string) (*Response, error) {
	url := fmt.Sprintf("/api/v1/authorizationServers/%v/clients/%v/tokens/%v", authServerId, clientId, tokenId)

	rq := m.client.CloneRequestExecutor()

	req, err := rq.WithAccept("application/json").WithContentType("application/json").NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := m.client.requestExecutor.Do(ctx, req, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

func (m *AuthorizationServerResource) GetRefreshTokenForAuthorizationServerAndClient(ctx context.Context, authServerId, clientId, tokenId string, qp *query.Params) (*OAuth2RefreshToken, *Response, error) {
	url := fmt.Sprintf("/api/v1/authorizationServers/%v/clients/%v/tokens/%v", authServerId, clientId, tokenId)
	if qp != nil {
		url = url + qp.String()
	}

	rq := m.client.CloneRequestExecutor()

	req, err := rq.WithAccept("application/json").WithContentType("application/json").NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	var oAuth2RefreshToken *OAuth2RefreshToken

	resp, err := rq.Do(ctx, req, &oAuth2RefreshToken)
	if err != nil {
		return nil, resp, err
	}

	return oAuth2RefreshToken, resp, nil
}

func (m *AuthorizationServerResource) ListAuthorizationServerKeys(ctx context.Context, authServerId string) ([]*JsonWebKey, *Response, error) {
	url := fmt.Sprintf("/api/v1/authorizationServers/%v/credentials/keys", authServerId)

	rq := m.client.CloneRequestExecutor()

	req, err := rq.WithAccept("application/json").WithContentType("application/json").NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	var jsonWebKey []*JsonWebKey

	resp, err := rq.Do(ctx, req, &jsonWebKey)
	if err != nil {
		return nil, resp, err
	}

	return jsonWebKey, resp, nil
}

func (m *AuthorizationServerResource) RotateAuthorizationServerKeys(ctx context.Context, authServerId string, body JwkUse) ([]*JsonWebKey, *Response, error) {
	url := fmt.Sprintf("/api/v1/authorizationServers/%v/credentials/lifecycle/keyRotate", authServerId)

	rq := m.client.CloneRequestExecutor()

	req, err := rq.WithAccept("application/json").WithContentType("application/json").NewRequest("POST", url, body)
	if err != nil {
		return nil, nil, err
	}

	var jsonWebKey []*JsonWebKey

	resp, err := rq.Do(ctx, req, &jsonWebKey)
	if err != nil {
		return nil, resp, err
	}

	return jsonWebKey, resp, nil
}

func (m *AuthorizationServerResource) ActivateAuthorizationServer(ctx context.Context, authServerId string) (*Response, error) {
	url := fmt.Sprintf("/api/v1/authorizationServers/%v/lifecycle/activate", authServerId)

	rq := m.client.CloneRequestExecutor()

	req, err := rq.WithAccept("application/json").WithContentType("application/json").NewRequest("POST", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := m.client.requestExecutor.Do(ctx, req, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

func (m *AuthorizationServerResource) DeactivateAuthorizationServer(ctx context.Context, authServerId string) (*Response, error) {
	url := fmt.Sprintf("/api/v1/authorizationServers/%v/lifecycle/deactivate", authServerId)

	rq := m.client.CloneRequestExecutor()

	req, err := rq.WithAccept("application/json").WithContentType("application/json").NewRequest("POST", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := m.client.requestExecutor.Do(ctx, req, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

func (m *AuthorizationServerResource) ListAuthorizationServerPolicies(ctx context.Context, authServerId string) ([]*AuthorizationServerPolicy, *Response, error) {
	url := fmt.Sprintf("/api/v1/authorizationServers/%v/policies", authServerId)

	rq := m.client.CloneRequestExecutor()

	req, err := rq.WithAccept("application/json").WithContentType("application/json").NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	var authorizationServerPolicy []*AuthorizationServerPolicy

	resp, err := rq.Do(ctx, req, &authorizationServerPolicy)
	if err != nil {
		return nil, resp, err
	}

	return authorizationServerPolicy, resp, nil
}

func (m *AuthorizationServerResource) CreateAuthorizationServerPolicy(ctx context.Context, authServerId string, body AuthorizationServerPolicy) (*AuthorizationServerPolicy, *Response, error) {
	url := fmt.Sprintf("/api/v1/authorizationServers/%v/policies", authServerId)

	rq := m.client.CloneRequestExecutor()

	req, err := rq.WithAccept("application/json").WithContentType("application/json").NewRequest("POST", url, body)
	if err != nil {
		return nil, nil, err
	}

	var authorizationServerPolicy *AuthorizationServerPolicy

	resp, err := rq.Do(ctx, req, &authorizationServerPolicy)
	if err != nil {
		return nil, resp, err
	}

	return authorizationServerPolicy, resp, nil
}

func (m *AuthorizationServerResource) DeleteAuthorizationServerPolicy(ctx context.Context, authServerId, policyId string) (*Response, error) {
	url := fmt.Sprintf("/api/v1/authorizationServers/%v/policies/%v", authServerId, policyId)

	rq := m.client.CloneRequestExecutor()

	req, err := rq.WithAccept("application/json").WithContentType("application/json").NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := m.client.requestExecutor.Do(ctx, req, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

func (m *AuthorizationServerResource) GetAuthorizationServerPolicy(ctx context.Context, authServerId, policyId string) (*AuthorizationServerPolicy, *Response, error) {
	url := fmt.Sprintf("/api/v1/authorizationServers/%v/policies/%v", authServerId, policyId)

	rq := m.client.CloneRequestExecutor()

	req, err := rq.WithAccept("application/json").WithContentType("application/json").NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	var authorizationServerPolicy *AuthorizationServerPolicy

	resp, err := rq.Do(ctx, req, &authorizationServerPolicy)
	if err != nil {
		return nil, resp, err
	}

	return authorizationServerPolicy, resp, nil
}

func (m *AuthorizationServerResource) UpdateAuthorizationServerPolicy(ctx context.Context, authServerId, policyId string, body AuthorizationServerPolicy) (*AuthorizationServerPolicy, *Response, error) {
	url := fmt.Sprintf("/api/v1/authorizationServers/%v/policies/%v", authServerId, policyId)

	rq := m.client.CloneRequestExecutor()

	req, err := rq.WithAccept("application/json").WithContentType("application/json").NewRequest("PUT", url, body)
	if err != nil {
		return nil, nil, err
	}

	var authorizationServerPolicy *AuthorizationServerPolicy

	resp, err := rq.Do(ctx, req, &authorizationServerPolicy)
	if err != nil {
		return nil, resp, err
	}

	return authorizationServerPolicy, resp, nil
}

// Activate Authorization Server Policy
func (m *AuthorizationServerResource) ActivateAuthorizationServerPolicy(ctx context.Context, authServerId, policyId string) (*Response, error) {
	url := fmt.Sprintf("/api/v1/authorizationServers/%v/policies/%v/lifecycle/activate", authServerId, policyId)

	rq := m.client.CloneRequestExecutor()

	req, err := rq.WithAccept("application/json").WithContentType("application/json").NewRequest("POST", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := m.client.requestExecutor.Do(ctx, req, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// Deactivate Authorization Server Policy
func (m *AuthorizationServerResource) DeactivateAuthorizationServerPolicy(ctx context.Context, authServerId, policyId string) (*Response, error) {
	url := fmt.Sprintf("/api/v1/authorizationServers/%v/policies/%v/lifecycle/deactivate", authServerId, policyId)

	rq := m.client.CloneRequestExecutor()

	req, err := rq.WithAccept("application/json").WithContentType("application/json").NewRequest("POST", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := m.client.requestExecutor.Do(ctx, req, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// Enumerates all policy rules for the specified Custom Authorization Server and Policy.
func (m *AuthorizationServerResource) ListAuthorizationServerPolicyRules(ctx context.Context, authServerId, policyId string) ([]*AuthorizationServerPolicyRule, *Response, error) {
	url := fmt.Sprintf("/api/v1/authorizationServers/%v/policies/%v/rules", authServerId, policyId)

	rq := m.client.CloneRequestExecutor()

	req, err := rq.WithAccept("application/json").WithContentType("application/json").NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	var authorizationServerPolicyRule []*AuthorizationServerPolicyRule

	resp, err := rq.Do(ctx, req, &authorizationServerPolicyRule)
	if err != nil {
		return nil, resp, err
	}

	return authorizationServerPolicyRule, resp, nil
}

// Creates a policy rule for the specified Custom Authorization Server and Policy.
func (m *AuthorizationServerResource) CreateAuthorizationServerPolicyRule(ctx context.Context, authServerId, policyId string, body AuthorizationServerPolicyRule) (*AuthorizationServerPolicyRule, *Response, error) {
	url := fmt.Sprintf("/api/v1/authorizationServers/%v/policies/%v/rules", authServerId, policyId)

	rq := m.client.CloneRequestExecutor()

	req, err := rq.WithAccept("application/json").WithContentType("application/json").NewRequest("POST", url, body)
	if err != nil {
		return nil, nil, err
	}

	var authorizationServerPolicyRule *AuthorizationServerPolicyRule

	resp, err := rq.Do(ctx, req, &authorizationServerPolicyRule)
	if err != nil {
		return nil, resp, err
	}

	return authorizationServerPolicyRule, resp, nil
}

// Deletes a Policy Rule defined in the specified Custom Authorization Server and Policy.
func (m *AuthorizationServerResource) DeleteAuthorizationServerPolicyRule(ctx context.Context, authServerId, policyId, ruleId string) (*Response, error) {
	url := fmt.Sprintf("/api/v1/authorizationServers/%v/policies/%v/rules/%v", authServerId, policyId, ruleId)

	rq := m.client.CloneRequestExecutor()

	req, err := rq.WithAccept("application/json").WithContentType("application/json").NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := m.client.requestExecutor.Do(ctx, req, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// Returns a Policy Rule by ID that is defined in the specified Custom Authorization Server and Policy.
func (m *AuthorizationServerResource) GetAuthorizationServerPolicyRule(ctx context.Context, authServerId, policyId, ruleId string) (*AuthorizationServerPolicyRule, *Response, error) {
	url := fmt.Sprintf("/api/v1/authorizationServers/%v/policies/%v/rules/%v", authServerId, policyId, ruleId)

	rq := m.client.CloneRequestExecutor()

	req, err := rq.WithAccept("application/json").WithContentType("application/json").NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	var authorizationServerPolicyRule *AuthorizationServerPolicyRule

	resp, err := rq.Do(ctx, req, &authorizationServerPolicyRule)
	if err != nil {
		return nil, resp, err
	}

	return authorizationServerPolicyRule, resp, nil
}

// Updates the configuration of the Policy Rule defined in the specified Custom Authorization Server and Policy.
func (m *AuthorizationServerResource) UpdateAuthorizationServerPolicyRule(ctx context.Context, authServerId, policyId, ruleId string, body AuthorizationServerPolicyRule) (*AuthorizationServerPolicyRule, *Response, error) {
	url := fmt.Sprintf("/api/v1/authorizationServers/%v/policies/%v/rules/%v", authServerId, policyId, ruleId)

	rq := m.client.CloneRequestExecutor()

	req, err := rq.WithAccept("application/json").WithContentType("application/json").NewRequest("PUT", url, body)
	if err != nil {
		return nil, nil, err
	}

	var authorizationServerPolicyRule *AuthorizationServerPolicyRule

	resp, err := rq.Do(ctx, req, &authorizationServerPolicyRule)
	if err != nil {
		return nil, resp, err
	}

	return authorizationServerPolicyRule, resp, nil
}

// Activate Authorization Server Policy Rule
func (m *AuthorizationServerResource) ActivateAuthorizationServerPolicyRule(ctx context.Context, authServerId, policyId, ruleId string) (*Response, error) {
	url := fmt.Sprintf("/api/v1/authorizationServers/%v/policies/%v/rules/%v/lifecycle/activate", authServerId, policyId, ruleId)

	rq := m.client.CloneRequestExecutor()

	req, err := rq.WithAccept("application/json").WithContentType("application/json").NewRequest("POST", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := m.client.requestExecutor.Do(ctx, req, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// Deactivate Authorization Server Policy Rule
func (m *AuthorizationServerResource) DeactivateAuthorizationServerPolicyRule(ctx context.Context, authServerId, policyId, ruleId string) (*Response, error) {
	url := fmt.Sprintf("/api/v1/authorizationServers/%v/policies/%v/rules/%v/lifecycle/deactivate", authServerId, policyId, ruleId)

	rq := m.client.CloneRequestExecutor()

	req, err := rq.WithAccept("application/json").WithContentType("application/json").NewRequest("POST", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := m.client.requestExecutor.Do(ctx, req, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

func (m *AuthorizationServerResource) ListOAuth2Scopes(ctx context.Context, authServerId string, qp *query.Params) ([]*OAuth2Scope, *Response, error) {
	url := fmt.Sprintf("/api/v1/authorizationServers/%v/scopes", authServerId)
	if qp != nil {
		url = url + qp.String()
	}

	rq := m.client.CloneRequestExecutor()

	req, err := rq.WithAccept("application/json").WithContentType("application/json").NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	var oAuth2Scope []*OAuth2Scope

	resp, err := rq.Do(ctx, req, &oAuth2Scope)
	if err != nil {
		return nil, resp, err
	}

	return oAuth2Scope, resp, nil
}

func (m *AuthorizationServerResource) CreateOAuth2Scope(ctx context.Context, authServerId string, body OAuth2Scope) (*OAuth2Scope, *Response, error) {
	url := fmt.Sprintf("/api/v1/authorizationServers/%v/scopes", authServerId)

	rq := m.client.CloneRequestExecutor()

	req, err := rq.WithAccept("application/json").WithContentType("application/json").NewRequest("POST", url, body)
	if err != nil {
		return nil, nil, err
	}

	var oAuth2Scope *OAuth2Scope

	resp, err := rq.Do(ctx, req, &oAuth2Scope)
	if err != nil {
		return nil, resp, err
	}

	return oAuth2Scope, resp, nil
}

func (m *AuthorizationServerResource) DeleteOAuth2Scope(ctx context.Context, authServerId, scopeId string) (*Response, error) {
	url := fmt.Sprintf("/api/v1/authorizationServers/%v/scopes/%v", authServerId, scopeId)

	rq := m.client.CloneRequestExecutor()

	req, err := rq.WithAccept("application/json").WithContentType("application/json").NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := m.client.requestExecutor.Do(ctx, req, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

func (m *AuthorizationServerResource) GetOAuth2Scope(ctx context.Context, authServerId, scopeId string) (*OAuth2Scope, *Response, error) {
	url := fmt.Sprintf("/api/v1/authorizationServers/%v/scopes/%v", authServerId, scopeId)

	rq := m.client.CloneRequestExecutor()

	req, err := rq.WithAccept("application/json").WithContentType("application/json").NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	var oAuth2Scope *OAuth2Scope

	resp, err := rq.Do(ctx, req, &oAuth2Scope)
	if err != nil {
		return nil, resp, err
	}

	return oAuth2Scope, resp, nil
}

func (m *AuthorizationServerResource) UpdateOAuth2Scope(ctx context.Context, authServerId, scopeId string, body OAuth2Scope) (*OAuth2Scope, *Response, error) {
	url := fmt.Sprintf("/api/v1/authorizationServers/%v/scopes/%v", authServerId, scopeId)

	rq := m.client.CloneRequestExecutor()

	req, err := rq.WithAccept("application/json").WithContentType("application/json").NewRequest("PUT", url, body)
	if err != nil {
		return nil, nil, err
	}

	var oAuth2Scope *OAuth2Scope

	resp, err := rq.Do(ctx, req, &oAuth2Scope)
	if err != nil {
		return nil, resp, err
	}

	return oAuth2Scope, resp, nil
}
