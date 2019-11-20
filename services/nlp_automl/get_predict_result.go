package nlp_automl

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

// GetPredictResult invokes the nlp_automl.GetPredictResult API synchronously
// api document: https://help.aliyun.com/api/nlp-automl/getpredictresult.html
func (client *Client) GetPredictResult(request *GetPredictResultRequest) (response *GetPredictResultResponse, err error) {
	response = CreateGetPredictResultResponse()
	err = client.DoAction(request, response)
	return
}

// GetPredictResultWithChan invokes the nlp_automl.GetPredictResult API asynchronously
// api document: https://help.aliyun.com/api/nlp-automl/getpredictresult.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetPredictResultWithChan(request *GetPredictResultRequest) (<-chan *GetPredictResultResponse, <-chan error) {
	responseChan := make(chan *GetPredictResultResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetPredictResult(request)
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

// GetPredictResultWithCallback invokes the nlp_automl.GetPredictResult API asynchronously
// api document: https://help.aliyun.com/api/nlp-automl/getpredictresult.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetPredictResultWithCallback(request *GetPredictResultRequest, callback func(response *GetPredictResultResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetPredictResultResponse
		var err error
		defer close(result)
		response, err = client.GetPredictResult(request)
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

// GetPredictResultRequest is the request struct for api GetPredictResult
type GetPredictResultRequest struct {
	*requests.RpcRequest
	TopK         requests.Integer `position:"Body" name:"TopK"`
	ModelId      requests.Integer `position:"Body" name:"ModelId"`
	DetailTag    string           `position:"Body" name:"DetailTag"`
	Content      string           `position:"Body" name:"Content"`
	ModelVersion string           `position:"Body" name:"ModelVersion"`
}

// GetPredictResultResponse is the response struct for api GetPredictResult
type GetPredictResultResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Content   string `json:"Content" xml:"Content"`
	Code      int    `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	Success   bool   `json:"Success" xml:"Success"`
}

// CreateGetPredictResultRequest creates a request to invoke GetPredictResult API
func CreateGetPredictResultRequest() (request *GetPredictResultRequest) {
	request = &GetPredictResultRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("nlp-automl", "2019-11-11", "GetPredictResult", "nlpautoml", "openAPI")
	return
}

// CreateGetPredictResultResponse creates a response to parse from GetPredictResult response
func CreateGetPredictResultResponse() (response *GetPredictResultResponse) {
	response = &GetPredictResultResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}