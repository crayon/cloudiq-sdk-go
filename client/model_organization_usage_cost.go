/*
 * Crayon Group Customer API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type OrganizationUsageCost struct {
	Supplier string `json:"supplier,omitempty"`
	AccountId int32 `json:"accountId,omitempty"`
	AccountName string `json:"accountName,omitempty"`
	SubscriptionName string `json:"subscriptionName,omitempty"`
	SubscriptionId string `json:"subscriptionId,omitempty"`
	Amount float64 `json:"amount,omitempty"`
	CurrencyCode string `json:"currencyCode,omitempty"`
}
