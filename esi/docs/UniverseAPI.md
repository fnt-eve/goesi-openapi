# \UniverseAPI

All URIs are relative to *https://esi.evetech.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetUniverseAncestries**](UniverseAPI.md#GetUniverseAncestries) | **Get** /universe/ancestries | Get ancestries
[**GetUniverseAsteroidBeltsAsteroidBeltId**](UniverseAPI.md#GetUniverseAsteroidBeltsAsteroidBeltId) | **Get** /universe/asteroid_belts/{asteroid_belt_id} | Get asteroid belt information
[**GetUniverseBloodlines**](UniverseAPI.md#GetUniverseBloodlines) | **Get** /universe/bloodlines | Get bloodlines
[**GetUniverseCategories**](UniverseAPI.md#GetUniverseCategories) | **Get** /universe/categories | Get item categories
[**GetUniverseCategoriesCategoryId**](UniverseAPI.md#GetUniverseCategoriesCategoryId) | **Get** /universe/categories/{category_id} | Get item category information
[**GetUniverseConstellations**](UniverseAPI.md#GetUniverseConstellations) | **Get** /universe/constellations | Get constellations
[**GetUniverseConstellationsConstellationId**](UniverseAPI.md#GetUniverseConstellationsConstellationId) | **Get** /universe/constellations/{constellation_id} | Get constellation information
[**GetUniverseFactions**](UniverseAPI.md#GetUniverseFactions) | **Get** /universe/factions | Get factions
[**GetUniverseGraphics**](UniverseAPI.md#GetUniverseGraphics) | **Get** /universe/graphics | Get graphics
[**GetUniverseGraphicsGraphicId**](UniverseAPI.md#GetUniverseGraphicsGraphicId) | **Get** /universe/graphics/{graphic_id} | Get graphic information
[**GetUniverseGroups**](UniverseAPI.md#GetUniverseGroups) | **Get** /universe/groups | Get item groups
[**GetUniverseGroupsGroupId**](UniverseAPI.md#GetUniverseGroupsGroupId) | **Get** /universe/groups/{group_id} | Get item group information
[**GetUniverseMoonsMoonId**](UniverseAPI.md#GetUniverseMoonsMoonId) | **Get** /universe/moons/{moon_id} | Get moon information
[**GetUniversePlanetsPlanetId**](UniverseAPI.md#GetUniversePlanetsPlanetId) | **Get** /universe/planets/{planet_id} | Get planet information
[**GetUniverseRaces**](UniverseAPI.md#GetUniverseRaces) | **Get** /universe/races | Get character races
[**GetUniverseRegions**](UniverseAPI.md#GetUniverseRegions) | **Get** /universe/regions | Get regions
[**GetUniverseRegionsRegionId**](UniverseAPI.md#GetUniverseRegionsRegionId) | **Get** /universe/regions/{region_id} | Get region information
[**GetUniverseStargatesStargateId**](UniverseAPI.md#GetUniverseStargatesStargateId) | **Get** /universe/stargates/{stargate_id} | Get stargate information
[**GetUniverseStarsStarId**](UniverseAPI.md#GetUniverseStarsStarId) | **Get** /universe/stars/{star_id} | Get star information
[**GetUniverseStationsStationId**](UniverseAPI.md#GetUniverseStationsStationId) | **Get** /universe/stations/{station_id} | Get station information
[**GetUniverseStructures**](UniverseAPI.md#GetUniverseStructures) | **Get** /universe/structures | List all public structures
[**GetUniverseStructuresStructureId**](UniverseAPI.md#GetUniverseStructuresStructureId) | **Get** /universe/structures/{structure_id} | Get structure information
[**GetUniverseSystemJumps**](UniverseAPI.md#GetUniverseSystemJumps) | **Get** /universe/system_jumps | Get system jumps
[**GetUniverseSystemKills**](UniverseAPI.md#GetUniverseSystemKills) | **Get** /universe/system_kills | Get system kills
[**GetUniverseSystems**](UniverseAPI.md#GetUniverseSystems) | **Get** /universe/systems | Get solar systems
[**GetUniverseSystemsSystemId**](UniverseAPI.md#GetUniverseSystemsSystemId) | **Get** /universe/systems/{system_id} | Get solar system information
[**GetUniverseTypes**](UniverseAPI.md#GetUniverseTypes) | **Get** /universe/types | Get types
[**GetUniverseTypesTypeId**](UniverseAPI.md#GetUniverseTypesTypeId) | **Get** /universe/types/{type_id} | Get type information
[**PostUniverseIds**](UniverseAPI.md#PostUniverseIds) | **Post** /universe/ids | Bulk names to IDs
[**PostUniverseNames**](UniverseAPI.md#PostUniverseNames) | **Post** /universe/names | Get names and categories for a set of IDs



## GetUniverseAncestries

> []UniverseAncestriesGetInner GetUniverseAncestries(ctx).XCompatibilityDate(xCompatibilityDate).AcceptLanguage(acceptLanguage).IfNoneMatch(ifNoneMatch).XTenant(xTenant).Execute()

Get ancestries



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
	resp, r, err := apiClient.UniverseAPI.GetUniverseAncestries(context.Background()).XCompatibilityDate(xCompatibilityDate).AcceptLanguage(acceptLanguage).IfNoneMatch(ifNoneMatch).XTenant(xTenant).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UniverseAPI.GetUniverseAncestries``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUniverseAncestries`: []UniverseAncestriesGetInner
	fmt.Fprintf(os.Stdout, "Response from `UniverseAPI.GetUniverseAncestries`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUniverseAncestriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xCompatibilityDate** | **string** | The compatibility date for the request. | 
 **acceptLanguage** | **string** | The language to use for the response. | [default to &quot;en&quot;]
 **ifNoneMatch** | **string** | The ETag of the previous request. A 304 will be returned if this matches the current ETag. | 
 **xTenant** | **string** | The tenant ID for the request. | [default to &quot;tranquility&quot;]

### Return type

[**[]UniverseAncestriesGetInner**](UniverseAncestriesGetInner.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUniverseAsteroidBeltsAsteroidBeltId

> UniverseAsteroidBeltsAsteroidBeltIdGet GetUniverseAsteroidBeltsAsteroidBeltId(ctx, asteroidBeltId).XCompatibilityDate(xCompatibilityDate).AcceptLanguage(acceptLanguage).IfNoneMatch(ifNoneMatch).XTenant(xTenant).Execute()

Get asteroid belt information



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
	asteroidBeltId := int64(789) // int64 | 
	xCompatibilityDate := time.Now() // string | The compatibility date for the request.
	acceptLanguage := "acceptLanguage_example" // string | The language to use for the response. (optional) (default to "en")
	ifNoneMatch := "ifNoneMatch_example" // string | The ETag of the previous request. A 304 will be returned if this matches the current ETag. (optional)
	xTenant := "xTenant_example" // string | The tenant ID for the request. (optional) (default to "tranquility")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UniverseAPI.GetUniverseAsteroidBeltsAsteroidBeltId(context.Background(), asteroidBeltId).XCompatibilityDate(xCompatibilityDate).AcceptLanguage(acceptLanguage).IfNoneMatch(ifNoneMatch).XTenant(xTenant).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UniverseAPI.GetUniverseAsteroidBeltsAsteroidBeltId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUniverseAsteroidBeltsAsteroidBeltId`: UniverseAsteroidBeltsAsteroidBeltIdGet
	fmt.Fprintf(os.Stdout, "Response from `UniverseAPI.GetUniverseAsteroidBeltsAsteroidBeltId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**asteroidBeltId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUniverseAsteroidBeltsAsteroidBeltIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xCompatibilityDate** | **string** | The compatibility date for the request. | 
 **acceptLanguage** | **string** | The language to use for the response. | [default to &quot;en&quot;]
 **ifNoneMatch** | **string** | The ETag of the previous request. A 304 will be returned if this matches the current ETag. | 
 **xTenant** | **string** | The tenant ID for the request. | [default to &quot;tranquility&quot;]

### Return type

[**UniverseAsteroidBeltsAsteroidBeltIdGet**](UniverseAsteroidBeltsAsteroidBeltIdGet.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUniverseBloodlines

> []UniverseBloodlinesGetInner GetUniverseBloodlines(ctx).XCompatibilityDate(xCompatibilityDate).AcceptLanguage(acceptLanguage).IfNoneMatch(ifNoneMatch).XTenant(xTenant).Execute()

Get bloodlines



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
	resp, r, err := apiClient.UniverseAPI.GetUniverseBloodlines(context.Background()).XCompatibilityDate(xCompatibilityDate).AcceptLanguage(acceptLanguage).IfNoneMatch(ifNoneMatch).XTenant(xTenant).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UniverseAPI.GetUniverseBloodlines``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUniverseBloodlines`: []UniverseBloodlinesGetInner
	fmt.Fprintf(os.Stdout, "Response from `UniverseAPI.GetUniverseBloodlines`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUniverseBloodlinesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xCompatibilityDate** | **string** | The compatibility date for the request. | 
 **acceptLanguage** | **string** | The language to use for the response. | [default to &quot;en&quot;]
 **ifNoneMatch** | **string** | The ETag of the previous request. A 304 will be returned if this matches the current ETag. | 
 **xTenant** | **string** | The tenant ID for the request. | [default to &quot;tranquility&quot;]

### Return type

[**[]UniverseBloodlinesGetInner**](UniverseBloodlinesGetInner.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUniverseCategories

> []int64 GetUniverseCategories(ctx).XCompatibilityDate(xCompatibilityDate).AcceptLanguage(acceptLanguage).IfNoneMatch(ifNoneMatch).XTenant(xTenant).Execute()

Get item categories



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
	resp, r, err := apiClient.UniverseAPI.GetUniverseCategories(context.Background()).XCompatibilityDate(xCompatibilityDate).AcceptLanguage(acceptLanguage).IfNoneMatch(ifNoneMatch).XTenant(xTenant).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UniverseAPI.GetUniverseCategories``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUniverseCategories`: []int64
	fmt.Fprintf(os.Stdout, "Response from `UniverseAPI.GetUniverseCategories`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUniverseCategoriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xCompatibilityDate** | **string** | The compatibility date for the request. | 
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


## GetUniverseCategoriesCategoryId

> UniverseCategoriesCategoryIdGet GetUniverseCategoriesCategoryId(ctx, categoryId).XCompatibilityDate(xCompatibilityDate).AcceptLanguage(acceptLanguage).IfNoneMatch(ifNoneMatch).XTenant(xTenant).Execute()

Get item category information



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
	categoryId := int64(789) // int64 | 
	xCompatibilityDate := time.Now() // string | The compatibility date for the request.
	acceptLanguage := "acceptLanguage_example" // string | The language to use for the response. (optional) (default to "en")
	ifNoneMatch := "ifNoneMatch_example" // string | The ETag of the previous request. A 304 will be returned if this matches the current ETag. (optional)
	xTenant := "xTenant_example" // string | The tenant ID for the request. (optional) (default to "tranquility")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UniverseAPI.GetUniverseCategoriesCategoryId(context.Background(), categoryId).XCompatibilityDate(xCompatibilityDate).AcceptLanguage(acceptLanguage).IfNoneMatch(ifNoneMatch).XTenant(xTenant).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UniverseAPI.GetUniverseCategoriesCategoryId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUniverseCategoriesCategoryId`: UniverseCategoriesCategoryIdGet
	fmt.Fprintf(os.Stdout, "Response from `UniverseAPI.GetUniverseCategoriesCategoryId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**categoryId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUniverseCategoriesCategoryIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xCompatibilityDate** | **string** | The compatibility date for the request. | 
 **acceptLanguage** | **string** | The language to use for the response. | [default to &quot;en&quot;]
 **ifNoneMatch** | **string** | The ETag of the previous request. A 304 will be returned if this matches the current ETag. | 
 **xTenant** | **string** | The tenant ID for the request. | [default to &quot;tranquility&quot;]

### Return type

[**UniverseCategoriesCategoryIdGet**](UniverseCategoriesCategoryIdGet.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUniverseConstellations

> []int64 GetUniverseConstellations(ctx).XCompatibilityDate(xCompatibilityDate).AcceptLanguage(acceptLanguage).IfNoneMatch(ifNoneMatch).XTenant(xTenant).Execute()

Get constellations



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
	resp, r, err := apiClient.UniverseAPI.GetUniverseConstellations(context.Background()).XCompatibilityDate(xCompatibilityDate).AcceptLanguage(acceptLanguage).IfNoneMatch(ifNoneMatch).XTenant(xTenant).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UniverseAPI.GetUniverseConstellations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUniverseConstellations`: []int64
	fmt.Fprintf(os.Stdout, "Response from `UniverseAPI.GetUniverseConstellations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUniverseConstellationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xCompatibilityDate** | **string** | The compatibility date for the request. | 
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


## GetUniverseConstellationsConstellationId

> UniverseConstellationsConstellationIdGet GetUniverseConstellationsConstellationId(ctx, constellationId).XCompatibilityDate(xCompatibilityDate).AcceptLanguage(acceptLanguage).IfNoneMatch(ifNoneMatch).XTenant(xTenant).Execute()

Get constellation information



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
	constellationId := int64(789) // int64 | 
	xCompatibilityDate := time.Now() // string | The compatibility date for the request.
	acceptLanguage := "acceptLanguage_example" // string | The language to use for the response. (optional) (default to "en")
	ifNoneMatch := "ifNoneMatch_example" // string | The ETag of the previous request. A 304 will be returned if this matches the current ETag. (optional)
	xTenant := "xTenant_example" // string | The tenant ID for the request. (optional) (default to "tranquility")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UniverseAPI.GetUniverseConstellationsConstellationId(context.Background(), constellationId).XCompatibilityDate(xCompatibilityDate).AcceptLanguage(acceptLanguage).IfNoneMatch(ifNoneMatch).XTenant(xTenant).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UniverseAPI.GetUniverseConstellationsConstellationId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUniverseConstellationsConstellationId`: UniverseConstellationsConstellationIdGet
	fmt.Fprintf(os.Stdout, "Response from `UniverseAPI.GetUniverseConstellationsConstellationId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**constellationId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUniverseConstellationsConstellationIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xCompatibilityDate** | **string** | The compatibility date for the request. | 
 **acceptLanguage** | **string** | The language to use for the response. | [default to &quot;en&quot;]
 **ifNoneMatch** | **string** | The ETag of the previous request. A 304 will be returned if this matches the current ETag. | 
 **xTenant** | **string** | The tenant ID for the request. | [default to &quot;tranquility&quot;]

### Return type

[**UniverseConstellationsConstellationIdGet**](UniverseConstellationsConstellationIdGet.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUniverseFactions

> []UniverseFactionsGetInner GetUniverseFactions(ctx).XCompatibilityDate(xCompatibilityDate).AcceptLanguage(acceptLanguage).IfNoneMatch(ifNoneMatch).XTenant(xTenant).Execute()

Get factions



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
	resp, r, err := apiClient.UniverseAPI.GetUniverseFactions(context.Background()).XCompatibilityDate(xCompatibilityDate).AcceptLanguage(acceptLanguage).IfNoneMatch(ifNoneMatch).XTenant(xTenant).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UniverseAPI.GetUniverseFactions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUniverseFactions`: []UniverseFactionsGetInner
	fmt.Fprintf(os.Stdout, "Response from `UniverseAPI.GetUniverseFactions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUniverseFactionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xCompatibilityDate** | **string** | The compatibility date for the request. | 
 **acceptLanguage** | **string** | The language to use for the response. | [default to &quot;en&quot;]
 **ifNoneMatch** | **string** | The ETag of the previous request. A 304 will be returned if this matches the current ETag. | 
 **xTenant** | **string** | The tenant ID for the request. | [default to &quot;tranquility&quot;]

### Return type

[**[]UniverseFactionsGetInner**](UniverseFactionsGetInner.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUniverseGraphics

> []int64 GetUniverseGraphics(ctx).XCompatibilityDate(xCompatibilityDate).AcceptLanguage(acceptLanguage).IfNoneMatch(ifNoneMatch).XTenant(xTenant).Execute()

Get graphics



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
	resp, r, err := apiClient.UniverseAPI.GetUniverseGraphics(context.Background()).XCompatibilityDate(xCompatibilityDate).AcceptLanguage(acceptLanguage).IfNoneMatch(ifNoneMatch).XTenant(xTenant).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UniverseAPI.GetUniverseGraphics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUniverseGraphics`: []int64
	fmt.Fprintf(os.Stdout, "Response from `UniverseAPI.GetUniverseGraphics`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUniverseGraphicsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xCompatibilityDate** | **string** | The compatibility date for the request. | 
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


## GetUniverseGraphicsGraphicId

> UniverseGraphicsGraphicIdGet GetUniverseGraphicsGraphicId(ctx, graphicId).XCompatibilityDate(xCompatibilityDate).AcceptLanguage(acceptLanguage).IfNoneMatch(ifNoneMatch).XTenant(xTenant).Execute()

Get graphic information



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
	graphicId := int64(789) // int64 | 
	xCompatibilityDate := time.Now() // string | The compatibility date for the request.
	acceptLanguage := "acceptLanguage_example" // string | The language to use for the response. (optional) (default to "en")
	ifNoneMatch := "ifNoneMatch_example" // string | The ETag of the previous request. A 304 will be returned if this matches the current ETag. (optional)
	xTenant := "xTenant_example" // string | The tenant ID for the request. (optional) (default to "tranquility")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UniverseAPI.GetUniverseGraphicsGraphicId(context.Background(), graphicId).XCompatibilityDate(xCompatibilityDate).AcceptLanguage(acceptLanguage).IfNoneMatch(ifNoneMatch).XTenant(xTenant).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UniverseAPI.GetUniverseGraphicsGraphicId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUniverseGraphicsGraphicId`: UniverseGraphicsGraphicIdGet
	fmt.Fprintf(os.Stdout, "Response from `UniverseAPI.GetUniverseGraphicsGraphicId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**graphicId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUniverseGraphicsGraphicIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xCompatibilityDate** | **string** | The compatibility date for the request. | 
 **acceptLanguage** | **string** | The language to use for the response. | [default to &quot;en&quot;]
 **ifNoneMatch** | **string** | The ETag of the previous request. A 304 will be returned if this matches the current ETag. | 
 **xTenant** | **string** | The tenant ID for the request. | [default to &quot;tranquility&quot;]

### Return type

[**UniverseGraphicsGraphicIdGet**](UniverseGraphicsGraphicIdGet.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUniverseGroups

> []int64 GetUniverseGroups(ctx).XCompatibilityDate(xCompatibilityDate).Page(page).AcceptLanguage(acceptLanguage).IfNoneMatch(ifNoneMatch).XTenant(xTenant).Execute()

Get item groups



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
	page := int32(56) // int32 |  (optional)
	acceptLanguage := "acceptLanguage_example" // string | The language to use for the response. (optional) (default to "en")
	ifNoneMatch := "ifNoneMatch_example" // string | The ETag of the previous request. A 304 will be returned if this matches the current ETag. (optional)
	xTenant := "xTenant_example" // string | The tenant ID for the request. (optional) (default to "tranquility")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UniverseAPI.GetUniverseGroups(context.Background()).XCompatibilityDate(xCompatibilityDate).Page(page).AcceptLanguage(acceptLanguage).IfNoneMatch(ifNoneMatch).XTenant(xTenant).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UniverseAPI.GetUniverseGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUniverseGroups`: []int64
	fmt.Fprintf(os.Stdout, "Response from `UniverseAPI.GetUniverseGroups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUniverseGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xCompatibilityDate** | **string** | The compatibility date for the request. | 
 **page** | **int32** |  | 
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


## GetUniverseGroupsGroupId

> UniverseGroupsGroupIdGet GetUniverseGroupsGroupId(ctx, groupId).XCompatibilityDate(xCompatibilityDate).AcceptLanguage(acceptLanguage).IfNoneMatch(ifNoneMatch).XTenant(xTenant).Execute()

Get item group information



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
	groupId := int64(789) // int64 | 
	xCompatibilityDate := time.Now() // string | The compatibility date for the request.
	acceptLanguage := "acceptLanguage_example" // string | The language to use for the response. (optional) (default to "en")
	ifNoneMatch := "ifNoneMatch_example" // string | The ETag of the previous request. A 304 will be returned if this matches the current ETag. (optional)
	xTenant := "xTenant_example" // string | The tenant ID for the request. (optional) (default to "tranquility")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UniverseAPI.GetUniverseGroupsGroupId(context.Background(), groupId).XCompatibilityDate(xCompatibilityDate).AcceptLanguage(acceptLanguage).IfNoneMatch(ifNoneMatch).XTenant(xTenant).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UniverseAPI.GetUniverseGroupsGroupId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUniverseGroupsGroupId`: UniverseGroupsGroupIdGet
	fmt.Fprintf(os.Stdout, "Response from `UniverseAPI.GetUniverseGroupsGroupId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUniverseGroupsGroupIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xCompatibilityDate** | **string** | The compatibility date for the request. | 
 **acceptLanguage** | **string** | The language to use for the response. | [default to &quot;en&quot;]
 **ifNoneMatch** | **string** | The ETag of the previous request. A 304 will be returned if this matches the current ETag. | 
 **xTenant** | **string** | The tenant ID for the request. | [default to &quot;tranquility&quot;]

### Return type

[**UniverseGroupsGroupIdGet**](UniverseGroupsGroupIdGet.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUniverseMoonsMoonId

> UniverseMoonsMoonIdGet GetUniverseMoonsMoonId(ctx, moonId).XCompatibilityDate(xCompatibilityDate).AcceptLanguage(acceptLanguage).IfNoneMatch(ifNoneMatch).XTenant(xTenant).Execute()

Get moon information



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
	moonId := int64(789) // int64 | 
	xCompatibilityDate := time.Now() // string | The compatibility date for the request.
	acceptLanguage := "acceptLanguage_example" // string | The language to use for the response. (optional) (default to "en")
	ifNoneMatch := "ifNoneMatch_example" // string | The ETag of the previous request. A 304 will be returned if this matches the current ETag. (optional)
	xTenant := "xTenant_example" // string | The tenant ID for the request. (optional) (default to "tranquility")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UniverseAPI.GetUniverseMoonsMoonId(context.Background(), moonId).XCompatibilityDate(xCompatibilityDate).AcceptLanguage(acceptLanguage).IfNoneMatch(ifNoneMatch).XTenant(xTenant).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UniverseAPI.GetUniverseMoonsMoonId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUniverseMoonsMoonId`: UniverseMoonsMoonIdGet
	fmt.Fprintf(os.Stdout, "Response from `UniverseAPI.GetUniverseMoonsMoonId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**moonId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUniverseMoonsMoonIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xCompatibilityDate** | **string** | The compatibility date for the request. | 
 **acceptLanguage** | **string** | The language to use for the response. | [default to &quot;en&quot;]
 **ifNoneMatch** | **string** | The ETag of the previous request. A 304 will be returned if this matches the current ETag. | 
 **xTenant** | **string** | The tenant ID for the request. | [default to &quot;tranquility&quot;]

### Return type

[**UniverseMoonsMoonIdGet**](UniverseMoonsMoonIdGet.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUniversePlanetsPlanetId

> UniversePlanetsPlanetIdGet GetUniversePlanetsPlanetId(ctx, planetId).XCompatibilityDate(xCompatibilityDate).AcceptLanguage(acceptLanguage).IfNoneMatch(ifNoneMatch).XTenant(xTenant).Execute()

Get planet information



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
	planetId := int64(789) // int64 | 
	xCompatibilityDate := time.Now() // string | The compatibility date for the request.
	acceptLanguage := "acceptLanguage_example" // string | The language to use for the response. (optional) (default to "en")
	ifNoneMatch := "ifNoneMatch_example" // string | The ETag of the previous request. A 304 will be returned if this matches the current ETag. (optional)
	xTenant := "xTenant_example" // string | The tenant ID for the request. (optional) (default to "tranquility")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UniverseAPI.GetUniversePlanetsPlanetId(context.Background(), planetId).XCompatibilityDate(xCompatibilityDate).AcceptLanguage(acceptLanguage).IfNoneMatch(ifNoneMatch).XTenant(xTenant).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UniverseAPI.GetUniversePlanetsPlanetId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUniversePlanetsPlanetId`: UniversePlanetsPlanetIdGet
	fmt.Fprintf(os.Stdout, "Response from `UniverseAPI.GetUniversePlanetsPlanetId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**planetId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUniversePlanetsPlanetIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xCompatibilityDate** | **string** | The compatibility date for the request. | 
 **acceptLanguage** | **string** | The language to use for the response. | [default to &quot;en&quot;]
 **ifNoneMatch** | **string** | The ETag of the previous request. A 304 will be returned if this matches the current ETag. | 
 **xTenant** | **string** | The tenant ID for the request. | [default to &quot;tranquility&quot;]

### Return type

[**UniversePlanetsPlanetIdGet**](UniversePlanetsPlanetIdGet.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUniverseRaces

> []UniverseRacesGetInner GetUniverseRaces(ctx).XCompatibilityDate(xCompatibilityDate).AcceptLanguage(acceptLanguage).IfNoneMatch(ifNoneMatch).XTenant(xTenant).Execute()

Get character races



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
	resp, r, err := apiClient.UniverseAPI.GetUniverseRaces(context.Background()).XCompatibilityDate(xCompatibilityDate).AcceptLanguage(acceptLanguage).IfNoneMatch(ifNoneMatch).XTenant(xTenant).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UniverseAPI.GetUniverseRaces``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUniverseRaces`: []UniverseRacesGetInner
	fmt.Fprintf(os.Stdout, "Response from `UniverseAPI.GetUniverseRaces`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUniverseRacesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xCompatibilityDate** | **string** | The compatibility date for the request. | 
 **acceptLanguage** | **string** | The language to use for the response. | [default to &quot;en&quot;]
 **ifNoneMatch** | **string** | The ETag of the previous request. A 304 will be returned if this matches the current ETag. | 
 **xTenant** | **string** | The tenant ID for the request. | [default to &quot;tranquility&quot;]

### Return type

[**[]UniverseRacesGetInner**](UniverseRacesGetInner.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUniverseRegions

> []int64 GetUniverseRegions(ctx).XCompatibilityDate(xCompatibilityDate).AcceptLanguage(acceptLanguage).IfNoneMatch(ifNoneMatch).XTenant(xTenant).Execute()

Get regions



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
	resp, r, err := apiClient.UniverseAPI.GetUniverseRegions(context.Background()).XCompatibilityDate(xCompatibilityDate).AcceptLanguage(acceptLanguage).IfNoneMatch(ifNoneMatch).XTenant(xTenant).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UniverseAPI.GetUniverseRegions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUniverseRegions`: []int64
	fmt.Fprintf(os.Stdout, "Response from `UniverseAPI.GetUniverseRegions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUniverseRegionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xCompatibilityDate** | **string** | The compatibility date for the request. | 
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


## GetUniverseRegionsRegionId

> UniverseRegionsRegionIdGet GetUniverseRegionsRegionId(ctx, regionId).XCompatibilityDate(xCompatibilityDate).AcceptLanguage(acceptLanguage).IfNoneMatch(ifNoneMatch).XTenant(xTenant).Execute()

Get region information



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
	regionId := int64(789) // int64 | 
	xCompatibilityDate := time.Now() // string | The compatibility date for the request.
	acceptLanguage := "acceptLanguage_example" // string | The language to use for the response. (optional) (default to "en")
	ifNoneMatch := "ifNoneMatch_example" // string | The ETag of the previous request. A 304 will be returned if this matches the current ETag. (optional)
	xTenant := "xTenant_example" // string | The tenant ID for the request. (optional) (default to "tranquility")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UniverseAPI.GetUniverseRegionsRegionId(context.Background(), regionId).XCompatibilityDate(xCompatibilityDate).AcceptLanguage(acceptLanguage).IfNoneMatch(ifNoneMatch).XTenant(xTenant).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UniverseAPI.GetUniverseRegionsRegionId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUniverseRegionsRegionId`: UniverseRegionsRegionIdGet
	fmt.Fprintf(os.Stdout, "Response from `UniverseAPI.GetUniverseRegionsRegionId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**regionId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUniverseRegionsRegionIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xCompatibilityDate** | **string** | The compatibility date for the request. | 
 **acceptLanguage** | **string** | The language to use for the response. | [default to &quot;en&quot;]
 **ifNoneMatch** | **string** | The ETag of the previous request. A 304 will be returned if this matches the current ETag. | 
 **xTenant** | **string** | The tenant ID for the request. | [default to &quot;tranquility&quot;]

### Return type

[**UniverseRegionsRegionIdGet**](UniverseRegionsRegionIdGet.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUniverseStargatesStargateId

> UniverseStargatesStargateIdGet GetUniverseStargatesStargateId(ctx, stargateId).XCompatibilityDate(xCompatibilityDate).AcceptLanguage(acceptLanguage).IfNoneMatch(ifNoneMatch).XTenant(xTenant).Execute()

Get stargate information



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
	stargateId := int64(789) // int64 | 
	xCompatibilityDate := time.Now() // string | The compatibility date for the request.
	acceptLanguage := "acceptLanguage_example" // string | The language to use for the response. (optional) (default to "en")
	ifNoneMatch := "ifNoneMatch_example" // string | The ETag of the previous request. A 304 will be returned if this matches the current ETag. (optional)
	xTenant := "xTenant_example" // string | The tenant ID for the request. (optional) (default to "tranquility")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UniverseAPI.GetUniverseStargatesStargateId(context.Background(), stargateId).XCompatibilityDate(xCompatibilityDate).AcceptLanguage(acceptLanguage).IfNoneMatch(ifNoneMatch).XTenant(xTenant).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UniverseAPI.GetUniverseStargatesStargateId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUniverseStargatesStargateId`: UniverseStargatesStargateIdGet
	fmt.Fprintf(os.Stdout, "Response from `UniverseAPI.GetUniverseStargatesStargateId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stargateId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUniverseStargatesStargateIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xCompatibilityDate** | **string** | The compatibility date for the request. | 
 **acceptLanguage** | **string** | The language to use for the response. | [default to &quot;en&quot;]
 **ifNoneMatch** | **string** | The ETag of the previous request. A 304 will be returned if this matches the current ETag. | 
 **xTenant** | **string** | The tenant ID for the request. | [default to &quot;tranquility&quot;]

### Return type

[**UniverseStargatesStargateIdGet**](UniverseStargatesStargateIdGet.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUniverseStarsStarId

> UniverseStarsStarIdGet GetUniverseStarsStarId(ctx, starId).XCompatibilityDate(xCompatibilityDate).AcceptLanguage(acceptLanguage).IfNoneMatch(ifNoneMatch).XTenant(xTenant).Execute()

Get star information



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
	starId := int64(789) // int64 | 
	xCompatibilityDate := time.Now() // string | The compatibility date for the request.
	acceptLanguage := "acceptLanguage_example" // string | The language to use for the response. (optional) (default to "en")
	ifNoneMatch := "ifNoneMatch_example" // string | The ETag of the previous request. A 304 will be returned if this matches the current ETag. (optional)
	xTenant := "xTenant_example" // string | The tenant ID for the request. (optional) (default to "tranquility")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UniverseAPI.GetUniverseStarsStarId(context.Background(), starId).XCompatibilityDate(xCompatibilityDate).AcceptLanguage(acceptLanguage).IfNoneMatch(ifNoneMatch).XTenant(xTenant).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UniverseAPI.GetUniverseStarsStarId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUniverseStarsStarId`: UniverseStarsStarIdGet
	fmt.Fprintf(os.Stdout, "Response from `UniverseAPI.GetUniverseStarsStarId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**starId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUniverseStarsStarIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xCompatibilityDate** | **string** | The compatibility date for the request. | 
 **acceptLanguage** | **string** | The language to use for the response. | [default to &quot;en&quot;]
 **ifNoneMatch** | **string** | The ETag of the previous request. A 304 will be returned if this matches the current ETag. | 
 **xTenant** | **string** | The tenant ID for the request. | [default to &quot;tranquility&quot;]

### Return type

[**UniverseStarsStarIdGet**](UniverseStarsStarIdGet.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUniverseStationsStationId

> UniverseStationsStationIdGet GetUniverseStationsStationId(ctx, stationId).XCompatibilityDate(xCompatibilityDate).AcceptLanguage(acceptLanguage).IfNoneMatch(ifNoneMatch).XTenant(xTenant).Execute()

Get station information



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
	stationId := int64(789) // int64 | 
	xCompatibilityDate := time.Now() // string | The compatibility date for the request.
	acceptLanguage := "acceptLanguage_example" // string | The language to use for the response. (optional) (default to "en")
	ifNoneMatch := "ifNoneMatch_example" // string | The ETag of the previous request. A 304 will be returned if this matches the current ETag. (optional)
	xTenant := "xTenant_example" // string | The tenant ID for the request. (optional) (default to "tranquility")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UniverseAPI.GetUniverseStationsStationId(context.Background(), stationId).XCompatibilityDate(xCompatibilityDate).AcceptLanguage(acceptLanguage).IfNoneMatch(ifNoneMatch).XTenant(xTenant).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UniverseAPI.GetUniverseStationsStationId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUniverseStationsStationId`: UniverseStationsStationIdGet
	fmt.Fprintf(os.Stdout, "Response from `UniverseAPI.GetUniverseStationsStationId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stationId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUniverseStationsStationIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xCompatibilityDate** | **string** | The compatibility date for the request. | 
 **acceptLanguage** | **string** | The language to use for the response. | [default to &quot;en&quot;]
 **ifNoneMatch** | **string** | The ETag of the previous request. A 304 will be returned if this matches the current ETag. | 
 **xTenant** | **string** | The tenant ID for the request. | [default to &quot;tranquility&quot;]

### Return type

[**UniverseStationsStationIdGet**](UniverseStationsStationIdGet.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUniverseStructures

> []int64 GetUniverseStructures(ctx).XCompatibilityDate(xCompatibilityDate).Filter(filter).AcceptLanguage(acceptLanguage).IfNoneMatch(ifNoneMatch).XTenant(xTenant).Execute()

List all public structures



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
	filter := "filter_example" // string |  (optional)
	acceptLanguage := "acceptLanguage_example" // string | The language to use for the response. (optional) (default to "en")
	ifNoneMatch := "ifNoneMatch_example" // string | The ETag of the previous request. A 304 will be returned if this matches the current ETag. (optional)
	xTenant := "xTenant_example" // string | The tenant ID for the request. (optional) (default to "tranquility")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UniverseAPI.GetUniverseStructures(context.Background()).XCompatibilityDate(xCompatibilityDate).Filter(filter).AcceptLanguage(acceptLanguage).IfNoneMatch(ifNoneMatch).XTenant(xTenant).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UniverseAPI.GetUniverseStructures``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUniverseStructures`: []int64
	fmt.Fprintf(os.Stdout, "Response from `UniverseAPI.GetUniverseStructures`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUniverseStructuresRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xCompatibilityDate** | **string** | The compatibility date for the request. | 
 **filter** | **string** |  | 
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


## GetUniverseStructuresStructureId

> UniverseStructuresStructureIdGet GetUniverseStructuresStructureId(ctx, structureId).XCompatibilityDate(xCompatibilityDate).AcceptLanguage(acceptLanguage).IfNoneMatch(ifNoneMatch).XTenant(xTenant).Execute()

Get structure information



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
	structureId := int64(789) // int64 | 
	xCompatibilityDate := time.Now() // string | The compatibility date for the request.
	acceptLanguage := "acceptLanguage_example" // string | The language to use for the response. (optional) (default to "en")
	ifNoneMatch := "ifNoneMatch_example" // string | The ETag of the previous request. A 304 will be returned if this matches the current ETag. (optional)
	xTenant := "xTenant_example" // string | The tenant ID for the request. (optional) (default to "tranquility")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UniverseAPI.GetUniverseStructuresStructureId(context.Background(), structureId).XCompatibilityDate(xCompatibilityDate).AcceptLanguage(acceptLanguage).IfNoneMatch(ifNoneMatch).XTenant(xTenant).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UniverseAPI.GetUniverseStructuresStructureId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUniverseStructuresStructureId`: UniverseStructuresStructureIdGet
	fmt.Fprintf(os.Stdout, "Response from `UniverseAPI.GetUniverseStructuresStructureId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**structureId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUniverseStructuresStructureIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xCompatibilityDate** | **string** | The compatibility date for the request. | 
 **acceptLanguage** | **string** | The language to use for the response. | [default to &quot;en&quot;]
 **ifNoneMatch** | **string** | The ETag of the previous request. A 304 will be returned if this matches the current ETag. | 
 **xTenant** | **string** | The tenant ID for the request. | [default to &quot;tranquility&quot;]

### Return type

[**UniverseStructuresStructureIdGet**](UniverseStructuresStructureIdGet.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUniverseSystemJumps

> []UniverseSystemJumpsGetInner GetUniverseSystemJumps(ctx).XCompatibilityDate(xCompatibilityDate).AcceptLanguage(acceptLanguage).IfNoneMatch(ifNoneMatch).XTenant(xTenant).Execute()

Get system jumps



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
	resp, r, err := apiClient.UniverseAPI.GetUniverseSystemJumps(context.Background()).XCompatibilityDate(xCompatibilityDate).AcceptLanguage(acceptLanguage).IfNoneMatch(ifNoneMatch).XTenant(xTenant).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UniverseAPI.GetUniverseSystemJumps``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUniverseSystemJumps`: []UniverseSystemJumpsGetInner
	fmt.Fprintf(os.Stdout, "Response from `UniverseAPI.GetUniverseSystemJumps`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUniverseSystemJumpsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xCompatibilityDate** | **string** | The compatibility date for the request. | 
 **acceptLanguage** | **string** | The language to use for the response. | [default to &quot;en&quot;]
 **ifNoneMatch** | **string** | The ETag of the previous request. A 304 will be returned if this matches the current ETag. | 
 **xTenant** | **string** | The tenant ID for the request. | [default to &quot;tranquility&quot;]

### Return type

[**[]UniverseSystemJumpsGetInner**](UniverseSystemJumpsGetInner.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUniverseSystemKills

> []UniverseSystemKillsGetInner GetUniverseSystemKills(ctx).XCompatibilityDate(xCompatibilityDate).AcceptLanguage(acceptLanguage).IfNoneMatch(ifNoneMatch).XTenant(xTenant).Execute()

Get system kills



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
	resp, r, err := apiClient.UniverseAPI.GetUniverseSystemKills(context.Background()).XCompatibilityDate(xCompatibilityDate).AcceptLanguage(acceptLanguage).IfNoneMatch(ifNoneMatch).XTenant(xTenant).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UniverseAPI.GetUniverseSystemKills``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUniverseSystemKills`: []UniverseSystemKillsGetInner
	fmt.Fprintf(os.Stdout, "Response from `UniverseAPI.GetUniverseSystemKills`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUniverseSystemKillsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xCompatibilityDate** | **string** | The compatibility date for the request. | 
 **acceptLanguage** | **string** | The language to use for the response. | [default to &quot;en&quot;]
 **ifNoneMatch** | **string** | The ETag of the previous request. A 304 will be returned if this matches the current ETag. | 
 **xTenant** | **string** | The tenant ID for the request. | [default to &quot;tranquility&quot;]

### Return type

[**[]UniverseSystemKillsGetInner**](UniverseSystemKillsGetInner.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUniverseSystems

> []int64 GetUniverseSystems(ctx).XCompatibilityDate(xCompatibilityDate).AcceptLanguage(acceptLanguage).IfNoneMatch(ifNoneMatch).XTenant(xTenant).Execute()

Get solar systems



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
	resp, r, err := apiClient.UniverseAPI.GetUniverseSystems(context.Background()).XCompatibilityDate(xCompatibilityDate).AcceptLanguage(acceptLanguage).IfNoneMatch(ifNoneMatch).XTenant(xTenant).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UniverseAPI.GetUniverseSystems``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUniverseSystems`: []int64
	fmt.Fprintf(os.Stdout, "Response from `UniverseAPI.GetUniverseSystems`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUniverseSystemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xCompatibilityDate** | **string** | The compatibility date for the request. | 
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


## GetUniverseSystemsSystemId

> UniverseSystemsSystemIdGet GetUniverseSystemsSystemId(ctx, systemId).XCompatibilityDate(xCompatibilityDate).AcceptLanguage(acceptLanguage).IfNoneMatch(ifNoneMatch).XTenant(xTenant).Execute()

Get solar system information



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
	systemId := int64(789) // int64 | 
	xCompatibilityDate := time.Now() // string | The compatibility date for the request.
	acceptLanguage := "acceptLanguage_example" // string | The language to use for the response. (optional) (default to "en")
	ifNoneMatch := "ifNoneMatch_example" // string | The ETag of the previous request. A 304 will be returned if this matches the current ETag. (optional)
	xTenant := "xTenant_example" // string | The tenant ID for the request. (optional) (default to "tranquility")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UniverseAPI.GetUniverseSystemsSystemId(context.Background(), systemId).XCompatibilityDate(xCompatibilityDate).AcceptLanguage(acceptLanguage).IfNoneMatch(ifNoneMatch).XTenant(xTenant).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UniverseAPI.GetUniverseSystemsSystemId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUniverseSystemsSystemId`: UniverseSystemsSystemIdGet
	fmt.Fprintf(os.Stdout, "Response from `UniverseAPI.GetUniverseSystemsSystemId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**systemId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUniverseSystemsSystemIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xCompatibilityDate** | **string** | The compatibility date for the request. | 
 **acceptLanguage** | **string** | The language to use for the response. | [default to &quot;en&quot;]
 **ifNoneMatch** | **string** | The ETag of the previous request. A 304 will be returned if this matches the current ETag. | 
 **xTenant** | **string** | The tenant ID for the request. | [default to &quot;tranquility&quot;]

### Return type

[**UniverseSystemsSystemIdGet**](UniverseSystemsSystemIdGet.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUniverseTypes

> []int64 GetUniverseTypes(ctx).XCompatibilityDate(xCompatibilityDate).Page(page).AcceptLanguage(acceptLanguage).IfNoneMatch(ifNoneMatch).XTenant(xTenant).Execute()

Get types



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
	page := int32(56) // int32 |  (optional)
	acceptLanguage := "acceptLanguage_example" // string | The language to use for the response. (optional) (default to "en")
	ifNoneMatch := "ifNoneMatch_example" // string | The ETag of the previous request. A 304 will be returned if this matches the current ETag. (optional)
	xTenant := "xTenant_example" // string | The tenant ID for the request. (optional) (default to "tranquility")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UniverseAPI.GetUniverseTypes(context.Background()).XCompatibilityDate(xCompatibilityDate).Page(page).AcceptLanguage(acceptLanguage).IfNoneMatch(ifNoneMatch).XTenant(xTenant).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UniverseAPI.GetUniverseTypes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUniverseTypes`: []int64
	fmt.Fprintf(os.Stdout, "Response from `UniverseAPI.GetUniverseTypes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUniverseTypesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xCompatibilityDate** | **string** | The compatibility date for the request. | 
 **page** | **int32** |  | 
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


## GetUniverseTypesTypeId

> UniverseTypesTypeIdGet GetUniverseTypesTypeId(ctx, typeId).XCompatibilityDate(xCompatibilityDate).AcceptLanguage(acceptLanguage).IfNoneMatch(ifNoneMatch).XTenant(xTenant).Execute()

Get type information



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
	typeId := int64(789) // int64 | 
	xCompatibilityDate := time.Now() // string | The compatibility date for the request.
	acceptLanguage := "acceptLanguage_example" // string | The language to use for the response. (optional) (default to "en")
	ifNoneMatch := "ifNoneMatch_example" // string | The ETag of the previous request. A 304 will be returned if this matches the current ETag. (optional)
	xTenant := "xTenant_example" // string | The tenant ID for the request. (optional) (default to "tranquility")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UniverseAPI.GetUniverseTypesTypeId(context.Background(), typeId).XCompatibilityDate(xCompatibilityDate).AcceptLanguage(acceptLanguage).IfNoneMatch(ifNoneMatch).XTenant(xTenant).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UniverseAPI.GetUniverseTypesTypeId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUniverseTypesTypeId`: UniverseTypesTypeIdGet
	fmt.Fprintf(os.Stdout, "Response from `UniverseAPI.GetUniverseTypesTypeId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**typeId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUniverseTypesTypeIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xCompatibilityDate** | **string** | The compatibility date for the request. | 
 **acceptLanguage** | **string** | The language to use for the response. | [default to &quot;en&quot;]
 **ifNoneMatch** | **string** | The ETag of the previous request. A 304 will be returned if this matches the current ETag. | 
 **xTenant** | **string** | The tenant ID for the request. | [default to &quot;tranquility&quot;]

### Return type

[**UniverseTypesTypeIdGet**](UniverseTypesTypeIdGet.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostUniverseIds

> UniverseIdsPost PostUniverseIds(ctx).XCompatibilityDate(xCompatibilityDate).RequestBody(requestBody).AcceptLanguage(acceptLanguage).IfNoneMatch(ifNoneMatch).XTenant(xTenant).Execute()

Bulk names to IDs



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
	requestBody := []string{"Property_example"} // []string | 
	acceptLanguage := "acceptLanguage_example" // string | The language to use for the response. (optional) (default to "en")
	ifNoneMatch := "ifNoneMatch_example" // string | The ETag of the previous request. A 304 will be returned if this matches the current ETag. (optional)
	xTenant := "xTenant_example" // string | The tenant ID for the request. (optional) (default to "tranquility")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UniverseAPI.PostUniverseIds(context.Background()).XCompatibilityDate(xCompatibilityDate).RequestBody(requestBody).AcceptLanguage(acceptLanguage).IfNoneMatch(ifNoneMatch).XTenant(xTenant).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UniverseAPI.PostUniverseIds``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostUniverseIds`: UniverseIdsPost
	fmt.Fprintf(os.Stdout, "Response from `UniverseAPI.PostUniverseIds`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostUniverseIdsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xCompatibilityDate** | **string** | The compatibility date for the request. | 
 **requestBody** | **[]string** |  | 
 **acceptLanguage** | **string** | The language to use for the response. | [default to &quot;en&quot;]
 **ifNoneMatch** | **string** | The ETag of the previous request. A 304 will be returned if this matches the current ETag. | 
 **xTenant** | **string** | The tenant ID for the request. | [default to &quot;tranquility&quot;]

### Return type

[**UniverseIdsPost**](UniverseIdsPost.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostUniverseNames

> []UniverseNamesPostInner PostUniverseNames(ctx).XCompatibilityDate(xCompatibilityDate).RequestBody(requestBody).AcceptLanguage(acceptLanguage).IfNoneMatch(ifNoneMatch).XTenant(xTenant).Execute()

Get names and categories for a set of IDs



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
	requestBody := []int64{int64(123)} // []int64 | 
	acceptLanguage := "acceptLanguage_example" // string | The language to use for the response. (optional) (default to "en")
	ifNoneMatch := "ifNoneMatch_example" // string | The ETag of the previous request. A 304 will be returned if this matches the current ETag. (optional)
	xTenant := "xTenant_example" // string | The tenant ID for the request. (optional) (default to "tranquility")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UniverseAPI.PostUniverseNames(context.Background()).XCompatibilityDate(xCompatibilityDate).RequestBody(requestBody).AcceptLanguage(acceptLanguage).IfNoneMatch(ifNoneMatch).XTenant(xTenant).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UniverseAPI.PostUniverseNames``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostUniverseNames`: []UniverseNamesPostInner
	fmt.Fprintf(os.Stdout, "Response from `UniverseAPI.PostUniverseNames`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostUniverseNamesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xCompatibilityDate** | **string** | The compatibility date for the request. | 
 **requestBody** | **[]int64** |  | 
 **acceptLanguage** | **string** | The language to use for the response. | [default to &quot;en&quot;]
 **ifNoneMatch** | **string** | The ETag of the previous request. A 304 will be returned if this matches the current ETag. | 
 **xTenant** | **string** | The tenant ID for the request. | [default to &quot;tranquility&quot;]

### Return type

[**[]UniverseNamesPostInner**](UniverseNamesPostInner.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

