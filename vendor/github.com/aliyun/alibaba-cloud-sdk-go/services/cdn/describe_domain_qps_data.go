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

// DescribeDomainQpsData invokes the cdn.DescribeDomainQpsData API synchronously
// api document: https://help.aliyun.com/api/cdn/describedomainqpsdata.html
func (client *Client) DescribeDomainQpsData(request *DescribeDomainQpsDataRequest) (response *DescribeDomainQpsDataResponse, err error) {
	response = CreateDescribeDomainQpsDataResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDomainQpsDataWithChan invokes the cdn.DescribeDomainQpsData API asynchronously
// api document: https://help.aliyun.com/api/cdn/describedomainqpsdata.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeDomainQpsDataWithChan(request *DescribeDomainQpsDataRequest) (<-chan *DescribeDomainQpsDataResponse, <-chan error) {
	responseChan := make(chan *DescribeDomainQpsDataResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDomainQpsData(request)
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

// DescribeDomainQpsDataWithCallback invokes the cdn.DescribeDomainQpsData API asynchronously
// api document: https://help.aliyun.com/api/cdn/describedomainqpsdata.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeDomainQpsDataWithCallback(request *DescribeDomainQpsDataRequest, callback func(response *DescribeDomainQpsDataResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDomainQpsDataResponse
		var err error
		defer close(result)
		response, err = client.DescribeDomainQpsData(request)
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

// DescribeDomainQpsDataRequest is the request struct for api DescribeDomainQpsData
type DescribeDomainQpsDataRequest struct {
	*requests.RpcRequest
	LocationNameEn string           `position:"Query" name:"LocationNameEn"`
	StartTime      string           `position:"Query" name:"StartTime"`
	IspNameEn      string           `position:"Query" name:"IspNameEn"`
	DomainName     string           `position:"Query" name:"DomainName"`
	EndTime        string           `position:"Query" name:"EndTime"`
	OwnerId        requests.Integer `position:"Query" name:"OwnerId"`
	Interval       string           `position:"Query" name:"Interval"`
}

// DescribeDomainQpsDataResponse is the response struct for api DescribeDomainQpsData
type DescribeDomainQpsDataResponse struct {
	*responses.BaseResponse
	RequestId       string                                 `json:"RequestId" xml:"RequestId"`
	DomainName      string                                 `json:"DomainName" xml:"DomainName"`
	StartTime       string                                 `json:"StartTime" xml:"StartTime"`
	EndTime         string                                 `json:"EndTime" xml:"EndTime"`
	DataInterval    string                                 `json:"DataInterval" xml:"DataInterval"`
	QpsDataInterval QpsDataIntervalInDescribeDomainQpsData `json:"QpsDataInterval" xml:"QpsDataInterval"`
}

// CreateDescribeDomainQpsDataRequest creates a request to invoke DescribeDomainQpsData API
func CreateDescribeDomainQpsDataRequest() (request *DescribeDomainQpsDataRequest) {
	request = &DescribeDomainQpsDataRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cdn", "2018-05-10", "DescribeDomainQpsData", "cdn", "openAPI")
	return
}

// CreateDescribeDomainQpsDataResponse creates a response to parse from DescribeDomainQpsData response
func CreateDescribeDomainQpsDataResponse() (response *DescribeDomainQpsDataResponse) {
	response = &DescribeDomainQpsDataResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
