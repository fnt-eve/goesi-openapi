# \LoyaltyAPI

All URIs are relative to *https://esi.evetech.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCharactersCharacterIdLoyaltyPoints**](LoyaltyAPI.md#GetCharactersCharacterIdLoyaltyPoints) | **Get** /characters/{character_id}/loyalty/points | Get loyalty points
[**GetLoyaltyStoresCorporationIdOffers**](LoyaltyAPI.md#GetLoyaltyStoresCorporationIdOffers) | **Get** /loyalty/stores/{corporation_id}/offers | List loyalty store offers



## GetCharactersCharacterIdLoyaltyPoints

> []CharactersCharacterIdLoyaltyPointsGetInner GetCharactersCharacterIdLoyaltyPoints(ctx, characterId).XCompatibilityDate(xCompatibilityDate).AcceptLanguage(acceptLanguage).IfNoneMatch(ifNoneMatch).XTenant(xTenant).Execute()

Get loyalty points



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LoyaltyAPI.GetCharactersCharacterIdLoyaltyPoints(context.Background(), characterId).XCompatibilityDate(xCompatibilityDate).AcceptLanguage(acceptLanguage).IfNoneMatch(ifNoneMatch).XTenant(xTenant).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LoyaltyAPI.GetCharactersCharacterIdLoyaltyPoints``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCharactersCharacterIdLoyaltyPoints`: []CharactersCharacterIdLoyaltyPointsGetInner
	fmt.Fprintf(os.Stdout, "Response from `LoyaltyAPI.GetCharactersCharacterIdLoyaltyPoints`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**characterId** | **int64** | The ID of the character | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCharactersCharacterIdLoyaltyPointsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xCompatibilityDate** | **string** | The compatibility date for the request. | 
 **acceptLanguage** | **string** | The language to use for the response. | [default to &quot;en&quot;]
 **ifNoneMatch** | **string** | The ETag of the previous request. A 304 will be returned if this matches the current ETag. | 
 **xTenant** | **string** | The tenant ID for the request. | [default to &quot;tranquility&quot;]

### Return type

[**[]CharactersCharacterIdLoyaltyPointsGetInner**](CharactersCharacterIdLoyaltyPointsGetInner.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLoyaltyStoresCorporationIdOffers

> []LoyaltyStoresCorporationIdOffersGetInner GetLoyaltyStoresCorporationIdOffers(ctx, corporationId).XCompatibilityDate(xCompatibilityDate).AcceptLanguage(acceptLanguage).IfNoneMatch(ifNoneMatch).XTenant(xTenant).Execute()

List loyalty store offers



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
	acceptLanguage := "acceptLanguage_example" // string | The language to use for the response. (optional) (default to "en")
	ifNoneMatch := "ifNoneMatch_example" // string | The ETag of the previous request. A 304 will be returned if this matches the current ETag. (optional)
	xTenant := "xTenant_example" // string | The tenant ID for the request. (optional) (default to "tranquility")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LoyaltyAPI.GetLoyaltyStoresCorporationIdOffers(context.Background(), corporationId).XCompatibilityDate(xCompatibilityDate).AcceptLanguage(acceptLanguage).IfNoneMatch(ifNoneMatch).XTenant(xTenant).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LoyaltyAPI.GetLoyaltyStoresCorporationIdOffers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLoyaltyStoresCorporationIdOffers`: []LoyaltyStoresCorporationIdOffersGetInner
	fmt.Fprintf(os.Stdout, "Response from `LoyaltyAPI.GetLoyaltyStoresCorporationIdOffers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**corporationId** | **int64** | The ID of the corporation | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLoyaltyStoresCorporationIdOffersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xCompatibilityDate** | **string** | The compatibility date for the request. | 
 **acceptLanguage** | **string** | The language to use for the response. | [default to &quot;en&quot;]
 **ifNoneMatch** | **string** | The ETag of the previous request. A 304 will be returned if this matches the current ETag. | 
 **xTenant** | **string** | The tenant ID for the request. | [default to &quot;tranquility&quot;]

### Return type

[**[]LoyaltyStoresCorporationIdOffersGetInner**](LoyaltyStoresCorporationIdOffersGetInner.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

