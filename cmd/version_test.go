package cmd

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVersionCmd(t *testing.T) {
	// Set up a buffer to capture output
	var buf bytes.Buffer
	rootCmd.SetOut(&buf) // Ensure output goes to the buffer

	// Set the arguments to simulate running the "version" command
	rootCmd.SetArgs([]string{"version"})

	// Execute the command
	err := rootCmd.Execute()

	// Ensure no errors occurred
	assert.NoError(t, err, "The version command should run without errors")

	// Capture the output from the buffer
	output := buf.String()

	// Define the expected output
	expectedOutput := "Aegis v0.1\n"

	// Compare the actual output with the expected output
	assert.Equal(t, expectedOutput, output, "Version output should match the expected string")
}
