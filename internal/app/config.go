package app

type Config struct {
	// The port the application is listening to.
	Port string

	// If set to "local" some rules are loosen to make local development easier
	Environment string

	// Version tag
	Version string

	// GCP Project ID
	ProjectID string

	// When set, activates emulator mode for the firestore client
	FirestoreEmulatorHost string

	// CSRF secret
	CSRFSecret string

	// The root url of the auth callback. Must be configured accordingly in GCP API Credentials (Authorized redirect URIs)
	AuthCallbackHost string

	// Session passphrase
	SessionPassphrase string

	// Google client key
	GoogleClientKey string

	// Google secret
	GoogleSecret string

	// Target GCS Buckets
	Buckets []string

	// List of authorised user's email addresses
	UserWhitelist []string
}
