# STIR/SHAKEN CA Ecosystem Compliance

## Ribbon Communications
Name: w_pki_ca_subject_rdn_unknown\
Source: SHAKEN PKI Best Practice\
Citation: SHAKEN PKI Best Practice\
Effective Date: 01 Jan 00 00:00 UTC\
Description: Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject. However, the names should be meaningful meaningful enough to represent the SP to whom the certificate is being issued, in a manner similar to that used to identify SP’s equipment in the network.

### Leaf Certificates

| Status | Subject | Link | Details |
|--------|---------|------|---------|

no warning, or error, or not effective date level issues were found

### CA Certificates

| Status | Subject | Link | Details |
|--------|---------|------|---------|
| warn | CN=SHAKEN Ribbon Root CA, OU=Certification Authority, O=Ribbon Communications, C=US | [view](../../CERTIFICATES/5e00d019ee3884ec7d8513deb8f4e8671a1db86e/README.md) | Only CN, C, L, and O should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| warn | CN=SHAKEN Ribbon Issuing CA, OU=Certification Authorities, O=Ribbon Communications, C=US | [view](../../CERTIFICATES/cac26ee453b887be4c545426c19733add7c138c0/README.md) | Only CN, C, L, and O should be included. Additional RNDs may introduce ambiguity and may not be verifiable |


Generated: 28/10/2022 at 19:22:10