package main

import (
	"io"

	chglog "github.com/git-chglog/git-chglog"
)

// Generator ...
type Generator interface {
	Generate(*chglog.Logger, io.Writer, string, *chglog.Config, bool) error
}

type generatorImpl struct{}

// NewGenerator ...
func NewGenerator() Generator {
	return &generatorImpl{}
}

// Generate ...
func (*generatorImpl) Generate(logger *chglog.Logger, w io.Writer, query string, config *chglog.Config, currentBranch bool) error {
	return chglog.NewGenerator(logger, config, currentBranch).Generate(w, query)
}
