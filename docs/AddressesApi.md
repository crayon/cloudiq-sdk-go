# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV1OrganizationsOrganizationIdAddressesGet**](AddressesApi.md#ApiV1OrganizationsOrganizationIdAddressesGet) | **Get** /api/v1/organizations/{organizationId}/Addresses | 
[**ApiV1OrganizationsOrganizationIdAddressesIdGet**](AddressesApi.md#ApiV1OrganizationsOrganizationIdAddressesIdGet) | **Get** /api/v1/organizations/{organizationId}/Addresses/{id} | 

# **ApiV1OrganizationsOrganizationIdAddressesGet**
> []Address ApiV1OrganizationsOrganizationIdAddressesGet(ctx, organizationId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **organizationId** | **int32**|  | 
 **optional** | ***AddressesApiApiV1OrganizationsOrganizationIdAddressesGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AddressesApiApiV1OrganizationsOrganizationIdAddressesGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | [**optional.Interface of AddressType**](.md)|  | 

### Return type

[**[]Address**](Address.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV1OrganizationsOrganizationIdAddressesIdGet**
> Address ApiV1OrganizationsOrganizationIdAddressesIdGet(ctx, organizationId, id)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **organizationId** | **int32**|  | 
  **id** | **int64**|  | 

### Return type

[**Address**](Address.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

