# STIR/SHAKEN CA Ecosystem Compliance

## Metaswitch
Name: w_pki_subject_rdn_unknown\
Source: SHAKEN PKI Best Practice\
Citation: SHAKEN PKI Best Practice\
Effective Date: 01 Jan 00 00:00 UTC\
Description: Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject. However, the names should be meaningful meaningful enough to represent the SP to whom the certificate is being issued, in a manner similar to that used to identify SP’s equipment in the network.

### Leaf Certificates

| Status | Subject | Link | Details |
|--------|---------|------|---------|
| warn | CN=Verizon SHAKEN cert 5807, OU=NNO CDS, O=Verizon Data Services LLC, L=Southlake, ST=Texas, C=US | [view](../../CERTIFICATES/b3b4954ec8815349931cd0c6947d42a12569872f/README.md) | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| warn | CN=Allstream SHAKEN cert 4130, O=Allstream Business US\\, LLC, L=Vancouver, ST=WA, C=US | [view](../../CERTIFICATES/ff4fd6ee8ba51ca3158a8f6e11a2d6ddef2effb7/README.md) | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| warn | CN=USCellular SHAKEN Cert 6349, L=Ch, ST=IL, C=US | [view](../../CERTIFICATES/9215e4a768eaf6b5209dfeccc846e5d604061979/README.md) | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| warn | CN=U. S. Telepacific Corp SHAKEN 7453, O=U. S. Telepacific Corp, L=Los Angeles, ST=California, C=US | [view](../../CERTIFICATES/7deaf409e0c7859925f43a64b96cdd8b0eb1be89/README.md) | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| warn | CN=Mediacom 846F, O=Mediacom Communications Corporation, L=Chester, ST=New York, C=US | [view](../../CERTIFICATES/82b4c63aa246ac8c380b66179ec2e381c89dde4a/README.md) | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| warn | CN=Sonic Telecom SHAKEN cert 433E, O=Sonic Telecom, ST=CA, C=US | [view](../../CERTIFICATES/954ca8988aa2a569c678c36777914fd6f1cd551a/README.md) | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |

### CA Certificates

| Status | Subject | Link | Details |
|--------|---------|------|---------|

no warning, or error, or not effective date level issues were found


Generated: 28/10/2022 at 18:55:00