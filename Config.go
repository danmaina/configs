package configs

import (
	"github.com/danmaina/logger"
	"github.com/go-yaml/yaml"
	"io/ioutil"
	"os"
)

const (
	path = "./Config.yaml"
)



// Read Configs from file or create default Configs
func ReadConfigs(defaultConfigString string) (interface{}, error) {
	logger.DEBUG("Reading Config File or Creating Config File if not exists")

	// Fetch/ Create Yaml config file
	configFile, errFetchFile := os.OpenFile(path, os.O_RDWR|os.O_CREATE, os.ModePerm)

	if errFetchFile != nil {
		logger.ERR("An Error Occurred while initializing configs: ", errFetchFile)
		return nil, errFetchFile
	}

	// Read Contents of config file
	configFileByteArr, errReadingByteArr := ioutil.ReadAll(configFile)

	if errReadingByteArr != nil {
		logger.ERR("An Error Occurred while reading contents of config file: ", errReadingByteArr)
		return nil, errReadingByteArr
	}

	var configStruct interface{}

	// Get config from yaml
	// Get Configuration from file yaml
	errDecodingFileYaml := yaml.Unmarshal(configFileByteArr, &configStruct)

	if errDecodingFileYaml != nil {
		logger.ERR("An Error Occurred while converting yaml to Config Struct: ", errDecodingFileYaml)
		return nil, errDecodingFileYaml
	}

	defer configFile.Close()

	// Check if config file is empty? write default configs to file: Return configs from file
	if configStruct == nil {
		logger.ERR("Config File Does Not Contain any information, Loading Default Configs")

		errDecodingDefaultYaml := yaml.Unmarshal([]byte(defaultConfigString), &configStruct)

		lenConfigs, errWritingDefaultConfigs := configFile.WriteString(defaultConfigString)

		if errWritingDefaultConfigs != nil {
			logger.ERR("Could not write default configs to config file")
		} else {
			logger.INFO("Wrote Default Configs to file. Bytes Written: ", lenConfigs)
		}

		if errDecodingDefaultYaml != nil {
			logger.ERR("Error Decoding Default Configs Yaml: ", errDecodingDefaultYaml)
		}
	}

	return configStruct, nil
}
