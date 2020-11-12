package communication

type CommunicationInterface interface {
	SendRequest()
	SendResponseResponse()
}

const (
	unaryServerAddress  = ":7000"
	streamServerAddress = ":7001"
)

func GetCommunicationClient(communicationType string) CommunicationInterface {
	switch communicationType {
	case "unary":
		return CreateUnaryClient(unaryServerAddress)
	case "stream":
		return CreateStreamClient(streamServerAddress)
	}
	return nil
}

func StartServer(communicationType string) CommunicationInterface {
	switch communicationType {
	case "unary":
		StartUnaryServerInstance(unaryServerAddress)
		break
	case "stream":
		StartStreamServerInstance(streamServerAddress)
		break
	}
	return nil
}
