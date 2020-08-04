package image

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ParseImageTags(t *testing.T) {
	t.Run("Parse valid image name without registry info", func(t *testing.T) {
		image := NewFromIdentifier("jannfis/test-image:0.1")
		assert.Empty(t, image.RegistryURL)
		assert.Empty(t, image.SymbolicName)
		assert.Equal(t, "jannfis/test-image", image.ImageName)
		assert.Equal(t, "0.1", image.ImageTag)
	})

	t.Run("Parse valid image name with registry info", func(t *testing.T) {
		image := NewFromIdentifier("gcr.io/jannfis/test-image:0.1")
		assert.Equal(t, "gcr.io", image.RegistryURL)
		assert.Empty(t, image.SymbolicName)
		assert.Equal(t, "jannfis/test-image", image.ImageName)
		assert.Equal(t, "0.1", image.ImageTag)
	})

	t.Run("Parse valid image name with source name and registry info", func(t *testing.T) {
		image := NewFromIdentifier("jannfis/orig-image=gcr.io/jannfis/test-image:0.1")
		assert.Equal(t, "gcr.io", image.RegistryURL)
		assert.Equal(t, "jannfis/orig-image", image.SymbolicName)
		assert.Equal(t, "jannfis/test-image", image.ImageName)
		assert.Equal(t, "0.1", image.ImageTag)
	})

	t.Run("Parse image without version source name and registry info", func(t *testing.T) {
		image := NewFromIdentifier("jannfis/orig-image=gcr.io/jannfis/test-image")
		assert.Equal(t, "gcr.io", image.RegistryURL)
		assert.Equal(t, "jannfis/orig-image", image.SymbolicName)
		assert.Equal(t, "jannfis/test-image", image.ImageName)
		assert.Empty(t, image.ImageTag)
	})
}

func Test_ImageToString(t *testing.T) {
	t.Run("Get string representation of full-qualified image name", func(t *testing.T) {
		imageName := "jannfis/argocd=jannfis/orig-image:0.1"
		img := NewFromIdentifier(imageName)
		assert.Equal(t, imageName, img.String())
	})
	t.Run("Get string representation of full-qualified image name with registry", func(t *testing.T) {
		imageName := "jannfis/argocd=gcr.io/jannfis/orig-image:0.1"
		img := NewFromIdentifier(imageName)
		assert.Equal(t, imageName, img.String())
	})
	t.Run("Get string representation of full-qualified image name with registry", func(t *testing.T) {
		imageName := "jannfis/argocd=gcr.io/jannfis/orig-image"
		img := NewFromIdentifier(imageName)
		assert.Equal(t, imageName, img.String())
	})
	t.Run("Get original value", func(t *testing.T) {
		imageName := "invalid==foo"
		img := NewFromIdentifier(imageName)
		assert.Equal(t, imageName, img.Original())
	})
}

func Test_WithTag(t *testing.T) {
	t.Run("Get string representation of full-qualified image name", func(t *testing.T) {
		imageName := "jannfis/argocd=jannfis/orig-image:0.1"
		nimageName := "jannfis/argocd=jannfis/orig-image:0.2"
		oImg := NewFromIdentifier(imageName)
		nImg := oImg.WithTag("0.2")
		assert.Equal(t, nimageName, nImg.String())
	})
}

func Test_ContainerList(t *testing.T) {
	t.Run("Test whether image is contained in list", func(t *testing.T) {
		images := make(ContainerImageList, 0)
		image_names := []string{"a/a:0.1", "a/b:1.2", "x/y=foo.bar/a/c:0.23"}
		for _, n := range image_names {
			images = append(images, NewFromIdentifier(n))
		}
		assert.NotNil(t, images.ContainsImage(NewFromIdentifier(image_names[0]), false))
		assert.NotNil(t, images.ContainsImage(NewFromIdentifier(image_names[1]), false))
		assert.NotNil(t, images.ContainsImage(NewFromIdentifier(image_names[2]), false))
		assert.Nil(t, images.ContainsImage(NewFromIdentifier("foo/bar"), false))
	})
}