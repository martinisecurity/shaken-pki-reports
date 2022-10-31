# STIR/SHAKEN CA Ecosystem Compliance

## Peeringhub

Name: w_pki_subject_rdn_unknown\
Source: SHAKEN PKI Best Practice\
Citation: SHAKEN PKI Best Practice\
Effective Date: 01 Jan 00 00:00 UTC\
Description: Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject. However, the names should be meaningful meaningful enough to represent the SP to whom the certificate is being issued, in a manner similar to that used to identify SP’s equipment in the network.

### Leaf Certificates

| Status | Subject | Link | Details |
|--------|---------|------|---------|
| warn | Voiceterm SHAKEN 240K | [view](../../CERTS/84819934eb8b7f347d7365133dba4376d162d283fa9d09858aa1706ec88487c6/README.md) | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |

### CA Certificates

no warning, or error, or not effective date level issues were found


Generated: 31/10/2022 at 16:43:22