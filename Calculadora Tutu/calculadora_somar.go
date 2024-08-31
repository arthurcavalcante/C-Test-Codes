package main
	
type SomarOperacao struct {
	
	Operation

}

func (SomarOperacao) CollectOperation(float32, float32, error) (float32, error) {

	return Operando_01 + Operando_02, nil

}
