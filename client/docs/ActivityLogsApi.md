# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV1ActivityLogsGet**](ActivityLogsApi.md#ApiV1ActivityLogsGet) | **Get** /api/v1/ActivityLogs | 

# **ApiV1ActivityLogsGet**
> []ActivityLogItem ApiV1ActivityLogsGet(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ActivityLogsApiApiV1ActivityLogsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ActivityLogsApiApiV1ActivityLogsGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **entity** | **optional.String**|  | 
 **id** | **optional.Int32**|  | 
 **ids** | [**optional.Interface of []int32**](int32.md)|  | 
 **searchDate** | **optional.Time**|  | 
 **from** | **optional.Time**|  | 
 **to** | **optional.Time**|  | 
 **page** | **optional.Int32**|  | 
 **pageSize** | **optional.Int32**|  | 

### Return type

[**[]ActivityLogItem**](ActivityLogItem.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

