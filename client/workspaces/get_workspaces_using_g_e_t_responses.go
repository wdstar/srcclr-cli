// Code generated by go-swagger; DO NOT EDIT.

package workspaces

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/wdstar/srcclr-cli/models"
)

// GetWorkspacesUsingGETReader is a Reader for the GetWorkspacesUsingGET structure.
type GetWorkspacesUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetWorkspacesUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetWorkspacesUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetWorkspacesUsingGETOK creates a GetWorkspacesUsingGETOK with default headers values
func NewGetWorkspacesUsingGETOK() *GetWorkspacesUsingGETOK {
	return &GetWorkspacesUsingGETOK{}
}

/*GetWorkspacesUsingGETOK handles this case with default header values.

You have successfully submitted your request.
*/
type GetWorkspacesUsingGETOK struct {
	Payload *models.PagedResourcesWorkspace
}

func (o *GetWorkspacesUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /v3/workspaces][%d] getWorkspacesUsingGETOK  %+v", 200, o.Payload)
}

func (o *GetWorkspacesUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PagedResourcesWorkspace)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}