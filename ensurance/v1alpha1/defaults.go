package v1alpha1

func SetDefaults_NodeQOSEnsurancePolicy(nep *NodeQOSEnsurancePolicy) {
	if nep.Spec.NodeQualityProbe.NodeLocalGet.LocalCacheTTLSeconds == nil {
		var i int32 = 60
		nep.Spec.NodeQualityProbe.NodeLocalGet.LocalCacheTTLSeconds = &i
	}
}

func SetDefaults_ObjectiveEnsurance(obj *ObjectiveEnsurance) {
	if obj.AvoidanceThreshold == 0 {
		obj.AvoidanceThreshold = 2
	}
}
