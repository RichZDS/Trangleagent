package twwidth

import "github.com/clipperhouse/displaywidth"

// These functions provide a minimal, compatible subset of the upstream
// github.com/olekukonko/tablewriter/pkg/twwidth package, without relying on
// the newer StrictEmojiNeutral field that is unavailable in the current
// displaywidth.Options version.

var defaultOptions = displaywidth.Options{
	// Keep default zero-values; callers rely mainly on Width() behaviour.
}

// DefaultCondition implements the width calculation used by tablewriter to
// truncate and align runes when printing tables.
func DefaultCondition(s string) int {
	return displaywidth.StringWidth(s, defaultOptions)
}

