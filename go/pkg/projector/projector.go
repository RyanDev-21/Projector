package projector

import (
	"encoding/json"
	"os"
	"path"
)

type Data struct {
	Projector map[string]map[string]string `json:"projector"`
}

type Projector struct {
	config *Config
	data   *Data
}

func NewProjector(config *Config, data *Data) *Projector {
	return &Projector{
		config: config,
		data:   data,
	}
}

func CreateProjector(config *Config, data *Data) *Projector {
	return &Projector{
		config: config,
		data:   data,
	}
}

func newDefaultProjector(config *Config) *Projector {
	return &Projector{
		config: config,
		data: &Data{
			Projector: map[string]map[string]string{},
		},
	}
}

func NewProjectorWConfig(config *Config) *Projector {
	if _, err := os.Stat(config.Config); err == nil {
		contents, err := os.ReadFile(config.Config)
		if err != nil {
			return newDefaultProjector(config)
		}
		var data Data
		err = json.Unmarshal(contents, &data)
		if err != nil {
			return newDefaultProjector(config)
		}
		return &Projector{
			config: config,
			data:   &data,
		}
	}
	return newDefaultProjector(config)
}

func (p *Projector) GetValue(key string) (string, bool) {
	curr := p.config.Pwd
	prev := ""
	out := ""
	found := false
	for curr != prev {
		if dir, ok := p.data.Projector[curr]; ok {
			if value, ok := dir[key]; ok {
				out = value
				found = true
				break
			}
		}
		prev = curr
		curr = path.Dir(curr)
	}
	return out, found
}

func (p *Projector) GetValueAll() map[string]string {
	out := map[string]string{}
	paths := []string{}
	curr := p.config.Pwd
	prev := ""
	for curr != prev {
		paths = append(paths, curr)
		prev = curr
		curr = path.Dir(curr)
	}

	for i := len(paths) - 1; i >= 0; i-- {
		if dir, ok := p.data.Projector[paths[i]]; ok {
			for k, v := range dir {
				out[k] = v
			}
		}
	}
	return out
}

func (p *Projector) SetValue(key, value string) {
	if _, ok := p.data.Projector[p.config.Pwd]; !ok {
		p.data.Projector[p.config.Pwd] = map[string]string{}
	}

	p.data.Projector[p.config.Pwd][key] = value
}

func (p *Projector) RemoveValue(key string) {
	if dir, ok := p.data.Projector[p.config.Pwd]; ok {
		delete(dir, key)
	}
}
