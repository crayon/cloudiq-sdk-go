/*
 * Crayon Group Customer API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type Product struct {
	ProductId string `json:"productId,omitempty"`
	SkuId string `json:"skuId,omitempty"`
	AvailabilityId string `json:"availabilityId,omitempty"`
	BillingCycle string `json:"billingCycle,omitempty"`
	TermDuration string `json:"termDuration,omitempty"`
}