package config

import (
	"time"

	"github.com/zeromicro/go-zero/core/service"
)

type (
	Condition struct {
		Key   string
		Value string
		Type  string `json:",default=match,options=match|contains"`
		Op    string `json:",default=and,options=and|or"`
	}

	ElasticSearchConf struct {
		Hosts         []string
		Index         string
		DocType       string `json:",default=doc"`
		TimeZone      string `json:",optional"`
		MaxChunkBytes int    `json:",default=15728640"` // default 15M
		Compress      bool   `json:",default=false"`
		UserName   string    `json:",default=elastic"`
		PassWord   string    `json:",default=z123456"`
	}

	Filter struct {
		Action     string      `json:",options=drop|remove_field|transfer"`
		Conditions []Condition `json:",optional"`
		Fields     []string    `json:",optional"`
		Field      string      `json:",optional"`
		Target     string      `json:",optional"`
	}

	KafkaConf struct {
		service.ServiceConf
		Brokers    []string
		Group      string
		Topics     []string
		Offset     string `json:",options=first|last,default=last"`
		Conns      int    `json:",default=1"`
		Consumers  int    `json:",default=8"`
		Processors int    `json:",default=8"`
		MinBytes   int    `json:",default=10240"`    // 10K
		MaxBytes   int    `json:",default=10485760"` // 10M
	}

	Cluster struct {
		Input struct {
			Kafka KafkaConf
		}
		Filters []Filter `json:",optional"`
		Output  struct {
			ElasticSearch ElasticSearchConf
		}
	}

	Config struct {
		Clusters    []Cluster
		GracePeriod time.Duration `json:",default=10s"`
	}
)
