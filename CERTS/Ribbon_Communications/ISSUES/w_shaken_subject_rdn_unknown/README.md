# STIR/SHAKEN CA Ecosystem Compliance

## Ribbon Communications

Name: w_shaken_subject_rdn_unknown\
Source: SHAKEN_PKI_BEST_PRACTICES\
Citation: SHAKEN PKI Best Practice\
Effective Date: 01 Jan 00 00:00 UTC\
Description: Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject. However, the names should be meaningful meaningful enough to represent the SP to whom the certificate is being issued, in a manner similar to that used to identify SP’s equipment in the network.

### Leaf Certificates

| Status | Subject | Link | Details |
|--------|---------|------|---------|
| warn | Nuwave Communications SHAKEN 620J | [view](../../CERTS/885b7fea5177d66a28bd5dea525144479428e27d2ed1ff7c493d767c73173fbc/README.md) | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| warn | Veracity SHAKEN 716D | [view](../../CERTS/fccfbb719ccc9513231e9ea6f38906f0271f4640f253e1be0780da1e7b5f03ff/README.md) | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| warn | Netfortris SHAKEN 8886 | [view](../../CERTS/64b7a3eed364e863b36e050a95b35799c594ea1c15c17562611907a2f0dd8bbe/README.md) | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |

### CA Certificates

No error, warning, or notice level issues were found


Generated: 02 Nov 22 15:41 UTC