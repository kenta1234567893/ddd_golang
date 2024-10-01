package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"

	"github.com/kenta1234567893/upsider-coding-test/src/controller"
)

func TestHealthCheck(t *testing.T) {
	e := httptest.NewServer(setupServer())
	r, err := http.Get(e.URL + "/healthcheck")
	if err != nil {
		t.Errorf("Error making request: %s", err)
		return
	}
	defer r.Body.Close()
	if r.StatusCode != 200 {
		t.Errorf("Expected 200, got %d", r.StatusCode)
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		t.Errorf("Error reading response body: %s", err)
		return
	}

	if string(body) != "OK" {
		t.Errorf("Expected 'Hello, World!', got '%s'", string(body))
	}
}

func TestIssueInvoice(t *testing.T) {

	e := httptest.NewServer(setupServer())
	defer e.Close()

	c := &http.Client{}

	// ログイン処理
	loginRequest, err := http.NewRequest("POST", e.URL+"/login", nil)
	if err != nil {
		t.Errorf("Error making request: %s", err)
		return
	}
	loginResponse, err := c.Do(loginRequest)
	if err != nil {
		t.Errorf("Error making request: %s", err)
		return
	}
	defer loginResponse.Body.Close()

	// リクエストデータ
	postData := []byte(`{
    "company_id": 1,
    "customer_id": 1,
    "payment_amount": 10000.0,
    "fee_rate": 0.04,
    "tax_rate": 0.1,
    "payment_due_date": "2024-10-17T13:30:00Z"
}`)
	request, err := http.NewRequest("POST", e.URL+"/api/invoices", bytes.NewBuffer(postData))
	if err != nil {
		t.Errorf("Error making request: %s", err)
		return
	}

	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Cookie", loginResponse.Header.Get("Set-Cookie"))

	r, err := c.Do(request)
	if err != nil {
		t.Errorf("Error making request: %s", err)
		return
	}
	defer r.Body.Close()
	if r.StatusCode != 200 {
		t.Errorf("Expected 200, got %d", r.StatusCode)
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		t.Errorf("Error reading response body: %s", err)
		return
	}
	responseData := &controller.IssueInvoiceResponse{}
	if err := json.Unmarshal(body, responseData); err != nil {
		t.Errorf("Error unmarshalling response body: %s", err)
		return
	}

	if responseData.InvoiceID == 0 {
		t.Errorf("Expected invoice ID, got %d", responseData.InvoiceID)
	}
	if responseData.CustomerID != 1 {

	}

	now := time.Now()
	nowDate := now.Truncate(time.Hour).Add(-time.Duration(now.Hour()) * time.Hour)
	if !responseData.IssueDate.Equal(nowDate) {
		t.Errorf("Expected IssueDate, got %s", responseData.IssueDate)

	}
	if responseData.PaymentAmount != 10000 {
		t.Errorf("Expected PaymentAmount, got %f", responseData.PaymentAmount)

	}
	if responseData.Fee != 400 {
		t.Errorf("Expected Fee, got %f", responseData.Fee)

	}
	if responseData.FeeRate != 0.04 {
		t.Errorf("Expected FeeRate, got %f", responseData.FeeRate)

	}
	if responseData.Tax != 40 {
		t.Errorf("Expected Tax, got %f", responseData.Tax)

	}
	if responseData.TaxRate != 0.1 {
		t.Errorf("Expected TaxRate, got %f", responseData.TaxRate)

	}
	if responseData.BillingAmount != 10440 {
		t.Errorf("Expected BillingAmount, got %f", responseData.BillingAmount)

	}
	target, err := time.Parse(time.RFC3339, "2024-10-17T13:30:00Z")
	if err != nil || !responseData.PaymentDueDate.Equal(target) {
		t.Errorf("Expected PaymentDueDate, got %s", responseData.PaymentDueDate)

	}
	if responseData.Status != "pending" {
		t.Errorf("Expected Status, got %s", responseData.Status)

	}
}

func TestSearchInvoice(t *testing.T) {

	e := httptest.NewServer(setupServer())
	defer e.Close()

	c := &http.Client{}

	// ログイン処理
	loginRequest, err := http.NewRequest("POST", e.URL+"/login", nil)
	if err != nil {
		t.Errorf("Error making request: %s", err)
		return
	}
	loginResponse, err := c.Do(loginRequest)
	if err != nil {
		t.Errorf("Error making request: %s", err)
		return
	}
	defer loginResponse.Body.Close()

	u, err := url.Parse(e.URL + "/api/invoices")
	if err != nil {
		t.Errorf("Error making request: %s", err)
		return
	}
	// クエリパラメータ
	q := u.Query()
	q.Set("company_id", "1")
	q.Set("payment_due_date_start", "2024-10-16T00:00:00Z")
	q.Set("payment_due_date_end", "2024-11-16T00:00:00Z")
	u.RawQuery = q.Encode()

	request, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		t.Errorf("Error making request: %s", err)
		return
	}

	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Cookie", loginResponse.Header.Get("Set-Cookie"))

	r, err := c.Do(request)
	if err != nil {
		t.Errorf("Error making request: %s", err)
		return
	}
	defer r.Body.Close()
	if r.StatusCode != 200 {
		t.Errorf("Expected 200, got %d", r.StatusCode)
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		t.Errorf("Error reading response body: %s", err)
		return
	}

	fmt.Println(string(body))
}
