package instance

import (
	"context"
	"fmt"

	"github.com/IBM-Cloud/power-go-client/errors"
	"github.com/IBM-Cloud/power-go-client/helpers"
	"github.com/go-openapi/runtime"

	"github.com/IBM-Cloud/power-go-client/ibmpisession"
	"github.com/IBM-Cloud/power-go-client/power/client/p_cloud_service_d_h_c_p"
	"github.com/IBM-Cloud/power-go-client/power/models"
)

// NewIBMPIDhcpClient ...
type IBMPIDhcpClient struct {
	session         *ibmpisession.IBMPISession
	cloudInstanceID string
	authInfo        runtime.ClientAuthInfoWriter
	ctx             context.Context
}

// NewIBMPIDhcpClient ...
func NewIBMPIDhcpClient(ctx context.Context, sess *ibmpisession.IBMPISession, cloudInstanceID string) *IBMPIDhcpClient {
	authInfo := ibmpisession.NewAuth(sess, cloudInstanceID)
	return &IBMPIDhcpClient{
		session:         sess,
		cloudInstanceID: cloudInstanceID,
		authInfo:        authInfo,
		ctx:             ctx,
	}
}

// Create
func (f *IBMPIDhcpClient) Create() (*models.DHCPServer, error) {
	params := p_cloud_service_d_h_c_p.NewPcloudDhcpPostParams().
		WithContext(f.ctx).WithTimeout(helpers.PICreateTimeOut).
		WithCloudInstanceID(f.cloudInstanceID)
	postaccepted, err := f.session.Power.PCloudServiceDHCP.PcloudDhcpPost(params, f.authInfo)
	if err != nil {
		return nil, fmt.Errorf(errors.CreateDchpOperationFailed, f.cloudInstanceID, err)
	}
	if postaccepted != nil && postaccepted.Payload != nil {
		return postaccepted.Payload, nil
	}
	return nil, fmt.Errorf("failed to Create DHCP")
}

// Get
func (f *IBMPIDhcpClient) Get(id string) (*models.DHCPServerDetail, error) {
	params := p_cloud_service_d_h_c_p.NewPcloudDhcpGetParams().
		WithContext(f.ctx).WithTimeout(helpers.PIGetTimeOut).
		WithCloudInstanceID(f.cloudInstanceID).WithDhcpID(id)
	resp, err := f.session.Power.PCloudServiceDHCP.PcloudDhcpGet(params, f.authInfo)
	if err != nil {
		return nil, fmt.Errorf(errors.GetDhcpOperationFailed, id, err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to Get DHCP %s", id)
	}
	return resp.Payload, nil
}

// GetAll
func (f *IBMPIDhcpClient) GetAll() (models.DHCPServers, error) {
	params := p_cloud_service_d_h_c_p.NewPcloudDhcpGetallParams().
		WithContext(f.ctx).WithTimeout(helpers.PIGetTimeOut).
		WithCloudInstanceID(f.cloudInstanceID)
	resp, err := f.session.Power.PCloudServiceDHCP.PcloudDhcpGetall(params, f.authInfo)
	if err != nil {
		return nil, fmt.Errorf("failed to Get all DHCP servers: %v", err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to Get all DHCP servers")
	}
	return resp.Payload, nil
}

// Delete
func (f *IBMPIDhcpClient) Delete(id string) error {
	params := p_cloud_service_d_h_c_p.NewPcloudDhcpDeleteParams().
		WithContext(f.ctx).WithTimeout(helpers.PIDeleteTimeOut).
		WithCloudInstanceID(f.cloudInstanceID).WithDhcpID(id)
	_, err := f.session.Power.PCloudServiceDHCP.PcloudDhcpDelete(params, f.authInfo)
	if err != nil {
		return fmt.Errorf(errors.DeleteDhcpOperationFailed, id, err)
	}
	return nil
}
