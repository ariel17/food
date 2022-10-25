package repositories

import (
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v3"

	"github.com/ariel17/food/configs"
	"github.com/ariel17/food/internal/entities"
)

type repositoryYAML struct {
	File   *os.File
	Plates []struct {
		Name   string `yaml:"name"`
		OnlyOn string `yaml:"only_on"`
		Steps  []struct {
			Name   string  `yaml:"name"`
			Type   string  `yaml:"type"`
			Amount float64 `yaml:"amount"`
			Unit   string  `yaml:"unit"`
		} `yaml:"steps"`
	} `yaml:"plates"`
}

func NewRepositoryYAML() Repository {
	r := repositoryYAML{}
	var err error
	r.File, err = os.Open(configs.GetYAMLPath())
	if err != nil {
		panic(err)
	}
	b, err := ioutil.ReadAll(r.File)
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(b, &r)
	if err != nil {
		panic(err)
	}
	return &r
}

func (r *repositoryYAML) JoinPlatesSteps(plates []entities.Plate) ([]entities.Step, error) {
	mapSteps := map[string]entities.Step{}
	for _, p := range plates {
		for _, pp := range r.Plates {
			if p.Name == pp.Name {
				for _, s := range pp.Steps {
					key := fmt.Sprintf("%s-%s", s.Name, s.Type)
					if v, exists := mapSteps[key]; exists {
						// TODO: assuming same unit... what about different units?
						v.Amount += s.Amount
					} else {
						mapSteps[key] = entities.Step{
							Ingredient: entities.Ingredient{
								Name: s.Name,
								Type: s.Type,
							},
							Amount: s.Amount,
							Unit:   s.Unit,
						}
					}
				}
			}
		}
	}
	steps := []entities.Step{}
	for _, v := range mapSteps {
		steps = append(steps, v)
	}
	return steps, nil
}

func (r *repositoryYAML) GetStepsForPlate(plate entities.Plate) ([]entities.Step, error) {
	steps := []entities.Step{}
	for _, p := range r.Plates {
		if plate.Name == p.Name {
			for _, s := range p.Steps {
				steps = append(steps, entities.Step{
					Ingredient: entities.Ingredient{
						Name: s.Name,
						Type: s.Type,
					},
					Amount: s.Amount,
					Unit:   s.Unit,
				})
			}
		}
	}
	return steps, nil
}

func (r *repositoryYAML) GetAllPlates() ([]entities.Plate, error) {
	plates := []entities.Plate{}
	for _, p := range r.Plates {
		plates = append(plates, entities.Plate{
			Name:   p.Name,
			OnlyOn: p.OnlyOn,
		})
	}
	return plates, nil
}

func (r *repositoryYAML) Close() {
	if r.File != nil {
		if err := r.File.Close(); err != nil {
			panic(err)
		}
	}
}
