package oos

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

// ListSecretParameterVersions invokes the oos.ListSecretParameterVersions API synchronously
func (client *Client) ListSecretParameterVersions(request *ListSecretParameterVersionsRequest) (response *ListSecretParameterVersionsResponse, err error) {
	response = CreateListSecretParameterVersionsResponse()
	err = client.DoAction(request, response)
	return
}

// ListSecretParameterVersionsWithChan invokes the oos.ListSecretParameterVersions API asynchronously
func (client *Client) ListSecretParameterVersionsWithChan(request *ListSecretParameterVersionsRequest) (<-chan *ListSecretParameterVersionsResponse, <-chan error) {
	responseChan := make(chan *ListSecretParameterVersionsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListSecretParameterVersions(request)
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

// ListSecretParameterVersionsWithCallback invokes the oos.ListSecretParameterVersions API asynchronously
func (client *Client) ListSecretParameterVersionsWithCallback(request *ListSecretParameterVersionsRequest, callback func(response *ListSecretParameterVersionsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListSecretParameterVersionsResponse
		var err error
		defer close(result)
		response, err = client.ListSecretParameterVersions(request)
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

// ListSecretParameterVersionsRequest is the request struct for api ListSecretParameterVersions
type ListSecretParameterVersionsRequest struct {
	*requests.RpcRequest
	WithDecryption requests.Boolean `position:"Query" name:"WithDecryption"`
	NextToken      string           `position:"Query" name:"NextToken"`
	Name           string           `position:"Query" name:"Name"`
	MaxResults     requests.Integer `position:"Query" name:"MaxResults"`
	ShareType      string           `position:"Query" name:"ShareType"`
}

// ListSecretParameterVersionsResponse is the response struct for api ListSecretParameterVersions
type ListSecretParameterVersionsResponse struct {
	*responses.BaseResponse
	RequestId         string             `json:"RequestId" xml:"RequestId"`
	MaxResults        int                `json:"MaxResults" xml:"MaxResults"`
	NextToken         string             `json:"NextToken" xml:"NextToken"`
	TotalCount        int                `json:"TotalCount" xml:"TotalCount"`
	Name              string             `json:"Name" xml:"Name"`
	Id                string             `json:"Id" xml:"Id"`
	Type              string             `json:"Type" xml:"Type"`
	Description       string             `json:"Description" xml:"Description"`
	CreatedDate       string             `json:"CreatedDate" xml:"CreatedDate"`
	CreatedBy         string             `json:"CreatedBy" xml:"CreatedBy"`
	ParameterVersions []ParameterVersion `json:"ParameterVersions" xml:"ParameterVersions"`
}

// CreateListSecretParameterVersionsRequest creates a request to invoke ListSecretParameterVersions API
func CreateListSecretParameterVersionsRequest() (request *ListSecretParameterVersionsRequest) {
	request = &ListSecretParameterVersionsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("oos", "2019-06-01", "ListSecretParameterVersions", "oos", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListSecretParameterVersionsResponse creates a response to parse from ListSecretParameterVersions response
func CreateListSecretParameterVersionsResponse() (response *ListSecretParameterVersionsResponse) {
	response = &ListSecretParameterVersionsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
