# \MilitaryCampaignsAPI

All URIs are relative to *https://esi.evetech.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCharactersMilitaryCampaignsObjectivesListing**](MilitaryCampaignsAPI.md#GetCharactersMilitaryCampaignsObjectivesListing) | **Get** /characters/{character_id}/military-campaigns/objectives | List character participation in military campaigns
[**GetCharactersMilitaryCampaignsObjectivesParticipation**](MilitaryCampaignsAPI.md#GetCharactersMilitaryCampaignsObjectivesParticipation) | **Get** /characters/{character_id}/military-campaigns/objectives/{objective_id} | Get character military campaign objective participation
[**GetMilitaryCampaignsDetail**](MilitaryCampaignsAPI.md#GetMilitaryCampaignsDetail) | **Get** /military-campaigns/{campaign_id} | Get military campaign details
[**GetMilitaryCampaignsListing**](MilitaryCampaignsAPI.md#GetMilitaryCampaignsListing) | **Get** /military-campaigns | List military campaigns
[**GetMilitaryCampaignsObjectivesDetail**](MilitaryCampaignsAPI.md#GetMilitaryCampaignsObjectivesDetail) | **Get** /military-campaigns/{campaign_id}/objectives/{objective_id} | Get military campaign objective details
[**GetMilitaryCampaignsObjectivesListing**](MilitaryCampaignsAPI.md#GetMilitaryCampaignsObjectivesListing) | **Get** /military-campaigns/{campaign_id}/objectives | List military campaign objectives



## GetCharactersMilitaryCampaignsObjectivesListing

> CharactersMilitaryCampaignsObjectivesListing GetCharactersMilitaryCampaignsObjectivesListing(ctx, characterId).XCompatibilityDate(xCompatibilityDate).After(after).Before(before).Limit(limit).AcceptLanguage(acceptLanguage).IfNoneMatch(ifNoneMatch).XTenant(xTenant).IfModifiedSince(ifModifiedSince).Execute()

List character participation in military campaigns



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
	after := "after_example" // string | Return records from after this cursor (mutual exclusive with 'before'). '0' to start from the beginning. (optional)
	before := "before_example" // string | Return records from before this cursor (mutual exclusive with 'after'). '0' to start from the end. (optional)
	limit := int64(789) // int64 | The amount of records to retrieve per request. (optional) (default to 10)
	acceptLanguage := "acceptLanguage_example" // string | The language to use for the response. (optional) (default to "en")
	ifNoneMatch := "ifNoneMatch_example" // string | The ETag of the previous request. A 304 will be returned if this matches the current ETag. (optional)
	xTenant := "xTenant_example" // string | The tenant ID for the request. (optional) (default to "tranquility")
	ifModifiedSince := "ifModifiedSince_example" // string | The date the resource was last modified. A 304 will be returned if the resource has not been modified since this date. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MilitaryCampaignsAPI.GetCharactersMilitaryCampaignsObjectivesListing(context.Background(), characterId).XCompatibilityDate(xCompatibilityDate).After(after).Before(before).Limit(limit).AcceptLanguage(acceptLanguage).IfNoneMatch(ifNoneMatch).XTenant(xTenant).IfModifiedSince(ifModifiedSince).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MilitaryCampaignsAPI.GetCharactersMilitaryCampaignsObjectivesListing``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCharactersMilitaryCampaignsObjectivesListing`: CharactersMilitaryCampaignsObjectivesListing
	fmt.Fprintf(os.Stdout, "Response from `MilitaryCampaignsAPI.GetCharactersMilitaryCampaignsObjectivesListing`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**characterId** | **int64** | The ID of the character | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCharactersMilitaryCampaignsObjectivesListingRequest struct via the builder pattern


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

[**CharactersMilitaryCampaignsObjectivesListing**](CharactersMilitaryCampaignsObjectivesListing.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCharactersMilitaryCampaignsObjectivesParticipation

> CharactersMilitaryCampaignsObjectivesParticipation GetCharactersMilitaryCampaignsObjectivesParticipation(ctx, characterId, objectiveId).XCompatibilityDate(xCompatibilityDate).AcceptLanguage(acceptLanguage).IfNoneMatch(ifNoneMatch).XTenant(xTenant).IfModifiedSince(ifModifiedSince).Execute()

Get character military campaign objective participation



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
	objectiveId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of the objective
	xCompatibilityDate := time.Now() // string | The compatibility date for the request.
	acceptLanguage := "acceptLanguage_example" // string | The language to use for the response. (optional) (default to "en")
	ifNoneMatch := "ifNoneMatch_example" // string | The ETag of the previous request. A 304 will be returned if this matches the current ETag. (optional)
	xTenant := "xTenant_example" // string | The tenant ID for the request. (optional) (default to "tranquility")
	ifModifiedSince := "ifModifiedSince_example" // string | The date the resource was last modified. A 304 will be returned if the resource has not been modified since this date. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MilitaryCampaignsAPI.GetCharactersMilitaryCampaignsObjectivesParticipation(context.Background(), characterId, objectiveId).XCompatibilityDate(xCompatibilityDate).AcceptLanguage(acceptLanguage).IfNoneMatch(ifNoneMatch).XTenant(xTenant).IfModifiedSince(ifModifiedSince).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MilitaryCampaignsAPI.GetCharactersMilitaryCampaignsObjectivesParticipation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCharactersMilitaryCampaignsObjectivesParticipation`: CharactersMilitaryCampaignsObjectivesParticipation
	fmt.Fprintf(os.Stdout, "Response from `MilitaryCampaignsAPI.GetCharactersMilitaryCampaignsObjectivesParticipation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**characterId** | **int64** | The ID of the character | 
**objectiveId** | **string** | The ID of the objective | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCharactersMilitaryCampaignsObjectivesParticipationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xCompatibilityDate** | **string** | The compatibility date for the request. | 
 **acceptLanguage** | **string** | The language to use for the response. | [default to &quot;en&quot;]
 **ifNoneMatch** | **string** | The ETag of the previous request. A 304 will be returned if this matches the current ETag. | 
 **xTenant** | **string** | The tenant ID for the request. | [default to &quot;tranquility&quot;]
 **ifModifiedSince** | **string** | The date the resource was last modified. A 304 will be returned if the resource has not been modified since this date. | 

### Return type

[**CharactersMilitaryCampaignsObjectivesParticipation**](CharactersMilitaryCampaignsObjectivesParticipation.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMilitaryCampaignsDetail

> MilitaryCampaignsDetail GetMilitaryCampaignsDetail(ctx, campaignId).XCompatibilityDate(xCompatibilityDate).AcceptLanguage(acceptLanguage).IfNoneMatch(ifNoneMatch).XTenant(xTenant).IfModifiedSince(ifModifiedSince).Execute()

Get military campaign details



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
	campaignId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of the military campaign
	xCompatibilityDate := time.Now() // string | The compatibility date for the request.
	acceptLanguage := "acceptLanguage_example" // string | The language to use for the response. (optional) (default to "en")
	ifNoneMatch := "ifNoneMatch_example" // string | The ETag of the previous request. A 304 will be returned if this matches the current ETag. (optional)
	xTenant := "xTenant_example" // string | The tenant ID for the request. (optional) (default to "tranquility")
	ifModifiedSince := "ifModifiedSince_example" // string | The date the resource was last modified. A 304 will be returned if the resource has not been modified since this date. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MilitaryCampaignsAPI.GetMilitaryCampaignsDetail(context.Background(), campaignId).XCompatibilityDate(xCompatibilityDate).AcceptLanguage(acceptLanguage).IfNoneMatch(ifNoneMatch).XTenant(xTenant).IfModifiedSince(ifModifiedSince).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MilitaryCampaignsAPI.GetMilitaryCampaignsDetail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMilitaryCampaignsDetail`: MilitaryCampaignsDetail
	fmt.Fprintf(os.Stdout, "Response from `MilitaryCampaignsAPI.GetMilitaryCampaignsDetail`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **string** | The ID of the military campaign | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMilitaryCampaignsDetailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xCompatibilityDate** | **string** | The compatibility date for the request. | 
 **acceptLanguage** | **string** | The language to use for the response. | [default to &quot;en&quot;]
 **ifNoneMatch** | **string** | The ETag of the previous request. A 304 will be returned if this matches the current ETag. | 
 **xTenant** | **string** | The tenant ID for the request. | [default to &quot;tranquility&quot;]
 **ifModifiedSince** | **string** | The date the resource was last modified. A 304 will be returned if the resource has not been modified since this date. | 

### Return type

[**MilitaryCampaignsDetail**](MilitaryCampaignsDetail.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMilitaryCampaignsListing

> MilitaryCampaignsListing GetMilitaryCampaignsListing(ctx).XCompatibilityDate(xCompatibilityDate).AcceptLanguage(acceptLanguage).IfNoneMatch(ifNoneMatch).XTenant(xTenant).IfModifiedSince(ifModifiedSince).Execute()

List military campaigns



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
	ifModifiedSince := "ifModifiedSince_example" // string | The date the resource was last modified. A 304 will be returned if the resource has not been modified since this date. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MilitaryCampaignsAPI.GetMilitaryCampaignsListing(context.Background()).XCompatibilityDate(xCompatibilityDate).AcceptLanguage(acceptLanguage).IfNoneMatch(ifNoneMatch).XTenant(xTenant).IfModifiedSince(ifModifiedSince).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MilitaryCampaignsAPI.GetMilitaryCampaignsListing``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMilitaryCampaignsListing`: MilitaryCampaignsListing
	fmt.Fprintf(os.Stdout, "Response from `MilitaryCampaignsAPI.GetMilitaryCampaignsListing`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMilitaryCampaignsListingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xCompatibilityDate** | **string** | The compatibility date for the request. | 
 **acceptLanguage** | **string** | The language to use for the response. | [default to &quot;en&quot;]
 **ifNoneMatch** | **string** | The ETag of the previous request. A 304 will be returned if this matches the current ETag. | 
 **xTenant** | **string** | The tenant ID for the request. | [default to &quot;tranquility&quot;]
 **ifModifiedSince** | **string** | The date the resource was last modified. A 304 will be returned if the resource has not been modified since this date. | 

### Return type

[**MilitaryCampaignsListing**](MilitaryCampaignsListing.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMilitaryCampaignsObjectivesDetail

> MilitaryCampaignsObjectivesDetail GetMilitaryCampaignsObjectivesDetail(ctx, campaignId, objectiveId).XCompatibilityDate(xCompatibilityDate).AcceptLanguage(acceptLanguage).IfNoneMatch(ifNoneMatch).XTenant(xTenant).IfModifiedSince(ifModifiedSince).Execute()

Get military campaign objective details



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
	campaignId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of the military campaign
	objectiveId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of the objective
	xCompatibilityDate := time.Now() // string | The compatibility date for the request.
	acceptLanguage := "acceptLanguage_example" // string | The language to use for the response. (optional) (default to "en")
	ifNoneMatch := "ifNoneMatch_example" // string | The ETag of the previous request. A 304 will be returned if this matches the current ETag. (optional)
	xTenant := "xTenant_example" // string | The tenant ID for the request. (optional) (default to "tranquility")
	ifModifiedSince := "ifModifiedSince_example" // string | The date the resource was last modified. A 304 will be returned if the resource has not been modified since this date. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MilitaryCampaignsAPI.GetMilitaryCampaignsObjectivesDetail(context.Background(), campaignId, objectiveId).XCompatibilityDate(xCompatibilityDate).AcceptLanguage(acceptLanguage).IfNoneMatch(ifNoneMatch).XTenant(xTenant).IfModifiedSince(ifModifiedSince).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MilitaryCampaignsAPI.GetMilitaryCampaignsObjectivesDetail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMilitaryCampaignsObjectivesDetail`: MilitaryCampaignsObjectivesDetail
	fmt.Fprintf(os.Stdout, "Response from `MilitaryCampaignsAPI.GetMilitaryCampaignsObjectivesDetail`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **string** | The ID of the military campaign | 
**objectiveId** | **string** | The ID of the objective | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMilitaryCampaignsObjectivesDetailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xCompatibilityDate** | **string** | The compatibility date for the request. | 
 **acceptLanguage** | **string** | The language to use for the response. | [default to &quot;en&quot;]
 **ifNoneMatch** | **string** | The ETag of the previous request. A 304 will be returned if this matches the current ETag. | 
 **xTenant** | **string** | The tenant ID for the request. | [default to &quot;tranquility&quot;]
 **ifModifiedSince** | **string** | The date the resource was last modified. A 304 will be returned if the resource has not been modified since this date. | 

### Return type

[**MilitaryCampaignsObjectivesDetail**](MilitaryCampaignsObjectivesDetail.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMilitaryCampaignsObjectivesListing

> MilitaryCampaignsObjectivesListing GetMilitaryCampaignsObjectivesListing(ctx, campaignId).XCompatibilityDate(xCompatibilityDate).After(after).Before(before).Limit(limit).AcceptLanguage(acceptLanguage).IfNoneMatch(ifNoneMatch).XTenant(xTenant).IfModifiedSince(ifModifiedSince).Execute()

List military campaign objectives



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
	campaignId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of the military campaign
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
	resp, r, err := apiClient.MilitaryCampaignsAPI.GetMilitaryCampaignsObjectivesListing(context.Background(), campaignId).XCompatibilityDate(xCompatibilityDate).After(after).Before(before).Limit(limit).AcceptLanguage(acceptLanguage).IfNoneMatch(ifNoneMatch).XTenant(xTenant).IfModifiedSince(ifModifiedSince).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MilitaryCampaignsAPI.GetMilitaryCampaignsObjectivesListing``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMilitaryCampaignsObjectivesListing`: MilitaryCampaignsObjectivesListing
	fmt.Fprintf(os.Stdout, "Response from `MilitaryCampaignsAPI.GetMilitaryCampaignsObjectivesListing`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **string** | The ID of the military campaign | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMilitaryCampaignsObjectivesListingRequest struct via the builder pattern


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

[**MilitaryCampaignsObjectivesListing**](MilitaryCampaignsObjectivesListing.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

