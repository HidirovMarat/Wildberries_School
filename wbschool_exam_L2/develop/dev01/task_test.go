package main

import (
	"testing"
)

// Тест получения времени с NTP-сервера по умолчанию
func TestGetNTPTime_DefaultServer(t *testing.T) {
	ntpTime, err := GetNTPTime("")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if ntpTime.IsZero() {
		t.Fatalf("Expected a valid time, got zero time")
	}
}

// Тест обработки ошибки при передаче невалидного NTP-сервера
func TestGetNTPTime_InvalidServer(t *testing.T) {
	_, err := GetNTPTime("invalid.ntp.server")
	if err == nil {
		t.Fatalf("Expected an error, got nil")
	}
}
