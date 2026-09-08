package valueobjects

import (
	"errors"
	"regexp"
)

var uuidRegex = regexp.MustCompile(
	`^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$`,
)

func parseUUID(value string) error {
	if !uuidRegex.MatchString(value) {
		return errors.New("invalid uuid format")
	}
	return nil
}
