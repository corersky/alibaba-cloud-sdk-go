package emr

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

func (client *Client) ListClusterServiceConfigHistory(request *ListClusterServiceConfigHistoryRequest) (response *ListClusterServiceConfigHistoryResponse, err error) {
	response = CreateListClusterServiceConfigHistoryResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) ListClusterServiceConfigHistoryWithChan(request *ListClusterServiceConfigHistoryRequest) (<-chan *ListClusterServiceConfigHistoryResponse, <-chan error) {
	responseChan := make(chan *ListClusterServiceConfigHistoryResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListClusterServiceConfigHistory(request)
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

func (client *Client) ListClusterServiceConfigHistoryWithCallback(request *ListClusterServiceConfigHistoryRequest, callback func(response *ListClusterServiceConfigHistoryResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListClusterServiceConfigHistoryResponse
		var err error
		defer close(result)
		response, err = client.ListClusterServiceConfigHistory(request)
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

type ListClusterServiceConfigHistoryRequest struct {
	*requests.RpcRequest
	PageSize        string `position:"Query" name:"PageSize"`
	PageNumber      string `position:"Query" name:"PageNumber"`
	ConfigVersion   string `position:"Query" name:"ConfigVersion"`
	ClusterId       string `position:"Query" name:"ClusterId"`
	ServiceName     string `position:"Query" name:"ServiceName"`
	ResourceOwnerId string `position:"Query" name:"ResourceOwnerId"`
}

type ListClusterServiceConfigHistoryResponse struct {
	*responses.BaseResponse
	RequestId         string `json:"RequestId" xml:"RequestId"`
	TotalCount        int    `json:"TotalCount" xml:"TotalCount"`
	PageNumber        int    `json:"PageNumber" xml:"PageNumber"`
	PageSize          int    `json:"PageSize" xml:"PageSize"`
	ConfigHistoryList []struct {
		ServiceName    string `json:"ServiceName" xml:"ServiceName"`
		ConfigVersion  string `json:"ConfigVersion" xml:"ConfigVersion"`
		ConfigFileName string `json:"ConfigFileName" xml:"ConfigFileName"`
		ConfigItemName string `json:"ConfigItemName" xml:"ConfigItemName"`
		NewValue       string `json:"NewValue" xml:"NewValue"`
		OldValue       string `json:"OldValue" xml:"OldValue"`
		Applied        bool   `json:"Applied" xml:"Applied"`
		CreateTime     int64  `json:"CreateTime" xml:"CreateTime"`
		Author         string `json:"Author" xml:"Author"`
		Comment        string `json:"Comment" xml:"Comment"`
	} `json:"ConfigHistoryList" xml:"ConfigHistoryList"`
}

func CreateListClusterServiceConfigHistoryRequest() (request *ListClusterServiceConfigHistoryRequest) {
	request = &ListClusterServiceConfigHistoryRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Emr", "2016-04-08", "ListClusterServiceConfigHistory", "", "")
	return
}

func CreateListClusterServiceConfigHistoryResponse() (response *ListClusterServiceConfigHistoryResponse) {
	response = &ListClusterServiceConfigHistoryResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}