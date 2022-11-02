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
| warn | VOCALTRANSIT SHAKEN 783J | [view](../../CERTS/81b78fff8a772249d72d4854d97672d7ac69a83c4900beaac699d28d220d8c13/README.md) | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| warn | Definitely SHAKEN 709J but not SHAKEN 0007 | [view](../../CERTS/8960e7cbd26618777e4018fdaf9853ff67f176b10954cb37596a75c2dd0a4ddb/README.md) | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| warn | Definitely SHAKEN 709J but not SHAKEN 0007 | [view](../../CERTS/148fa003df4d2ae06a49e002c44e3dd1a78feedd963cfb762b7a07829c496be8/README.md) | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| warn | Definitely SHAKEN 709J but not SHAKEN 0007 | [view](../../CERTS/d35cdea11102a752f34309c8ecae597b4b7c1caf7f41dea46415471eb5317cce/README.md) | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |

### CA Certificates

No error, warning, or notice level issues were found


Generated: 02 Nov 22 15:15 UTC