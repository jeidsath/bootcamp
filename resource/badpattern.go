package main

type Resource struct {
}

func (*ResourceManager) getResourceFromPool() *Resource {
	return new(Resource)
}

func (*ResourceManager) returnResourceToPool(rr *Resource) {
}

// START OMIT

type ResourceManager struct {
	//Buffered channel with capacity == max available resources
	availableResources chan bool
}

func (rm *ResourceManager) Initialize() {
	//Does nothing
}

func (rm *ResourceManager) GetResource() *Resource {
	rm.availableResources <- true
	return rm.getResourceFromPool()
}

func (rm *ResourceManager) ReturnResource(rr *Resource) {
	rm.returnResourceToPool(rr)
	<-rm.availableResources
}

// END OMIT

func main() {
}
