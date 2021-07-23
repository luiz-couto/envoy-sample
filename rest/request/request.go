package request

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

// Send sends a request to the url specified in instantiation, with the given
// route and method, using
// the encoder to encode the body and the decoder to decode the response into
// the responsePtr
func (c Client) Send(ctx context.Context, route, method, hostname string, body, responsePtr interface{}) (err error) {
	buf, err := c.encoder(body)
	if err != nil {
		return err
	}

	req, err := http.NewRequestWithContext(
		ctx,
		method,
		c.routeToURL(route),
		bytes.NewBuffer(buf),
	)

	// Specifies the hostname of the request, default is inspr.com
	// it is useful for cluster ingresses reverse DNS
	req.Host = hostname

	if err != nil {
		return err
	}

	for key, values := range c.headers {
		req.Header[key] = values
	}

	resp, err := c.c.Do(req)
	if err != nil {
		return err
	}

	err = c.handleResponseErr(resp)
	if err != nil {
		return err
	}

	if responsePtr != nil {
		decoder := json.NewDecoder(resp.Body)
		err = decoder.Decode(responsePtr)

		if err == io.EOF {
			return nil
		}
	}

	return err
}

func (c Client) handleResponseErr(resp *http.Response) error {
	decoder := c.decoderGenerator(resp.Body)
	var err error
	defaultErr := errors.New("server Error")

	switch resp.StatusCode {
	case http.StatusOK:
		return nil
	case http.StatusUnauthorized:
		decoder.Decode(&err)
		if err != nil {
			return errors.New("Unauthorized")
		}
		return defaultErr
	case http.StatusForbidden:
		decoder.Decode(&err)
		if err != nil {
			return errors.New("Forbidden")
		}
		return defaultErr
	case http.StatusNotFound:
		return errors.New("StatusNotFound")

	default:
		decoder.Decode(&err)
		if err == nil {
			return defaultErr
		}
		return err
	}
}
