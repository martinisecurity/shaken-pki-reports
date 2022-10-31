# STIR/SHAKEN CA Ecosystem Compliance

## TransNexus

Name: w_pki_subject_rdn_unknown\
Source: SHAKEN PKI Best Practice\
Citation: SHAKEN PKI Best Practice\
Effective Date: 01 Jan 00 00:00 UTC\
Description: Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject. However, the names should be meaningful meaningful enough to represent the SP to whom the certificate is being issued, in a manner similar to that used to identify SP’s equipment in the network.

### Leaf Certificates

| Status | Subject | Link | Details |
|--------|---------|------|---------|
| warn | SHAKEN 012K | [view](../../CERTS/788e828c8fe34b9e0116e8f5d3a7e6d71a22a83803056a63b5ab35f2ec52f6cf/README.md) | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |

### CA Certificates

no warning, or error, or not effective date level issues were found


Generated: 31/10/2022 at 16:43:22