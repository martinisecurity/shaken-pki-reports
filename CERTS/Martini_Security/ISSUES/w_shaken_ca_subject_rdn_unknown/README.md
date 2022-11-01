# STIR/SHAKEN CA Ecosystem Compliance

## Martini Security

Name: w_shaken_ca_subject_rdn_unknown\
Source: SHAKEN PKI Best Practice\
Citation: SHAKEN PKI Best Practice\
Effective Date: 01 Jan 00 00:00 UTC\
Description: Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject. However, the names should be meaningful meaningful enough to represent the SP to whom the certificate is being issued, in a manner similar to that used to identify SP’s equipment in the network.

### Leaf Certificates

no warning, or error, or not effective date level issues were found

### CA Certificates

| Status | Subject | Link | Details |
|--------|---------|------|---------|
| warn | Martini Security SHAKEN G2 | [view](../../CERTS/bf818ddbd3ae492e4a85331b85b52f4d2cdef8287bf910b59e247b6c132fa7fd/README.md) | Only CN, C, L, and O should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| warn | Martini Security SHAKEN G1 | [view](../../CERTS/46e98b1599688c83928b66bcdd49e723ee4207128b5ba4488046edf2879970ef/README.md) | Only CN, C, L, and O should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| warn | Martini Security SHAKEN R1 | [view](../../CERTS/6077e368b0a0e4b6076eaa07ce67d6652ef310c4757776b76af84a6a9e003cdf/README.md) | Only CN, C, L, and O should be included. Additional RNDs may introduce ambiguity and may not be verifiable |


Generated: 01/11/2022 at 16:30:07