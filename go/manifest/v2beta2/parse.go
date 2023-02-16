package v2beta2

import (
	"errors"
	"fmt"
	"strings"

	corev1 "k8s.io/api/core/v1"

	types "github.com/akash-network/akash-api/go/node/types/v1beta3"
)

var (
	errUnknownServiceProtocol     = errors.New("unknown service protocol")
	ErrUnsupportedServiceProtocol = errors.New("unsupported service protocol")
)

func ParseServiceProtocol(input string) (ServiceProtocol, error) {
	var result ServiceProtocol

	// This is not a case-sensitive parse, so make all input uppercase
	input = strings.ToUpper(input)

	switch input {
	case "TCP", "": // The empty string (no input) implies TCP
		result = TCP
	case "UDP":
		result = UDP
	default:
		return result, ErrUnsupportedServiceProtocol
	}

	return result, nil
}

func (sp ServiceProtocol) ToString() string {
	return string(sp)
}

func (sp ServiceProtocol) ToKube() (corev1.Protocol, error) {
	switch sp {
	case TCP:
		return corev1.ProtocolTCP, nil
	case UDP:
		return corev1.ProtocolUDP, nil
	}

	return "", fmt.Errorf("%w: %v", errUnknownServiceProtocol, sp)
}

// GetResources returns list of resources in a group
func (g *Group) GetResources() []types.Resources {
	resources := make([]types.Resources, 0, len(g.Services))
	for _, s := range g.Services {
		resources = append(resources, types.Resources{
			Resources: s.Resources,
			Count:     s.Count,
		})
	}

	return resources
}
