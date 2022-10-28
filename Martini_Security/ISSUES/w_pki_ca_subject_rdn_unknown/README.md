# STIR/SHAKEN CA Ecosystem Compliance

## Martini Security
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
| warn | CN=Martini Security SHAKEN R1, O=Martini Security\\, LLC, L=Seattle, ST=WA, C=US | [view](../../CERTIFICATES/0ffc3156067f94ee63a0ff6f33fce9cff5f51d39/README.md) | Only CN, C, L, and O should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| warn | CN=Martini Security SHAKEN G1, O=Martini Security\\, LLC, L=Seattle, ST=WA, C=US | [view](../../CERTIFICATES/0fe086f321e93ca9ae08a19a89bf9049b7625fcf/README.md) | Only CN, C, L, and O should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| warn | CN=Martini Security SHAKEN G2, O=Martini Security\\, LLC, L=Seattle, ST=WA, C=US | [view](../../CERTIFICATES/6af848efe4c680e44114ee32350826704eca8135/README.md) | Only CN, C, L, and O should be included. Additional RNDs may introduce ambiguity and may not be verifiable |


Generated: 28/10/2022 at 19:22:10