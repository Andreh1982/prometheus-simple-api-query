package models

import (
	"time"

	"github.com/prometheus/common/model"
)

type ResponsePayloadList struct {
	Status string `json:"status"`
	Data   struct {
		ResultType string `json:"resultType"`
		Result     []struct {
			Metric struct {
				Alias string `json:"alias"`
			} `json:"metric"`
			Value []interface{} `json:"value"`
		} `json:"result"`
	} `json:"data"`
}

type RulesList struct {
	Status string `json:"status"`
	Data   struct {
		Groups []struct {
			Name  string `json:"name"`
			File  string `json:"file"`
			Rules []struct {
				State    string `json:"state"`
				Name     string `json:"name"`
				Query    string `json:"query"`
				Duration int    `json:"duration"`
				Labels   struct {
					Severity string `json:"severity"`
				} `json:"labels"`
				Annotations struct {
					Description string `json:"description"`
				} `json:"annotations"`
				Alerts         []interface{} `json:"alerts"`
				Health         string        `json:"health"`
				EvaluationTime float64       `json:"evaluationTime"`
				LastEvaluation time.Time     `json:"lastEvaluation"`
				Type           string        `json:"type"`
			} `json:"rules"`
			Interval       int       `json:"interval"`
			EvaluationTime float64   `json:"evaluationTime"`
			LastEvaluation time.Time `json:"lastEvaluation"`
		} `json:"groups"`
	} `json:"data"`
}

type Alerts struct {
	Data struct {
		Alerts []struct {
			ActiveAt    time.Time `json:"activeAt"`
			Annotations struct {
			} `json:"annotations"`
			Labels struct {
				Alertname string `json:"alertname"`
			} `json:"labels"`
			State string `json:"state"`
			Value string `json:"value"`
		} `json:"alerts"`
	} `json:"data"`
	Status string `json:"status"`
}

type Configuration struct {
	ApiUrlHttp string `json:"apiurlhttp"`
	ApiUrlCLI  string `json:"apiurlcli"`
}

type TargetsResult struct {
	Active  []ActiveTarget  `json:"activeTargets"`
	Dropped []DroppedTarget `json:"droppedTargets"`
}

type ActiveTarget struct {
	DiscoveredLabels   map[string]string `json:"discoveredLabels"`
	Labels             model.LabelSet    `json:"labels"`
	ScrapePool         string            `json:"scrapePool"`
	ScrapeURL          string            `json:"scrapeUrl"`
	GlobalURL          string            `json:"globalUrl"`
	LastError          string            `json:"lastError"`
	LastScrape         time.Time         `json:"lastScrape"`
	LastScrapeDuration float64           `json:"lastScrapeDuration"`
	Health             HealthStatus      `json:"health"`
}

type DroppedTarget struct {
	DiscoveredLabels map[string]string `json:"discoveredLabels"`
}

type HealthStatus string
