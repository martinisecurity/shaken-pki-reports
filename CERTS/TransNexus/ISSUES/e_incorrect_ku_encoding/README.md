# STIR/SHAKEN CA Ecosystem Compliance

## TransNexus

Name: e_incorrect_ku_encoding\
Source: RFC5280\
Citation: Where ITU-T Rec. X.680 | ISO/IEC 8824-1, 21.7, applies, the bitstring shall have all trailing 0 bits removed before it is encoded.\
Effective Date: 01 Jan 00 00:00 UTC\
Description: RFC 5280 Section 4.2.1.3 describes the value of a KeyUsage to be a DER encoded BitString, which itself defines that all trailing 0 bits be counted as being "unused".

### Leaf Certificates

| Status | Subject | Link | Details |
|--------|---------|------|---------|
| error | SHAKEN 012K | [view](../../CERTS/788e828c8fe34b9e0116e8f5d3a7e6d71a22a83803056a63b5ab35f2ec52f6cf/README.md) | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |

### CA Certificates

| Status | Subject | Link | Details |
|--------|---------|------|---------|
| error | TransNexus, Inc. SHAKEN Issuing CA3 | [view](../../CERTS/51b51431ffd8055cd7e30c213540f4868258058d4c4c9f23f1ed8bea75101e37/README.md) | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 1. Raw Bytes: [3 2 0 6], Raw Binary: [00000011 00000010 00000000 00000110] |


Generated: 31/10/2022 at 16:43:22