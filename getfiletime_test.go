package getfiletime_test

import (
    "fmt"
    "os"
    "path"
    "testing"
    "time"

    "github.com/skillian/getfiletime"
)

func getTempFileName() string {
    tempdir := os.TempDir()
    for tries := 3; tries > 0; tries-- {
        tempName := fmt.Sprintf("%d.tmp", time.Now().Nanosecond())
        tempPath := path.Join(tempdir, tempName)
        if _, err := os.Stat(tempPath); os.IsNotExist(err) {
            return tempPath
        }
    }
    return ""
}

func TestGetFileTime(t *testing.T) {
    tempFilename := getTempFileName()
    if tempFilename == "" {
        t.Fatal("Failed to get a temporary file name")
    }
    now := time.Now().UTC()
    tempFile, err := os.Create(tempFilename)
    if err != nil {
        t.Fatalf("failed to open temp file %v: %v", tempFilename, err)
    }
    if err := tempFile.Close(); err != nil {
        t.Fatalf("error while closing temporary file %v: %v", tempFilename, err)
    }
    fileTime, err := getfiletime.GetFileTime(tempFilename)
    if err != nil {
        t.Fatalf("error getting file time: %v", err)
    }
    if !fileTime.CreationTime.Round(1 * time.Minute).Equal(now.Round(1 * time.Minute)) {
        t.Errorf("Got an unexpected CreationTime: %v", fileTime.CreationTime)
    }
    if !fileTime.LastAccessTime.Round(1 * time.Minute).Equal(now.Round(1 * time.Minute)) {
        t.Errorf("Got an unexpected LastAccessTime: %v", fileTime.LastAccessTime)
    }
    if !fileTime.LastWriteTime.Round(1 * time.Minute).Equal(now.Round(1 * time.Minute)) {
        t.Errorf("Got an unexpected LastWriteTime: %v", fileTime.LastWriteTime)
    }
    if err := os.Remove(tempFilename); err != nil {
        t.Errorf("Failed to remove temporary file %v: %v", tempFilename, err)
    }
    t.Log("GetFileTime probably works")
}