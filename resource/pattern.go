package main

type Resource struct {
}

func (*Resource) activate() {}

func (*Resource) deactivate() {}

// START OMIT

type ResourceManager struct {
	//Buffered channel with capacity == max available resources
	availableResources chan *Resource
}

func (rm *ResourceManager) Initialize(numResources int) {
	rm.availableResources = make(chan *Resource, numResources)
	for ii := 0; ii < cap(rm.availableResources); ii++ {
		rm.availableResources <- new(Resource)
	}
}

func (rm *ResourceManager) GetResource() *Resource {
	rr := <-rm.availableResources
	rr.activate()
	return rr
}

func (rm *ResourceManager) ReturnResource(rr *Resource) {
	rr.deactivate()
	rm.availableResources <- rr
}

// END OMIT

func main() {
}
