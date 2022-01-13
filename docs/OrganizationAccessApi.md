# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV1OrganizationAccessGet**](OrganizationAccessApi.md#ApiV1OrganizationAccessGet) | **Get** /api/v1/OrganizationAccess | 
[**ApiV1OrganizationAccessGrantGet**](OrganizationAccessApi.md#ApiV1OrganizationAccessGrantGet) | **Get** /api/v1/OrganizationAccess/grant | 
[**ApiV1OrganizationAccessPut**](OrganizationAccessApi.md#ApiV1OrganizationAccessPut) | **Put** /api/v1/OrganizationAccess | 

# **ApiV1OrganizationAccessGet**
> []OrganizationAccess ApiV1OrganizationAccessGet(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***OrganizationAccessApiApiV1OrganizationAccessGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrganizationAccessApiApiV1OrganizationAccessGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **optional.String**|  | 
 **organizationId** | **optional.Int32**|  | [default to 0]
 **page** | **optional.Int32**|  | [default to 1]
 **pageSize** | **optional.Int32**|  | [default to 10000]

### Return type

[**[]OrganizationAccess**](OrganizationAccess.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV1OrganizationAccessGrantGet**
> []OrganizationAccess ApiV1OrganizationAccessGrantGet(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***OrganizationAccessApiApiV1OrganizationAccessGrantGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrganizationAccessApiApiV1OrganizationAccessGrantGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **optional.String**|  | 
 **organizationId** | **optional.Int32**|  | [default to 0]
 **page** | **optional.Int32**|  | [default to 1]
 **pageSize** | **optional.Int32**|  | [default to 10000]

### Return type

[**[]OrganizationAccess**](OrganizationAccess.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV1OrganizationAccessPut**
> []OrganizationAccess ApiV1OrganizationAccessPut(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***OrganizationAccessApiApiV1OrganizationAccessPutOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrganizationAccessApiApiV1OrganizationAccessPutOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of []OrganizationAccess**](OrganizationAccess.md)|  | 

### Return type

[**[]OrganizationAccess**](OrganizationAccess.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

