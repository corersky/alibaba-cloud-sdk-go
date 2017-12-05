package jaq

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

func (client *Client) MessagePrevention(request *MessagePreventionRequest) (response *MessagePreventionResponse, err error) {
	response = CreateMessagePreventionResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) MessagePreventionWithChan(request *MessagePreventionRequest) (<-chan *MessagePreventionResponse, <-chan error) {
	responseChan := make(chan *MessagePreventionResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.MessagePrevention(request)
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

func (client *Client) MessagePreventionWithCallback(request *MessagePreventionRequest, callback func(response *MessagePreventionResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *MessagePreventionResponse
		var err error
		defer close(result)
		response, err = client.MessagePrevention(request)
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

type MessagePreventionRequest struct {
	*requests.RpcRequest
	CurrentUrl      string `position:"Query" name:"CurrentUrl"`
	IdType          string `position:"Query" name:"IdType"`
	LoginDate       string `position:"Query" name:"LoginDate"`
	MacAddress      string `position:"Query" name:"MacAddress"`
	UserId          string `position:"Query" name:"UserId"`
	UserName        string `position:"Query" name:"UserName"`
	BankCardNumber  string `position:"Query" name:"BankCardNumber"`
	CallerName      string `position:"Query" name:"CallerName"`
	SDKToken        string `position:"Query" name:"SDKToken"`
	PhoneNumber     string `position:"Query" name:"PhoneNumber"`
	RegisterIp      string `position:"Query" name:"RegisterIp"`
	IDNumber        string `position:"Query" name:"IDNumber"`
	PasswordHash    string `position:"Query" name:"PasswordHash"`
	SessionId       string `position:"Query" name:"SessionId"`
	Agent           string `position:"Query" name:"Agent"`
	CompanyName     string `position:"Query" name:"CompanyName"`
	Ip              string `position:"Query" name:"Ip"`
	Cookie          string `position:"Query" name:"Cookie"`
	JsToken         string `position:"Query" name:"JsToken"`
	Source          string `position:"Query" name:"Source"`
	RegisterDate    string `position:"Query" name:"RegisterDate"`
	Email           string `position:"Query" name:"Email"`
	Address         string `position:"Query" name:"Address"`
	Referer         string `position:"Query" name:"Referer"`
	ExtendData      string `position:"Query" name:"ExtendData"`
	ProtocolVersion string `position:"Query" name:"ProtocolVersion"`
	LoginIp         string `position:"Query" name:"LoginIp"`
}

type MessagePreventionResponse struct {
	*responses.BaseResponse
	ErrorCode int    `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMsg  string `json:"ErrorMsg" xml:"ErrorMsg"`
	Data      struct {
		FnalDecision     int    `json:"FnalDecision" xml:"FnalDecision"`
		EventId          string `json:"EventId" xml:"EventId"`
		UserId           string `json:"UserId" xml:"UserId"`
		FinalScore       int    `json:"FinalScore" xml:"FinalScore"`
		FinalDesc        string `json:"FinalDesc" xml:"FinalDesc"`
		Detail           string `json:"Detail" xml:"Detail"`
		CaptchaCheckData string `json:"CaptchaCheckData" xml:"CaptchaCheckData"`
	} `json:"Data" xml:"Data"`
}

func CreateMessagePreventionRequest() (request *MessagePreventionRequest) {
	request = &MessagePreventionRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("jaq", "2017-08-25", "MessagePrevention", "", "")
	return
}

func CreateMessagePreventionResponse() (response *MessagePreventionResponse) {
	response = &MessagePreventionResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}