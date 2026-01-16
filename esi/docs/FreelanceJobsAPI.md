# \FreelanceJobsAPI

All URIs are relative to *https://esi.evetech.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCharactersFreelanceJobsListing**](FreelanceJobsAPI.md#GetCharactersFreelanceJobsListing) | **Get** /characters/{character_id}/freelance-jobs | List character freelance jobs
[**GetCharactersFreelanceJobsParticipation**](FreelanceJobsAPI.md#GetCharactersFreelanceJobsParticipation) | **Get** /characters/{character_id}/freelance-jobs/{job_id}/participation | Get character freelance job participation
[**GetCorporationsFreelanceJobsListing**](FreelanceJobsAPI.md#GetCorporationsFreelanceJobsListing) | **Get** /corporations/{corporation_id}/freelance-jobs | List corporation freelance jobs
[**GetCorporationsFreelanceJobsParticipants**](FreelanceJobsAPI.md#GetCorporationsFreelanceJobsParticipants) | **Get** /corporations/{corporation_id}/freelance-jobs/{job_id}/participants | List participants of a freelance job
[**GetFreelanceJobsDetail**](FreelanceJobsAPI.md#GetFreelanceJobsDetail) | **Get** /freelance-jobs/{job_id} | Get freelance job details
[**GetFreelanceJobsListing**](FreelanceJobsAPI.md#GetFreelanceJobsListing) | **Get** /freelance-jobs | List freelance jobs



## GetCharactersFreelanceJobsListing

> CharactersFreelanceJobsListing GetCharactersFreelanceJobsListing(ctx, characterId).XCompatibilityDate(xCompatibilityDate).AcceptLanguage(acceptLanguage).IfNoneMatch(ifNoneMatch).XTenant(xTenant).IfModifiedSince(ifModifiedSince).Execute()

List character freelance jobs



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
	characterId := int64(789) // int64 | The ID of the character
	xCompatibilityDate := time.Now() // string | The compatibility date for the request.
	acceptLanguage := "acceptLanguage_example" // string | The language to use for the response. (optional) (default to "en")
	ifNoneMatch := "ifNoneMatch_example" // string | The ETag of the previous request. A 304 will be returned if this matches the current ETag. (optional)
	xTenant := "xTenant_example" // string | The tenant ID for the request. (optional) (default to "tranquility")
	ifModifiedSince := "ifModifiedSince_example" // string | The date the resource was last modified. A 304 will be returned if the resource has not been modified since this date. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FreelanceJobsAPI.GetCharactersFreelanceJobsListing(context.Background(), characterId).XCompatibilityDate(xCompatibilityDate).AcceptLanguage(acceptLanguage).IfNoneMatch(ifNoneMatch).XTenant(xTenant).IfModifiedSince(ifModifiedSince).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FreelanceJobsAPI.GetCharactersFreelanceJobsListing``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCharactersFreelanceJobsListing`: CharactersFreelanceJobsListing
	fmt.Fprintf(os.Stdout, "Response from `FreelanceJobsAPI.GetCharactersFreelanceJobsListing`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**characterId** | **int64** | The ID of the character | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCharactersFreelanceJobsListingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xCompatibilityDate** | **string** | The compatibility date for the request. | 
 **acceptLanguage** | **string** | The language to use for the response. | [default to &quot;en&quot;]
 **ifNoneMatch** | **string** | The ETag of the previous request. A 304 will be returned if this matches the current ETag. | 
 **xTenant** | **string** | The tenant ID for the request. | [default to &quot;tranquility&quot;]
 **ifModifiedSince** | **string** | The date the resource was last modified. A 304 will be returned if the resource has not been modified since this date. | 

### Return type

[**CharactersFreelanceJobsListing**](CharactersFreelanceJobsListing.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCharactersFreelanceJobsParticipation

> CharactersFreelanceJobsParticipation GetCharactersFreelanceJobsParticipation(ctx, characterId, jobId).XCompatibilityDate(xCompatibilityDate).AcceptLanguage(acceptLanguage).IfNoneMatch(ifNoneMatch).XTenant(xTenant).IfModifiedSince(ifModifiedSince).Execute()

Get character freelance job participation



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
	characterId := int64(789) // int64 | The ID of the character
	jobId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of the freelance job
	xCompatibilityDate := time.Now() // string | The compatibility date for the request.
	acceptLanguage := "acceptLanguage_example" // string | The language to use for the response. (optional) (default to "en")
	ifNoneMatch := "ifNoneMatch_example" // string | The ETag of the previous request. A 304 will be returned if this matches the current ETag. (optional)
	xTenant := "xTenant_example" // string | The tenant ID for the request. (optional) (default to "tranquility")
	ifModifiedSince := "ifModifiedSince_example" // string | The date the resource was last modified. A 304 will be returned if the resource has not been modified since this date. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FreelanceJobsAPI.GetCharactersFreelanceJobsParticipation(context.Background(), characterId, jobId).XCompatibilityDate(xCompatibilityDate).AcceptLanguage(acceptLanguage).IfNoneMatch(ifNoneMatch).XTenant(xTenant).IfModifiedSince(ifModifiedSince).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FreelanceJobsAPI.GetCharactersFreelanceJobsParticipation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCharactersFreelanceJobsParticipation`: CharactersFreelanceJobsParticipation
	fmt.Fprintf(os.Stdout, "Response from `FreelanceJobsAPI.GetCharactersFreelanceJobsParticipation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**characterId** | **int64** | The ID of the character | 
**jobId** | **string** | The ID of the freelance job | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCharactersFreelanceJobsParticipationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xCompatibilityDate** | **string** | The compatibility date for the request. | 
 **acceptLanguage** | **string** | The language to use for the response. | [default to &quot;en&quot;]
 **ifNoneMatch** | **string** | The ETag of the previous request. A 304 will be returned if this matches the current ETag. | 
 **xTenant** | **string** | The tenant ID for the request. | [default to &quot;tranquility&quot;]
 **ifModifiedSince** | **string** | The date the resource was last modified. A 304 will be returned if the resource has not been modified since this date. | 

### Return type

[**CharactersFreelanceJobsParticipation**](CharactersFreelanceJobsParticipation.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCorporationsFreelanceJobsListing

> CorporationsFreelanceJobsListing GetCorporationsFreelanceJobsListing(ctx, corporationId).XCompatibilityDate(xCompatibilityDate).After(after).Before(before).Limit(limit).AcceptLanguage(acceptLanguage).IfNoneMatch(ifNoneMatch).XTenant(xTenant).IfModifiedSince(ifModifiedSince).Execute()

List corporation freelance jobs



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
	acceptLanguage := "acceptLanguage_example" // string | The language to use for the response. (optional) (default to "en")
	ifNoneMatch := "ifNoneMatch_example" // string | The ETag of the previous request. A 304 will be returned if this matches the current ETag. (optional)
	xTenant := "xTenant_example" // string | The tenant ID for the request. (optional) (default to "tranquility")
	ifModifiedSince := "ifModifiedSince_example" // string | The date the resource was last modified. A 304 will be returned if the resource has not been modified since this date. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FreelanceJobsAPI.GetCorporationsFreelanceJobsListing(context.Background(), corporationId).XCompatibilityDate(xCompatibilityDate).After(after).Before(before).Limit(limit).AcceptLanguage(acceptLanguage).IfNoneMatch(ifNoneMatch).XTenant(xTenant).IfModifiedSince(ifModifiedSince).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FreelanceJobsAPI.GetCorporationsFreelanceJobsListing``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCorporationsFreelanceJobsListing`: CorporationsFreelanceJobsListing
	fmt.Fprintf(os.Stdout, "Response from `FreelanceJobsAPI.GetCorporationsFreelanceJobsListing`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**corporationId** | **int64** | The ID of the corporation | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCorporationsFreelanceJobsListingRequest struct via the builder pattern


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

[**CorporationsFreelanceJobsListing**](CorporationsFreelanceJobsListing.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCorporationsFreelanceJobsParticipants

> CorporationsFreelanceJobsParticipants GetCorporationsFreelanceJobsParticipants(ctx, corporationId, jobId).XCompatibilityDate(xCompatibilityDate).After(after).Before(before).Limit(limit).AcceptLanguage(acceptLanguage).IfNoneMatch(ifNoneMatch).XTenant(xTenant).IfModifiedSince(ifModifiedSince).Execute()

List participants of a freelance job



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
	jobId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of the job
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
	resp, r, err := apiClient.FreelanceJobsAPI.GetCorporationsFreelanceJobsParticipants(context.Background(), corporationId, jobId).XCompatibilityDate(xCompatibilityDate).After(after).Before(before).Limit(limit).AcceptLanguage(acceptLanguage).IfNoneMatch(ifNoneMatch).XTenant(xTenant).IfModifiedSince(ifModifiedSince).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FreelanceJobsAPI.GetCorporationsFreelanceJobsParticipants``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCorporationsFreelanceJobsParticipants`: CorporationsFreelanceJobsParticipants
	fmt.Fprintf(os.Stdout, "Response from `FreelanceJobsAPI.GetCorporationsFreelanceJobsParticipants`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**corporationId** | **int64** | The ID of the corporation | 
**jobId** | **string** | The ID of the job | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCorporationsFreelanceJobsParticipantsRequest struct via the builder pattern


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

[**CorporationsFreelanceJobsParticipants**](CorporationsFreelanceJobsParticipants.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFreelanceJobsDetail

> FreelanceJobsDetail GetFreelanceJobsDetail(ctx, jobId).XCompatibilityDate(xCompatibilityDate).AcceptLanguage(acceptLanguage).IfNoneMatch(ifNoneMatch).XTenant(xTenant).IfModifiedSince(ifModifiedSince).Execute()

Get freelance job details



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
	jobId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of the freelance job
	xCompatibilityDate := time.Now() // string | The compatibility date for the request.
	acceptLanguage := "acceptLanguage_example" // string | The language to use for the response. (optional) (default to "en")
	ifNoneMatch := "ifNoneMatch_example" // string | The ETag of the previous request. A 304 will be returned if this matches the current ETag. (optional)
	xTenant := "xTenant_example" // string | The tenant ID for the request. (optional) (default to "tranquility")
	ifModifiedSince := "ifModifiedSince_example" // string | The date the resource was last modified. A 304 will be returned if the resource has not been modified since this date. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FreelanceJobsAPI.GetFreelanceJobsDetail(context.Background(), jobId).XCompatibilityDate(xCompatibilityDate).AcceptLanguage(acceptLanguage).IfNoneMatch(ifNoneMatch).XTenant(xTenant).IfModifiedSince(ifModifiedSince).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FreelanceJobsAPI.GetFreelanceJobsDetail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFreelanceJobsDetail`: FreelanceJobsDetail
	fmt.Fprintf(os.Stdout, "Response from `FreelanceJobsAPI.GetFreelanceJobsDetail`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** | The ID of the freelance job | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFreelanceJobsDetailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xCompatibilityDate** | **string** | The compatibility date for the request. | 
 **acceptLanguage** | **string** | The language to use for the response. | [default to &quot;en&quot;]
 **ifNoneMatch** | **string** | The ETag of the previous request. A 304 will be returned if this matches the current ETag. | 
 **xTenant** | **string** | The tenant ID for the request. | [default to &quot;tranquility&quot;]
 **ifModifiedSince** | **string** | The date the resource was last modified. A 304 will be returned if the resource has not been modified since this date. | 

### Return type

[**FreelanceJobsDetail**](FreelanceJobsDetail.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFreelanceJobsListing

> FreelanceJobsListing GetFreelanceJobsListing(ctx).XCompatibilityDate(xCompatibilityDate).After(after).Before(before).Limit(limit).CorporationId(corporationId).AcceptLanguage(acceptLanguage).IfNoneMatch(ifNoneMatch).XTenant(xTenant).IfModifiedSince(ifModifiedSince).Execute()

List freelance jobs



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
	xCompatibilityDate := time.Now() // string | The compatibility date for the request.
	after := "after_example" // string | Return records from after this cursor (mutual exclusive with 'before'). '0' to start from the beginning. (optional)
	before := "before_example" // string | Return records from before this cursor (mutual exclusive with 'after'). '0' to start from the end. (optional)
	limit := int64(789) // int64 | The amount of records to retrieve per request. (optional) (default to 10)
	corporationId := int64(789) // int64 | Filter on corporation ID (optional)
	acceptLanguage := "acceptLanguage_example" // string | The language to use for the response. (optional) (default to "en")
	ifNoneMatch := "ifNoneMatch_example" // string | The ETag of the previous request. A 304 will be returned if this matches the current ETag. (optional)
	xTenant := "xTenant_example" // string | The tenant ID for the request. (optional) (default to "tranquility")
	ifModifiedSince := "ifModifiedSince_example" // string | The date the resource was last modified. A 304 will be returned if the resource has not been modified since this date. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FreelanceJobsAPI.GetFreelanceJobsListing(context.Background()).XCompatibilityDate(xCompatibilityDate).After(after).Before(before).Limit(limit).CorporationId(corporationId).AcceptLanguage(acceptLanguage).IfNoneMatch(ifNoneMatch).XTenant(xTenant).IfModifiedSince(ifModifiedSince).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FreelanceJobsAPI.GetFreelanceJobsListing``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFreelanceJobsListing`: FreelanceJobsListing
	fmt.Fprintf(os.Stdout, "Response from `FreelanceJobsAPI.GetFreelanceJobsListing`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFreelanceJobsListingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xCompatibilityDate** | **string** | The compatibility date for the request. | 
 **after** | **string** | Return records from after this cursor (mutual exclusive with &#39;before&#39;). &#39;0&#39; to start from the beginning. | 
 **before** | **string** | Return records from before this cursor (mutual exclusive with &#39;after&#39;). &#39;0&#39; to start from the end. | 
 **limit** | **int64** | The amount of records to retrieve per request. | [default to 10]
 **corporationId** | **int64** | Filter on corporation ID | 
 **acceptLanguage** | **string** | The language to use for the response. | [default to &quot;en&quot;]
 **ifNoneMatch** | **string** | The ETag of the previous request. A 304 will be returned if this matches the current ETag. | 
 **xTenant** | **string** | The tenant ID for the request. | [default to &quot;tranquility&quot;]
 **ifModifiedSince** | **string** | The date the resource was last modified. A 304 will be returned if the resource has not been modified since this date. | 

### Return type

[**FreelanceJobsListing**](FreelanceJobsListing.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

