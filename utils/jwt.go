package utils

import "os"

var secretKey = []byte(os.Getenv("SECRETE_KEY"))