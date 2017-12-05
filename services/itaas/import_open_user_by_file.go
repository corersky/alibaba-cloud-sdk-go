package itaas

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

func (client *Client) ImportOpenUserByFile(request *ImportOpenUserByFileRequest) (response *ImportOpenUserByFileResponse, err error) {
	response = CreateImportOpenUserByFileResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) ImportOpenUserByFileWithChan(request *ImportOpenUserByFileRequest) (<-chan *ImportOpenUserByFileResponse, <-chan error) {
	responseChan := make(chan *ImportOpenUserByFileResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ImportOpenUserByFile(request)
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

func (client *Client) ImportOpenUserByFileWithCallback(request *ImportOpenUserByFileRequest, callback func(response *ImportOpenUserByFileResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ImportOpenUserByFileResponse
		var err error
		defer close(result)
		response, err = client.ImportOpenUserByFile(request)
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

type ImportOpenUserByFileRequest struct {
	*requests.RpcRequest
	AuthType      string `position:"Query" name:"AuthType"`
	Sid           string `position:"Query" name:"Sid"`
	OsType        string `position:"Query" name:"OsType"`
	AppName       string `position:"Query" name:"AppName"`
	AppVersion    string `position:"Query" name:"AppVersion"`
	Language      string `position:"Query" name:"Language"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	FileContext   string `position:"Query" name:"FileContext"`
}

type ImportOpenUserByFileResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	Success      bool   `json:"Success" xml:"Success"`
	Code         string `json:"Code" xml:"Code"`
	BusinessCode string `json:"BusinessCode" xml:"BusinessCode"`
	Message      string `json:"Message" xml:"Message"`
	ResultData   struct {
	} `json:"ResultData" xml:"ResultData"`
}

func CreateImportOpenUserByFileRequest() (request *ImportOpenUserByFileRequest) {
	request = &ImportOpenUserByFileRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ITaaS", "2017-05-12", "ImportOpenUserByFile", "", "")
	return
}

func CreateImportOpenUserByFileResponse() (response *ImportOpenUserByFileResponse) {
	response = &ImportOpenUserByFileResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}