# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV1BlogItemsGet**](BlogItemsApi.md#ApiV1BlogItemsGet) | **Get** /api/v1/BlogItems | 

# **ApiV1BlogItemsGet**
> []BlogItem ApiV1BlogItemsGet(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***BlogItemsApiApiV1BlogItemsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BlogItemsApiApiV1BlogItemsGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **count** | **optional.Int32**|  | [default to 6]
 **organizationId** | **optional.Int32**|  | [default to 0]

### Return type

[**[]BlogItem**](BlogItem.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

