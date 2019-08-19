package cdn

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

// DescribeUserUsageDataExportTask invokes the cdn.DescribeUserUsageDataExportTask API synchronously
// api document: https://help.aliyun.com/api/cdn/describeuserusagedataexporttask.html
func (client *Client) DescribeUserUsageDataExportTask(request *DescribeUserUsageDataExportTaskRequest) (response *DescribeUserUsageDataExportTaskResponse, err error) {
	response = CreateDescribeUserUsageDataExportTaskResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeUserUsageDataExportTaskWithChan invokes the cdn.DescribeUserUsageDataExportTask API asynchronously
// api document: https://help.aliyun.com/api/cdn/describeuserusagedataexporttask.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeUserUsageDataExportTaskWithChan(request *DescribeUserUsageDataExportTaskRequest) (<-chan *DescribeUserUsageDataExportTaskResponse, <-chan error) {
	responseChan := make(chan *DescribeUserUsageDataExportTaskResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeUserUsageDataExportTask(request)
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

// DescribeUserUsageDataExportTaskWithCallback invokes the cdn.DescribeUserUsageDataExportTask API asynchronously
// api document: https://help.aliyun.com/api/cdn/describeuserusagedataexporttask.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeUserUsageDataExportTaskWithCallback(request *DescribeUserUsageDataExportTaskRequest, callback func(response *DescribeUserUsageDataExportTaskResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeUserUsageDataExportTaskResponse
		var err error
		defer close(result)
		response, err = client.DescribeUserUsageDataExportTask(request)
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

// DescribeUserUsageDataExportTaskRequest is the request struct for api DescribeUserUsageDataExportTask
type DescribeUserUsageDataExportTaskRequest struct {
	*requests.RpcRequest
	PageNumber string           `position:"Query" name:"PageNumber"`
	PageSize   string           `position:"Query" name:"PageSize"`
	OwnerId    requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribeUserUsageDataExportTaskResponse is the response struct for api DescribeUserUsageDataExportTask
type DescribeUserUsageDataExportTaskResponse struct {
	*responses.BaseResponse
	RequestId        string           `json:"RequestId" xml:"RequestId"`
	UsageDataPerPage UsageDataPerPage `json:"UsageDataPerPage" xml:"UsageDataPerPage"`
}

// CreateDescribeUserUsageDataExportTaskRequest creates a request to invoke DescribeUserUsageDataExportTask API
func CreateDescribeUserUsageDataExportTaskRequest() (request *DescribeUserUsageDataExportTaskRequest) {
	request = &DescribeUserUsageDataExportTaskRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cdn", "2018-05-10", "DescribeUserUsageDataExportTask", "cdn", "openAPI")
	return
}

// CreateDescribeUserUsageDataExportTaskResponse creates a response to parse from DescribeUserUsageDataExportTask response
func CreateDescribeUserUsageDataExportTaskResponse() (response *DescribeUserUsageDataExportTaskResponse) {
	response = &DescribeUserUsageDataExportTaskResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
