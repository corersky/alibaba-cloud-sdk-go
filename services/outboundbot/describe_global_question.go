package outboundbot

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

// DescribeGlobalQuestion invokes the outboundbot.DescribeGlobalQuestion API synchronously
// api document: https://help.aliyun.com/api/outboundbot/describeglobalquestion.html
func (client *Client) DescribeGlobalQuestion(request *DescribeGlobalQuestionRequest) (response *DescribeGlobalQuestionResponse, err error) {
	response = CreateDescribeGlobalQuestionResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeGlobalQuestionWithChan invokes the outboundbot.DescribeGlobalQuestion API asynchronously
// api document: https://help.aliyun.com/api/outboundbot/describeglobalquestion.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeGlobalQuestionWithChan(request *DescribeGlobalQuestionRequest) (<-chan *DescribeGlobalQuestionResponse, <-chan error) {
	responseChan := make(chan *DescribeGlobalQuestionResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeGlobalQuestion(request)
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

// DescribeGlobalQuestionWithCallback invokes the outboundbot.DescribeGlobalQuestion API asynchronously
// api document: https://help.aliyun.com/api/outboundbot/describeglobalquestion.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeGlobalQuestionWithCallback(request *DescribeGlobalQuestionRequest, callback func(response *DescribeGlobalQuestionResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeGlobalQuestionResponse
		var err error
		defer close(result)
		response, err = client.DescribeGlobalQuestion(request)
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

// DescribeGlobalQuestionRequest is the request struct for api DescribeGlobalQuestion
type DescribeGlobalQuestionRequest struct {
	*requests.RpcRequest
	GlobalQuestionId string `position:"Query" name:"GlobalQuestionId"`
	ScriptId         string `position:"Query" name:"ScriptId"`
	InstanceId       string `position:"Query" name:"InstanceId"`
}

// DescribeGlobalQuestionResponse is the response struct for api DescribeGlobalQuestion
type DescribeGlobalQuestionResponse struct {
	*responses.BaseResponse
	RequestId      string         `json:"RequestId" xml:"RequestId"`
	Success        bool           `json:"Success" xml:"Success"`
	Code           string         `json:"Code" xml:"Code"`
	Message        string         `json:"Message" xml:"Message"`
	HttpStatusCode int            `json:"HttpStatusCode" xml:"HttpStatusCode"`
	GlobalQuestion GlobalQuestion `json:"GlobalQuestion" xml:"GlobalQuestion"`
}

// CreateDescribeGlobalQuestionRequest creates a request to invoke DescribeGlobalQuestion API
func CreateDescribeGlobalQuestionRequest() (request *DescribeGlobalQuestionRequest) {
	request = &DescribeGlobalQuestionRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("OutboundBot", "2019-12-26", "DescribeGlobalQuestion", "outboundbot", "openAPI")
	return
}

// CreateDescribeGlobalQuestionResponse creates a response to parse from DescribeGlobalQuestion response
func CreateDescribeGlobalQuestionResponse() (response *DescribeGlobalQuestionResponse) {
	response = &DescribeGlobalQuestionResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
