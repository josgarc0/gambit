package secretm

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/josgarc0/gambituser/awsgo"
	"github.com/josgarc0/gambituser/models"
)

func GetSecret(nombreSecret string) (models.SecretRDSJson, error) {
	var datosSecret models.SecretRDSJson
	fmt.Println(" > Pido Secreto " + nombreSecret)

	svc := secretsmanager.NewFromConfig(awsgo.Cfg)

	clave, err := svc.GetSecretValue(awsgo.Ctx, &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(nombreSecret),
	})
	if err != nil {
		fmt.Println(err.Error())
		return datosSecret, err
	}

	json.Unmarshal([]byte(*clave.SecretString), &datosSecret)
	fmt.Println(" > Lectura Secret Ok " + nombreSecret)
	return datosSecret, nil
}
