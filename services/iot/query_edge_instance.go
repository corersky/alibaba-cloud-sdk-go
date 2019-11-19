package iot

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

// QueryEdgeInstance invokes the iot.QueryEdgeInstance API synchronously
// api document: https://help.aliyun.com/api/iot/queryedgeinstance.html
func (client *Client) QueryEdgeInstance(request *QueryEdgeInstanceRequest) (response *QueryEdgeInstanceResponse, err error) {
	response = CreateQueryEdgeInstanceResponse()
	err = client.DoAction(request, response)
	return
}

// QueryEdgeInstanceWithChan invokes the iot.QueryEdgeInstance API asynchronously
// api document: https://help.aliyun.com/api/iot/queryedgeinstance.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryEdgeInstanceWithChan(request *QueryEdgeInstanceRequest) (<-chan *QueryEdgeInstanceResponse, <-chan error) {
	responseChan := make(chan *QueryEdgeInstanceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryEdgeInstance(request)
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

// QueryEdgeInstanceWithCallback invokes the iot.QueryEdgeInstance API asynchronously
// api document: https://help.aliyun.com/api/iot/queryedgeinstance.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryEdgeInstanceWithCallback(request *QueryEdgeInstanceRequest, callback func(response *QueryEdgeInstanceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryEdgeInstanceResponse
		var err error
		defer close(result)
		response, err = client.QueryEdgeInstance(request)
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

// QueryEdgeInstanceRequest is the request struct for api QueryEdgeInstance
type QueryEdgeInstanceRequest struct {
	*requests.RpcRequest
	CurrentPage   requests.Integer `position:"Query" name:"CurrentPage"`
	IotInstanceId string           `position:"Query" name:"IotInstanceId"`
	PageSize      requests.Integer `position:"Query" name:"PageSize"`
	Name          string           `position:"Query" name:"Name"`
}

// QueryEdgeInstanceResponse is the response struct for api QueryEdgeInstance
type QueryEdgeInstanceResponse struct {
	*responses.BaseResponse
	RequestId    string                  `json:"RequestId" xml:"RequestId"`
	Success      bool                    `json:"Success" xml:"Success"`
	Code         string                  `json:"Code" xml:"Code"`
	ErrorMessage string                  `json:"ErrorMessage" xml:"ErrorMessage"`
	Data         DataInQueryEdgeInstance `json:"Data" xml:"Data"`
}

// CreateQueryEdgeInstanceRequest creates a request to invoke QueryEdgeInstance API
func CreateQueryEdgeInstanceRequest() (request *QueryEdgeInstanceRequest) {
	request = &QueryEdgeInstanceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "QueryEdgeInstance", "iot", "openAPI")
	return
}

// CreateQueryEdgeInstanceResponse creates a response to parse from QueryEdgeInstance response
func CreateQueryEdgeInstanceResponse() (response *QueryEdgeInstanceResponse) {
	response = &QueryEdgeInstanceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
