package gcpcomputepricing

type Pricing struct {
	Gcp struct {
		Compute struct {
			Gce struct {
				VmsOnDemand struct {
					AdvancedNetworking struct {
						Advnet100Gbpstotalondemand struct {
							Description string            `json:"description"`
							Regions     map[string]Region `json:"regions"`
						} `json:"advnet100gbpstotalondemand"`
						Advnet50Gbpstotalondemand struct {
							Description string            `json:"description"`
							Regions     map[string]Region `json:"regions"`
						} `json:"advnet50gbpstotalondemand"`
						Advnet75Gbpstotalondemand struct {
							Description string            `json:"description"`
							Regions     map[string]Region `json:"regions"`
						} `json:"advnet75gbpstotalondemand"`
					} `json:"advanced_networking"`
					MemoryPerGb struct {
						Vmstatesuspendedram struct {
							Description string            `json:"description"`
							Regions     map[string]Region `json:"regions"`
						} `json:"vmstatesuspendedram"`
						Vmimagea2Highgpuram struct {
							Description string            `json:"description"`
							Regions     map[string]Region `json:"regions"`
						} `json:"vmimagea2highgpuram"`
						Vmimagec2Dcustomextendedram struct {
							Description string            `json:"description"`
							Regions     map[string]Region `json:"regions"`
						} `json:"vmimagec2dcustomextendedram"`
						Vmimagec2Dcustomram struct {
							Description string            `json:"description"`
							Regions     map[string]Region `json:"regions"`
						} `json:"vmimagec2dcustomram"`
						Vmimagec2Dstandardram struct {
							Description string            `json:"description"`
							Regions     map[string]Region `json:"regions"`
						} `json:"vmimagec2dstandardram"`
						C3 struct {
							Vmimagec3Soletenancyram struct {
								Description string            `json:"description"`
								Regions     map[string]Region `json:"regions"`
							} `json:"vmimagec3soletenancyram"`
							Vmimagec3Soletenancyramsoletenancypremium struct {
								Description string            `json:"description"`
								Regions     map[string]Region `json:"regions"`
							} `json:"vmimagec3soletenancyramsoletenancypremium"`
							Vmimagec3Standardram struct {
								Description string            `json:"description"`
								Regions     map[string]Region `json:"regions"`
							} `json:"vmimagec3standardram"`
						} `json:"c3"`
						Vmimagecomputeoptimizedram struct {
							Description string            `json:"description"`
							Regions     map[string]Region `json:"regions"`
						} `json:"vmimagecomputeoptimizedram"`
						Vmimagecustomextendedram struct {
							Description string            `json:"description"`
							Regions     map[string]Region `json:"regions"`
						} `json:"vmimagecustomextendedram"`
						Vmimagecustomram struct {
							Description string            `json:"description"`
							Regions     map[string]Region `json:"regions"`
						} `json:"vmimagecustomram"`
						Vmimagee2RAM struct {
							Description string            `json:"description"`
							Regions     map[string]Region `json:"regions"`
						} `json:"vmimagee2ram"`
						G2 struct {
							Vmimageg2Soletenancyram struct {
								Description string            `json:"description"`
								Regions     map[string]Region `json:"regions"`
							} `json:"vmimageg2soletenancyram"`
							Vmimageg2Soletenancyramsoletenancypremium struct {
								Description string            `json:"description"`
								Regions     map[string]Region `json:"regions"`
							} `json:"vmimageg2soletenancyramsoletenancypremium"`
							Vmimageg2Standardram struct {
								Description string            `json:"description"`
								Regions     map[string]Region `json:"regions"`
							} `json:"vmimageg2standardram"`
						} `json:"g2"`
						Vmimagelargeram struct {
							Description string            `json:"description"`
							Regions     map[string]Region `json:"regions"`
						} `json:"vmimagelargeram"`
						Vmimagelargerammemoryoptimizedupgradepremium struct {
							Description string            `json:"description"`
							Regions     map[string]Region `json:"regions"`
						} `json:"vmimagelargerammemoryoptimizedupgradepremium"`
						M3 struct {
							Vmimagem3Soletenancyram struct {
								Description string            `json:"description"`
								Regions     map[string]Region `json:"regions"`
							} `json:"vmimagem3soletenancyram"`
							Vmimagem3Soletenancyramsoletenancypremium struct {
								Description string            `json:"description"`
								Regions     map[string]Region `json:"regions"`
							} `json:"vmimagem3soletenancyramsoletenancypremium"`
							Vmimagem3Standardram struct {
								Description string            `json:"description"`
								Regions     map[string]Region `json:"regions"`
							} `json:"vmimagem3standardram"`
						} `json:"m3"`
						Vmimagen1Standardram struct {
							Description string            `json:"description"`
							Regions     map[string]Region `json:"regions"`
						} `json:"vmimagen1standardram"`
						Vmimagen2Customextendedram struct {
							Description string            `json:"description"`
							Regions     map[string]Region `json:"regions"`
						} `json:"vmimagen2customextendedram"`
						Vmimagen2Customram struct {
							Description string            `json:"description"`
							Regions     map[string]Region `json:"regions"`
						} `json:"vmimagen2customram"`
						Vmimagen2Dcustomextendedram struct {
							Description string            `json:"description"`
							Regions     map[string]Region `json:"regions"`
						} `json:"vmimagen2dcustomextendedram"`
						Vmimagen2Dcustomram struct {
							Description string            `json:"description"`
							Regions     map[string]Region `json:"regions"`
						} `json:"vmimagen2dcustomram"`
						Vmimagen2Dsoletenancyram struct {
							Description string            `json:"description"`
							Regions     map[string]Region `json:"regions"`
						} `json:"vmimagen2dsoletenancyram"`
						Vmimagen2Dsoletenancyramsoletenancypremium struct {
							Description string            `json:"description"`
							Regions     map[string]Region `json:"regions"`
						} `json:"vmimagen2dsoletenancyramsoletenancypremium"`
						Vmimagen2Dstandardram struct {
							Description string            `json:"description"`
							Regions     map[string]Region `json:"regions"`
						} `json:"vmimagen2dstandardram"`
						Vmimagen2Soletenancyram struct {
							Description string            `json:"description"`
							Regions     map[string]Region `json:"regions"`
						} `json:"vmimagen2soletenancyram"`
						Vmimagen2Standardram struct {
							Description string            `json:"description"`
							Regions     map[string]Region `json:"regions"`
						} `json:"vmimagen2standardram"`
						Vmimagesoletenancyram struct {
							Description string            `json:"description"`
							Regions     map[string]Region `json:"regions"`
						} `json:"vmimagesoletenancyram"`
						Vmimagesoletenancyramsoletenancypremium struct {
							Description string            `json:"description"`
							Regions     map[string]Region `json:"regions"`
						} `json:"vmimagesoletenancyramsoletenancypremium"`
						T2A struct {
							Vmimaget2Astandardram struct {
								Description string            `json:"description"`
								Regions     map[string]Region `json:"regions"`
							} `json:"vmimaget2astandardram"`
						} `json:"t2a"`
						Vmimaget2Dstandardram struct {
							Description string            `json:"description"`
							Regions     map[string]Region `json:"regions"`
						} `json:"vmimaget2dstandardram"`
					} `json:"memory:_per_gb"`
					CoresPerCore struct {
						Vmimagea2Highgpucore struct {
							Description string            `json:"description"`
							Regions     map[string]Region `json:"regions"`
						} `json:"vmimagea2highgpucore"`
						Vmimagec2Dcustomcore struct {
							Description string            `json:"description"`
							Regions     map[string]Region `json:"regions"`
						} `json:"vmimagec2dcustomcore"`
						Vmimagec2Dstandardcore struct {
							Description string            `json:"description"`
							Regions     map[string]Region `json:"regions"`
						} `json:"vmimagec2dstandardcore"`
						C3 struct {
							Vmimagec3Soletenancycore struct {
								Description string            `json:"description"`
								Regions     map[string]Region `json:"regions"`
							} `json:"vmimagec3soletenancycore"`
							Vmimagec3Soletenancycoresoletenancypremium struct {
								Description string            `json:"description"`
								Regions     map[string]Region `json:"regions"`
							} `json:"vmimagec3soletenancycoresoletenancypremium"`
							Vmimagec3Standardcore struct {
								Description string            `json:"description"`
								Regions     map[string]Region `json:"regions"`
							} `json:"vmimagec3standardcore"`
						} `json:"c3"`
						Vmimagecomputeoptimizedcore struct {
							Description string            `json:"description"`
							Regions     map[string]Region `json:"regions"`
						} `json:"vmimagecomputeoptimizedcore"`
						Vmimagecustomcore struct {
							Description string            `json:"description"`
							Regions     map[string]Region `json:"regions"`
						} `json:"vmimagecustomcore"`
						Vmimagee2Core struct {
							Description string            `json:"description"`
							Regions     map[string]Region `json:"regions"`
						} `json:"vmimagee2core"`
						Vmimagef1Micro struct {
							Description string            `json:"description"`
							Regions     map[string]Region `json:"regions"`
						} `json:"vmimagef1micro"`
						Vmimageg1Small struct {
							Description string            `json:"description"`
							Regions     map[string]Region `json:"regions"`
						} `json:"vmimageg1small"`
						G2 struct {
							Vmimageg2Soletenancycore struct {
								Description string            `json:"description"`
								Regions     map[string]Region `json:"regions"`
							} `json:"vmimageg2soletenancycore"`
							Vmimageg2Soletenancycoresoletenancypremium struct {
								Description string            `json:"description"`
								Regions     map[string]Region `json:"regions"`
							} `json:"vmimageg2soletenancycoresoletenancypremium"`
							Vmimageg2Standardcore struct {
								Description string            `json:"description"`
								Regions     map[string]Region `json:"regions"`
							} `json:"vmimageg2standardcore"`
						} `json:"g2"`
						Vmimagelargecore struct {
							Description string            `json:"description"`
							Regions     map[string]Region `json:"regions"`
						} `json:"vmimagelargecore"`
						Vmimagelargecorememoryoptimizedupgradepremium struct {
							Description string            `json:"description"`
							Regions     map[string]Region `json:"regions"`
						} `json:"vmimagelargecorememoryoptimizedupgradepremium"`
						M3 struct {
							Vmimagem3Soletenancycore struct {
								Description string            `json:"description"`
								Regions     map[string]Region `json:"regions"`
							} `json:"vmimagem3soletenancycore"`
							Vmimagem3Soletenancycoresoletenancypremium struct {
								Description string            `json:"description"`
								Regions     map[string]Region `json:"regions"`
							} `json:"vmimagem3soletenancycoresoletenancypremium"`
							Vmimagem3Standardcore struct {
								Description string            `json:"description"`
								Regions     map[string]Region `json:"regions"`
							} `json:"vmimagem3standardcore"`
						} `json:"m3"`
						Vmimagen1Standardcore struct {
							Description string            `json:"description"`
							Regions     map[string]Region `json:"regions"`
						} `json:"vmimagen1standardcore"`
						Vmimagen2Customcore struct {
							Description string            `json:"description"`
							Regions     map[string]Region `json:"regions"`
						} `json:"vmimagen2customcore"`
						Vmimagen2Customextendedcore struct {
							Description string            `json:"description"`
							Regions     map[string]Region `json:"regions"`
						} `json:"vmimagen2customextendedcore"`
						Vmimagen2Dcustomcore struct {
							Description string            `json:"description"`
							Regions     map[string]Region `json:"regions"`
						} `json:"vmimagen2dcustomcore"`
						Vmimagen2Dsoletenancycore struct {
							Description string            `json:"description"`
							Regions     map[string]Region `json:"regions"`
						} `json:"vmimagen2dsoletenancycore"`
						Vmimagen2Dsoletenancycoresoletenancypremium struct {
							Description string            `json:"description"`
							Regions     map[string]Region `json:"regions"`
						} `json:"vmimagen2dsoletenancycoresoletenancypremium"`
						Vmimagen2Dstandardcore struct {
							Description string            `json:"description"`
							Regions     map[string]Region `json:"regions"`
						} `json:"vmimagen2dstandardcore"`
						Vmimagen2Soletenancycore struct {
							Description string            `json:"description"`
							Regions     map[string]Region `json:"regions"`
						} `json:"vmimagen2soletenancycore"`
						Vmimagen2Soletenancycoresoletenancypremium struct {
							Description string            `json:"description"`
							Regions     map[string]Region `json:"regions"`
						} `json:"vmimagen2soletenancycoresoletenancypremium"`
						Vmimagen2Standardcore struct {
							Description string            `json:"description"`
							Regions     map[string]Region `json:"regions"`
						} `json:"vmimagen2standardcore"`
						Vmimagesoletenancycore struct {
							Description string            `json:"description"`
							Regions     map[string]Region `json:"regions"`
						} `json:"vmimagesoletenancycore"`
						Vmimagesoletenancycoresoletenancypremium struct {
							Description string            `json:"description"`
							Regions     map[string]Region `json:"regions"`
						} `json:"vmimagesoletenancycoresoletenancypremium"`
						T2A struct {
							Vmimaget2Astandardcore struct {
								Description string            `json:"description"`
								Regions     map[string]Region `json:"regions"`
							} `json:"vmimaget2astandardcore"`
						} `json:"t2a"`
						Vmimaget2Dstandardcore struct {
							Description string            `json:"description"`
							Regions     map[string]Region `json:"regions"`
						} `json:"vmimaget2dstandardcore"`
					} `json:"cores:_per_core"`
					Vmimagecomputeoptimizedsoletenancycore struct {
						Description string            `json:"description"`
						Regions     map[string]Region `json:"regions"`
					} `json:"vmimagecomputeoptimizedsoletenancycore"`
					Vmimagecomputeoptimizedsoletenancycoresoletenancypremium struct {
						Description string            `json:"description"`
						Regions     map[string]Region `json:"regions"`
					} `json:"vmimagecomputeoptimizedsoletenancycoresoletenancypremium"`
					Vmimagecomputeoptimizedsoletenancyram struct {
						Description string            `json:"description"`
						Regions     map[string]Region `json:"regions"`
					} `json:"vmimagecomputeoptimizedsoletenancyram"`
					Vmimagecomputeoptimizedsoletenancyramsoletenancypremium struct {
						Description string            `json:"description"`
						Regions     map[string]Region `json:"regions"`
					} `json:"vmimagecomputeoptimizedsoletenancyramsoletenancypremium"`
					Vmimagelargesoletenancycore struct {
						Description string            `json:"description"`
						Regions     map[string]Region `json:"regions"`
					} `json:"vmimagelargesoletenancycore"`
					Vmimagelargesoletenancycoresoletenancypremium struct {
						Description string            `json:"description"`
						Regions     map[string]Region `json:"regions"`
					} `json:"vmimagelargesoletenancycoresoletenancypremium"`
					Vmimagelargesoletenancyram struct {
						Description string            `json:"description"`
						Regions     map[string]Region `json:"regions"`
					} `json:"vmimagelargesoletenancyram"`
					Vmimagelargesoletenancyramsoletenancypremium struct {
						Description string            `json:"description"`
						Regions     map[string]Region `json:"regions"`
					} `json:"vmimagelargesoletenancyramsoletenancypremium"`
					Vmimagen2Soletenancycoresoletenancyovercommitpremium struct {
						Description string            `json:"description"`
						Regions     map[string]Region `json:"regions"`
					} `json:"vmimagen2soletenancycoresoletenancyovercommitpremium"`
					Vmimagen2Soletenancyramsoletenancyovercommitpremium struct {
						Description string            `json:"description"`
						Regions     map[string]Region `json:"regions"`
					} `json:"vmimagen2soletenancyramsoletenancyovercommitpremium"`
					Vmimagen2Soletenancyramsoletenancypremium struct {
						Description string            `json:"description"`
						Regions     map[string]Region `json:"regions"`
					} `json:"vmimagen2soletenancyramsoletenancypremium"`
					Vmimagesoletenancycoresoletenancyovercommitpremium struct {
						Description string            `json:"description"`
						Regions     map[string]Region `json:"regions"`
					} `json:"vmimagesoletenancycoresoletenancyovercommitpremium"`
					Vmimagesoletenancyramsoletenancyovercommitpremium struct {
						Description string            `json:"description"`
						Regions     map[string]Region `json:"regions"`
					} `json:"vmimagesoletenancyramsoletenancyovercommitpremium"`
				} `json:"vms_on_demand"`
				VmsPreemptible struct {
					AdvancedNetworking struct {
						Advnet100Gbpstotalpreemptible struct {
							Description string            `json:"description"`
							Regions     map[string]Region `json:"regions"`
						} `json:"advnet100gbpstotalpreemptible"`
						Advnet50Gbpstotalpreemptible struct {
							Description string            `json:"description"`
							Regions     map[string]Region `json:"regions"`
						} `json:"advnet50gbpstotalpreemptible"`
						Advnet75Gbpstotalpreemptible struct {
							Description string            `json:"description"`
							Regions     map[string]Region `json:"regions"`
						} `json:"advnet75gbpstotalpreemptible"`
					} `json:"advanced_networking"`
					Cores1To64 struct {
						Highcpu struct {
							Vmimagepreemptiblea2Highgpucore struct {
								Description string            `json:"description"`
								Regions     map[string]Region `json:"regions"`
							} `json:"vmimagepreemptiblea2highgpucore"`
						} `json:"highcpu"`
					} `json:"cores:_1_to_64"`
					MemoryPerGb struct {
						Vmimagepreemptiblea2Highgpuram struct {
							Description string            `json:"description"`
							Regions     map[string]Region `json:"regions"`
						} `json:"vmimagepreemptiblea2highgpuram"`
						Vmimagepreemptiblec2Dcustomextendedram struct {
							Description string            `json:"description"`
							Regions     map[string]Region `json:"regions"`
						} `json:"vmimagepreemptiblec2dcustomextendedram"`
						Vmimagepreemptiblec2Dcustomram struct {
							Description string            `json:"description"`
							Regions     map[string]Region `json:"regions"`
						} `json:"vmimagepreemptiblec2dcustomram"`
						Vmimagepreemptiblec2Dstandardram struct {
							Description string            `json:"description"`
							Regions     map[string]Region `json:"regions"`
						} `json:"vmimagepreemptiblec2dstandardram"`
						C3 struct {
							Vmimagepreemptiblec3Standardram struct {
								Description string            `json:"description"`
								Regions     map[string]Region `json:"regions"`
							} `json:"vmimagepreemptiblec3standardram"`
						} `json:"c3"`
						Vmimagepreemptiblecomputeoptimizedram struct {
							Description string            `json:"description"`
							Regions     map[string]Region `json:"regions"`
						} `json:"vmimagepreemptiblecomputeoptimizedram"`
						Vmimagepreemptiblecustomextendedram struct {
							Description string            `json:"description"`
							Regions     map[string]Region `json:"regions"`
						} `json:"vmimagepreemptiblecustomextendedram"`
						Vmimagepreemptiblecustomram struct {
							Description string            `json:"description"`
							Regions     map[string]Region `json:"regions"`
						} `json:"vmimagepreemptiblecustomram"`
						Vmimagepreemptiblee2RAM struct {
							Description string            `json:"description"`
							Regions     map[string]Region `json:"regions"`
						} `json:"vmimagepreemptiblee2ram"`
						G2 struct {
							Vmimagepreemptibleg2Standardram struct {
								Description string            `json:"description"`
								Regions     map[string]Region `json:"regions"`
							} `json:"vmimagepreemptibleg2standardram"`
						} `json:"g2"`
						Vmimagepreemptiblelargeram struct {
							Description string            `json:"description"`
							Regions     map[string]Region `json:"regions"`
						} `json:"vmimagepreemptiblelargeram"`
						M3 struct {
							Vmimagepreemptiblem3Standardram struct {
								Description string            `json:"description"`
								Regions     map[string]Region `json:"regions"`
							} `json:"vmimagepreemptiblem3standardram"`
						} `json:"m3"`
						Vmimagepreemptiblen1Standardram struct {
							Description string            `json:"description"`
							Regions     map[string]Region `json:"regions"`
						} `json:"vmimagepreemptiblen1standardram"`
						Vmimagepreemptiblen2Customextendedram struct {
							Description string            `json:"description"`
							Regions     map[string]Region `json:"regions"`
						} `json:"vmimagepreemptiblen2customextendedram"`
						Vmimagepreemptiblen2Customram struct {
							Description string            `json:"description"`
							Regions     map[string]Region `json:"regions"`
						} `json:"vmimagepreemptiblen2customram"`
						Vmimagepreemptiblen2Dcustomextendedram struct {
							Description string            `json:"description"`
							Regions     map[string]Region `json:"regions"`
						} `json:"vmimagepreemptiblen2dcustomextendedram"`
						Vmimagepreemptiblen2Dcustomram struct {
							Description string            `json:"description"`
							Regions     map[string]Region `json:"regions"`
						} `json:"vmimagepreemptiblen2dcustomram"`
						Vmimagepreemptiblen2Dstandardram struct {
							Description string            `json:"description"`
							Regions     map[string]Region `json:"regions"`
						} `json:"vmimagepreemptiblen2dstandardram"`
						Vmimagepreemptiblen2Standardram struct {
							Description string            `json:"description"`
							Regions     map[string]Region `json:"regions"`
						} `json:"vmimagepreemptiblen2standardram"`
						T2A struct {
							Vmimagepreemptiblet2Astandardram struct {
								Description string            `json:"description"`
								Regions     map[string]Region `json:"regions"`
							} `json:"vmimagepreemptiblet2astandardram"`
						} `json:"t2a"`
						Vmimagepreemptiblet2Dstandardram struct {
							Description string            `json:"description"`
							Regions     map[string]Region `json:"regions"`
						} `json:"vmimagepreemptiblet2dstandardram"`
					} `json:"memory:_per_gb"`
					CoresPerCore struct {
						Vmimagepreemptiblec2Dcustomcore struct {
							Description string            `json:"description"`
							Regions     map[string]Region `json:"regions"`
						} `json:"vmimagepreemptiblec2dcustomcore"`
						Vmimagepreemptiblec2Dstandardcore struct {
							Description string            `json:"description"`
							Regions     map[string]Region `json:"regions"`
						} `json:"vmimagepreemptiblec2dstandardcore"`
						C3 struct {
							Vmimagepreemptiblec3Standardcore struct {
								Description string            `json:"description"`
								Regions     map[string]Region `json:"regions"`
							} `json:"vmimagepreemptiblec3standardcore"`
						} `json:"c3"`
						Vmimagepreemptiblecomputeoptimizedcore struct {
							Description string            `json:"description"`
							Regions     map[string]Region `json:"regions"`
						} `json:"vmimagepreemptiblecomputeoptimizedcore"`
						Vmimagepreemptiblecustomcore struct {
							Description string            `json:"description"`
							Regions     map[string]Region `json:"regions"`
						} `json:"vmimagepreemptiblecustomcore"`
						Vmimagepreemptiblecustomextendedcore struct {
							Description string            `json:"description"`
							Regions     map[string]Region `json:"regions"`
						} `json:"vmimagepreemptiblecustomextendedcore"`
						Vmimagepreemptiblee2Core struct {
							Description string            `json:"description"`
							Regions     map[string]Region `json:"regions"`
						} `json:"vmimagepreemptiblee2core"`
						Vmimagepreemptiblef1Micro struct {
							Description string            `json:"description"`
							Regions     map[string]Region `json:"regions"`
						} `json:"vmimagepreemptiblef1micro"`
						Vmimagepreemptibleg1Small struct {
							Description string            `json:"description"`
							Regions     map[string]Region `json:"regions"`
						} `json:"vmimagepreemptibleg1small"`
						G2 struct {
							Vmimagepreemptibleg2Standardcore struct {
								Description string            `json:"description"`
								Regions     map[string]Region `json:"regions"`
							} `json:"vmimagepreemptibleg2standardcore"`
						} `json:"g2"`
						Vmimagepreemptiblelargecore struct {
							Description string            `json:"description"`
							Regions     map[string]Region `json:"regions"`
						} `json:"vmimagepreemptiblelargecore"`
						M3 struct {
							Vmimagepreemptiblem3Standardcore struct {
								Description string            `json:"description"`
								Regions     map[string]Region `json:"regions"`
							} `json:"vmimagepreemptiblem3standardcore"`
						} `json:"m3"`
						Vmimagepreemptiblen1Standardcore struct {
							Description string            `json:"description"`
							Regions     map[string]Region `json:"regions"`
						} `json:"vmimagepreemptiblen1standardcore"`
						Vmimagepreemptiblen2Customcore struct {
							Description string            `json:"description"`
							Regions     map[string]Region `json:"regions"`
						} `json:"vmimagepreemptiblen2customcore"`
						Vmimagepreemptiblen2Customextendedcore struct {
							Description string            `json:"description"`
							Regions     map[string]Region `json:"regions"`
						} `json:"vmimagepreemptiblen2customextendedcore"`
						Vmimagepreemptiblen2Dcustomcore struct {
							Description string            `json:"description"`
							Regions     map[string]Region `json:"regions"`
						} `json:"vmimagepreemptiblen2dcustomcore"`
						Vmimagepreemptiblen2Dstandardcore struct {
							Description string            `json:"description"`
							Regions     map[string]Region `json:"regions"`
						} `json:"vmimagepreemptiblen2dstandardcore"`
						Vmimagepreemptiblen2Standardcore struct {
							Description string            `json:"description"`
							Regions     map[string]Region `json:"regions"`
						} `json:"vmimagepreemptiblen2standardcore"`
						T2A struct {
							Vmimagepreemptiblet2Astandardcore struct {
								Description string            `json:"description"`
								Regions     map[string]Region `json:"regions"`
							} `json:"vmimagepreemptiblet2astandardcore"`
						} `json:"t2a"`
						Vmimagepreemptiblet2Dstandardcore struct {
							Description string            `json:"description"`
							Regions     map[string]Region `json:"regions"`
						} `json:"vmimagepreemptiblet2dstandardcore"`
					} `json:"cores:_per_core"`
				} `json:"vms_preemptible"`
				VmsCommit1Year struct {
					CoresPerCore struct {
						Commitmenta2Highgpucpu1Yv1 struct {
							Description string            `json:"description"`
							Regions     map[string]Region `json:"regions"`
						} `json:"commitmenta2highgpucpu1yv1"`
						Commitmentc2Dcpu1Yv1 struct {
							Description string            `json:"description"`
							Regions     map[string]Region `json:"regions"`
						} `json:"commitmentc2dcpu1yv1"`
						C3 struct {
							Commitmentc3CPU1Yv1 struct {
								Description string            `json:"description"`
								Regions     map[string]Region `json:"regions"`
							} `json:"commitmentc3cpu1yv1"`
						} `json:"c3"`
						Commitmentcpucomputeoptimized1Yv1 struct {
							Description string            `json:"description"`
							Regions     map[string]Region `json:"regions"`
						} `json:"commitmentcpucomputeoptimized1yv1"`
						Commitmentcpulargeinstance1Yv1 struct {
							Description string            `json:"description"`
							Regions     map[string]Region `json:"regions"`
						} `json:"commitmentcpulargeinstance1yv1"`
						Commitmentcpu1Yv1 struct {
							Description string            `json:"description"`
							Regions     map[string]Region `json:"regions"`
						} `json:"commitmentcpu1yv1"`
						Commitmente2CPU1Yv1 struct {
							Description string            `json:"description"`
							Regions     map[string]Region `json:"regions"`
						} `json:"commitmente2cpu1yv1"`
						G2 struct {
							Commitmentg2CPU1Yv1 struct {
								Description string            `json:"description"`
								Regions     map[string]Region `json:"regions"`
							} `json:"commitmentg2cpu1yv1"`
						} `json:"g2"`
						M3 struct {
							Commitmentm3CPU1Yv1 struct {
								Description string            `json:"description"`
								Regions     map[string]Region `json:"regions"`
							} `json:"commitmentm3cpu1yv1"`
						} `json:"m3"`
						Commitmentn2CPU1Yv1 struct {
							Description string            `json:"description"`
							Regions     map[string]Region `json:"regions"`
						} `json:"commitmentn2cpu1yv1"`
						Commitmentn2Dcpu1Yv1 struct {
							Description string            `json:"description"`
							Regions     map[string]Region `json:"regions"`
						} `json:"commitmentn2dcpu1yv1"`
						Commitmentt2Dcpu1Yv1 struct {
							Description string            `json:"description"`
							Regions     map[string]Region `json:"regions"`
						} `json:"commitmentt2dcpu1yv1"`
					} `json:"cores:_per_core"`
					MemoryPerGb struct {
						Commitmenta2Highgpuram1Yv1 struct {
							Description string            `json:"description"`
							Regions     map[string]Region `json:"regions"`
						} `json:"commitmenta2highgpuram1yv1"`
						Commitmentc2Dram1Yv1 struct {
							Description string            `json:"description"`
							Regions     map[string]Region `json:"regions"`
						} `json:"commitmentc2dram1yv1"`
						C3 struct {
							Commitmentc3RAM1Yv1 struct {
								Description string            `json:"description"`
								Regions     map[string]Region `json:"regions"`
							} `json:"commitmentc3ram1yv1"`
						} `json:"c3"`
						Commitmente2RAM1Yv1 struct {
							Description string            `json:"description"`
							Regions     map[string]Region `json:"regions"`
						} `json:"commitmente2ram1yv1"`
						G2 struct {
							Commitmentg2RAM1Yv1 struct {
								Description string            `json:"description"`
								Regions     map[string]Region `json:"regions"`
							} `json:"commitmentg2ram1yv1"`
						} `json:"g2"`
						M3 struct {
							Commitmentm3RAM1Yv1 struct {
								Description string            `json:"description"`
								Regions     map[string]Region `json:"regions"`
							} `json:"commitmentm3ram1yv1"`
						} `json:"m3"`
						Commitmentn2Dram1Yv1 struct {
							Description string            `json:"description"`
							Regions     map[string]Region `json:"regions"`
						} `json:"commitmentn2dram1yv1"`
						Commitmentn2RAM1Yv1 struct {
							Description string            `json:"description"`
							Regions     map[string]Region `json:"regions"`
						} `json:"commitmentn2ram1yv1"`
						Commitmentramcomputeoptimized1Yv1 struct {
							Description string            `json:"description"`
							Regions     map[string]Region `json:"regions"`
						} `json:"commitmentramcomputeoptimized1yv1"`
						Commitmentramlargeinstance1Yv1 struct {
							Description string            `json:"description"`
							Regions     map[string]Region `json:"regions"`
						} `json:"commitmentramlargeinstance1yv1"`
						Commitmentram1Yv1 struct {
							Description string            `json:"description"`
							Regions     map[string]Region `json:"regions"`
						} `json:"commitmentram1yv1"`
						Commitmentt2Dram1Yv1 struct {
							Description string            `json:"description"`
							Regions     map[string]Region `json:"regions"`
						} `json:"commitmentt2dram1yv1"`
					} `json:"memory:_per_gb"`
					Vmwareengineucs12Moprepay struct {
						Description string            `json:"description"`
						Regions     map[string]Region `json:"regions"`
					} `json:"vmwareengineucs12moprepay"`
				} `json:"vms_commit_1_year"`
				VmsCommit3Year struct {
					CoresPerCore struct {
						Commitmenta2Highgpucpu3Yv1 struct {
							Description string            `json:"description"`
							Regions     map[string]Region `json:"regions"`
						} `json:"commitmenta2highgpucpu3yv1"`
						Commitmentc2Dcpu3Yv1 struct {
							Description string            `json:"description"`
							Regions     map[string]Region `json:"regions"`
						} `json:"commitmentc2dcpu3yv1"`
						C3 struct {
							Commitmentc3CPU3Yv1 struct {
								Description string            `json:"description"`
								Regions     map[string]Region `json:"regions"`
							} `json:"commitmentc3cpu3yv1"`
						} `json:"c3"`
						Commitmentcpucomputeoptimized3Yv1 struct {
							Description string            `json:"description"`
							Regions     map[string]Region `json:"regions"`
						} `json:"commitmentcpucomputeoptimized3yv1"`
						Commitmentcpulargeinstance1Yv1 struct {
							Description string            `json:"description"`
							Regions     map[string]Region `json:"regions"`
						} `json:"commitmentcpulargeinstance1yv1"`
						Commitmentcpulargeinstance3Yv1 struct {
							Description string            `json:"description"`
							Regions     map[string]Region `json:"regions"`
						} `json:"commitmentcpulargeinstance3yv1"`
						Commitmentcpu3Yv1 struct {
							Description string            `json:"description"`
							Regions     map[string]Region `json:"regions"`
						} `json:"commitmentcpu3yv1"`
						Commitmente2CPU3Yv1 struct {
							Description string            `json:"description"`
							Regions     map[string]Region `json:"regions"`
						} `json:"commitmente2cpu3yv1"`
						G2 struct {
							Commitmentg2CPU3Yv1 struct {
								Description string            `json:"description"`
								Regions     map[string]Region `json:"regions"`
							} `json:"commitmentg2cpu3yv1"`
						} `json:"g2"`
						M3 struct {
							Commitmentm3CPU3Yv1 struct {
								Description string            `json:"description"`
								Regions     map[string]Region `json:"regions"`
							} `json:"commitmentm3cpu3yv1"`
						} `json:"m3"`
						Commitmentn2CPU3Yv1 struct {
							Description string            `json:"description"`
							Regions     map[string]Region `json:"regions"`
						} `json:"commitmentn2cpu3yv1"`
						Commitmentn2Dcpu3Yv1 struct {
							Description string            `json:"description"`
							Regions     map[string]Region `json:"regions"`
						} `json:"commitmentn2dcpu3yv1"`
						Commitmentt2Dcpu3Yv1 struct {
							Description string            `json:"description"`
							Regions     map[string]Region `json:"regions"`
						} `json:"commitmentt2dcpu3yv1"`
					} `json:"cores:_per_core"`
					MemoryPerGb struct {
						Commitmenta2Highgpuram3Yv1 struct {
							Description string            `json:"description"`
							Regions     map[string]Region `json:"regions"`
						} `json:"commitmenta2highgpuram3yv1"`
						Commitmentc2Dram3Yv1 struct {
							Description string            `json:"description"`
							Regions     map[string]Region `json:"regions"`
						} `json:"commitmentc2dram3yv1"`
						C3 struct {
							Commitmentc3RAM3Yv1 struct {
								Description string            `json:"description"`
								Regions     map[string]Region `json:"regions"`
							} `json:"commitmentc3ram3yv1"`
						} `json:"c3"`
						Commitmente2RAM3Yv1 struct {
							Description string            `json:"description"`
							Regions     map[string]Region `json:"regions"`
						} `json:"commitmente2ram3yv1"`
						G2 struct {
							Commitmentg2RAM3Yv1 struct {
								Description string            `json:"description"`
								Regions     map[string]Region `json:"regions"`
							} `json:"commitmentg2ram3yv1"`
						} `json:"g2"`
						M3 struct {
							Commitmentm3RAM3Yv1 struct {
								Description string            `json:"description"`
								Regions     map[string]Region `json:"regions"`
							} `json:"commitmentm3ram3yv1"`
						} `json:"m3"`
						Commitmentn2Dram3Yv1 struct {
							Description string            `json:"description"`
							Regions     map[string]Region `json:"regions"`
						} `json:"commitmentn2dram3yv1"`
						Commitmentn2RAM3Yv1 struct {
							Description string            `json:"description"`
							Regions     map[string]Region `json:"regions"`
						} `json:"commitmentn2ram3yv1"`
						Commitmentramcomputeoptimized3Yv1 struct {
							Description string            `json:"description"`
							Regions     map[string]Region `json:"regions"`
						} `json:"commitmentramcomputeoptimized3yv1"`
						Commitmentramlargeinstance3Yv1 struct {
							Description string            `json:"description"`
							Regions     map[string]Region `json:"regions"`
						} `json:"commitmentramlargeinstance3yv1"`
						Commitmentram3Yv1 struct {
							Description string            `json:"description"`
							Regions     map[string]Region `json:"regions"`
						} `json:"commitmentram3yv1"`
						Commitmentt2Dram3Yv1 struct {
							Description string            `json:"description"`
							Regions     map[string]Region `json:"regions"`
						} `json:"commitmentt2dram3yv1"`
					} `json:"memory:_per_gb"`
					Vmwareengineucs36Moprepay struct {
						Description string            `json:"description"`
						Regions     map[string]Region `json:"regions"`
					} `json:"vmwareengineucs36moprepay"`
				} `json:"vms_commit_3_year"`
				NetworkOther struct {
					Externalip struct {
						Description string            `json:"description"`
						Regions     map[string]Region `json:"regions"`
					} `json:"externalip"`
					Externalippreemptible struct {
						Description string            `json:"description"`
						Regions     map[string]Region `json:"regions"`
					} `json:"externalippreemptible"`
				} `json:"network_other"`
				FlexCommit1Year struct {
					Gcecommitmentsucs12Mo struct {
						Description string            `json:"description"`
						Regions     map[string]Region `json:"regions"`
					} `json:"gcecommitmentsucs12mo"`
				} `json:"flex_commit_1_year"`
				FlexCommit3Year struct {
					Gcecommitmentsucs36Mo struct {
						Description string            `json:"description"`
						Regions     map[string]Region `json:"regions"`
					} `json:"gcecommitmentsucs36mo"`
				} `json:"flex_commit_3_year"`
				PremiumImage struct {
					Microsoft struct {
						Windows struct {
							Licensed1656378918552316916Core struct {
								Description string            `json:"description"`
								Regions     map[string]Region `json:"regions"`
							} `json:"licensed1656378918552316916core"`
							Licensed1656378918552316916F1Micro struct {
								Description string            `json:"description"`
								Regions     map[string]Region `json:"regions"`
							} `json:"licensed1656378918552316916f1micro"`
							Licensed1656378918552316916G1Small struct {
								Description string            `json:"description"`
								Regions     map[string]Region `json:"regions"`
							} `json:"licensed1656378918552316916g1small"`
							Licensed3284763237085719542Core struct {
								Description string            `json:"description"`
								Regions     map[string]Region `json:"regions"`
							} `json:"licensed3284763237085719542core"`
							Licensed3284763237085719542F1Micro struct {
								Description string            `json:"description"`
								Regions     map[string]Region `json:"regions"`
							} `json:"licensed3284763237085719542f1micro"`
							Licensed3284763237085719542G1Small struct {
								Description string            `json:"description"`
								Regions     map[string]Region `json:"regions"`
							} `json:"licensed3284763237085719542g1small"`
							Licensed4819555115818134498Core struct {
								Description string            `json:"description"`
								Regions     map[string]Region `json:"regions"`
							} `json:"licensed4819555115818134498core"`
							Licensed4819555115818134498F1Micro struct {
								Description string            `json:"description"`
								Regions     map[string]Region `json:"regions"`
							} `json:"licensed4819555115818134498f1micro"`
							Licensed4819555115818134498G1Small struct {
								Description string            `json:"description"`
								Regions     map[string]Region `json:"regions"`
							} `json:"licensed4819555115818134498g1small"`
							Licensed4874454843789519845Core struct {
								Description string            `json:"description"`
								Regions     map[string]Region `json:"regions"`
							} `json:"licensed4874454843789519845core"`
							Licensed4874454843789519845F1Micro struct {
								Description string            `json:"description"`
								Regions     map[string]Region `json:"regions"`
							} `json:"licensed4874454843789519845f1micro"`
							Licensed4874454843789519845G1Small struct {
								Description string            `json:"description"`
								Regions     map[string]Region `json:"regions"`
							} `json:"licensed4874454843789519845g1small"`
							Licensed6107784707477449232Core struct {
								Description string            `json:"description"`
								Regions     map[string]Region `json:"regions"`
							} `json:"licensed6107784707477449232core"`
							Licensed6107784707477449232F1Micro struct {
								Description string            `json:"description"`
								Regions     map[string]Region `json:"regions"`
							} `json:"licensed6107784707477449232f1micro"`
							Licensed6107784707477449232G1Small struct {
								Description string            `json:"description"`
								Regions     map[string]Region `json:"regions"`
							} `json:"licensed6107784707477449232g1small"`
							Licensed7695108898142923768Core struct {
								Description string            `json:"description"`
								Regions     map[string]Region `json:"regions"`
							} `json:"licensed7695108898142923768core"`
							Licensed7695108898142923768F1Micro struct {
								Description string            `json:"description"`
								Regions     map[string]Region `json:"regions"`
							} `json:"licensed7695108898142923768f1micro"`
							Licensed7695108898142923768G1Small struct {
								Description string            `json:"description"`
								Regions     map[string]Region `json:"regions"`
							} `json:"licensed7695108898142923768g1small"`
							Licensed7798417859637521376Core struct {
								Description string            `json:"description"`
								Regions     map[string]Region `json:"regions"`
							} `json:"licensed7798417859637521376core"`
							Licensed7798417859637521376F1Micro struct {
								Description string            `json:"description"`
								Regions     map[string]Region `json:"regions"`
							} `json:"licensed7798417859637521376f1micro"`
							Licensed7798417859637521376G1Small struct {
								Description string            `json:"description"`
								Regions     map[string]Region `json:"regions"`
							} `json:"licensed7798417859637521376g1small"`
						} `json:"windows"`
						SQLServer struct {
							Licensed1741222371620352982Core5Ormore struct {
								Description string            `json:"description"`
								Regions     map[string]Region `json:"regions"`
							} `json:"licensed1741222371620352982core5ormore"`
							Licensed3039072951948447844Core5Ormore struct {
								Description string            `json:"description"`
								Regions     map[string]Region `json:"regions"`
							} `json:"licensed3039072951948447844core5ormore"`
							Licensed3042936622923550835Core5Ormore struct {
								Description string            `json:"description"`
								Regions     map[string]Region `json:"regions"`
							} `json:"licensed3042936622923550835core5ormore"`
							Licensed3398668354433905558Core5Ormore struct {
								Description string            `json:"description"`
								Regions     map[string]Region `json:"regions"`
							} `json:"licensed3398668354433905558core5ormore"`
							Licensed6213885950785916969Core5Ormore struct {
								Description string            `json:"description"`
								Regions     map[string]Region `json:"regions"`
							} `json:"licensed6213885950785916969core5ormore"`
							Licensed6795597790302237536Core5Ormore struct {
								Description string            `json:"description"`
								Regions     map[string]Region `json:"regions"`
							} `json:"licensed6795597790302237536core5ormore"`
						} `json:"sql_server"`
					} `json:"microsoft"`
					Rhel struct {
						Licensed7883559014960410759Corerange04 struct {
							Description string            `json:"description"`
							Regions     map[string]Region `json:"regions"`
						} `json:"licensed7883559014960410759corerange04"`
						Licensed7883559014960410759Corerange5Ormore struct {
							Description string            `json:"description"`
							Regions     map[string]Region `json:"regions"`
						} `json:"licensed7883559014960410759corerange5ormore"`
					} `json:"rhel"`
				} `json:"premium_image"`
				Ingress struct {
					Premium struct {
						Networkgoogleingress struct {
							Description string            `json:"description"`
							Regions     map[string]Region `json:"regions"`
						} `json:"networkgoogleingress"`
						Networkinternetingress struct {
							Description string            `json:"description"`
							Regions     map[string]Region `json:"regions"`
						} `json:"networkinternetingress"`
					} `json:"premium"`
					InterRegion struct {
						Networkinterregioningress struct {
							Description string            `json:"description"`
							Regions     map[string]Region `json:"regions"`
						} `json:"networkinterregioningress"`
					} `json:"inter-region"`
					InterZone struct {
						Networkinterzoneingress struct {
							Description string            `json:"description"`
							Regions     map[string]Region `json:"regions"`
						} `json:"networkinterzoneingress"`
					} `json:"inter-zone"`
					Standard struct {
						Networkinternetstandardtieringress struct {
							Description string            `json:"description"`
							Regions     map[string]Region `json:"regions"`
						} `json:"networkinternetstandardtieringress"`
					} `json:"standard"`
					IntraZone struct {
						Networkintrazoneingress struct {
							Description string            `json:"description"`
							Regions     map[string]Region `json:"regions"`
						} `json:"networkintrazoneingress"`
					} `json:"intra-zone"`
				} `json:"ingress"`
				VmsReservation struct {
					CoresPerCore struct {
						Reservationa2Highgpucore struct {
							Description string            `json:"description"`
							Regions     map[string]Region `json:"regions"`
						} `json:"reservationa2highgpucore"`
					} `json:"cores:_per_core"`
					MemoryPerGb struct {
						Reservationa2Highgpuram struct {
							Description string            `json:"description"`
							Regions     map[string]Region `json:"regions"`
						} `json:"reservationa2highgpuram"`
					} `json:"memory:_per_gb"`
				} `json:"vms_reservation"`
				Storagemultiregionalarchivesnapshotdatastorage struct {
					Description string            `json:"description"`
					Regions     map[string]Region `json:"regions"`
				} `json:"storagemultiregionalarchivesnapshotdatastorage"`
				Storagemultiregionalarchivesnapshotearlydeletion struct {
					Description string            `json:"description"`
					Regions     map[string]Region `json:"regions"`
				} `json:"storagemultiregionalarchivesnapshotearlydeletion"`
				Storagemultiregionalarchivesnapshotretrieval struct {
					Description string            `json:"description"`
					Regions     map[string]Region `json:"regions"`
				} `json:"storagemultiregionalarchivesnapshotretrieval"`
				Storageregionalarchivesnapshotearlydeletion struct {
					Description string            `json:"description"`
					Regions     map[string]Region `json:"regions"`
				} `json:"storageregionalarchivesnapshotearlydeletion"`
				Storageregionalarchivesnapshotretrieval struct {
					Description string            `json:"description"`
					Regions     map[string]Region `json:"regions"`
				} `json:"storageregionalarchivesnapshotretrieval"`
				MemoryPerGb struct {
					Vmimagen2Soletenancyramsoletenancyovercommitpremium struct {
						Description string            `json:"description"`
						Regions     map[string]Region `json:"regions"`
					} `json:"vmimagen2soletenancyramsoletenancyovercommitpremium"`
					Vmimagen2Soletenancyramsoletenancypremium struct {
						Description string            `json:"description"`
						Regions     map[string]Region `json:"regions"`
					} `json:"vmimagen2soletenancyramsoletenancypremium"`
					Vmimagesoletenancyramsoletenancyovercommitpremium struct {
						Description string            `json:"description"`
						Regions     map[string]Region `json:"regions"`
					} `json:"vmimagesoletenancyramsoletenancyovercommitpremium"`
				} `json:"memory:_per_gb"`
				Management struct {
					Agentshourscount struct {
						Description string            `json:"description"`
						Regions     map[string]Region `json:"regions"`
					} `json:"agentshourscount"`
					Cloudopsagentshourscount struct {
						Description string            `json:"description"`
						Regions     map[string]Region `json:"regions"`
					} `json:"cloudopsagentshourscount"`
				} `json:"management"`
				WorkloadManager struct {
					BillingScannedResources struct {
						Description string            `json:"description"`
						Regions     map[string]Region `json:"regions"`
					} `json:"billing/scanned_resources"`
				} `json:"workload_manager"`
			} `json:"gce"`
			PersistentDisk struct {
				AsyncReplication struct {
					BalancedProtection struct {
						Asyncreplicationprotectionpdbalanced struct {
							Description string            `json:"description"`
							Regions     map[string]Region `json:"regions"`
						} `json:"asyncreplicationprotectionpdbalanced"`
						Regional struct {
							Asyncreplicationprotectionregionalpdbalanced struct {
								Description string            `json:"description"`
								Regions     map[string]Region `json:"regions"`
							} `json:"asyncreplicationprotectionregionalpdbalanced"`
						} `json:"regional"`
					} `json:"balanced_protection"`
					SsdProtection struct {
						Asyncreplicationprotectionpdssd struct {
							Description string            `json:"description"`
							Regions     map[string]Region `json:"regions"`
						} `json:"asyncreplicationprotectionpdssd"`
						Regional struct {
							Asyncreplicationprotectionregionalpdssd struct {
								Description string            `json:"description"`
								Regions     map[string]Region `json:"regions"`
							} `json:"asyncreplicationprotectionregionalpdssd"`
						} `json:"regional"`
					} `json:"ssd_protection"`
					Networking struct {
						Asia struct {
							Asynchronousreplicationinterregionnetworkegress struct {
								Description string            `json:"description"`
								Regions     map[string]Region `json:"regions"`
							} `json:"asynchronousreplicationinterregionnetworkegress"`
						} `json:"asia"`
						Europe struct {
							Asynchronousreplicationinterregionnetworkegress struct {
								Description string            `json:"description"`
								Regions     map[string]Region `json:"regions"`
							} `json:"asynchronousreplicationinterregionnetworkegress"`
						} `json:"europe"`
						NorthAmerica struct {
							Asynchronousreplicationinterregionnetworkegress struct {
								Description string            `json:"description"`
								Regions     map[string]Region `json:"regions"`
							} `json:"asynchronousreplicationinterregionnetworkegress"`
						} `json:"north_america"`
						Oceania struct {
							Asynchronousreplicationinterregionnetworkegress struct {
								Description string            `json:"description"`
								Regions     map[string]Region `json:"regions"`
							} `json:"asynchronousreplicationinterregionnetworkegress"`
						} `json:"oceania"`
					} `json:"networking"`
				} `json:"async_replication"`
				Snapshots struct {
					Multiregionalsnapshotdownload struct {
						Description string            `json:"description"`
						Regions     map[string]Region `json:"regions"`
					} `json:"multiregionalsnapshotdownload"`
					Multiregionalsnapshotupload struct {
						Description string            `json:"description"`
						Regions     map[string]Region `json:"regions"`
					} `json:"multiregionalsnapshotupload"`
					Storagemultiregionalstandardsnapshotearlydeletion struct {
						Description string            `json:"description"`
						Regions     map[string]Region `json:"regions"`
					} `json:"storagemultiregionalstandardsnapshotearlydeletion"`
					Storageregionalarchivesnapshotdatastorage struct {
						Description string            `json:"description"`
						Regions     map[string]Region `json:"regions"`
					} `json:"storageregionalarchivesnapshotdatastorage"`
					Storageregionalarchivesnapshotearlydeletion struct {
						Description string            `json:"description"`
						Regions     map[string]Region `json:"regions"`
					} `json:"storageregionalarchivesnapshotearlydeletion"`
					Storageregionalarchivesnapshotretrieval struct {
						Description string            `json:"description"`
						Regions     map[string]Region `json:"regions"`
					} `json:"storageregionalarchivesnapshotretrieval"`
					Storageregionalstandardsnapshotearlydeletion struct {
						Description string            `json:"description"`
						Regions     map[string]Region `json:"regions"`
					} `json:"storageregionalstandardsnapshotearlydeletion"`
				} `json:"snapshots"`
				Diskops struct {
					Pdiorequests struct {
						Description string            `json:"description"`
						Regions     map[string]Region `json:"regions"`
					} `json:"pdiorequests"`
				} `json:"diskops"`
				Standard struct {
					Capacity struct {
						Regional struct {
							Regionalstoragepdcapacity struct {
								Description string            `json:"description"`
								Regions     map[string]Region `json:"regions"`
							} `json:"regionalstoragepdcapacity"`
						} `json:"regional"`
						Storagepdcapacity struct {
							Description string            `json:"description"`
							Regions     map[string]Region `json:"regions"`
						} `json:"storagepdcapacity"`
					} `json:"capacity"`
					Snapshot struct {
						Storagepdsnapshot struct {
							Description string            `json:"description"`
							Regions     map[string]Region `json:"regions"`
						} `json:"storagepdsnapshot"`
					} `json:"snapshot"`
				} `json:"standard"`
				Ssd struct {
					Capacity struct {
						Regional struct {
							Regionalstoragepdssd struct {
								Description string            `json:"description"`
								Regions     map[string]Region `json:"regions"`
							} `json:"regionalstoragepdssd"`
						} `json:"regional"`
						Storagepdssd struct {
							Description string            `json:"description"`
							Regions     map[string]Region `json:"regions"`
						} `json:"storagepdssd"`
						Extreme struct {
							Storagepdssdextremecapacity struct {
								Description string            `json:"description"`
								Regions     map[string]Region `json:"regions"`
							} `json:"storagepdssdextremecapacity"`
							Storagepdssdextremeiops struct {
								Description string            `json:"description"`
								Regions     map[string]Region `json:"regions"`
							} `json:"storagepdssdextremeiops"`
						} `json:"extreme"`
						Lite struct {
							Storagepdssdlitecapacity struct {
								Description string            `json:"description"`
								Regions     map[string]Region `json:"regions"`
							} `json:"storagepdssdlitecapacity"`
						} `json:"lite"`
						RegionalLite struct {
							Storageregionalpdssdlitecapacity struct {
								Description string            `json:"description"`
								Regions     map[string]Region `json:"regions"`
							} `json:"storageregionalpdssdlitecapacity"`
						} `json:"regional_lite"`
					} `json:"capacity"`
				} `json:"ssd"`
			} `json:"persistent_disk"`
			Gpus struct {
				GpusCommit1Year struct {
					L4 struct {
						Commitmentgpunvidial41Yv1 struct {
							Description string            `json:"description"`
							Regions     map[string]Region `json:"regions"`
						} `json:"commitmentgpunvidial41yv1"`
					} `json:"l4"`
					A100 struct {
						Commitmentgpunvidiateslaa1001Yv1 struct {
							Description string            `json:"description"`
							Regions     map[string]Region `json:"regions"`
						} `json:"commitmentgpunvidiateslaa1001yv1"`
					} `json:"a100"`
					K80 struct {
						Commitmentgpunvidiateslak801Yv1 struct {
							Description string            `json:"description"`
							Regions     map[string]Region `json:"regions"`
						} `json:"commitmentgpunvidiateslak801yv1"`
					} `json:"k80"`
					P100 struct {
						Commitmentgpunvidiateslap1001Yv1 struct {
							Description string            `json:"description"`
							Regions     map[string]Region `json:"regions"`
						} `json:"commitmentgpunvidiateslap1001yv1"`
					} `json:"p100"`
					P4 struct {
						Commitmentgpunvidiateslap41Yv1 struct {
							Description string            `json:"description"`
							Regions     map[string]Region `json:"regions"`
						} `json:"commitmentgpunvidiateslap41yv1"`
					} `json:"p4"`
					T4 struct {
						Commitmentgpunvidiateslat41Yv1 struct {
							Description string            `json:"description"`
							Regions     map[string]Region `json:"regions"`
						} `json:"commitmentgpunvidiateslat41yv1"`
					} `json:"t4"`
					V100 struct {
						Commitmentgpunvidiateslav1001Yv1 struct {
							Description string            `json:"description"`
							Regions     map[string]Region `json:"regions"`
						} `json:"commitmentgpunvidiateslav1001yv1"`
					} `json:"v100"`
				} `json:"gpus_commit_1_year"`
				GpusCommit3Year struct {
					L4 struct {
						Commitmentgpunvidial43Yv1 struct {
							Description string            `json:"description"`
							Regions     map[string]Region `json:"regions"`
						} `json:"commitmentgpunvidial43yv1"`
					} `json:"l4"`
					A100 struct {
						Commitmentgpunvidiateslaa1003Yv1 struct {
							Description string            `json:"description"`
							Regions     map[string]Region `json:"regions"`
						} `json:"commitmentgpunvidiateslaa1003yv1"`
					} `json:"a100"`
					K80 struct {
						Commitmentgpunvidiateslak803Yv1 struct {
							Description string            `json:"description"`
							Regions     map[string]Region `json:"regions"`
						} `json:"commitmentgpunvidiateslak803yv1"`
					} `json:"k80"`
					P100 struct {
						Commitmentgpunvidiateslap1003Yv1 struct {
							Description string            `json:"description"`
							Regions     map[string]Region `json:"regions"`
						} `json:"commitmentgpunvidiateslap1003yv1"`
					} `json:"p100"`
					P4 struct {
						Commitmentgpunvidiateslap43Yv1 struct {
							Description string            `json:"description"`
							Regions     map[string]Region `json:"regions"`
						} `json:"commitmentgpunvidiateslap43yv1"`
					} `json:"p4"`
					T4 struct {
						Commitmentgpunvidiateslat43Yv1 struct {
							Description string            `json:"description"`
							Regions     map[string]Region `json:"regions"`
						} `json:"commitmentgpunvidiateslat43yv1"`
					} `json:"t4"`
					V100 struct {
						Commitmentgpunvidiateslav1003Yv1 struct {
							Description string            `json:"description"`
							Regions     map[string]Region `json:"regions"`
						} `json:"commitmentgpunvidiateslav1003yv1"`
					} `json:"v100"`
				} `json:"gpus_commit_3_year"`
				GpusOnDemand struct {
					L4 struct {
						Gpunvidial4 struct {
							Description string            `json:"description"`
							Regions     map[string]Region `json:"regions"`
						} `json:"gpunvidial4"`
					} `json:"l4"`
					A100 struct {
						Gpunvidiateslaa100 struct {
							Description string            `json:"description"`
							Regions     map[string]Region `json:"regions"`
						} `json:"gpunvidiateslaa100"`
					} `json:"a100"`
					K80 struct {
						Gpunvidiateslak80 struct {
							Description string            `json:"description"`
							Regions     map[string]Region `json:"regions"`
						} `json:"gpunvidiateslak80"`
					} `json:"k80"`
					P100 struct {
						Gpunvidiateslap100 struct {
							Description string            `json:"description"`
							Regions     map[string]Region `json:"regions"`
						} `json:"gpunvidiateslap100"`
					} `json:"p100"`
					P4 struct {
						Gpunvidiateslap4 struct {
							Description string            `json:"description"`
							Regions     map[string]Region `json:"regions"`
						} `json:"gpunvidiateslap4"`
					} `json:"p4"`
					T4 struct {
						Gpunvidiateslat4 struct {
							Description string            `json:"description"`
							Regions     map[string]Region `json:"regions"`
						} `json:"gpunvidiateslat4"`
					} `json:"t4"`
					V100 struct {
						Gpunvidiateslav100 struct {
							Description string            `json:"description"`
							Regions     map[string]Region `json:"regions"`
						} `json:"gpunvidiateslav100"`
					} `json:"v100"`
				} `json:"gpus_on_demand"`
				GpusPreemptible struct {
					L4 struct {
						Gpupreemptiblenvidial4 struct {
							Description string            `json:"description"`
							Regions     map[string]Region `json:"regions"`
						} `json:"gpupreemptiblenvidial4"`
					} `json:"l4"`
					A100 struct {
						Gpupreemptiblenvidiateslaa100 struct {
							Description string            `json:"description"`
							Regions     map[string]Region `json:"regions"`
						} `json:"gpupreemptiblenvidiateslaa100"`
					} `json:"a100"`
					K80 struct {
						Gpupreemptiblenvidiateslak80 struct {
							Description string            `json:"description"`
							Regions     map[string]Region `json:"regions"`
						} `json:"gpupreemptiblenvidiateslak80"`
					} `json:"k80"`
					P100 struct {
						Gpupreemptiblenvidiateslap100 struct {
							Description string            `json:"description"`
							Regions     map[string]Region `json:"regions"`
						} `json:"gpupreemptiblenvidiateslap100"`
					} `json:"p100"`
					P4 struct {
						Gpupreemptiblenvidiateslap4 struct {
							Description string            `json:"description"`
							Regions     map[string]Region `json:"regions"`
						} `json:"gpupreemptiblenvidiateslap4"`
					} `json:"p4"`
					T4 struct {
						Gpupreemptiblenvidiateslat4 struct {
							Description string            `json:"description"`
							Regions     map[string]Region `json:"regions"`
						} `json:"gpupreemptiblenvidiateslat4"`
					} `json:"t4"`
					V100 struct {
						Gpupreemptiblenvidiateslav100 struct {
							Description string            `json:"description"`
							Regions     map[string]Region `json:"regions"`
						} `json:"gpupreemptiblenvidiateslav100"`
					} `json:"v100"`
				} `json:"gpus_preemptible"`
				GpusReservation struct {
					A100 struct {
						Reservationgpunvidiateslaa100 struct {
							Description string            `json:"description"`
							Regions     map[string]Region `json:"regions"`
						} `json:"reservationgpunvidiateslaa100"`
					} `json:"a100"`
				} `json:"gpus_reservation"`
			} `json:"gpus"`
			LocalSsd struct {
				Commit1Year struct {
					Commitmentlocalssd1Yv1 struct {
						Description string            `json:"description"`
						Regions     map[string]Region `json:"regions"`
					} `json:"commitmentlocalssd1yv1"`
				} `json:"commit_1_year"`
				Commit3Year struct {
					Commitmentlocalssd3Yv1 struct {
						Description string            `json:"description"`
						Regions     map[string]Region `json:"regions"`
					} `json:"commitmentlocalssd3yv1"`
				} `json:"commit_3_year"`
				OnDemand struct {
					Storagelocalssd struct {
						Description string            `json:"description"`
						Regions     map[string]Region `json:"regions"`
					} `json:"storagelocalssd"`
					Storagepreemptiblelocalssd struct {
						Description string            `json:"description"`
						Regions     map[string]Region `json:"regions"`
					} `json:"storagepreemptiblelocalssd"`
				} `json:"on_demand"`
			} `json:"local_ssd"`
			Hyperdisk struct {
				HyperdiskVolumes struct {
					Extreme struct {
						Capacity struct {
							Storagehyperdiskextremecapacity struct {
								Description string            `json:"description"`
								Regions     map[string]Region `json:"regions"`
							} `json:"storagehyperdiskextremecapacity"`
						} `json:"capacity"`
						Iops struct {
							Storagehyperdiskextremeiops struct {
								Description string            `json:"description"`
								Regions     map[string]Region `json:"regions"`
							} `json:"storagehyperdiskextremeiops"`
						} `json:"iops"`
					} `json:"extreme"`
				} `json:"hyperdisk_volumes"`
			} `json:"hyperdisk"`
			VMImageStorage struct {
				ImageStorage struct {
					Storageimage struct {
						Description string            `json:"description"`
						Regions     map[string]Region `json:"regions"`
					} `json:"storageimage"`
					Storagemachineimage struct {
						Description string            `json:"description"`
						Regions     map[string]Region `json:"regions"`
					} `json:"storagemachineimage"`
				} `json:"image_storage"`
			} `json:"vm_image_storage"`
			CloudTpu struct {
				Compute struct {
					Tpu struct {
						Tpuv2Preemptiblesecondsdefault struct {
							Description string            `json:"description"`
							Regions     map[string]Region `json:"regions"`
						} `json:"tpuv2preemptiblesecondsdefault"`
						Tpuv2Preemptiblesecondspoddefault struct {
							Description string            `json:"description"`
							Regions     map[string]Region `json:"regions"`
						} `json:"tpuv2preemptiblesecondspoddefault"`
						Tpuv2Secondsdefault struct {
							Description string            `json:"description"`
							Regions     map[string]Region `json:"regions"`
						} `json:"tpuv2secondsdefault"`
						Tpuv2Secondspoddefault struct {
							Description string            `json:"description"`
							Regions     map[string]Region `json:"regions"`
						} `json:"tpuv2secondspoddefault"`
						Tpuv3Preemptiblesecondsdefault struct {
							Description string            `json:"description"`
							Regions     map[string]Region `json:"regions"`
						} `json:"tpuv3preemptiblesecondsdefault"`
						Tpuv3Preemptiblesecondspoddefault struct {
							Description string            `json:"description"`
							Regions     map[string]Region `json:"regions"`
						} `json:"tpuv3preemptiblesecondspoddefault"`
						Tpuv3Secondsdefault struct {
							Description string            `json:"description"`
							Regions     map[string]Region `json:"regions"`
						} `json:"tpuv3secondsdefault"`
						Tpuv3Secondspoddefault struct {
							Description string            `json:"description"`
							Regions     map[string]Region `json:"regions"`
						} `json:"tpuv3secondspoddefault"`
						Tpuv4Preemptiblesecondspoddefault struct {
							Description string            `json:"description"`
							Regions     map[string]Region `json:"regions"`
						} `json:"tpuv4preemptiblesecondspoddefault"`
						Tpuv4Secondspoddefault struct {
							Description string            `json:"description"`
							Regions     map[string]Region `json:"regions"`
						} `json:"tpuv4secondspoddefault"`
					} `json:"tpu"`
				} `json:"compute"`
			} `json:"cloud_tpu"`
			ComputeSolutions struct {
				GoogleCloudVmwareEngine struct {
					Standard72Hourlyv2 struct {
						Description string            `json:"description"`
						Regions     map[string]Region `json:"regions"`
					} `json:"standard72hourlyv2"`
					Standard72Hourly struct {
						Description string            `json:"description"`
						Regions     map[string]Region `json:"regions"`
					} `json:"standard72hourly"`
					Vmwareengineucs12Mo struct {
						Description string            `json:"description"`
						Regions     map[string]Region `json:"regions"`
					} `json:"vmwareengineucs12mo"`
					Vmwareengineucs36Mo struct {
						Description string            `json:"description"`
						Regions     map[string]Region `json:"regions"`
					} `json:"vmwareengineucs36mo"`
				} `json:"google_cloud_vmware_engine"`
			} `json:"compute_solutions"`
		} `json:"compute"`
	} `json:"gcp"`
}
