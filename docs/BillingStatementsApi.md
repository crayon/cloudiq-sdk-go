# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV1BillingStatementsFileIdGet**](BillingStatementsApi.md#ApiV1BillingStatementsFileIdGet) | **Get** /api/v1/BillingStatements/file/{id} | 
[**ApiV1BillingStatementsGet**](BillingStatementsApi.md#ApiV1BillingStatementsGet) | **Get** /api/v1/BillingStatements | 
[**ApiV1BillingStatementsGroupedGet**](BillingStatementsApi.md#ApiV1BillingStatementsGroupedGet) | **Get** /api/v1/BillingStatements/grouped | 
[**ApiV1BillingStatementsIdBillingrecordsfileGet**](BillingStatementsApi.md#ApiV1BillingStatementsIdBillingrecordsfileGet) | **Get** /api/v1/BillingStatements/{id}/billingrecordsfile | 
[**ApiV1BillingStatementsIdReconciliationfileGet**](BillingStatementsApi.md#ApiV1BillingStatementsIdReconciliationfileGet) | **Get** /api/v1/BillingStatements/{id}/reconciliationfile | 

# **ApiV1BillingStatementsFileIdGet**
> ApiV1BillingStatementsFileIdGet(ctx, id)


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

# **ApiV1BillingStatementsGet**
> []BillingStatement ApiV1BillingStatementsGet(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***BillingStatementsApiApiV1BillingStatementsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BillingStatementsApiApiV1BillingStatementsGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **invoiceProfileId** | **optional.Int32**|  | 
 **organizationId** | **optional.Int32**|  | 
 **provisionType** | [**optional.Interface of ProvisionType**](.md)|  | 
 **from** | **optional.Time**|  | 
 **to** | **optional.Time**|  | 
 **page** | **optional.Int32**|  | 
 **pageSize** | **optional.Int32**|  | 

### Return type

[**[]BillingStatement**](BillingStatement.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV1BillingStatementsGroupedGet**
> []GroupedBillingStatement ApiV1BillingStatementsGroupedGet(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***BillingStatementsApiApiV1BillingStatementsGroupedGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BillingStatementsApiApiV1BillingStatementsGroupedGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **invoiceProfileId** | **optional.Int32**|  | 
 **organizationId** | **optional.Int32**|  | 
 **provisionType** | [**optional.Interface of ProvisionType**](.md)|  | 
 **from** | **optional.Time**|  | 
 **to** | **optional.Time**|  | 
 **page** | **optional.Int32**|  | 
 **pageSize** | **optional.Int32**|  | 

### Return type

[**[]GroupedBillingStatement**](GroupedBillingStatement.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV1BillingStatementsIdBillingrecordsfileGet**
> ApiV1BillingStatementsIdBillingrecordsfileGet(ctx, id)


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

# **ApiV1BillingStatementsIdReconciliationfileGet**
> ApiV1BillingStatementsIdReconciliationfileGet(ctx, id)


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

