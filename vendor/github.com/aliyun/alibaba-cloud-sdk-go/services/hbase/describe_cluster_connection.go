package hbase

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

// DescribeClusterConnection invokes the hbase.DescribeClusterConnection API synchronously
func (client *Client) DescribeClusterConnection(request *DescribeClusterConnectionRequest) (response *DescribeClusterConnectionResponse, err error) {
	response = CreateDescribeClusterConnectionResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeClusterConnectionWithChan invokes the hbase.DescribeClusterConnection API asynchronously
func (client *Client) DescribeClusterConnectionWithChan(request *DescribeClusterConnectionRequest) (<-chan *DescribeClusterConnectionResponse, <-chan error) {
	responseChan := make(chan *DescribeClusterConnectionResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeClusterConnection(request)
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

// DescribeClusterConnectionWithCallback invokes the hbase.DescribeClusterConnection API asynchronously
func (client *Client) DescribeClusterConnectionWithCallback(request *DescribeClusterConnectionRequest, callback func(response *DescribeClusterConnectionResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeClusterConnectionResponse
		var err error
		defer close(result)
		response, err = client.DescribeClusterConnection(request)
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

// DescribeClusterConnectionRequest is the request struct for api DescribeClusterConnection
type DescribeClusterConnectionRequest struct {
	*requests.RpcRequest
	ClusterId string `position:"Query" name:"ClusterId"`
}

// DescribeClusterConnectionResponse is the response struct for api DescribeClusterConnection
type DescribeClusterConnectionResponse struct {
	*responses.BaseResponse
	RequestId           string              `json:"RequestId" xml:"RequestId"`
	NetType             string              `json:"NetType" xml:"NetType"`
	VpcId               string              `json:"VpcId" xml:"VpcId"`
	VSwitchId           string              `json:"VSwitchId" xml:"VSwitchId"`
	DbType              string              `json:"DbType" xml:"DbType"`
	IsMultimod          string              `json:"IsMultimod" xml:"IsMultimod"`
	UiProxyConnAddrInfo UiProxyConnAddrInfo `json:"UiProxyConnAddrInfo" xml:"UiProxyConnAddrInfo"`
	ThriftConn          ThriftConn          `json:"ThriftConn" xml:"ThriftConn"`
	ZkConnAddrs         []ZkConnAddr        `json:"ZkConnAddrs" xml:"ZkConnAddrs"`
	SlbConnAddrs        []SlbConnAddr       `json:"SlbConnAddrs" xml:"SlbConnAddrs"`
	ServiceConnAddrs    []ServiceConnAddr   `json:"ServiceConnAddrs" xml:"ServiceConnAddrs"`
}

// CreateDescribeClusterConnectionRequest creates a request to invoke DescribeClusterConnection API
func CreateDescribeClusterConnectionRequest() (request *DescribeClusterConnectionRequest) {
	request = &DescribeClusterConnectionRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("HBase", "2019-01-01", "DescribeClusterConnection", "hbase", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeClusterConnectionResponse creates a response to parse from DescribeClusterConnection response
func CreateDescribeClusterConnectionResponse() (response *DescribeClusterConnectionResponse) {
	response = &DescribeClusterConnectionResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
