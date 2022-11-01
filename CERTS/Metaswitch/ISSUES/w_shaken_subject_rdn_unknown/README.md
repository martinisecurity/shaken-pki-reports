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
| warn | Allstream SHAKEN cert 4130 | [view](../../CERTS/a9eec584cf66eade89fa84e043e0eb05900dc40653ea963a96c079ef92d9e2ad/README.md) | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| warn | Sonic Telecom SHAKEN cert 433E | [view](../../CERTS/a5082b808b3bf200a634e273b7988a617099c9c897a9c596858e6d4ffe4fd352/README.md) | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| warn | USCellular SHAKEN Cert 6349 | [view](../../CERTS/c75937e7c843b0054b1a51f6149f47661696b3cfcbcf6bab4dbc708c4453244f/README.md) | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| warn | U. S. Telepacific Corp SHAKEN 7453 | [view](../../CERTS/9ed03dac797a5a27d52aa5209a4caa6a3ec9c3943d55a2cbfb69416480787da0/README.md) | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| warn | Mediacom 846F | [view](../../CERTS/e6c9e9fd411d8174b3ffe1af4d569c6919f4b98a5d0c6e429cd3682d82284e7e/README.md) | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| warn | Verizon SHAKEN cert 5807 | [view](../../CERTS/d7b413267be2d050d516af4f4a864ffdc2feacc774a1a6264b9cfe68c796f43f/README.md) | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |

### CA Certificates

no warning, or error, or not effective date level issues were found


Generated: 01/11/2022 at 20:34:21