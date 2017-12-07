package config

import (
	"log"
	"os"
)

const (
	envStripeSecretKeyName  = "PAYMENT_SECRET_KEY"
	envStripePublishKeyName = "PAYMENT_PUBLISH_KEY"

	envJWTSaltName = "JWT_SALT"
	// it's better NOT to change it. if change it, all user's sessions are expired.
	additionalJWTSalt = "jfskdjfkd9jsfsdjkj97fdskjk988SAsdfs"
	//
	//envDBSaltName = "DB_SALT"

	additionalDBSalt = "MkdjksfjdkjsXjkdjsfkjdsj8jfsdjfsdkf"
)

var jwtSalt string

//var dbSalt string
var envChecked bool
var stripeSecretKey string
var stripePublishKey string

const (
	clientID     = "98fdskafjdksjaf8787687fsdfsjjkj7"
	clientSecret = "6787fdsfkjskjfkjsdkfjkskjkj7ynfd"
)

func init() {
	CheckEnvs()
}

// CheckEnvs ...
func CheckEnvs() error {
	log.Print("")
	if envChecked {
		log.Print("")
		return nil
	}
	log.Print("")

	stripeSecretKey = os.Getenv(envStripeSecretKeyName)
	if stripeSecretKey == "" {
		stripeSecretKey = "sk_test_dsfdsafdsf32fEfsdfwdfsfd"
	}
	log.Print(stripeSecretKey)
	stripePublishKey = os.Getenv(envStripePublishKeyName)
	if stripePublishKey == "" {
		stripePublishKey = "pk_test_Afsdafdawfdsf32sdfdsfsdC"
	}
	log.Print(stripePublishKey)

	envJWTSalt := os.Getenv(envJWTSaltName)

	if envJWTSalt == "" {
		envJWTSalt = "jkfjdlkajkwekknncihkfdskjkjafkdsjkljfdskjiuekjkjssd"
	}

	jwtSalt = envJWTSalt + additionalJWTSalt
	envChecked = true
	log.Print(envChecked)

	return nil
}

// GetClientID ...
func GetClientID() string {
	return clientID
}

// GetClientSecret ...
func GetClientSecret() string {
	return clientSecret
}

// GetJWTSalt ...
func GetJWTSalt() string {
	if envChecked == false {
		log.Print("NO CONFIG ENVS CHECKED")
		os.Exit(1)
	}
	return jwtSalt
}

// GetStripePublishKey ...
func GetStripePublishKey() string {
	if envChecked == false {
		log.Print("NO CONFIG ENVS CHECKED")
		os.Exit(1)
	}
	return stripePublishKey
}

// GetStripeSecretKey ...
func GetStripeSecretKey() string {
	if envChecked == false {
		log.Print("NO CONFIG ENVS CHECKED")
		os.Exit(1)
	}
	return stripeSecretKey
}
