# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV1CustomerTenantsCustomerTenantIdAzurePlanGet**](CustomerTenantsApi.md#ApiV1CustomerTenantsCustomerTenantIdAzurePlanGet) | **Get** /api/v1/CustomerTenants/{customerTenantId}/azurePlan | 
[**ApiV1CustomerTenantsExistingPost**](CustomerTenantsApi.md#ApiV1CustomerTenantsExistingPost) | **Post** /api/v1/CustomerTenants/existing | 
[**ApiV1CustomerTenantsGet**](CustomerTenantsApi.md#ApiV1CustomerTenantsGet) | **Get** /api/v1/CustomerTenants | 
[**ApiV1CustomerTenantsIdDelete**](CustomerTenantsApi.md#ApiV1CustomerTenantsIdDelete) | **Delete** /api/v1/CustomerTenants/{id} | 
[**ApiV1CustomerTenantsIdDetailedGet**](CustomerTenantsApi.md#ApiV1CustomerTenantsIdDetailedGet) | **Get** /api/v1/CustomerTenants/{id}/detailed | 
[**ApiV1CustomerTenantsIdGet**](CustomerTenantsApi.md#ApiV1CustomerTenantsIdGet) | **Get** /api/v1/CustomerTenants/{id} | 
[**ApiV1CustomerTenantsIdPut**](CustomerTenantsApi.md#ApiV1CustomerTenantsIdPut) | **Put** /api/v1/CustomerTenants/{id} | 
[**ApiV1CustomerTenantsPost**](CustomerTenantsApi.md#ApiV1CustomerTenantsPost) | **Post** /api/v1/CustomerTenants | 

# **ApiV1CustomerTenantsCustomerTenantIdAzurePlanGet**
> AzurePlan ApiV1CustomerTenantsCustomerTenantIdAzurePlanGet(ctx, customerTenantId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **customerTenantId** | **int32**|  | 

### Return type

[**AzurePlan**](AzurePlan.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV1CustomerTenantsExistingPost**
> CustomerTenantDetailed ApiV1CustomerTenantsExistingPost(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CustomerTenantsApiApiV1CustomerTenantsExistingPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerTenantsApiApiV1CustomerTenantsExistingPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of CustomerTenantDetailed**](CustomerTenantDetailed.md)|  | 
 **syncFromPublisher** | **optional.**|  | [default to false]

### Return type

[**CustomerTenantDetailed**](CustomerTenantDetailed.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV1CustomerTenantsGet**
> []CustomerTenant ApiV1CustomerTenantsGet(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CustomerTenantsApiApiV1CustomerTenantsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerTenantsApiApiV1CustomerTenantsGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **organizationId** | **optional.Int32**|  | 
 **publisherId** | **optional.Int32**|  | 
 **programId** | **optional.Int32**|  | 
 **consumerId** | **optional.Int32**|  | 
 **domain** | **optional.String**|  | 
 **domainPrefix** | **optional.String**|  | 
 **customerTenantType** | [**optional.Interface of CustomerTenantType**](.md)|  | 
 **invoiceProfileId** | **optional.Int32**|  | 
 **page** | **optional.Int32**|  | 
 **pageSize** | **optional.Int32**|  | 
 **search** | **optional.String**|  | 

### Return type

[**[]CustomerTenant**](CustomerTenant.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV1CustomerTenantsIdDelete**
> ApiV1CustomerTenantsIdDelete(ctx, id)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV1CustomerTenantsIdDetailedGet**
> CustomerTenantDetailed ApiV1CustomerTenantsIdDetailedGet(ctx, id)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 

### Return type

[**CustomerTenantDetailed**](CustomerTenantDetailed.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV1CustomerTenantsIdGet**
> CustomerTenant ApiV1CustomerTenantsIdGet(ctx, id)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 

### Return type

[**CustomerTenant**](CustomerTenant.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV1CustomerTenantsIdPut**
> CustomerTenantDetailed ApiV1CustomerTenantsIdPut(ctx, id, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 
 **optional** | ***CustomerTenantsApiApiV1CustomerTenantsIdPutOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerTenantsApiApiV1CustomerTenantsIdPutOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of CustomerTenantDetailed**](CustomerTenantDetailed.md)|  | 

### Return type

[**CustomerTenantDetailed**](CustomerTenantDetailed.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV1CustomerTenantsPost**
> CustomerTenantDetailed ApiV1CustomerTenantsPost(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CustomerTenantsApiApiV1CustomerTenantsPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerTenantsApiApiV1CustomerTenantsPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of CustomerTenantDetailed**](CustomerTenantDetailed.md)|  | 

### Return type

[**CustomerTenantDetailed**](CustomerTenantDetailed.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

