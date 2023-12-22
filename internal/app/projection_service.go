package app

import (
	"context"
	"gopkg.in/yaml.v3"
	"os"
)

func NewProjectionService() *ProjectionService {
	return &ProjectionService{}
}

type ProjectionService struct {
}

type ProjectionsReq struct {
	AppID string
}

func (s *ProjectionService) AppProjections(ctx context.Context, _ ProjectionsReq) ([]Projection, error) {
	cfg, err := s.loadProjections()
	if err != nil {
		return nil, err
	}

	return cfg.Projections, nil
}

type ProjectionReq struct {
	AppID        string
	ProjectionID string
}

func (s *ProjectionService) AppProjection(ctx context.Context, req ProjectionReq) (*Projection, error) {
	p, err := s.loadProjection(req.AppID, req.ProjectionID)
	if err != nil {
		return nil, err
	}

	if p == nil {
		return nil, nil
	}

	// TODO - Augment state info from nats and docker

	return nil, nil
}

func (s *ProjectionService) loadProjection(appID, projectionID string) (*Projection, error) {
	cfg, err := s.loadProjections()
	if err != nil {
		return nil, err
	}

	for _, p := range cfg.Projections {
		if p.ID == projectionID {
			return &p, nil
		}
	}

	return nil, nil
}

func (s *ProjectionService) loadProjections() (*ProjectionConfig, error) {
	data, err := os.ReadFile("data/projections.cfg")
	if err != nil {
		// TODO - Return typed err
		return nil, err
	}

	var cfg ProjectionConfig

	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		return nil, err
	}

	for i, p := range cfg.Projections {
		if p.Name == "" {
			p.Name = p.TypeName
		}

		if p.Description == "" {
			p.Description = "Aperture Projection (see docs on how to add desc)"
		}

		if p.ID == "" {
			p.ID = p.TypeName
		}

		cfg.Projections[i] = p
	}

	return &cfg, nil
}
