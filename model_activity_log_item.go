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

type ActivityLogItem struct {
	Entity string `json:"entity,omitempty"`
	Id int32 `json:"id,omitempty"`
	ChangedBy string `json:"changedBy,omitempty"`
	ChangedDate time.Time `json:"changedDate,omitempty"`
	ChangedColumn string `json:"changedColumn,omitempty"`
	ChangedFrom string `json:"changedFrom,omitempty"`
	ChangedTo string `json:"changedTo,omitempty"`
	CustomValues map[string]string `json:"customValues,omitempty"`
	ChangeType *ChangeTypeEnum `json:"changeType,omitempty"`
}