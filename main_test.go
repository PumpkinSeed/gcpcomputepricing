package gcpcomputepricing

import (
	"testing"
)

func TestGet(t *testing.T) {
	pricing, err := Get()
	if err != nil {
		t.Fatal(err)
	}
	//fmt.Println(*pricing)
	if pricing.Gcp.Compute.GCE.VmsOnDemand.CoresPerCore.Vmimagee2Core.Regions["us-central1"].Prices[0].Nanos != 21811590 {
		t.Error("Broken")
	}
	//for k1, v1 := range pricing.GCP.Compute.GCE {
	//	fmt.Println(k1)
	//	for k2, v2 := range v1 {
	//		fmt.Println("\t", k2)
	//		fmt.Println("\t\t", v2)
	//	}
	//}
}
