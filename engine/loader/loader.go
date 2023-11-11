package loader

type DefaultLoader {
}

func (d DefaultLoader) Load() domain.State {
	layout := 'm'
	/*
	 XXXXXXX
	 X     X
	 X  XXXX
	XX  X
	 X  X
	 X  X
	 XXXX
	*/
	bp := blueprintFromString(layout)
	
	return domain.State{
		
	}
}

func blueprintFromString(layout string) domain.Blueprint{
	return initBlueprint(7,7)
}

func initBlueprint(x,y int){
	bp := make(domain.Blueprint, x - 1)
	for i := range(x-1){
        bp[i] = make([]domain.Cell, y- 1)
	}

	return bp
}