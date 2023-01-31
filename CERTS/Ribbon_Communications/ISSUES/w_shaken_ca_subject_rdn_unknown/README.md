# STIR/SHAKEN CA Ecosystem Compliance

## Ribbon Communications

Name: w_shaken_ca_subject_rdn_unknown\
Source: SHAKEN_PKI_BEST_PRACTICES\
Citation: SHAKEN PKI Best Practice\
Effective Date: 01 Jan 00 00:00 UTC\
Description: Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject. However, the names should be meaningful meaningful enough to represent the SP to whom the certificate is being issued, in a manner similar to that used to identify SP’s equipment in the network.

### Leaf Certificates

No error, warning, or notice level issues were found

### CA Certificates

| Status | Subject | Link | Details |
|--------|---------|------|---------|
| warn | SHAKEN Ribbon Issuing CA | [view](../../CERTS/05a71a04eaedbdf4b0534f40768616d7c19c8deb5a3aefd1f4a04b3aab55a48f/README.md) | Only CN, C, L, and O should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| warn | SHAKEN Ribbon Root CA | [view](../../CERTS/7c2799d3642d04f04fe667c3ab251c18689af323acdc43b2fa5f3dc89e3a0f14/README.md) | Only CN, C, L, and O should be included. Additional RNDs may introduce ambiguity and may not be verifiable |


Generated: 31 Jan 23 21:50 UTC