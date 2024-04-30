package main

import (
	"os"
)

type StaticSiteGenerator struct {}

func NewStaticSiteGenerator() *StaticSiteGenerator {
	s := &StaticSiteGenerator{}
	return s
}

func (s *StaticSiteGenerator) generateStaticSite() error {
	err := os.MkdirAll("deploy/static", 0755)
	if err != nil {
		return err
	}
	return nil
}
