package services

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func GetGoogleRecaptchaSiteKey() string {
	key := os.Getenv("GOOGLE_RECAPTCHA_SITE_KEY")
	if key == "" {
		// This is a development default. The environment variable should always be set in production.
		key = "6LeY_QsUAAAAAOlpVw4MhoLEr50h-dM80oz6M2AX"
	}
	return key
}
func GetGoogleRecaptchaSiteSecret() string {
	key := os.Getenv("GOOGLE_RECAPTCHA_SITE_SECRET")
	if key == "" {
		// This is a development default. The environment variable should always be set in production.
		key = "6LeY_QsUAAAAAHIALCtm0GKfk-UhtXoyJKarnRV8"
	}

	return key
}

type recaptchaResponse struct {
	Success bool `json:"success"`
}

func IsHuman(req *http.Request) bool {
        return true
}
