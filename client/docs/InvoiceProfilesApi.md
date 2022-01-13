# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV1InvoiceProfilesGet**](InvoiceProfilesApi.md#ApiV1InvoiceProfilesGet) | **Get** /api/v1/InvoiceProfiles | 
[**ApiV1InvoiceProfilesIdDelete**](InvoiceProfilesApi.md#ApiV1InvoiceProfilesIdDelete) | **Delete** /api/v1/InvoiceProfiles/{id} | 
[**ApiV1InvoiceProfilesIdGet**](InvoiceProfilesApi.md#ApiV1InvoiceProfilesIdGet) | **Get** /api/v1/InvoiceProfiles/{id} | 
[**ApiV1InvoiceProfilesIdPut**](InvoiceProfilesApi.md#ApiV1InvoiceProfilesIdPut) | **Put** /api/v1/InvoiceProfiles/{id} | 
[**ApiV1InvoiceProfilesPost**](InvoiceProfilesApi.md#ApiV1InvoiceProfilesPost) | **Post** /api/v1/InvoiceProfiles | 

# **ApiV1InvoiceProfilesGet**
> []InvoiceProfile ApiV1InvoiceProfilesGet(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***InvoiceProfilesApiApiV1InvoiceProfilesGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InvoiceProfilesApiApiV1InvoiceProfilesGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **organizationId** | **optional.Int32**|  | 
 **page** | **optional.Int32**|  | 
 **pageSize** | **optional.Int32**|  | 
 **search** | **optional.String**|  | 

### Return type

[**[]InvoiceProfile**](InvoiceProfile.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV1InvoiceProfilesIdDelete**
> bool ApiV1InvoiceProfilesIdDelete(ctx, id)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 

### Return type

**bool**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV1InvoiceProfilesIdGet**
> InvoiceProfile ApiV1InvoiceProfilesIdGet(ctx, id)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 

### Return type

[**InvoiceProfile**](InvoiceProfile.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV1InvoiceProfilesIdPut**
> InvoiceProfile ApiV1InvoiceProfilesIdPut(ctx, id, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 
 **optional** | ***InvoiceProfilesApiApiV1InvoiceProfilesIdPutOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InvoiceProfilesApiApiV1InvoiceProfilesIdPutOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of InvoiceProfile**](InvoiceProfile.md)|  | 

### Return type

[**InvoiceProfile**](InvoiceProfile.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV1InvoiceProfilesPost**
> InvoiceProfile ApiV1InvoiceProfilesPost(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***InvoiceProfilesApiApiV1InvoiceProfilesPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InvoiceProfilesApiApiV1InvoiceProfilesPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of InvoiceProfile**](InvoiceProfile.md)|  | 

### Return type

[**InvoiceProfile**](InvoiceProfile.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

