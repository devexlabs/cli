package docker

import (
	"archive/tar"
	"bytes"
	"context"
	_ "embed"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"os"

	"github.com/devexlabs/cli/internal/utils"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

type ToolsMap struct {
	Tools map[string]struct{ Version string }
}

//go:embed resources/Dockerfile.tmpl
var dockerfilePath string

// https://medium.com/@Frikkylikeme/controlling-docker-with-golang-code-b213d9699998
func buildImage(client *client.Client, tags []string, dockerfile string) error {
	fmt.Println("Building docker cli container")
	ctx := context.Background()

	// Create a buffer
	buf := new(bytes.Buffer)
	tw := tar.NewWriter(buf)
	defer tw.Close()

	// Create a filereader
	dockerFileReader, err := os.Open(dockerfile)
	utils.Check(err)

	// Read the actual Dockerfile
	readDockerFile, err := ioutil.ReadAll(dockerFileReader)
	utils.Check(err)

	// Make a TAR header for the file
	tarHeader := &tar.Header{
		Name: dockerfile,
		Size: int64(len(readDockerFile)),
	}

	// Writes the header described for the TAR file
	err = tw.WriteHeader(tarHeader)
	utils.Check(err)

	// Writes the dockerfile data to the TAR file
	_, err = tw.Write(readDockerFile)
	utils.Check(err)

	dockerFileTarReader := bytes.NewReader(buf.Bytes())

	// Define the build options to use for the file
	// https://godoc.org/github.com/docker/docker/api/types#ImageBuildOptions
	buildOptions := types.ImageBuildOptions{
		Context:        dockerFileTarReader,
		SuppressOutput: true,
		Dockerfile:     dockerfile,
		Remove:         true,
		Tags:           tags,
	}

	// Build the actual image
	imageBuildResponse, err := client.ImageBuild(
		ctx,
		dockerFileTarReader,
		buildOptions,
	)

	utils.Check(err)

	// Read the STDOUT from the build process
	defer imageBuildResponse.Body.Close()
	_, err = io.Copy(os.Stdout, imageBuildResponse.Body)
	utils.Check(err)

	fmt.Println("Docker cli container builded")

	return nil
}

func WriteDockerfile(tools map[string]struct{ Version string }) {
	t, err := template.New("Dockerfile").Parse(dockerfilePath)

	utils.Check(err)

	f, err := os.Create("./Dockerfile")

	utils.Check(err)

	err = t.Execute(f, ToolsMap{tools})

	utils.Check(err)

	f.Close()

	fmt.Println("Write Dockerfile")
}

func Build() {
	client, err := client.NewEnvClient()
	if err != nil {
		log.Fatalf("Unable to create docker client: %s", err)
	}

	// Client, imagename and Dockerfile location
	tags := []string{"devexlabs/cli"}
	dockerfile := "Dockerfile"
	err = buildImage(client, tags, dockerfile)
	if err != nil {
		log.Println(err)
	}
}
