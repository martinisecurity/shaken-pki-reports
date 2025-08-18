# STIR/SHAKEN CA Ecosystem Compliance

## Ribbon Communications

### Summary

\* The percent of errors, warnings and notices is calculated against total observed unexpired and trusted certificates from the specified issuer.\
\*\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran.

#### Leaf Certificates

- 31 certificates were included in the corpus being tested
- 0 certificates in the corpus were skipped because they are duplicates
- 19 certificates in the corpus were skipped because they are expired
- 0 certificates in the corpus were skipped because they are not currently trusted
- 12 certificates being tested against the remaining rules
- 1.08 issues on average found in unexpired, trusted, and non-compliant certificates
- 100.00% of certificates contain one or more Error level issue
- 0.00% of certificates contain one or more Warning level issue
- 0.00% of certificates contain one or more Notice level issue
- 0.00% of certificates are too old to be assessed against currently enforced expectations
- 347 days is the average remaining validity for the certificates in the corpus
- 365 days is the average initial validity for the certificates in the corpus
- 1 certificates expire in the next 30 days
- 1.00 average number of unexpired certificates per OCN observed
- 12 unique OCNs observed in unexpired and valid certificate corpus

| Instances | Test | Source |
|-----------|------|--------|
| 12 | [e_atis_ext_crl_distribution_struct](ISSUES/e_atis_ext_crl_distribution_struct/README.md) | ATIS1000080 |
| 1 | [e_atis_subject_cn_spc](ISSUES/e_atis_subject_cn_spc/README.md) | ATIS1000080 |

#### CA Certificates

- 3 certificates were included in the corpus being tested
- 0 certificates in the corpus were skipped because they are duplicates
- 0 certificates in the corpus were skipped because they are expired
- 0 certificates in the corpus were skipped because they are not currently trusted
- 3 certificates being tested against the remaining rules
- 1.00 issues on average found in unexpired, trusted, and non-compliant certificates
- 66.67% of certificates contain one or more Error level issue
- 0.00% of certificates contain one or more Warning level issue
- 0.00% of certificates contain one or more Notice level issue
- 66.67% of certificates are too old to be assessed against currently enforced expectations
- 6368 days is the average remaining validity for the certificates in the corpus
- 5966 days is the average initial validity for the certificates in the corpus
- 0 certificates expire in the next 30 days

| Instances | Test | Source |
|-----------|------|--------|
| 1 | [e_atis_ext_crl_distribution_ca](ISSUES/e_atis_ext_crl_distribution_ca/README.md) | ATIS1000080 |
| 1 | [e_atis_ext_crl_distribution_struct_ca](ISSUES/e_atis_ext_crl_distribution_struct_ca/README.md) | ATIS1000080 |

### Details

#### Leaf Certificates

| Created At | Subject | Not After | Problems | Link |
|------------|---------|-----------|----------|------|
| 23&#160;Aug&#160;24&#160;18:34&#160;UTC | CBTS STIR SHAKEN 797G | 23&#160;Aug&#160;25&#160;18:34&#160;UTC | true | [view](CERTS/f555850de2d64a537e25bde110b289355dd1da825c16b9183e3fcd928e56ba17/README.md) |
| 02&#160;Jan&#160;25&#160;15:10&#160;UTC | Empire SHAKEN 087H | 02&#160;Jan&#160;26&#160;15:10&#160;UTC | true | [view](CERTS/aef5bf41cadd87ad794165085f619fe1d4cd5756d6634e6b481d8d2a9b15903e/README.md) |
| 07&#160;Jan&#160;25&#160;16:51&#160;UTC | Nuwave Communications SHAKEN 620J | 07&#160;Jan&#160;26&#160;16:51&#160;UTC | true | [view](CERTS/6ceeeef74a5105b07fcd2d6eda97acbe7172835ff8ce8f1ae70def5c86047bf1/README.md) |
| 27&#160;Mar&#160;25&#160;03:09&#160;UTC | Etex STI SHAKEN 3196 | 27&#160;Mar&#160;26&#160;03:09&#160;UTC | true | [view](CERTS/66999f4d7f9b10ca4dc5cbb1b3a8d729e1a0d42b63a884bcbfafa0faeb92e7f8/README.md) |
| 05&#160;Apr&#160;25&#160;05:18&#160;UTC | FirstDigital SHAKEN 5727 | 05&#160;Apr&#160;26&#160;05:18&#160;UTC | true | [view](CERTS/38e9d97c418a538d2d2b0535fd7bca410919ead6562c2bfbbd17a8c7ece32f6a/README.md) |
| 02&#160;May&#160;25&#160;15:28&#160;UTC | Netfortris SHAKEN 8886 | 02&#160;May&#160;26&#160;15:28&#160;UTC | true | [view](CERTS/7145382dacf58b945036cffa86c1cc8e743362a7adf1ab946bf85351bd19ea7f/README.md) |
| 05&#160;May&#160;25&#160;13:42&#160;UTC | Ribbon SHAKEN 2086 | 05&#160;May&#160;26&#160;13:42&#160;UTC | true | [view](CERTS/622e779cda1e66eb9c352333c842da7b6527eb3dc7a9309fd290dab423c3be70/README.md) |
| 07&#160;May&#160;25&#160;13:44&#160;UTC | GreatWave SHAKEN 032H | 07&#160;May&#160;26&#160;13:44&#160;UTC | true | [view](CERTS/ca8d19bd46fa3583085af2e69827f24d0cbfc4000a52b4193d512ae2fabdc66f/README.md) |
| 23&#160;May&#160;25&#160;19:17&#160;UTC | Ironton SHAKEN 1234 0175 | 23&#160;May&#160;26&#160;19:17&#160;UTC | true | [view](CERTS/c51d8e73a75e31750389f8c938c01a9954650165849da6af2241c7ae364431b2/README.md) |
| 03&#160;Jun&#160;25&#160;16:02&#160;UTC | UniVoIP SHAKEN 629J | 03&#160;Jun&#160;26&#160;16:02&#160;UTC | true | [view](CERTS/04ff1c8d7bebd8132d4abbb64e912d4e172b38f89787809abcd4c2e0ce067cf1/README.md) |
| 12&#160;Jun&#160;25&#160;17:28&#160;UTC | Fulton SHAKEN 0455 | 12&#160;Jun&#160;26&#160;17:28&#160;UTC | true | [view](CERTS/f0ab223c1b2aef45b1286f5b2a435813915c01a8d1438cef9afb6b1a0406ca3c/README.md) |
| 13&#160;Jun&#160;25&#160;16:23&#160;UTC | SWTEXAS SHAKEN 2135 | 13&#160;Jun&#160;26&#160;16:23&#160;UTC | true | [view](CERTS/417aa4d76c9dd30594df52c1118360169dd07f3efb93d807d9de8614098ed99c/README.md) |

#### CA Certificates

| Created At | Subject | Not After | Problems | Link |
|------------|---------|-----------|----------|------|
| 13&#160;May&#160;21&#160;00:00&#160;UTC | SHAKEN Ribbon Issuing CA | 12&#160;May&#160;33&#160;23:59&#160;UTC | true | [view](CERTS/05a71a04eaedbdf4b0534f40768616d7c19c8deb5a3aefd1f4a04b3aab55a48f/README.md) |
| 13&#160;May&#160;21&#160;00:00&#160;UTC | SHAKEN Ribbon Root CA | 12&#160;May&#160;46&#160;23:59&#160;UTC | false | [view](CERTS/7c2799d3642d04f04fe667c3ab251c18689af323acdc43b2fa5f3dc89e3a0f14/README.md) |
| 12&#160;Mar&#160;24&#160;00:00&#160;UTC | SHAKEN Ribbon Issuing CA 2 | 11&#160;Mar&#160;36&#160;23:59&#160;UTC | true | [view](CERTS/2b48949bed753b3edc3645a638704bd1e708cd047117e3fe8af6cb144fa3a8c5/README.md) |


Generated: 18 Aug 25 21:13 UTC