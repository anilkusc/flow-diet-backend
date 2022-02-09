package app

import (
	"encoding/json"

	"github.com/anilkusc/flow-diet-backend/pkg/recipe/ingredient/material"
)

func (app *App) CreateMaterial(materialJson string) error {
	var material material.Material
	var err error
	err = json.Unmarshal([]byte(materialJson), &material)
	if err != nil {
		return err
	}

	err = material.Create(app.DB)
	if err != nil {
		return err
	}
	return nil
}

func (app *App) GetMaterial(materialJson string) (string, error) {

	var material material.Material
	var err error
	err = json.Unmarshal([]byte(materialJson), &material)
	if err != nil {
		return "", err
	}
	material.Read(app.DB)
	if err != nil {
		return "", err
	}
	materialStr, err := json.Marshal(material)
	if err != nil {
		return "", err
	}
	return string(materialStr), nil
}

func (app *App) UpdateMaterial(materialJson string) error {

	var material material.Material
	err := json.Unmarshal([]byte(materialJson), &material)
	if err != nil {
		return err
	}

	material.Update(app.DB)
	if err != nil {
		return err
	}

	return nil
}

func (app *App) DeleteMaterial(materialJson string) error {

	var material material.Material
	err := json.Unmarshal([]byte(materialJson), &material)
	if err != nil {
		return err
	}

	material.Delete(app.DB)
	if err != nil {
		return err
	}

	return nil
}
