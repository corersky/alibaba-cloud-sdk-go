package cms

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

func (client *Client) ListMyGroups(request *ListMyGroupsRequest) (response *ListMyGroupsResponse, err error) {
	response = CreateListMyGroupsResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) ListMyGroupsWithChan(request *ListMyGroupsRequest) (<-chan *ListMyGroupsResponse, <-chan error) {
	responseChan := make(chan *ListMyGroupsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListMyGroups(request)
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

func (client *Client) ListMyGroupsWithCallback(request *ListMyGroupsRequest, callback func(response *ListMyGroupsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListMyGroupsResponse
		var err error
		defer close(result)
		response, err = client.ListMyGroups(request)
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

type ListMyGroupsRequest struct {
	*requests.RpcRequest
	PageNumber          requests.Integer `position:"Query" name:"PageNumber"`
	PageSize            requests.Integer `position:"Query" name:"PageSize"`
	Keyword             string           `position:"Query" name:"Keyword"`
	Type                string           `position:"Query" name:"Type"`
	SelectContactGroups requests.Boolean `position:"Query" name:"SelectContactGroups"`
	InstanceId          string           `position:"Query" name:"InstanceId"`
	BindUrls            string           `position:"Query" name:"BindUrls"`
	GroupName           string           `position:"Query" name:"GroupName"`
}

type ListMyGroupsResponse struct {
	*responses.BaseResponse
	RequestId    string                  `json:"RequestId" xml:"RequestId"`
	Success      bool                    `json:"Success" xml:"Success"`
	ErrorCode    int                     `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMessage string                  `json:"ErrorMessage" xml:"ErrorMessage"`
	PageNumber   int                     `json:"PageNumber" xml:"PageNumber"`
	PageSize     int                     `json:"PageSize" xml:"PageSize"`
	Total        int                     `json:"Total" xml:"Total"`
	Resources    ResourcesInListMyGroups `json:"Resources" xml:"Resources"`
}

func CreateListMyGroupsRequest() (request *ListMyGroupsRequest) {
	request = &ListMyGroupsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cms", "2017-03-01", "ListMyGroups", "", "")
	return
}

func CreateListMyGroupsResponse() (response *ListMyGroupsResponse) {
	response = &ListMyGroupsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}