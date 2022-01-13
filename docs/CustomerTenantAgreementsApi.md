# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV1CustomertenantsCustomerTenantIdAgreementsGet**](CustomerTenantAgreementsApi.md#ApiV1CustomertenantsCustomerTenantIdAgreementsGet) | **Get** /api/v1/customertenants/{customerTenantId}/agreements | 
[**ApiV1CustomertenantsCustomerTenantIdAgreementsPost**](CustomerTenantAgreementsApi.md#ApiV1CustomertenantsCustomerTenantIdAgreementsPost) | **Post** /api/v1/customertenants/{customerTenantId}/agreements | 

# **ApiV1CustomertenantsCustomerTenantIdAgreementsGet**
> []ServiceAccountAgreement ApiV1CustomertenantsCustomerTenantIdAgreementsGet(ctx, customerTenantId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **customerTenantId** | **int32**|  | 
 **optional** | ***CustomerTenantAgreementsApiApiV1CustomertenantsCustomerTenantIdAgreementsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerTenantAgreementsApiApiV1CustomertenantsCustomerTenantIdAgreementsGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **agreementTypeConsent** | [**optional.Interface of AgreementTypeConsent**](.md)|  | 

### Return type

[**[]ServiceAccountAgreement**](ServiceAccountAgreement.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV1CustomertenantsCustomerTenantIdAgreementsPost**
> ServiceAccountAgreement ApiV1CustomertenantsCustomerTenantIdAgreementsPost(ctx, customerTenantId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **customerTenantId** | **int32**|  | 
 **optional** | ***CustomerTenantAgreementsApiApiV1CustomertenantsCustomerTenantIdAgreementsPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerTenantAgreementsApiApiV1CustomertenantsCustomerTenantIdAgreementsPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of ServiceAccountAgreement**](ServiceAccountAgreement.md)|  | 

### Return type

[**ServiceAccountAgreement**](ServiceAccountAgreement.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

