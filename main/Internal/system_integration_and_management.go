package internal

import (
	"log"
	"time"
)

// SystemIntegrationAndManagement is a struct that represents the system integration and management module
type SystemIntegrationAndManagement struct {
	environmentControlSystem *EnvironmentControlSystem
	nutrientDeliverySystem *NutrientDeliverySystem
	pestAndDiseaseManagementSystem *PestAndDiseaseManagementSystem
	pruningAndTrainingSystem *PruningAndTrainingSystem
	dataCollectionAndAnalysisSystem *DataCollectionAndAnalysisSystem
	machineLearningSystem *MachineLearningSystem
	predictiveMaintenanceSystem *PredictiveMaintenanceSystem
}

// NewSystemIntegrationAndManagement creates a new instance of the system integration and management module
func NewSystemIntegrationAndManagement() (*SystemIntegrationAndManagement, error) {
	environmentControlSystem, err := NewEnvironmentControlSystem()
	if err != nil {
		return nil, err
	}
	nutrientDeliverySystem, err := NewNutrientDeliverySystem()
	if err != nil {
		return nil, err
	}
	pestAndDiseaseManagementSystem, err := NewPestAndDiseaseManagementSystem()
	if err != nil {
		return nil, err
	}
	pruningAndTrainingSystem, err := NewPruningAndTrainingSystem()
	if err != nil {
		return nil, err
	}
	dataCollectionAndAnalysisSystem, err := NewDataCollectionAndAnalysisSystem()
	if err != nil {
		return nil, err
	}
	machineLearningSystem, err := NewMachineLearningSystem()
	if err != nil {
		return nil, err
	}
	predictiveMaintenanceSystem, err := NewPredictiveMaintenanceSystem()
	if err != nil {
		return nil, err
	}
	return &SystemIntegrationAndManagement{
		environmentControlSystem: environmentControlSystem,
		nutrientDeliverySystem: nutrientDeliverySystem,
		pestAndDiseaseManagementSystem: pestAndDiseaseManagementSystem,
		pruningAndTrainingSystem: pruningAndTrainingSystem,
		dataCollectionAndAnalysisSystem: dataCollectionAndAnalysisSystem,
		machineLearningSystem: machineLearningSystem,
		predictiveMaintenanceSystem: predictiveMaintenanceSystem,
	}, nil
}

// Update updates the system integration and management module
func (s *SystemIntegrationAndManagement) Update() error {
	log.Println("Updating system integration and management module...")
	// Update the various systems
	s.environmentControlSystem.Update()
	s.nutrientDeliverySystem.Update()
	s.pestAndDiseaseManagementSystem.Update()
	s.pruningAndTrainingSystem.Update()
	s.dataCollectionAndAnalysisSystem.Update()
	s.machineLearningSystem.Update()
	s.predictiveMaintenanceSystem.Update()
	return nil
}
