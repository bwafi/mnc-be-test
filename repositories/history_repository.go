package repositories

import (
	"encoding/json"
	"io/ioutil"
	"mnc-golang-test/models"
)

type HistoryRepository struct {
	FilePath string
}

func (r *HistoryRepository) AddHistory(history models.History) error {
	histories, err := r.GetHistories()
	if err != nil {
		return err
	}
	histories = append(histories, history)
	data, err := json.MarshalIndent(histories, "", "  ")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(r.FilePath, data, 0644)
}

func (r *HistoryRepository) GetHistories() ([]models.History, error) {
	file, err := ioutil.ReadFile(r.FilePath)
	if err != nil {
		return nil, err
	}
	var histories []models.History
	json.Unmarshal(file, &histories)
	return histories, nil
}
