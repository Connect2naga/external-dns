package pvtz

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// AddZone invokes the pvtz.AddZone API synchronously
// api document: https://help.aliyun.com/api/pvtz/addzone.html
func (client *Client) AddZone(request *AddZoneRequest) (response *AddZoneResponse, err error) {
	response = CreateAddZoneResponse()
	err = client.DoAction(request, response)
	return
}

// AddZoneWithChan invokes the pvtz.AddZone API asynchronously
// api document: https://help.aliyun.com/api/pvtz/addzone.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) AddZoneWithChan(request *AddZoneRequest) (<-chan *AddZoneResponse, <-chan error) {
	responseChan := make(chan *AddZoneResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AddZone(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// AddZoneWithCallback invokes the pvtz.AddZone API asynchronously
// api document: https://help.aliyun.com/api/pvtz/addzone.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) AddZoneWithCallback(request *AddZoneRequest, callback func(response *AddZoneResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AddZoneResponse
		var err error
		defer close(result)
		response, err = client.AddZone(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// AddZoneRequest is the request struct for api AddZone
type AddZoneRequest struct {
	*requests.RpcRequest
	ProxyPattern    string `position:"Query" name:"ProxyPattern"`
	ZoneName        string `position:"Query" name:"ZoneName"`
	ResourceGroupId string `position:"Query" name:"ResourceGroupId"`
	UserClientIp    string `position:"Query" name:"UserClientIp"`
	Lang            string `position:"Query" name:"Lang"`
}

// AddZoneResponse is the response struct for api AddZone
type AddZoneResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	ZoneId    string `json:"ZoneId" xml:"ZoneId"`
	ZoneName  string `json:"ZoneName" xml:"ZoneName"`
}

// CreateAddZoneRequest creates a request to invoke AddZone API
func CreateAddZoneRequest() (request *AddZoneRequest) {
	request = &AddZoneRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("pvtz", "2018-01-01", "AddZone", "pvtz", "openAPI")
	return
}

// CreateAddZoneResponse creates a response to parse from AddZone response
func CreateAddZoneResponse() (response *AddZoneResponse) {
	response = &AddZoneResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
