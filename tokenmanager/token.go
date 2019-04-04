package tokenmanager

import (
	"time"

	"bitbucket.org/animor/vetstoria/configuration"
	client "bitbucket.org/animor/vetstoria/openapi"
	"github.com/antihax/optional"
	log "github.com/sirupsen/logrus"
)

func ReadToken() string {
	myClient := configuration.GetRedisClient()
	myConfig := configuration.GetConfig()

	var inlineResponse client.InlineResponse200

	redisReadError := myClient.Get("token").Scan(&inlineResponse)

	if redisReadError != nil {
		log.WithFields(log.Fields{
			"redis": myConfig.Redis.Host,
		}).Info("Token not found in cache")

		opts := &client.PartnersAuthenticationPostOpts{
			InlineObject: optional.NewInterface(client.InlineObject{
				Key:    myConfig.API.Key,
				Secret: myConfig.API.Secret,
			}),
		}

		var APIClient = client.NewAPIClient(configuration.GetVetstoriaClientConfiguration())
		inlineResponse, _, apiCallError := APIClient.AuthenticationApi.PartnersAuthenticationPost(nil, opts)

		if apiCallError != nil {
			log.WithFields(log.Fields{
				"api_url": myConfig.API.URL,
				"err":     apiCallError,
			}).Fatal("Unable to get token from API")
		}

		today := time.Now().Sub(inlineResponse.Expiry)
		redisSetValueError := myClient.Set("token", inlineResponse, today).Err()

		if redisSetValueError != nil {
			log.WithFields(log.Fields{
				"redis": myConfig.Redis.Host,
				"err":   redisSetValueError,
			}).Error("Unable to write token to Redis")
		}
	} else {
		log.WithFields(log.Fields{
			"redis": myConfig.Redis.Host,
		}).Info("Token found in cache")
	}

	return inlineResponse.Token
}
