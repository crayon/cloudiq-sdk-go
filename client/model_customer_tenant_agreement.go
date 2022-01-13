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

type CustomerTenantAgreement struct {
	FirstName string `json:"firstName,omitempty"`
	LastName string `json:"lastName,omitempty"`
	PhoneNumber string `json:"phoneNumber,omitempty"`
	Email string `json:"email,omitempty"`
	SameAsPrimaryContact bool `json:"sameAsPrimaryContact,omitempty"`
	DateAgreed time.Time `json:"dateAgreed,omitempty"`
	Accepted bool `json:"accepted,omitempty"`
	AgreementType *AgreementTypeConsent `json:"agreementType,omitempty"`
}