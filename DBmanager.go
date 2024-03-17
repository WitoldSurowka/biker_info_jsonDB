package main

import (
	"encoding/json"
	"os"
)

func NewJsonFileRepository(filePath string) *JsonFileRepository {
	return &JsonFileRepository{filePath: filePath}
}

func (repo *JsonFileRepository) AddRecord(phoneNumber string, city string) error {
	data, err := repo.ReadData()
	if err != nil {
		return err
	}

	for _, record := range data.Records {
		if record.PhoneNumber == phoneNumber && record.City == city {
			repo.DeleteRecord(record.ID)
			InfoStatusDeleted(phoneNumber, city)
			return err
		}
	}
	currentID := data.CurrentID
	data.CurrentID++
	data.Records = append(data.Records, Record{ID: currentID, PhoneNumber: phoneNumber, City: city})
	repo.saveData(data)
	InfoStatusAdded(phoneNumber, city)
	return err
}

func (repo *JsonFileRepository) DeleteRecord(id int) error {
	data, err := repo.ReadData()
	if err != nil {
		return err
	}

	for i, record := range data.Records {
		if record.ID == id {
			data.Records = append(data.Records[:i], data.Records[i+1:]...)
			repo.saveData(data)
			return err
		}
	}

	return nil // or some error indicating "record not found"
}

func (repo *JsonFileRepository) ReadData() (Data, error) {
	bytes, err := os.ReadFile(repo.filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return Data{CurrentID: 0, Records: []Record{}}, nil
		}
		return Data{}, err
	}
	if len(bytes) == 0 {
		// If the file is empty, treat it as an empty list of records
		return Data{CurrentID: 0, Records: []Record{}}, nil
	}

	var data Data
	if err := json.Unmarshal(bytes, &data); err != nil {
		return Data{}, nil
	}

	return data, nil
}

func (repo *JsonFileRepository) saveData(data Data) error {
	bytes, err := json.Marshal(data)
	if err != nil {
		return err
	}

	return os.WriteFile(repo.filePath, bytes, 0644)
}
