# \WarsAPI

All URIs are relative to *https://esi.evetech.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetWars**](WarsAPI.md#GetWars) | **Get** /wars | List wars
[**GetWarsWarId**](WarsAPI.md#GetWarsWarId) | **Get** /wars/{war_id} | Get war information
[**GetWarsWarIdKillmails**](WarsAPI.md#GetWarsWarIdKillmails) | **Get** /wars/{war_id}/killmails | List kills for a war



## GetWars

> []int64 GetWars(ctx).XCompatibilityDate(xCompatibilityDate).MaxWarId(maxWarId).AcceptLanguage(acceptLanguage).IfNoneMatch(ifNoneMatch).XTenant(xTenant).Execute()

List wars



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
	maxWarId := int64(789) // int64 |  (optional)
	acceptLanguage := "acceptLanguage_example" // string | The language to use for the response. (optional) (default to "en")
	ifNoneMatch := "ifNoneMatch_example" // string | The ETag of the previous request. A 304 will be returned if this matches the current ETag. (optional)
	xTenant := "xTenant_example" // string | The tenant ID for the request. (optional) (default to "tranquility")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WarsAPI.GetWars(context.Background()).XCompatibilityDate(xCompatibilityDate).MaxWarId(maxWarId).AcceptLanguage(acceptLanguage).IfNoneMatch(ifNoneMatch).XTenant(xTenant).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WarsAPI.GetWars``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWars`: []int64
	fmt.Fprintf(os.Stdout, "Response from `WarsAPI.GetWars`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetWarsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xCompatibilityDate** | **string** | The compatibility date for the request. | 
 **maxWarId** | **int64** |  | 
 **acceptLanguage** | **string** | The language to use for the response. | [default to &quot;en&quot;]
 **ifNoneMatch** | **string** | The ETag of the previous request. A 304 will be returned if this matches the current ETag. | 
 **xTenant** | **string** | The tenant ID for the request. | [default to &quot;tranquility&quot;]

### Return type

**[]int64**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWarsWarId

> WarsWarIdGet GetWarsWarId(ctx, warId).XCompatibilityDate(xCompatibilityDate).AcceptLanguage(acceptLanguage).IfNoneMatch(ifNoneMatch).XTenant(xTenant).Execute()

Get war information



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
	warId := int64(789) // int64 | 
	xCompatibilityDate := time.Now() // string | The compatibility date for the request.
	acceptLanguage := "acceptLanguage_example" // string | The language to use for the response. (optional) (default to "en")
	ifNoneMatch := "ifNoneMatch_example" // string | The ETag of the previous request. A 304 will be returned if this matches the current ETag. (optional)
	xTenant := "xTenant_example" // string | The tenant ID for the request. (optional) (default to "tranquility")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WarsAPI.GetWarsWarId(context.Background(), warId).XCompatibilityDate(xCompatibilityDate).AcceptLanguage(acceptLanguage).IfNoneMatch(ifNoneMatch).XTenant(xTenant).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WarsAPI.GetWarsWarId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWarsWarId`: WarsWarIdGet
	fmt.Fprintf(os.Stdout, "Response from `WarsAPI.GetWarsWarId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**warId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWarsWarIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xCompatibilityDate** | **string** | The compatibility date for the request. | 
 **acceptLanguage** | **string** | The language to use for the response. | [default to &quot;en&quot;]
 **ifNoneMatch** | **string** | The ETag of the previous request. A 304 will be returned if this matches the current ETag. | 
 **xTenant** | **string** | The tenant ID for the request. | [default to &quot;tranquility&quot;]

### Return type

[**WarsWarIdGet**](WarsWarIdGet.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWarsWarIdKillmails

> []CharactersCharacterIdKillmailsRecentGetInner GetWarsWarIdKillmails(ctx, warId).XCompatibilityDate(xCompatibilityDate).Page(page).AcceptLanguage(acceptLanguage).IfNoneMatch(ifNoneMatch).XTenant(xTenant).Execute()

List kills for a war



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
	warId := int64(789) // int64 | 
	xCompatibilityDate := time.Now() // string | The compatibility date for the request.
	page := int32(56) // int32 |  (optional)
	acceptLanguage := "acceptLanguage_example" // string | The language to use for the response. (optional) (default to "en")
	ifNoneMatch := "ifNoneMatch_example" // string | The ETag of the previous request. A 304 will be returned if this matches the current ETag. (optional)
	xTenant := "xTenant_example" // string | The tenant ID for the request. (optional) (default to "tranquility")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WarsAPI.GetWarsWarIdKillmails(context.Background(), warId).XCompatibilityDate(xCompatibilityDate).Page(page).AcceptLanguage(acceptLanguage).IfNoneMatch(ifNoneMatch).XTenant(xTenant).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WarsAPI.GetWarsWarIdKillmails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWarsWarIdKillmails`: []CharactersCharacterIdKillmailsRecentGetInner
	fmt.Fprintf(os.Stdout, "Response from `WarsAPI.GetWarsWarIdKillmails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**warId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWarsWarIdKillmailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xCompatibilityDate** | **string** | The compatibility date for the request. | 
 **page** | **int32** |  | 
 **acceptLanguage** | **string** | The language to use for the response. | [default to &quot;en&quot;]
 **ifNoneMatch** | **string** | The ETag of the previous request. A 304 will be returned if this matches the current ETag. | 
 **xTenant** | **string** | The tenant ID for the request. | [default to &quot;tranquility&quot;]

### Return type

[**[]CharactersCharacterIdKillmailsRecentGetInner**](CharactersCharacterIdKillmailsRecentGetInner.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

