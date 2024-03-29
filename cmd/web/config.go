package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/larry-lw-chan/goti/internal/filestore"
)

func loadEnvConfig(c *Config) {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	// Override the default template path with user configuration
	if os.Getenv("TEMPLATE_PATH") != "" {
		c.TmplPath = os.Getenv("TEMPLATE_PATH")
	}

	// Override the default port with user configuration
	if os.Getenv("PORT") != "" {
		c.Port = os.Getenv("PORT")
	}

	// Override the filestore local folder with user configuration
	if os.Getenv("LOCAL_STORE") != "" {
		filestore.FS = filestore.LocalStore{
			LocalFolder: os.Getenv("LOCAL_STORE"),
		}
	}

	/*********************************************
	* Session Store Configurations
	*********************************************/
	if os.Getenv("AUTH_KEY") != "" {
		c.AuthKey = os.Getenv("AUTH_KEY")
	}

	if os.Getenv("ENCRYPTION_KEY") != "" {
		c.EncryptionKey = os.Getenv("ENCRYPTION_KEY")
	}

	if os.Getenv("MAX_AGE") != "" {
		c.MaxAge = os.Getenv("MAX_AGE")
	}
}
