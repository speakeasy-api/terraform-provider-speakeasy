// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package sdk

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"speakeasy/internal/sdk/pkg/models/operations"
	"speakeasy/internal/sdk/pkg/models/shared"
	"speakeasy/internal/sdk/pkg/utils"
)

// apiEndpoints - REST APIs for managing ApiEndpoint entities
type apiEndpoints struct {
	defaultClient  HTTPClient
	securityClient HTTPClient
	serverURL      string
	language       string
	sdkVersion     string
	genVersion     string
}

func newAPIEndpoints(defaultClient, securityClient HTTPClient, serverURL, language, sdkVersion, genVersion string) *apiEndpoints {
	return &apiEndpoints{
		defaultClient:  defaultClient,
		securityClient: securityClient,
		serverURL:      serverURL,
		language:       language,
		sdkVersion:     sdkVersion,
		genVersion:     genVersion,
	}
}

// DeleteAPIEndpoint - Delete an ApiEndpoint.
// Delete an ApiEndpoint. This will also delete all associated Request Logs (if using a Postgres datastore).
func (s *apiEndpoints) DeleteAPIEndpoint(ctx context.Context, request operations.DeleteAPIEndpointRequest) (*operations.DeleteAPIEndpointResponse, error) {
	baseURL := s.serverURL
	url := utils.GenerateURL(ctx, baseURL, "/v1/apis/{apiID}/version/{versionID}/api_endpoints/{apiEndpointID}", request, nil)

	req, err := http.NewRequestWithContext(ctx, "DELETE", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	client := s.securityClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}
	defer httpRes.Body.Close()

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.DeleteAPIEndpointResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
	default:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.Error
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.Error = out
		}
	}

	return res, nil
}

// FindAPIEndpoint - Find an ApiEndpoint via its displayName.
// Find an ApiEndpoint via its displayName (set by operationId from a registered OpenAPI schema).
// This is useful for finding the ID of an ApiEndpoint to use in the /v1/apis/{apiID}/version/{versionID}/api_endpoints/{apiEndpointID} endpoints.
func (s *apiEndpoints) FindAPIEndpoint(ctx context.Context, request operations.FindAPIEndpointRequest) (*operations.FindAPIEndpointResponse, error) {
	baseURL := s.serverURL
	url := utils.GenerateURL(ctx, baseURL, "/v1/apis/{apiID}/version/{versionID}/api_endpoints/find/{displayName}", request, nil)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	client := s.securityClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}
	defer httpRes.Body.Close()

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.FindAPIEndpointResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.APIEndpoint
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.APIEndpoint = out
		}
	default:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.Error
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.Error = out
		}
	}

	return res, nil
}

// GenerateOpenAPISpecForAPIEndpoint - Generate an OpenAPI specification for a particular ApiEndpoint.
// This endpoint will generate a new operation in any registered OpenAPI document if the operation does not already exist in the document.
// Returns the original document and the newly generated document allowing a diff to be performed to see what has changed.
func (s *apiEndpoints) GenerateOpenAPISpecForAPIEndpoint(ctx context.Context, request operations.GenerateOpenAPISpecForAPIEndpointRequest) (*operations.GenerateOpenAPISpecForAPIEndpointResponse, error) {
	baseURL := s.serverURL
	url := utils.GenerateURL(ctx, baseURL, "/v1/apis/{apiID}/version/{versionID}/api_endpoints/{apiEndpointID}/generate/openapi", request, nil)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	client := s.securityClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}
	defer httpRes.Body.Close()

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.GenerateOpenAPISpecForAPIEndpointResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.GenerateOpenAPISpecDiff
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.GenerateOpenAPISpecDiff = out
		}
	default:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.Error
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.Error = out
		}
	}

	return res, nil
}

// GeneratePostmanCollectionForAPIEndpoint - Generate a Postman collection for a particular ApiEndpoint.
// Generates a postman collection that allows the endpoint to be called from postman variables produced for any path/query/header parameters included in the OpenAPI document.
func (s *apiEndpoints) GeneratePostmanCollectionForAPIEndpoint(ctx context.Context, request operations.GeneratePostmanCollectionForAPIEndpointRequest) (*operations.GeneratePostmanCollectionForAPIEndpointResponse, error) {
	baseURL := s.serverURL
	url := utils.GenerateURL(ctx, baseURL, "/v1/apis/{apiID}/version/{versionID}/api_endpoints/{apiEndpointID}/generate/postman", request, nil)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	client := s.securityClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}
	defer httpRes.Body.Close()

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.GeneratePostmanCollectionForAPIEndpointResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/octet-stream`):
			out, err := io.ReadAll(httpRes.Body)
			if err != nil {
				return nil, fmt.Errorf("error reading response body: %w", err)
			}

			res.PostmanCollection = out
		}
	default:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.Error
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.Error = out
		}
	}

	return res, nil
}

// GetAllAPIEndpoints - Get all Api endpoints for a particular apiID.
func (s *apiEndpoints) GetAllAPIEndpoints(ctx context.Context, request operations.GetAllAPIEndpointsRequest) (*operations.GetAllAPIEndpointsResponse, error) {
	baseURL := s.serverURL
	url := utils.GenerateURL(ctx, baseURL, "/v1/apis/{apiID}/api_endpoints", request, nil)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	client := s.securityClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}
	defer httpRes.Body.Close()

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.GetAllAPIEndpointsResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out []shared.APIEndpoint
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.APIEndpoints = out
		}
	default:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.Error
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.Error = out
		}
	}

	return res, nil
}

// GetAllForVersionAPIEndpoints - Get all ApiEndpoints for a particular apiID and versionID.
func (s *apiEndpoints) GetAllForVersionAPIEndpoints(ctx context.Context, request operations.GetAllForVersionAPIEndpointsRequest) (*operations.GetAllForVersionAPIEndpointsResponse, error) {
	baseURL := s.serverURL
	url := utils.GenerateURL(ctx, baseURL, "/v1/apis/{apiID}/version/{versionID}/api_endpoints", request, nil)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	client := s.securityClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}
	defer httpRes.Body.Close()

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.GetAllForVersionAPIEndpointsResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out []shared.APIEndpoint
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.APIEndpoints = out
		}
	default:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.Error
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.Error = out
		}
	}

	return res, nil
}

// GetAPIEndpoint - Get an ApiEndpoint.
func (s *apiEndpoints) GetAPIEndpoint(ctx context.Context, request operations.GetAPIEndpointRequest) (*operations.GetAPIEndpointResponse, error) {
	baseURL := s.serverURL
	url := utils.GenerateURL(ctx, baseURL, "/v1/apis/{apiID}/version/{versionID}/api_endpoints/{apiEndpointID}", request, nil)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	client := s.securityClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}
	defer httpRes.Body.Close()

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.GetAPIEndpointResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.APIEndpoint
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.APIEndpoint = out
		}
	default:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.Error
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.Error = out
		}
	}

	return res, nil
}

// UpsertAPIEndpoint - Upsert an ApiEndpoint.
// Upsert an ApiEndpoint. If the ApiEndpoint does not exist it will be created, otherwise it will be updated.
func (s *apiEndpoints) UpsertAPIEndpoint(ctx context.Context, request operations.UpsertAPIEndpointRequest) (*operations.UpsertAPIEndpointResponse, error) {
	baseURL := s.serverURL
	url := utils.GenerateURL(ctx, baseURL, "/v1/apis/{apiID}/version/{versionID}/api_endpoints/{apiEndpointID}", request, nil)

	bodyReader, reqContentType, err := utils.SerializeRequestBody(ctx, request, "APIEndpointInput", "json")
	if err != nil {
		return nil, fmt.Errorf("error serializing request body: %w", err)
	}
	if bodyReader == nil {
		return nil, fmt.Errorf("request body is required")
	}

	req, err := http.NewRequestWithContext(ctx, "PUT", url, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("Content-Type", reqContentType)

	client := s.securityClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}
	defer httpRes.Body.Close()

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.UpsertAPIEndpointResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.APIEndpoint
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.APIEndpoint = out
		}
	default:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.Error
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.Error = out
		}
	}

	return res, nil
}
