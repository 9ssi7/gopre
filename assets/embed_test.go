package assets

import (
	"fmt"
	"reflect"
	"testing"
)

func TestEmbedMailTemplate(t *testing.T) {
	// Get the embedded FS
	fs := EmbedMailTemplate()

	// Test if the template files exist

	for i := 0; i < reflect.ValueOf(Templates).NumField(); i++ {
		templateName := reflect.ValueOf(Templates).Type().Field(i).Name
		path := reflect.ValueOf(Templates).Field(i).String()
		t.Run(templateName, func(t *testing.T) {
			_, err := fs.Open(fmt.Sprintf("mail/%s.html", path))
			if err != nil {
				t.Errorf("Failed to open embedded template %s: %v", templateName, err)
			}
		})
	}
}
