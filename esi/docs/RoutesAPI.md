# \RoutesAPI

All URIs are relative to *https://esi.evetech.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetRouteOriginDestination**](RoutesAPI.md#GetRouteOriginDestination) | **Get** /route/{origin}/{destination} | Get route



## GetRouteOriginDestination

> []int64 GetRouteOriginDestination(ctx, destination, origin).XCompatibilityDate(xCompatibilityDate).Avoid(avoid).Connections(connections).Flag(flag).AcceptLanguage(acceptLanguage).IfNoneMatch(ifNoneMatch).XTenant(xTenant).Execute()

Get route



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
	destination := int64(789) // int64 | 
	origin := int64(789) // int64 | 
	xCompatibilityDate := time.Now() // string | The compatibility date for the request.
	avoid := []int64{int64(123)} // []int64 |  (optional)
	connections := [][]int64{[]int64{int64(123)}} // [][]int64 |  (optional)
	flag := "flag_example" // string |  (optional) (default to "shortest")
	acceptLanguage := "acceptLanguage_example" // string | The language to use for the response. (optional) (default to "en")
	ifNoneMatch := "ifNoneMatch_example" // string | The ETag of the previous request. A 304 will be returned if this matches the current ETag. (optional)
	xTenant := "xTenant_example" // string | The tenant ID for the request. (optional) (default to "tranquility")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoutesAPI.GetRouteOriginDestination(context.Background(), destination, origin).XCompatibilityDate(xCompatibilityDate).Avoid(avoid).Connections(connections).Flag(flag).AcceptLanguage(acceptLanguage).IfNoneMatch(ifNoneMatch).XTenant(xTenant).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoutesAPI.GetRouteOriginDestination``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRouteOriginDestination`: []int64
	fmt.Fprintf(os.Stdout, "Response from `RoutesAPI.GetRouteOriginDestination`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**destination** | **int64** |  | 
**origin** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRouteOriginDestinationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xCompatibilityDate** | **string** | The compatibility date for the request. | 
 **avoid** | **[]int64** |  | 
 **connections** | **[][]int64** |  | 
 **flag** | **string** |  | [default to &quot;shortest&quot;]
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

