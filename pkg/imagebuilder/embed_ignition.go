package imagebuilder

import (
	"io"
	"os"

	"github.com/openshift/assisted-image-service/pkg/isoeditor"
)

const (
	outputImage = "output/fleeting.iso"
)

func getIgnition() ([]byte, error) {
	ignition, err := os.Open("data/test_ignition.ign")
	if err != nil {
		return nil, err
	}
	defer ignition.Close()

	return io.ReadAll(ignition)
}

func BuildImage(baseImage string) error {
	ignition, err := getIgnition()
	if err != nil {
		return err
	}
	ignitionContent := &isoeditor.IgnitionContent{Config: ignition}

	custom, err := isoeditor.NewRHCOSStreamReader(baseImage, ignitionContent, nil)
	if err != nil {
		return err
	}
	defer custom.Close()

	output, err := os.Create(outputImage)
	if err != nil {
		return err
	}
	defer output.Close()

	_, err = io.Copy(output, custom)
	return err
}