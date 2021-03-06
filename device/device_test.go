package device

import (
	"github.com/BPing/aliyun-live-go-sdk/aliyun"
	"github.com/BPing/aliyun-live-go-sdk/device/cdn"
	"github.com/BPing/aliyun-live-go-sdk/device/live"
	"testing"
)

const (
	AccessKeyId     = ""
	AccessKeySecret = ""
	DomainName      = "DomainName"
	AppName         = "AppName"
	PrivateKey      = ""
)

//
func TestDevice(t *testing.T) {

	cert := aliyun.NewCredentials(AccessKeyId, AccessKeySecret)
	streamCert := live.NewStreamCredentials(PrivateKey, live.DefaultStreamTimeout)

	cdnDev, err := GetDevice(CdnDevice, Config{Credentials: cert})
	if _, ok := cdnDev.(*cdn.CDN); err != nil || !ok {
		t.Fatal("get cdn device fail", err, ok)
	}

	cdnDev, err = GetDevice(CdnDevice, Config{})
	if cdnDev != nil || err == nil {
		t.Fatal("get cdn device : param error")
	}

	liveDev, err := GetDevice(LiveDevice,
		Config{Credentials: cert,
			StreamCert: streamCert,
			DomainName: DomainName,
			AppName:    AppName})
	if _, ok := liveDev.(*live.Live); err != nil || !ok {
		t.Fatal("get live device fail", err, ok)
	}

	_, err = GetDevice(LiveDevice,
		Config{Credentials: cert,
			StreamCert: streamCert,
			DomainName: "",
			AppName:    AppName})
	if cdnDev != nil || err == nil {
		t.Fatal("get cdn device : param error")
	}

}
