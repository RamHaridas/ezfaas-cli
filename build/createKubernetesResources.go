package build

import (
	"ezfaas/models"
)

func CreateServiceYaml(yaml_obj models.Config) {
	var service models.Service

	service.APIVersion = "v1"
	service.Kind = "Service"
	service.Metadata.Name = yaml_obj.Provider.Name
	service.Spec.Selector.App = yaml_obj.Provider.Name
	service.Spec.Type = "ClusterIP"
	port := models.Port{Name: "http"}
	port.Name = "http"
	port.Port = 80
	port.TargetPort = 8080
	service.Spec.Ports = append(service.Spec.Ports, port)
}
