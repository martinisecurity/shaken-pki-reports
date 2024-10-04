# STIR/SHAKEN CA Ecosystem Compliance

## Ribbon Communications

### Summary

\* The percent of errors, warnings and notices is calculated against total observed unexpired and trusted certificates from the specified issuer.\
\*\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran.

#### Leaf Certificates

- 31 certificates were included in the corpus being tested
- 0 certificates in the corpus were skipped because they are duplicates
- 17 certificates in the corpus were skipped because they are expired
- 0 certificates in the corpus were skipped because they are not currently trusted
- 14 certificates being tested against the remaining rules
- 1.64 issues on average found in unexpired, trusted, and non-compliant certificates
- 100.00% of certificates contain one or more Error level issue
- 0.00% of certificates contain one or more Warning level issue
- 0.00% of certificates contain one or more Notice level issue
- 0.00% of certificates are too old to be assessed against currently enforced expectations
- 343 days is the average remaining validity for the certificates in the corpus
- 365 days is the average initial validity for the certificates in the corpus
- 1 certificates expire in the next 30 days
- 1.00 average number of unexpired certificates per OCN observed
- 14 unique OCNs observed in unexpired and valid certificate corpus

| Instances | Test | Source |
|-----------|------|--------|
| 7 | [e_atis_ext_crl_distribution_struct](ISSUES/e_atis_ext_crl_distribution_struct/README.md) | ATIS1000080 |
| 8 | [e_atis_ext_not_specified](ISSUES/e_atis_ext_not_specified/README.md) | ATIS1000080 |
| 1 | [e_atis_subject_cn_spc](ISSUES/e_atis_subject_cn_spc/README.md) | ATIS1000080 |
| 7 | [e_shaken_certificate_policies_id](ISSUES/e_shaken_certificate_policies_id/README.md) | US_SHAKEN_CP |

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
- 6447 days is the average remaining validity for the certificates in the corpus
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
| 02&#160;Nov&#160;23&#160;13:37&#160;UTC | Pierce SHAKEN 1581 | 01&#160;Nov&#160;24&#160;13:37&#160;UTC | true | [view](CERTS/18c68e30226bed3373645c691175d5b2c3b83d51a4132914ced3a25b9f28aa2e/README.md) |
| 04&#160;Jan&#160;24&#160;17:59&#160;UTC | Empire SHAKEN 087H | 03&#160;Jan&#160;25&#160;17:59&#160;UTC | true | [view](CERTS/ff4a083ed7527187359cb180383afae989c9a1144ad498e79ad3e3e05f5091fb/README.md) |
| 22&#160;Jan&#160;24&#160;14:57&#160;UTC | Eastex SHAKEN 2068 | 21&#160;Jan&#160;25&#160;14:57&#160;UTC | true | [view](CERTS/edc37467b2db082bbed1e00e885cf81f7bb179ad1ac3a75fe5f82f7c284da595/README.md) |
| 30&#160;Jan&#160;24&#160;18:59&#160;UTC | NCTC SHAKEN 0573 | 29&#160;Jan&#160;25&#160;18:59&#160;UTC | true | [view](CERTS/061fc531a3791e3dc8f6ab92fa774fcd34159cdaca5806df7c2f1ef77e1f9541/README.md) |
| 02&#160;Feb&#160;24&#160;14:07&#160;UTC | UniVoIP SHAKEN 629J | 01&#160;Feb&#160;25&#160;14:07&#160;UTC | true | [view](CERTS/aa242e08e8ae190ac6bb6f4592db929e3ff13278dac0629a9db7be151e253914/README.md) |
| 16&#160;Feb&#160;24&#160;19:41&#160;UTC | Nuwave Communications SHAKEN 620J | 15&#160;Feb&#160;25&#160;19:41&#160;UTC | true | [view](CERTS/e9e2df4bebfaedd61373b3a9ddc2be32e19d81d2431f40ddb790185ded4b6783/README.md) |
| 07&#160;Mar&#160;24&#160;22:41&#160;UTC | SWAT SHAKEN 1724 | 07&#160;Mar&#160;25&#160;22:41&#160;UTC | true | [view](CERTS/ba78c647ff4975e21e1538d4ff4f3a36d8ce65ed7ed2cc918370df2c02b2eee5/README.md) |
| 24&#160;Mar&#160;24&#160;18:27&#160;UTC | ColoradoValley SHAKEN 2059 | 24&#160;Mar&#160;25&#160;18:27&#160;UTC | true | [view](CERTS/328011b837c2166f113664ab01dc719816e4bfcb734782636e9abb7bbf68199a/README.md) |
| 01&#160;Apr&#160;24&#160;13:12&#160;UTC | Etex STI SHAKEN 3196 | 01&#160;Apr&#160;25&#160;13:12&#160;UTC | true | [view](CERTS/9cb289044279fde9ca74c37807eb10061a1098c816b89e2d146e37c0d377e044/README.md) |
| 29&#160;Apr&#160;24&#160;05:30&#160;UTC | FirstDigital SHAKEN 5727 | 29&#160;Apr&#160;25&#160;05:30&#160;UTC | true | [view](CERTS/cf281053153a12955be41e5b0f99caf1572a1c30b21f62a5c22543bf337740f7/README.md) |
| 08&#160;May&#160;24&#160;17:05&#160;UTC | Ribbon SHAKEN 2086 | 08&#160;May&#160;25&#160;17:05&#160;UTC | true | [view](CERTS/a8850907d694bd9b9787d39dd18aff77488f3ff1d12a045b5a112dde73d8d5b2/README.md) |
| 13&#160;May&#160;24&#160;22:02&#160;UTC | Localtel SHAKEN 3229 | 13&#160;May&#160;25&#160;22:02&#160;UTC | true | [view](CERTS/92063be4763362c7e542a41e3f4ebd030f1816f5ec8600b444be908c646307e5/README.md) |
| 30&#160;May&#160;24&#160;19:04&#160;UTC | Ironton SHAKEN 1234 0175 | 30&#160;May&#160;25&#160;19:04&#160;UTC | true | [view](CERTS/897d53482f78c037e9dd54cc3e62a1a3dba25a5abff81f1faf3ec09ef58bb5a0/README.md) |
| 06&#160;Jun&#160;24&#160;13:33&#160;UTC | Peoples SHAKEN 2130 | 06&#160;Jun&#160;25&#160;13:33&#160;UTC | true | [view](CERTS/506e5980d9b8f24ac2c781f15c58fcc50dd50c9f82d8aac2ac72967211f262db/README.md) |

#### CA Certificates

| Created At | Subject | Not After | Problems | Link |
|------------|---------|-----------|----------|------|
| 13&#160;May&#160;21&#160;00:00&#160;UTC | SHAKEN Ribbon Issuing CA | 12&#160;May&#160;33&#160;23:59&#160;UTC | true | [view](CERTS/05a71a04eaedbdf4b0534f40768616d7c19c8deb5a3aefd1f4a04b3aab55a48f/README.md) |
| 13&#160;May&#160;21&#160;00:00&#160;UTC | SHAKEN Ribbon Root CA | 12&#160;May&#160;46&#160;23:59&#160;UTC | false | [view](CERTS/7c2799d3642d04f04fe667c3ab251c18689af323acdc43b2fa5f3dc89e3a0f14/README.md) |
| 12&#160;Mar&#160;24&#160;00:00&#160;UTC | SHAKEN Ribbon Issuing CA 2 | 11&#160;Mar&#160;36&#160;23:59&#160;UTC | true | [view](CERTS/2b48949bed753b3edc3645a638704bd1e708cd047117e3fe8af6cb144fa3a8c5/README.md) |


Generated: 04 Oct 24 16:29 UTC