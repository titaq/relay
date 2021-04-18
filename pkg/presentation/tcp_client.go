package presentation

import (
	"net"

	"github.com/sirupsen/logrus"
)

type tcpClient struct {
	conn    net.Conn
	mailbox chan []byte

	id string
}

func (c *tcpClient) GetId() string {
	return c.id
}

func (c *tcpClient) Send(data []byte) error {
	_, err := c.conn.Write(data)
	if err != nil {
		logrus.Warn(err.Error())
	}
	logrus.Debugf("Sent a message to client %s", c.GetId())
	return err
}
