package cmd_test

import (
	"bytes"
	"github.com/corverroos/stingoftheviper/cmd"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Note that these tests fail if any test arguments are provided (so you cannot run it via the IDE).
//go:generate go test .

func TestPrecedence(t *testing.T) {
	t.Run("no args", func(t *testing.T) {
		// Run ./stingoftheviper
		root := cmd.New()
		var output bytes.Buffer
		root.SetOut(&output)
		root.SetArgs([]string{""})

		err := root.Execute()
		require.NoError(t, err, "error executing command")

		gotOutput := output.String()

		assert.Contains(t, gotOutput, "Usage:")
		assert.Contains(t, gotOutput, "foo")
		assert.Contains(t, gotOutput, "bar")
	})

	// Set favorite-color with the config file
	t.Run("config file", func(t *testing.T) {
		// Run the tests in a temporary directory
		testDir, err := os.Getwd()
		require.NoError(t, err, "error getting the current working directory")
		defer os.Chdir(testDir)

		// Copy the config file into our temporary test directory
		configB, err := ioutil.ReadFile(filepath.Join("..", "stingoftheviper.toml"))
		require.NoError(t, err, "error reading test config file")

		tmpDir, err := ioutil.TempDir("", "stingoftheviper")
		require.NoError(t, err, "error creating a temporary test directory")
		err = os.Chdir(tmpDir)
		require.NoError(t, err, "error changing to the temporary test directory")

		err = ioutil.WriteFile(filepath.Join(tmpDir, "stingoftheviper.toml"), configB, 0644)
		require.NoError(t, err, "error writing test config file")
		defer os.Remove(filepath.Join(tmpDir, "stingoftheviper.toml"))

		// Run ./stingoftheviper
		root := cmd.New()
		var output bytes.Buffer
		root.SetOut(&output)
		root.SetArgs([]string{"foo"})

		err = root.Execute()
		require.NoError(t, err, "error executing command")

		gotOutput := output.String()
		wantOutput := `Foo config:
{
 "Bar": {
  "String": "barbeque",
  "Float": 0.00001,
  "Bool": false
 },
 "String": "foo ftw",
 "Float": 12.3,
 "Bool": true
}`
		assert.Equal(t, wantOutput, gotOutput, "expected the color from the config file and the number from the flag default")
	})

	// Set favorite-color with an environment variable
	t.Run("env var", func(t *testing.T) {
		// Run STING_FAVORITE_COLOR=purple ./stingoftheviper
		os.Setenv("STING_BAR_STRING", "barbara")
		defer os.Unsetenv("STING_BAR_STRING")

		root := cmd.New()
		output := &bytes.Buffer{}
		root.SetOut(output)
		root.SetArgs([]string{"bar"})

		err := root.Execute()
		require.NoError(t, err, "error executing command")

		gotOutput := output.String()
		wantOutput := `Bar config:
{
 "String": "barbara",
 "Float": 0.1,
 "Bool": true
}`
		assert.Equal(t, wantOutput, gotOutput, "expected the bar string to us the environment variable value and the rest to use the flag default")
	})

	// Set number with a flag
	t.Run("cobra flag", func(t *testing.T) {
		// Run ./stingoftheviper --number 2
		root := cmd.New()
		output := &bytes.Buffer{}
		root.SetOut(output)
		root.SetArgs([]string{"foo", "--foo_string", "cobra"})

		err := root.Execute()
		require.NoError(t, err, "error executing command")

		gotOutput := output.String()
		wantContains := `"String": "cobra",`
		assert.Contains(t, gotOutput, wantContains)
	})
}
