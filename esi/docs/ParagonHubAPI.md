# \ParagonHubAPI

All URIs are relative to *https://esi.evetech.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCharactersParagonHubSkinr**](ParagonHubAPI.md#GetCharactersParagonHubSkinr) | **Get** /characters/{character_id}/paragon-hub/skinr | List a character&#39;s Paragon Hub SKINR listings
[**GetParagonHubSkinr**](ParagonHubAPI.md#GetParagonHubSkinr) | **Get** /paragon-hub/skinr | List public Paragon Hub SKINR listings
[**GetParagonHubSkinrAlliances**](ParagonHubAPI.md#GetParagonHubSkinrAlliances) | **Get** /paragon-hub/skinr/alliances/{alliance_id} | List Paragon Hub SKINR listings targeted at an alliance
[**GetParagonHubSkinrCharacters**](ParagonHubAPI.md#GetParagonHubSkinrCharacters) | **Get** /paragon-hub/skinr/characters/{character_id} | List Paragon Hub SKINR listings targeted at a character
[**GetParagonHubSkinrCorporations**](ParagonHubAPI.md#GetParagonHubSkinrCorporations) | **Get** /paragon-hub/skinr/corporations/{corporation_id} | List Paragon Hub SKINR listings targeted at a corporation



## GetCharactersParagonHubSkinr

> CharactersParagonHubSkinr GetCharactersParagonHubSkinr(ctx, characterId).XCompatibilityDate(xCompatibilityDate).After(after).Before(before).Limit(limit).AcceptLanguage(acceptLanguage).IfNoneMatch(ifNoneMatch).XTenant(xTenant).IfModifiedSince(ifModifiedSince).Execute()

List a character's Paragon Hub SKINR listings



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
	characterId := int64(789) // int64 | The ID of the character whose listings to return
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
	resp, r, err := apiClient.ParagonHubAPI.GetCharactersParagonHubSkinr(context.Background(), characterId).XCompatibilityDate(xCompatibilityDate).After(after).Before(before).Limit(limit).AcceptLanguage(acceptLanguage).IfNoneMatch(ifNoneMatch).XTenant(xTenant).IfModifiedSince(ifModifiedSince).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ParagonHubAPI.GetCharactersParagonHubSkinr``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCharactersParagonHubSkinr`: CharactersParagonHubSkinr
	fmt.Fprintf(os.Stdout, "Response from `ParagonHubAPI.GetCharactersParagonHubSkinr`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**characterId** | **int64** | The ID of the character whose listings to return | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCharactersParagonHubSkinrRequest struct via the builder pattern


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

[**CharactersParagonHubSkinr**](CharactersParagonHubSkinr.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetParagonHubSkinr

> ParagonHubSkinr GetParagonHubSkinr(ctx).XCompatibilityDate(xCompatibilityDate).After(after).Before(before).Limit(limit).AcceptLanguage(acceptLanguage).IfNoneMatch(ifNoneMatch).XTenant(xTenant).IfModifiedSince(ifModifiedSince).Execute()

List public Paragon Hub SKINR listings



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
	acceptLanguage := "acceptLanguage_example" // string | The language to use for the response. (optional) (default to "en")
	ifNoneMatch := "ifNoneMatch_example" // string | The ETag of the previous request. A 304 will be returned if this matches the current ETag. (optional)
	xTenant := "xTenant_example" // string | The tenant ID for the request. (optional) (default to "tranquility")
	ifModifiedSince := "ifModifiedSince_example" // string | The date the resource was last modified. A 304 will be returned if the resource has not been modified since this date. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ParagonHubAPI.GetParagonHubSkinr(context.Background()).XCompatibilityDate(xCompatibilityDate).After(after).Before(before).Limit(limit).AcceptLanguage(acceptLanguage).IfNoneMatch(ifNoneMatch).XTenant(xTenant).IfModifiedSince(ifModifiedSince).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ParagonHubAPI.GetParagonHubSkinr``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetParagonHubSkinr`: ParagonHubSkinr
	fmt.Fprintf(os.Stdout, "Response from `ParagonHubAPI.GetParagonHubSkinr`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetParagonHubSkinrRequest struct via the builder pattern


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

[**ParagonHubSkinr**](ParagonHubSkinr.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetParagonHubSkinrAlliances

> ParagonHubSkinrAlliances GetParagonHubSkinrAlliances(ctx, allianceId).XCompatibilityDate(xCompatibilityDate).After(after).Before(before).Limit(limit).AcceptLanguage(acceptLanguage).IfNoneMatch(ifNoneMatch).XTenant(xTenant).IfModifiedSince(ifModifiedSince).Execute()

List Paragon Hub SKINR listings targeted at an alliance



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
	allianceId := int64(789) // int64 | The ID of the alliance the listings are targeted at
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
	resp, r, err := apiClient.ParagonHubAPI.GetParagonHubSkinrAlliances(context.Background(), allianceId).XCompatibilityDate(xCompatibilityDate).After(after).Before(before).Limit(limit).AcceptLanguage(acceptLanguage).IfNoneMatch(ifNoneMatch).XTenant(xTenant).IfModifiedSince(ifModifiedSince).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ParagonHubAPI.GetParagonHubSkinrAlliances``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetParagonHubSkinrAlliances`: ParagonHubSkinrAlliances
	fmt.Fprintf(os.Stdout, "Response from `ParagonHubAPI.GetParagonHubSkinrAlliances`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**allianceId** | **int64** | The ID of the alliance the listings are targeted at | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetParagonHubSkinrAlliancesRequest struct via the builder pattern


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

[**ParagonHubSkinrAlliances**](ParagonHubSkinrAlliances.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetParagonHubSkinrCharacters

> ParagonHubSkinrCharacters GetParagonHubSkinrCharacters(ctx, characterId).XCompatibilityDate(xCompatibilityDate).After(after).Before(before).Limit(limit).AcceptLanguage(acceptLanguage).IfNoneMatch(ifNoneMatch).XTenant(xTenant).IfModifiedSince(ifModifiedSince).Execute()

List Paragon Hub SKINR listings targeted at a character



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
	characterId := int64(789) // int64 | The ID of the character the listings are targeted at
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
	resp, r, err := apiClient.ParagonHubAPI.GetParagonHubSkinrCharacters(context.Background(), characterId).XCompatibilityDate(xCompatibilityDate).After(after).Before(before).Limit(limit).AcceptLanguage(acceptLanguage).IfNoneMatch(ifNoneMatch).XTenant(xTenant).IfModifiedSince(ifModifiedSince).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ParagonHubAPI.GetParagonHubSkinrCharacters``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetParagonHubSkinrCharacters`: ParagonHubSkinrCharacters
	fmt.Fprintf(os.Stdout, "Response from `ParagonHubAPI.GetParagonHubSkinrCharacters`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**characterId** | **int64** | The ID of the character the listings are targeted at | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetParagonHubSkinrCharactersRequest struct via the builder pattern


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

[**ParagonHubSkinrCharacters**](ParagonHubSkinrCharacters.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetParagonHubSkinrCorporations

> ParagonHubSkinrCorporations GetParagonHubSkinrCorporations(ctx, corporationId).XCompatibilityDate(xCompatibilityDate).After(after).Before(before).Limit(limit).AcceptLanguage(acceptLanguage).IfNoneMatch(ifNoneMatch).XTenant(xTenant).IfModifiedSince(ifModifiedSince).Execute()

List Paragon Hub SKINR listings targeted at a corporation



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
	corporationId := int64(789) // int64 | The ID of the corporation the listings are targeted at
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
	resp, r, err := apiClient.ParagonHubAPI.GetParagonHubSkinrCorporations(context.Background(), corporationId).XCompatibilityDate(xCompatibilityDate).After(after).Before(before).Limit(limit).AcceptLanguage(acceptLanguage).IfNoneMatch(ifNoneMatch).XTenant(xTenant).IfModifiedSince(ifModifiedSince).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ParagonHubAPI.GetParagonHubSkinrCorporations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetParagonHubSkinrCorporations`: ParagonHubSkinrCorporations
	fmt.Fprintf(os.Stdout, "Response from `ParagonHubAPI.GetParagonHubSkinrCorporations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**corporationId** | **int64** | The ID of the corporation the listings are targeted at | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetParagonHubSkinrCorporationsRequest struct via the builder pattern


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

[**ParagonHubSkinrCorporations**](ParagonHubSkinrCorporations.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

