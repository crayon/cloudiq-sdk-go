# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV1AwsAccountsGet**](AwsAccountsApi.md#ApiV1AwsAccountsGet) | **Get** /api/v1/AwsAccounts | 
[**ApiV1AwsAccountsIdGet**](AwsAccountsApi.md#ApiV1AwsAccountsIdGet) | **Get** /api/v1/AwsAccounts/{id} | 
[**ApiV1AwsAccountsIdPut**](AwsAccountsApi.md#ApiV1AwsAccountsIdPut) | **Put** /api/v1/AwsAccounts/{id} | 

# **ApiV1AwsAccountsGet**
> []AwsAccount ApiV1AwsAccountsGet(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AwsAccountsApiApiV1AwsAccountsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AwsAccountsApiApiV1AwsAccountsGetOpts struct
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

[**[]AwsAccount**](AwsAccount.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV1AwsAccountsIdGet**
> AwsAccount ApiV1AwsAccountsIdGet(ctx, id)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 

### Return type

[**AwsAccount**](AwsAccount.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV1AwsAccountsIdPut**
> AwsAccount ApiV1AwsAccountsIdPut(ctx, id, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 
 **optional** | ***AwsAccountsApiApiV1AwsAccountsIdPutOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AwsAccountsApiApiV1AwsAccountsIdPutOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of AwsAccount**](AwsAccount.md)|  | 

### Return type

[**AwsAccount**](AwsAccount.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

