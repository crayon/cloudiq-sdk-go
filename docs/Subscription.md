# Subscription

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [optional] [default to null]
**PublisherSubscriptionId** | **string** |  | [optional] [default to null]
**EntitlementId** | **string** |  | [optional] [default to null]
**Publisher** | [***ObjectReference**](ObjectReference.md) |  | [optional] [default to null]
**Organization** | [***ObjectReference**](ObjectReference.md) |  | [optional] [default to null]
**CustomerTenant** | [***CustomerTenantReference**](CustomerTenantReference.md) |  | [optional] [default to null]
**Product** | [***ProductReference**](ProductReference.md) |  | [optional] [default to null]
**Quantity** | **int32** |  | [optional] [default to null]
**Name** | **string** |  | [optional] [default to null]
**Status** | [***SubscriptionStatus**](SubscriptionStatus.md) |  | [optional] [default to null]
**ProvisionType** | [***ProvisionType**](ProvisionType.md) |  | [optional] [default to null]
**AvailableAddonsCount** | **int32** |  | [optional] [default to null]
**Subscriptions** | [**[]SubscriptionAddOn**](SubscriptionAddOn.md) |  | [optional] [default to null]
**OrderId** | **string** |  | [optional] [default to null]
**CreationDate** | [**time.Time**](time.Time.md) |  | [optional] [default to null]
**BillingCycle** | [***BillingCycleEnum**](BillingCycleEnum.md) |  | [optional] [default to null]
**Markup** | **float64** |  | [optional] [default to null]
**IsTrial** | **bool** |  | [optional] [default to null]
**StartDate** | [**time.Time**](time.Time.md) |  | [optional] [default to null]
**EndDate** | [**time.Time**](time.Time.md) |  | [optional] [default to null]
**PriceType** | [***SubscriptionPriceType**](SubscriptionPriceType.md) |  | [optional] [default to null]
**SalesPrice** | **float64** |  | [optional] [default to null]
**RegisteredForReservedInstance** | **bool** |  | [optional] [default to null]
**InvoiceProfile** | [***ObjectReference**](ObjectReference.md) |  | [optional] [default to null]
**SubscriptionTags** | [***SubscriptionTags**](SubscriptionTags.md) |  | [optional] [default to null]
**AcceptAutoSuspension** | **bool** |  | [optional] [default to null]
**AutoSuspensionDate** | [**time.Time**](time.Time.md) |  | [optional] [default to null]
**PartNumber** | **string** |  | [optional] [default to null]
**TermDuration** | **string** |  | [optional] [default to null]
**AutoRenewEnabled** | **bool** |  | [optional] [default to null]
**CancellationAllowedUntilDate** | [**time.Time**](time.Time.md) |  | [optional] [default to null]
**ScheduledNextTermInstructions** | [***ScheduledNextTermInstructions**](ScheduledNextTermInstructions.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

