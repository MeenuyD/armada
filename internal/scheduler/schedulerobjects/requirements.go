package schedulerobjects

func (info *JobSchedulingInfo) GetTotalResourceRequest() ResourceList {
	rv := ResourceList{}
	for _, oreq := range info.ObjectRequirements {
		if preq := oreq.GetPodRequirements(); preq != nil {
			rv.Add(ResourceListFromV1ResourceList(preq.ResourceRequirements.Requests))
		}
	}
	return rv
}