# STIR/SHAKEN CA Ecosystem Compliance

## Comcast

### Summary

\* The percent of errors, warnings and notices is calculated against total observed unexpired and trusted certificates from the specified issuer.\
\*\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran.

#### Leaf Certificates

- 30 certificates were included in the corpus being tested
- 0 certificates in the corpus were skipped because they are duplicates
- 5 certificates in the corpus were skipped because they are expired
- 0 certificates in the corpus were skipped because they are not currently trusted
- 25 certificates being tested against the remaining rules
- 2.00 issues on average found in unexpired, trusted, and non-compliant certificates
- 100.00% of certificates contain one or more Error level issue
- 0.00% of certificates contain one or more Warning level issue
- 0.00% of certificates contain one or more Notice level issue
- 0.00% of certificates are too old to be assessed against currently enforced expectations
- 29 days is the average remaining validity for the certificates in the corpus
- 30 days is the average initial validity for the certificates in the corpus
- 25 certificates expire in the next 30 days
- 25.00 average number of unexpired certificates per OCN observed
- 1 unique OCNs observed in unexpired and valid certificate corpus

| Instances | Test | Source |
|-----------|------|--------|
| 25 | [e_atis_ext_crl_distribution_struct](ISSUES/e_atis_ext_crl_distribution_struct/README.md) | ATIS1000080 |
| 25 | [e_atis_serial_number_size](ISSUES/e_atis_serial_number_size/README.md) | ATIS1000080 |

#### CA Certificates

- 2 certificates were included in the corpus being tested
- 0 certificates in the corpus were skipped because they are duplicates
- 0 certificates in the corpus were skipped because they are expired
- 0 certificates in the corpus were skipped because they are not currently trusted
- 2 certificates being tested against the remaining rules
- 0.00 issues on average found in unexpired, trusted, and non-compliant certificates
- 0.00% of certificates contain one or more Error level issue
- 0.00% of certificates contain one or more Warning level issue
- 0.00% of certificates contain one or more Notice level issue
- 100.00% of certificates are too old to be assessed against currently enforced expectations
- 6724 days is the average remaining validity for the certificates in the corpus
- 7150 days is the average initial validity for the certificates in the corpus
- 0 certificates expire in the next 30 days

No error, warning, or notice level issues were found

### Details

#### Leaf Certificates

| Created At | Subject | Not After | Problems | Link |
|------------|---------|-----------|----------|------|
| 14&#160;Jan&#160;24&#160;10:17&#160;UTC | SHAKEN 318J | 13&#160;Feb&#160;24&#160;10:17&#160;UTC | true | [view](CERTS/b082fc47b4ffb10aa4075f60d3cd03be6ae848871243d2c417a5e1d873ff3fc3/README.md) |
| 15&#160;Jan&#160;24&#160;10:17&#160;UTC | SHAKEN 318J | 14&#160;Feb&#160;24&#160;10:17&#160;UTC | true | [view](CERTS/7690ea272adb1c661dd09b8f1f75562d8f02a5b342893cc62091835bdddfcaec/README.md) |
| 16&#160;Jan&#160;24&#160;10:17&#160;UTC | SHAKEN 318J | 15&#160;Feb&#160;24&#160;10:17&#160;UTC | true | [view](CERTS/2077f8bd4ccbf52827581a6bcc761e6d9a4e161dc325bcecda126728d90d3fa9/README.md) |
| 17&#160;Jan&#160;24&#160;10:17&#160;UTC | SHAKEN 318J | 16&#160;Feb&#160;24&#160;10:17&#160;UTC | true | [view](CERTS/a0f500926341bdaed813391e9181173ad5e86e11c07602d90589f91c3d6f31c4/README.md) |
| 18&#160;Jan&#160;24&#160;10:17&#160;UTC | SHAKEN 318J | 17&#160;Feb&#160;24&#160;10:17&#160;UTC | true | [view](CERTS/aaed6f2f09770c1abeb6955d4838f48f63ee34d6203175751073918ad9663171/README.md) |
| 19&#160;Jan&#160;24&#160;10:17&#160;UTC | SHAKEN 318J | 18&#160;Feb&#160;24&#160;10:17&#160;UTC | true | [view](CERTS/8b534b0e09eee5f8b20bb6781e1e4875af23044bfc31ee450d5634d44d54171b/README.md) |
| 20&#160;Jan&#160;24&#160;10:17&#160;UTC | SHAKEN 318J | 19&#160;Feb&#160;24&#160;10:17&#160;UTC | true | [view](CERTS/158d2e35252a7783593644a9776ef9564cfbd74361b84ff62abaac09ff04207c/README.md) |
| 21&#160;Jan&#160;24&#160;10:17&#160;UTC | SHAKEN 318J | 20&#160;Feb&#160;24&#160;10:17&#160;UTC | true | [view](CERTS/bb5d594a41b634fe31b3abd511f4d9600c2ba464debbab8bd26928fcb5dd8c58/README.md) |
| 22&#160;Jan&#160;24&#160;10:17&#160;UTC | SHAKEN 318J | 21&#160;Feb&#160;24&#160;10:17&#160;UTC | true | [view](CERTS/aa049910ad85b3bfd7cab7aae5efcc19fa0718dfbde676819bdc9ce83a56f3bd/README.md) |
| 23&#160;Jan&#160;24&#160;10:17&#160;UTC | SHAKEN 318J | 22&#160;Feb&#160;24&#160;10:17&#160;UTC | true | [view](CERTS/628f5a72cb09e7f7756830bf78f913cbbd32b5c82b0acde27018624e2ffa5fa5/README.md) |
| 24&#160;Jan&#160;24&#160;10:17&#160;UTC | SHAKEN 318J | 23&#160;Feb&#160;24&#160;10:17&#160;UTC | true | [view](CERTS/3216d6bf30b710bb99b84cc7beec5ed6e88c8c7f300414e05093bae50359ec14/README.md) |
| 25&#160;Jan&#160;24&#160;10:17&#160;UTC | SHAKEN 318J | 24&#160;Feb&#160;24&#160;10:17&#160;UTC | true | [view](CERTS/fefb9b90faaf8c57928930c1596a6db407a97a7aebab6031419b738946a9b34e/README.md) |
| 26&#160;Jan&#160;24&#160;10:17&#160;UTC | SHAKEN 318J | 25&#160;Feb&#160;24&#160;10:17&#160;UTC | true | [view](CERTS/6a34dccaf071bc4775132c3bfd6269be5bd555c844cc7f912bba1aace19cac85/README.md) |
| 27&#160;Jan&#160;24&#160;10:17&#160;UTC | SHAKEN 318J | 26&#160;Feb&#160;24&#160;10:17&#160;UTC | true | [view](CERTS/160c4fd2444f07506a5488c7d61bff59e64d839d7f5a8213e5741e3830e6f7f2/README.md) |
| 28&#160;Jan&#160;24&#160;10:17&#160;UTC | SHAKEN 318J | 27&#160;Feb&#160;24&#160;10:17&#160;UTC | true | [view](CERTS/8768f5364f07ead041b0721de61f071ffc1ff5e3a0fdcb22dc5cd7cd5b2a76a9/README.md) |
| 29&#160;Jan&#160;24&#160;10:17&#160;UTC | SHAKEN 318J | 28&#160;Feb&#160;24&#160;10:17&#160;UTC | true | [view](CERTS/cb7badd234f97c05f49519d8a1ee4d463f0a8d458717a7f3d2ef4a4084a4c3b8/README.md) |
| 30&#160;Jan&#160;24&#160;10:17&#160;UTC | SHAKEN 318J | 29&#160;Feb&#160;24&#160;10:17&#160;UTC | true | [view](CERTS/165993ed23c71f3a02a77f75d72a8335b34039038b82656a2bee7bb01167b73b/README.md) |
| 31&#160;Jan&#160;24&#160;10:17&#160;UTC | SHAKEN 318J | 01&#160;Mar&#160;24&#160;10:17&#160;UTC | true | [view](CERTS/9f7fc1299969151b1861d8478fa1ba800e04c5411d93fbb9c42389328820d694/README.md) |
| 01&#160;Feb&#160;24&#160;10:17&#160;UTC | SHAKEN 318J | 02&#160;Mar&#160;24&#160;10:17&#160;UTC | true | [view](CERTS/a1148584470b1dee13d002c7664935fd9fc6dbc8fe65ccbd1196965f56bce75c/README.md) |
| 02&#160;Feb&#160;24&#160;10:17&#160;UTC | SHAKEN 318J | 03&#160;Mar&#160;24&#160;10:17&#160;UTC | true | [view](CERTS/59099357a7ed8b45e3d055445d58bedcbbb8857942759db501f640d3fb5caf4c/README.md) |
| 03&#160;Feb&#160;24&#160;10:17&#160;UTC | SHAKEN 318J | 04&#160;Mar&#160;24&#160;10:17&#160;UTC | true | [view](CERTS/6bcb493f674a977f34a9489f1bb4634bc61e4b9ea4dc9907b0463ada38152d1c/README.md) |
| 04&#160;Feb&#160;24&#160;10:17&#160;UTC | SHAKEN 318J | 05&#160;Mar&#160;24&#160;10:17&#160;UTC | true | [view](CERTS/daa93e4b4c191a2ba7d3114e8d91efe87f6396fc6c1f957ae0630f48b3881830/README.md) |
| 05&#160;Feb&#160;24&#160;10:17&#160;UTC | SHAKEN 318J | 06&#160;Mar&#160;24&#160;10:17&#160;UTC | true | [view](CERTS/bf29d7f48d0a50e5daf1c3e3c1d7a2308c7631fb3faeac926cad9428340c518e/README.md) |
| 06&#160;Feb&#160;24&#160;10:17&#160;UTC | SHAKEN 318J | 07&#160;Mar&#160;24&#160;10:17&#160;UTC | true | [view](CERTS/0f2a493e65fa391ecfab2ebd134f995195716c3930aa35d01a135503691600d8/README.md) |
| 07&#160;Feb&#160;24&#160;10:17&#160;UTC | SHAKEN 318J | 08&#160;Mar&#160;24&#160;10:17&#160;UTC | true | [view](CERTS/3859aa11bba76d88f7c85259b4d89e8a0838386e066e245d894118d14850fd44/README.md) |

#### CA Certificates

| Created At | Subject | Not After | Problems | Link |
|------------|---------|-----------|----------|------|
| 17&#160;Mar&#160;20&#160;14:45&#160;UTC | Comcast SHAKEN Root CA | 12&#160;Mar&#160;40&#160;14:45&#160;UTC | false | [view](CERTS/b1132c5f12c3ca4d2ff119f2df99544336eb1703512ac99cc42d596e25125bbd/README.md) |
| 06&#160;Apr&#160;20&#160;13:48&#160;UTC | Comcast SHAKEN Intermediate CA | 06&#160;Jun&#160;39&#160;13:48&#160;UTC | false | [view](CERTS/2f3abdfe711377f4d59f31d941962797c61f021c3924af31de99031f9ac54f77/README.md) |


Generated: 12 Feb 24 19:26 UTC