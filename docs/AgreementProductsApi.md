# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV1AgreementProductsFileXlsxPost**](AgreementProductsApi.md#ApiV1AgreementProductsFileXlsxPost) | **Post** /api/v1/AgreementProducts/file/xlsx | 
[**ApiV1AgreementProductsGet**](AgreementProductsApi.md#ApiV1AgreementProductsGet) | **Get** /api/v1/AgreementProducts | 
[**ApiV1AgreementProductsPartNumberSupportedbillingcyclesGet**](AgreementProductsApi.md#ApiV1AgreementProductsPartNumberSupportedbillingcyclesGet) | **Get** /api/v1/AgreementProducts/{partNumber}/supportedbillingcycles | 
[**ApiV1AgreementProductsPost**](AgreementProductsApi.md#ApiV1AgreementProductsPost) | **Post** /api/v1/AgreementProducts | 

# **ApiV1AgreementProductsFileXlsxPost**
> ApiV1AgreementProductsFileXlsxPost(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AgreementProductsApiApiV1AgreementProductsFileXlsxPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AgreementProductsApiApiV1AgreementProductsFileXlsxPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of AgreementProductFilter**](AgreementProductFilter.md)|  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV1AgreementProductsGet**
> []AgreementProduct ApiV1AgreementProductsGet(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AgreementProductsApiApiV1AgreementProductsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AgreementProductsApiApiV1AgreementProductsGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **agreementTypeIds** | [**optional.Interface of []AgreementType**](AgreementType.md)|  | 
 **page** | **optional.Int32**|  | 
 **pageSize** | **optional.Int32**|  | 
 **search** | **optional.String**|  | 
 **priceListId** | **optional.Int32**|  | 
 **organizationId** | **optional.Int32**|  | 
 **customerTenantId** | **optional.Int32**|  | 
 **agreementId** | **optional.Int32**|  | 
 **isTrial** | **optional.Bool**|  | 
 **agreementIds** | [**optional.Interface of []int32**](int32.md)|  | 
 **searchDate** | **optional.Time**|  | 
 **includePartNumbers** | [**optional.Interface of []string**](string.md)|  | 
 **includePublisherIds** | [**optional.Interface of []int32**](int32.md)|  | 
 **includePublisherNames** | [**optional.Interface of []string**](string.md)|  | 
 **includePoolNames** | [**optional.Interface of []string**](string.md)|  | 
 **includeOperatingSystemNames** | [**optional.Interface of []string**](string.md)|  | 
 **includeLevelNames** | [**optional.Interface of []string**](string.md)|  | 
 **includeLanguageNames** | [**optional.Interface of []string**](string.md)|  | 
 **includeLicenseAgreementTypeNames** | [**optional.Interface of []string**](string.md)|  | 
 **includeLicenseTypeNames** | [**optional.Interface of []string**](string.md)|  | 
 **includeProductFamilyNames** | [**optional.Interface of []string**](string.md)|  | 
 **includeProductTypeNames** | [**optional.Interface of []string**](string.md)|  | 
 **includeProgramNames** | [**optional.Interface of []string**](string.md)|  | 
 **includeOfferingNames** | [**optional.Interface of []string**](string.md)|  | 
 **includePurchasePeriodNames** | [**optional.Interface of []string**](string.md)|  | 
 **includePurchaseUnitNames** | [**optional.Interface of []string**](string.md)|  | 
 **includeVersionNames** | [**optional.Interface of []string**](string.md)|  | 
 **includeRegionNames** | [**optional.Interface of []string**](string.md)|  | 
 **includeProductCategoryNames** | [**optional.Interface of []string**](string.md)|  | 
 **includeCatalogNames** | [**optional.Interface of []string**](string.md)|  | 
 **includeBillingCycles** | [**optional.Interface of []string**](string.md)|  | 
 **excludePartNumbers** | [**optional.Interface of []string**](string.md)|  | 
 **excludePublisherIds** | [**optional.Interface of []int32**](int32.md)|  | 
 **excludePublisherNames** | [**optional.Interface of []string**](string.md)|  | 
 **excludePoolNames** | [**optional.Interface of []string**](string.md)|  | 
 **excludeOperatingSystemNames** | [**optional.Interface of []string**](string.md)|  | 
 **excludeLevelNames** | [**optional.Interface of []string**](string.md)|  | 
 **excludeLanguageNames** | [**optional.Interface of []string**](string.md)|  | 
 **excludeLicenseAgreementTypeNames** | [**optional.Interface of []string**](string.md)|  | 
 **excludeLicenseTypeNames** | [**optional.Interface of []string**](string.md)|  | 
 **excludeProductFamilyNames** | [**optional.Interface of []string**](string.md)|  | 
 **excludeProductTypeNames** | [**optional.Interface of []string**](string.md)|  | 
 **excludeProgramNames** | [**optional.Interface of []string**](string.md)|  | 
 **excludeOfferingNames** | [**optional.Interface of []string**](string.md)|  | 
 **excludePurchasePeriodNames** | [**optional.Interface of []string**](string.md)|  | 
 **excludePurchaseUnitNames** | [**optional.Interface of []string**](string.md)|  | 
 **excludeVersionNames** | [**optional.Interface of []string**](string.md)|  | 
 **excludeRegionNames** | [**optional.Interface of []string**](string.md)|  | 
 **excludeProductCategoryNames** | [**optional.Interface of []string**](string.md)|  | 
 **excludeCatalogNames** | [**optional.Interface of []string**](string.md)|  | 
 **excludeBillingCycles** | [**optional.Interface of []string**](string.md)|  | 
 **sortKey** | **optional.String**|  | 
 **includeProductInformation** | **optional.Bool**|  | 
 **sortOrder** | [**optional.Interface of SortOrder**](.md)|  | 

### Return type

[**[]AgreementProduct**](AgreementProduct.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV1AgreementProductsPartNumberSupportedbillingcyclesGet**
> []BillingCycleEnum ApiV1AgreementProductsPartNumberSupportedbillingcyclesGet(ctx, partNumber, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **partNumber** | **string**|  | 
 **optional** | ***AgreementProductsApiApiV1AgreementProductsPartNumberSupportedbillingcyclesGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AgreementProductsApiApiV1AgreementProductsPartNumberSupportedbillingcyclesGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **resellerCustomerId** | **optional.Int32**|  | 

### Return type

[**[]BillingCycleEnum**](BillingCycleEnum.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV1AgreementProductsPost**
> []AgreementProduct ApiV1AgreementProductsPost(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AgreementProductsApiApiV1AgreementProductsPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AgreementProductsApiApiV1AgreementProductsPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of AgreementProductFilter**](AgreementProductFilter.md)|  | 

### Return type

[**[]AgreementProduct**](AgreementProduct.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

