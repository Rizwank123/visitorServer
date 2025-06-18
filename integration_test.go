package main

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
)

// MockDB is a simple in-memory mock for the database
type MockDB struct {
	visitors []Visitor
}

func (m *MockDB) Exec(ctx context.Context, sql string, args ...interface{}) (interface{}, error) {
	// Simulate TRUNCATE by clearing the visitors slice
	if sql == "TRUNCATE visitors" {
		m.visitors = nil
		return nil, nil
	}
	return nil, nil
}

func (m *MockDB) QueryRow(ctx context.Context, sql string, args ...interface{}) (interface{}, error) {
	// Simulate COUNT query
	if sql == "SELECT COUNT(*) FROM visitors" {
		return len(m.visitors), nil
	}
	return nil, nil
}

func (m *MockDB) Query(ctx context.Context, sql string, args ...interface{}) (interface{}, error) {
	// Simulate SELECT query for all visitors
	if sql == "SELECT ip, network, version, city, region, region_code, country, country_name, country_code, country_code_iso3, country_capital, country_tld, continent_code, in_eu, postal, latitude, longitude, timezone, utc_offset, country_calling_code, currency, currency_name, languages, country_area, country_population, asn, org FROM visitors" {
		return m.visitors, nil
	}
	return nil, nil
}

func setupMockDB() *MockDB {
	return &MockDB{visitors: []Visitor{}}
}

func setupRouter(mockDB *MockDB) http.Handler {
	r := chi.NewRouter()
	r.Post("/visitor", func(w http.ResponseWriter, r *http.Request) {
		var v Visitor
		if err := json.NewDecoder(r.Body).Decode(&v); err != nil {
			http.Error(w, "Invalid input", http.StatusBadRequest)
			return
		}
		mockDB.visitors = append(mockDB.visitors, v)
		w.WriteHeader(http.StatusCreated)
	})
	r.Get("/visitor/count", func(w http.ResponseWriter, r *http.Request) {
		count := len(mockDB.visitors)
		w.Write([]byte(strconv.Itoa(count)))
	})
	r.Get("/visitor", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(mockDB.visitors)
	})
	return r
}

func TestVisitorIntegration(t *testing.T) {
	mockDB := setupMockDB()
	router := setupRouter(mockDB)

	visitor := Visitor{
		IP:                 "10.0.0.1",
		Network:            "10.0.0.0/24",
		Version:            "IPv4",
		City:               "Test City",
		Region:             "Test Region",
		RegionCode:         "TR",
		Country:            "Test Country",
		CountryName:        "Test Country",
		CountryCode:        "TC",
		CountryCodeISO3:    "TST",
		CountryCapital:     "Test Capital",
		CountryTLD:         ".tc",
		ContinentCode:      strPtr("TS"),
		InEU:               false,
		Postal:             "12345",
		Latitude:           0.0,
		Longitude:          0.0,
		Timezone:           "UTC",
		UTCOffset:          "+0000",
		CountryCallingCode: "+1",
		Currency:           strPtr("USD"),
		CurrencyName:       strPtr("US Dollar"),
		Languages:          "en",
		CountryArea:        1000,
		CountryPopulation:  1000000,
		ASN:                "AS12345",
		Org:                "Test Org",
	}

	// Add visitor
	body, _ := json.Marshal(visitor)
	req := httptest.NewRequest("POST", "/visitor", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)

	// Get visitor count
	req = httptest.NewRequest("GET", "/visitor/count", nil)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "1", w.Body.String())

	// Get all visitors
	req = httptest.NewRequest("GET", "/visitor", nil)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	var visitors []Visitor
	err := json.NewDecoder(w.Body).Decode(&visitors)
	assert.NoError(t, err)
	assert.Len(t, visitors, 1)
	assert.Equal(t, visitor.IP, visitors[0].IP)
}

func strPtr(s string) *string {
	return &s
}
