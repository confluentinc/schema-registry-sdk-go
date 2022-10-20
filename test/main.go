package main

import (
	"context"
	"crypto/tls"
	"github.com/antihax/optional"
	srsdk "github.com/confluentinc/schema-registry-sdk-go"
	"log"
	"net"
	"net/http"
	"time"
)

func main() {
	config := getConfig()

	srClient := srsdk.NewAPIClient(config)

	ctx := context.WithValue(context.Background(), srsdk.ContextBasicAuth,
		srsdk.BasicAuth{
			UserName: "**",
			Password: "**",
		})

	//updateTopic(ctx, srClient)
	listTopic(ctx, srClient)
}

func getConfig() *srsdk.Configuration {
	return &srsdk.Configuration{
		BasePath:      "https://psrc-6k7wdj.us-west-2.aws.devel.cpdev.cloud",
		DefaultHeader: make(map[string]string),
		UserAgent:     "OpenAPI-Generator/1.0.0/go",
		Debug:         false,
		Servers: []srsdk.ServerConfiguration{
			{
				Url:         "/",
				Description: "No description provided",
			},
		},
		HTTPClient: &http.Client{
			Transport: DefaultTransport(),
		},
	}
}

func DefaultTransport() *http.Transport {
	return &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		}).DialContext,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
		TLSClientConfig: &tls.Config{
			MinVersion: tls.VersionTLS12,
		},
		ForceAttemptHTTP2: true,
	}
}

func listSubject(ctx context.Context, srClient *srsdk.APIClient) {
	subjects, response, err := srClient.DefaultApi.List(ctx, &srsdk.ListOpts{
		SubjectPrefix: optional.NewString(""),
		Deleted:       optional.NewBool(false),
	})
	if err != nil {
		log.Println("Got error from http list ", err)
	}
	log.Println(subjects)
	log.Println(response)
}

func createAndListBMDef(ctx context.Context, srClient *srsdk.APIClient) {
	bmAttributeDefs := []srsdk.AtlasAttributeDef{
		{
			Name:        "display_name",
			TypeName:    "string",
			IsOptional:  true,
			Cardinality: "SINGLE",
			Options:     map[string]string{"applicableEntityTypes": "[\"kafka_topic\"]"},
		},
		{
			Name:        "organization_description",
			TypeName:    "string",
			IsOptional:  true,
			Cardinality: "SINGLE",
			Options:     map[string]string{"applicableEntityTypes": "[\"kafka_topic\"]"},
		},
		{
			Name:        "organization_contact_email",
			TypeName:    "string",
			IsOptional:  true,
			Cardinality: "SINGLE",
			Options:     map[string]string{"applicableEntityTypes": "[\"kafka_topic\"]"},
		},
	}

	log.Println("creating bmDef ...")
	bmdefResponse, response, err := srClient.DefaultApi.CreateBusinessMetadataDefs(
		ctx,
		&srsdk.CreateBusinessMetadataDefsOpts{
			AtlasBusinessMetadataDef: optional.NewInterface([]srsdk.AtlasBusinessMetadataDef{
				{
					Name:          "stream_share",
					Description:   "Business Metadata for Stream Sharing",
					AttributeDefs: bmAttributeDefs,
				},
			}),
		})

	if err != nil {
		log.Println("Got error from http create ", err)
		panic(err)
	}
	log.Println(bmdefResponse)
	log.Println(response)

	time.Sleep(1 * time.Second)

	log.Println("getting bmDef ...")
	bmDef, response, err := srClient.DefaultApi.GetAllBusinessMetadataDefs(ctx,
		&srsdk.GetAllBusinessMetadataDefsOpts{Prefix: optional.NewString("stream_share")})
	if err != nil {
		log.Println("Got error from http get ", err)
		panic(err)
	}
	log.Println(bmDef)
	log.Println(response)
}

func listBM(ctx context.Context, srClient *srsdk.APIClient) {
	log.Println("getting bm ...")
	bmResponse, response, err := srClient.DefaultApi.GetBusinessMetadata(ctx, "kafka_topic", "lkc-pwj91m:topic_1")
	if err != nil {
		log.Println("Got error from http get ", err)
		panic(err)
	}
	log.Println(bmResponse)
	log.Println(response)
}

func createBM(ctx context.Context, srClient *srsdk.APIClient) {
	log.Println("creating bm ...")
	bmResponse, response, err := srClient.DefaultApi.CreateBusinessMetadata(ctx, &srsdk.CreateBusinessMetadataOpts{
		BusinessMetadata: optional.NewInterface([]srsdk.BusinessMetadata{
			{
				TypeName:   "stream_share",
				EntityType: "kafka_topic",
				EntityName: "lkc-pwj91m:topic_1",
				Attributes: map[string]interface{}{
					"display_name":               "confluent",
					"organization_description":   "confluent inc",
					"organization_contact_email": "test@confluent.io",
				},
			},
		}),
	})
	if err != nil {
		log.Println("Got error from http post ", err)
		panic(err)
	}
	log.Println(bmResponse)
	log.Println(response)
}

func updateTopic(ctx context.Context, srClient *srsdk.APIClient) {
	log.Println("updating topic")
	response, err := srClient.DefaultApi.PartialUpdateByUniqueAttributes(ctx,
		&srsdk.PartialUpdateByUniqueAttributesOpts{
			AtlasEntityWithExtInfo: optional.NewInterface(srsdk.AtlasEntityWithExtInfo{
				Entity: srsdk.AtlasEntity{
					TypeName: "kafka_topic",
					Attributes: map[string]interface{}{
						"qualifiedName":       "lkc-5w0p1q:test-stream-share",
						"description":         "Test update topic description 2",
						"__ss_orgDescription": "test add org description",
					},
				},
			}),
		})

	if err != nil {
		log.Println("Got error from http post ", err)
		panic(err)
	}
	log.Println(response)
}

func listTopic(ctx context.Context, srClient *srsdk.APIClient) {
	log.Println("listing topic")
	listResponse, response, err := srClient.DefaultApi.GetByUniqueAttributes(ctx,
		"kafka_topic",
		"lkc-35dnvo:topic_0",
		&srsdk.GetByUniqueAttributesOpts{
			MinExtInfo:            optional.NewBool(false),
			IgnoreRelationships:   optional.NewBool(false),
			IncludeInternalPrefix: optional.NewString("ss_"),
		},
	)
	if err != nil {
		log.Println("Got error from http post ", err)
		panic(err)
	}
	log.Println(listResponse)
	log.Println(response)
}
