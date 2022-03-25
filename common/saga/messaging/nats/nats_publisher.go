package nats

import (
	"github.com/nats-io/nats.go"
	saga "github.com/tamararankovic/microservices_demo/common/saga/messaging"
)

type NATSPublisher struct {
	conn    *nats.EncodedConn
	subject string
}

func NewNATSPublisher(host, port, user, password, subject string) (saga.Publisher, error) {
	conn, err := getConnection(host, port, user, password)
	encConn, err := nats.NewEncodedConn(conn, nats.JSON_ENCODER)
	if err != nil {
		return nil, err
	}
	return &NATSPublisher{
		conn:    encConn,
		subject: subject,
	}, nil
}

func (p *NATSPublisher) Publish(message interface{}) error {
	err := p.conn.Publish(p.subject, message)
	if err != nil {
		return err
	}
	return nil
}