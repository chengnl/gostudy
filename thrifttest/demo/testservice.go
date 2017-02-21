// Autogenerated by Thrift Compiler (0.9.3)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package demo

import (
	"bytes"
	"fmt"

	"git.apache.org/thrift.git/lib/go/thrift"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = bytes.Equal

type TestService interface {
	HelloWorld() (r string, err error)
	// Parameters:
	//  - Name
	HelloWorldForString(name string) (r string, err error)
	// Parameters:
	//  - Name
	HelloWorldForMap(name map[string]int32) (r string, err error)
	// Parameters:
	//  - Name
	HelloWorldForStruct(name *Name) (r string, err error)
}

type TestServiceClient struct {
	Transport       thrift.TTransport
	ProtocolFactory thrift.TProtocolFactory
	InputProtocol   thrift.TProtocol
	OutputProtocol  thrift.TProtocol
	SeqId           int32
}

func NewTestServiceClientFactory(t thrift.TTransport, f thrift.TProtocolFactory) *TestServiceClient {
	return &TestServiceClient{Transport: t,
		ProtocolFactory: f,
		InputProtocol:   f.GetProtocol(t),
		OutputProtocol:  f.GetProtocol(t),
		SeqId:           0,
	}
}

func NewTestServiceClientProtocol(t thrift.TTransport, iprot thrift.TProtocol, oprot thrift.TProtocol) *TestServiceClient {
	return &TestServiceClient{Transport: t,
		ProtocolFactory: nil,
		InputProtocol:   iprot,
		OutputProtocol:  oprot,
		SeqId:           0,
	}
}

func (p *TestServiceClient) HelloWorld() (r string, err error) {
	if err = p.sendHelloWorld(); err != nil {
		return
	}
	return p.recvHelloWorld()
}

func (p *TestServiceClient) sendHelloWorld() (err error) {
	oprot := p.OutputProtocol
	if oprot == nil {
		oprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.OutputProtocol = oprot
	}
	p.SeqId++
	if err = oprot.WriteMessageBegin("HelloWorld", thrift.CALL, p.SeqId); err != nil {
		return
	}
	args := TestServiceHelloWorldArgs{}
	if err = args.Write(oprot); err != nil {
		return
	}
	if err = oprot.WriteMessageEnd(); err != nil {
		return
	}
	return oprot.Flush()
}

func (p *TestServiceClient) recvHelloWorld() (value string, err error) {
	iprot := p.InputProtocol
	if iprot == nil {
		iprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.InputProtocol = iprot
	}
	method, mTypeId, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return
	}
	if method != "HelloWorld" {
		err = thrift.NewTApplicationException(thrift.WRONG_METHOD_NAME, "HelloWorld failed: wrong method name")
		return
	}
	if p.SeqId != seqId {
		err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "HelloWorld failed: out of sequence response")
		return
	}
	if mTypeId == thrift.EXCEPTION {
		error0 := thrift.NewTApplicationException(thrift.UNKNOWN_APPLICATION_EXCEPTION, "Unknown Exception")
		var error1 error
		error1, err = error0.Read(iprot)
		if err != nil {
			return
		}
		if err = iprot.ReadMessageEnd(); err != nil {
			return
		}
		err = error1
		return
	}
	if mTypeId != thrift.REPLY {
		err = thrift.NewTApplicationException(thrift.INVALID_MESSAGE_TYPE_EXCEPTION, "HelloWorld failed: invalid message type")
		return
	}
	result := TestServiceHelloWorldResult{}
	if err = result.Read(iprot); err != nil {
		return
	}
	if err = iprot.ReadMessageEnd(); err != nil {
		return
	}
	value = result.GetSuccess()
	return
}

// Parameters:
//  - Name
func (p *TestServiceClient) HelloWorldForString(name string) (r string, err error) {
	if err = p.sendHelloWorldForString(name); err != nil {
		return
	}
	return p.recvHelloWorldForString()
}

func (p *TestServiceClient) sendHelloWorldForString(name string) (err error) {
	oprot := p.OutputProtocol
	if oprot == nil {
		oprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.OutputProtocol = oprot
	}
	p.SeqId++
	if err = oprot.WriteMessageBegin("HelloWorldForString", thrift.CALL, p.SeqId); err != nil {
		return
	}
	args := TestServiceHelloWorldForStringArgs{
		Name: name,
	}
	if err = args.Write(oprot); err != nil {
		return
	}
	if err = oprot.WriteMessageEnd(); err != nil {
		return
	}
	return oprot.Flush()
}

func (p *TestServiceClient) recvHelloWorldForString() (value string, err error) {
	iprot := p.InputProtocol
	if iprot == nil {
		iprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.InputProtocol = iprot
	}
	method, mTypeId, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return
	}
	if method != "HelloWorldForString" {
		err = thrift.NewTApplicationException(thrift.WRONG_METHOD_NAME, "HelloWorldForString failed: wrong method name")
		return
	}
	if p.SeqId != seqId {
		err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "HelloWorldForString failed: out of sequence response")
		return
	}
	if mTypeId == thrift.EXCEPTION {
		error2 := thrift.NewTApplicationException(thrift.UNKNOWN_APPLICATION_EXCEPTION, "Unknown Exception")
		var error3 error
		error3, err = error2.Read(iprot)
		if err != nil {
			return
		}
		if err = iprot.ReadMessageEnd(); err != nil {
			return
		}
		err = error3
		return
	}
	if mTypeId != thrift.REPLY {
		err = thrift.NewTApplicationException(thrift.INVALID_MESSAGE_TYPE_EXCEPTION, "HelloWorldForString failed: invalid message type")
		return
	}
	result := TestServiceHelloWorldForStringResult{}
	if err = result.Read(iprot); err != nil {
		return
	}
	if err = iprot.ReadMessageEnd(); err != nil {
		return
	}
	value = result.GetSuccess()
	return
}

// Parameters:
//  - Name
func (p *TestServiceClient) HelloWorldForMap(name map[string]int32) (r string, err error) {
	if err = p.sendHelloWorldForMap(name); err != nil {
		return
	}
	return p.recvHelloWorldForMap()
}

func (p *TestServiceClient) sendHelloWorldForMap(name map[string]int32) (err error) {
	oprot := p.OutputProtocol
	if oprot == nil {
		oprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.OutputProtocol = oprot
	}
	p.SeqId++
	if err = oprot.WriteMessageBegin("HelloWorldForMap", thrift.CALL, p.SeqId); err != nil {
		return
	}
	args := TestServiceHelloWorldForMapArgs{
		Name: name,
	}
	if err = args.Write(oprot); err != nil {
		return
	}
	if err = oprot.WriteMessageEnd(); err != nil {
		return
	}
	return oprot.Flush()
}

func (p *TestServiceClient) recvHelloWorldForMap() (value string, err error) {
	iprot := p.InputProtocol
	if iprot == nil {
		iprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.InputProtocol = iprot
	}
	method, mTypeId, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return
	}
	if method != "HelloWorldForMap" {
		err = thrift.NewTApplicationException(thrift.WRONG_METHOD_NAME, "HelloWorldForMap failed: wrong method name")
		return
	}
	if p.SeqId != seqId {
		err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "HelloWorldForMap failed: out of sequence response")
		return
	}
	if mTypeId == thrift.EXCEPTION {
		error4 := thrift.NewTApplicationException(thrift.UNKNOWN_APPLICATION_EXCEPTION, "Unknown Exception")
		var error5 error
		error5, err = error4.Read(iprot)
		if err != nil {
			return
		}
		if err = iprot.ReadMessageEnd(); err != nil {
			return
		}
		err = error5
		return
	}
	if mTypeId != thrift.REPLY {
		err = thrift.NewTApplicationException(thrift.INVALID_MESSAGE_TYPE_EXCEPTION, "HelloWorldForMap failed: invalid message type")
		return
	}
	result := TestServiceHelloWorldForMapResult{}
	if err = result.Read(iprot); err != nil {
		return
	}
	if err = iprot.ReadMessageEnd(); err != nil {
		return
	}
	value = result.GetSuccess()
	return
}

// Parameters:
//  - Name
func (p *TestServiceClient) HelloWorldForStruct(name *Name) (r string, err error) {
	if err = p.sendHelloWorldForStruct(name); err != nil {
		return
	}
	return p.recvHelloWorldForStruct()
}

func (p *TestServiceClient) sendHelloWorldForStruct(name *Name) (err error) {
	oprot := p.OutputProtocol
	if oprot == nil {
		oprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.OutputProtocol = oprot
	}
	p.SeqId++
	if err = oprot.WriteMessageBegin("HelloWorldForStruct", thrift.CALL, p.SeqId); err != nil {
		return
	}
	args := TestServiceHelloWorldForStructArgs{
		Name: name,
	}
	if err = args.Write(oprot); err != nil {
		return
	}
	if err = oprot.WriteMessageEnd(); err != nil {
		return
	}
	return oprot.Flush()
}

func (p *TestServiceClient) recvHelloWorldForStruct() (value string, err error) {
	iprot := p.InputProtocol
	if iprot == nil {
		iprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.InputProtocol = iprot
	}
	method, mTypeId, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return
	}
	if method != "HelloWorldForStruct" {
		err = thrift.NewTApplicationException(thrift.WRONG_METHOD_NAME, "HelloWorldForStruct failed: wrong method name")
		return
	}
	if p.SeqId != seqId {
		err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "HelloWorldForStruct failed: out of sequence response")
		return
	}
	if mTypeId == thrift.EXCEPTION {
		error6 := thrift.NewTApplicationException(thrift.UNKNOWN_APPLICATION_EXCEPTION, "Unknown Exception")
		var error7 error
		error7, err = error6.Read(iprot)
		if err != nil {
			return
		}
		if err = iprot.ReadMessageEnd(); err != nil {
			return
		}
		err = error7
		return
	}
	if mTypeId != thrift.REPLY {
		err = thrift.NewTApplicationException(thrift.INVALID_MESSAGE_TYPE_EXCEPTION, "HelloWorldForStruct failed: invalid message type")
		return
	}
	result := TestServiceHelloWorldForStructResult{}
	if err = result.Read(iprot); err != nil {
		return
	}
	if err = iprot.ReadMessageEnd(); err != nil {
		return
	}
	value = result.GetSuccess()
	return
}

type TestServiceProcessor struct {
	processorMap map[string]thrift.TProcessorFunction
	handler      TestService
}

func (p *TestServiceProcessor) AddToProcessorMap(key string, processor thrift.TProcessorFunction) {
	p.processorMap[key] = processor
}

func (p *TestServiceProcessor) GetProcessorFunction(key string) (processor thrift.TProcessorFunction, ok bool) {
	processor, ok = p.processorMap[key]
	return processor, ok
}

func (p *TestServiceProcessor) ProcessorMap() map[string]thrift.TProcessorFunction {
	return p.processorMap
}

func NewTestServiceProcessor(handler TestService) *TestServiceProcessor {

	self8 := &TestServiceProcessor{handler: handler, processorMap: make(map[string]thrift.TProcessorFunction)}
	self8.processorMap["HelloWorld"] = &testServiceProcessorHelloWorld{handler: handler}
	self8.processorMap["HelloWorldForString"] = &testServiceProcessorHelloWorldForString{handler: handler}
	self8.processorMap["HelloWorldForMap"] = &testServiceProcessorHelloWorldForMap{handler: handler}
	self8.processorMap["HelloWorldForStruct"] = &testServiceProcessorHelloWorldForStruct{handler: handler}
	return self8
}

func (p *TestServiceProcessor) Process(iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	name, _, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return false, err
	}
	if processor, ok := p.GetProcessorFunction(name); ok {
		return processor.Process(seqId, iprot, oprot)
	}
	iprot.Skip(thrift.STRUCT)
	iprot.ReadMessageEnd()
	x9 := thrift.NewTApplicationException(thrift.UNKNOWN_METHOD, "Unknown function "+name)
	oprot.WriteMessageBegin(name, thrift.EXCEPTION, seqId)
	x9.Write(oprot)
	oprot.WriteMessageEnd()
	oprot.Flush()
	return false, x9

}

type testServiceProcessorHelloWorld struct {
	handler TestService
}

func (p *testServiceProcessorHelloWorld) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := TestServiceHelloWorldArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("HelloWorld", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return false, err
	}

	iprot.ReadMessageEnd()
	result := TestServiceHelloWorldResult{}
	var retval string
	var err2 error
	if retval, err2 = p.handler.HelloWorld(); err2 != nil {
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing HelloWorld: "+err2.Error())
		oprot.WriteMessageBegin("HelloWorld", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return true, err2
	} else {
		result.Success = &retval
	}
	if err2 = oprot.WriteMessageBegin("HelloWorld", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

type testServiceProcessorHelloWorldForString struct {
	handler TestService
}

func (p *testServiceProcessorHelloWorldForString) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := TestServiceHelloWorldForStringArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("HelloWorldForString", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return false, err
	}

	iprot.ReadMessageEnd()
	result := TestServiceHelloWorldForStringResult{}
	var retval string
	var err2 error
	if retval, err2 = p.handler.HelloWorldForString(args.Name); err2 != nil {
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing HelloWorldForString: "+err2.Error())
		oprot.WriteMessageBegin("HelloWorldForString", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return true, err2
	} else {
		result.Success = &retval
	}
	if err2 = oprot.WriteMessageBegin("HelloWorldForString", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

type testServiceProcessorHelloWorldForMap struct {
	handler TestService
}

func (p *testServiceProcessorHelloWorldForMap) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := TestServiceHelloWorldForMapArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("HelloWorldForMap", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return false, err
	}

	iprot.ReadMessageEnd()
	result := TestServiceHelloWorldForMapResult{}
	var retval string
	var err2 error
	if retval, err2 = p.handler.HelloWorldForMap(args.Name); err2 != nil {
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing HelloWorldForMap: "+err2.Error())
		oprot.WriteMessageBegin("HelloWorldForMap", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return true, err2
	} else {
		result.Success = &retval
	}
	if err2 = oprot.WriteMessageBegin("HelloWorldForMap", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

type testServiceProcessorHelloWorldForStruct struct {
	handler TestService
}

func (p *testServiceProcessorHelloWorldForStruct) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := TestServiceHelloWorldForStructArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("HelloWorldForStruct", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return false, err
	}

	iprot.ReadMessageEnd()
	result := TestServiceHelloWorldForStructResult{}
	var retval string
	var err2 error
	if retval, err2 = p.handler.HelloWorldForStruct(args.Name); err2 != nil {
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing HelloWorldForStruct: "+err2.Error())
		oprot.WriteMessageBegin("HelloWorldForStruct", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return true, err2
	} else {
		result.Success = &retval
	}
	if err2 = oprot.WriteMessageBegin("HelloWorldForStruct", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

// HELPER FUNCTIONS AND STRUCTURES

type TestServiceHelloWorldArgs struct {
}

func NewTestServiceHelloWorldArgs() *TestServiceHelloWorldArgs {
	return &TestServiceHelloWorldArgs{}
}

func (p *TestServiceHelloWorldArgs) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		if err := iprot.Skip(fieldTypeId); err != nil {
			return err
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *TestServiceHelloWorldArgs) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("HelloWorld_args"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *TestServiceHelloWorldArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("TestServiceHelloWorldArgs(%+v)", *p)
}

// Attributes:
//  - Success
type TestServiceHelloWorldResult struct {
	Success *string `thrift:"success,0" json:"success,omitempty"`
}

func NewTestServiceHelloWorldResult() *TestServiceHelloWorldResult {
	return &TestServiceHelloWorldResult{}
}

var TestServiceHelloWorldResult_Success_DEFAULT string

func (p *TestServiceHelloWorldResult) GetSuccess() string {
	if !p.IsSetSuccess() {
		return TestServiceHelloWorldResult_Success_DEFAULT
	}
	return *p.Success
}
func (p *TestServiceHelloWorldResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *TestServiceHelloWorldResult) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 0:
			if err := p.readField0(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *TestServiceHelloWorldResult) readField0(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 0: ", err)
	} else {
		p.Success = &v
	}
	return nil
}

func (p *TestServiceHelloWorldResult) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("HelloWorld_result"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField0(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *TestServiceHelloWorldResult) writeField0(oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err := oprot.WriteFieldBegin("success", thrift.STRING, 0); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err)
		}
		if err := oprot.WriteString(string(*p.Success)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.success (0) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err)
		}
	}
	return err
}

func (p *TestServiceHelloWorldResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("TestServiceHelloWorldResult(%+v)", *p)
}

// Attributes:
//  - Name
type TestServiceHelloWorldForStringArgs struct {
	Name string `thrift:"name,1" json:"name"`
}

func NewTestServiceHelloWorldForStringArgs() *TestServiceHelloWorldForStringArgs {
	return &TestServiceHelloWorldForStringArgs{}
}

func (p *TestServiceHelloWorldForStringArgs) GetName() string {
	return p.Name
}
func (p *TestServiceHelloWorldForStringArgs) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.readField1(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *TestServiceHelloWorldForStringArgs) readField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.Name = v
	}
	return nil
}

func (p *TestServiceHelloWorldForStringArgs) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("HelloWorldForString_args"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *TestServiceHelloWorldForStringArgs) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("name", thrift.STRING, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:name: ", p), err)
	}
	if err := oprot.WriteString(string(p.Name)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.name (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:name: ", p), err)
	}
	return err
}

func (p *TestServiceHelloWorldForStringArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("TestServiceHelloWorldForStringArgs(%+v)", *p)
}

// Attributes:
//  - Success
type TestServiceHelloWorldForStringResult struct {
	Success *string `thrift:"success,0" json:"success,omitempty"`
}

func NewTestServiceHelloWorldForStringResult() *TestServiceHelloWorldForStringResult {
	return &TestServiceHelloWorldForStringResult{}
}

var TestServiceHelloWorldForStringResult_Success_DEFAULT string

func (p *TestServiceHelloWorldForStringResult) GetSuccess() string {
	if !p.IsSetSuccess() {
		return TestServiceHelloWorldForStringResult_Success_DEFAULT
	}
	return *p.Success
}
func (p *TestServiceHelloWorldForStringResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *TestServiceHelloWorldForStringResult) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 0:
			if err := p.readField0(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *TestServiceHelloWorldForStringResult) readField0(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 0: ", err)
	} else {
		p.Success = &v
	}
	return nil
}

func (p *TestServiceHelloWorldForStringResult) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("HelloWorldForString_result"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField0(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *TestServiceHelloWorldForStringResult) writeField0(oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err := oprot.WriteFieldBegin("success", thrift.STRING, 0); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err)
		}
		if err := oprot.WriteString(string(*p.Success)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.success (0) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err)
		}
	}
	return err
}

func (p *TestServiceHelloWorldForStringResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("TestServiceHelloWorldForStringResult(%+v)", *p)
}

// Attributes:
//  - Name
type TestServiceHelloWorldForMapArgs struct {
	Name map[string]int32 `thrift:"name,1" json:"name"`
}

func NewTestServiceHelloWorldForMapArgs() *TestServiceHelloWorldForMapArgs {
	return &TestServiceHelloWorldForMapArgs{}
}

func (p *TestServiceHelloWorldForMapArgs) GetName() map[string]int32 {
	return p.Name
}
func (p *TestServiceHelloWorldForMapArgs) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.readField1(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *TestServiceHelloWorldForMapArgs) readField1(iprot thrift.TProtocol) error {
	_, _, size, err := iprot.ReadMapBegin()
	if err != nil {
		return thrift.PrependError("error reading map begin: ", err)
	}
	tMap := make(map[string]int32, size)
	p.Name = tMap
	for i := 0; i < size; i++ {
		var _key10 string
		if v, err := iprot.ReadString(); err != nil {
			return thrift.PrependError("error reading field 0: ", err)
		} else {
			_key10 = v
		}
		var _val11 int32
		if v, err := iprot.ReadI32(); err != nil {
			return thrift.PrependError("error reading field 0: ", err)
		} else {
			_val11 = v
		}
		p.Name[_key10] = _val11
	}
	if err := iprot.ReadMapEnd(); err != nil {
		return thrift.PrependError("error reading map end: ", err)
	}
	return nil
}

func (p *TestServiceHelloWorldForMapArgs) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("HelloWorldForMap_args"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *TestServiceHelloWorldForMapArgs) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("name", thrift.MAP, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:name: ", p), err)
	}
	if err := oprot.WriteMapBegin(thrift.STRING, thrift.I32, len(p.Name)); err != nil {
		return thrift.PrependError("error writing map begin: ", err)
	}
	for k, v := range p.Name {
		if err := oprot.WriteString(string(k)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T. (0) field write error: ", p), err)
		}
		if err := oprot.WriteI32(int32(v)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T. (0) field write error: ", p), err)
		}
	}
	if err := oprot.WriteMapEnd(); err != nil {
		return thrift.PrependError("error writing map end: ", err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:name: ", p), err)
	}
	return err
}

func (p *TestServiceHelloWorldForMapArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("TestServiceHelloWorldForMapArgs(%+v)", *p)
}

// Attributes:
//  - Success
type TestServiceHelloWorldForMapResult struct {
	Success *string `thrift:"success,0" json:"success,omitempty"`
}

func NewTestServiceHelloWorldForMapResult() *TestServiceHelloWorldForMapResult {
	return &TestServiceHelloWorldForMapResult{}
}

var TestServiceHelloWorldForMapResult_Success_DEFAULT string

func (p *TestServiceHelloWorldForMapResult) GetSuccess() string {
	if !p.IsSetSuccess() {
		return TestServiceHelloWorldForMapResult_Success_DEFAULT
	}
	return *p.Success
}
func (p *TestServiceHelloWorldForMapResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *TestServiceHelloWorldForMapResult) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 0:
			if err := p.readField0(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *TestServiceHelloWorldForMapResult) readField0(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 0: ", err)
	} else {
		p.Success = &v
	}
	return nil
}

func (p *TestServiceHelloWorldForMapResult) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("HelloWorldForMap_result"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField0(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *TestServiceHelloWorldForMapResult) writeField0(oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err := oprot.WriteFieldBegin("success", thrift.STRING, 0); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err)
		}
		if err := oprot.WriteString(string(*p.Success)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.success (0) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err)
		}
	}
	return err
}

func (p *TestServiceHelloWorldForMapResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("TestServiceHelloWorldForMapResult(%+v)", *p)
}

// Attributes:
//  - Name
type TestServiceHelloWorldForStructArgs struct {
	Name *Name `thrift:"name,1" json:"name"`
}

func NewTestServiceHelloWorldForStructArgs() *TestServiceHelloWorldForStructArgs {
	return &TestServiceHelloWorldForStructArgs{}
}

var TestServiceHelloWorldForStructArgs_Name_DEFAULT *Name

func (p *TestServiceHelloWorldForStructArgs) GetName() *Name {
	if !p.IsSetName() {
		return TestServiceHelloWorldForStructArgs_Name_DEFAULT
	}
	return p.Name
}
func (p *TestServiceHelloWorldForStructArgs) IsSetName() bool {
	return p.Name != nil
}

func (p *TestServiceHelloWorldForStructArgs) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.readField1(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *TestServiceHelloWorldForStructArgs) readField1(iprot thrift.TProtocol) error {
	p.Name = &Name{}
	if err := p.Name.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Name), err)
	}
	return nil
}

func (p *TestServiceHelloWorldForStructArgs) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("HelloWorldForStruct_args"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *TestServiceHelloWorldForStructArgs) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("name", thrift.STRUCT, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:name: ", p), err)
	}
	if err := p.Name.Write(oprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Name), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:name: ", p), err)
	}
	return err
}

func (p *TestServiceHelloWorldForStructArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("TestServiceHelloWorldForStructArgs(%+v)", *p)
}

// Attributes:
//  - Success
type TestServiceHelloWorldForStructResult struct {
	Success *string `thrift:"success,0" json:"success,omitempty"`
}

func NewTestServiceHelloWorldForStructResult() *TestServiceHelloWorldForStructResult {
	return &TestServiceHelloWorldForStructResult{}
}

var TestServiceHelloWorldForStructResult_Success_DEFAULT string

func (p *TestServiceHelloWorldForStructResult) GetSuccess() string {
	if !p.IsSetSuccess() {
		return TestServiceHelloWorldForStructResult_Success_DEFAULT
	}
	return *p.Success
}
func (p *TestServiceHelloWorldForStructResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *TestServiceHelloWorldForStructResult) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 0:
			if err := p.readField0(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *TestServiceHelloWorldForStructResult) readField0(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 0: ", err)
	} else {
		p.Success = &v
	}
	return nil
}

func (p *TestServiceHelloWorldForStructResult) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("HelloWorldForStruct_result"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField0(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *TestServiceHelloWorldForStructResult) writeField0(oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err := oprot.WriteFieldBegin("success", thrift.STRING, 0); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err)
		}
		if err := oprot.WriteString(string(*p.Success)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.success (0) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err)
		}
	}
	return err
}

func (p *TestServiceHelloWorldForStructResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("TestServiceHelloWorldForStructResult(%+v)", *p)
}
