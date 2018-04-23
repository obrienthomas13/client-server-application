package tcp

type TCPHeader struct {
  SourcePort uint16
  DestinationPort uint16
  SequenceNumber uint32
  AcknowledgementNumber uint32
  DataOffSet uint8
  Reserved uint8
  ECN uint8
  Control uint8
  WindowSize uint16
  TCPCheckSum uint16
  UrgentPointer uint16
  Options []TCPOptions
}

type TCPOptions struct {
  Kind uint8
  Length uint8
  Data []byte
}
