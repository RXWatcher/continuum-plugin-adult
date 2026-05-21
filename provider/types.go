package provider

import "strings"

// CapabilityID is the manifest capability_id this plugin registers under.
// It is the key the host uses to look up "this plugin's ID" in a ProviderIDs
// map and the URL scheme used for image paths.
const CapabilityID = "adult"

// EncodeProviderID combines a source slug and source-local ID into the single
// canonical string the host stores under ProviderIDs["adult"].
//
// Format: "<slug>:<id>". Slugs must not contain ":".
func EncodeProviderID(sourceSlug, sourceID string) string {
	if sourceSlug == "" || sourceID == "" {
		return ""
	}
	return sourceSlug + ":" + sourceID
}

// DecodeProviderID is the inverse of EncodeProviderID.
//
// Returns (sourceSlug, sourceID, ok). ok is false if the value lacks a colon
// separator or has an empty slug or id.
func DecodeProviderID(encoded string) (string, string, bool) {
	idx := strings.Index(encoded, ":")
	if idx <= 0 || idx == len(encoded)-1 {
		return "", "", false
	}
	return encoded[:idx], encoded[idx+1:], true
}

// EncodeImagePath wraps a source-relative path with the adult:// URL scheme so
// the host can later dispatch ResolveImageURL back to the right source.
//
// Format: "adult://<sourceSlug>/<role>/<path>". Empty path returns empty string.
func EncodeImagePath(sourceSlug, role, path string) string {
	if path == "" || sourceSlug == "" || role == "" {
		return ""
	}
	if strings.HasPrefix(path, "/") {
		path = path[1:]
	}
	return "adult://" + sourceSlug + "/" + role + "/" + path
}

// DecodeImagePath is the inverse of EncodeImagePath.
//
// Returns (sourceSlug, role, rawPath, ok). ok is false if the input does not
// look like an adult:// URL with all three components.
func DecodeImagePath(encoded string) (string, string, string, bool) {
	const scheme = "adult://"
	if !strings.HasPrefix(encoded, scheme) {
		return "", "", "", false
	}
	rest := encoded[len(scheme):]
	slashIdx := strings.Index(rest, "/")
	if slashIdx <= 0 {
		return "", "", "", false
	}
	sourceSlug := rest[:slashIdx]
	rest = rest[slashIdx+1:]
	slashIdx = strings.Index(rest, "/")
	if slashIdx <= 0 {
		return "", "", "", false
	}
	role := rest[:slashIdx]
	rawPath := rest[slashIdx+1:]
	if rawPath == "" {
		return "", "", "", false
	}
	return sourceSlug, role, rawPath, true
}
