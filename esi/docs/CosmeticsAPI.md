# \CosmeticsAPI

All URIs are relative to *https://esi.evetech.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCharactersCosmeticsSkinr**](CosmeticsAPI.md#GetCharactersCosmeticsSkinr) | **Get** /characters/{character_id}/cosmetics/skinr | List a character&#39;s owned SKINR licenses
[**GetCharactersCosmeticsSkinrComponents**](CosmeticsAPI.md#GetCharactersCosmeticsSkinrComponents) | **Get** /characters/{character_id}/cosmetics/skinr/components | List a character&#39;s owned SKINR component licenses
[**GetCosmeticsSkinr**](CosmeticsAPI.md#GetCosmeticsSkinr) | **Get** /cosmetics/skinr/{skinr_id} | Get SKINR attributes



## GetCharactersCosmeticsSkinr

> CharactersCosmeticsSkinr GetCharactersCosmeticsSkinr(ctx, characterId).XCompatibilityDate(xCompatibilityDate).AcceptLanguage(acceptLanguage).IfNoneMatch(ifNoneMatch).XTenant(xTenant).IfModifiedSince(ifModifiedSince).Execute()

List a character's owned SKINR licenses



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
	characterId := int64(789) // int64 | The ID of the character whose SKINR licenses to return
	xCompatibilityDate := time.Now() // string | The compatibility date for the request.
	acceptLanguage := "acceptLanguage_example" // string | The language to use for the response. (optional) (default to "en")
	ifNoneMatch := "ifNoneMatch_example" // string | The ETag of the previous request. A 304 will be returned if this matches the current ETag. (optional)
	xTenant := "xTenant_example" // string | The tenant ID for the request. (optional) (default to "tranquility")
	ifModifiedSince := "ifModifiedSince_example" // string | The date the resource was last modified. A 304 will be returned if the resource has not been modified since this date. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CosmeticsAPI.GetCharactersCosmeticsSkinr(context.Background(), characterId).XCompatibilityDate(xCompatibilityDate).AcceptLanguage(acceptLanguage).IfNoneMatch(ifNoneMatch).XTenant(xTenant).IfModifiedSince(ifModifiedSince).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CosmeticsAPI.GetCharactersCosmeticsSkinr``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCharactersCosmeticsSkinr`: CharactersCosmeticsSkinr
	fmt.Fprintf(os.Stdout, "Response from `CosmeticsAPI.GetCharactersCosmeticsSkinr`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**characterId** | **int64** | The ID of the character whose SKINR licenses to return | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCharactersCosmeticsSkinrRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xCompatibilityDate** | **string** | The compatibility date for the request. | 
 **acceptLanguage** | **string** | The language to use for the response. | [default to &quot;en&quot;]
 **ifNoneMatch** | **string** | The ETag of the previous request. A 304 will be returned if this matches the current ETag. | 
 **xTenant** | **string** | The tenant ID for the request. | [default to &quot;tranquility&quot;]
 **ifModifiedSince** | **string** | The date the resource was last modified. A 304 will be returned if the resource has not been modified since this date. | 

### Return type

[**CharactersCosmeticsSkinr**](CharactersCosmeticsSkinr.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCharactersCosmeticsSkinrComponents

> CharactersCosmeticsSkinrComponents GetCharactersCosmeticsSkinrComponents(ctx, characterId).XCompatibilityDate(xCompatibilityDate).AcceptLanguage(acceptLanguage).IfNoneMatch(ifNoneMatch).XTenant(xTenant).IfModifiedSince(ifModifiedSince).Execute()

List a character's owned SKINR component licenses



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
	characterId := int64(789) // int64 | The ID of the character whose SKINR component licenses to return
	xCompatibilityDate := time.Now() // string | The compatibility date for the request.
	acceptLanguage := "acceptLanguage_example" // string | The language to use for the response. (optional) (default to "en")
	ifNoneMatch := "ifNoneMatch_example" // string | The ETag of the previous request. A 304 will be returned if this matches the current ETag. (optional)
	xTenant := "xTenant_example" // string | The tenant ID for the request. (optional) (default to "tranquility")
	ifModifiedSince := "ifModifiedSince_example" // string | The date the resource was last modified. A 304 will be returned if the resource has not been modified since this date. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CosmeticsAPI.GetCharactersCosmeticsSkinrComponents(context.Background(), characterId).XCompatibilityDate(xCompatibilityDate).AcceptLanguage(acceptLanguage).IfNoneMatch(ifNoneMatch).XTenant(xTenant).IfModifiedSince(ifModifiedSince).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CosmeticsAPI.GetCharactersCosmeticsSkinrComponents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCharactersCosmeticsSkinrComponents`: CharactersCosmeticsSkinrComponents
	fmt.Fprintf(os.Stdout, "Response from `CosmeticsAPI.GetCharactersCosmeticsSkinrComponents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**characterId** | **int64** | The ID of the character whose SKINR component licenses to return | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCharactersCosmeticsSkinrComponentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xCompatibilityDate** | **string** | The compatibility date for the request. | 
 **acceptLanguage** | **string** | The language to use for the response. | [default to &quot;en&quot;]
 **ifNoneMatch** | **string** | The ETag of the previous request. A 304 will be returned if this matches the current ETag. | 
 **xTenant** | **string** | The tenant ID for the request. | [default to &quot;tranquility&quot;]
 **ifModifiedSince** | **string** | The date the resource was last modified. A 304 will be returned if the resource has not been modified since this date. | 

### Return type

[**CharactersCosmeticsSkinrComponents**](CharactersCosmeticsSkinrComponents.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCosmeticsSkinr

> CosmeticsSkinr GetCosmeticsSkinr(ctx, skinrId).XCompatibilityDate(xCompatibilityDate).AcceptLanguage(acceptLanguage).IfNoneMatch(ifNoneMatch).XTenant(xTenant).IfModifiedSince(ifModifiedSince).Execute()

Get SKINR attributes



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
	skinrId := "skinrId_example" // string | SKINR identifier
	xCompatibilityDate := time.Now() // string | The compatibility date for the request.
	acceptLanguage := "acceptLanguage_example" // string | The language to use for the response. (optional) (default to "en")
	ifNoneMatch := "ifNoneMatch_example" // string | The ETag of the previous request. A 304 will be returned if this matches the current ETag. (optional)
	xTenant := "xTenant_example" // string | The tenant ID for the request. (optional) (default to "tranquility")
	ifModifiedSince := "ifModifiedSince_example" // string | The date the resource was last modified. A 304 will be returned if the resource has not been modified since this date. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CosmeticsAPI.GetCosmeticsSkinr(context.Background(), skinrId).XCompatibilityDate(xCompatibilityDate).AcceptLanguage(acceptLanguage).IfNoneMatch(ifNoneMatch).XTenant(xTenant).IfModifiedSince(ifModifiedSince).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CosmeticsAPI.GetCosmeticsSkinr``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCosmeticsSkinr`: CosmeticsSkinr
	fmt.Fprintf(os.Stdout, "Response from `CosmeticsAPI.GetCosmeticsSkinr`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**skinrId** | **string** | SKINR identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCosmeticsSkinrRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xCompatibilityDate** | **string** | The compatibility date for the request. | 
 **acceptLanguage** | **string** | The language to use for the response. | [default to &quot;en&quot;]
 **ifNoneMatch** | **string** | The ETag of the previous request. A 304 will be returned if this matches the current ETag. | 
 **xTenant** | **string** | The tenant ID for the request. | [default to &quot;tranquility&quot;]
 **ifModifiedSince** | **string** | The date the resource was last modified. A 304 will be returned if the resource has not been modified since this date. | 

### Return type

[**CosmeticsSkinr**](CosmeticsSkinr.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

