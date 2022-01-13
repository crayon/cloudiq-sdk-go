# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV1FacebookOrdersCheckoutPost**](FacebookOrdersApi.md#ApiV1FacebookOrdersCheckoutPost) | **Post** /api/v1/FacebookOrders/checkout | 

# **ApiV1FacebookOrdersCheckoutPost**
> ApiV1FacebookOrdersCheckoutPost(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***FacebookOrdersApiApiV1FacebookOrdersCheckoutPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FacebookOrdersApiApiV1FacebookOrdersCheckoutPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of FacebookOrder**](FacebookOrder.md)|  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

