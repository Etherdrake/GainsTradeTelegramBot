package wallet

import (
	"github.com/divan/qrlogo"
	"github.com/skip2/go-qrcode"
	"image/png"
	"io/ioutil"
	"log"
	"os"
)

func CreateQR(pubkey string) (nameOfFile string, err error) {
	filePath := "./qr_codes/" + pubkey + ".png"
	err = qrcode.WriteFile(pubkey, qrcode.Medium, 256, filePath)
	if err != nil {
		return "", err
	}
	return filePath, nil
}

// CreateQRWithLogo generates a QR code with an embedded logo and saves it to a file.
func CreateQRWithLogo(str string) (filePath string, err error) {
	// Open the logo.png image from the current directory
	logoFile, err := os.Open("logo_64.png")
	if err != nil {
		log.Println("Error opening ./logo_64.png: ", err)
		return "", err
	}
	defer logoFile.Close()

	logo, err := png.Decode(logoFile)
	if err != nil {
		log.Println("Error decoding logo: ", logoFile.Name(), err)
		return "", err
	}

	// Assuming size is defined or passed somewhere else in your code
	size := 256

	// Generate the QR code with the logo
	encode, err := qrlogo.Encode(str, logo, size)
	if err != nil {
		log.Println("Error encoding QR code: ", err)
		return "", err
	}

	// Define the filePath
	filePath = "./qr_codes/" + str + ".png"

	// Log the filePath for debugging
	log.Println("Writing QR to:", filePath)

	// Check if the qr_codes directory exists; if not, create it
	if _, err := os.Stat("./qr_codes/"); os.IsNotExist(err) {
		err = os.Mkdir("./qr_codes/", 0755)
		if err != nil {
			log.Println("Error creating directory:", err)
			return "", err
		}
	}

	// Write buffer contents to the file
	err = ioutil.WriteFile(filePath, encode.Bytes(), 0644)
	if err != nil {
		log.Println("Error writing file:", err)
		return "", err
	}

	log.Println("QR code generated successfully:", filePath)
	return filePath, nil
}
