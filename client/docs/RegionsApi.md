# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV1RegionsBycodeGet**](RegionsApi.md#ApiV1RegionsBycodeGet) | **Get** /api/v1/Regions/bycode | 
[**ApiV1RegionsGet**](RegionsApi.md#ApiV1RegionsGet) | **Get** /api/v1/Regions | 

# **ApiV1RegionsBycodeGet**
> Region ApiV1RegionsBycodeGet(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***RegionsApiApiV1RegionsBycodeGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RegionsApiApiV1RegionsBycodeGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **regionCode** | **optional.String**|  | 
 **regionList** | [**optional.Interface of RegionList**](.md)|  | 

### Return type

[**Region**](Region.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV1RegionsGet**
> []Region ApiV1RegionsGet(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***RegionsApiApiV1RegionsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RegionsApiApiV1RegionsGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **regionList** | [**optional.Interface of RegionList**](.md)|  | 
 **organizationId** | **optional.Int32**|  | 
 **page** | **optional.Int32**|  | 
 **pageSize** | **optional.Int32**|  | 
 **search** | **optional.String**|  | 

### Return type

[**[]Region**](Region.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

