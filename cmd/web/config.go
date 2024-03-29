package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/larry-lw-chan/goti/internal/filestore"
)

func loadConfig(c *Config) {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found.  Please create one based on the .env.example file")
	}

	// Override the default template path with user configuration
	if os.Getenv("TEMPLATE_PATH") != "" {
		c.TmplPath = os.Getenv("TEMPLATE_PATH")
	}

	// Override the default port with user configuration
	if os.Getenv("PORT") != "" {
		c.Port = os.Getenv("PORT")
	}

	/*********************************************
	* Filestore Configurations
	*********************************************/
	if os.Getenv("FILESTORE") != "" {
		if os.Getenv("FILESTORE") == "local" {
			// Create a new local filestore
			localFS := filestore.LocalStore{}

			// Override the default local folder with user configuration
			if os.Getenv("LOCAL_FOLDER") != "" {
				localFS.LocalFolder = os.Getenv("LOCAL_FOLDER")
			}

			// Set the filestore to the local filestore
			c.Filestore = localFS
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
