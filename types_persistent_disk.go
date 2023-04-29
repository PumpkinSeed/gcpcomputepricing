package gcpcomputepricing

type PersistentDisk struct {
	AsyncReplication AsyncReplication        `json:"async_replication"`
	Snapshots        PersistentDiskSnapshots `json:"snapshots"`
	Diskops          Diskops                 `json:"diskops"`
	Standard         struct {
		Capacity struct {
			Regional struct {
				Regionalstoragepdcapacity Subtype `json:"regionalstoragepdcapacity"`
			} `json:"regional"`
			Storagepdcapacity Subtype `json:"storagepdcapacity"`
		} `json:"capacity"`
		Snapshot struct {
			Storagepdsnapshot Subtype `json:"storagepdsnapshot"`
		} `json:"snapshot"`
	} `json:"standard"`
	Ssd struct {
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
	} `json:"ssd"`
}

type AsyncReplication struct {
	BalancedProtection struct {
		Asyncreplicationprotectionpdbalanced Subtype `json:"asyncreplicationprotectionpdbalanced"`
		Regional                             struct {
			Asyncreplicationprotectionregionalpdbalanced Subtype `json:"asyncreplicationprotectionregionalpdbalanced"`
		} `json:"regional"`
	} `json:"balanced_protection"`
	SsdProtection struct {
		Asyncreplicationprotectionpdssd Subtype `json:"asyncreplicationprotectionpdssd"`
		Regional                        struct {
			Asyncreplicationprotectionregionalpdssd Subtype `json:"asyncreplicationprotectionregionalpdssd"`
		} `json:"regional"`
	} `json:"ssd_protection"`
	Networking struct {
		Asia struct {
			Asynchronousreplicationinterregionnetworkegress Subtype `json:"asynchronousreplicationinterregionnetworkegress"`
		} `json:"asia"`
		Europe struct {
			Asynchronousreplicationinterregionnetworkegress Subtype `json:"asynchronousreplicationinterregionnetworkegress"`
		} `json:"europe"`
		NorthAmerica struct {
			Asynchronousreplicationinterregionnetworkegress Subtype `json:"asynchronousreplicationinterregionnetworkegress"`
		} `json:"north_america"`
		Oceania struct {
			Asynchronousreplicationinterregionnetworkegress Subtype `json:"asynchronousreplicationinterregionnetworkegress"`
		} `json:"oceania"`
	} `json:"networking"`
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
