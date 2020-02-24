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

// ModifyIntent invokes the outboundbot.ModifyIntent API synchronously
// api document: https://help.aliyun.com/api/outboundbot/modifyintent.html
func (client *Client) ModifyIntent(request *ModifyIntentRequest) (response *ModifyIntentResponse, err error) {
	response = CreateModifyIntentResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyIntentWithChan invokes the outboundbot.ModifyIntent API asynchronously
// api document: https://help.aliyun.com/api/outboundbot/modifyintent.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyIntentWithChan(request *ModifyIntentRequest) (<-chan *ModifyIntentResponse, <-chan error) {
	responseChan := make(chan *ModifyIntentResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyIntent(request)
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

// ModifyIntentWithCallback invokes the outboundbot.ModifyIntent API asynchronously
// api document: https://help.aliyun.com/api/outboundbot/modifyintent.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyIntentWithCallback(request *ModifyIntentRequest, callback func(response *ModifyIntentResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyIntentResponse
		var err error
		defer close(result)
		response, err = client.ModifyIntent(request)
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

// ModifyIntentRequest is the request struct for api ModifyIntent
type ModifyIntentRequest struct {
	*requests.RpcRequest
	Utterances        string `position:"Query" name:"Utterances"`
	Keywords          string `position:"Query" name:"Keywords"`
	IntentDescription string `position:"Query" name:"IntentDescription"`
	IntentId          string `position:"Query" name:"IntentId"`
	ScriptId          string `position:"Query" name:"ScriptId"`
	InstanceId        string `position:"Query" name:"InstanceId"`
	IntentName        string `position:"Query" name:"IntentName"`
}

// ModifyIntentResponse is the response struct for api ModifyIntent
type ModifyIntentResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Success        bool   `json:"Success" xml:"Success"`
	Code           string `json:"Code" xml:"Code"`
	Message        string `json:"Message" xml:"Message"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	IntentId       string `json:"IntentId" xml:"IntentId"`
}

// CreateModifyIntentRequest creates a request to invoke ModifyIntent API
func CreateModifyIntentRequest() (request *ModifyIntentRequest) {
	request = &ModifyIntentRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("OutboundBot", "2019-12-26", "ModifyIntent", "outboundbot", "openAPI")
	return
}

// CreateModifyIntentResponse creates a response to parse from ModifyIntent response
func CreateModifyIntentResponse() (response *ModifyIntentResponse) {
	response = &ModifyIntentResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
