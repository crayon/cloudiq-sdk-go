/*
 * Crayon Group Customer API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type ProductRow struct {
	Id int32 `json:"id,omitempty"`
	ProductContainerId int32 `json:"productContainerId,omitempty"`
	Quantity int32 `json:"quantity,omitempty"`
	Comment string `json:"comment,omitempty"`
	UsageCountryCode string `json:"usageCountryCode,omitempty"`
	SalesUnitPrice *Price `json:"salesUnitPrice,omitempty"`
	AlternativeSalesUnitPrice *Price `json:"alternativeSalesUnitPrice,omitempty"`
	Publisher *ObjectReference `json:"publisher,omitempty"`
	Program *ObjectReference `json:"program,omitempty"`
	Agreement *AgreementIdentityReference `json:"agreement,omitempty"`
	Product *ProductReference `json:"product,omitempty"`
	User *ProductRowUser `json:"user,omitempty"`
	ProductVariant *ObjectReference `json:"productVariant,omitempty"`
	InvoiceProfile *ObjectReference `json:"invoiceProfile,omitempty"`
	Grouping *Grouping `json:"grouping,omitempty"`
	Issues []ProductContainerIssue `json:"issues,omitempty"`
	OfferingType string `json:"offeringType,omitempty"`
	PriceCalculation *PriceCalculationType `json:"priceCalculation,omitempty"`
	InvoiceReference string `json:"invoiceReference,omitempty"`
	CustomerReference string `json:"customerReference,omitempty"`
	SalesPricePerAlternativeUnit float64 `json:"salesPricePerAlternativeUnit,omitempty"`
	LevelValue float64 `json:"levelValue,omitempty"`
}
