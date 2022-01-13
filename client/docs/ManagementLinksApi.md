# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV1ManagementLinksGet**](ManagementLinksApi.md#ApiV1ManagementLinksGet) | **Get** /api/v1/ManagementLinks | 
[**ApiV1ManagementLinksGroupedGet**](ManagementLinksApi.md#ApiV1ManagementLinksGroupedGet) | **Get** /api/v1/ManagementLinks/grouped | 

# **ApiV1ManagementLinksGet**
> []ManagementLink ApiV1ManagementLinksGet(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ManagementLinksApiApiV1ManagementLinksGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementLinksApiApiV1ManagementLinksGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **subscriptionIds** | [**optional.Interface of []int32**](int32.md)|  | 
 **resellerCustomerIds** | [**optional.Interface of []int32**](int32.md)|  | 
 **page** | **optional.Int32**|  | 
 **pageSize** | **optional.Int32**|  | 

### Return type

[**[]ManagementLink**](ManagementLink.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV1ManagementLinksGroupedGet**
> []ManagementLinkGrouped ApiV1ManagementLinksGroupedGet(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ManagementLinksApiApiV1ManagementLinksGroupedGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementLinksApiApiV1ManagementLinksGroupedGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **subscriptionIds** | [**optional.Interface of []int32**](int32.md)|  | 
 **resellerCustomerIds** | [**optional.Interface of []int32**](int32.md)|  | 
 **page** | **optional.Int32**|  | 
 **pageSize** | **optional.Int32**|  | 

### Return type

[**[]ManagementLinkGrouped**](ManagementLinkGrouped.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

