/*
 * Crayon Group Customer API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type PostSubscriptionAddOn struct {
	Quantity int32 `json:"quantity,omitempty"`
	OfferId string `json:"offerId,omitempty"`
	SubscriptionTags *SubscriptionTags `json:"subscriptionTags,omitempty"`
}
