# STIR/SHAKEN CA Ecosystem Compliance

## NetNumber

### Summary

\* The percent of errors, warnings and notices is calculated against total observed unexpired and trusted certificates from the specified issuer.\
\*\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran.

#### Leaf Certificates

- 51 certificates were included in the corpus being tested
- 0 certificates in the corpus were skipped because they are duplicates
- 45 certificates in the corpus were skipped because they are expired
- 0 certificates in the corpus were skipped because they are not currently trusted
- 6 certificates being tested against the remaining rules
- 4.00 issues on average found in unexpired, trusted, and non-compliant certificates
- 100.00% of certificates contain one or more Error level issue
- 0.00% of certificates contain one or more Warning level issue
- 0.00% of certificates contain one or more Notice level issue
- 0.00% of certificates are too old to be assessed against currently enforced expectations
- 125 days is the average remaining validity for the certificates in the corpus
- 142 days is the average initial validity for the certificates in the corpus
- 4 certificates expire in the next 30 days
- 1.20 average number of unexpired certificates per OCN observed
- 5 unique OCNs observed in unexpired and valid certificate corpus

| Instances | Test | Source |
|-----------|------|--------|
| 6 | [e_atis_ext_certificate_policies](ISSUES/e_atis_ext_certificate_policies/README.md) | ATIS1000080 |
| 2 | [e_atis_ext_crl_distribution](ISSUES/e_atis_ext_crl_distribution/README.md) | ATIS1000080 |
| 2 | [e_atis_serial_number_size](ISSUES/e_atis_serial_number_size/README.md) | ATIS1000080 |
| 2 | [e_atis_subject_cn](ISSUES/e_atis_subject_cn/README.md) | ATIS1000080 |
| 6 | [e_atis_subject_cn_spc](ISSUES/e_atis_subject_cn_spc/README.md) | ATIS1000080 |
| 6 | [e_shaken_certificate_policies_id](ISSUES/e_shaken_certificate_policies_id/README.md) | US_SHAKEN_CP |

#### CA Certificates

- 3 certificates were included in the corpus being tested
- 0 certificates in the corpus were skipped because they are duplicates
- 0 certificates in the corpus were skipped because they are expired
- 0 certificates in the corpus were skipped because they are not currently trusted
- 3 certificates being tested against the remaining rules
- 3.00 issues on average found in unexpired, trusted, and non-compliant certificates
- 66.67% of certificates contain one or more Error level issue
- 0.00% of certificates contain one or more Warning level issue
- 0.00% of certificates contain one or more Notice level issue
- 100.00% of certificates are too old to be assessed against currently enforced expectations
- 7266 days is the average remaining validity for the certificates in the corpus
- 6935 days is the average initial validity for the certificates in the corpus
- 0 certificates expire in the next 30 days

| Instances | Test | Source |
|-----------|------|--------|
| 1 | [e_atis_ext_certificate_policies_ca](ISSUES/e_atis_ext_certificate_policies_ca/README.md) | ATIS1000080 |
| 2 | [e_atis_signature_algorithm_ca](ISSUES/e_atis_signature_algorithm_ca/README.md) | ATIS1000080 |
| 2 | [e_atis_subject_public_key_ca](ISSUES/e_atis_subject_public_key_ca/README.md) | ATIS1000080 |
| 1 | [e_shaken_certificate_policies_id_ca](ISSUES/e_shaken_certificate_policies_id_ca/README.md) | US_SHAKEN_CP |

### Details

#### Leaf Certificates

| Created At | Subject | Not After | Problems | Link |
|------------|---------|-----------|----------|------|
| 18&#160;May&#160;23&#160;00:00&#160;UTC | Plivo Inc | 17&#160;May&#160;24&#160;00:00&#160;UTC | true | [view](CERTS/fed50200daa631dd0cd7b74969c780f8d456dfd31db156c6cbea276f5a9a4cbf/README.md) |
| 25&#160;Oct&#160;23&#160;15:48&#160;UTC | DISH Wireless L.L.C.SHAKEN.490J | 24&#160;Oct&#160;24&#160;15:48&#160;UTC | true | [view](CERTS/2943713c56f0705ed027ecffced5eb89cb1c36bb5386bdc36a6b8e5618ca2c9c/README.md) |
| 14&#160;Jan&#160;24&#160;08:02&#160;UTC | Google SHAKEN cert 969H | 13&#160;Feb&#160;24&#160;08:02&#160;UTC | true | [view](CERTS/219583be0a46879699b163cfc1c6ee25f2fae2ea6c83a5889fd211a03ce88d44/README.md) |
| 16&#160;Jan&#160;24&#160;00:00&#160;UTC | HD CARRIER LLC | 14&#160;Feb&#160;24&#160;23:59&#160;UTC | true | [view](CERTS/d61794931ed919fde06b2ea87b472356ae5b2a87476516b62bcf6c5d58a1f996/README.md) |
| 02&#160;Feb&#160;24&#160;15:00&#160;UTC | Baltimore-Washington Telephone Company SHAKEN cert 8697 | 02&#160;Mar&#160;24&#160;14:59&#160;UTC | true | [view](CERTS/a037bd3bca14bd52993d5449264f63b11eb8cda01885268c849b97c043444820/README.md) |
| 06&#160;Feb&#160;24&#160;12:05&#160;UTC | Google SHAKEN cert 969H | 07&#160;Mar&#160;24&#160;12:05&#160;UTC | true | [view](CERTS/bcb2b91c1a3df613ba4f9e298d10bc26c44554cf155845b60106c0aea06588cb/README.md) |

#### CA Certificates

| Created At | Subject | Not After | Problems | Link |
|------------|---------|-----------|----------|------|
| 12&#160;Jul&#160;21&#160;23:25&#160;UTC | NetNumber SHAKEN Root CA | 07&#160;Jul&#160;41&#160;23:25&#160;UTC | false | [view](CERTS/7ac80e8481ecb019dc95484016842db78686069efbc0f703e7f39310217b6157/README.md) |
| 27&#160;Sep&#160;21&#160;19:45&#160;UTC | NetNumber SHAKEN Root CA 1 | 21&#160;Sep&#160;46&#160;19:45&#160;UTC | true | [view](CERTS/2dd1386ca717f31d550b35b9bce9daa9b02483bcdb98bdfcfca07202276136d7/README.md) |
| 29&#160;Sep&#160;21&#160;13:22&#160;UTC | NetNumber SHAKEN Root Intermediate CA 1 | 26&#160;Sep&#160;33&#160;13:22&#160;UTC | true | [view](CERTS/e449803766edf02ab50b034dd7e89e54efd332cce87688a032f89b340d039878/README.md) |


Generated: 12 Feb 24 17:02 UTC