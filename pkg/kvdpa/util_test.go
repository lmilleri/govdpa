package kvdpa

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExtractBusAndMgmtDevice(t *testing.T) {
	tests := []struct {
		testName      string
		deviceAddress string
		busName       string
		devName       string
		err           bool
	}{
		{
			testName:      "regular PCI address",
			deviceAddress: "pci/0000:65:00.1",
			busName:       "pci",
			devName:       "0000:65:00.1",
			err:           false,
		},
		{
			testName:      "no bus",
			deviceAddress: "vdpa_sim",
			busName:       "",
			devName:       "vdpa_sim",
			err:           false,
		},
		{
			testName:      "wrong address",
			deviceAddress: "pci/0000:65:00.1/0",
			busName:       "",
			devName:       "",
			err:           true,
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%s_%s", "TestExtractBusAndMgmtDevice", tt.testName), func(t *testing.T) {
			busName, devName, err := ExtractBusAndMgmtDevice(tt.deviceAddress)
			if tt.err {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
				assert.Equal(t, tt.busName, busName)
				assert.Equal(t, tt.devName, devName)
			}
		})
	}
}
