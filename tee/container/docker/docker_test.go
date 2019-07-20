package docker

import (
	"crypto/sha256"
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/gaeanetwork/gaea-core/common"
	"github.com/stretchr/testify/assert"
)

func Test_Create(t *testing.T) {
	c, err := Create()
	assert.NoError(t, err)
	defer c.Destroy()
	assert.NotNil(t, c.client)
	assert.NotNil(t, c.id)

	// Repeat create
	c, err1 := Create()
	assert.NoError(t, err1)

	cmd := fmt.Sprintf("echo 'hello world'")
	container, err2 := c.startFunc(cmd)
	assert.NoError(t, err2)

	container1, err3 := c.client.InspectContainer(container.ID)
	assert.NoError(t, err3)
	assert.Equal(t, container1.Args[1], cmd)
}

func Test_Upload(t *testing.T) {
	c, err := Create()
	assert.NoError(t, err)
	defer c.Destroy()

	algorithm, err := ioutil.ReadFile("/home/rabbit/teetest/client/resume")
	assert.NoError(t, err)
	A, err := ioutil.ReadFile("/home/rabbit/teetest/A/A_resume.txt")
	assert.NoError(t, err)
	B, err := ioutil.ReadFile("/home/rabbit/teetest/B/B_resume.txt")
	assert.NoError(t, err)
	C, err := ioutil.ReadFile("/home/rabbit/teetest/C/C_resume.txt")
	assert.NoError(t, err)
	dataList := [][]byte{A, B, C}
	err = c.Upload(algorithm, dataList)
	assert.NoError(t, err)
	hash := sha256.Sum256(algorithm)
	algorithmHash := common.BytesToHex(hash[:])
	assert.Equal(t, algorithmHash, c.algorithmHash)
}
