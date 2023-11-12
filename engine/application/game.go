package application

type Loader interface {
	Load()domain.State
}
type game struct {
   loader Loader,
}
func NewGame (loader Loader){
	return game{
		loader: loader,
	}
}
func (g game) StartGame(){
   _ := g.loader.Load()
}