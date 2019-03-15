package rtc

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

// RemoveParticipants invokes the rtc.RemoveParticipants API synchronously
// api document: https://help.aliyun.com/api/rtc/removeparticipants.html
func (client *Client) RemoveParticipants(request *RemoveParticipantsRequest) (response *RemoveParticipantsResponse, err error) {
	response = CreateRemoveParticipantsResponse()
	err = client.DoAction(request, response)
	return
}

// RemoveParticipantsWithChan invokes the rtc.RemoveParticipants API asynchronously
// api document: https://help.aliyun.com/api/rtc/removeparticipants.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) RemoveParticipantsWithChan(request *RemoveParticipantsRequest) (<-chan *RemoveParticipantsResponse, <-chan error) {
	responseChan := make(chan *RemoveParticipantsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RemoveParticipants(request)
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

// RemoveParticipantsWithCallback invokes the rtc.RemoveParticipants API asynchronously
// api document: https://help.aliyun.com/api/rtc/removeparticipants.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) RemoveParticipantsWithCallback(request *RemoveParticipantsRequest, callback func(response *RemoveParticipantsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RemoveParticipantsResponse
		var err error
		defer close(result)
		response, err = client.RemoveParticipants(request)
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

// RemoveParticipantsRequest is the request struct for api RemoveParticipants
type RemoveParticipantsRequest struct {
	*requests.RpcRequest
	ParticipantIds *[]string        `position:"Query" name:"ParticipantIds"  type:"Repeated"`
	OwnerId        requests.Integer `position:"Query" name:"OwnerId"`
	ConferenceId   string           `position:"Query" name:"ConferenceId"`
	AppId          string           `position:"Query" name:"AppId"`
}

// RemoveParticipantsResponse is the response struct for api RemoveParticipants
type RemoveParticipantsResponse struct {
	*responses.BaseResponse
	RequestId    string                           `json:"RequestId" xml:"RequestId"`
	ConferenceId string                           `json:"ConferenceId" xml:"ConferenceId"`
	Participants ParticipantsInRemoveParticipants `json:"Participants" xml:"Participants"`
}

// CreateRemoveParticipantsRequest creates a request to invoke RemoveParticipants API
func CreateRemoveParticipantsRequest() (request *RemoveParticipantsRequest) {
	request = &RemoveParticipantsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("rtc", "2018-01-11", "RemoveParticipants", "rtc", "openAPI")
	return
}

// CreateRemoveParticipantsResponse creates a response to parse from RemoveParticipants response
func CreateRemoveParticipantsResponse() (response *RemoveParticipantsResponse) {
	response = &RemoveParticipantsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
