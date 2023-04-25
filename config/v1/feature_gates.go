package v1

// FeatureGateDescription is a golang-only interface used to contains details for a feature gate.
type FeatureGateDescription struct {
	// FeatureGateAttributes is the information that appears in the API
	FeatureGateAttributes FeatureGateAttributes

	// OwningJiraComponent is the jira component that owns most of the impl and first assignment for the bug.
	// This is the team that owns the feature long term.
	OwningJiraComponent string
	// ResponsiblePerson is the person who is on the hook for first contact.  This is often, but not always, a team lead.
	// It is someone who can make the promise on the behalf of the team.
	ResponsiblePerson string
	// OwningProduct is the product that owns the lifecycle of the gate.
	OwningProduct OwningProduct
}

type OwningProduct string

var (
	ocpSpecific = OwningProduct("OCP")
	kubernetes  = OwningProduct("Kubernetes")
)

var (
	FeatureGateOpenShiftPodSecurityAdmission = FeatureGateName("OpenShiftPodSecurityAdmission")
	openShiftPodSecurityAdmission            = FeatureGateDescription{
		FeatureGateAttributes: FeatureGateAttributes{
			Name: FeatureGateOpenShiftPodSecurityAdmission,
		},
		OwningJiraComponent: "auth",
		ResponsiblePerson:   "stlaz",
		OwningProduct:       ocpSpecific,
	}

	FeatureGateRetroactiveDefaultStorageClass = FeatureGateName("RetroactiveDefaultStorageClass")
	retroactiveDefaultStorageClass            = FeatureGateDescription{
		FeatureGateAttributes: FeatureGateAttributes{
			Name: FeatureGateRetroactiveDefaultStorageClass,
		},
		OwningJiraComponent: "storage",
		ResponsiblePerson:   "RomanBednar",
		OwningProduct:       kubernetes,
	}

	FeatureGateExternalCloudProvider = FeatureGateName("ExternalCloudProvider")
	externalCloudProvider            = FeatureGateDescription{
		FeatureGateAttributes: FeatureGateAttributes{
			Name: FeatureGateExternalCloudProvider,
		},
		OwningJiraComponent: "cloud-provider",
		ResponsiblePerson:   "jspeed",
		OwningProduct:       ocpSpecific,
	}

	FeatureGateExternalCloudProviderAzure = FeatureGateName("ExternalCloudProviderAzure")
	externalCloudProviderAzure            = FeatureGateDescription{
		FeatureGateAttributes: FeatureGateAttributes{
			Name: FeatureGateExternalCloudProviderAzure,
		},
		OwningJiraComponent: "cloud-provider",
		ResponsiblePerson:   "jspeed",
		OwningProduct:       ocpSpecific,
	}

	FeatureGateExternalCloudProviderGCP = FeatureGateName("ExternalCloudProviderGCP")
	externalCloudProviderGCP            = FeatureGateDescription{
		FeatureGateAttributes: FeatureGateAttributes{
			Name: FeatureGateExternalCloudProviderGCP,
		},
		OwningJiraComponent: "cloud-provider",
		ResponsiblePerson:   "jspeed",
		OwningProduct:       ocpSpecific,
	}

	FeatureGateCSIDriverSharedResource = FeatureGateName("CSIDriverSharedResource")
	csiDriverSharedResource            = FeatureGateDescription{
		FeatureGateAttributes: FeatureGateAttributes{
			Name: FeatureGateCSIDriverSharedResource,
		},
		OwningJiraComponent: "builds",
		ResponsiblePerson:   "adkaplan",
		OwningProduct:       ocpSpecific,
	}

	FeatureGateBuildCSIVolumes = FeatureGateName("BuildCSIVolumes")
	buildCSIVolumes            = FeatureGateDescription{
		FeatureGateAttributes: FeatureGateAttributes{
			Name: FeatureGateBuildCSIVolumes,
		},
		OwningJiraComponent: "builds",
		ResponsiblePerson:   "adkaplan",
		OwningProduct:       ocpSpecific,
	}

	FeatureGateNodeSwap = FeatureGateName("NodeSwap")
	nodeSwap            = FeatureGateDescription{
		FeatureGateAttributes: FeatureGateAttributes{
			Name: FeatureGateNodeSwap,
		},
		OwningJiraComponent: "node",
		ResponsiblePerson:   "ehashman",
		OwningProduct:       kubernetes,
	}

	FeatureGateMachineAPIProviderOpenStack = FeatureGateName("MachineAPIProviderOpenStack")
	machineAPIProviderOpenStack            = FeatureGateDescription{
		FeatureGateAttributes: FeatureGateAttributes{
			Name: FeatureGateMachineAPIProviderOpenStack,
		},
		OwningJiraComponent: "openstack",
		ResponsiblePerson:   "egarcia",
		OwningProduct:       ocpSpecific,
	}

	FeatureGateInsightsConfigAPI = FeatureGateName("InsightsConfigAPI")
	insightsConfigAPI            = FeatureGateDescription{
		FeatureGateAttributes: FeatureGateAttributes{
			Name: FeatureGateInsightsConfigAPI,
		},
		OwningJiraComponent: "insights",
		ResponsiblePerson:   "tremes",
		OwningProduct:       ocpSpecific,
	}

	FeatureGateMatchLabelKeysInPodTopologySpread = FeatureGateName("MatchLabelKeysInPodTopologySpread")
	matchLabelKeysInPodTopologySpread            = FeatureGateDescription{
		FeatureGateAttributes: FeatureGateAttributes{
			Name: FeatureGateMatchLabelKeysInPodTopologySpread,
		},
		OwningJiraComponent: "scheduling",
		ResponsiblePerson:   "ingvagabund",
		OwningProduct:       kubernetes,
	}

	FeatureGatePDBUnhealthyPodEvictionPolicy = FeatureGateName("PDBUnhealthyPodEvictionPolicy")
	pdbUnhealthyPodEvictionPolicy            = FeatureGateDescription{
		FeatureGateAttributes: FeatureGateAttributes{
			Name: FeatureGatePDBUnhealthyPodEvictionPolicy,
		},
		OwningJiraComponent: "apps",
		ResponsiblePerson:   "atiratree",
		OwningProduct:       kubernetes,
	}

	FeatureGateDynamicResourceAllocation = FeatureGateName("DynamicResourceAllocation")
	dynamicResourceAllocation            = FeatureGateDescription{
		FeatureGateAttributes: FeatureGateAttributes{
			Name: FeatureGateDynamicResourceAllocation,
		},
		OwningJiraComponent: "scheduling",
		ResponsiblePerson:   "jchaloup",
		OwningProduct:       kubernetes,
	}

	FeatureGateAdmissionWebhookMatchConditions = FeatureGateName("AdmissionWebhookMatchConditions")
	admissionWebhookMatchConditions            = FeatureGateDescription{
		FeatureGateAttributes: FeatureGateAttributes{
			Name: FeatureGateDynamicResourceAllocation,
		},
		OwningJiraComponent: "kube-apiserver",
		ResponsiblePerson:   "benluddy",
		OwningProduct:       kubernetes,
	}

	FeatureGateAzureWorkloadIdentity = FeatureGateName("AzureWorkloadIdentity")
	azureWorkloadIdentity            = FeatureGateDescription{
		FeatureGateAttributes: FeatureGateAttributes{
			Name: FeatureGateDynamicResourceAllocation,
		},
		OwningJiraComponent: "cloud-credential-operator",
		ResponsiblePerson:   "abutcher",
		OwningProduct:       ocpSpecific,
	}
)
