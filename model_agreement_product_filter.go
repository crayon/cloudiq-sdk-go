/*
 * Crayon Group Customer API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger
import (
	"time"
)

type AgreementProductFilter struct {
	AgreementTypeIds []AgreementType `json:"agreementTypeIds,omitempty"`
	Page int32 `json:"page,omitempty"`
	PageSize int32 `json:"pageSize,omitempty"`
	Search string `json:"search,omitempty"`
	PriceListId int32 `json:"priceListId,omitempty"`
	OrganizationId int32 `json:"organizationId,omitempty"`
	CustomerTenantId int32 `json:"customerTenantId,omitempty"`
	AgreementId int32 `json:"agreementId,omitempty"`
	IsTrial bool `json:"isTrial,omitempty"`
	IgnoreTermBillingCycleCombinationsCheck bool `json:"ignoreTermBillingCycleCombinationsCheck,omitempty"`
	AgreementIds []int32 `json:"agreementIds,omitempty"`
	SearchDate time.Time `json:"searchDate,omitempty"`
	Include *AgreementProductsSubFilter `json:"include,omitempty"`
	Exclude *AgreementProductsSubFilter `json:"exclude,omitempty"`
	SortKey string `json:"sortKey,omitempty"`
	IncludeProductInformation bool `json:"includeProductInformation,omitempty"`
	SortOrder *SortOrder `json:"sortOrder,omitempty"`
}
