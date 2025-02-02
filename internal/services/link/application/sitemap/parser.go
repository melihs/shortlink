package sitemap

import (
	"context"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"

	"github.com/gogo/protobuf/proto"

	http_client "github.com/shortlink-org/shortlink/internal/pkg/http/client"
	"github.com/shortlink-org/shortlink/internal/pkg/logger"
	"github.com/shortlink-org/shortlink/internal/pkg/mq"
	link "github.com/shortlink-org/shortlink/internal/services/link/domain/link/v1"
	domain "github.com/shortlink-org/shortlink/internal/services/link/domain/sitemap/v1"
)

type Service struct {
	log logger.Logger

	// Delivery
	mq mq.MQ
}

func New(log logger.Logger, dataBus mq.MQ) (*Service, error) {
	service := &Service{
		log: log,

		// Delivery
		mq: dataBus,
	}

	return service, nil
}

func (s *Service) Parse(ctx context.Context, url string) error {
	// Request the HTML page.
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, http.NoBody)
	if err != nil {
		return err
	}

	client := http_client.New()

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf(`incorrect response code: %d for %s`, resp.StatusCode, url)
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var payload domain.Sitemap
	err = xml.Unmarshal(bodyBytes, &payload)
	if err != nil {
		return err
	}

	// send to link_rpc.add
	for key := range payload.GetUrl() {
		data, errMarshal := proto.Marshal(&link.Link{Url: payload.GetUrl()[key].GetLoc()})
		if errMarshal != nil {
			return errMarshal
		}

		errPublish := s.mq.Publish(ctx, link.MQ_EVENT_LINK_NEW, nil, data)
		if errPublish != nil {
			return errPublish
		}
	}

	return nil
}
