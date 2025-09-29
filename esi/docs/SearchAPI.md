# \SearchAPI

All URIs are relative to *https://esi.evetech.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCharactersCharacterIdSearch**](SearchAPI.md#GetCharactersCharacterIdSearch) | **Get** /characters/{character_id}/search | Search on a string



## GetCharactersCharacterIdSearch

> CharactersCharacterIdSearchGet GetCharactersCharacterIdSearch(ctx, characterId).Categories(categories).Search(search).XCompatibilityDate(xCompatibilityDate).Strict(strict).AcceptLanguage(acceptLanguage).IfNoneMatch(ifNoneMatch).XTenant(xTenant).Execute()

Search on a string



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
	categories := []string{"Categories_example"} // []string | 
	characterId := int64(789) // int64 | The ID of the character
	search := "search_example" // string | 
	xCompatibilityDate := time.Now() // string | The compatibility date for the request.
	strict := true // bool |  (optional) (default to false)
	acceptLanguage := "acceptLanguage_example" // string | The language to use for the response. (optional) (default to "en")
	ifNoneMatch := "ifNoneMatch_example" // string | The ETag of the previous request. A 304 will be returned if this matches the current ETag. (optional)
	xTenant := "xTenant_example" // string | The tenant ID for the request. (optional) (default to "tranquility")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SearchAPI.GetCharactersCharacterIdSearch(context.Background(), characterId).Categories(categories).Search(search).XCompatibilityDate(xCompatibilityDate).Strict(strict).AcceptLanguage(acceptLanguage).IfNoneMatch(ifNoneMatch).XTenant(xTenant).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchAPI.GetCharactersCharacterIdSearch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCharactersCharacterIdSearch`: CharactersCharacterIdSearchGet
	fmt.Fprintf(os.Stdout, "Response from `SearchAPI.GetCharactersCharacterIdSearch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**characterId** | **int64** | The ID of the character | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCharactersCharacterIdSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **categories** | **[]string** |  | 

 **search** | **string** |  | 
 **xCompatibilityDate** | **string** | The compatibility date for the request. | 
 **strict** | **bool** |  | [default to false]
 **acceptLanguage** | **string** | The language to use for the response. | [default to &quot;en&quot;]
 **ifNoneMatch** | **string** | The ETag of the previous request. A 304 will be returned if this matches the current ETag. | 
 **xTenant** | **string** | The tenant ID for the request. | [default to &quot;tranquility&quot;]

### Return type

[**CharactersCharacterIdSearchGet**](CharactersCharacterIdSearchGet.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

