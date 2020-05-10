package database

import (
	"gotest.tools/assert"
	"testing"
	"time"
)

func TestIsLastFriday(t *testing.T) {
	location, _ := time.LoadLocation("Europe/Rome")
	lastFriday := time.Date(2020, 4, 24, 14, 0, 0, 0, location)
	notLastFriday := time.Date(2020, 4, 17, 14, 0, 0, 0, location)
	notFriday := time.Date(2020, 4, 23, 14, 0, 0, 0, location)

	assert.Equal(t, isLastFriday(lastFriday), true)
	assert.Equal(t, isLastFriday(notFriday), false)
	assert.Equal(t, isLastFriday(notLastFriday), false)
}
