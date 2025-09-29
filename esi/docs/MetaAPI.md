# \MetaAPI

All URIs are relative to *https://esi.evetech.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMetaChangelog**](MetaAPI.md#GetMetaChangelog) | **Get** /meta/changelog | Get changelog
[**GetMetaCompatibilityDates**](MetaAPI.md#GetMetaCompatibilityDates) | **Get** /meta/compatibility-dates | Get compatibility dates



## GetMetaChangelog

> MetaChangelog GetMetaChangelog(ctx).XCompatibilityDate(xCompatibilityDate).AcceptLanguage(acceptLanguage).IfNoneMatch(ifNoneMatch).XTenant(xTenant).Execute()

Get changelog



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
	acceptLanguage := "acceptLanguage_example" // string | The language to use for the response. (optional) (default to "en")
	ifNoneMatch := "ifNoneMatch_example" // string | The ETag of the previous request. A 304 will be returned if this matches the current ETag. (optional)
	xTenant := "xTenant_example" // string | The tenant ID for the request. (optional) (default to "tranquility")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetaAPI.GetMetaChangelog(context.Background()).XCompatibilityDate(xCompatibilityDate).AcceptLanguage(acceptLanguage).IfNoneMatch(ifNoneMatch).XTenant(xTenant).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetaAPI.GetMetaChangelog``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMetaChangelog`: MetaChangelog
	fmt.Fprintf(os.Stdout, "Response from `MetaAPI.GetMetaChangelog`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMetaChangelogRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xCompatibilityDate** | **string** | The compatibility date for the request. | 
 **acceptLanguage** | **string** | The language to use for the response. | [default to &quot;en&quot;]
 **ifNoneMatch** | **string** | The ETag of the previous request. A 304 will be returned if this matches the current ETag. | 
 **xTenant** | **string** | The tenant ID for the request. | [default to &quot;tranquility&quot;]

### Return type

[**MetaChangelog**](MetaChangelog.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMetaCompatibilityDates

> MetaCompatibilityDates GetMetaCompatibilityDates(ctx).XCompatibilityDate(xCompatibilityDate).AcceptLanguage(acceptLanguage).IfNoneMatch(ifNoneMatch).XTenant(xTenant).Execute()

Get compatibility dates



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
	acceptLanguage := "acceptLanguage_example" // string | The language to use for the response. (optional) (default to "en")
	ifNoneMatch := "ifNoneMatch_example" // string | The ETag of the previous request. A 304 will be returned if this matches the current ETag. (optional)
	xTenant := "xTenant_example" // string | The tenant ID for the request. (optional) (default to "tranquility")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetaAPI.GetMetaCompatibilityDates(context.Background()).XCompatibilityDate(xCompatibilityDate).AcceptLanguage(acceptLanguage).IfNoneMatch(ifNoneMatch).XTenant(xTenant).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetaAPI.GetMetaCompatibilityDates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMetaCompatibilityDates`: MetaCompatibilityDates
	fmt.Fprintf(os.Stdout, "Response from `MetaAPI.GetMetaCompatibilityDates`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMetaCompatibilityDatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xCompatibilityDate** | **string** | The compatibility date for the request. | 
 **acceptLanguage** | **string** | The language to use for the response. | [default to &quot;en&quot;]
 **ifNoneMatch** | **string** | The ETag of the previous request. A 304 will be returned if this matches the current ETag. | 
 **xTenant** | **string** | The tenant ID for the request. | [default to &quot;tranquility&quot;]

### Return type

[**MetaCompatibilityDates**](MetaCompatibilityDates.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

