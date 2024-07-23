package helper

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt/v5"
	database "github.com/nk-hung/ecom-go/global"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type SignedDetails struct {
	Email string
	jwt.RegisteredClaims
	Uid primitive.ObjectID
}

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "user")

var SECRET_KEY string = os.Getenv("SECRET_KEY")

func GenerateKeyPairTokens(email string, uid primitive.ObjectID, publichKey string, privateKey string) (accessToken string, refreshToken string, err error) {
	claims := &SignedDetails{
		Email: email,
		Uid:   uid,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Local().Add(time.Hour * time.Duration(24))),
		},
	}

	refreshClaims := SignedDetails{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Local().Add(time.Hour * time.Duration(168))),
		},
	}

	accessToken, err = jwt.NewWithClaims(jwt.SigningMethodRS256, claims).SignedString(privateKey)
	if err != nil {
		log.Panic(err)
		return
	}

	refreshToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString(privateKey)
	if err != nil {
		log.Panic(err)
		return
	}

	// Verify accessToken
	token, err := jwt.ParseWithClaims(accessToken, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(publichKey), nil
	})

	if err != nil {
		log.Fatal(err)
	} else if claims, ok := token.Claims.(*SignedDetails); ok {
		fmt.Println(claims.Uid)
	} else {
		log.Fatal("unknown claims type, cannot process")
	}

	return
}

/* GenerateKeyPairs creates a PEM Private Key and Publich Key of specified byte size */
func GenerateKeyPairs(bitSize int) ([]byte, []byte, error) {
	// GeneratePrivateKey
	privateKey, err := rsa.GenerateKey(rand.Reader, bitSize)
	if err != nil {
		return nil, nil, err
	}

	// encodes Private Key from RSA to PEM format
	privateKeyByte := x509.MarshalPKCS1PrivateKey(privateKey)
	privateKeyPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: privateKeyByte,
	})

	// encodes Public Key from RSA to PEM format
	publicKey := &privateKey.PublicKey
	publicKeyByte := x509.MarshalPKCS1PublicKey(publicKey)
	publicKeyPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: publicKeyByte,
	})

	return privateKeyPEM, publicKeyPEM, nil
}

func CheckUserType(c *gin.Context, role string) (err error) {
	userType := c.GetString("user_type")
	err = nil
	if userType != role {
		err = errors.New("unauthorized to access the resouce")
		return err
	}
	return err
}

func MatchUserTypeToUid(c *gin.Context, userId string) (err error) {
	userType := c.GetString("user_type")
	uid := c.GetString("uid")

	if userType == "USER" && uid != userId {
		err = errors.New("unauthorized to access the resource")
		return err
	}

	err = CheckUserType(c, userType)
	return err
}
