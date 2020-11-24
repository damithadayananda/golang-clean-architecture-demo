package producer

type TestProducer struct {

}

func (t TestProducer)Produce(interface{})bool{
	return true
}