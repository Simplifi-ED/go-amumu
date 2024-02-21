package primary

import (
	"fmt"
	"os"
	"testing"
)

func TestValidateConfigPath(t *testing.T) {
	// Create a temporary directory for testing
	tempDir := os.TempDir()

	// Test case 1: Path does not exist
	nonExistentPath := tempDir + "/nonexistentfile.txt"
	err := ValidateConfigPath(nonExistentPath)
	if err == nil {
		t.Errorf("Expected an error for non-existent path, got nil")
	}

	// Test case 2: Path is a directory
	dirPath := tempDir
	err = ValidateConfigPath(dirPath)
	expectedErr := fmt.Sprintf("'%s' is a directory, not a normal file", dirPath)
	if err == nil || err.Error() != expectedErr {
		t.Errorf("Expected error message '%s', got '%v'", expectedErr, err)
	}

	// Test case 3: Path exists and is a file
	filePath := tempDir + "/testfile.txt"
	file, err := os.Create(filePath)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}
	file.Close()

	err = ValidateConfigPath(filePath)
	if err != nil {
		t.Errorf("Expected no error for file path, got '%v'", err)
	}

	// Clean up: remove created file
	err = os.Remove(filePath)
	if err != nil {
		t.Fatalf("Failed to remove test file: %v", err)
	}
}
