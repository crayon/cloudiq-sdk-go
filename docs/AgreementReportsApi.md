# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV1AgreementReportsAgreementIdPut**](AgreementReportsApi.md#ApiV1AgreementReportsAgreementIdPut) | **Put** /api/v1/AgreementReports/{agreementId} | 
[**ApiV1AgreementReportsProductContainerIdGet**](AgreementReportsApi.md#ApiV1AgreementReportsProductContainerIdGet) | **Get** /api/v1/AgreementReports/{productContainerId} | 

# **ApiV1AgreementReportsAgreementIdPut**
> AgreementReport ApiV1AgreementReportsAgreementIdPut(ctx, agreementId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **agreementId** | **int32**|  | 
 **optional** | ***AgreementReportsApiApiV1AgreementReportsAgreementIdPutOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AgreementReportsApiApiV1AgreementReportsAgreementIdPutOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of AgreementReport**](AgreementReport.md)|  | 

### Return type

[**AgreementReport**](AgreementReport.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV1AgreementReportsProductContainerIdGet**
> []AgreementReport ApiV1AgreementReportsProductContainerIdGet(ctx, productContainerId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **productContainerId** | **int32**|  | 

### Return type

[**[]AgreementReport**](AgreementReport.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

