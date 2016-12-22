package handler

import "vrv/im/service/vrv"

type VRVHandler struct {
}

// return service name
func (v *VRVHandler) GetName() (r string, err error) {
	return "", nil
}

// return service version
func (v *VRVHandler) GetVersion() (r string, err error) {
	return "", nil
}

// Get all the business  method of the service
func (v *VRVHandler) GetServiceBizMethods() (r []*vrv.BizMethodInfo, err error) {
	return nil, nil
}

// Get all the business  method invocation of services provide information
func (v *VRVHandler) GetBizMethodsInvokeInfo() (r map[string]*vrv.BizMethodInvokeInfo, err error) {
	return nil, nil
}

// Get the business  method invocation of services provide information
//
// Parameters:
//  - MethodName
func (v *VRVHandler) GetBizMethodInvokeInfo(methodName string) (r *vrv.BizMethodInvokeInfo, err error) {
	return nil, nil
}

// Access to the service call total number
func (v *VRVHandler) GetServiceCount() (r int64, err error) {
	return 0, nil
}

// Access to services run duration unit of seconds
func (v *VRVHandler) AliveSince() (r int64, err error) {
	return 0, nil
}

// Tell the server to reload its configuration, reopen log files, etc
func (v *VRVHandler) Reinitialize() (err error) {
	return nil
}

// Suggest a shutdown to the server
func (v *VRVHandler) Shutdown() (err error) {
	return nil
}

// Sets an option
//
// Parameters:
//  - Key
//  - Value
func (v *VRVHandler) SetOption(key string, value string) (err error) {
	return nil
}

// Gets all options
func (v *VRVHandler) GetOptions() (r map[string]string, err error) {
	return nil, nil
}
