# STIR/SHAKEN CA Ecosystem Compliance

## Peeringhub

Name: w_shaken_subject_rdn_unknown\
Source: SHAKEN_PKI_BEST_PRACTICES\
Citation: SHAKEN PKI Best Practice\
Effective Date: 01 Jan 00 00:00 UTC\
Description: Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject. However, the names should be meaningful meaningful enough to represent the SP to whom the certificate is being issued, in a manner similar to that used to identify SP’s equipment in the network.

### Leaf Certificates

| Status | Subject | Link | Details |
|--------|---------|------|---------|
| warn | VOCALTRANSIT SHAKEN 783J | [view](../../CERTS/e8c92f75c60632193af34b988570281cbbcbafce7a9c3ae534b8406d62aab179/README.md) | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |

### CA Certificates

No error, warning, or notice level issues were found


Generated: 15 Nov 23 18:10 UTC