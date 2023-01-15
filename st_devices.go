package smartthings

import (
	"time"
)

type StDevices struct {
	Links struct{} `json:"_links"`
	Items []struct {
		Allowed      []interface{} `json:"allowed"`
		ChildDevices []struct {
			Allowed  interface{} `json:"allowed"`
			DeviceID string      `json:"deviceId"`
			Profile  struct{}    `json:"profile"`
		} `json:"childDevices,omitempty"`
		Components []struct {
			Capabilities []struct {
				ID      string  `json:"id"`
				Version float64 `json:"version"`
			} `json:"capabilities"`
			Categories []struct {
				CategoryType string `json:"categoryType"`
				Name         string `json:"name"`
			} `json:"categories"`
			ID    string `json:"id"`
			Label string `json:"label"`
		} `json:"components"`
		CreateTime             time.Time `json:"createTime"`
		DeviceID               string    `json:"deviceId"`
		DeviceManufacturerCode string    `json:"deviceManufacturerCode,omitempty"`
		DeviceNetworkType      string    `json:"deviceNetworkType,omitempty"`
		DeviceTypeID           string    `json:"deviceTypeId,omitempty"`
		DeviceTypeName         string    `json:"deviceTypeName,omitempty"`
		Dth                    *struct {
			CompletedSetup       bool   `json:"completedSetup"`
			DeviceNetworkType    string `json:"deviceNetworkType"`
			DeviceTypeID         string `json:"deviceTypeId"`
			DeviceTypeName       string `json:"deviceTypeName"`
			ExecutingLocally     bool   `json:"executingLocally"`
			HubID                string `json:"hubId,omitempty"`
			InstalledGroovyAppID string `json:"installedGroovyAppId,omitempty"`
			NetworkSecurityLevel string `json:"networkSecurityLevel"`
		} `json:"dth,omitempty"`
		Label            string `json:"label"`
		LocationID       string `json:"locationId"`
		ManufacturerName string `json:"manufacturerName"`
		Name             string `json:"name"`
		Ocf              *struct {
			FirmwareVersion           string    `json:"firmwareVersion"`
			HwVersion                 string    `json:"hwVersion"`
			LastSignupTime            time.Time `json:"lastSignupTime"`
			Locale                    string    `json:"locale"`
			ManufacturerName          string    `json:"manufacturerName"`
			ModelNumber               string    `json:"modelNumber"`
			Name                      string    `json:"name"`
			OcfDeviceType             string    `json:"ocfDeviceType"`
			PlatformOS                string    `json:"platformOS"`
			PlatformVersion           string    `json:"platformVersion"`
			SpecVersion               string    `json:"specVersion"`
			VendorID                  string    `json:"vendorId"`
			VerticalDomainSpecVersion string    `json:"verticalDomainSpecVersion"`
		} `json:"ocf,omitempty"`
		OwnerID        string `json:"ownerId,omitempty"`
		ParentDeviceID string `json:"parentDeviceId,omitempty"`
		PresentationID string `json:"presentationId"`
		Profile        *struct {
			ID string `json:"id"`
		} `json:"profile,omitempty"`
		RestrictionTier float64 `json:"restrictionTier"`
		RoomID          string  `json:"roomId,omitempty"`
		Type            string  `json:"type"`
		Viper           *struct {
			HwVersion        string `json:"hwVersion,omitempty"`
			ManufacturerName string `json:"manufacturerName"`
			ModelName        string `json:"modelName"`
			SwVersion        string `json:"swVersion,omitempty"`
			UniqueIdentifier string `json:"uniqueIdentifier,omitempty"`
		} `json:"viper,omitempty"`
		Zwave *struct {
			DriverID             string `json:"driverId"`
			ExecutingLocally     bool   `json:"executingLocally"`
			HubID                string `json:"hubId"`
			NetworkID            string `json:"networkId"`
			NetworkSecurityLevel string `json:"networkSecurityLevel"`
			ProvisioningState    string `json:"provisioningState"`
		} `json:"zwave,omitempty"`
	} `json:"items"`
}
