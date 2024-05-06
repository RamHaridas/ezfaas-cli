package build

import (
	"fmt"
)

func BuildEzfaasFunction(path string) (int, error) {
	yaml_obj, err := ReadYaml(path)

	if err != nil {
		fmt.Println(err.Error())
		return 1, err
	}
	fmt.Println(yaml_obj)
	return 0, nil
}
