/*
 * Crayon Group Customer API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type AzureSubscription struct {
	Id int32 `json:"id,omitempty"`
	AzurePlanId int32 `json:"azurePlanId,omitempty"`
	PublisherSubscriptionId string `json:"publisherSubscriptionId,omitempty"`
	FriendlyName string `json:"friendlyName,omitempty"`
	Status string `json:"status,omitempty"`
	InvoiceProfile *ObjectReference `json:"invoiceProfile,omitempty"`
	Tags *AzureSubscriptionTags `json:"tags,omitempty"`
}
