package collector

// ThermalFanCollection represents an expanded collection of ThermalFan resources.
type ThermalFanCollection struct {
	Odata
	Members []ThermalFan `json:"Members"`
}

// ProcessorCollection represents an expanded collection of Processor resources.
type ProcessorCollection struct {
	Odata
	Members []Processor `json:"Members"`
}

// H3CProcessorCollection represents an expanded collection of H3CProcessor resources.
type H3CProcessorCollection struct {
	Odata
	Members []H3CProcessor `json:"Members"`
}

// NetworkAdapterCollection represents an expanded collection of NetworkAdapter resources.
type NetworkAdapterCollection struct {
	Odata
	Members []NetworkAdapter `json:"Members"`
}

// NetworkPortCollection represents an expanded collection of NetworkPort resources.
type NetworkPortCollection struct {
	Odata
	Members []NetworkPort `json:"Members"`
}

// PowerSupplyCollection represents an expanded collection of PowerSupply resources.
type PowerSupplyCollection struct {
	Odata
	Members []PowerSupply `json:"Members"`
}

// StorageCollection represents an expanded collection of Storage resources.
type StorageCollection struct {
	Odata
	Members []Storage `json:"Members"`
}

// StorageDriveCollection represents an expanded collection of StorageDrive resources.
type StorageDriveCollection struct {
	Odata
	Members []StorageDrive `json:"Members"`
}

// StorageControllerCollection represents an expanded collection of StorageController resources.
type StorageControllerCollection struct {
	Odata
	Members []StorageController `json:"Members"`
}

// StorageVolumeCollection represents an expanded collection of StorageVolume resources.
type StorageVolumeCollection struct {
	Odata
	Members []StorageVolume `json:"Members"`
}

// MemoryCollection represents an expanded collection of Memory resources.
type MemoryCollection struct {
	Odata
	Members []Memory `json:"Members"`
}
