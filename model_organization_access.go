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

type OrganizationAccess struct {
	Id int32 `json:"id,omitempty"`
	AllAgreements bool `json:"allAgreements,omitempty"`
	Agreements []AgreementAccess `json:"agreements,omitempty"`
	Organization *Organization `json:"organization,omitempty"`
	User *UserProfile `json:"user,omitempty"`
	Role *OrganizationAccessRole `json:"role,omitempty"`
	CrayonCompanyName string `json:"crayonCompanyName,omitempty"`
	Timestamp time.Time `json:"timestamp,omitempty"`
}
