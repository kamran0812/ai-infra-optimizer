package optimizer

import (
	"fmt"

	"github.com/kamran0812/ai-infra-optimizer/internal/cloud"
	"github.com/kamran0812/ai-infra-optimizer/internal/ml"
	"github.com/kamran0812/ai-infra-optimizer/internal/storage"
)

type Optimizer struct {
	db        *storage.Database
	predictor *ml.Predictor
	providers []cloud.Provider
}

func NewOptimizer(db *storage.Database, predictor *ml.Predictor, providers ...cloud.Provider) *Optimizer {
	return &Optimizer{
		db:        db,
		predictor: predictor,
		providers: providers,
	}
}

func (o *Optimizer) GenerateRecommendations() ([]string, error) {
	var allUsage []cloud.ResourceUsage
	for _, provider := range o.providers {
		usage, err := provider.GetResourceUsage()
		if err != nil {
			return nil, fmt.Errorf("error getting resource usage from %s: %v", provider.GetName(), err)
		}
		allUsage = append(allUsage, usage...)
	}

	if err := o.db.SaveResourceUsage(allUsage); err != nil {
		return nil, fmt.Errorf("error saving resource usage: %v", err)
	}

	historicalUsage, err := o.db.GetHistoricalUsage()
	if err != nil {
		return nil, fmt.Errorf("error getting historical usage: %v", err)
	}

	if err := o.predictor.Train(historicalUsage); err != nil {
		return nil, fmt.Errorf("error training predictor: %v", err)
	}

	var recommendations []string
	for _, usage := range allUsage {
		predictedCPU, err := o.predictor.Predict(usage.ResourceID, 24) // Predict 24 hours ahead
		if err != nil {
			return nil, fmt.Errorf("error predicting CPU usage: %v", err)
		}

		if predictedCPU < 30 && usage.CPU < 30 {
			recommendations = append(recommendations, fmt.Sprintf("Consider downsizing %s instance %s (current CPU: %.2f%%, predicted: %.2f%%)", usage.Provider, usage.ResourceID, usage.CPU, predictedCPU))
		} else if predictedCPU > 80 {
			recommendations = append(recommendations, fmt.Sprintf("Consider upsizing %s instance %s (current CPU: %.2f%%, predicted: %.2f%%)", usage.Provider, usage.ResourceID, usage.CPU, predictedCPU))
		}
	}

	return recommendations, nil
}
