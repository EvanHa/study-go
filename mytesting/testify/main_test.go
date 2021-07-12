// godoc.org/github.com/stretchr/testify/assert
package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilterUnique(t *testing.T) {
	input := []Developer{
		Developer{Name: "Elliot"},
		Developer{Name: "Elliot"},
		Developer{Name: "David"},
		Developer{Name: "Alexander"},
		Developer{Name: "Eva"},
		Developer{Name: "Alan"},
	}

	expected := []string{
		"Elliot",
		"David",
		"Alexander",
		"Eva",
		"Alan",
	}

	assert.Equal(t, expected, FilterUnique(input))
}

func TestNegativeFilterUnique(t *testing.T) {
	input := []Developer{
		Developer{Name: "Elliot"},
		Developer{Name: "Elliot"},
		Developer{Name: "David"},
		Developer{Name: "Alexander"},
		Developer{Name: "Eva"},
		Developer{Name: "Alan"},
	}

	expected := []string{
		"Elliot",
		"Eva",
		"Alan",
	}

	assert.NotEqual(t, expected, FilterUnique(input))
}
