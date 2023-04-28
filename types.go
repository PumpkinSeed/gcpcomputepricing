package gcpcomputepricing

type Pricing struct {
	Gcp struct {
		Compute struct {
			Gce struct {
				VmsOnDemand struct {
					AdvancedNetworking struct {
						Advnet100Gbpstotalondemand Subtype `json:"advnet100gbpstotalondemand"`
						Advnet50Gbpstotalondemand  Subtype `json:"advnet50gbpstotalondemand"`
						Advnet75Gbpstotalondemand  Subtype `json:"advnet75gbpstotalondemand"`
					} `json:"advanced_networking"`
					MemoryPerGb struct {
						Vmstatesuspendedram         Subtype `json:"vmstatesuspendedram"`
						Vmimagea2Highgpuram         Subtype `json:"vmimagea2highgpuram"`
						Vmimagec2Dcustomextendedram Subtype `json:"vmimagec2dcustomextendedram"`
						Vmimagec2Dcustomram         Subtype `json:"vmimagec2dcustomram"`
						Vmimagec2Dstandardram       Subtype `json:"vmimagec2dstandardram"`
						C3                          struct {
							Vmimagec3Soletenancyram                   Subtype `json:"vmimagec3soletenancyram"`
							Vmimagec3Soletenancyramsoletenancypremium Subtype `json:"vmimagec3soletenancyramsoletenancypremium"`
							Vmimagec3Standardram                      Subtype `json:"vmimagec3standardram"`
						} `json:"c3"`
						Vmimagecomputeoptimizedram Subtype `json:"vmimagecomputeoptimizedram"`
						Vmimagecustomextendedram   Subtype `json:"vmimagecustomextendedram"`
						Vmimagecustomram           Subtype `json:"vmimagecustomram"`
						Vmimagee2RAM               Subtype `json:"vmimagee2ram"`
						G2                         struct {
							Vmimageg2Soletenancyram                   Subtype `json:"vmimageg2soletenancyram"`
							Vmimageg2Soletenancyramsoletenancypremium Subtype `json:"vmimageg2soletenancyramsoletenancypremium"`
							Vmimageg2Standardram                      Subtype `json:"vmimageg2standardram"`
						} `json:"g2"`
						Vmimagelargeram                              Subtype `json:"vmimagelargeram"`
						Vmimagelargerammemoryoptimizedupgradepremium Subtype `json:"vmimagelargerammemoryoptimizedupgradepremium"`
						M3                                           struct {
							Vmimagem3Soletenancyram                   Subtype `json:"vmimagem3soletenancyram"`
							Vmimagem3Soletenancyramsoletenancypremium Subtype `json:"vmimagem3soletenancyramsoletenancypremium"`
							Vmimagem3Standardram                      Subtype `json:"vmimagem3standardram"`
						} `json:"m3"`
						Vmimagen1Standardram                       Subtype `json:"vmimagen1standardram"`
						Vmimagen2Customextendedram                 Subtype `json:"vmimagen2customextendedram"`
						Vmimagen2Customram                         Subtype `json:"vmimagen2customram"`
						Vmimagen2Dcustomextendedram                Subtype `json:"vmimagen2dcustomextendedram"`
						Vmimagen2Dcustomram                        Subtype `json:"vmimagen2dcustomram"`
						Vmimagen2Dsoletenancyram                   Subtype `json:"vmimagen2dsoletenancyram"`
						Vmimagen2Dsoletenancyramsoletenancypremium Subtype `json:"vmimagen2dsoletenancyramsoletenancypremium"`
						Vmimagen2Dstandardram                      Subtype `json:"vmimagen2dstandardram"`
						Vmimagen2Soletenancyram                    Subtype `json:"vmimagen2soletenancyram"`
						Vmimagen2Standardram                       Subtype `json:"vmimagen2standardram"`
						Vmimagesoletenancyram                      Subtype `json:"vmimagesoletenancyram"`
						Vmimagesoletenancyramsoletenancypremium    Subtype `json:"vmimagesoletenancyramsoletenancypremium"`
						T2A                                        struct {
							Vmimaget2Astandardram Subtype `json:"vmimaget2astandardram"`
						} `json:"t2a"`
						Vmimaget2Dstandardram Subtype `json:"vmimaget2dstandardram"`
					} `json:"memory:_per_gb"`
					CoresPerCore struct {
						Vmimagea2Highgpucore   Subtype `json:"vmimagea2highgpucore"`
						Vmimagec2Dcustomcore   Subtype `json:"vmimagec2dcustomcore"`
						Vmimagec2Dstandardcore Subtype `json:"vmimagec2dstandardcore"`
						C3                     struct {
							Vmimagec3Soletenancycore                   Subtype `json:"vmimagec3soletenancycore"`
							Vmimagec3Soletenancycoresoletenancypremium Subtype `json:"vmimagec3soletenancycoresoletenancypremium"`
							Vmimagec3Standardcore                      Subtype `json:"vmimagec3standardcore"`
						} `json:"c3"`
						Vmimagecomputeoptimizedcore Subtype `json:"vmimagecomputeoptimizedcore"`
						Vmimagecustomcore           Subtype `json:"vmimagecustomcore"`
						Vmimagee2Core               Subtype `json:"vmimagee2core"`
						Vmimagef1Micro              Subtype `json:"vmimagef1micro"`
						Vmimageg1Small              Subtype `json:"vmimageg1small"`
						G2                          struct {
							Vmimageg2Soletenancycore                   Subtype `json:"vmimageg2soletenancycore"`
							Vmimageg2Soletenancycoresoletenancypremium Subtype `json:"vmimageg2soletenancycoresoletenancypremium"`
							Vmimageg2Standardcore                      Subtype `json:"vmimageg2standardcore"`
						} `json:"g2"`
						Vmimagelargecore                              Subtype `json:"vmimagelargecore"`
						Vmimagelargecorememoryoptimizedupgradepremium Subtype `json:"vmimagelargecorememoryoptimizedupgradepremium"`
						M3                                            struct {
							Vmimagem3Soletenancycore                   Subtype `json:"vmimagem3soletenancycore"`
							Vmimagem3Soletenancycoresoletenancypremium Subtype `json:"vmimagem3soletenancycoresoletenancypremium"`
							Vmimagem3Standardcore                      Subtype `json:"vmimagem3standardcore"`
						} `json:"m3"`
						Vmimagen1Standardcore                       Subtype `json:"vmimagen1standardcore"`
						Vmimagen2Customcore                         Subtype `json:"vmimagen2customcore"`
						Vmimagen2Customextendedcore                 Subtype `json:"vmimagen2customextendedcore"`
						Vmimagen2Dcustomcore                        Subtype `json:"vmimagen2dcustomcore"`
						Vmimagen2Dsoletenancycore                   Subtype `json:"vmimagen2dsoletenancycore"`
						Vmimagen2Dsoletenancycoresoletenancypremium Subtype `json:"vmimagen2dsoletenancycoresoletenancypremium"`
						Vmimagen2Dstandardcore                      Subtype `json:"vmimagen2dstandardcore"`
						Vmimagen2Soletenancycore                    Subtype `json:"vmimagen2soletenancycore"`
						Vmimagen2Soletenancycoresoletenancypremium  Subtype `json:"vmimagen2soletenancycoresoletenancypremium"`
						Vmimagen2Standardcore                       Subtype `json:"vmimagen2standardcore"`
						Vmimagesoletenancycore                      Subtype `json:"vmimagesoletenancycore"`
						Vmimagesoletenancycoresoletenancypremium    Subtype `json:"vmimagesoletenancycoresoletenancypremium"`
						T2A                                         struct {
							Vmimaget2Astandardcore Subtype `json:"vmimaget2astandardcore"`
						} `json:"t2a"`
						Vmimaget2Dstandardcore Subtype `json:"vmimaget2dstandardcore"`
					} `json:"cores:_per_core"`
					Vmimagecomputeoptimizedsoletenancycore                   Subtype `json:"vmimagecomputeoptimizedsoletenancycore"`
					Vmimagecomputeoptimizedsoletenancycoresoletenancypremium Subtype `json:"vmimagecomputeoptimizedsoletenancycoresoletenancypremium"`
					Vmimagecomputeoptimizedsoletenancyram                    Subtype `json:"vmimagecomputeoptimizedsoletenancyram"`
					Vmimagecomputeoptimizedsoletenancyramsoletenancypremium  Subtype `json:"vmimagecomputeoptimizedsoletenancyramsoletenancypremium"`
					Vmimagelargesoletenancycore                              Subtype `json:"vmimagelargesoletenancycore"`
					Vmimagelargesoletenancycoresoletenancypremium            Subtype `json:"vmimagelargesoletenancycoresoletenancypremium"`
					Vmimagelargesoletenancyram                               Subtype `json:"vmimagelargesoletenancyram"`
					Vmimagelargesoletenancyramsoletenancypremium             Subtype `json:"vmimagelargesoletenancyramsoletenancypremium"`
					Vmimagen2Soletenancycoresoletenancyovercommitpremium     Subtype `json:"vmimagen2soletenancycoresoletenancyovercommitpremium"`
					Vmimagen2Soletenancyramsoletenancyovercommitpremium      Subtype `json:"vmimagen2soletenancyramsoletenancyovercommitpremium"`
					Vmimagen2Soletenancyramsoletenancypremium                Subtype `json:"vmimagen2soletenancyramsoletenancypremium"`
					Vmimagesoletenancycoresoletenancyovercommitpremium       Subtype `json:"vmimagesoletenancycoresoletenancyovercommitpremium"`
					Vmimagesoletenancyramsoletenancyovercommitpremium        Subtype `json:"vmimagesoletenancyramsoletenancyovercommitpremium"`
				} `json:"vms_on_demand"`
				VmsPreemptible struct {
					AdvancedNetworking struct {
						Advnet100Gbpstotalpreemptible Subtype `json:"advnet100gbpstotalpreemptible"`
						Advnet50Gbpstotalpreemptible  Subtype `json:"advnet50gbpstotalpreemptible"`
						Advnet75Gbpstotalpreemptible  Subtype `json:"advnet75gbpstotalpreemptible"`
					} `json:"advanced_networking"`
					Cores1To64 struct {
						Highcpu struct {
							Vmimagepreemptiblea2Highgpucore Subtype `json:"vmimagepreemptiblea2highgpucore"`
						} `json:"highcpu"`
					} `json:"cores:_1_to_64"`
					MemoryPerGb struct {
						Vmimagepreemptiblea2Highgpuram         Subtype `json:"vmimagepreemptiblea2highgpuram"`
						Vmimagepreemptiblec2Dcustomextendedram Subtype `json:"vmimagepreemptiblec2dcustomextendedram"`
						Vmimagepreemptiblec2Dcustomram         Subtype `json:"vmimagepreemptiblec2dcustomram"`
						Vmimagepreemptiblec2Dstandardram       Subtype `json:"vmimagepreemptiblec2dstandardram"`
						C3                                     struct {
							Vmimagepreemptiblec3Standardram Subtype `json:"vmimagepreemptiblec3standardram"`
						} `json:"c3"`
						Vmimagepreemptiblecomputeoptimizedram Subtype `json:"vmimagepreemptiblecomputeoptimizedram"`
						Vmimagepreemptiblecustomextendedram   Subtype `json:"vmimagepreemptiblecustomextendedram"`
						Vmimagepreemptiblecustomram           Subtype `json:"vmimagepreemptiblecustomram"`
						Vmimagepreemptiblee2RAM               Subtype `json:"vmimagepreemptiblee2ram"`
						G2                                    struct {
							Vmimagepreemptibleg2Standardram Subtype `json:"vmimagepreemptibleg2standardram"`
						} `json:"g2"`
						Vmimagepreemptiblelargeram Subtype `json:"vmimagepreemptiblelargeram"`
						M3                         struct {
							Vmimagepreemptiblem3Standardram Subtype `json:"vmimagepreemptiblem3standardram"`
						} `json:"m3"`
						Vmimagepreemptiblen1Standardram        Subtype `json:"vmimagepreemptiblen1standardram"`
						Vmimagepreemptiblen2Customextendedram  Subtype `json:"vmimagepreemptiblen2customextendedram"`
						Vmimagepreemptiblen2Customram          Subtype `json:"vmimagepreemptiblen2customram"`
						Vmimagepreemptiblen2Dcustomextendedram Subtype `json:"vmimagepreemptiblen2dcustomextendedram"`
						Vmimagepreemptiblen2Dcustomram         Subtype `json:"vmimagepreemptiblen2dcustomram"`
						Vmimagepreemptiblen2Dstandardram       Subtype `json:"vmimagepreemptiblen2dstandardram"`
						Vmimagepreemptiblen2Standardram        Subtype `json:"vmimagepreemptiblen2standardram"`
						T2A                                    struct {
							Vmimagepreemptiblet2Astandardram Subtype `json:"vmimagepreemptiblet2astandardram"`
						} `json:"t2a"`
						Vmimagepreemptiblet2Dstandardram Subtype `json:"vmimagepreemptiblet2dstandardram"`
					} `json:"memory:_per_gb"`
					CoresPerCore struct {
						Vmimagepreemptiblec2Dcustomcore   Subtype `json:"vmimagepreemptiblec2dcustomcore"`
						Vmimagepreemptiblec2Dstandardcore Subtype `json:"vmimagepreemptiblec2dstandardcore"`
						C3                                struct {
							Vmimagepreemptiblec3Standardcore Subtype `json:"vmimagepreemptiblec3standardcore"`
						} `json:"c3"`
						Vmimagepreemptiblecomputeoptimizedcore Subtype `json:"vmimagepreemptiblecomputeoptimizedcore"`
						Vmimagepreemptiblecustomcore           Subtype `json:"vmimagepreemptiblecustomcore"`
						Vmimagepreemptiblecustomextendedcore   Subtype `json:"vmimagepreemptiblecustomextendedcore"`
						Vmimagepreemptiblee2Core               Subtype `json:"vmimagepreemptiblee2core"`
						Vmimagepreemptiblef1Micro              Subtype `json:"vmimagepreemptiblef1micro"`
						Vmimagepreemptibleg1Small              Subtype `json:"vmimagepreemptibleg1small"`
						G2                                     struct {
							Vmimagepreemptibleg2Standardcore Subtype `json:"vmimagepreemptibleg2standardcore"`
						} `json:"g2"`
						Vmimagepreemptiblelargecore Subtype `json:"vmimagepreemptiblelargecore"`
						M3                          struct {
							Vmimagepreemptiblem3Standardcore Subtype `json:"vmimagepreemptiblem3standardcore"`
						} `json:"m3"`
						Vmimagepreemptiblen1Standardcore       Subtype `json:"vmimagepreemptiblen1standardcore"`
						Vmimagepreemptiblen2Customcore         Subtype `json:"vmimagepreemptiblen2customcore"`
						Vmimagepreemptiblen2Customextendedcore Subtype `json:"vmimagepreemptiblen2customextendedcore"`
						Vmimagepreemptiblen2Dcustomcore        Subtype `json:"vmimagepreemptiblen2dcustomcore"`
						Vmimagepreemptiblen2Dstandardcore      Subtype `json:"vmimagepreemptiblen2dstandardcore"`
						Vmimagepreemptiblen2Standardcore       Subtype `json:"vmimagepreemptiblen2standardcore"`
						T2A                                    struct {
							Vmimagepreemptiblet2Astandardcore Subtype `json:"vmimagepreemptiblet2astandardcore"`
						} `json:"t2a"`
						Vmimagepreemptiblet2Dstandardcore Subtype `json:"vmimagepreemptiblet2dstandardcore"`
					} `json:"cores:_per_core"`
				} `json:"vms_preemptible"`
				VmsCommit1Year struct {
					CoresPerCore struct {
						Commitmenta2Highgpucpu1Yv1 Subtype `json:"commitmenta2highgpucpu1yv1"`
						Commitmentc2Dcpu1Yv1       Subtype `json:"commitmentc2dcpu1yv1"`
						C3                         struct {
							Commitmentc3CPU1Yv1 Subtype `json:"commitmentc3cpu1yv1"`
						} `json:"c3"`
						Commitmentcpucomputeoptimized1Yv1 Subtype `json:"commitmentcpucomputeoptimized1yv1"`
						Commitmentcpulargeinstance1Yv1    Subtype `json:"commitmentcpulargeinstance1yv1"`
						Commitmentcpu1Yv1                 Subtype `json:"commitmentcpu1yv1"`
						Commitmente2CPU1Yv1               Subtype `json:"commitmente2cpu1yv1"`
						G2                                struct {
							Commitmentg2CPU1Yv1 Subtype `json:"commitmentg2cpu1yv1"`
						} `json:"g2"`
						M3 struct {
							Commitmentm3CPU1Yv1 Subtype `json:"commitmentm3cpu1yv1"`
						} `json:"m3"`
						Commitmentn2CPU1Yv1  Subtype `json:"commitmentn2cpu1yv1"`
						Commitmentn2Dcpu1Yv1 Subtype `json:"commitmentn2dcpu1yv1"`
						Commitmentt2Dcpu1Yv1 Subtype `json:"commitmentt2dcpu1yv1"`
					} `json:"cores:_per_core"`
					MemoryPerGb struct {
						Commitmenta2Highgpuram1Yv1 Subtype `json:"commitmenta2highgpuram1yv1"`
						Commitmentc2Dram1Yv1       Subtype `json:"commitmentc2dram1yv1"`
						C3                         struct {
							Commitmentc3RAM1Yv1 Subtype `json:"commitmentc3ram1yv1"`
						} `json:"c3"`
						Commitmente2RAM1Yv1 Subtype `json:"commitmente2ram1yv1"`
						G2                  struct {
							Commitmentg2RAM1Yv1 Subtype `json:"commitmentg2ram1yv1"`
						} `json:"g2"`
						M3 struct {
							Commitmentm3RAM1Yv1 Subtype `json:"commitmentm3ram1yv1"`
						} `json:"m3"`
						Commitmentn2Dram1Yv1              Subtype `json:"commitmentn2dram1yv1"`
						Commitmentn2RAM1Yv1               Subtype `json:"commitmentn2ram1yv1"`
						Commitmentramcomputeoptimized1Yv1 Subtype `json:"commitmentramcomputeoptimized1yv1"`
						Commitmentramlargeinstance1Yv1    Subtype `json:"commitmentramlargeinstance1yv1"`
						Commitmentram1Yv1                 Subtype `json:"commitmentram1yv1"`
						Commitmentt2Dram1Yv1              Subtype `json:"commitmentt2dram1yv1"`
					} `json:"memory:_per_gb"`
					Vmwareengineucs12Moprepay Subtype `json:"vmwareengineucs12moprepay"`
				} `json:"vms_commit_1_year"`
				VmsCommit3Year struct {
					CoresPerCore struct {
						Commitmenta2Highgpucpu3Yv1 Subtype `json:"commitmenta2highgpucpu3yv1"`
						Commitmentc2Dcpu3Yv1       Subtype `json:"commitmentc2dcpu3yv1"`
						C3                         struct {
							Commitmentc3CPU3Yv1 Subtype `json:"commitmentc3cpu3yv1"`
						} `json:"c3"`
						Commitmentcpucomputeoptimized3Yv1 Subtype `json:"commitmentcpucomputeoptimized3yv1"`
						Commitmentcpulargeinstance1Yv1    Subtype `json:"commitmentcpulargeinstance1yv1"`
						Commitmentcpulargeinstance3Yv1    Subtype `json:"commitmentcpulargeinstance3yv1"`
						Commitmentcpu3Yv1                 Subtype `json:"commitmentcpu3yv1"`
						Commitmente2CPU3Yv1               Subtype `json:"commitmente2cpu3yv1"`
						G2                                struct {
							Commitmentg2CPU3Yv1 Subtype `json:"commitmentg2cpu3yv1"`
						} `json:"g2"`
						M3 struct {
							Commitmentm3CPU3Yv1 Subtype `json:"commitmentm3cpu3yv1"`
						} `json:"m3"`
						Commitmentn2CPU3Yv1  Subtype `json:"commitmentn2cpu3yv1"`
						Commitmentn2Dcpu3Yv1 Subtype `json:"commitmentn2dcpu3yv1"`
						Commitmentt2Dcpu3Yv1 Subtype `json:"commitmentt2dcpu3yv1"`
					} `json:"cores:_per_core"`
					MemoryPerGb struct {
						Commitmenta2Highgpuram3Yv1 Subtype `json:"commitmenta2highgpuram3yv1"`
						Commitmentc2Dram3Yv1       Subtype `json:"commitmentc2dram3yv1"`
						C3                         struct {
							Commitmentc3RAM3Yv1 Subtype `json:"commitmentc3ram3yv1"`
						} `json:"c3"`
						Commitmente2RAM3Yv1 Subtype `json:"commitmente2ram3yv1"`
						G2                  struct {
							Commitmentg2RAM3Yv1 Subtype `json:"commitmentg2ram3yv1"`
						} `json:"g2"`
						M3 struct {
							Commitmentm3RAM3Yv1 Subtype `json:"commitmentm3ram3yv1"`
						} `json:"m3"`
						Commitmentn2Dram3Yv1              Subtype `json:"commitmentn2dram3yv1"`
						Commitmentn2RAM3Yv1               Subtype `json:"commitmentn2ram3yv1"`
						Commitmentramcomputeoptimized3Yv1 Subtype `json:"commitmentramcomputeoptimized3yv1"`
						Commitmentramlargeinstance3Yv1    Subtype `json:"commitmentramlargeinstance3yv1"`
						Commitmentram3Yv1                 Subtype `json:"commitmentram3yv1"`
						Commitmentt2Dram3Yv1              Subtype `json:"commitmentt2dram3yv1"`
					} `json:"memory:_per_gb"`
					Vmwareengineucs36Moprepay Subtype `json:"vmwareengineucs36moprepay"`
				} `json:"vms_commit_3_year"`
				NetworkOther struct {
					Externalip            Subtype `json:"externalip"`
					Externalippreemptible Subtype `json:"externalippreemptible"`
				} `json:"network_other"`
				FlexCommit1Year struct {
					Gcecommitmentsucs12Mo Subtype `json:"gcecommitmentsucs12mo"`
				} `json:"flex_commit_1_year"`
				FlexCommit3Year struct {
					Gcecommitmentsucs36Mo Subtype `json:"gcecommitmentsucs36mo"`
				} `json:"flex_commit_3_year"`
				PremiumImage struct {
					Microsoft struct {
						Windows struct {
							Licensed1656378918552316916Core    Subtype `json:"licensed1656378918552316916core"`
							Licensed1656378918552316916F1Micro Subtype `json:"licensed1656378918552316916f1micro"`
							Licensed1656378918552316916G1Small Subtype `json:"licensed1656378918552316916g1small"`
							Licensed3284763237085719542Core    Subtype `json:"licensed3284763237085719542core"`
							Licensed3284763237085719542F1Micro Subtype `json:"licensed3284763237085719542f1micro"`
							Licensed3284763237085719542G1Small Subtype `json:"licensed3284763237085719542g1small"`
							Licensed4819555115818134498Core    Subtype `json:"licensed4819555115818134498core"`
							Licensed4819555115818134498F1Micro Subtype `json:"licensed4819555115818134498f1micro"`
							Licensed4819555115818134498G1Small Subtype `json:"licensed4819555115818134498g1small"`
							Licensed4874454843789519845Core    Subtype `json:"licensed4874454843789519845core"`
							Licensed4874454843789519845F1Micro Subtype `json:"licensed4874454843789519845f1micro"`
							Licensed4874454843789519845G1Small Subtype `json:"licensed4874454843789519845g1small"`
							Licensed6107784707477449232Core    Subtype `json:"licensed6107784707477449232core"`
							Licensed6107784707477449232F1Micro Subtype `json:"licensed6107784707477449232f1micro"`
							Licensed6107784707477449232G1Small Subtype `json:"licensed6107784707477449232g1small"`
							Licensed7695108898142923768Core    Subtype `json:"licensed7695108898142923768core"`
							Licensed7695108898142923768F1Micro Subtype `json:"licensed7695108898142923768f1micro"`
							Licensed7695108898142923768G1Small Subtype `json:"licensed7695108898142923768g1small"`
							Licensed7798417859637521376Core    Subtype `json:"licensed7798417859637521376core"`
							Licensed7798417859637521376F1Micro Subtype `json:"licensed7798417859637521376f1micro"`
							Licensed7798417859637521376G1Small Subtype `json:"licensed7798417859637521376g1small"`
						} `json:"windows"`
						SQLServer struct {
							Licensed1741222371620352982Core5Ormore Subtype `json:"licensed1741222371620352982core5ormore"`
							Licensed3039072951948447844Core5Ormore Subtype `json:"licensed3039072951948447844core5ormore"`
							Licensed3042936622923550835Core5Ormore Subtype `json:"licensed3042936622923550835core5ormore"`
							Licensed3398668354433905558Core5Ormore Subtype `json:"licensed3398668354433905558core5ormore"`
							Licensed6213885950785916969Core5Ormore Subtype `json:"licensed6213885950785916969core5ormore"`
							Licensed6795597790302237536Core5Ormore Subtype `json:"licensed6795597790302237536core5ormore"`
						} `json:"sql_server"`
					} `json:"microsoft"`
					Rhel struct {
						Licensed7883559014960410759Corerange04      Subtype `json:"licensed7883559014960410759corerange04"`
						Licensed7883559014960410759Corerange5Ormore Subtype `json:"licensed7883559014960410759corerange5ormore"`
					} `json:"rhel"`
				} `json:"premium_image"`
				Ingress struct {
					Premium struct {
						Networkgoogleingress   Subtype `json:"networkgoogleingress"`
						Networkinternetingress Subtype `json:"networkinternetingress"`
					} `json:"premium"`
					InterRegion struct {
						Networkinterregioningress Subtype `json:"networkinterregioningress"`
					} `json:"inter-region"`
					InterZone struct {
						Networkinterzoneingress Subtype `json:"networkinterzoneingress"`
					} `json:"inter-zone"`
					Standard struct {
						Networkinternetstandardtieringress Subtype `json:"networkinternetstandardtieringress"`
					} `json:"standard"`
					IntraZone struct {
						Networkintrazoneingress Subtype `json:"networkintrazoneingress"`
					} `json:"intra-zone"`
				} `json:"ingress"`
				VmsReservation struct {
					CoresPerCore struct {
						Reservationa2Highgpucore Subtype `json:"reservationa2highgpucore"`
					} `json:"cores:_per_core"`
					MemoryPerGb struct {
						Reservationa2Highgpuram Subtype `json:"reservationa2highgpuram"`
					} `json:"memory:_per_gb"`
				} `json:"vms_reservation"`
				Storagemultiregionalarchivesnapshotdatastorage   Subtype `json:"storagemultiregionalarchivesnapshotdatastorage"`
				Storagemultiregionalarchivesnapshotearlydeletion Subtype `json:"storagemultiregionalarchivesnapshotearlydeletion"`
				Storagemultiregionalarchivesnapshotretrieval     Subtype `json:"storagemultiregionalarchivesnapshotretrieval"`
				Storageregionalarchivesnapshotearlydeletion      Subtype `json:"storageregionalarchivesnapshotearlydeletion"`
				Storageregionalarchivesnapshotretrieval          Subtype `json:"storageregionalarchivesnapshotretrieval"`
				MemoryPerGb                                      struct {
					Vmimagen2Soletenancyramsoletenancyovercommitpremium Subtype `json:"vmimagen2soletenancyramsoletenancyovercommitpremium"`
					Vmimagen2Soletenancyramsoletenancypremium           Subtype `json:"vmimagen2soletenancyramsoletenancypremium"`
					Vmimagesoletenancyramsoletenancyovercommitpremium   Subtype `json:"vmimagesoletenancyramsoletenancyovercommitpremium"`
				} `json:"memory:_per_gb"`
				Management struct {
					Agentshourscount         Subtype `json:"agentshourscount"`
					Cloudopsagentshourscount Subtype `json:"cloudopsagentshourscount"`
				} `json:"management"`
				WorkloadManager struct {
					BillingScannedResources Subtype `json:"billing/scanned_resources"`
				} `json:"workload_manager"`
			} `json:"gce"`
			PersistentDisk struct {
				AsyncReplication struct {
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
				} `json:"async_replication"`
				Snapshots struct {
					Multiregionalsnapshotdownload                     Subtype `json:"multiregionalsnapshotdownload"`
					Multiregionalsnapshotupload                       Subtype `json:"multiregionalsnapshotupload"`
					Storagemultiregionalstandardsnapshotearlydeletion Subtype `json:"storagemultiregionalstandardsnapshotearlydeletion"`
					Storageregionalarchivesnapshotdatastorage         Subtype `json:"storageregionalarchivesnapshotdatastorage"`
					Storageregionalarchivesnapshotearlydeletion       Subtype `json:"storageregionalarchivesnapshotearlydeletion"`
					Storageregionalarchivesnapshotretrieval           Subtype `json:"storageregionalarchivesnapshotretrieval"`
					Storageregionalstandardsnapshotearlydeletion      Subtype `json:"storageregionalstandardsnapshotearlydeletion"`
				} `json:"snapshots"`
				Diskops struct {
					Pdiorequests Subtype `json:"pdiorequests"`
				} `json:"diskops"`
				Standard struct {
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
			} `json:"persistent_disk"`
			Gpus struct {
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
			} `json:"gpus"`
			LocalSsd struct {
				Commit1Year struct {
					Commitmentlocalssd1Yv1 Subtype `json:"commitmentlocalssd1yv1"`
				} `json:"commit_1_year"`
				Commit3Year struct {
					Commitmentlocalssd3Yv1 Subtype `json:"commitmentlocalssd3yv1"`
				} `json:"commit_3_year"`
				OnDemand struct {
					Storagelocalssd            Subtype `json:"storagelocalssd"`
					Storagepreemptiblelocalssd Subtype `json:"storagepreemptiblelocalssd"`
				} `json:"on_demand"`
			} `json:"local_ssd"`
			Hyperdisk struct {
				HyperdiskVolumes struct {
					Extreme struct {
						Capacity struct {
							Storagehyperdiskextremecapacity Subtype `json:"storagehyperdiskextremecapacity"`
						} `json:"capacity"`
						Iops struct {
							Storagehyperdiskextremeiops Subtype `json:"storagehyperdiskextremeiops"`
						} `json:"iops"`
					} `json:"extreme"`
				} `json:"hyperdisk_volumes"`
			} `json:"hyperdisk"`
			VMImageStorage struct {
				ImageStorage struct {
					Storageimage        Subtype `json:"storageimage"`
					Storagemachineimage Subtype `json:"storagemachineimage"`
				} `json:"image_storage"`
			} `json:"vm_image_storage"`
			CloudTpu struct {
				Compute struct {
					Tpu struct {
						Tpuv2Preemptiblesecondsdefault    Subtype `json:"tpuv2preemptiblesecondsdefault"`
						Tpuv2Preemptiblesecondspoddefault Subtype `json:"tpuv2preemptiblesecondspoddefault"`
						Tpuv2Secondsdefault               Subtype `json:"tpuv2secondsdefault"`
						Tpuv2Secondspoddefault            Subtype `json:"tpuv2secondspoddefault"`
						Tpuv3Preemptiblesecondsdefault    Subtype `json:"tpuv3preemptiblesecondsdefault"`
						Tpuv3Preemptiblesecondspoddefault Subtype `json:"tpuv3preemptiblesecondspoddefault"`
						Tpuv3Secondsdefault               Subtype `json:"tpuv3secondsdefault"`
						Tpuv3Secondspoddefault            Subtype `json:"tpuv3secondspoddefault"`
						Tpuv4Preemptiblesecondspoddefault Subtype `json:"tpuv4preemptiblesecondspoddefault"`
						Tpuv4Secondspoddefault            Subtype `json:"tpuv4secondspoddefault"`
					} `json:"tpu"`
				} `json:"compute"`
			} `json:"cloud_tpu"`
			ComputeSolutions struct {
				GoogleCloudVmwareEngine struct {
					Standard72Hourlyv2  Subtype `json:"standard72hourlyv2"`
					Standard72Hourly    Subtype `json:"standard72hourly"`
					Vmwareengineucs12Mo Subtype `json:"vmwareengineucs12mo"`
					Vmwareengineucs36Mo Subtype `json:"vmwareengineucs36mo"`
				} `json:"google_cloud_vmware_engine"`
			} `json:"compute_solutions"`
		} `json:"compute"`
	} `json:"gcp"`
}

type Subtype struct {
	Description string            `json:"description"`
	Regions     map[string]Region `json:"regions"`
}

type Region struct {
	Prices       []Price `json:"price"`
	Name         string  `json:"name"`
	InternalName string  `json:"internal_name"`
}

type Price struct {
	Val      uint64 `json:"val"`
	Currency string `json:"currency"`
	Nanos    uint64 `json:"nanos"`
}
