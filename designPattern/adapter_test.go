package designPattern

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

/**
适配器模式
*/

type Records struct {
	Items []string
}

type Consumer interface {
	Poll() Records
}

type KafkaInput struct {
	status   Status
	consumer Consumer
}

func (k *KafkaInput) Receive() string {
	records := k.consumer.Poll()
	if k.status != Started {
		fmt.Println("Kafka input plugin is not running, input nothing.")
		return ""
	}
	return strings.Join(records.Items,",")
}

func (k *KafkaInput) Start() {
	k.status = Started
	fmt.Println("Kafka input plugin started.")
}


func (k *KafkaInput) Stop() {
	k.status = Stopped
	fmt.Println("Kafka input plugin stopped.")
}

func (k *KafkaInput) Status() Status {
	return k.status
}


type MockConsumer struct {}

func (m *MockConsumer) Poll() Records {
	records := Records{}
	records.Items = append(records.Items, "i am mock consumer.")
	return records
}

func (k *KafkaInput) Init() {
	k.consumer = &MockConsumer{}
}

func init() {
	inputNames["kafka"] = reflect.TypeOf(KafkaInput{})
}

func Test_AdapterPipeline(t *testing.T) {
	conf := PipelineConfig {
		Input: Config{
			Name:       "kafka",
			PluginType: InputType,
		},
		Filter: Config{
			Name:       "upper",
			PluginType: FilterType,
		},
		Output: Config{
			Name:       "console",
			PluginType: OutputType,
		},
	}
	p := Of(conf)
	p.Start()
	p.Exec()
	p.Stop()
}