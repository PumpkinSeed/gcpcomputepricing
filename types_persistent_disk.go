package gcpcomputepricing

type PersistentDisk struct {
	AsyncReplication AsyncReplication        `json:"async_replication"`
	Snapshots        PersistentDiskSnapshots `json:"snapshots"`
	Diskops          Diskops                 `json:"diskops"`
	Standard         PersistentDiskStandard  `json:"standard"`
	SSD              PersistentDiskSSD       `json:"ssd"`
}

type AsyncReplication struct {
	BalancedProtection BalancedProtection         `json:"balanced_protection"`
	SsdProtection      SSDProtection              `json:"ssd_protection"`
	Networking         AsyncReplicationNetworking `json:"networking"`
}

type BalancedProtection struct {
	Asyncreplicationprotectionpdbalanced Subtype                    `json:"asyncreplicationprotectionpdbalanced"`
	Regional                             BalancedProtectionRegional `json:"regional"`
}

type BalancedProtectionRegional struct {
	Asyncreplicationprotectionregionalpdbalanced Subtype `json:"asyncreplicationprotectionregionalpdbalanced"`
}

type SSDProtection struct {
	Asyncreplicationprotectionpdssd Subtype               `json:"asyncreplicationprotectionpdssd"`
	Regional                        SSDProtectionRegional `json:"regional"`
}

type SSDProtectionRegional struct {
	Asyncreplicationprotectionregionalpdssd Subtype `json:"asyncreplicationprotectionregionalpdssd"`
}

type AsyncReplicationNetworking struct {
	Asia         AsyncReplicationNetworkingContinent `json:"asia"`
	Europe       AsyncReplicationNetworkingContinent `json:"europe"`
	NorthAmerica AsyncReplicationNetworkingContinent `json:"north_america"`
	Oceania      AsyncReplicationNetworkingContinent `json:"oceania"`
}

type AsyncReplicationNetworkingContinent struct {
	Asynchronousreplicationinterregionnetworkegress Subtype `json:"asynchronousreplicationinterregionnetworkegress"`
}

type PersistentDiskSnapshots struct {
	Multiregionalsnapshotdownload                     Subtype `json:"multiregionalsnapshotdownload"`
	Multiregionalsnapshotupload                       Subtype `json:"multiregionalsnapshotupload"`
	Storagemultiregionalstandardsnapshotearlydeletion Subtype `json:"storagemultiregionalstandardsnapshotearlydeletion"`
	Storageregionalarchivesnapshotdatastorage         Subtype `json:"storageregionalarchivesnapshotdatastorage"`
	Storageregionalarchivesnapshotearlydeletion       Subtype `json:"storageregionalarchivesnapshotearlydeletion"`
	Storageregionalarchivesnapshotretrieval           Subtype `json:"storageregionalarchivesnapshotretrieval"`
	Storageregionalstandardsnapshotearlydeletion      Subtype `json:"storageregionalstandardsnapshotearlydeletion"`
}

type Diskops struct {
	Pdiorequests Subtype `json:"pdiorequests"`
}

type PersistentDiskStandard struct {
	Capacity struct {
		Regional struct {
			Regionalstoragepdcapacity Subtype `json:"regionalstoragepdcapacity"`
		} `json:"regional"`
		Storagepdcapacity Subtype `json:"storagepdcapacity"`
	} `json:"capacity"`
	Snapshot struct {
		Storagepdsnapshot Subtype `json:"storagepdsnapshot"`
	} `json:"snapshot"`
}

type PersistentDiskSSD struct {
	Capacity struct {
		Regional struct {
			Regionalstoragepdssd Subtype `json:"regionalstoragepdssd"`
		} `json:"regional"`
		Storagepdssd Subtype `json:"storagepdssd"`
		Extreme      struct {
			Storagepdssdextremecapacity Subtype `json:"storagepdssdextremecapacity"`
			Storagepdssdextremeiops     Subtype `json:"storagepdssdextremeiops"`
		} `json:"extreme"`
		Lite struct {
			Storagepdssdlitecapacity Subtype `json:"storagepdssdlitecapacity"`
		} `json:"lite"`
		RegionalLite struct {
			Storageregionalpdssdlitecapacity Subtype `json:"storageregionalpdssdlitecapacity"`
		} `json:"regional_lite"`
	} `json:"capacity"`
}
