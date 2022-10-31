# STIR/SHAKEN CA Ecosystem Compliance

## Peeringhub

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
| warn | Peeringhub Inc SHAKEN Intermediate CA 2 | [view](../../CERTS/f00871963a40b04269c4b019968e42f9f40964cbfb512ff5342307e9942874ce/README.md) | Only CN, C, L, and O should be included. Additional RNDs may introduce ambiguity and may not be verifiable |


Generated: 31/10/2022 at 16:43:22