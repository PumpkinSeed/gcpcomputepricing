package gcpcomputepricing

type GPUs struct {
	GpusCommit1Year struct {
		L4 struct {
			Commitmentgpunvidial41Yv1 Subtype `json:"commitmentgpunvidial41yv1"`
		} `json:"l4"`
		A100 struct {
			Commitmentgpunvidiateslaa1001Yv1 Subtype `json:"commitmentgpunvidiateslaa1001yv1"`
		} `json:"a100"`
		K80 struct {
			Commitmentgpunvidiateslak801Yv1 Subtype `json:"commitmentgpunvidiateslak801yv1"`
		} `json:"k80"`
		P100 struct {
			Commitmentgpunvidiateslap1001Yv1 Subtype `json:"commitmentgpunvidiateslap1001yv1"`
		} `json:"p100"`
		P4 struct {
			Commitmentgpunvidiateslap41Yv1 Subtype `json:"commitmentgpunvidiateslap41yv1"`
		} `json:"p4"`
		T4 struct {
			Commitmentgpunvidiateslat41Yv1 Subtype `json:"commitmentgpunvidiateslat41yv1"`
		} `json:"t4"`
		V100 struct {
			Commitmentgpunvidiateslav1001Yv1 Subtype `json:"commitmentgpunvidiateslav1001yv1"`
		} `json:"v100"`
	} `json:"gpus_commit_1_year"`
	GpusCommit3Year struct {
		L4 struct {
			Commitmentgpunvidial43Yv1 Subtype `json:"commitmentgpunvidial43yv1"`
		} `json:"l4"`
		A100 struct {
			Commitmentgpunvidiateslaa1003Yv1 Subtype `json:"commitmentgpunvidiateslaa1003yv1"`
		} `json:"a100"`
		K80 struct {
			Commitmentgpunvidiateslak803Yv1 Subtype `json:"commitmentgpunvidiateslak803yv1"`
		} `json:"k80"`
		P100 struct {
			Commitmentgpunvidiateslap1003Yv1 Subtype `json:"commitmentgpunvidiateslap1003yv1"`
		} `json:"p100"`
		P4 struct {
			Commitmentgpunvidiateslap43Yv1 Subtype `json:"commitmentgpunvidiateslap43yv1"`
		} `json:"p4"`
		T4 struct {
			Commitmentgpunvidiateslat43Yv1 Subtype `json:"commitmentgpunvidiateslat43yv1"`
		} `json:"t4"`
		V100 struct {
			Commitmentgpunvidiateslav1003Yv1 Subtype `json:"commitmentgpunvidiateslav1003yv1"`
		} `json:"v100"`
	} `json:"gpus_commit_3_year"`
	GpusOnDemand struct {
		L4 struct {
			Gpunvidial4 Subtype `json:"gpunvidial4"`
		} `json:"l4"`
		A100 struct {
			Gpunvidiateslaa100 Subtype `json:"gpunvidiateslaa100"`
		} `json:"a100"`
		K80 struct {
			Gpunvidiateslak80 Subtype `json:"gpunvidiateslak80"`
		} `json:"k80"`
		P100 struct {
			Gpunvidiateslap100 Subtype `json:"gpunvidiateslap100"`
		} `json:"p100"`
		P4 struct {
			Gpunvidiateslap4 Subtype `json:"gpunvidiateslap4"`
		} `json:"p4"`
		T4 struct {
			Gpunvidiateslat4 Subtype `json:"gpunvidiateslat4"`
		} `json:"t4"`
		V100 struct {
			Gpunvidiateslav100 Subtype `json:"gpunvidiateslav100"`
		} `json:"v100"`
	} `json:"gpus_on_demand"`
	GpusPreemptible struct {
		L4 struct {
			Gpupreemptiblenvidial4 Subtype `json:"gpupreemptiblenvidial4"`
		} `json:"l4"`
		A100 struct {
			Gpupreemptiblenvidiateslaa100 Subtype `json:"gpupreemptiblenvidiateslaa100"`
		} `json:"a100"`
		K80 struct {
			Gpupreemptiblenvidiateslak80 Subtype `json:"gpupreemptiblenvidiateslak80"`
		} `json:"k80"`
		P100 struct {
			Gpupreemptiblenvidiateslap100 Subtype `json:"gpupreemptiblenvidiateslap100"`
		} `json:"p100"`
		P4 struct {
			Gpupreemptiblenvidiateslap4 Subtype `json:"gpupreemptiblenvidiateslap4"`
		} `json:"p4"`
		T4 struct {
			Gpupreemptiblenvidiateslat4 Subtype `json:"gpupreemptiblenvidiateslat4"`
		} `json:"t4"`
		V100 struct {
			Gpupreemptiblenvidiateslav100 Subtype `json:"gpupreemptiblenvidiateslav100"`
		} `json:"v100"`
	} `json:"gpus_preemptible"`
	GpusReservation struct {
		A100 struct {
			Reservationgpunvidiateslaa100 Subtype `json:"reservationgpunvidiateslaa100"`
		} `json:"a100"`
	} `json:"gpus_reservation"`
}
