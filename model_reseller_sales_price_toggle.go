/*
 * Crayon Group Customer API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type ResellerSalesPriceToggle struct {
	Toggle bool `json:"toggle,omitempty"`
	ObjectId int32 `json:"objectId,omitempty"`
	ObjectType *ResellerSalesPriceObjectType `json:"objectType,omitempty"`
}