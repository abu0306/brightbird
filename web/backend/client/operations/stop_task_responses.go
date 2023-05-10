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

// StopTaskReader is a Reader for the StopTask structure.
type StopTaskReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StopTaskReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStopTaskOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 503:
		result := NewStopTaskServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewStopTaskOK creates a StopTaskOK with default headers values
func NewStopTaskOK() *StopTaskOK {
	return &StopTaskOK{}
}

/*
StopTaskOK describes a response with status code 200, with default header values.

StopTaskOK stop task o k
*/
type StopTaskOK struct {
}

// IsSuccess returns true when this stop task o k response has a 2xx status code
func (o *StopTaskOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this stop task o k response has a 3xx status code
func (o *StopTaskOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stop task o k response has a 4xx status code
func (o *StopTaskOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this stop task o k response has a 5xx status code
func (o *StopTaskOK) IsServerError() bool {
	return false
}

// IsCode returns true when this stop task o k response a status code equal to that given
func (o *StopTaskOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the stop task o k response
func (o *StopTaskOK) Code() int {
	return 200
}

func (o *StopTaskOK) Error() string {
	return fmt.Sprintf("[DELETE /task/stop/{id}][%d] stopTaskOK ", 200)
}

func (o *StopTaskOK) String() string {
	return fmt.Sprintf("[DELETE /task/stop/{id}][%d] stopTaskOK ", 200)
}

func (o *StopTaskOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewStopTaskServiceUnavailable creates a StopTaskServiceUnavailable with default headers values
func NewStopTaskServiceUnavailable() *StopTaskServiceUnavailable {
	return &StopTaskServiceUnavailable{}
}

/*
StopTaskServiceUnavailable describes a response with status code 503, with default header values.

apiError
*/
type StopTaskServiceUnavailable struct {
	Payload *models.APIError
}

// IsSuccess returns true when this stop task service unavailable response has a 2xx status code
func (o *StopTaskServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stop task service unavailable response has a 3xx status code
func (o *StopTaskServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stop task service unavailable response has a 4xx status code
func (o *StopTaskServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this stop task service unavailable response has a 5xx status code
func (o *StopTaskServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this stop task service unavailable response a status code equal to that given
func (o *StopTaskServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

// Code gets the status code for the stop task service unavailable response
func (o *StopTaskServiceUnavailable) Code() int {
	return 503
}

func (o *StopTaskServiceUnavailable) Error() string {
	return fmt.Sprintf("[DELETE /task/stop/{id}][%d] stopTaskServiceUnavailable  %+v", 503, o.Payload)
}

func (o *StopTaskServiceUnavailable) String() string {
	return fmt.Sprintf("[DELETE /task/stop/{id}][%d] stopTaskServiceUnavailable  %+v", 503, o.Payload)
}

func (o *StopTaskServiceUnavailable) GetPayload() *models.APIError {
	return o.Payload
}

func (o *StopTaskServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
