/*
 * Crayon Group Customer API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type SubscriptionPricing struct {
	CurrencyCode string `json:"currencyCode,omitempty"`
	PurchasePrice float64 `json:"purchasePrice,omitempty"`
	RecommendedRetailPrice float64 `json:"recommendedRetailPrice,omitempty"`
}