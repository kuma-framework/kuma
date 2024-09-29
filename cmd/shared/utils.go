package shared

import (
	"encoding/json"
	"fmt"
	"path/filepath"

	"github.com/arthurbcp/kuma-cli/internal/debug"
	"github.com/arthurbcp/kuma-cli/internal/helpers"
	"github.com/arthurbcp/kuma-cli/pkg/filesystem"
	"github.com/spf13/afero"
	"gopkg.in/yaml.v3"
)

func UnmarshalFile(fileName string) (map[string]interface{}, error) {
	fs := filesystem.NewFileSystem(afero.NewOsFs())
	// Read the content of the OpenAPI file.
	fileContent, err := fs.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	// Unmarshal the JSON or YAML content into a generic map.
	fileData, err := UnmarshalByExt(fileName, []byte(fileContent))
	if err != nil {
		return nil, err
	}
	if debug.Debug {
		helpers := helpers.NewHelpers()
		j, err := helpers.PrettyMarshal(fileData)
		if err != nil {
			return nil, err
		}
		helpers.DebugPrint(fmt.Sprintf("JSON from %s", fileName), j)
	}
	return fileData, nil
}

func UnmarshalByExt(file string, configData []byte) (map[string]interface{}, error) {
	// Determine the file type based on its extension and unmarshal accordingly.
	switch filepath.Ext(file) {
	case ".yaml", ".yml":
		data, err := UnmarshalYaml(configData)
		if err != nil {
			return nil, err
		}
		return data, nil
	case ".json":
		data, err := UnmarshalJson(configData)
		if err != nil {
			return nil, err
		}
		return data, nil
	default:
		return nil, fmt.Errorf("invalid file extension: %s", file)
	}
}

// UnmarshalJson parses JSON configuration data into BuilderData.
//
// Parameters:
//   - configData: A byte slice containing JSON-formatted configuration data.
//
// Returns:
//
//	A pointer to BuilderData and an error if unmarshaling fails.
func UnmarshalJson(configData []byte) (map[string]interface{}, error) {
	fileData := make(map[string]interface{})
	err := json.Unmarshal(configData, &fileData)
	if err != nil {
		return fileData, err
	}
	return fileData, nil
}

// UnmarshalYaml parses YAML configuration data into BuilderData.
//
// Parameters:
//   - configData: A byte slice containing YAML-formatted configuration data.
//
// Returns:
//
//	A pointer to BuilderData and an error if unmarshaling fails.
func UnmarshalYaml(configData []byte) (map[string]interface{}, error) {
	fileData := make(map[string]interface{})
	err := yaml.Unmarshal(configData, &fileData)
	if err != nil {
		return fileData, err
	}
	return fileData, nil
}
