package gcpcomputepricing

import "errors"

const (
	E2Standard2 = "e2-standard-2"
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
	Type   string
	Region string
}

type machineTypeGetter func(p *Pricing, opts Opts) (uint64, error)

var machineTypeGetters = map[string]machineTypeGetter{
	E2Standard2: getE2standard2,
}

func Get(typ string, p *Pricing, opts Opts) (uint64, error) {
	return machineTypeGetters[typ](p, opts)
}

func getE2standard2(p *Pricing, opts Opts) (uint64, error) {
	return getE2standard(p, opts, 2, 8)
}

func getE2standard4(p *Pricing, opts Opts) (uint64, error) {
	return getE2standard(p, opts, 4, 16)
}

func getE2standard8(p *Pricing, opts Opts) (uint64, error) {
	return getE2standard(p, opts, 8, 32)
}

func getE2standard16(p *Pricing, opts Opts) (uint64, error) {
	return getE2standard(p, opts, 16, 64)
}

func getE2standard32(p *Pricing, opts Opts) (uint64, error) {
	return getE2standard(p, opts, 32, 128)
}

func getE2standard(p *Pricing, opts Opts, numOfCPU, numOfMemory uint64) (uint64, error) {
	e2Core := p.Gcp.Compute.GCE.VmsOnDemand.CoresPerCore.Vmimagee2Core
	e2Memory := p.Gcp.Compute.GCE.VmsOnDemand.MemoryPerGb.Vmimagee2RAM

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
	sum += e2CorePricePerRegion * numOfCPU
	sum += e2MemoryPricePerRegion * numOfMemory
	return sum, nil
}
