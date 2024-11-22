package client

import (
	"fmt"
	"net/url"
	"os"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

//go:generate go run ../../../cmd/bitsnap-accounting-freshbooks-codegen/main.go

var defaultToken = "test_Bf88SnJGFJZaYc8V9sFJxQPRLsU475oj"

func ClientToken() string {
	if tok, ok := os.LookupEnv("FRESHBOOKS_APIKEY"); ok {
		return tok
	}

	return defaultToken
}

func getQuery(client *fiber.Client, parsedUrl *url.URL, sortBy *SortBy) ([]byte, error) {
	return getQueryPage(client, parsedUrl, 1, sortBy)
}

func getQueryPage(client *fiber.Client, parsedUrl *url.URL, page uint64, sortBy *SortBy) ([]byte, error) {
	parsedUrl.Query().Add("per_page", "100")
	if page > 0 {
		q := parsedUrl.Query()
		q.Add("page", strconv.Itoa(int(page)))
		parsedUrl.RawQuery = q.Encode()
	}

	if sortBy != nil {
		q := parsedUrl.Query()
		if sortBy.Name != "" {
			if sortBy.Asc {
				q.Add("sort", sortBy.Name+"_asc")
			} else {
				q.Add("sort", sortBy.Name+"_desc")
			}
		}

		if sortBy.AfterDate != nil {
			if sortBy.UpdatedAfter {
				q.Add("search[updated_min]", sortBy.AfterDate.Format("2006-01-02"))
				q.Add("search[updated_max]", time.Now().Format("2006-01-02"))
			} else {
				q.Add("search[date_min]", sortBy.AfterDate.Format("2006-01-02"))
				q.Add("search[date_max]", time.Now().Format("2006-01-02"))
			}
		}

		parsedUrl.RawQuery = q.Encode()
	}

	statusCode, content, errs := client.Get(parsedUrl.String()).
		Set("Authorization", "Bearer "+ClientToken()).
		MaxRedirectsCount(5).
		Timeout(time.Second * 10).Bytes()
	if len(errs) > 0 || statusCode != 200 {
		return nil, fmt.Errorf("%v \nget failed, status code: %d, body: %s", parsedUrl.String(), statusCode, content)
	}

	return content, nil
}
