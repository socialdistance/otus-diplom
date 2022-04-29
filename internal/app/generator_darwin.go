package app

var generators []generator

func init() {
	generators = []generator{
		&loadavgGenerator{},
		//&cpuGenerator{},
		//&diskGenerator{},
		//&memoryGenerator{},
		//&networkGenerator{},
		//&talkersGenerator{},
	}
}
