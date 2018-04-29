package handshake

import (
  tcp "../tcp"
)

func InitilizeConnection(packet tcp.TCPHeader) (tcp.TCPHeader, bool) {
  if (packet.SequenceNumber == 0x0001 && packet.AcknowledgementNumber == 0x0000) {
    packet.SequenceNumber = 0x0001
    packet.AcknowledgementNumber = 0x0001
    return packet, true
  }
  return packet, false
}

func ConfirmInitConnection(packet tcp.TCPHeader) (tcp.TCPHeader, bool) {
  if (packet.SequenceNumber == 0x0001 && packet.AcknowledgementNumber == 0x0001) {
    packet.SequenceNumber = 0x0000
    packet.AcknowledgementNumber = 0x0001
    return packet, true
  }
  return packet, false
}

func ConfirmPacket(packet tcp.TCPHeader) (tcp.TCPHeader, bool) {
  if (packet.SequenceNumber == 0x0000 && packet.AcknowledgementNumber == 0x0001) {
    return packet, true
  }
  return packet, false
}
