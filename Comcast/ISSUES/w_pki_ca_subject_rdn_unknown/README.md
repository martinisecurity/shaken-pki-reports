# STIR/SHAKEN CA Ecosystem Compliance

## Comcast
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
| warn | CN=Comcast SHAKEN Root CA, O=Comcast, L=Philadelphia, ST=Pennsylvania, C=US | [view](../../CERTIFICATES/e341fff079ef701a75085e21aaa915d84a27a52a/README.md) | Only CN, C, L, and O should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| warn | CN=Comcast SHAKEN Intermediate CA, O=Comcast, ST=Pennsylvania, C=US | [view](../../CERTIFICATES/39011602b92be825bea3e29648f2e1866d60d0c6/README.md) | Only CN, C, L, and O should be included. Additional RNDs may introduce ambiguity and may not be verifiable |


Generated: 27/10/2022 at 22:44:50