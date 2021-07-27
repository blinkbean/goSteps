package designPattern

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

type Plugin interface {
	Start()
	Stop()
	Status() Status
}

type Input interface {
	Plugin
	Receive() string
}

type Filter interface {
	Plugin
	Process(msg string) string
}

type Output interface {
	Plugin
	Send(msg string)
}

type Pipeline struct {
	status Status
	input  Input
	filter Filter
	output Output
}

func (p *Pipeline) Exec() {
	msg := p.input.Receive()
	msg = p.filter.Process(msg)
	p.output.Send(msg)
}

func (p *Pipeline) Start() {
	p.output.Start()
	p.filter.Start()
	p.input.Start()
	p.status = Started
	fmt.Println("Hello input plugin started.")
}

func (p *Pipeline) Stop() {
	p.input.Stop()
	p.filter.Stop()
	p.output.Stop()
	p.status = Stopped
	fmt.Println("Hello input plugin stopped.")
}

func (p *Pipeline) Status() Status {
	return p.status
}

var inputNames = make(map[string]reflect.Type)
var filterNames = make(map[string]reflect.Type)
var outputNames = make(map[string]reflect.Type)

func init() {
	inputNames["hello"] = reflect.TypeOf(HelloInput{})
	filterNames["upper"] = reflect.TypeOf(UpperFilter{})
	outputNames["console"] = reflect.TypeOf(ConsoleOutPut{})
}

type HelloInput struct {
	status Status
}

func (h *HelloInput) Receive() string {
	return "Hello World"
}

func (h *HelloInput) Start() {
	h.status = Started
	fmt.Println("Hello input plugin started.")
}

func (h *HelloInput) Stop() {
	h.status = Stopped
	fmt.Println("Hello input plugin stopped.")
}

func (h *HelloInput) Status() Status {
	return h.status
}

type UpperFilter struct {
	status Status
}

func (u *UpperFilter) Process(msg string) string {
	return strings.ToUpper(msg)
}

func (u *UpperFilter) Start() {
	u.status = Started
	fmt.Println("Upper input plugin started.")
}

func (u *UpperFilter) Stop() {
	u.status = Stopped
	fmt.Println("Upper input plugin stopped.")
}

func (u *UpperFilter) Status() Status {
	return u.status
}

type ConsoleOutPut struct {
	status Status
}

func (c *ConsoleOutPut) Send(msg string) {
	fmt.Println(msg)
}

func (c *ConsoleOutPut) Start() {
	c.status = Started
	fmt.Println("Console input plugin started.")
}

func (c *ConsoleOutPut) Stop() {
	c.status = Stopped
	fmt.Println("Console input plugin stopped.")
}

func (c *ConsoleOutPut) Status() Status {
	return c.status
}

type Config struct {
	PluginType PluginType
	Name       string
}

type PipelineConfig struct {
	Input  Config
	Filter Config
	Output Config
}

// 抽象工厂接口
type AbstractFactory interface {
	Create(conf Config) Plugin
}

// input工厂对象
type InputFactory struct{}

// 通过反射进行对象实例化
func (i *InputFactory) Create(conf Config) Plugin {
	t, _ := inputNames[conf.Name]
	return reflect.New(t).Interface().(Plugin)
}

type FilterFactory struct{}

func (f FilterFactory) Create(conf Config) Plugin {
	t, _ := filterNames[conf.Name]
	return reflect.New(t).Interface().(Plugin)
}

type OutputFactory struct{}

func (o *OutputFactory) Create(conf Config) Plugin {
	t, _ := outputNames[conf.Name]
	return reflect.New(t).Interface().(Plugin)
}

type PluginType int

const (
	InputType PluginType = iota
	FilterType
	OutputType
)

var pluginFactories = make(map[PluginType]AbstractFactory)

func factoryOf(t PluginType) AbstractFactory {
	factory, _ := pluginFactories[t]
	return factory
}

// pipeline工厂方法，根据配置创建一个Pipeline实例
func Of(conf PipelineConfig) *Pipeline {
	p := &Pipeline{}
	p.input = factoryOf(conf.Input.PluginType).Create(conf.Input).(Input)
	p.filter = factoryOf(conf.Filter.PluginType).Create(conf.Filter).(Filter)
	p.output = factoryOf(conf.Output.PluginType).Create(conf.Output).(Output)
	return p
}

// 初始化插件工厂对象
func init() {
	pluginFactories[InputType] = &InputFactory{}
	pluginFactories[FilterType] = &FilterFactory{}
	pluginFactories[OutputType] = &OutputFactory{}
}

func Test_Pipeline(t *testing.T) {
	conf := PipelineConfig {
		Input: Config{
			Name:       "hello",
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
