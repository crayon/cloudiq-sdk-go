# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV1ProgramsGet**](ProgramsApi.md#ApiV1ProgramsGet) | **Get** /api/v1/Programs | 
[**ApiV1ProgramsIdGet**](ProgramsApi.md#ApiV1ProgramsIdGet) | **Get** /api/v1/Programs/{id} | 

# **ApiV1ProgramsGet**
> []Program ApiV1ProgramsGet(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ProgramsApiApiV1ProgramsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProgramsApiApiV1ProgramsGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **publisherId** | **optional.Int32**|  | 
 **programType** | [**optional.Interface of ProgramType**](.md)|  | 
 **page** | **optional.Int32**|  | 
 **pageSize** | **optional.Int32**|  | 
 **search** | **optional.String**|  | 
 **organizationId** | **optional.Int32**|  | 

### Return type

[**[]Program**](Program.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV1ProgramsIdGet**
> Program ApiV1ProgramsIdGet(ctx, id)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 

### Return type

[**Program**](Program.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

