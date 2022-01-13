# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV1SubscriptionsGet**](SubscriptionsApi.md#ApiV1SubscriptionsGet) | **Get** /api/v1/Subscriptions | 
[**ApiV1SubscriptionsIdActivationlinkGet**](SubscriptionsApi.md#ApiV1SubscriptionsIdActivationlinkGet) | **Get** /api/v1/Subscriptions/{id}/activationlink | 
[**ApiV1SubscriptionsIdGet**](SubscriptionsApi.md#ApiV1SubscriptionsIdGet) | **Get** /api/v1/Subscriptions/{id} | 
[**ApiV1SubscriptionsIdPut**](SubscriptionsApi.md#ApiV1SubscriptionsIdPut) | **Put** /api/v1/Subscriptions/{id} | 
[**ApiV1SubscriptionsIdTransitionEligibilitiesEligibilityTypeGet**](SubscriptionsApi.md#ApiV1SubscriptionsIdTransitionEligibilitiesEligibilityTypeGet) | **Get** /api/v1/Subscriptions/{id}/transition-eligibilities/{eligibilityType} | 
[**ApiV1SubscriptionsIdTransitionPost**](SubscriptionsApi.md#ApiV1SubscriptionsIdTransitionPost) | **Post** /api/v1/Subscriptions/{id}/transition | 
[**ApiV1SubscriptionsIdTransitionsGet**](SubscriptionsApi.md#ApiV1SubscriptionsIdTransitionsGet) | **Get** /api/v1/Subscriptions/{id}/transitions | 
[**ApiV1SubscriptionsPost**](SubscriptionsApi.md#ApiV1SubscriptionsPost) | **Post** /api/v1/Subscriptions | 
[**ApiV1SubscriptionsReservedInstanceSubscriptionIdIdGet**](SubscriptionsApi.md#ApiV1SubscriptionsReservedInstanceSubscriptionIdIdGet) | **Get** /api/v1/Subscriptions/{reservedInstance}/subscriptionId/{id} | 
[**ApiV1SubscriptionsReservedInstanceSubscriptionIdIdPost**](SubscriptionsApi.md#ApiV1SubscriptionsReservedInstanceSubscriptionIdIdPost) | **Post** /api/v1/Subscriptions/{reservedInstance}/subscriptionId/{id} | 
[**ApiV1SubscriptionsSubscriptionIdAddonOffersGet**](SubscriptionsApi.md#ApiV1SubscriptionsSubscriptionIdAddonOffersGet) | **Get** /api/v1/Subscriptions/{subscriptionId}/addon-offers | 
[**ApiV1SubscriptionsSubscriptionIdAddonsPost**](SubscriptionsApi.md#ApiV1SubscriptionsSubscriptionIdAddonsPost) | **Post** /api/v1/Subscriptions/{subscriptionId}/addons | 
[**ApiV1SubscriptionsSubscriptionIdConversionsGet**](SubscriptionsApi.md#ApiV1SubscriptionsSubscriptionIdConversionsGet) | **Get** /api/v1/Subscriptions/{subscriptionId}/conversions | 
[**ApiV1SubscriptionsSubscriptionIdConversionsPost**](SubscriptionsApi.md#ApiV1SubscriptionsSubscriptionIdConversionsPost) | **Post** /api/v1/Subscriptions/{subscriptionId}/conversions | 
[**ApiV1SubscriptionsSubscriptionIdTagsDelete**](SubscriptionsApi.md#ApiV1SubscriptionsSubscriptionIdTagsDelete) | **Delete** /api/v1/Subscriptions/{subscriptionId}/tags | 
[**ApiV1SubscriptionsSubscriptionIdTagsGet**](SubscriptionsApi.md#ApiV1SubscriptionsSubscriptionIdTagsGet) | **Get** /api/v1/Subscriptions/{subscriptionId}/tags | 
[**ApiV1SubscriptionsSubscriptionIdTagsPost**](SubscriptionsApi.md#ApiV1SubscriptionsSubscriptionIdTagsPost) | **Post** /api/v1/Subscriptions/{subscriptionId}/tags | 
[**ApiV1SubscriptionsSubscriptionpricetypesGet**](SubscriptionsApi.md#ApiV1SubscriptionsSubscriptionpricetypesGet) | **Get** /api/v1/Subscriptions/subscriptionpricetypes | 

# **ApiV1SubscriptionsGet**
> []Subscription ApiV1SubscriptionsGet(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SubscriptionsApiApiV1SubscriptionsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SubscriptionsApiApiV1SubscriptionsGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **organizationId** | **optional.Int32**|  | 
 **customerTenantId** | **optional.Int32**|  | 
 **publisherId** | **optional.Int32**|  | 
 **refresh** | **optional.Bool**|  | 
 **statuses** | [**optional.Interface of SubscriptionStatus**](.md)|  | 
 **isTrial** | **optional.Bool**|  | 
 **page** | **optional.Int32**|  | 
 **pageSize** | **optional.Int32**|  | 
 **search** | **optional.String**|  | 
 **registeredForReservedInstance** | **optional.Bool**|  | 
 **sortBy** | [**optional.Interface of SubscriptionSortBy**](.md)|  | 
 **sortOrder** | [**optional.Interface of SortOrder**](.md)|  | 

### Return type

[**[]Subscription**](Subscription.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV1SubscriptionsIdActivationlinkGet**
> ActivationLink ApiV1SubscriptionsIdActivationlinkGet(ctx, id)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 

### Return type

[**ActivationLink**](ActivationLink.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV1SubscriptionsIdGet**
> SubscriptionDetailed ApiV1SubscriptionsIdGet(ctx, id)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 

### Return type

[**SubscriptionDetailed**](SubscriptionDetailed.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV1SubscriptionsIdPut**
> SubscriptionDetailed ApiV1SubscriptionsIdPut(ctx, id, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 
 **optional** | ***SubscriptionsApiApiV1SubscriptionsIdPutOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SubscriptionsApiApiV1SubscriptionsIdPutOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of SubscriptionDetailed**](SubscriptionDetailed.md)|  | 

### Return type

[**SubscriptionDetailed**](SubscriptionDetailed.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV1SubscriptionsIdTransitionEligibilitiesEligibilityTypeGet**
> []SubscriptionTransitionEligibility ApiV1SubscriptionsIdTransitionEligibilitiesEligibilityTypeGet(ctx, id, eligibilityType)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 
  **eligibilityType** | **string**|  | 

### Return type

[**[]SubscriptionTransitionEligibility**](SubscriptionTransitionEligibility.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV1SubscriptionsIdTransitionPost**
> SubscriptionTransitionResponse ApiV1SubscriptionsIdTransitionPost(ctx, id, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 
 **optional** | ***SubscriptionsApiApiV1SubscriptionsIdTransitionPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SubscriptionsApiApiV1SubscriptionsIdTransitionPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of SubscriptionTransition**](SubscriptionTransition.md)|  | 

### Return type

[**SubscriptionTransitionResponse**](SubscriptionTransitionResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV1SubscriptionsIdTransitionsGet**
> []SubscriptionTransitionResponse ApiV1SubscriptionsIdTransitionsGet(ctx, id)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 

### Return type

[**[]SubscriptionTransitionResponse**](SubscriptionTransitionResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV1SubscriptionsPost**
> SubscriptionDetailed ApiV1SubscriptionsPost(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SubscriptionsApiApiV1SubscriptionsPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SubscriptionsApiApiV1SubscriptionsPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of SubscriptionDetailed**](SubscriptionDetailed.md)|  | 

### Return type

[**SubscriptionDetailed**](SubscriptionDetailed.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV1SubscriptionsReservedInstanceSubscriptionIdIdGet**
> []bool ApiV1SubscriptionsReservedInstanceSubscriptionIdIdGet(ctx, id, reservedInstance)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 
  **reservedInstance** | **bool**|  | 

### Return type

**[]bool**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV1SubscriptionsReservedInstanceSubscriptionIdIdPost**
> bool ApiV1SubscriptionsReservedInstanceSubscriptionIdIdPost(ctx, id, reservedInstance)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 
  **reservedInstance** | **bool**|  | 

### Return type

**bool**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV1SubscriptionsSubscriptionIdAddonOffersGet**
> []SubscriptionAddOnOffer ApiV1SubscriptionsSubscriptionIdAddonOffersGet(ctx, subscriptionId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **subscriptionId** | **int32**|  | 

### Return type

[**[]SubscriptionAddOnOffer**](SubscriptionAddOnOffer.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV1SubscriptionsSubscriptionIdAddonsPost**
> bool ApiV1SubscriptionsSubscriptionIdAddonsPost(ctx, subscriptionId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **subscriptionId** | **int32**|  | 
 **optional** | ***SubscriptionsApiApiV1SubscriptionsSubscriptionIdAddonsPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SubscriptionsApiApiV1SubscriptionsSubscriptionIdAddonsPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of PostSubscriptionAddOn**](PostSubscriptionAddOn.md)|  | 

### Return type

**bool**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV1SubscriptionsSubscriptionIdConversionsGet**
> []SubscriptionConversion ApiV1SubscriptionsSubscriptionIdConversionsGet(ctx, subscriptionId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **subscriptionId** | **int32**|  | 

### Return type

[**[]SubscriptionConversion**](SubscriptionConversion.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV1SubscriptionsSubscriptionIdConversionsPost**
> SubscriptionDetailed ApiV1SubscriptionsSubscriptionIdConversionsPost(ctx, subscriptionId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **subscriptionId** | **int32**|  | 
 **optional** | ***SubscriptionsApiApiV1SubscriptionsSubscriptionIdConversionsPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SubscriptionsApiApiV1SubscriptionsSubscriptionIdConversionsPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of SubscriptionConversion**](SubscriptionConversion.md)|  | 

### Return type

[**SubscriptionDetailed**](SubscriptionDetailed.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV1SubscriptionsSubscriptionIdTagsDelete**
> ApiV1SubscriptionsSubscriptionIdTagsDelete(ctx, subscriptionId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **subscriptionId** | **int32**|  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV1SubscriptionsSubscriptionIdTagsGet**
> SubscriptionTags ApiV1SubscriptionsSubscriptionIdTagsGet(ctx, subscriptionId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **subscriptionId** | **int32**|  | 

### Return type

[**SubscriptionTags**](SubscriptionTags.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV1SubscriptionsSubscriptionIdTagsPost**
> bool ApiV1SubscriptionsSubscriptionIdTagsPost(ctx, subscriptionId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **subscriptionId** | **string**|  | 
 **optional** | ***SubscriptionsApiApiV1SubscriptionsSubscriptionIdTagsPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SubscriptionsApiApiV1SubscriptionsSubscriptionIdTagsPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of SubscriptionTags**](SubscriptionTags.md)|  | 

### Return type

**bool**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV1SubscriptionsSubscriptionpricetypesGet**
> []ObjectReference ApiV1SubscriptionsSubscriptionpricetypesGet(ctx, )


### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]ObjectReference**](ObjectReference.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

