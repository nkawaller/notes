package main

import (
	"testing"
)

func TestSSG(t *testing.T) {

	t.Run("testing the static site generator", func(t *testing.T) {
		generateStaticSite()
	})
}
