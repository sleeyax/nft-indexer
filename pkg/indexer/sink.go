package indexer

import "nft-indexer/pkg/database"

type Sink struct {
	ch chan IndexResult
}

func NewSink(ch chan IndexResult) *Sink {
	return &Sink{ch: ch}
}

func (s *Sink) Write(result IndexResult) {
	s.ch <- result
}

func (s *Sink) readStep(steps ...database.CreationStep) database.CreationStep {
	var step database.CreationStep
	if len(steps) > 0 {
		step = steps[0]
	}
	return step
}

func (s *Sink) WriteError(err error, steps ...database.CreationStep) {
	s.Write(IndexResult{
		Error: err,
		Step:  s.readStep(steps...),
	})
}

func (s *Sink) WriteWarning(err error, steps ...database.CreationStep) {
	s.Write(IndexResult{
		Warning: err,
		Step:    s.readStep(steps...),
	})
}
