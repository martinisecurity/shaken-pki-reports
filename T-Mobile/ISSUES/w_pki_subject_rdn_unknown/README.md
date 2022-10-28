# STIR/SHAKEN CA Ecosystem Compliance

## T-Mobile
Name: w_pki_subject_rdn_unknown\
Source: SHAKEN PKI Best Practice\
Citation: SHAKEN PKI Best Practice\
Effective Date: 01 Jan 00 00:00 UTC\
Description: Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject. However, the names should be meaningful meaningful enough to represent the SP to whom the certificate is being issued, in a manner similar to that used to identify SP’s equipment in the network.

### Leaf Certificates

| Status | Subject | Link | Details |
|--------|---------|------|---------|
| warn | CN=cert.stir.t-mobile.com, OU=T-Mobile USA\\, Inc, O=T-Mobile USA\\, Inc., L=Bothell, ST=Washington, C=US | [view](../../CERTIFICATES/9ac0def10180cb22aad9e8885715c9409031697a/README.md) | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |

### CA Certificates

| Status | Subject | Link | Details |
|--------|---------|------|---------|

no warning, or error, or not effective date level issues were found


Generated: 28/10/2022 at 10:33:25