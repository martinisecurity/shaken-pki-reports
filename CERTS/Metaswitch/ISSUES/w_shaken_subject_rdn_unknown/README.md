# STIR/SHAKEN CA Ecosystem Compliance

## Metaswitch

Name: w_shaken_subject_rdn_unknown\
Source: SHAKEN_PKI_BEST_PRACTICES\
Citation: SHAKEN PKI Best Practice\
Effective Date: 01 Jan 00 00:00 UTC\
Description: Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject. However, the names should be meaningful meaningful enough to represent the SP to whom the certificate is being issued, in a manner similar to that used to identify SP’s equipment in the network.

### Leaf Certificates

| Status | Subject | Link | Details |
|--------|---------|------|---------|
| warn | Verizon SHAKEN cert 5807 | [view](../../CERTS/d7b413267be2d050d516af4f4a864ffdc2feacc774a1a6264b9cfe68c796f43f/README.md) | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |

### CA Certificates

No error, warning, or notice level issues were found


Generated: 17 Dec 22 17:07 UTC