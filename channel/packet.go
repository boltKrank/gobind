package channel

/* All DNS packets have a structure that is

+---------------------+
| Header              |
+---------------------+
| Question            | Question for the name server
+---------------------+
| Answer              | Answers to the question
+---------------------+
| Authority           | Not used in this project
+---------------------+
| Additional          | Not used in this project
+---------------------+ */

type Packet struct {
	Header     Header
	Question   []Question
	Answer     []Answer
	Authority  []Authority
	Additional []Additional
}

/*

HEADER STRUCTURE
                                1  1  1  1  1  1
  0  1  2  3  4  5  6  7  8  9  0  1  2  3  4  5
+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
|                     ID                        |
+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
|QR|   Opcode |AA|TC|RD|RA|     Z      |  RCODE |
+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
|                 QDCOUNT                       |
+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
|                 ANCOUNT                       |
+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
|                 NSCOUNT                       |
+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
|                 ARCOUNT                       |
+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
*/

// This is how we push the bits into a single 16 bit integer
/*
func (h Header) IsResponse() bool {
    return h.Flags&FlagQR != 0
}

func (h *Header) SetResponse() {
    h.Flags |= FlagQR
}
*/

const (
	FlagQR uint16 = 1 << 15
	FlagAA uint16 = 1 << 10
	FlagTC uint16 = 1 << 9
	FlagRD uint16 = 1 << 8
	FlagRA uint16 = 1 << 7
)

type Header struct {
	// ID - 16 bit identifier
	Id uint16

	// QR - 1 bit
	// OPCODE - 4 bit
	// AA - Auth Answer - 1 bit
	// TC - TrunCation - 1 bit
	// RD - Recursion Desired - 1 bit
	// RA - Recursion Available - 1 bit
	// Z  - Reserved for future use - 1 bit (usually set to 0)
	// RCODE - Response code - 4 bit

	// We'll combine these all into a single 16 bit u-int
	// The bit modifications are done as const above
	Flags uint16

	// QDCOUNT - Number of entries in the questions section - 16 bit
	QDCount uint16

	// ANCOUNT - Number of record resources in the answer section - 16 bit
	ANCount uint16

	// NSCOUNT - Number of name server resource recors in the auth section - 16 bit
	NSCount uint16

	// ARCOUNT - Number of resource records in the addtional records section - 16 bit
	ARCount uint16
}

/*
QUESTION STRUCTURE:
                               1  1  1  1  1  1
 0  1  2  3  4   5 6   7 8  9  0  1  2  3  4  5
+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
|                                               |
/                    QNAME                      /
/                                               /
+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
|                    QTYPE                      |
+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
|                   QCLASS                      |
+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+

*/

type Question struct {
	// QNAME - The domain name, sequence of labels, variable length, zero terminated
	QName []byte

	// QTYPE - Query type - 2 x octets, 16 bit
	QType uint16

	// QCLASS - Query class - 2 x octets, 16 bit
	QClass uint16
}

/*

ANSWER STRUCTURE:
                               1  1  1  1  1  1
 0  1  2  3  4  5  6  7  8  9  0  1  2  3  4  5
+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
|                                               |
/                                               /
/                    NAME                       /
|                                               |
+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
|                    TYPE                       |
+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
|                    CLASS                      |
+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
|                    TTL                        |
|                                               |
+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
|                    RDLENGTH                   |
+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--|
/                    RDATA                      /
/                                               /
+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
*/

type Answer struct {
	// NAME - Domain name queried - same format as QNAME
	Name []byte

	// TYPE - The type of data in the RDATA field - 2 x octets
	Type uint16

	// CLASS - The class of data in the RDATA field - 2 x octets
	Class uint16

	// TTL - The seconds a result can be cached - 16 bit
	Ttl uint16

	// RDLENGTH - Length of the RDATA field - 16 bit
	Rdlength uint16

	// RDATA - The data being returned, format and length are dependent on TYPE
	RData []byte
}

type Authority struct {
}

type Additional struct {
}
