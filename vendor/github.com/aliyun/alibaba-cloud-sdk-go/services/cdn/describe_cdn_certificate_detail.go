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

// DescribeCdnCertificateDetail invokes the cdn.DescribeCdnCertificateDetail API synchronously
// api document: https://help.aliyun.com/api/cdn/describecdncertificatedetail.html
func (client *Client) DescribeCdnCertificateDetail(request *DescribeCdnCertificateDetailRequest) (response *DescribeCdnCertificateDetailResponse, err error) {
	response = CreateDescribeCdnCertificateDetailResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeCdnCertificateDetailWithChan invokes the cdn.DescribeCdnCertificateDetail API asynchronously
// api document: https://help.aliyun.com/api/cdn/describecdncertificatedetail.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeCdnCertificateDetailWithChan(request *DescribeCdnCertificateDetailRequest) (<-chan *DescribeCdnCertificateDetailResponse, <-chan error) {
	responseChan := make(chan *DescribeCdnCertificateDetailResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeCdnCertificateDetail(request)
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

// DescribeCdnCertificateDetailWithCallback invokes the cdn.DescribeCdnCertificateDetail API asynchronously
// api document: https://help.aliyun.com/api/cdn/describecdncertificatedetail.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeCdnCertificateDetailWithCallback(request *DescribeCdnCertificateDetailRequest, callback func(response *DescribeCdnCertificateDetailResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeCdnCertificateDetailResponse
		var err error
		defer close(result)
		response, err = client.DescribeCdnCertificateDetail(request)
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

// DescribeCdnCertificateDetailRequest is the request struct for api DescribeCdnCertificateDetail
type DescribeCdnCertificateDetailRequest struct {
	*requests.RpcRequest
	SecurityToken string           `position:"Query" name:"SecurityToken"`
	CertName      string           `position:"Query" name:"CertName"`
	OwnerId       requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribeCdnCertificateDetailResponse is the response struct for api DescribeCdnCertificateDetail
type DescribeCdnCertificateDetailResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Cert      string `json:"Cert" xml:"Cert"`
	Key       string `json:"Key" xml:"Key"`
	CertId    int64  `json:"CertId" xml:"CertId"`
	CertName  string `json:"CertName" xml:"CertName"`
}

// CreateDescribeCdnCertificateDetailRequest creates a request to invoke DescribeCdnCertificateDetail API
func CreateDescribeCdnCertificateDetailRequest() (request *DescribeCdnCertificateDetailRequest) {
	request = &DescribeCdnCertificateDetailRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cdn", "2018-05-10", "DescribeCdnCertificateDetail", "cdn", "openAPI")
	return
}

// CreateDescribeCdnCertificateDetailResponse creates a response to parse from DescribeCdnCertificateDetail response
func CreateDescribeCdnCertificateDetailResponse() (response *DescribeCdnCertificateDetailResponse) {
	response = &DescribeCdnCertificateDetailResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
