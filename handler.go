package main

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"time"
)

type Visitor struct {
	IP                 string  `json:"ip"`
	Network            string  `json:"network"`
	Version            string  `json:"version"`
	City               string  `json:"city"`
	Region             string  `json:"region"`
	RegionCode         string  `json:"region_code"`
	Country            string  `json:"country"`
	CountryName        string  `json:"country_name"`
	CountryCode        string  `json:"country_code"`
	CountryCodeISO3    string  `json:"country_code_iso3"`
	CountryCapital     string  `json:"country_capital"`
	CountryTLD         string  `json:"country_tld"`
	ContinentCode      string  `json:"continent_code"`
	InEU               bool    `json:"in_eu"`
	Postal             string  `json:"postal"`
	Latitude           float64 `json:"latitude"`
	Longitude          float64 `json:"longitude"`
	Timezone           string  `json:"timezone"`
	UTCOffset          string  `json:"utc_offset"`
	CountryCallingCode string  `json:"country_calling_code"`
	Currency           string  `json:"currency"`
	CurrencyName       string  `json:"currency_name"`
	Languages          string  `json:"languages"`
	CountryArea        int     `json:"country_area"`
	CountryPopulation  int64   `json:"country_population"`
	ASN                string  `json:"asn"`
	Org                string  `json:"org"`
}

func AddVisitor(w http.ResponseWriter, r *http.Request) {
	var v Visitor
	if err := json.NewDecoder(r.Body).Decode(&v); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	_, err := DB.Exec(ctx, `
		INSERT INTO visitors (
			ip, network, version, city, region, region_code, country, country_name, country_code,
			country_code_iso3, country_capital, country_tld, continent_code, in_eu, postal, latitude,
			longitude, timezone, utc_offset, country_calling_code, currency, currency_name, languages,
			country_area, country_population, asn, org
		) VALUES (
			$1, $2, $3, $4, $5, $6, $7, $8, $9,
			$10, $11, $12, $13, $14, $15, $16,
			$17, $18, $19, $20, $21, $22, $23,
			$24, $25, $26, $27
		);`,
		v.IP, v.Network, v.Version, v.City, v.Region, v.RegionCode, v.Country, v.CountryName, v.CountryCode,
		v.CountryCodeISO3, v.CountryCapital, v.CountryTLD, v.ContinentCode, v.InEU, v.Postal, v.Latitude,
		v.Longitude, v.Timezone, v.UTCOffset, v.CountryCallingCode, v.Currency, v.CurrencyName, v.Languages,
		v.CountryArea, v.CountryPopulation, v.ASN, v.Org,
	)
	if err != nil {
		http.Error(w, "DB error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
func GetVisitorCount(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	var count int
	if err := DB.QueryRow(ctx, "SELECT COUNT(*) FROM visitors").Scan(&count); err != nil {
		http.Error(w, "DB error", http.StatusInternalServerError)
		return
	}

	w.Write([]byte(strconv.Itoa(count)))
}
func GetAllVisitors(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	rows, err := DB.Query(ctx, "SELECT ip, network, version, city, region, region_code, country, country_name, country_code, country_code_iso3, country_capital, country_tld, continent_code, in_eu, postal, latitude, longitude, timezone, utc_offset, country_calling_code, currency, currency_name, languages, country_area, country_population, asn, org FROM visitors")
	if err != nil {
		http.Error(w, "DB error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var visitors []Visitor

	for rows.Next() {
		var v Visitor
		err := rows.Scan(
			&v.IP, &v.Network, &v.Version, &v.City, &v.Region, &v.RegionCode, &v.Country, &v.CountryName, &v.CountryCode,
			&v.CountryCodeISO3, &v.CountryCapital, &v.CountryTLD, &v.ContinentCode, &v.InEU, &v.Postal, &v.Latitude,
			&v.Longitude, &v.Timezone, &v.UTCOffset, &v.CountryCallingCode, &v.Currency, &v.CurrencyName, &v.Languages,
			&v.CountryArea, &v.CountryPopulation, &v.ASN, &v.Org,
		)
		if err != nil {
			http.Error(w, "Error scanning row", http.StatusInternalServerError)
			return
		}
		visitors = append(visitors, v)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(visitors)
}
