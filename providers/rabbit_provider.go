package providers

import (
	"HttpApiService/config"
	"HttpApiService/utils"
	"fmt"

	"github.com/google/uuid"
	logger "github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
	googleProto "google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

var (
	rabbitUrl = "amqp://%s:%s@%s:%s/"

	createdQueueMsg             = "HttpApiService created queue: %s in ex: %s with rk: %s\n"
	publishedMsg                = "Published message: to %s with %s"
	gotResponseMsg              = "HttpApiService got response: %+v"
	successfullyDeletedQueueMsg = "Successfully deleted queue: %s"
)

type RabbitProvider struct {
	connection *amqp.Connection
}

func NewRabbitProvider(envCfg *config.EnvConfig) *RabbitProvider {
	conn, err := amqp.Dial(fmt.Sprintf(rabbitUrl, envCfg.RabbitUser, envCfg.RabbitPassword, envCfg.RabbitHost, envCfg.RabbitPort))

	utils.CheckErrorWithPanic(err)
	rabbitProvider := &RabbitProvider{connection: conn}

	return rabbitProvider
}

func (r *RabbitProvider) getNewChannel() *amqp.Channel {
	ch, err := r.connection.Channel()
	utils.CheckErrorWithPanic(err)

	return ch
}

func (r *RabbitProvider) GetQueueConsumer(exName string, rk string, queueName string) (<-chan amqp.Delivery, *amqp.Channel) {
	ch := r.getNewChannel()
	_, err := ch.QueueDeclare(
		queueName,
		false,
		false,
		false,
		false,
		nil,
	)
	utils.CheckErrorWithPanic(err)

	err = ch.QueueBind(
		queueName,
		rk,
		exName,
		false,
		nil,
	)
	utils.CheckErrorWithPanic(err)

	msgs, err := ch.Consume(
		queueName,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	utils.CheckErrorWithPanic(err)

	logger.Debugf(createdQueueMsg, queueName, exName, rk)
	return msgs, ch
}

func (r *RabbitProvider) DeclareExchange(exName string) {
	err := r.getNewChannel().ExchangeDeclare(
		exName,
		"topic",
		false,
		false,
		false,
		false,
		nil,
	)
	utils.CheckErrorWithPanic(err)
}

func (r *RabbitProvider) SendMessage(exName string, rk string, message []byte) {
	err := r.getNewChannel().Publish(
		exName,
		rk,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		},
	)
	utils.CheckErrorWithPanic(err)
	logger.Infof(publishedMsg, exName, rk)
}

func (r *RabbitProvider) SendMessageAndHandleResponse(exName, rkSend, rkListen, queueName string, request []byte, response protoreflect.ProtoMessage) error {
	uniqueQueueName := queueName + "." + uuid.NewString()
	msgs, ch := r.GetQueueConsumer(exName, rkListen, uniqueQueueName)

	r.SendMessage(exName, rkSend, request)

	defer ch.Close()
	for msg := range msgs {
		err := googleProto.Unmarshal(msg.Body, response)
		if err != nil {
			return err
		}
		logger.Infof(gotResponseMsg, response)

		if _, err := ch.QueueDelete(uniqueQueueName, false, false, false); err != nil {
			return err
		}
		logger.Debugf(successfullyDeletedQueueMsg, uniqueQueueName)
		return nil
	}
	return nil
}
