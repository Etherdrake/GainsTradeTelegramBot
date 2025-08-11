package tg_utils

import (
	tgbotapi "github.com/Etherdrake/telegram-bot-api/v7"
	"io"
	"os"
	"path/filepath"
)

type FileRequestData struct {
	FilePath string
}

func (f *FileRequestData) NeedsUpload() bool {
	return f.FilePath != ""
}

func (f *FileRequestData) UploadData() (string, io.Reader, error) {
	file, err := os.Open(f.FilePath)
	if err != nil {
		return "", nil, err
	}
	return filepath.Base(f.FilePath), file, nil
}

func (f *FileRequestData) SendData() string {
	// Return the attachment format for sending
	return "attach://" + filepath.Base(f.FilePath)
}

func NewInputMediaPhotoFromFile(qrFileName string) tgbotapi.InputMediaPhoto {
	attachmentName := filepath.Base(qrFileName) // Extract just the filename without directory paths
	fileData := &FileRequestData{
		FilePath: "qr_codes/" + attachmentName, // Storing the correct relative path
	}

	// Constructing the BaseInputMedia for InputMediaPhoto
	media := tgbotapi.BaseInputMedia{
		Type:  "photo",
		Media: fileData,
	}

	return tgbotapi.InputMediaPhoto{
		BaseInputMedia: media,
	}
}
