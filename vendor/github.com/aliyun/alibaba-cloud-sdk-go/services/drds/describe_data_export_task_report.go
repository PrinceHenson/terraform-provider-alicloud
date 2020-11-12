package drds

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

// DescribeDataExportTaskReport invokes the drds.DescribeDataExportTaskReport API synchronously
func (client *Client) DescribeDataExportTaskReport(request *DescribeDataExportTaskReportRequest) (response *DescribeDataExportTaskReportResponse, err error) {
	response = CreateDescribeDataExportTaskReportResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDataExportTaskReportWithChan invokes the drds.DescribeDataExportTaskReport API asynchronously
func (client *Client) DescribeDataExportTaskReportWithChan(request *DescribeDataExportTaskReportRequest) (<-chan *DescribeDataExportTaskReportResponse, <-chan error) {
	responseChan := make(chan *DescribeDataExportTaskReportResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDataExportTaskReport(request)
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

// DescribeDataExportTaskReportWithCallback invokes the drds.DescribeDataExportTaskReport API asynchronously
func (client *Client) DescribeDataExportTaskReportWithCallback(request *DescribeDataExportTaskReportRequest, callback func(response *DescribeDataExportTaskReportResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDataExportTaskReportResponse
		var err error
		defer close(result)
		response, err = client.DescribeDataExportTaskReport(request)
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

// DescribeDataExportTaskReportRequest is the request struct for api DescribeDataExportTaskReport
type DescribeDataExportTaskReportRequest struct {
	*requests.RpcRequest
	TaskId requests.Integer `position:"Query" name:"TaskId"`
}

// DescribeDataExportTaskReportResponse is the response struct for api DescribeDataExportTaskReport
type DescribeDataExportTaskReportResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Data      string `json:"Data" xml:"Data"`
}

// CreateDescribeDataExportTaskReportRequest creates a request to invoke DescribeDataExportTaskReport API
func CreateDescribeDataExportTaskReportRequest() (request *DescribeDataExportTaskReportRequest) {
	request = &DescribeDataExportTaskReportRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Drds", "2019-01-23", "DescribeDataExportTaskReport", "Drds", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeDataExportTaskReportResponse creates a response to parse from DescribeDataExportTaskReport response
func CreateDescribeDataExportTaskReportResponse() (response *DescribeDataExportTaskReportResponse) {
	response = &DescribeDataExportTaskReportResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
