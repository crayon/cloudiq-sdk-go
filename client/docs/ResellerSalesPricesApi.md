# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV1ResellerSalesPricesCurrentGet**](ResellerSalesPricesApi.md#ApiV1ResellerSalesPricesCurrentGet) | **Get** /api/v1/ResellerSalesPrices/current | 
[**ApiV1ResellerSalesPricesDelete**](ResellerSalesPricesApi.md#ApiV1ResellerSalesPricesDelete) | **Delete** /api/v1/ResellerSalesPrices | 
[**ApiV1ResellerSalesPricesGet**](ResellerSalesPricesApi.md#ApiV1ResellerSalesPricesGet) | **Get** /api/v1/ResellerSalesPrices | 
[**ApiV1ResellerSalesPricesOldFromDatePut**](ResellerSalesPricesApi.md#ApiV1ResellerSalesPricesOldFromDatePut) | **Put** /api/v1/ResellerSalesPrices/{oldFromDate} | 
[**ApiV1ResellerSalesPricesPost**](ResellerSalesPricesApi.md#ApiV1ResellerSalesPricesPost) | **Post** /api/v1/ResellerSalesPrices | 
[**ApiV1ResellerSalesPricesTogglePost**](ResellerSalesPricesApi.md#ApiV1ResellerSalesPricesTogglePost) | **Post** /api/v1/ResellerSalesPrices/toggle | 

# **ApiV1ResellerSalesPricesCurrentGet**
> ResellerSalesPrice ApiV1ResellerSalesPricesCurrentGet(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ResellerSalesPricesApiApiV1ResellerSalesPricesCurrentGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ResellerSalesPricesApiApiV1ResellerSalesPricesCurrentGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | [**optional.Interface of ResellerSalesPriceType**](.md)|  | 
 **objectId** | **optional.Int32**|  | 
 **objectType** | [**optional.Interface of ResellerSalesPriceObjectType**](.md)|  | 
 **fromDate** | **optional.Time**|  | 

### Return type

[**ResellerSalesPrice**](ResellerSalesPrice.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV1ResellerSalesPricesDelete**
> ApiV1ResellerSalesPricesDelete(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ResellerSalesPricesApiApiV1ResellerSalesPricesDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ResellerSalesPricesApiApiV1ResellerSalesPricesDeleteOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | [**optional.Interface of ResellerSalesPriceType**](.md)|  | 
 **objectId** | **optional.Int32**|  | 
 **objectType** | [**optional.Interface of ResellerSalesPriceObjectType**](.md)|  | 
 **fromDate** | **optional.Time**|  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV1ResellerSalesPricesGet**
> []ResellerSalesPrice ApiV1ResellerSalesPricesGet(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ResellerSalesPricesApiApiV1ResellerSalesPricesGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ResellerSalesPricesApiApiV1ResellerSalesPricesGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | [**optional.Interface of ResellerSalesPriceType**](.md)|  | 
 **objectId** | **optional.Int32**|  | 
 **objectType** | [**optional.Interface of ResellerSalesPriceObjectType**](.md)|  | 
 **fromDate** | **optional.Time**|  | 

### Return type

[**[]ResellerSalesPrice**](ResellerSalesPrice.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV1ResellerSalesPricesOldFromDatePut**
> ResellerSalesPrice ApiV1ResellerSalesPricesOldFromDatePut(ctx, oldFromDate, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **oldFromDate** | **time.Time**|  | 
 **optional** | ***ResellerSalesPricesApiApiV1ResellerSalesPricesOldFromDatePutOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ResellerSalesPricesApiApiV1ResellerSalesPricesOldFromDatePutOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of ResellerSalesPrice**](ResellerSalesPrice.md)|  | 

### Return type

[**ResellerSalesPrice**](ResellerSalesPrice.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV1ResellerSalesPricesPost**
> ResellerSalesPrice ApiV1ResellerSalesPricesPost(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ResellerSalesPricesApiApiV1ResellerSalesPricesPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ResellerSalesPricesApiApiV1ResellerSalesPricesPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of ResellerSalesPrice**](ResellerSalesPrice.md)|  | 

### Return type

[**ResellerSalesPrice**](ResellerSalesPrice.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV1ResellerSalesPricesTogglePost**
> ApiV1ResellerSalesPricesTogglePost(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ResellerSalesPricesApiApiV1ResellerSalesPricesTogglePostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ResellerSalesPricesApiApiV1ResellerSalesPricesTogglePostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of ResellerSalesPriceToggle**](ResellerSalesPriceToggle.md)|  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

