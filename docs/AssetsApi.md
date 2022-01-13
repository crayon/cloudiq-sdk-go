# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV1AssetsAssetIdPut**](AssetsApi.md#ApiV1AssetsAssetIdPut) | **Put** /api/v1/Assets/{assetId} | 
[**ApiV1AssetsAssetIdTagsDelete**](AssetsApi.md#ApiV1AssetsAssetIdTagsDelete) | **Delete** /api/v1/Assets/{assetId}/tags | 
[**ApiV1AssetsAssetIdTagsPost**](AssetsApi.md#ApiV1AssetsAssetIdTagsPost) | **Post** /api/v1/Assets/{assetId}/tags | 
[**ApiV1AssetsAssetIdTagsPut**](AssetsApi.md#ApiV1AssetsAssetIdTagsPut) | **Put** /api/v1/Assets/{assetId}/tags | 
[**ApiV1AssetsCheckoutPost**](AssetsApi.md#ApiV1AssetsCheckoutPost) | **Post** /api/v1/Assets/checkout | 
[**ApiV1AssetsGet**](AssetsApi.md#ApiV1AssetsGet) | **Get** /api/v1/Assets | 
[**ApiV1AssetsIdGet**](AssetsApi.md#ApiV1AssetsIdGet) | **Get** /api/v1/Assets/{id} | 
[**ApiV1AssetsOrdersGet**](AssetsApi.md#ApiV1AssetsOrdersGet) | **Get** /api/v1/Assets/orders | 
[**ApiV1AssetsVerifyPost**](AssetsApi.md#ApiV1AssetsVerifyPost) | **Post** /api/v1/Assets/verify | 

# **ApiV1AssetsAssetIdPut**
> ApiV1AssetsAssetIdPut(ctx, assetId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **assetId** | **int32**|  | 
 **optional** | ***AssetsApiApiV1AssetsAssetIdPutOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AssetsApiApiV1AssetsAssetIdPutOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of Asset**](Asset.md)|  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV1AssetsAssetIdTagsDelete**
> ApiV1AssetsAssetIdTagsDelete(ctx, assetId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **assetId** | **int32**|  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV1AssetsAssetIdTagsPost**
> ApiV1AssetsAssetIdTagsPost(ctx, assetId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **assetId** | **int32**|  | 
 **optional** | ***AssetsApiApiV1AssetsAssetIdTagsPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AssetsApiApiV1AssetsAssetIdTagsPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of AssetTags**](AssetTags.md)|  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV1AssetsAssetIdTagsPut**
> ApiV1AssetsAssetIdTagsPut(ctx, assetId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **assetId** | **int32**|  | 
 **optional** | ***AssetsApiApiV1AssetsAssetIdTagsPutOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AssetsApiApiV1AssetsAssetIdTagsPutOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of AssetTags**](AssetTags.md)|  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV1AssetsCheckoutPost**
> ApiV1AssetsCheckoutPost(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AssetsApiApiV1AssetsCheckoutPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AssetsApiApiV1AssetsCheckoutPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of AssetOrder**](AssetOrder.md)|  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV1AssetsGet**
> []Asset ApiV1AssetsGet(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AssetsApiApiV1AssetsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AssetsApiApiV1AssetsGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **resellerCustomerId** | **optional.Int32**|  | 
 **publisherId** | **optional.Int32**|  | 
 **externalOrderId** | **optional.String**|  | 
 **externalOrderIds** | [**optional.Interface of []string**](string.md)|  | 
 **reservationId** | **optional.String**|  | 
 **assetType** | [**optional.Interface of AssetType**](.md)|  | 
 **status** | [**optional.Interface of AssetStatus**](.md)|  | 
 **page** | **optional.Int32**|  | 
 **pageSize** | **optional.Int32**|  | 
 **search** | **optional.String**|  | 
 **sortBy** | [**optional.Interface of AssetSortBy**](.md)|  | 
 **sortOrder** | [**optional.Interface of SortOrder**](.md)|  | 

### Return type

[**[]Asset**](Asset.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV1AssetsIdGet**
> Asset ApiV1AssetsIdGet(ctx, id)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 

### Return type

[**Asset**](Asset.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV1AssetsOrdersGet**
> []AssetOrder ApiV1AssetsOrdersGet(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AssetsApiApiV1AssetsOrdersGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AssetsApiApiV1AssetsOrdersGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **resellerCustomerId** | **optional.Int32**|  | 
 **publisherId** | **optional.Int32**|  | 
 **externalOrderId** | **optional.String**|  | 
 **externalOrderIds** | [**optional.Interface of []string**](string.md)|  | 
 **reservationId** | **optional.String**|  | 
 **assetType** | [**optional.Interface of AssetType**](.md)|  | 
 **status** | [**optional.Interface of AssetStatus**](.md)|  | 
 **page** | **optional.Int32**|  | 
 **pageSize** | **optional.Int32**|  | 
 **search** | **optional.String**|  | 
 **sortBy** | [**optional.Interface of AssetSortBy**](.md)|  | 
 **sortOrder** | [**optional.Interface of SortOrder**](.md)|  | 

### Return type

[**[]AssetOrder**](AssetOrder.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV1AssetsVerifyPost**
> AssetOrder ApiV1AssetsVerifyPost(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AssetsApiApiV1AssetsVerifyPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AssetsApiApiV1AssetsVerifyPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of AssetOrder**](AssetOrder.md)|  | 

### Return type

[**AssetOrder**](AssetOrder.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

