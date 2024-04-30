package main

import (
	"github.com/nkawaller/notes/internal/utils"
)

type StaticSiteGenerator struct {
	fileSystem utils.FileSystem
}

func NewStaticSiteGenerator(fileSystem utils.FileSystem) *StaticSiteGenerator {
	s := &StaticSiteGenerator{
		fileSystem: fileSystem,
	}
	return s
}

func (s *StaticSiteGenerator) generateStaticSite() error {
	err := s.fileSystem.MkdirAll("deploy/static", 0755)
	if err != nil {
		return err
	}
	return nil
}
