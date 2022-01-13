# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV1AgreementsGet**](AgreementsApi.md#ApiV1AgreementsGet) | **Get** /api/v1/Agreements | 

# **ApiV1AgreementsGet**
> []Agreement ApiV1AgreementsGet(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AgreementsApiApiV1AgreementsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AgreementsApiApiV1AgreementsGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **organizationId** | **optional.Int32**|  | 
 **organizationIds** | [**optional.Interface of []int32**](int32.md)|  | 
 **pricelistIds** | [**optional.Interface of []int32**](int32.md)|  | 
 **status** | [**optional.Interface of AgreementStatus**](.md)|  | 
 **agreementTypes** | [**optional.Interface of []AgreementType**](AgreementType.md)|  | 
 **publisherIds** | [**optional.Interface of []int32**](int32.md)|  | 
 **programIds** | [**optional.Interface of []int32**](int32.md)|  | 
 **searchDate** | **optional.Time**|  | 
 **agreementIds** | [**optional.Interface of []int32**](int32.md)|  | 
 **salesPriceCurrency** | **optional.String**|  | 
 **termRequired** | **optional.Bool**|  | 
 **publisherId** | **optional.Int32**|  | 
 **endDateFrom** | **optional.Time**|  | 
 **endDateTo** | **optional.Time**|  | 
 **page** | **optional.Int32**|  | 
 **pageSize** | **optional.Int32**|  | 
 **search** | **optional.String**|  | 

### Return type

[**[]Agreement**](Agreement.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

