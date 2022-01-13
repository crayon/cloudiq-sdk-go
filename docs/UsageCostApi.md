# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV1UsageCostGetForCategoryPost**](UsageCostApi.md#ApiV1UsageCostGetForCategoryPost) | **Post** /api/v1/UsageCost/getForCategory | 
[**ApiV1UsageCostGetForResourceGroupPost**](UsageCostApi.md#ApiV1UsageCostGetForResourceGroupPost) | **Post** /api/v1/UsageCost/getForResourceGroup | 
[**ApiV1UsageCostGetForSubcategoryPost**](UsageCostApi.md#ApiV1UsageCostGetForSubcategoryPost) | **Post** /api/v1/UsageCost/getForSubcategory | 
[**ApiV1UsageCostGetForSubscriptionPost**](UsageCostApi.md#ApiV1UsageCostGetForSubscriptionPost) | **Post** /api/v1/UsageCost/getForSubscription | 
[**ApiV1UsageCostGetForSubscriptionResourceGroupsPost**](UsageCostApi.md#ApiV1UsageCostGetForSubscriptionResourceGroupsPost) | **Post** /api/v1/UsageCost/getForSubscription/resourceGroups | 
[**ApiV1UsageCostOrganizationOrganizationIdGet**](UsageCostApi.md#ApiV1UsageCostOrganizationOrganizationIdGet) | **Get** /api/v1/UsageCost/organization/{organizationId} | 

# **ApiV1UsageCostGetForCategoryPost**
> []CategoryUsageCost ApiV1UsageCostGetForCategoryPost(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UsageCostApiApiV1UsageCostGetForCategoryPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsageCostApiApiV1UsageCostGetForCategoryPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of CategoryUsageCostRequest**](CategoryUsageCostRequest.md)|  | 

### Return type

[**[]CategoryUsageCost**](CategoryUsageCost.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV1UsageCostGetForResourceGroupPost**
> []ResourceGroupUsageCost ApiV1UsageCostGetForResourceGroupPost(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UsageCostApiApiV1UsageCostGetForResourceGroupPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsageCostApiApiV1UsageCostGetForResourceGroupPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of ResourceGroupUsageCostRequest**](ResourceGroupUsageCostRequest.md)|  | 

### Return type

[**[]ResourceGroupUsageCost**](ResourceGroupUsageCost.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV1UsageCostGetForSubcategoryPost**
> []SubcategoryUsageCost ApiV1UsageCostGetForSubcategoryPost(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UsageCostApiApiV1UsageCostGetForSubcategoryPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsageCostApiApiV1UsageCostGetForSubcategoryPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of SubcategoryUsageCostRequest**](SubcategoryUsageCostRequest.md)|  | 

### Return type

[**[]SubcategoryUsageCost**](SubcategoryUsageCost.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV1UsageCostGetForSubscriptionPost**
> []SubscriptionUsageCost ApiV1UsageCostGetForSubscriptionPost(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UsageCostApiApiV1UsageCostGetForSubscriptionPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsageCostApiApiV1UsageCostGetForSubscriptionPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of SubscriptionUsageCostRequest**](SubscriptionUsageCostRequest.md)|  | 

### Return type

[**[]SubscriptionUsageCost**](SubscriptionUsageCost.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV1UsageCostGetForSubscriptionResourceGroupsPost**
> []SubscriptionResourceGroupUsageCost ApiV1UsageCostGetForSubscriptionResourceGroupsPost(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UsageCostApiApiV1UsageCostGetForSubscriptionResourceGroupsPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsageCostApiApiV1UsageCostGetForSubscriptionResourceGroupsPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of SubscriptionUsageCostRequest**](SubscriptionUsageCostRequest.md)|  | 

### Return type

[**[]SubscriptionResourceGroupUsageCost**](SubscriptionResourceGroupUsageCost.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV1UsageCostOrganizationOrganizationIdGet**
> []OrganizationUsageCost ApiV1UsageCostOrganizationOrganizationIdGet(ctx, organizationId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **organizationId** | **int32**|  | 
 **optional** | ***UsageCostApiApiV1UsageCostOrganizationOrganizationIdGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsageCostApiApiV1UsageCostOrganizationOrganizationIdGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **from** | **optional.Time**|  | 
 **to** | **optional.Time**|  | 

### Return type

[**[]OrganizationUsageCost**](OrganizationUsageCost.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

