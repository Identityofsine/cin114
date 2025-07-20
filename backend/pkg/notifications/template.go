package notifications

import (
	"bytes"
	"os"
	"path/filepath"
	tmpltext "text/template"
)

type _template struct {
	Path string // this should be loaded everytime
}

type TemplateStore struct {
	templates map[string]_template
}

type H map[string]any

func NewTemplateStore() *TemplateStore {
	return &TemplateStore{}
}

func (t *TemplateStore) LoadTemplates(path string) error {
	// load all templates from the path and store them their path
	files, err := os.ReadDir(path)
	if err != nil {
		return err
	}

	stripped_path := filepath.Base(path)
	if stripped_path == path {
		stripped_path = ""
	}
	for _, file := range files {
		if file.IsDir() {
			// act recursively ?
			t.LoadTemplates(filepath.Join(path, file.Name()))
			continue
		}
		if filepath.Ext(file.Name()) != ".tmpl" {
			continue
		} else {
			if t.templates == nil {
				t.templates = make(map[string]_template)
			}
			// check if the template already exists but with a different path
			if tmpl, ok := t.templates[filepath.Join(stripped_path, file.Name())]; ok && tmpl.Path != filepath.Join(stripped_path, file.Name()) {
				// handle duplicate file names
				t.templates[filepath.Join(stripped_path, file.Name())] = _template{Path: filepath.Join(stripped_path, file.Name())}
				continue
			}
			// add the template to the store
			t.templates[filepath.Join(stripped_path, file.Name())] = _template{Path: filepath.Join(path, file.Name())}
		}
	}

	return nil
}

// get the template
func (t *TemplateStore) GetTemplate(path string) (string, error) {
	if tmpl, ok := t.templates[path]; ok {
		return tmpl.Path, nil
	}
	return "", os.ErrNotExist
}

func (t *TemplateStore) ParseTemplate(path string, data H) (string, error) {
	// get the template
	tmpl, err := t.GetTemplate(path)
	if err != nil {
		return "", err
	}
	// parse the template
	tmplate, err := tmpltext.ParseFiles(tmpl)
	if err != nil {
		return "", err
	}
	buf := new(bytes.Buffer)
	err = tmplate.Execute(buf, data)
	return buf.String(), err
}
