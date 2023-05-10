// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/hunjixin/brightbird/models"
)

// RunJobImmediatelyReader is a Reader for the RunJobImmediately structure.
type RunJobImmediatelyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RunJobImmediatelyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRunJobImmediatelyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 503:
		result := NewRunJobImmediatelyServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewRunJobImmediatelyOK creates a RunJobImmediatelyOK with default headers values
func NewRunJobImmediatelyOK() *RunJobImmediatelyOK {
	return &RunJobImmediatelyOK{}
}

/*
RunJobImmediatelyOK describes a response with status code 200, with default header values.

RunJobImmediatelyOK run job immediately o k
*/
type RunJobImmediatelyOK struct {
}

// IsSuccess returns true when this run job immediately o k response has a 2xx status code
func (o *RunJobImmediatelyOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this run job immediately o k response has a 3xx status code
func (o *RunJobImmediatelyOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this run job immediately o k response has a 4xx status code
func (o *RunJobImmediatelyOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this run job immediately o k response has a 5xx status code
func (o *RunJobImmediatelyOK) IsServerError() bool {
	return false
}

// IsCode returns true when this run job immediately o k response a status code equal to that given
func (o *RunJobImmediatelyOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the run job immediately o k response
func (o *RunJobImmediatelyOK) Code() int {
	return 200
}

func (o *RunJobImmediatelyOK) Error() string {
	return fmt.Sprintf("[POST /run/{jobid}][%d] runJobImmediatelyOK ", 200)
}

func (o *RunJobImmediatelyOK) String() string {
	return fmt.Sprintf("[POST /run/{jobid}][%d] runJobImmediatelyOK ", 200)
}

func (o *RunJobImmediatelyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRunJobImmediatelyServiceUnavailable creates a RunJobImmediatelyServiceUnavailable with default headers values
func NewRunJobImmediatelyServiceUnavailable() *RunJobImmediatelyServiceUnavailable {
	return &RunJobImmediatelyServiceUnavailable{}
}

/*
RunJobImmediatelyServiceUnavailable describes a response with status code 503, with default header values.

apiError
*/
type RunJobImmediatelyServiceUnavailable struct {
	Payload *models.APIError
}

// IsSuccess returns true when this run job immediately service unavailable response has a 2xx status code
func (o *RunJobImmediatelyServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this run job immediately service unavailable response has a 3xx status code
func (o *RunJobImmediatelyServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this run job immediately service unavailable response has a 4xx status code
func (o *RunJobImmediatelyServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this run job immediately service unavailable response has a 5xx status code
func (o *RunJobImmediatelyServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this run job immediately service unavailable response a status code equal to that given
func (o *RunJobImmediatelyServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

// Code gets the status code for the run job immediately service unavailable response
func (o *RunJobImmediatelyServiceUnavailable) Code() int {
	return 503
}

func (o *RunJobImmediatelyServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /run/{jobid}][%d] runJobImmediatelyServiceUnavailable  %+v", 503, o.Payload)
}

func (o *RunJobImmediatelyServiceUnavailable) String() string {
	return fmt.Sprintf("[POST /run/{jobid}][%d] runJobImmediatelyServiceUnavailable  %+v", 503, o.Payload)
}

func (o *RunJobImmediatelyServiceUnavailable) GetPayload() *models.APIError {
	return o.Payload
}

func (o *RunJobImmediatelyServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
