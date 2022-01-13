/*
 * Crayon Group Customer API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type AgreementIdentityReferenceDto struct {
	AgreementNumber string `json:"agreementNumber,omitempty"`
	CommitmentLevel float64 `json:"commitmentLevel,omitempty"`
	IsCustomCommitment bool `json:"isCustomCommitment,omitempty"`
	Id int32 `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}
