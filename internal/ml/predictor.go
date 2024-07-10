package ml

import (
	"fmt"
	"math"
	"time"

	"github.com/kamran0812/ai-infra-optimizer/internal/cloud"
	"github.com/sajari/regression"
)

type Predictor struct {
	models map[string]*regression.Regression
}

func NewPredictor() *Predictor {
	return &Predictor{
		models: make(map[string]*regression.Regression),
	}
}

func (p *Predictor) Train(data []cloud.ResourceUsage) error {
	for _, usage := range data {
		if _, exists := p.models[usage.ResourceID]; !exists {
			p.models[usage.ResourceID] = new(regression.Regression)
			p.models[usage.ResourceID].SetObserved("CPU")
			p.models[usage.ResourceID].SetVar(0, "Time")
		}

		p.models[usage.ResourceID].Train(regression.DataPoint(usage.CPU, []float64{float64(usage.Timestamp.Unix())}))
	}

	for _, model := range p.models {
		if err := model.Run(); err != nil {
			return fmt.Errorf("error running regression: %v", err)
		}
	}

	return nil
}

func (p *Predictor) Predict(resourceID string, hoursAhead float64) (float64, error) {
	model, exists := p.models[resourceID]
	if !exists {
		return 0, fmt.Errorf("no model found for resource %s", resourceID)
	}

	prediction, err := model.Predict([]float64{float64(time.Now().Unix()) + hoursAhead*3600})
	if err != nil {
		return 0, fmt.Errorf("error making prediction: %v", err)
	}

	return math.Max(0, math.Min(100, prediction)), nil
}
