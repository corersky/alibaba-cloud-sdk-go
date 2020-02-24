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

// AssignJobs invokes the outboundbot.AssignJobs API synchronously
// api document: https://help.aliyun.com/api/outboundbot/assignjobs.html
func (client *Client) AssignJobs(request *AssignJobsRequest) (response *AssignJobsResponse, err error) {
	response = CreateAssignJobsResponse()
	err = client.DoAction(request, response)
	return
}

// AssignJobsWithChan invokes the outboundbot.AssignJobs API asynchronously
// api document: https://help.aliyun.com/api/outboundbot/assignjobs.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) AssignJobsWithChan(request *AssignJobsRequest) (<-chan *AssignJobsResponse, <-chan error) {
	responseChan := make(chan *AssignJobsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AssignJobs(request)
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

// AssignJobsWithCallback invokes the outboundbot.AssignJobs API asynchronously
// api document: https://help.aliyun.com/api/outboundbot/assignjobs.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) AssignJobsWithCallback(request *AssignJobsRequest, callback func(response *AssignJobsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AssignJobsResponse
		var err error
		defer close(result)
		response, err = client.AssignJobs(request)
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

// AssignJobsRequest is the request struct for api AssignJobs
type AssignJobsRequest struct {
	*requests.RpcRequest
	JobsJson      *[]string `position:"Query" name:"JobsJson"  type:"Repeated"`
	CallingNumber *[]string `position:"Query" name:"CallingNumber"  type:"Repeated"`
	InstanceId    string    `position:"Query" name:"InstanceId"`
	StrategyJson  string    `position:"Query" name:"StrategyJson"`
	JobGroupId    string    `position:"Query" name:"JobGroupId"`
}

// AssignJobsResponse is the response struct for api AssignJobs
type AssignJobsResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Success        bool   `json:"Success" xml:"Success"`
	Code           string `json:"Code" xml:"Code"`
	Message        string `json:"Message" xml:"Message"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	JobGroupId     string `json:"JobGroupId" xml:"JobGroupId"`
}

// CreateAssignJobsRequest creates a request to invoke AssignJobs API
func CreateAssignJobsRequest() (request *AssignJobsRequest) {
	request = &AssignJobsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("OutboundBot", "2019-12-26", "AssignJobs", "outboundbot", "openAPI")
	return
}

// CreateAssignJobsResponse creates a response to parse from AssignJobs response
func CreateAssignJobsResponse() (response *AssignJobsResponse) {
	response = &AssignJobsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
