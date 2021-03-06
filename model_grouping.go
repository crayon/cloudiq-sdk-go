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

type Grouping struct {
	Id int32 `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	InvoiceProfile *ObjectReference `json:"invoiceProfile,omitempty"`
	Organization *ObjectReference `json:"organization,omitempty"`
	CreatedDate time.Time `json:"createdDate,omitempty"`
	ModifiedDate time.Time `json:"modifiedDate,omitempty"`
	IsDisabled bool `json:"isDisabled,omitempty"`
	IsRemoved bool `json:"isRemoved,omitempty"`
	InvoiceReference string `json:"invoiceReference,omitempty"`
}
