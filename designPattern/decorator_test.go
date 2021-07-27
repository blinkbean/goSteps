package designPattern

import (
	"fmt"
	"math"
	"sync"
	"testing"
)

/**
装饰模式
 */

// 创建单例计数器
type metricCount struct {
	metrics map[string]uint64
	mu *sync.Mutex
}

func (m *metricCount) Inc(name string) {
	m.mu.Lock()
	if _, ok := m.metrics[name]; !ok {
		m.metrics[name] = 0
	}
	m.metrics[name] = m.metrics[name] + 1
	m.mu.Unlock()
}

func (m *metricCount) Show() {
	fmt.Printf("metric: %vln", m.metrics)
}

var metricInstance = &metricCount{
	metrics: make(map[string]uint64),
	mu:      &sync.Mutex{},
}

func MetricCount() *metricCount {
	return metricInstance
}


type DecoratorInput interface {
	DecoratorReceive(string)
}

type DecoratorHelloInput struct {}
func (d *DecoratorHelloInput) DecoratorReceive(name string) {
	fmt.Printf("DecoratorReceive: %v\n", name)
}


type InputMetric struct {
	input DecoratorInput
}

func (i *InputMetric) DecoratorReceive(name string)  {
	i.input.DecoratorReceive(name)
	MetricCount().Inc(name)
}

func CreateMetricDecorator(input DecoratorInput) *InputMetric {
	return &InputMetric{input: input}
}

type MyInput struct {
	input DecoratorInput
}

func TestMetricDecorator(t *testing.T) {
	aa := math.Sqrt(40000000)
	fmt.Println(aa)
	myInput := MyInput{input: &DecoratorHelloInput{}}
	myInput.input.DecoratorReceive("hello")
	metricInstance.Show()
	// 装饰
	myInput.input = CreateMetricDecorator(myInput.input)
	myInput.input.DecoratorReceive("hello")
	myInput.input.DecoratorReceive("hello")
	metricInstance.Show()
	myInput.input.DecoratorReceive("input")
	metricInstance.Show()
}