# \CorporationProjectsAPI

All URIs are relative to *https://esi.evetech.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCorporationsProjectsContribution**](CorporationProjectsAPI.md#GetCorporationsProjectsContribution) | **Get** /corporations/{corporation_id}/projects/{project_id}/contribution/{character_id} | Get your project contribution
[**GetCorporationsProjectsContributors**](CorporationProjectsAPI.md#GetCorporationsProjectsContributors) | **Get** /corporations/{corporation_id}/projects/{project_id}/contributors | List project contributors
[**GetCorporationsProjectsDetail**](CorporationProjectsAPI.md#GetCorporationsProjectsDetail) | **Get** /corporations/{corporation_id}/projects/{project_id} | Get project details
[**GetCorporationsProjectsListing**](CorporationProjectsAPI.md#GetCorporationsProjectsListing) | **Get** /corporations/{corporation_id}/projects | List corporation projects



## GetCorporationsProjectsContribution

> CorporationsProjectsContribution GetCorporationsProjectsContribution(ctx, corporationId, projectId, characterId).XCompatibilityDate(xCompatibilityDate).AcceptLanguage(acceptLanguage).IfNoneMatch(ifNoneMatch).XTenant(xTenant).IfModifiedSince(ifModifiedSince).Execute()

Get your project contribution



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	corporationId := int64(789) // int64 | The ID of the corporation
	projectId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of the project
	characterId := int64(789) // int64 | The ID of the character
	xCompatibilityDate := time.Now() // string | The compatibility date for the request.
	acceptLanguage := "acceptLanguage_example" // string | The language to use for the response. (optional) (default to "en")
	ifNoneMatch := "ifNoneMatch_example" // string | The ETag of the previous request. A 304 will be returned if this matches the current ETag. (optional)
	xTenant := "xTenant_example" // string | The tenant ID for the request. (optional) (default to "tranquility")
	ifModifiedSince := "ifModifiedSince_example" // string | The date the resource was last modified. A 304 will be returned if the resource has not been modified since this date. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CorporationProjectsAPI.GetCorporationsProjectsContribution(context.Background(), corporationId, projectId, characterId).XCompatibilityDate(xCompatibilityDate).AcceptLanguage(acceptLanguage).IfNoneMatch(ifNoneMatch).XTenant(xTenant).IfModifiedSince(ifModifiedSince).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CorporationProjectsAPI.GetCorporationsProjectsContribution``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCorporationsProjectsContribution`: CorporationsProjectsContribution
	fmt.Fprintf(os.Stdout, "Response from `CorporationProjectsAPI.GetCorporationsProjectsContribution`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**corporationId** | **int64** | The ID of the corporation | 
**projectId** | **string** | The ID of the project | 
**characterId** | **int64** | The ID of the character | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCorporationsProjectsContributionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **xCompatibilityDate** | **string** | The compatibility date for the request. | 
 **acceptLanguage** | **string** | The language to use for the response. | [default to &quot;en&quot;]
 **ifNoneMatch** | **string** | The ETag of the previous request. A 304 will be returned if this matches the current ETag. | 
 **xTenant** | **string** | The tenant ID for the request. | [default to &quot;tranquility&quot;]
 **ifModifiedSince** | **string** | The date the resource was last modified. A 304 will be returned if the resource has not been modified since this date. | 

### Return type

[**CorporationsProjectsContribution**](CorporationsProjectsContribution.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCorporationsProjectsContributors

> CorporationsProjectsContributors GetCorporationsProjectsContributors(ctx, corporationId, projectId).XCompatibilityDate(xCompatibilityDate).After(after).Before(before).Limit(limit).AcceptLanguage(acceptLanguage).IfNoneMatch(ifNoneMatch).XTenant(xTenant).IfModifiedSince(ifModifiedSince).Execute()

List project contributors



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	corporationId := int64(789) // int64 | The ID of the corporation
	projectId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of the project
	xCompatibilityDate := time.Now() // string | The compatibility date for the request.
	after := "after_example" // string | Return records from after this cursor (mutual exclusive with 'before'). '0' to start from the beginning. (optional)
	before := "before_example" // string | Return records from before this cursor (mutual exclusive with 'after'). '0' to start from the end. (optional)
	limit := int64(789) // int64 | The amount of records to retrieve per request. (optional) (default to 10)
	acceptLanguage := "acceptLanguage_example" // string | The language to use for the response. (optional) (default to "en")
	ifNoneMatch := "ifNoneMatch_example" // string | The ETag of the previous request. A 304 will be returned if this matches the current ETag. (optional)
	xTenant := "xTenant_example" // string | The tenant ID for the request. (optional) (default to "tranquility")
	ifModifiedSince := "ifModifiedSince_example" // string | The date the resource was last modified. A 304 will be returned if the resource has not been modified since this date. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CorporationProjectsAPI.GetCorporationsProjectsContributors(context.Background(), corporationId, projectId).XCompatibilityDate(xCompatibilityDate).After(after).Before(before).Limit(limit).AcceptLanguage(acceptLanguage).IfNoneMatch(ifNoneMatch).XTenant(xTenant).IfModifiedSince(ifModifiedSince).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CorporationProjectsAPI.GetCorporationsProjectsContributors``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCorporationsProjectsContributors`: CorporationsProjectsContributors
	fmt.Fprintf(os.Stdout, "Response from `CorporationProjectsAPI.GetCorporationsProjectsContributors`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**corporationId** | **int64** | The ID of the corporation | 
**projectId** | **string** | The ID of the project | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCorporationsProjectsContributorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xCompatibilityDate** | **string** | The compatibility date for the request. | 
 **after** | **string** | Return records from after this cursor (mutual exclusive with &#39;before&#39;). &#39;0&#39; to start from the beginning. | 
 **before** | **string** | Return records from before this cursor (mutual exclusive with &#39;after&#39;). &#39;0&#39; to start from the end. | 
 **limit** | **int64** | The amount of records to retrieve per request. | [default to 10]
 **acceptLanguage** | **string** | The language to use for the response. | [default to &quot;en&quot;]
 **ifNoneMatch** | **string** | The ETag of the previous request. A 304 will be returned if this matches the current ETag. | 
 **xTenant** | **string** | The tenant ID for the request. | [default to &quot;tranquility&quot;]
 **ifModifiedSince** | **string** | The date the resource was last modified. A 304 will be returned if the resource has not been modified since this date. | 

### Return type

[**CorporationsProjectsContributors**](CorporationsProjectsContributors.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCorporationsProjectsDetail

> CorporationsProjectsDetail GetCorporationsProjectsDetail(ctx, corporationId, projectId).XCompatibilityDate(xCompatibilityDate).AcceptLanguage(acceptLanguage).IfNoneMatch(ifNoneMatch).XTenant(xTenant).IfModifiedSince(ifModifiedSince).Execute()

Get project details



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	corporationId := int64(789) // int64 | The ID of the corporation
	projectId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of the project
	xCompatibilityDate := time.Now() // string | The compatibility date for the request.
	acceptLanguage := "acceptLanguage_example" // string | The language to use for the response. (optional) (default to "en")
	ifNoneMatch := "ifNoneMatch_example" // string | The ETag of the previous request. A 304 will be returned if this matches the current ETag. (optional)
	xTenant := "xTenant_example" // string | The tenant ID for the request. (optional) (default to "tranquility")
	ifModifiedSince := "ifModifiedSince_example" // string | The date the resource was last modified. A 304 will be returned if the resource has not been modified since this date. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CorporationProjectsAPI.GetCorporationsProjectsDetail(context.Background(), corporationId, projectId).XCompatibilityDate(xCompatibilityDate).AcceptLanguage(acceptLanguage).IfNoneMatch(ifNoneMatch).XTenant(xTenant).IfModifiedSince(ifModifiedSince).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CorporationProjectsAPI.GetCorporationsProjectsDetail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCorporationsProjectsDetail`: CorporationsProjectsDetail
	fmt.Fprintf(os.Stdout, "Response from `CorporationProjectsAPI.GetCorporationsProjectsDetail`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**corporationId** | **int64** | The ID of the corporation | 
**projectId** | **string** | The ID of the project | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCorporationsProjectsDetailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xCompatibilityDate** | **string** | The compatibility date for the request. | 
 **acceptLanguage** | **string** | The language to use for the response. | [default to &quot;en&quot;]
 **ifNoneMatch** | **string** | The ETag of the previous request. A 304 will be returned if this matches the current ETag. | 
 **xTenant** | **string** | The tenant ID for the request. | [default to &quot;tranquility&quot;]
 **ifModifiedSince** | **string** | The date the resource was last modified. A 304 will be returned if the resource has not been modified since this date. | 

### Return type

[**CorporationsProjectsDetail**](CorporationsProjectsDetail.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCorporationsProjectsListing

> CorporationsProjectsListing GetCorporationsProjectsListing(ctx, corporationId).XCompatibilityDate(xCompatibilityDate).After(after).Before(before).Limit(limit).State(state).AcceptLanguage(acceptLanguage).IfNoneMatch(ifNoneMatch).XTenant(xTenant).IfModifiedSince(ifModifiedSince).Execute()

List corporation projects



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	corporationId := int64(789) // int64 | The ID of the corporation
	xCompatibilityDate := time.Now() // string | The compatibility date for the request.
	after := "after_example" // string | Return records from after this cursor (mutual exclusive with 'before'). '0' to start from the beginning. (optional)
	before := "before_example" // string | Return records from before this cursor (mutual exclusive with 'after'). '0' to start from the end. (optional)
	limit := int64(789) // int64 | The amount of records to retrieve per request. (optional) (default to 10)
	state := "state_example" // string | Filter by state (optional) (default to "Active")
	acceptLanguage := "acceptLanguage_example" // string | The language to use for the response. (optional) (default to "en")
	ifNoneMatch := "ifNoneMatch_example" // string | The ETag of the previous request. A 304 will be returned if this matches the current ETag. (optional)
	xTenant := "xTenant_example" // string | The tenant ID for the request. (optional) (default to "tranquility")
	ifModifiedSince := "ifModifiedSince_example" // string | The date the resource was last modified. A 304 will be returned if the resource has not been modified since this date. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CorporationProjectsAPI.GetCorporationsProjectsListing(context.Background(), corporationId).XCompatibilityDate(xCompatibilityDate).After(after).Before(before).Limit(limit).State(state).AcceptLanguage(acceptLanguage).IfNoneMatch(ifNoneMatch).XTenant(xTenant).IfModifiedSince(ifModifiedSince).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CorporationProjectsAPI.GetCorporationsProjectsListing``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCorporationsProjectsListing`: CorporationsProjectsListing
	fmt.Fprintf(os.Stdout, "Response from `CorporationProjectsAPI.GetCorporationsProjectsListing`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**corporationId** | **int64** | The ID of the corporation | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCorporationsProjectsListingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xCompatibilityDate** | **string** | The compatibility date for the request. | 
 **after** | **string** | Return records from after this cursor (mutual exclusive with &#39;before&#39;). &#39;0&#39; to start from the beginning. | 
 **before** | **string** | Return records from before this cursor (mutual exclusive with &#39;after&#39;). &#39;0&#39; to start from the end. | 
 **limit** | **int64** | The amount of records to retrieve per request. | [default to 10]
 **state** | **string** | Filter by state | [default to &quot;Active&quot;]
 **acceptLanguage** | **string** | The language to use for the response. | [default to &quot;en&quot;]
 **ifNoneMatch** | **string** | The ETag of the previous request. A 304 will be returned if this matches the current ETag. | 
 **xTenant** | **string** | The tenant ID for the request. | [default to &quot;tranquility&quot;]
 **ifModifiedSince** | **string** | The date the resource was last modified. A 304 will be returned if the resource has not been modified since this date. | 

### Return type

[**CorporationsProjectsListing**](CorporationsProjectsListing.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

