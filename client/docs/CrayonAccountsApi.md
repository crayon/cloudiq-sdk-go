# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV1CrayonAccountsGet**](CrayonAccountsApi.md#ApiV1CrayonAccountsGet) | **Get** /api/v1/CrayonAccounts | 
[**ApiV1CrayonAccountsIdGet**](CrayonAccountsApi.md#ApiV1CrayonAccountsIdGet) | **Get** /api/v1/CrayonAccounts/{id} | 
[**ApiV1CrayonAccountsIdPut**](CrayonAccountsApi.md#ApiV1CrayonAccountsIdPut) | **Put** /api/v1/CrayonAccounts/{id} | 
[**ApiV1CrayonAccountsPost**](CrayonAccountsApi.md#ApiV1CrayonAccountsPost) | **Post** /api/v1/CrayonAccounts | 

# **ApiV1CrayonAccountsGet**
> []CrayonAccount ApiV1CrayonAccountsGet(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CrayonAccountsApiApiV1CrayonAccountsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CrayonAccountsApiApiV1CrayonAccountsGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **organizationId** | **optional.Int32**|  | 
 **publisherId** | **optional.Int32**|  | 
 **consumerId** | **optional.Int32**|  | 
 **customerTenantType** | [**optional.Interface of CustomerTenantType**](.md)|  | 
 **page** | **optional.Int32**|  | 
 **pageSize** | **optional.Int32**|  | 
 **search** | **optional.String**|  | 

### Return type

[**[]CrayonAccount**](CrayonAccount.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV1CrayonAccountsIdGet**
> CrayonAccount ApiV1CrayonAccountsIdGet(ctx, id)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 

### Return type

[**CrayonAccount**](CrayonAccount.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV1CrayonAccountsIdPut**
> CrayonAccount ApiV1CrayonAccountsIdPut(ctx, id, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 
 **optional** | ***CrayonAccountsApiApiV1CrayonAccountsIdPutOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CrayonAccountsApiApiV1CrayonAccountsIdPutOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of CrayonAccount**](CrayonAccount.md)|  | 

### Return type

[**CrayonAccount**](CrayonAccount.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV1CrayonAccountsPost**
> CrayonAccount ApiV1CrayonAccountsPost(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CrayonAccountsApiApiV1CrayonAccountsPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CrayonAccountsApiApiV1CrayonAccountsPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of CrayonAccount**](CrayonAccount.md)|  | 

### Return type

[**CrayonAccount**](CrayonAccount.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

