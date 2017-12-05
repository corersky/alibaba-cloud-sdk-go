package ons

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

func (client *Client) OnsTopicDelete(request *OnsTopicDeleteRequest) (response *OnsTopicDeleteResponse, err error) {
	response = CreateOnsTopicDeleteResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) OnsTopicDeleteWithChan(request *OnsTopicDeleteRequest) (<-chan *OnsTopicDeleteResponse, <-chan error) {
	responseChan := make(chan *OnsTopicDeleteResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.OnsTopicDelete(request)
		responseChan <- response
		errChan <- err
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

func (client *Client) OnsTopicDeleteWithCallback(request *OnsTopicDeleteRequest, callback func(response *OnsTopicDeleteResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *OnsTopicDeleteResponse
		var err error
		defer close(result)
		response, err = client.OnsTopicDelete(request)
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

type OnsTopicDeleteRequest struct {
	*requests.RpcRequest
	Topic        string `position:"Query" name:"Topic"`
	OnsRegionId  string `position:"Query" name:"OnsRegionId"`
	PreventCache string `position:"Query" name:"PreventCache"`
	Cluster      string `position:"Query" name:"Cluster"`
	OnsPlatform  string `position:"Query" name:"OnsPlatform"`
}

type OnsTopicDeleteResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	HelpUrl   string `json:"HelpUrl" xml:"HelpUrl"`
}

func CreateOnsTopicDeleteRequest() (request *OnsTopicDeleteRequest) {
	request = &OnsTopicDeleteRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ons", "2017-09-18", "OnsTopicDelete", "", "")
	return
}

func CreateOnsTopicDeleteResponse() (response *OnsTopicDeleteResponse) {
	response = &OnsTopicDeleteResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}