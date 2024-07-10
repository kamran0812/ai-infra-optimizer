package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/kamran0812/ai-infra-optimizer/internal/cloud"
	"github.com/kamran0812/ai-infra-optimizer/internal/ml"
	"github.com/kamran0812/ai-infra-optimizer/internal/optimizer"
	"github.com/kamran0812/ai-infra-optimizer/internal/storage"
)

func main() {
	db, err := storage.NewDatabase("infradata.db")
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer db.Close()

	predictor := ml.NewPredictor()

	// Get cloud providers from command line arguments
	providerNames := os.Args[1:]
	if len(providerNames) == 0 {
		log.Fatal("Please specify at least one cloud provider (AWS, Azure, GCP)")
	}

	var providers []cloud.Provider
	for _, name := range providerNames {
		provider, err := cloud.ProviderFactory(strings.TrimSpace(name))
		if err != nil {
			log.Fatalf("Failed to create provider %s: %v", name, err)
		}
		providers = append(providers, provider)
	}

	opt := optimizer.NewOptimizer(db, predictor, providers...)

	recommendations, err := opt.GenerateRecommendations()
	if err != nil {
		log.Fatalf("Error generating recommendations: %v", err)
	}

	fmt.Println("Optimization Recommendations:")
	for _, rec := range recommendations {
		fmt.Printf("- %s\n", rec)
	}
}
