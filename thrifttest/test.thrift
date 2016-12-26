namespace go thrifttest.demo
struct Name{
    1:string name
}
service TestService{
    string HelloWorld()
    string HelloWorldForString(1:string name)
    string HelloWorldForMap(1:map<string,i32> name)
    string HelloWorldForStruct(1:Name name)
}