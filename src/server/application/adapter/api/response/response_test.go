package response

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

const (
	contentTypeString = "Content-Type"
)

func TestSuccessSend(t *testing.T) {
	result := map[string]string{"key": "value"}
	success := NewSuccess(result, http.StatusOK)

	recorder := httptest.NewRecorder()
	err := success.Send(recorder)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	expectedStatus := http.StatusOK
	if recorder.Code != expectedStatus {
		t.Errorf("expected status code %d, but got %d", expectedStatus, recorder.Code)
	}

	expectedContentType := "application/json"
	if recorder.Header().Get(contentTypeString) != expectedContentType {
		t.Errorf("expected Content-Type %s, but got %s", expectedContentType, recorder.Header().Get(contentTypeString))
	}

	expectedBody := recorder.Body.String()
	if expectedBody != "{\"key\":\"value\"}\n" {
		t.Errorf("expected Body %s, but got %s", expectedBody, "{\"key\":\"value\"}\n")
	}

}

func TestErrorSend(t *testing.T) {
	errMsg := "parameter invalid"
	err := NewError(ErrParameterInvalid, http.StatusBadRequest)

	rr := httptest.NewRecorder()
	err.Send(rr)

	if rr.Code != http.StatusBadRequest {
		t.Errorf("expected HTTP status code %d, got %d", http.StatusBadRequest, rr.Code)
	}

	var respBody Error
	if err := json.NewDecoder(rr.Body).Decode(&respBody); err != nil {
		t.Errorf("error decoding response body: %s", err)
	}

	if len(respBody.Errors) != 1 || respBody.Errors[0] != errMsg {
		t.Errorf("expected response body with errors [%s], got [%v]", errMsg, respBody.Errors)
	}

	if contentType := rr.Header().Get(contentTypeString); contentType != "application/json" {
		t.Errorf("expected Content-Type header 'application/json', got '%s'", contentType)
	}
}

func TestNewErrorMessage(t *testing.T) {
	expectedStatus := 400
	expectedMessages := []string{"Error message 1", "Error message 2"}

	err := NewErrorMessage(expectedMessages, expectedStatus)

	if err.statusCode != expectedStatus {
		t.Errorf("Expected status code %d, but got %d", expectedStatus, err.statusCode)
	}

	if !reflect.DeepEqual(err.Errors, expectedMessages) {
		t.Errorf("Expected error messages %v, but got %v", expectedMessages, err.Errors)
	}
}
