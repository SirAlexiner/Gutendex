package main

import (
	"gutendex_api/internal/handlers"
	"gutendex_api/internal/handlers/endpoint/library"
	"gutendex_api/internal/handlers/endpoint/statistics"
	"gutendex_api/internal/utils/constants/Endpoints"
	"gutendex_api/internal/utils/constants/Paths"
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestLibraryGet confirms that the Root Endpoint returns Status I'm a Teapot for All Request.
func TestRoot(t *testing.T) {
	// Create a request to your endpoint with the GET method
	req, err := http.NewRequest("GET", Paths.Root, nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.EmptyHandler)

	// Serve the request to the handler
	handler.ServeHTTP(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusTeapot {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusTeapot)
	}

	req, err = http.NewRequest("POST", Paths.Root, nil)
	if err != nil {
		t.Fatal(err)
	}

	// Serve the request to the handler
	handler.ServeHTTP(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusTeapot {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusTeapot)
	}

	req, err = http.NewRequest("PUT", Paths.Root, nil)
	if err != nil {
		t.Fatal(err)
	}

	// Serve the request to the handler
	handler.ServeHTTP(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusTeapot {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusTeapot)
	}
}

// TestLibraryGet confirms that the Library Endpoint returns Status OK for Get Request.
func TestLibraryGet(t *testing.T) {
	// Create a request to your endpoint with the GET method
	req, err := http.NewRequest("GET", Endpoints.Library, nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(library.LibHandler)

	// Serve the request to the handler
	handler.ServeHTTP(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

// TestLibraryGetLanguage confirms that the Library Endpoint returns Status OK for Get Request with language param.
func TestLibraryGetLanguage(t *testing.T) {
	// Create a request to your endpoint with the GET method
	req, err := http.NewRequest("GET", Endpoints.Library+"?languages=no", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(library.LibHandler)

	// Serve the request to the handler
	handler.ServeHTTP(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

// TestLibraryGetLanguage confirms that the Library Endpoint returns Status OK for Get Request with multiple language param.
func TestLibraryGetLanguages(t *testing.T) {
	// Create a request to your endpoint with the GET method
	req, err := http.NewRequest("GET", Endpoints.Library+"?languages=no,es", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(library.LibHandler)

	// Serve the request to the handler
	handler.ServeHTTP(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

// TestLibraryGetLanguageWrong confirms that the Library Endpoint returns Status Bad Request for Get Request with wrongful language param.
func TestLibraryGetLanguageWrong(t *testing.T) {
	// Create a request to your endpoint with the GET method
	req, err := http.NewRequest("GET", Endpoints.Library+"?languages=nog", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(library.LibHandler)

	// Serve the request to the handler
	handler.ServeHTTP(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}
}

// TestLibraryGetLanguagesWrong confirms that the Library Endpoint returns Status Bad Request for Get Request with the same language param.
func TestLibraryGetLanguagesWrong(t *testing.T) {
	// Create a request to your endpoint with the GET method
	req, err := http.NewRequest("GET", Endpoints.Library+"?languages=no,no", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(library.LibHandler)

	// Serve the request to the handler
	handler.ServeHTTP(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}
}

// TestLibraryGetpage confirms that the Library Endpoint returns Status OK for Get Request with page param.
func TestLibraryGetpage(t *testing.T) {
	// Create a request to your endpoint with the GET method
	req, err := http.NewRequest("GET", Endpoints.Library+"?page=2", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(library.LibHandler)

	// Serve the request to the handler
	handler.ServeHTTP(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

// TestLibraryGetpageWrong confirms that the Library Endpoint returns Status Bad Request for Get Request with wrong page param.
func TestLibraryGetpageWrong(t *testing.T) {
	// Create a request to your endpoint with the GET method
	req, err := http.NewRequest("GET", Endpoints.Library+"?page=one", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(library.LibHandler)

	// Serve the request to the handler
	handler.ServeHTTP(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}
}

// TestLibraryGetSearch confirms that the Library Endpoint returns Status OK for Get Request with search param.
func TestLibraryGetSearch(t *testing.T) {
	// Create a request to your endpoint with the GET method
	req, err := http.NewRequest("GET", Endpoints.Library+"?search=hamsun", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(library.LibHandler)

	// Serve the request to the handler
	handler.ServeHTTP(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

// TestLibraryMethodNotAllowed confirms that the Library Endpoint returns Status Not Implemented for Methods other than GET.
func TestLibraryMethodNotAllowed(t *testing.T) {
	// Create a request to your endpoint with a method other than GET
	req, err := http.NewRequest("POST", Endpoints.Library, nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(library.LibHandler)

	// Serve the request to the handler
	handler.ServeHTTP(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusNotImplemented {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusNotImplemented)
	}

	req, err = http.NewRequest("PUT", Endpoints.Library, nil)
	if err != nil {
		t.Fatal(err)
	}

	// Serve the request to the handler
	handler.ServeHTTP(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusNotImplemented {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusNotImplemented)
	}
}

// TestBookCountGetLanguage confirms that the Bookcount Endpoint returns Status Bas Request for Get Request without language param.
func TestBookCountGet(t *testing.T) {
	// Create a request to your endpoint with the GET method
	req, err := http.NewRequest("GET", Endpoints.BookCount, nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(statistics.BookCountHandler)

	// Serve the request to the handler
	handler.ServeHTTP(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}
}

// TestBookCountGetLanguage confirms that the Bookcount Endpoint returns Status OK for Get Request with language param.
func TestBookCountGetLanguage(t *testing.T) {
	// Create a request to your endpoint with the GET method
	req, err := http.NewRequest("GET", Endpoints.BookCount+"?languages=no", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(statistics.BookCountHandler)

	// Serve the request to the handler
	handler.ServeHTTP(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

// TestBookCountGetLanguageWrong confirms that the Bookcount Endpoint returns Status Bad Request for Get Request with wrongful language param.
func TestBookCountGetLanguageWrong(t *testing.T) {
	// Create a request to your endpoint with the GET method
	req, err := http.NewRequest("GET", Endpoints.BookCount+"?languages=nog", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(statistics.BookCountHandler)

	// Serve the request to the handler
	handler.ServeHTTP(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}
}

// TestBookCountGetLanguages confirms that the Bookcount Endpoint returns Status OK for Get Request with multiple language param.
func TestBookCountGetLanguages(t *testing.T) {
	// Create a request to your endpoint with the GET method
	req, err := http.NewRequest("GET", Endpoints.BookCount+"?languages=no,es", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(statistics.BookCountHandler)

	// Serve the request to the handler
	handler.ServeHTTP(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

// TestBookCountGetLanguagesWrong confirms that the Bookcount Endpoint returns Status Bad Request for Get Request with same language param.
func TestBookCountGetLanguagesWrong(t *testing.T) {
	// Create a request to your endpoint with the GET method
	req, err := http.NewRequest("GET", Endpoints.BookCount+"?languages=no,no", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(statistics.BookCountHandler)

	// Serve the request to the handler
	handler.ServeHTTP(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}
}

// TestBookcountMethodNotAllowed confirms that the Bookcount Endpoint returns Status Not Implemented for Methods other than GET.
func TestBookcountMethodNotAllowed(t *testing.T) {
	// Create a request to your endpoint with a method other than GET
	req, err := http.NewRequest("POST", Endpoints.BookCount, nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(statistics.BookCountHandler)

	// Serve the request to the handler
	handler.ServeHTTP(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusNotImplemented {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusNotImplemented)
	}

	req, err = http.NewRequest("PUT", Endpoints.BookCount, nil)
	if err != nil {
		t.Fatal(err)
	}

	// Serve the request to the handler
	handler.ServeHTTP(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusNotImplemented {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusNotImplemented)
	}
}

// TestReadershipGet confirms that the Readership Endpoint returns Status Bas Request for Get Request without language param.
func TestReadershipGet(t *testing.T) {
	// Create a request to your endpoint with the GET method
	req, err := http.NewRequest("GET", Endpoints.Readership, nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(statistics.ReadershipHandler)

	// Serve the request to the handler
	handler.ServeHTTP(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}
}

// TestReadershipGetLanguage confirms that the Readership Endpoint returns Status OK for Get Request with language param.
func TestReadershipGetLanguage(t *testing.T) {
	// Create a request to your endpoint with the GET method
	req, err := http.NewRequest("GET", Endpoints.Readership+"no", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(statistics.ReadershipHandler)

	// Serve the request to the handler
	handler.ServeHTTP(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

// TestReadershipGetWrong confirms that the Readership Endpoint returns Status Bad Request for Get Request with wrongful language param.
func TestReadershipGetWrong(t *testing.T) {
	// Create a request to your endpoint with the GET method
	req, err := http.NewRequest("GET", Endpoints.Readership+"nog", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(statistics.ReadershipHandler)

	// Serve the request to the handler
	handler.ServeHTTP(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}
}

// TestReadershipGetLanguages confirms that the Readership Endpoint returns Status Bad Request for Get Request with multiple language param.
func TestReadershipGetLanguages(t *testing.T) {
	// Create a request to your endpoint with the GET method
	req, err := http.NewRequest("GET", Endpoints.Readership+"no,es", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(statistics.ReadershipHandler)

	// Serve the request to the handler
	handler.ServeHTTP(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}
}

// TestReadershipGetLimit confirms that the Readership Endpoint returns Status OK for Get Request with limit param.
func TestReadershipGetLimit(t *testing.T) {
	// Create a request to your endpoint with the GET method
	req, err := http.NewRequest("GET", Endpoints.Readership+"no/?limit=1", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(statistics.ReadershipHandler)

	// Serve the request to the handler
	handler.ServeHTTP(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

// TestReadershipGetLimitWrong confirms that the Readership Endpoint returns Status Bad Request for Get Request with wrongful limit param.
func TestReadershipGetLimitWrong(t *testing.T) {
	// Create a request to your endpoint with the GET method
	req, err := http.NewRequest("GET", Endpoints.Readership+"no/?limit=one", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(statistics.ReadershipHandler)

	// Serve the request to the handler
	handler.ServeHTTP(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}
}

// TestReadershipMethodNotAllowed confirms that the Readership Endpoint returns Status Not Implemented for Methods other than GET.
func TestReadershipMethodNotAllowed(t *testing.T) {
	// Create a request to your endpoint with a method other than GET
	req, err := http.NewRequest("POST", Endpoints.Readership, nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(statistics.ReadershipHandler)

	// Serve the request to the handler
	handler.ServeHTTP(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusNotImplemented {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusNotImplemented)
	}

	req, err = http.NewRequest("PUT", Endpoints.Readership, nil)
	if err != nil {
		t.Fatal(err)
	}

	// Serve the request to the handler
	handler.ServeHTTP(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusNotImplemented {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusNotImplemented)
	}
}

// TestStatusGet confirms that the Status Endpoint returns Status OK for GET Method.
func TestStatusGet(t *testing.T) {
	// Create a request to your endpoint with the GET method
	req, err := http.NewRequest("GET", Endpoints.Status, nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(statistics.StatusHandler)

	// Serve the request to the handler
	handler.ServeHTTP(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

// TestStatusMethodNotAllowed confirms that the Status Endpoint returns Status Not Implemented for Methods other than GET.
func TestStatusMethodNotAllowed(t *testing.T) {
	// Create a request to your endpoint with a method other than GET
	req, err := http.NewRequest("POST", Endpoints.Status, nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(statistics.StatusHandler)

	// Serve the request to the handler
	handler.ServeHTTP(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusNotImplemented {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusNotImplemented)
	}

	req, err = http.NewRequest("PUT", Endpoints.Status, nil)
	if err != nil {
		t.Fatal(err)
	}

	// Serve the request to the handler
	handler.ServeHTTP(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusNotImplemented {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusNotImplemented)
	}
}

// TestSupportedLanguagesGet confirms that the Supported Languages Endpoint returns Status OK for GET Method.
func TestSupportedLanguagesGet(t *testing.T) {
	// Create a request to your endpoint with the GET method
	req, err := http.NewRequest("GET", Endpoints.SupportedLanguages, nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(statistics.SupportedLanguagesHandler)

	// Serve the request to the handler
	handler.ServeHTTP(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

// TestSupportedLanguagesMethodNotAllowed confirms that the Supported Languages Endpoint returns Status Not Implemented for Methods other than GET.
func TestSupportedLanguagesMethodNotAllowed(t *testing.T) {
	// Create a request to your endpoint with a method other than GET
	req, err := http.NewRequest("POST", Endpoints.SupportedLanguages, nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(statistics.SupportedLanguagesHandler)

	// Serve the request to the handler
	handler.ServeHTTP(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusNotImplemented {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusNotImplemented)
	}

	req, err = http.NewRequest("PUT", Endpoints.SupportedLanguages, nil)
	if err != nil {
		t.Fatal(err)
	}

	// Serve the request to the handler
	handler.ServeHTTP(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusNotImplemented {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusNotImplemented)
	}
}
