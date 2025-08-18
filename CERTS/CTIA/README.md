# STIR/SHAKEN CA Ecosystem Compliance

## CTIA

### Summary

\* The percent of errors, warnings and notices is calculated against total observed unexpired and trusted certificates from the specified issuer.\
\*\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran.

#### Leaf Certificates

- 2 certificates were included in the corpus being tested
- 0 certificates in the corpus were skipped because they are duplicates
- 0 certificates in the corpus were skipped because they are expired
- 0 certificates in the corpus were skipped because they are not currently trusted
- 2 certificates being tested against the remaining rules
- 0.00 issues on average found in unexpired, trusted, and non-compliant certificates
- 0.00% of certificates contain one or more Error level issue
- 0.00% of certificates contain one or more Warning level issue
- 0.00% of certificates contain one or more Notice level issue
- 0.00% of certificates are too old to be assessed against currently enforced expectations
- 301 days is the average remaining validity for the certificates in the corpus
- 366 days is the average initial validity for the certificates in the corpus
- 0 certificates expire in the next 30 days
- 1.00 average number of unexpired certificates per OCN observed
- 2 unique OCNs observed in unexpired and valid certificate corpus

No error, warning, or notice level issues were found

#### CA Certificates

- 2 certificates were included in the corpus being tested
- 0 certificates in the corpus were skipped because they are duplicates
- 0 certificates in the corpus were skipped because they are expired
- 0 certificates in the corpus were skipped because they are not currently trusted
- 2 certificates being tested against the remaining rules
- 1.00 issues on average found in unexpired, trusted, and non-compliant certificates
- 50.00% of certificates contain one or more Error level issue
- 0.00% of certificates contain one or more Warning level issue
- 0.00% of certificates contain one or more Notice level issue
- 0.00% of certificates are too old to be assessed against currently enforced expectations
- 7037 days is the average remaining validity for the certificates in the corpus
- 6388 days is the average initial validity for the certificates in the corpus
- 0 certificates expire in the next 30 days

| Instances | Test | Source |
|-----------|------|--------|
| 1 | [e_atis_subject_public_key_ca](ISSUES/e_atis_subject_public_key_ca/README.md) | ATIS1000080 |

### Details

#### Leaf Certificates

| Created At | Subject | Not After | Problems | Link |
|------------|---------|-----------|----------|------|
| 06&#160;Jan&#160;25&#160;20:18&#160;UTC | SHAKEN 583J | 07&#160;Jan&#160;26&#160;21:31&#160;UTC | false | [view](CERTS/e150327e7954833b2ed28e4388e0c1287b24175366c4af33b669f9087d8dfce5/README.md) |
| 03&#160;Feb&#160;25&#160;18:47&#160;UTC | SHAKEN 5807 | 03&#160;Feb&#160;26&#160;19:47&#160;UTC | false | [view](CERTS/f412ed04b36d888c837d018c98e98fab60cb1d9f343a50747d4d3aef1f3bd25b/README.md) |

#### CA Certificates

| Created At | Subject | Not After | Problems | Link |
|------------|---------|-----------|----------|------|
| 21&#160;Jun&#160;23&#160;13:15&#160;UTC | SHAKEN Root CA | 14&#160;Jun&#160;48&#160;13:15&#160;UTC | true | [view](CERTS/ee1cf83becad4777dcf250170efecc7fc7498d85097e9a570dfac542151ebf53/README.md) |
| 13&#160;Nov&#160;24&#160;13:51&#160;UTC | SHAKEN Intermediate CA 4 | 11&#160;Nov&#160;34&#160;13:51&#160;UTC | false | [view](CERTS/152da8c7466f18a068961de3058c34efd4eb23e68895aa79a15eb7fec88bda6f/README.md) |


Generated: 18 Aug 25 21:13 UTC