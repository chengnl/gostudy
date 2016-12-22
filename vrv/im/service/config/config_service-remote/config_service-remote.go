// Autogenerated by Thrift Compiler (0.9.3)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package main

import (
	"flag"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"math"
	"net"
	"net/url"
	"os"
	"strconv"
	"strings"
	"vrv/im/service/config"
)

func Usage() {
	fmt.Fprintln(os.Stderr, "Usage of ", os.Args[0], " [-h host:port] [-u url] [-f[ramed]] function [arg1 [arg2...]]:")
	flag.PrintDefaults()
	fmt.Fprintln(os.Stderr, "\nFunctions:")
	fmt.Fprintln(os.Stderr, "  ServiceConfigResult registerService(ServiceConfigBean config)")
	fmt.Fprintln(os.Stderr, "  byte serviceHeart(string ID)")
	fmt.Fprintln(os.Stderr, "   loadServices()")
	fmt.Fprintln(os.Stderr, "   queryService(string serviceID, string version)")
	fmt.Fprintln(os.Stderr, "  string getName()")
	fmt.Fprintln(os.Stderr, "  string getVersion()")
	fmt.Fprintln(os.Stderr, "   getServiceBizMethods()")
	fmt.Fprintln(os.Stderr, "   getBizMethodsInvokeInfo()")
	fmt.Fprintln(os.Stderr, "  BizMethodInvokeInfo getBizMethodInvokeInfo(string methodName)")
	fmt.Fprintln(os.Stderr, "  i64 getServiceCount()")
	fmt.Fprintln(os.Stderr, "  i64 aliveSince()")
	fmt.Fprintln(os.Stderr, "  void reinitialize()")
	fmt.Fprintln(os.Stderr, "  void shutdown()")
	fmt.Fprintln(os.Stderr, "  void setOption(string key, string value)")
	fmt.Fprintln(os.Stderr, "   getOptions()")
	fmt.Fprintln(os.Stderr)
	os.Exit(0)
}

func main() {
	flag.Usage = Usage
	var host string
	var port int
	var protocol string
	var urlString string
	var framed bool
	var useHttp bool
	var parsedUrl url.URL
	var trans thrift.TTransport
	_ = strconv.Atoi
	_ = math.Abs
	flag.Usage = Usage
	flag.StringVar(&host, "h", "localhost", "Specify host and port")
	flag.IntVar(&port, "p", 9090, "Specify port")
	flag.StringVar(&protocol, "P", "binary", "Specify the protocol (binary, compact, simplejson, json)")
	flag.StringVar(&urlString, "u", "", "Specify the url")
	flag.BoolVar(&framed, "framed", false, "Use framed transport")
	flag.BoolVar(&useHttp, "http", false, "Use http")
	flag.Parse()

	if len(urlString) > 0 {
		parsedUrl, err := url.Parse(urlString)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error parsing URL: ", err)
			flag.Usage()
		}
		host = parsedUrl.Host
		useHttp = len(parsedUrl.Scheme) <= 0 || parsedUrl.Scheme == "http"
	} else if useHttp {
		_, err := url.Parse(fmt.Sprint("http://", host, ":", port))
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error parsing URL: ", err)
			flag.Usage()
		}
	}

	cmd := flag.Arg(0)
	var err error
	if useHttp {
		trans, err = thrift.NewTHttpClient(parsedUrl.String())
	} else {
		portStr := fmt.Sprint(port)
		if strings.Contains(host, ":") {
			host, portStr, err = net.SplitHostPort(host)
			if err != nil {
				fmt.Fprintln(os.Stderr, "error with host:", err)
				os.Exit(1)
			}
		}
		trans, err = thrift.NewTSocket(net.JoinHostPort(host, portStr))
		if err != nil {
			fmt.Fprintln(os.Stderr, "error resolving address:", err)
			os.Exit(1)
		}
		if framed {
			trans = thrift.NewTFramedTransport(trans)
		}
	}
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error creating transport", err)
		os.Exit(1)
	}
	defer trans.Close()
	var protocolFactory thrift.TProtocolFactory
	switch protocol {
	case "compact":
		protocolFactory = thrift.NewTCompactProtocolFactory()
		break
	case "simplejson":
		protocolFactory = thrift.NewTSimpleJSONProtocolFactory()
		break
	case "json":
		protocolFactory = thrift.NewTJSONProtocolFactory()
		break
	case "binary", "":
		protocolFactory = thrift.NewTBinaryProtocolFactoryDefault()
		break
	default:
		fmt.Fprintln(os.Stderr, "Invalid protocol specified: ", protocol)
		Usage()
		os.Exit(1)
	}
	client := config.NewConfigServiceClientFactory(trans, protocolFactory)
	if err := trans.Open(); err != nil {
		fmt.Fprintln(os.Stderr, "Error opening socket to ", host, ":", port, " ", err)
		os.Exit(1)
	}

	switch cmd {
	case "registerService":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "RegisterService requires 1 args")
			flag.Usage()
		}
		arg11 := flag.Arg(1)
		mbTrans12 := thrift.NewTMemoryBufferLen(len(arg11))
		defer mbTrans12.Close()
		_, err13 := mbTrans12.WriteString(arg11)
		if err13 != nil {
			Usage()
			return
		}
		factory14 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt15 := factory14.GetProtocol(mbTrans12)
		argvalue0 := config.NewServiceConfigBean()
		err16 := argvalue0.Read(jsProt15)
		if err16 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.RegisterService(value0))
		fmt.Print("\n")
		break
	case "serviceHeart":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "ServiceHeart requires 1 args")
			flag.Usage()
		}
		argvalue0 := flag.Arg(1)
		value0 := argvalue0
		fmt.Print(client.ServiceHeart(value0))
		fmt.Print("\n")
		break
	case "loadServices":
		if flag.NArg()-1 != 0 {
			fmt.Fprintln(os.Stderr, "LoadServices requires 0 args")
			flag.Usage()
		}
		fmt.Print(client.LoadServices())
		fmt.Print("\n")
		break
	case "queryService":
		if flag.NArg()-1 != 2 {
			fmt.Fprintln(os.Stderr, "QueryService requires 2 args")
			flag.Usage()
		}
		argvalue0 := flag.Arg(1)
		value0 := argvalue0
		argvalue1 := flag.Arg(2)
		value1 := argvalue1
		fmt.Print(client.QueryService(value0, value1))
		fmt.Print("\n")
		break
	case "getName":
		if flag.NArg()-1 != 0 {
			fmt.Fprintln(os.Stderr, "GetName requires 0 args")
			flag.Usage()
		}
		fmt.Print(client.GetName())
		fmt.Print("\n")
		break
	case "getVersion":
		if flag.NArg()-1 != 0 {
			fmt.Fprintln(os.Stderr, "GetVersion requires 0 args")
			flag.Usage()
		}
		fmt.Print(client.GetVersion())
		fmt.Print("\n")
		break
	case "getServiceBizMethods":
		if flag.NArg()-1 != 0 {
			fmt.Fprintln(os.Stderr, "GetServiceBizMethods requires 0 args")
			flag.Usage()
		}
		fmt.Print(client.GetServiceBizMethods())
		fmt.Print("\n")
		break
	case "getBizMethodsInvokeInfo":
		if flag.NArg()-1 != 0 {
			fmt.Fprintln(os.Stderr, "GetBizMethodsInvokeInfo requires 0 args")
			flag.Usage()
		}
		fmt.Print(client.GetBizMethodsInvokeInfo())
		fmt.Print("\n")
		break
	case "getBizMethodInvokeInfo":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "GetBizMethodInvokeInfo requires 1 args")
			flag.Usage()
		}
		argvalue0 := flag.Arg(1)
		value0 := argvalue0
		fmt.Print(client.GetBizMethodInvokeInfo(value0))
		fmt.Print("\n")
		break
	case "getServiceCount":
		if flag.NArg()-1 != 0 {
			fmt.Fprintln(os.Stderr, "GetServiceCount requires 0 args")
			flag.Usage()
		}
		fmt.Print(client.GetServiceCount())
		fmt.Print("\n")
		break
	case "aliveSince":
		if flag.NArg()-1 != 0 {
			fmt.Fprintln(os.Stderr, "AliveSince requires 0 args")
			flag.Usage()
		}
		fmt.Print(client.AliveSince())
		fmt.Print("\n")
		break
	case "reinitialize":
		if flag.NArg()-1 != 0 {
			fmt.Fprintln(os.Stderr, "Reinitialize requires 0 args")
			flag.Usage()
		}
		fmt.Print(client.Reinitialize())
		fmt.Print("\n")
		break
	case "shutdown":
		if flag.NArg()-1 != 0 {
			fmt.Fprintln(os.Stderr, "Shutdown requires 0 args")
			flag.Usage()
		}
		fmt.Print(client.Shutdown())
		fmt.Print("\n")
		break
	case "setOption":
		if flag.NArg()-1 != 2 {
			fmt.Fprintln(os.Stderr, "SetOption requires 2 args")
			flag.Usage()
		}
		argvalue0 := flag.Arg(1)
		value0 := argvalue0
		argvalue1 := flag.Arg(2)
		value1 := argvalue1
		fmt.Print(client.SetOption(value0, value1))
		fmt.Print("\n")
		break
	case "getOptions":
		if flag.NArg()-1 != 0 {
			fmt.Fprintln(os.Stderr, "GetOptions requires 0 args")
			flag.Usage()
		}
		fmt.Print(client.GetOptions())
		fmt.Print("\n")
		break
	case "":
		Usage()
		break
	default:
		fmt.Fprintln(os.Stderr, "Invalid function ", cmd)
	}
}