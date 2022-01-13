# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV1BillingCyclesCspNameDictionaryGet**](BillingCyclesApi.md#ApiV1BillingCyclesCspNameDictionaryGet) | **Get** /api/v1/BillingCycles/cspNameDictionary | 
[**ApiV1BillingCyclesGet**](BillingCyclesApi.md#ApiV1BillingCyclesGet) | **Get** /api/v1/BillingCycles | 
[**ApiV1BillingCyclesProductVariantProductVariantIdGet**](BillingCyclesApi.md#ApiV1BillingCyclesProductVariantProductVariantIdGet) | **Get** /api/v1/BillingCycles/productVariant/{productVariantId} | 

# **ApiV1BillingCyclesCspNameDictionaryGet**
> map[string]string ApiV1BillingCyclesCspNameDictionaryGet(ctx, )


### Required Parameters
This endpoint does not need any parameter.

### Return type

**map[string]string**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV1BillingCyclesGet**
> []BillingCycle ApiV1BillingCyclesGet(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***BillingCyclesApiApiV1BillingCyclesGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BillingCyclesApiApiV1BillingCyclesGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **includeUnknown** | **optional.Bool**|  | [default to false]

### Return type

[**[]BillingCycle**](BillingCycle.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV1BillingCyclesProductVariantProductVariantIdGet**
> []BillingCycle ApiV1BillingCyclesProductVariantProductVariantIdGet(ctx, productVariantId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **productVariantId** | **int32**|  | 

### Return type

[**[]BillingCycle**](BillingCycle.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

