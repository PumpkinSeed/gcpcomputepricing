package gcpcomputepricing

import "errors"

const (
	E2 = "e2"
)

const (
	OnDemand                = "on-demand"
	Spot                    = "spot"
	Commitment1YearResource = "commitment-1-year-resource"
	Commitment3YearResource = "commitment-3-year-resource"
	Flexible1Year           = "flexible-1-year"
	Flexible3Year           = "flexible-3-year"
)

type Opts struct {
	Type        string
	Commitment  string
	Region      string
	NumOfCPU    uint64
	NumOfMemory uint64
}

func Calculate(p *Pricing, opts Opts) (uint64, error) {
	switch opts.Type {
	case E2:
		return getE2standard(p, opts)
	}
	return 0, errors.New("unknown type")
}

func getE2standard(p *Pricing, opts Opts) (uint64, error) {
	var e2Core Subtype
	var e2Memory Subtype
	switch opts.Commitment {
	case OnDemand:
		e2Core = p.Gcp.Compute.GCE.VmsOnDemand.CoresPerCore.Vmimagee2Core
		e2Memory = p.Gcp.Compute.GCE.VmsOnDemand.MemoryPerGb.Vmimagee2RAM
	case Spot:
		e2Core = p.Gcp.Compute.GCE.VmsPreemptible.CoresPerCore.Vmimagepreemptiblee2Core
		e2Memory = p.Gcp.Compute.GCE.VmsPreemptible.MemoryPerGb.Vmimagepreemptiblee2RAM
	case Commitment1YearResource:
		e2Core = p.Gcp.Compute.GCE.VmsCommit1Year.CoresPerCore.Commitmente2CPU1Yv1
		e2Memory = p.Gcp.Compute.GCE.VmsCommit1Year.MemoryPerGb.Commitmente2RAM1Yv1
	case Commitment3YearResource:
		e2Core = p.Gcp.Compute.GCE.VmsCommit3Year.CoresPerCore.Commitmente2CPU3Yv1
		e2Memory = p.Gcp.Compute.GCE.VmsCommit3Year.MemoryPerGb.Commitmente2RAM3Yv1
	default:
		return 0, errors.New("commitment not supported") // TODO improve error
	}

	var e2CorePricePerRegion uint64 = 0
	if region, ok := e2Core.Regions[opts.Region]; ok {
		if len(region.Prices) > 0 {
			e2CorePricePerRegion = region.Prices[0].Nanos
		}
	} else {
		return 0, errors.New("core not found for that region") // TODO improve error
	}

	var e2MemoryPricePerRegion uint64 = 0
	if region, ok := e2Memory.Regions[opts.Region]; ok {
		if len(region.Prices) > 0 {
			e2MemoryPricePerRegion = region.Prices[0].Nanos
		}
	} else {
		return 0, errors.New("memory not found for that region") // TODO improve error
	}

	var sum uint64 = 0
	sum += e2CorePricePerRegion * opts.NumOfCPU
	sum += e2MemoryPricePerRegion * opts.NumOfMemory
	return sum, nil
}
