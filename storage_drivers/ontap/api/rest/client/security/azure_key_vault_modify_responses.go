// Code generated by go-swagger; DO NOT EDIT.

package security

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// AzureKeyVaultModifyReader is a Reader for the AzureKeyVaultModify structure.
type AzureKeyVaultModifyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AzureKeyVaultModifyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAzureKeyVaultModifyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewAzureKeyVaultModifyAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAzureKeyVaultModifyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAzureKeyVaultModifyOK creates a AzureKeyVaultModifyOK with default headers values
func NewAzureKeyVaultModifyOK() *AzureKeyVaultModifyOK {
	return &AzureKeyVaultModifyOK{}
}

/* AzureKeyVaultModifyOK describes a response with status code 200, with default header values.

OK
*/
type AzureKeyVaultModifyOK struct {
}

func (o *AzureKeyVaultModifyOK) Error() string {
	return fmt.Sprintf("[PATCH /security/azure-key-vaults/{uuid}][%d] azureKeyVaultModifyOK ", 200)
}

func (o *AzureKeyVaultModifyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAzureKeyVaultModifyAccepted creates a AzureKeyVaultModifyAccepted with default headers values
func NewAzureKeyVaultModifyAccepted() *AzureKeyVaultModifyAccepted {
	return &AzureKeyVaultModifyAccepted{}
}

/* AzureKeyVaultModifyAccepted describes a response with status code 202, with default header values.

Accepted
*/
type AzureKeyVaultModifyAccepted struct {
}

func (o *AzureKeyVaultModifyAccepted) Error() string {
	return fmt.Sprintf("[PATCH /security/azure-key-vaults/{uuid}][%d] azureKeyVaultModifyAccepted ", 202)
}

func (o *AzureKeyVaultModifyAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAzureKeyVaultModifyDefault creates a AzureKeyVaultModifyDefault with default headers values
func NewAzureKeyVaultModifyDefault(code int) *AzureKeyVaultModifyDefault {
	return &AzureKeyVaultModifyDefault{
		_statusCode: code,
	}
}

/* AzureKeyVaultModifyDefault describes a response with status code -1, with default header values.

 ONTAP Error Response Codes
| Error Code | Description |
| ---------- | ----------- |
| 65537120 | Azure Key Vault is not configured for the given SVM. |
| 65537504 | Internal error. Failed to store configuration in internal database. |
| 65537517 | The field \"client_secret\" must be specified. |
| 65537540 | Invalid client secret. |
| 65537541 | No inputs were provided for the patch request. |
| 65537547 | One or more volume encryption keys for encrypted volumes of this data SVM are stored in the key manager configured for the admin SVM. Use the REST API POST method to migrate this data SVM's keys from the admin SVM's key manager to this data SVM's key manager before running the rekey operation. |
| 65537573 | Invalid client certificate. |
| 65537577 | The AKV certificate authentication method cannot be configured for the given SVM as not all nodes in the cluster support the AKV certificate authentication. |

*/
type AzureKeyVaultModifyDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the azure key vault modify default response
func (o *AzureKeyVaultModifyDefault) Code() int {
	return o._statusCode
}

func (o *AzureKeyVaultModifyDefault) Error() string {
	return fmt.Sprintf("[PATCH /security/azure-key-vaults/{uuid}][%d] azure_key_vault_modify default  %+v", o._statusCode, o.Payload)
}
func (o *AzureKeyVaultModifyDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *AzureKeyVaultModifyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
