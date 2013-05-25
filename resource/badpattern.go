package main

type Resource struct {
}

func (*Resource) activate() {}

func (*Resource) destroy() {}

// START OMIT

type ResourceManager struct {
	//Buffered channel with capacity == max available resources
	availableResources chan bool
}

func (rm *ResourceManager) Initialize(numResources int) {
	rm.availableResources = make(chan bool, numResources)
}

func (rm *ResourceManager) GetResource() *Resource {
	rm.availableResources <- true
	rr := new(Resource)
	rr.activate()
	return rr
}

func (rm *ResourceManager) ReturnResource(rr *Resource) {
	rr.destroy()
	<-rm.availableResources
}

// END OMIT

func main() {
}
