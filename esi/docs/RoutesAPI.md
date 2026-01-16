# \RoutesAPI

All URIs are relative to *https://esi.evetech.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostRoute**](RoutesAPI.md#PostRoute) | **Post** /route/{origin_system_id}/{destination_system_id} | Get route between two systems



## PostRoute

> Route PostRoute(ctx, originSystemId, destinationSystemId).XCompatibilityDate(xCompatibilityDate).RouteRequestBody(routeRequestBody).AcceptLanguage(acceptLanguage).IfNoneMatch(ifNoneMatch).XTenant(xTenant).IfModifiedSince(ifModifiedSince).Execute()

Get route between two systems



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
	originSystemId := int64(789) // int64 | Origin system
	destinationSystemId := int64(789) // int64 | Destination system
	xCompatibilityDate := time.Now() // string | The compatibility date for the request.
	routeRequestBody := *openapiclient.NewRouteRequestBody() // RouteRequestBody | 
	acceptLanguage := "acceptLanguage_example" // string | The language to use for the response. (optional) (default to "en")
	ifNoneMatch := "ifNoneMatch_example" // string | The ETag of the previous request. A 304 will be returned if this matches the current ETag. (optional)
	xTenant := "xTenant_example" // string | The tenant ID for the request. (optional) (default to "tranquility")
	ifModifiedSince := "ifModifiedSince_example" // string | The date the resource was last modified. A 304 will be returned if the resource has not been modified since this date. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoutesAPI.PostRoute(context.Background(), originSystemId, destinationSystemId).XCompatibilityDate(xCompatibilityDate).RouteRequestBody(routeRequestBody).AcceptLanguage(acceptLanguage).IfNoneMatch(ifNoneMatch).XTenant(xTenant).IfModifiedSince(ifModifiedSince).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoutesAPI.PostRoute``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostRoute`: Route
	fmt.Fprintf(os.Stdout, "Response from `RoutesAPI.PostRoute`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**originSystemId** | **int64** | Origin system | 
**destinationSystemId** | **int64** | Destination system | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostRouteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xCompatibilityDate** | **string** | The compatibility date for the request. | 
 **routeRequestBody** | [**RouteRequestBody**](RouteRequestBody.md) |  | 
 **acceptLanguage** | **string** | The language to use for the response. | [default to &quot;en&quot;]
 **ifNoneMatch** | **string** | The ETag of the previous request. A 304 will be returned if this matches the current ETag. | 
 **xTenant** | **string** | The tenant ID for the request. | [default to &quot;tranquility&quot;]
 **ifModifiedSince** | **string** | The date the resource was last modified. A 304 will be returned if the resource has not been modified since this date. | 

### Return type

[**Route**](Route.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

