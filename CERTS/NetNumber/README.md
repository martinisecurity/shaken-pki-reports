# STIR/SHAKEN CA Ecosystem Compliance

## NetNumber

### Summary

\* The percent of errors, warnings and notices is calculated against total observed unexpired and trusted certificates from the specified issuer.\
\*\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran.

#### Leaf Certificates

- 25 certificates were included in the corpus being tested
- 0 certificates in the corpus were skipped because they are duplicates
- 17 certificates in the corpus were skipped because they are expired
- 0 certificates in the corpus were skipped because they are not currently trusted
- 8 certificates being tested against the remaining rules
- 5.38 issues on average found in unexpired, trusted, and non-compliant certificates
- 100.00% of certificates contain one or more Error level issue
- 12.50% of certificates contain one or more Warning level issue
- 100.00% of certificates contain one or more Notice level issue
- 0.00% of certificates are too old to be assessed against currently enforced expectations
- 141 days is the average remaining validity for the certificates in the corpus
- 156 days is the average initial validity for the certificates in the corpus
- 6 certificates expire in the next 30 days
- 1.33 average number of unexpired certificates per OCN observed
- 6 unique OCNs observed in unexpired and valid certificate corpus

| Instances | Test | Source |
|-----------|------|--------|
| 8 | [e_atis_certificate_policies](ISSUES/e_atis_certificate_policies/README.md) | ATIS1000080 |
| 2 | [e_atis_crl_distribution](ISSUES/e_atis_crl_distribution/README.md) | ATIS1000080 |
| 1 | [e_atis_subject](ISSUES/e_atis_subject/README.md) | ATIS1000080 |
| 7 | [e_atis_subject_cn](ISSUES/e_atis_subject_cn/README.md) | ATIS1000080 |
| 8 | [e_us_cp_ambiguous_identifier](ISSUES/e_us_cp_ambiguous_identifier/README.md) | US_SHAKEN_CP |
| 8 | [e_us_cp_subject_sn](ISSUES/e_us_cp_subject_sn/README.md) | US_SHAKEN_CP |
| 8 | [n_atis_certificate_policy_critical](ISSUES/n_atis_certificate_policy_critical/README.md) | ATIS1000080 |
| 1 | [w_shaken_subject_rdn_unknown](ISSUES/w_shaken_subject_rdn_unknown/README.md) | SHAKEN_PKI_BEST_PRACTICES |

#### CA Certificates

- 3 certificates were included in the corpus being tested
- 0 certificates in the corpus were skipped because they are duplicates
- 0 certificates in the corpus were skipped because they are expired
- 0 certificates in the corpus were skipped because they are not currently trusted
- 3 certificates being tested against the remaining rules
- 3.00 issues on average found in unexpired, trusted, and non-compliant certificates
- 66.67% of certificates contain one or more Error level issue
- 66.67% of certificates contain one or more Warning level issue
- 0.00% of certificates contain one or more Notice level issue
- 100.00% of certificates are too old to be assessed against currently enforced expectations
- 7330 days is the average remaining validity for the certificates in the corpus
- 6935 days is the average initial validity for the certificates in the corpus
- 0 certificates expire in the next 30 days

| Instances | Test | Source |
|-----------|------|--------|
| 2 | [e_atis_ca_signature_algorithm](ISSUES/e_atis_ca_signature_algorithm/README.md) | ATIS1000080 |
| 2 | [e_atis_ca_subject_public_key](ISSUES/e_atis_ca_subject_public_key/README.md) | ATIS1000080 |
| 2 | [w_shaken_ca_subject_rdn_unknown](ISSUES/w_shaken_ca_subject_rdn_unknown/README.md) | SHAKEN_PKI_BEST_PRACTICES |

### Details

#### Leaf Certificates

| Created At | Subject | Not After | Problems | Link |
|------------|---------|-----------|----------|------|
| 10&#160;Jun&#160;22&#160;00:00&#160;UTC | Plivo Inc | 09&#160;Jun&#160;23&#160;00:00&#160;UTC | true | [view](CERTS/7dc750fb7aa68d2b67b8dbc89f65217f92db54504685058be016638011adf8bf/README.md) |
| 16&#160;Nov&#160;22&#160;17:07&#160;UTC | DISH Wireless L.L.C.SHAKEN.490J | 16&#160;Nov&#160;23&#160;17:07&#160;UTC | true | [view](CERTS/58eec25518890bf640e859b8e81eef58fae2b641d6b5f5b4f814ec074a842cbb/README.md) |
| 05&#160;May&#160;23&#160;18:59&#160;UTC | Google SHAKEN cert 969H | 04&#160;Jun&#160;23&#160;18:59&#160;UTC | true | [view](CERTS/bc8177375fc44146fc0f541f135a1b7e8a226083942d5d4ef18a799fe9846bcc/README.md) |
| 10&#160;May&#160;23&#160;00:10&#160;UTC | EssexTel INC SHAKEN 692J | 09&#160;Jun&#160;23&#160;00:10&#160;UTC | true | [view](CERTS/68addc9c2fee2ee1d2df7dbf1957dcbd1c022162edd957c9edaf73b87d070fd3/README.md) |
| 16&#160;May&#160;23&#160;00:00&#160;UTC | HD CARRIER LLC | 14&#160;Jun&#160;23&#160;23:59&#160;UTC | true | [view](CERTS/33231fbb0dfc562c8d9b084829fd82bd4c8412dd6cc8f130f60e990d1d7b596d/README.md) |
| 18&#160;May&#160;23&#160;00:00&#160;UTC | Plivo Inc | 17&#160;May&#160;24&#160;00:00&#160;UTC | true | [view](CERTS/fed50200daa631dd0cd7b74969c780f8d456dfd31db156c6cbea276f5a9a4cbf/README.md) |
| 19&#160;May&#160;23&#160;15:00&#160;UTC | Baltimore-Washington Telephone Company SHAKEN cert 8697 | 19&#160;Jun&#160;23&#160;14:59&#160;UTC | true | [view](CERTS/5e8837b7a292d1535d22bee99322d6752e937f2ed99f5865559c374659398f03/README.md) |
| 28&#160;May&#160;23&#160;18:59&#160;UTC | Google SHAKEN cert 969H | 27&#160;Jun&#160;23&#160;18:59&#160;UTC | true | [view](CERTS/06bba32cdbf4eefaa54f3f6357c482cecc356fc385955a329a13cfd13591502f/README.md) |

#### CA Certificates

| Created At | Subject | Not After | Problems | Link |
|------------|---------|-----------|----------|------|
| 12&#160;Jul&#160;21&#160;23:25&#160;UTC | NetNumber SHAKEN Root CA | 07&#160;Jul&#160;41&#160;23:25&#160;UTC | false | [view](CERTS/7ac80e8481ecb019dc95484016842db78686069efbc0f703e7f39310217b6157/README.md) |
| 27&#160;Sep&#160;21&#160;19:45&#160;UTC | NetNumber SHAKEN Root CA 1 | 21&#160;Sep&#160;46&#160;19:45&#160;UTC | true | [view](CERTS/2dd1386ca717f31d550b35b9bce9daa9b02483bcdb98bdfcfca07202276136d7/README.md) |
| 29&#160;Sep&#160;21&#160;13:22&#160;UTC | NetNumber SHAKEN Root Intermediate CA 1 | 26&#160;Sep&#160;33&#160;13:22&#160;UTC | true | [view](CERTS/e449803766edf02ab50b034dd7e89e54efd332cce87688a032f89b340d039878/README.md) |


Generated: 02 Jun 23 01:12 UTC