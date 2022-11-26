package main

import (
	"errors"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"net/http"
)

type options struct {
	port *int
}

type Option func(option *options) error

func WithPort(port int) Option { //Юзаем замыкание
	return func(opt *options) error {
		if port < 0 {
			return errors.New("negative port")
		}
		opt.port = &port

		return nil
	}
}

func NewServer1(host string, opts ...Option) (*http.Server, error) {
	var err error
	var optionsObj options

	for _, opt := range opts {
		err = opt(&optionsObj)
		if err != nil {
			return nil, err
		}
	}

	var port int
	optionsObj.portValidation(&port)

	return &http.Server{Addr: fmt.Sprintf("%s:%d", host, port)}, nil
}

func main() {
	server, err := NewServer1("localhost", WithPort(10)) //Красиво получается
	spew.Dump(server, err)
}

func (opt options) portValidation(port *int) {
	if opt.port == nil {
		*port = 80
	} else {
		if *opt.port == 0 {
			*port = 63287
		} else {
			*port = *opt.port
		}
	}
}
