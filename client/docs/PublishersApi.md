# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV1PublishersGet**](PublishersApi.md#ApiV1PublishersGet) | **Get** /api/v1/Publishers | 
[**ApiV1PublishersIdGet**](PublishersApi.md#ApiV1PublishersIdGet) | **Get** /api/v1/Publishers/{id} | 

# **ApiV1PublishersGet**
> []Publisher ApiV1PublishersGet(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PublishersApiApiV1PublishersGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PublishersApiApiV1PublishersGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **names** | [**optional.Interface of []string**](string.md)|  | 
 **page** | **optional.Int32**|  | 
 **pageSize** | **optional.Int32**|  | 
 **search** | **optional.String**|  | 
 **programType** | [**optional.Interface of ProgramType**](.md)|  | 

### Return type

[**[]Publisher**](Publisher.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV1PublishersIdGet**
> Publisher ApiV1PublishersIdGet(ctx, id)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 

### Return type

[**Publisher**](Publisher.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

