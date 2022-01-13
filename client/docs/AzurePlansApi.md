# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV1AzurePlansAzurePlanIdAzureSubscriptionsGet**](AzurePlansApi.md#ApiV1AzurePlansAzurePlanIdAzureSubscriptionsGet) | **Get** /api/v1/AzurePlans/{azurePlanId}/azureSubscriptions | 
[**ApiV1AzurePlansAzurePlanIdAzureSubscriptionsIdAssignUniqueAdminPut**](AzurePlansApi.md#ApiV1AzurePlansAzurePlanIdAzureSubscriptionsIdAssignUniqueAdminPut) | **Put** /api/v1/AzurePlans/{azurePlanId}/azureSubscriptions/{id}/assign-unique-admin | 
[**ApiV1AzurePlansAzurePlanIdAzureSubscriptionsIdCancelPost**](AzurePlansApi.md#ApiV1AzurePlansAzurePlanIdAzureSubscriptionsIdCancelPost) | **Post** /api/v1/AzurePlans/{azurePlanId}/azureSubscriptions/{id}/cancel | 
[**ApiV1AzurePlansAzurePlanIdAzureSubscriptionsIdEnablePost**](AzurePlansApi.md#ApiV1AzurePlansAzurePlanIdAzureSubscriptionsIdEnablePost) | **Post** /api/v1/AzurePlans/{azurePlanId}/azureSubscriptions/{id}/enable | 
[**ApiV1AzurePlansAzurePlanIdAzureSubscriptionsIdGet**](AzurePlansApi.md#ApiV1AzurePlansAzurePlanIdAzureSubscriptionsIdGet) | **Get** /api/v1/AzurePlans/{azurePlanId}/azureSubscriptions/{id} | 
[**ApiV1AzurePlansAzurePlanIdAzureSubscriptionsIdPut**](AzurePlansApi.md#ApiV1AzurePlansAzurePlanIdAzureSubscriptionsIdPut) | **Put** /api/v1/AzurePlans/{azurePlanId}/azureSubscriptions/{id} | 
[**ApiV1AzurePlansAzurePlanIdAzureSubscriptionsIdRenamePatch**](AzurePlansApi.md#ApiV1AzurePlansAzurePlanIdAzureSubscriptionsIdRenamePatch) | **Patch** /api/v1/AzurePlans/{azurePlanId}/azureSubscriptions/{id}/rename | 
[**ApiV1AzurePlansAzurePlanIdAzureSubscriptionsPost**](AzurePlansApi.md#ApiV1AzurePlansAzurePlanIdAzureSubscriptionsPost) | **Post** /api/v1/AzurePlans/{azurePlanId}/azureSubscriptions | 
[**ApiV1AzurePlansAzurePlanIdGet**](AzurePlansApi.md#ApiV1AzurePlansAzurePlanIdGet) | **Get** /api/v1/AzurePlans/{azurePlanId} | 

# **ApiV1AzurePlansAzurePlanIdAzureSubscriptionsGet**
> []AzureSubscription ApiV1AzurePlansAzurePlanIdAzureSubscriptionsGet(ctx, azurePlanId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **azurePlanId** | **int32**|  | 
 **optional** | ***AzurePlansApiApiV1AzurePlansAzurePlanIdAzureSubscriptionsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AzurePlansApiApiV1AzurePlansAzurePlanIdAzureSubscriptionsGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **search** | **optional.String**|  | 
 **page** | **optional.Int32**|  | 
 **pageSize** | **optional.Int32**|  | 

### Return type

[**[]AzureSubscription**](AzureSubscription.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV1AzurePlansAzurePlanIdAzureSubscriptionsIdAssignUniqueAdminPut**
> bool ApiV1AzurePlansAzurePlanIdAzureSubscriptionsIdAssignUniqueAdminPut(ctx, azurePlanId, id, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **azurePlanId** | **int32**|  | 
  **id** | **int32**|  | 
 **optional** | ***AzurePlansApiApiV1AzurePlansAzurePlanIdAzureSubscriptionsIdAssignUniqueAdminPutOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AzurePlansApiApiV1AzurePlansAzurePlanIdAzureSubscriptionsIdAssignUniqueAdminPutOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of AzureSubscriptionAssignAdmin**](AzureSubscriptionAssignAdmin.md)|  | 

### Return type

**bool**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV1AzurePlansAzurePlanIdAzureSubscriptionsIdCancelPost**
> AzureSubscriptionUpdated ApiV1AzurePlansAzurePlanIdAzureSubscriptionsIdCancelPost(ctx, azurePlanId, id)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **azurePlanId** | **int32**|  | 
  **id** | **int32**|  | 

### Return type

[**AzureSubscriptionUpdated**](AzureSubscriptionUpdated.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV1AzurePlansAzurePlanIdAzureSubscriptionsIdEnablePost**
> AzureSubscriptionUpdated ApiV1AzurePlansAzurePlanIdAzureSubscriptionsIdEnablePost(ctx, azurePlanId, id)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **azurePlanId** | **int32**|  | 
  **id** | **int32**|  | 

### Return type

[**AzureSubscriptionUpdated**](AzureSubscriptionUpdated.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV1AzurePlansAzurePlanIdAzureSubscriptionsIdGet**
> AzureSubscription ApiV1AzurePlansAzurePlanIdAzureSubscriptionsIdGet(ctx, azurePlanId, id)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **azurePlanId** | **int32**|  | 
  **id** | **int32**|  | 

### Return type

[**AzureSubscription**](AzureSubscription.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV1AzurePlansAzurePlanIdAzureSubscriptionsIdPut**
> AzureSubscription ApiV1AzurePlansAzurePlanIdAzureSubscriptionsIdPut(ctx, azurePlanId, id, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **azurePlanId** | **int32**|  | 
  **id** | **int32**|  | 
 **optional** | ***AzurePlansApiApiV1AzurePlansAzurePlanIdAzureSubscriptionsIdPutOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AzurePlansApiApiV1AzurePlansAzurePlanIdAzureSubscriptionsIdPutOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of PutAzureSubscription**](PutAzureSubscription.md)|  | 

### Return type

[**AzureSubscription**](AzureSubscription.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV1AzurePlansAzurePlanIdAzureSubscriptionsIdRenamePatch**
> AzureSubscriptionUpdated ApiV1AzurePlansAzurePlanIdAzureSubscriptionsIdRenamePatch(ctx, azurePlanId, id, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **azurePlanId** | **int32**|  | 
  **id** | **int32**|  | 
 **optional** | ***AzurePlansApiApiV1AzurePlansAzurePlanIdAzureSubscriptionsIdRenamePatchOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AzurePlansApiApiV1AzurePlansAzurePlanIdAzureSubscriptionsIdRenamePatchOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of AzureSubscriptionRename**](AzureSubscriptionRename.md)|  | 

### Return type

[**AzureSubscriptionUpdated**](AzureSubscriptionUpdated.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV1AzurePlansAzurePlanIdAzureSubscriptionsPost**
> ApiV1AzurePlansAzurePlanIdAzureSubscriptionsPost(ctx, azurePlanId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **azurePlanId** | **int32**|  | 
 **optional** | ***AzurePlansApiApiV1AzurePlansAzurePlanIdAzureSubscriptionsPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AzurePlansApiApiV1AzurePlansAzurePlanIdAzureSubscriptionsPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of CreateAzureSubscriptionRequest**](CreateAzureSubscriptionRequest.md)|  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV1AzurePlansAzurePlanIdGet**
> AzurePlan ApiV1AzurePlansAzurePlanIdGet(ctx, azurePlanId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **azurePlanId** | **int32**|  | 

### Return type

[**AzurePlan**](AzurePlan.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

