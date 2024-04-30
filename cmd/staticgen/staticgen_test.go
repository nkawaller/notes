package main

import (
	"testing"
)

func TestSSG(t *testing.T) {

	ssg := NewStaticSiteGenerator()

	t.Run("testing the static site generator", func(t *testing.T) {
		ssg.generateStaticSite()
	})
}
