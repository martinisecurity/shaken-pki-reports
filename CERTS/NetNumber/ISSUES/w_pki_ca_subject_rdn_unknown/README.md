# STIR/SHAKEN CA Ecosystem Compliance

## NetNumber

Name: w_pki_ca_subject_rdn_unknown\
Source: SHAKEN PKI Best Practice\
Citation: SHAKEN PKI Best Practice\
Effective Date: 01 Jan 00 00:00 UTC\
Description: Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject. However, the names should be meaningful meaningful enough to represent the SP to whom the certificate is being issued, in a manner similar to that used to identify SP’s equipment in the network.

### Leaf Certificates

no warning, or error, or not effective date level issues were found

### CA Certificates

| Status | Subject | Link | Details |
|--------|---------|------|---------|
| warn | NetNumber SHAKEN Root Intermediate CA 1 | [view](../../CERTS/e449803766edf02ab50b034dd7e89e54efd332cce87688a032f89b340d039878/README.md) | Only CN, C, L, and O should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| warn | NetNumber SHAKEN Root CA | [view](../../CERTS/7ac80e8481ecb019dc95484016842db78686069efbc0f703e7f39310217b6157/README.md) | Only CN, C, L, and O should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| warn | NetNumber SHAKEN Root CA 1 | [view](../../CERTS/2dd1386ca717f31d550b35b9bce9daa9b02483bcdb98bdfcfca07202276136d7/README.md) | Only CN, C, L, and O should be included. Additional RNDs may introduce ambiguity and may not be verifiable |


Generated: 01/11/2022 at 10:05:32