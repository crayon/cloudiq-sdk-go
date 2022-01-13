# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV1ProductContainersGet**](ProductContainersApi.md#ApiV1ProductContainersGet) | **Get** /api/v1/ProductContainers | 
[**ApiV1ProductContainersGetorcreateshoppingcartGet**](ProductContainersApi.md#ApiV1ProductContainersGetorcreateshoppingcartGet) | **Get** /api/v1/ProductContainers/getorcreateshoppingcart | 
[**ApiV1ProductContainersIdDelete**](ProductContainersApi.md#ApiV1ProductContainersIdDelete) | **Delete** /api/v1/ProductContainers/{id} | 
[**ApiV1ProductContainersIdGet**](ProductContainersApi.md#ApiV1ProductContainersIdGet) | **Get** /api/v1/ProductContainers/{id} | 
[**ApiV1ProductContainersIdPut**](ProductContainersApi.md#ApiV1ProductContainersIdPut) | **Put** /api/v1/ProductContainers/{id} | 
[**ApiV1ProductContainersProductContainerIdRowProductRowIdPatch**](ProductContainersApi.md#ApiV1ProductContainersProductContainerIdRowProductRowIdPatch) | **Patch** /api/v1/ProductContainers/{productContainerId}/row/{productRowId} | 
[**ApiV1ProductContainersReportbymonthPost**](ProductContainersApi.md#ApiV1ProductContainersReportbymonthPost) | **Post** /api/v1/ProductContainers/reportbymonth | 
[**ApiV1ProductContainersRowissuesIdGet**](ProductContainersApi.md#ApiV1ProductContainersRowissuesIdGet) | **Get** /api/v1/ProductContainers/rowissues/{id} | 

# **ApiV1ProductContainersGet**
> []ProductContainer ApiV1ProductContainersGet(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ProductContainersApiApiV1ProductContainersGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProductContainersApiApiV1ProductContainersGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **organizationId** | **optional.Int32**|  | 
 **search** | **optional.String**|  | 
 **page** | **optional.Int32**|  | 
 **pageSize** | **optional.Int32**|  | 
 **activeDraft** | **optional.Bool**|  | 
 **programId** | **optional.Int32**|  | 
 **year** | **optional.Int32**|  | 
 **month** | **optional.Int32**|  | 
 **userId** | **optional.String**|  | 
 **sentByUserId** | **optional.String**|  | 
 **type_** | [**optional.Interface of ProductContainerType**](.md)|  | 
 **category** | [**optional.Interface of ProductContainerCategory**](.md)|  | 
 **from** | **optional.Time**|  | 
 **to** | **optional.Time**|  | 
 **includeRemoved** | **optional.Bool**|  | 
 **includeSubsidiaries** | **optional.Bool**|  | 

### Return type

[**[]ProductContainer**](ProductContainer.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV1ProductContainersGetorcreateshoppingcartGet**
> ProductContainer ApiV1ProductContainersGetorcreateshoppingcartGet(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ProductContainersApiApiV1ProductContainersGetorcreateshoppingcartGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProductContainersApiApiV1ProductContainersGetorcreateshoppingcartGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **organizationId** | **optional.Int32**|  | 

### Return type

[**ProductContainer**](ProductContainer.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV1ProductContainersIdDelete**
> bool ApiV1ProductContainersIdDelete(ctx, id)


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

# **ApiV1ProductContainersIdGet**
> ProductContainer ApiV1ProductContainersIdGet(ctx, id)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 

### Return type

[**ProductContainer**](ProductContainer.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV1ProductContainersIdPut**
> ProductContainer ApiV1ProductContainersIdPut(ctx, id, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 
 **optional** | ***ProductContainersApiApiV1ProductContainersIdPutOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProductContainersApiApiV1ProductContainersIdPutOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of ProductContainer**](ProductContainer.md)|  | 
 **requireEulaAnalysis** | **optional.**|  | [default to false]

### Return type

[**ProductContainer**](ProductContainer.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV1ProductContainersProductContainerIdRowProductRowIdPatch**
> ProductContainer ApiV1ProductContainersProductContainerIdRowProductRowIdPatch(ctx, productContainerId, productRowId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **productContainerId** | **int32**|  | 
  **productRowId** | **int32**|  | 
 **optional** | ***ProductContainersApiApiV1ProductContainersProductContainerIdRowProductRowIdPatchOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProductContainersApiApiV1ProductContainersProductContainerIdRowProductRowIdPatchOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of ProductRowPatch**](ProductRowPatch.md)|  | 

### Return type

[**ProductContainer**](ProductContainer.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV1ProductContainersReportbymonthPost**
> ProductContainer ApiV1ProductContainersReportbymonthPost(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ProductContainersApiApiV1ProductContainersReportbymonthPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProductContainersApiApiV1ProductContainersReportbymonthPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **year** | **optional.Int32**|  | 
 **month** | **optional.Int32**|  | 
 **programId** | **optional.Int32**|  | 
 **organizationId** | **optional.Int32**|  | 
 **copyLast** | **optional.Bool**|  | 

### Return type

[**ProductContainer**](ProductContainer.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV1ProductContainersRowissuesIdGet**
> ProductContainer ApiV1ProductContainersRowissuesIdGet(ctx, id)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 

### Return type

[**ProductContainer**](ProductContainer.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

