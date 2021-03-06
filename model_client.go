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

type Client struct {
	ClientId string `json:"clientId,omitempty"`
	ClientName string `json:"clientName,omitempty"`
	ClientUri string `json:"clientUri,omitempty"`
	Enabled bool `json:"enabled,omitempty"`
	ClientSecrets []Secret `json:"clientSecrets,omitempty"`
	RedirectUris []string `json:"redirectUris,omitempty"`
	TimeStamp time.Time `json:"timeStamp,omitempty"`
	Flow *Flow `json:"flow,omitempty"`
}
